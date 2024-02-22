// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Identification: the mechanism of pinning a UE to a DNS request
// There are probably better ways of achieving this other than through
// IP, if the UE is able to speak the protocol and can provide information
// such as ISMI, or other UUIDs that can 1:1 map the UE.
type Identification struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	Ismi                 string   `protobuf:"bytes,4,opt,name=ismi,proto3" json:"ismi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identification) Reset()         { *m = Identification{} }
func (m *Identification) String() string { return proto.CompactTextString(m) }
func (*Identification) ProtoMessage()    {}
func (*Identification) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{0}
}

func (m *Identification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identification.Unmarshal(m, b)
}
func (m *Identification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identification.Marshal(b, m, deterministic)
}
func (m *Identification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identification.Merge(m, src)
}
func (m *Identification) XXX_Size() int {
	return xxx_messageInfo_Identification.Size(m)
}
func (m *Identification) XXX_DiscardUnknown() {
	xxx_messageInfo_Identification.DiscardUnknown(m)
}

var xxx_messageInfo_Identification proto.InternalMessageInfo

func (m *Identification) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Identification) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Identification) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

func (m *Identification) GetIsmi() string {
	if m != nil {
		return m.Ismi
	}
	return ""
}

type Record struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Rrtype               int64    `protobuf:"varint,2,opt,name=Rrtype,proto3" json:"Rrtype,omitempty"`
	Class                int64    `protobuf:"varint,3,opt,name=Class,proto3" json:"Class,omitempty"`
	Ttl                  int64    `protobuf:"varint,4,opt,name=Ttl,proto3" json:"Ttl,omitempty"`
	Rdlength             int64    `protobuf:"varint,5,opt,name=Rdlength,proto3" json:"Rdlength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetRrtype() int64 {
	if m != nil {
		return m.Rrtype
	}
	return 0
}

func (m *Record) GetClass() int64 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *Record) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Record) GetRdlength() int64 {
	if m != nil {
		return m.Rdlength
	}
	return 0
}

