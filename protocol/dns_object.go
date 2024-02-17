package protocol

import (
	"fmt"

	"github.com/isi-lincoln/avoid/pkg"
)

func (x *DNSEntry) Key() string {
	return fmt.Sprintf("%s/%s/%s", pkg.DNSEntryPrefix, x.Id, x.Recordtype)
}

func (x *DNSEntry) SetVersion(v int64) { x.Version = v }

func (x *DNSEntry) Value() interface{} { return x }
