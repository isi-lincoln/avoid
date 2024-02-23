package pkg

import (
	"fmt"
	"net/netip"

	avoid "github.com/isi-lincoln/avoid/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ToEndpointAddr(server string, port int) string {
	return fmt.Sprintf("%s:%d", server, port)
}

func WithAvoidDNS(endpoint string, f func(avoid.AVOIDDNSClient) error) error {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to moa service: %v", err)
	}
	client := avoid.NewAVOIDDNSClient(conn)
	defer conn.Close()

	return f(client)
}

func CheckDNSRecord(record *avoid.DNSEntry) error {

	log.Infof("In check record\n")

	for _, ip := range record.Arecords {
		addr, err := netip.ParseAddr(ip)
		if err != nil {
			return fmt.Errorf("Bad address provided %s: %v", ip, err)
		}
		if !addr.Is4() {
			return fmt.Errorf("Address is not ipv4 but record type is A: %s", ip)
		}
	}

	for _, ip := range record.Aaaarecords {
		addr, err := netip.ParseAddr(ip)
		if err != nil {
			return fmt.Errorf("Bad address provided %s: %v", ip, err)
		}
		if !addr.Is6() {
			return fmt.Errorf("Address is not ipv6 but record type is AAAA: %v", ip)
		}
	}

	if len(record.Arecords) <= 0 && len(record.Aaaarecords) <= 0 {
		return fmt.Errorf("DNS Entry does not have any values")
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
		log.Warnf("%s/%s: should use a TXT field less than 200 bytes: %s", record.Ue, record.Name, record.Txt)
	}

	log.Infof("check record okay\n")

	return nil
}