type DNSEntry struct {
	Ue                   string   `protobuf:"bytes,1,opt,name=ue,proto3" json:"ue,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Arecords             []string `protobuf:"bytes,3,rep,name=arecords,proto3" json:"arecords,omitempty"`
	Aaaarecords          []string `protobuf:"bytes,4,rep,name=aaaarecords,proto3" json:"aaaarecords,omitempty"`
	Ttl                  int64    `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Txt                  string   `protobuf:"bytes,6,opt,name=txt,proto3" json:"txt,omitempty"`
	Version              int64    `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSEntry) Reset()         { *m = DNSEntry{} }
func (m *DNSEntry) String() string { return proto.CompactTextString(m) }
func (*DNSEntry) ProtoMessage()    {}
func (*DNSEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{2}
}

func (m *DNSEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSEntry.Unmarshal(m, b)
}
func (m *DNSEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSEntry.Marshal(b, m, deterministic)
}
func (m *DNSEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSEntry.Merge(m, src)
}
func (m *DNSEntry) XXX_Size() int {
	return xxx_messageInfo_DNSEntry.Size(m)
}
func (m *DNSEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DNSEntry proto.InternalMessageInfo

func (m *DNSEntry) GetUe() string {
	if m != nil {
		return m.Ue
	}
	return ""
}

func (m *DNSEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DNSEntry) GetArecords() []string {
	if m != nil {
		return m.Arecords
	}
	return nil
}

func (m *DNSEntry) GetAaaarecords() []string {
	if m != nil {
		return m.Aaaarecords
	}
	return nil
}

func (m *DNSEntry) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *DNSEntry) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

func (m *DNSEntry) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type EntryRequest struct {
	Entries              []*DNSEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EntryRequest) Reset()         { *m = EntryRequest{} }
func (m *EntryRequest) String() string { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()    {}
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{3}
}

func (m *EntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryRequest.Unmarshal(m, b)
}
func (m *EntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryRequest.Marshal(b, m, deterministic)
}
func (m *EntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryRequest.Merge(m, src)
}
func (m *EntryRequest) XXX_Size() int {
	return xxx_messageInfo_EntryRequest.Size(m)
}
func (m *EntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntryRequest proto.InternalMessageInfo

func (m *EntryRequest) GetEntries() []*DNSEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type EntryResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Code                 int64    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryResponse) Reset()         { *m = EntryResponse{} }
func (m *EntryResponse) String() string { return proto.CompactTextString(m) }
func (*EntryResponse) ProtoMessage()    {}
func (*EntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{4}
}

func (m *EntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryResponse.Unmarshal(m, b)
}
func (m *EntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryResponse.Marshal(b, m, deterministic)
}
func (m *EntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryResponse.Merge(m, src)
}
func (m *EntryResponse) XXX_Size() int {
	return xxx_messageInfo_EntryResponse.Size(m)
}
func (m *EntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntryResponse proto.InternalMessageInfo

func (m *EntryResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *EntryResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{6}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ShowRequest struct {
	Ue                   string   `protobuf:"bytes,1,opt,name=ue,proto3" json:"ue,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowRequest) Reset()         { *m = ShowRequest{} }
func (m *ShowRequest) String() string { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()    {}
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{7}
}

func (m *ShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRequest.Unmarshal(m, b)
}
func (m *ShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRequest.Marshal(b, m, deterministic)
}
func (m *ShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRequest.Merge(m, src)
}
func (m *ShowRequest) XXX_Size() int {
	return xxx_messageInfo_ShowRequest.Size(m)
}
func (m *ShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRequest proto.InternalMessageInfo

func (m *ShowRequest) GetUe() string {
	if m != nil {
		return m.Ue
	}
	return ""
}

func (m *ShowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ShowResponse struct {
	Entry                *DNSEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{8}
}

func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (m *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(m, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetEntry() *DNSEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type ClearRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearRequest) Reset()         { *m = ClearRequest{} }
func (m *ClearRequest) String() string { return proto.CompactTextString(m) }
func (*ClearRequest) ProtoMessage()    {}
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_638ff8d8aaf3d8ae, []int{9}
}

func (m *ClearRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearRequest.Unmarshal(m, b)
}
func (m *ClearRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearRequest.Marshal(b, m, deterministic)
}
func (m *ClearRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearRequest.Merge(m, src)
}
func (m *ClearRequest) XXX_Size() int {
	return xxx_messageInfo_ClearRequest.Size(m)
}
func (m *ClearRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Identification)(nil), "protocol.Identification")
	proto.RegisterType((*Record)(nil), "protocol.Record")
	proto.RegisterType((*DNSEntry)(nil), "protocol.DNSEntry")
	proto.RegisterType((*EntryRequest)(nil), "protocol.EntryRequest")
	proto.RegisterType((*EntryResponse)(nil), "protocol.EntryResponse")
	proto.RegisterType((*ListRequest)(nil), "protocol.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "protocol.ListResponse")
	proto.RegisterType((*ShowRequest)(nil), "protocol.ShowRequest")
	proto.RegisterType((*ShowResponse)(nil), "protocol.ShowResponse")
	proto.RegisterType((*ClearRequest)(nil), "protocol.ClearRequest")
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor_638ff8d8aaf3d8ae) }

var fileDescriptor_638ff8d8aaf3d8ae = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8f, 0x12, 0x31,
	0x14, 0x5e, 0x18, 0x18, 0xe0, 0xc1, 0x12, 0xd3, 0x28, 0x4e, 0x38, 0x91, 0x26, 0x46, 0x0e, 0x0a,
	0x71, 0x35, 0xab, 0x51, 0x13, 0x13, 0x17, 0x0f, 0x26, 0x86, 0x43, 0xd1, 0x8b, 0x27, 0xbb, 0x33,
	0x75, 0x69, 0x1c, 0xda, 0xb1, 0xed, 0x20, 0xfc, 0x25, 0x13, 0xff, 0xa3, 0x69, 0xa7, 0x33, 0x34,
	0x6a, 0x8c, 0xf1, 0xc4, 0xf7, 0xbe, 0xbe, 0xef, 0xbd, 0xef, 0x7d, 0x64, 0x60, 0x90, 0x09, 0xbd,
	0x28, 0x94, 0x34, 0x12, 0xf5, 0xdd, 0x4f, 0x2a, 0x73, 0xfc, 0x09, 0xc6, 0x6f, 0x33, 0x26, 0x0c,
	0xff, 0xcc, 0x53, 0x6a, 0xb8, 0x14, 0x08, 0x41, 0xa7, 0x2c, 0x79, 0x96, 0xb4, 0x66, 0xad, 0xf9,
	0x80, 0x38, 0x6c, 0x39, 0x5e, 0xec, 0x9f, 0x24, 0xed, 0x8a, 0xb3, 0xd8, 0x73, 0x97, 0x49, 0xd4,
	0x70, 0x97, 0x8e, 0xd3, 0x3b, 0x9e, 0x74, 0x3c, 0xa7, 0x77, 0x1c, 0x1f, 0x20, 0x26, 0x2c, 0x95,
	0xca, 0x4d, 0x59, 0xd3, 0x1d, 0xab, 0x27, 0x5b, 0x8c, 0x26, 0x10, 0x13, 0x65, 0x8e, 0x05, 0x73,
	0xb3, 0x23, 0xe2, 0x2b, 0x74, 0x1b, 0xba, 0x57, 0x39, 0xd5, 0xda, 0x8d, 0x8f, 0x48, 0x55, 0xa0,
	0x5b, 0x10, 0xbd, 0x37, 0xb9, 0x1b, 0x1f, 0x11, 0x0b, 0xd1, 0x14, 0xfa, 0x24, 0xcb, 0x99, 0xb8,
	0x31, 0xdb, 0xa4, 0xeb, 0xe8, 0xa6, 0xc6, 0xdf, 0x5b, 0xd0, 0x5f, 0xad, 0x37, 0x6f, 0x84, 0x51,
	0x47, 0x34, 0x86, 0x76, 0x59, 0xaf, 0x6e, 0x97, 0xcc, 0x9a, 0x11, 0xd6, 0x8c, 0x3f, 0xc9, 0x62,
	0x3b, 0x8c, 0x2a, 0xe7, 0xd5, 0xee, 0x8d, 0xe6, 0x03, 0xd2, 0xd4, 0x68, 0x06, 0x43, 0x4a, 0x69,
	0xf3, 0xdc, 0x71, 0xcf, 0x21, 0x65, 0xcd, 0x19, 0x93, 0x7b, 0x17, 0x16, 0x3a, 0xe6, 0x60, 0x92,
	0xd8, 0xad, 0xb0, 0x10, 0x25, 0xd0, 0xdb, 0x33, 0xa5, 0xb9, 0x14, 0x49, 0xcf, 0xf5, 0xd5, 0x25,
	0x7e, 0x09, 0x23, 0x67, 0x94, 0xb0, 0xaf, 0x25, 0xd3, 0x06, 0x3d, 0x80, 0x1e, 0x13, 0x46, 0x71,
	0xa6, 0x93, 0xd6, 0x2c, 0x9a, 0x0f, 0x2f, 0xd0, 0xa2, 0xfe, 0xd3, 0x16, 0xf5, 0x51, 0xa4, 0x6e,
	0xc1, 0xaf, 0xe0, 0xdc, 0xab, 0x75, 0x21, 0x85, 0x76, 0xa7, 0x28, 0x8f, 0xfd, 0xd1, 0x4d, 0x6d,
	0x4f, 0x4f, 0x65, 0x56, 0x27, 0xee, 0x30, 0x3e, 0x87, 0xe1, 0x3b, 0xae, 0x8d, 0xdf, 0x8e, 0x31,
	0x8c, 0xaa, 0xf2, 0x24, 0xf9, 0xc2, 0x8e, 0x95, 0x95, 0x01, 0x71, 0x18, 0x3f, 0x82, 0xe1, 0x66,
	0x2b, 0xbf, 0xd5, 0x86, 0xff, 0x21, 0x60, 0xfc, 0x0c, 0x46, 0x95, 0xc4, 0x8f, 0x9d, 0x43, 0xd7,
	0x5e, 0x70, 0x74, 0xb2, 0x3f, 0x9f, 0x58, 0x35, 0xe0, 0x31, 0x8c, 0xae, 0x72, 0x46, 0x95, 0xdf,
	0x76, 0xf1, 0xa3, 0x0d, 0xd1, 0x6a, 0xbd, 0x41, 0x2f, 0x20, 0xfe, 0x50, 0x64, 0xd4, 0x30, 0x34,
	0x39, 0x89, 0xc3, 0x20, 0xa7, 0x77, 0x7f, 0xe3, 0xab, 0xe5, 0xf8, 0xcc, 0x8a, 0x57, 0x2c, 0x67,
	0xff, 0x27, 0x7e, 0x0a, 0x1d, 0x1b, 0x11, 0xba, 0x73, 0x6a, 0x09, 0x12, 0x9c, 0x4e, 0x7e, 0xa5,
	0x43, 0xa1, 0x0d, 0x21, 0x14, 0x06, 0x39, 0x86, 0xc2, 0x30, 0x2b, 0x7c, 0x86, 0x9e, 0xdb, 0x6f,
	0x82, 0x51, 0x15, 0xba, 0x0d, 0x43, 0xf9, 0x8b, 0xdb, 0xd7, 0xf7, 0x3f, 0xde, 0xbb, 0xe1, 0x66,
	0x5b, 0x5e, 0x2f, 0x52, 0xb9, 0x5b, 0x72, 0xcd, 0x1f, 0xe6, 0x5c, 0xa4, 0x32, 0x17, 0x4b, 0xba,
	0x97, 0x3c, 0x5b, 0xd6, 0xc2, 0xeb, 0xd8, 0xa1, 0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x87,
	0xb1, 0x19, 0x48, 0x2e, 0x04, 0x00, 0x00,
}
