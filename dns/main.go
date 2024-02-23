package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	pkg "github.com/isi-lincoln/avoid/pkg"
	avoid "github.com/isi-lincoln/avoid/protocol"
	"gitlab.com/mergetb/tech/stor"
)

type AVOIDDNSServer struct {
	avoid.UnimplementedAVOIDDNSServer
}

func (s *AVOIDDNSServer) Update(ctx context.Context, req *avoid.EntryRequest) (*avoid.EntryResponse, error) {

	if req == nil {
		errMsg := fmt.Sprintf("Update: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Entries == nil {
		errMsg := fmt.Sprintf("Update: Missing DNS Entry")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if len(req.Entries) < 1 {
		errMsg := fmt.Sprintf("Update called with no values")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	objs := make([]stor.Object, 0)
	for item, entry := range req.Entries {
		log.WithFields(log.Fields{
			"Index":    item,
			"Identity": entry.Ue,
			"Name":     entry.Name,
			"A":        entry.Arecords,
			"AAAA":     entry.Aaaarecords,
		}).Info("Update")

		err := pkg.CheckDNSRecord(entry)
		if err != nil {
			log.Errorf("%s", err)
			return nil, fmt.Errorf("%s", err)
		}

		objs = append(objs, entry)
	}

	err := stor.WriteObjects(objs, false)
	if err != nil {
		return nil, err
	}

	return &avoid.EntryResponse{Response: "", Code: int64(len(req.Entries))}, nil
}

func (s *AVOIDDNSServer) Delete(ctx context.Context, req *avoid.EntryRequest) (*avoid.EntryResponse, error) {
	if req == nil {
		errMsg := fmt.Sprintf("Delete: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Entries == nil {
		errMsg := fmt.Sprintf("Delete: Missing DNS Entry")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if len(req.Entries) < 1 {
		errMsg := fmt.Sprintf("Delete called with no values")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	objs := make([]stor.Object, 0)
	for item, entry := range req.Entries {
		log.WithFields(log.Fields{
			"Index": item,
			"Key":   entry.Key,
		}).Info("Delete")
		objs = append(objs, entry)
	}

	err := stor.DeleteObjects(objs)
	if err != nil {
		return nil, err
	}

	return &avoid.EntryResponse{Response: "", Code: int64(len(req.Entries))}, nil
}

func (s *AVOIDDNSServer) List(ctx context.Context, req *avoid.ListRequest) (*avoid.ListResponse, error) {

	log.Info("List DNS Entry Keys")

	prefix := fmt.Sprintf("%s/", avoid.DNSEntryPrefix)

	keys := make(map[string]string)
	err := stor.WithEtcd(func(c *clientv3.Client) error {
		//TODO:  arbitrary 1 second delay, should use config value
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
		resp, err := c.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
		cancel()
		if err != nil {
			return err
		}

		// keys are in the form: "/prefix/ue/dnsName
		// so we only want the keys portion
		for _, kv := range resp.Kvs {
			keyStr := strings.Split(string(kv.Key), "/")
			if len(keyStr) != 4 {
				log.Warnf("dns key issue: %s\n", kv.Key)
				continue
			}
			// returns ue/dnsName
			key := strings.Join(keyStr[2:], "/")
			keys[key] = ""
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	keyList := make([]string, 0)
	for k, _ := range keys {
		keyList = append(keyList, k)
	}

	return &avoid.ListResponse{
		Keys: keyList,
	}, nil
}

func (s *AVOIDDNSServer) Show(ctx context.Context, req *avoid.ShowRequest) (*avoid.ShowResponse, error) {
	if req == nil {
		errMsg := fmt.Sprintf("Show: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Ue == "" || req.Name == "" {
		errMsg := fmt.Sprintf("Show: Invalid request: {Ue: %s, Name: %s}", req.Ue, req.Name)
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.WithFields(log.Fields{
		"Ue":   req.Ue,
		"Name": req.Name,
	}).Info("Show DNS Item")

	entry := &avoid.DNSEntry{
		Name: req.Name,
		Ue:   req.Ue,
	}

	err := stor.Read(entry)
	if err != nil {
		return nil, err
	}

	return &avoid.ShowResponse{
		Entry: entry,
	}, nil

}

func (s *AVOIDDNSServer) Clear(ctx context.Context, req *avoid.ClearRequest) (*avoid.EntryResponse, error) {
	if req == nil {
		errMsg := fmt.Sprintf("Show: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.Info("Clear Lookup Table ")

	err := stor.WithEtcd(func(c *clientv3.Client) error {
		//TODO:  arbitrary 1 second delay, should use config value
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
		_, err := c.Delete(ctx, fmt.Sprintf("%s/", avoid.DNSEntryPrefix), clientv3.WithPrefix())
		cancel()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &avoid.EntryResponse{
		Response: "cleared",
		Code:     0,
	}, nil
}

func main() {

	var debug bool
	var port int
	var addr string
	var etcdHost string
	var etcdPort int

	flag.IntVar(&port, "port", pkg.DefaultAvoidDNSPort, "set the Avoid DNS Control Port")
	flag.StringVar(&addr, "addr", "0.0.0.0", "set the Avoid DNS Control Address")
	flag.BoolVar(&debug, "debug", false, "enable extra debug logging")

	// these values will only work/matter if no tls has been enabled
	flag.IntVar(&etcdPort, "etcdport", pkg.DefaultEtcdPort, "set the etcd backend port")
	flag.StringVar(&etcdHost, "etcdhost", pkg.DefaultEtcdHost, "set the etcd backend host")

	defaultPortEnv := pkg.DefaultAvoidDNSPortENV
	defaultAddrEnv := pkg.DefaultAvoidDNSAddrENV

	// Presendence: ENV > flags > defaults

	// If we have environment variables, load them in
	portStr := os.Getenv(defaultPortEnv)
	if portStr != "" {
		portInt, err := strconv.Atoi(portStr)
		if err != nil {
			log.Warningf("Failed to convert %s to int, ignored: %v", defaultPortEnv, portStr)
		} else {
			port = portInt
		}
	}

	if os.Getenv(defaultAddrEnv) != "" {
		addr = os.Getenv(defaultAddrEnv)
	}

	debugStr := os.Getenv("DEBUG")
	if debugStr != "" {
		debug = true
	}

	// set debug level logging
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// if we have a configuration file load it
	cfg, err := pkg.LoadConfig(pkg.DefaultConfigPath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// if the configuration file has environment settings, set them
	err = pkg.ReadENVSettings(cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Debugf("etcd values: %s:%d", etcdHost, etcdPort)

	// configure the backend database settings
	etcdCfg, err := pkg.SetEtcdSettings(cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	stor.SetConfig(*etcdCfg)

	netAddr, err := netip.ParseAddr(addr)
	if err != nil {
		log.Errorf("Bad address provided %s: %v", addr, err)
		return
	}

	if !netAddr.Is4() && !netAddr.Is6() {
		log.Errorf("Address is not ipv4 or ipv6: %s", addr)
		return
	}

	log.Info(fmt.Sprintf("Avoid dns starting up on %s:%d", addr, port))
	log.Infof("db settings: %+v\n", etcdCfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	avoid.RegisterAVOIDDNSServer(grpcServer, &AVOIDDNSServer{})
	grpcServer.Serve(lis)
}
