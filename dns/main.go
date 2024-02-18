package dns

import (
	"context"
	"encoding/json"
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

func checkRecord(record *avoid.DNSEntry) error {
	// TODO: other methods of identification
	switch record.Recordtype {
	case "A":
		for _, ip := range record.Records {
			addr, err := netip.ParseAddr(ip)
			if err != nil {
				return fmt.Errorf("Bad address provided %s: %v", ip, err)
			}
			if !addr.Is4() {
				return fmt.Errorf("Address is not ipv4 but record type is %s: %s", ip, record.Recordtype)
			}
		}
		break
	case "AAAA":
		for _, ip := range record.Records {
			addr, err := netip.ParseAddr(ip)
			if err != nil {
				return fmt.Errorf("Bad address provided %s: %v", ip, err)
			}
			if !addr.Is6() {
				return fmt.Errorf("Address is not ipv6 but record type is %s: %v", ip, record.Recordtype)
			}
		}
		break
	default:
		errMsg := fmt.Errorf("Unknown record type: %s", record.Recordtype)
		log.Errorf("%v", errMsg)
		return errMsg
	}

	if len(record.Records) <= 0 {
		return fmt.Errorf("DNS Entry does not have a value- delete instead")
	}

	// RFC 2181
	if record.Ttl < 0 || record.Ttl >= (2<<31)-1 {
		return fmt.Errorf("DNS Entry TTL must be between 0 and 2^31-1")
	}

	// RFC 6763
	if len(record.Txt) == 0 || len(record.Txt) >= 65535 {
		return fmt.Errorf("DNS TXT field must be between 1 and 65535")
	}

	if len(record.Txt) > 200 {
		log.Warnf("%s: should use a TXT field less than 200 bytes: %s", record.Id, record.Txt)
	}

	return nil
}

func update(entry *avoid.DNSEntry) error {
	err := checkRecord(entry)
	if err != nil {
		log.Errorf("%s", err)
		return fmt.Errorf("%s", err)
	}

	objs := stor.Object(entry)

	err = stor.Write(objs, false)
	if err != nil {
		return err
	}

	return nil
}

type DNSServer struct {
	avoid.UnimplementedDNSServer
}

func (s *DNSServer) Update(ctx context.Context, req *avoid.EntryRequest) (*avoid.EntryResponse, error) {

	if req == nil {
		errMsg := fmt.Sprintf("Update: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Entry == nil {
		errMsg := fmt.Sprintf("Update: Missing DNS Entry")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.WithFields(log.Fields{
		"Identity": req.Entry.Id,
		"Type":     req.Entry.Recordtype,
		"Records":  req.Entry.Records,
	}).Info("Update")

	err := update(req.Entry)
	if err != nil {
		return nil, err
	}

	return &avoid.EntryResponse{Response: "", Code: 0}, nil
}

func (s *DNSServer) Delete(ctx context.Context, req *avoid.EntryRequest) (*avoid.EntryResponse, error) {
	if req == nil {
		errMsg := fmt.Sprintf("Delete: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Entry == nil {
		errMsg := fmt.Sprintf("Delete: Missing DNS Entry")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.WithFields(log.Fields{
		"Key": req.Entry.Key,
	}).Info("Delete")

	objs := stor.Object(req.Entry)

	err := stor.Delete(objs)
	if err != nil {
		return nil, err
	}

	return &avoid.EntryResponse{Response: "", Code: 0}, nil
}

func (s *DNSServer) List(ctx context.Context, req *avoid.ListRequest) (*avoid.ListResponse, error) {

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

		// keys are in the form: "/prefix/key/record
		// so we only want the keys portion
		for _, kv := range resp.Kvs {
			keyStr := strings.Split(string(kv.Key), "/")
			if len(keyStr) != 4 {
				log.Warnf("dns key issue: %s\n", keyStr)
				continue
			}
			key := keyStr[2]
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

func (s *DNSServer) Show(ctx context.Context, req *avoid.ShowRequest) (*avoid.ShowResponse, error) {
	if req == nil {
		errMsg := fmt.Sprintf("Show: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Key == "" {
		errMsg := fmt.Sprintf("Show: Empty Key")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.WithFields(log.Fields{
		"Key": req.Key,
	}).Info("Show DNS Item")

	records := make([]*avoid.DNSEntry, 0)
	err := stor.WithEtcd(func(c *clientv3.Client) error {
		//TODO:  arbitrary 1 second delay, should use config value
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
		resp, err := c.Get(ctx, fmt.Sprintf("%s/%s", avoid.DNSEntryPrefix, req.Key), clientv3.WithPrefix())
		cancel()
		if err != nil {
			return err
		}

		// keys are in the form: "/prefix/key/record
		// so we only want the keys portion
		for _, kv := range resp.Kvs {
			tmp := &avoid.DNSEntry{}
			err = json.Unmarshal(kv.Value, &tmp)
			if err != nil {
				return err
			}

			records = append(records, tmp)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &avoid.ShowResponse{
		Entries: records,
	}, nil
}

func (s *DNSServer) Clear(ctx context.Context, req *avoid.ClearRequest) (*avoid.EntryResponse, error) {
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

func (s *DNSServer) Reload(ctx context.Context, req *avoid.ReloadRequest) (*avoid.ReloadResponse, error) {

	if req == nil {
		errMsg := fmt.Sprintf("Reload: Malformed Request")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	if req.Entries == nil {
		errMsg := fmt.Sprintf("Reload: Missing DNS Entry")
		log.Errorf("%s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	log.Info("Updating with %d entries", len(req.Entries))

	resp := make([]*avoid.EntryResponse, 0)
	for _, entry := range req.Entries {
		var err error = update(entry.Entry)
		var msg string = ""
		var code int64 = 0
		if err != nil {
			msg = fmt.Sprintf("%v", err)
			code = 1
		}
		resp = append(resp, &avoid.EntryResponse{Response: msg, Code: code})
	}

	return &avoid.ReloadResponse{Responses: resp}, nil
}

func main() {

	var debug bool
	var port int
	var addr string

	flag.IntVar(&port, "port", pkg.DefaultAvoidDNSPort, "set the Avoid DNS Control Port")
	flag.StringVar(&addr, "addr", "0.0.0.0", "set the Avoid DNS Control Address")
	flag.BoolVar(&debug, "debug", false, "enable extra debug logging")

	defaultPortEnv := pkg.DefaultAvoidDNSPortENV
	defaultAddrEnv := pkg.DefaultAvoidDNSAddrENV

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

	if !netAddr.Is4() || !netAddr.Is6() {
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
	avoid.RegisterDNSServer(grpcServer, &DNSServer{})
	grpcServer.Serve(lis)
}
