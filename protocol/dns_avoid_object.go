package protocol

import (
	"fmt"
)

var (
	DNSEntryPrefix = "/dns"
)

func (x *DNSEntry) Key() string {
	return fmt.Sprintf("%s/%s/%s", DNSEntryPrefix, x.Ue, x.Name)
}

func (x *DNSEntry) SetVersion(v int64) { x.Version = v }

func (x *DNSEntry) Value() interface{} { return x }
