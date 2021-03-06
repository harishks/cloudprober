// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/config.proto

/*
Package probes is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/config.proto

It has these top-level messages:
	ProbeDef
*/
package probes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_probes_http "github.com/google/cloudprober/probes/http"
import cloudprober_probes_dns "github.com/google/cloudprober/probes/dns"
import cloudprober_probes_external "github.com/google/cloudprober/probes/external"
import cloudprober_probes_ping "github.com/google/cloudprober/probes/ping"
import cloudprober_probes_udp "github.com/google/cloudprober/probes/udp"
import cloudprober_targets "github.com/google/cloudprober/targets"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProbeDef_Type int32

const (
	ProbeDef_PING         ProbeDef_Type = 0
	ProbeDef_HTTP         ProbeDef_Type = 1
	ProbeDef_DNS          ProbeDef_Type = 2
	ProbeDef_EXTERNAL     ProbeDef_Type = 3
	ProbeDef_UDP          ProbeDef_Type = 4
	ProbeDef_USER_DEFINED ProbeDef_Type = 5
)

var ProbeDef_Type_name = map[int32]string{
	0: "PING",
	1: "HTTP",
	2: "DNS",
	3: "EXTERNAL",
	4: "UDP",
	5: "USER_DEFINED",
}
var ProbeDef_Type_value = map[string]int32{
	"PING":         0,
	"HTTP":         1,
	"DNS":          2,
	"EXTERNAL":     3,
	"UDP":          4,
	"USER_DEFINED": 5,
}

func (x ProbeDef_Type) Enum() *ProbeDef_Type {
	p := new(ProbeDef_Type)
	*p = x
	return p
}
func (x ProbeDef_Type) String() string {
	return proto.EnumName(ProbeDef_Type_name, int32(x))
}
func (x *ProbeDef_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeDef_Type_value, data, "ProbeDef_Type")
	if err != nil {
		return err
	}
	*x = ProbeDef_Type(value)
	return nil
}
func (ProbeDef_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ProbeDef struct {
	Name *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type *ProbeDef_Type `protobuf:"varint,2,req,name=type,enum=cloudprober.probes.ProbeDef_Type" json:"type,omitempty"`
	// Which machines this probe should run on. If defined, cloudprober will run
	// this probe only if machine's hostname matches this value.
	RunOn *string `protobuf:"bytes,3,opt,name=run_on,json=runOn" json:"run_on,omitempty"`
	// Interval between two probes
	IntervalMsec *int32 `protobuf:"varint,4,opt,name=interval_msec,json=intervalMsec,def=2000" json:"interval_msec,omitempty"`
	// Timeout for each probe
	TimeoutMsec *int32 `protobuf:"varint,5,opt,name=timeout_msec,json=timeoutMsec,def=1000" json:"timeout_msec,omitempty"`
	// Targets for the probe
	Targets *cloudprober_targets.TargetsDef `protobuf:"bytes,6,req,name=targets" json:"targets,omitempty"`
	// Types that are valid to be assigned to Probe:
	//	*ProbeDef_PingProbe
	//	*ProbeDef_HttpProbe
	//	*ProbeDef_DnsProbe
	//	*ProbeDef_ExternalProbe
	//	*ProbeDef_UdpProbe
	//	*ProbeDef_UserDefinedProbe
	Probe            isProbeDef_Probe `protobuf_oneof:"probe"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ProbeDef) Reset()                    { *m = ProbeDef{} }
func (m *ProbeDef) String() string            { return proto.CompactTextString(m) }
func (*ProbeDef) ProtoMessage()               {}
func (*ProbeDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeDef_IntervalMsec int32 = 2000
const Default_ProbeDef_TimeoutMsec int32 = 1000

type isProbeDef_Probe interface {
	isProbeDef_Probe()
}

type ProbeDef_PingProbe struct {
	PingProbe *cloudprober_probes_ping.ProbeConf `protobuf:"bytes,20,opt,name=ping_probe,json=pingProbe,oneof"`
}
type ProbeDef_HttpProbe struct {
	HttpProbe *cloudprober_probes_http.ProbeConf `protobuf:"bytes,21,opt,name=http_probe,json=httpProbe,oneof"`
}
type ProbeDef_DnsProbe struct {
	DnsProbe *cloudprober_probes_dns.ProbeConf `protobuf:"bytes,22,opt,name=dns_probe,json=dnsProbe,oneof"`
}
type ProbeDef_ExternalProbe struct {
	ExternalProbe *cloudprober_probes_external.ProbeConf `protobuf:"bytes,23,opt,name=external_probe,json=externalProbe,oneof"`
}
type ProbeDef_UdpProbe struct {
	UdpProbe *cloudprober_probes_udp.ProbeConf `protobuf:"bytes,24,opt,name=udp_probe,json=udpProbe,oneof"`
}
type ProbeDef_UserDefinedProbe struct {
	UserDefinedProbe string `protobuf:"bytes,99,opt,name=user_defined_probe,json=userDefinedProbe,oneof"`
}

func (*ProbeDef_PingProbe) isProbeDef_Probe()        {}
func (*ProbeDef_HttpProbe) isProbeDef_Probe()        {}
func (*ProbeDef_DnsProbe) isProbeDef_Probe()         {}
func (*ProbeDef_ExternalProbe) isProbeDef_Probe()    {}
func (*ProbeDef_UdpProbe) isProbeDef_Probe()         {}
func (*ProbeDef_UserDefinedProbe) isProbeDef_Probe() {}

func (m *ProbeDef) GetProbe() isProbeDef_Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeDef) GetType() ProbeDef_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ProbeDef_PING
}

func (m *ProbeDef) GetRunOn() string {
	if m != nil && m.RunOn != nil {
		return *m.RunOn
	}
	return ""
}

func (m *ProbeDef) GetIntervalMsec() int32 {
	if m != nil && m.IntervalMsec != nil {
		return *m.IntervalMsec
	}
	return Default_ProbeDef_IntervalMsec
}

func (m *ProbeDef) GetTimeoutMsec() int32 {
	if m != nil && m.TimeoutMsec != nil {
		return *m.TimeoutMsec
	}
	return Default_ProbeDef_TimeoutMsec
}

func (m *ProbeDef) GetTargets() *cloudprober_targets.TargetsDef {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ProbeDef) GetPingProbe() *cloudprober_probes_ping.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_PingProbe); ok {
		return x.PingProbe
	}
	return nil
}

func (m *ProbeDef) GetHttpProbe() *cloudprober_probes_http.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_HttpProbe); ok {
		return x.HttpProbe
	}
	return nil
}

func (m *ProbeDef) GetDnsProbe() *cloudprober_probes_dns.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_DnsProbe); ok {
		return x.DnsProbe
	}
	return nil
}

func (m *ProbeDef) GetExternalProbe() *cloudprober_probes_external.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_ExternalProbe); ok {
		return x.ExternalProbe
	}
	return nil
}

func (m *ProbeDef) GetUdpProbe() *cloudprober_probes_udp.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_UdpProbe); ok {
		return x.UdpProbe
	}
	return nil
}

func (m *ProbeDef) GetUserDefinedProbe() string {
	if x, ok := m.GetProbe().(*ProbeDef_UserDefinedProbe); ok {
		return x.UserDefinedProbe
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProbeDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProbeDef_OneofMarshaler, _ProbeDef_OneofUnmarshaler, _ProbeDef_OneofSizer, []interface{}{
		(*ProbeDef_PingProbe)(nil),
		(*ProbeDef_HttpProbe)(nil),
		(*ProbeDef_DnsProbe)(nil),
		(*ProbeDef_ExternalProbe)(nil),
		(*ProbeDef_UdpProbe)(nil),
		(*ProbeDef_UserDefinedProbe)(nil),
	}
}

func _ProbeDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProbeDef)
	// probe
	switch x := m.Probe.(type) {
	case *ProbeDef_PingProbe:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PingProbe); err != nil {
			return err
		}
	case *ProbeDef_HttpProbe:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpProbe); err != nil {
			return err
		}
	case *ProbeDef_DnsProbe:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DnsProbe); err != nil {
			return err
		}
	case *ProbeDef_ExternalProbe:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExternalProbe); err != nil {
			return err
		}
	case *ProbeDef_UdpProbe:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UdpProbe); err != nil {
			return err
		}
	case *ProbeDef_UserDefinedProbe:
		b.EncodeVarint(99<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.UserDefinedProbe)
	case nil:
	default:
		return fmt.Errorf("ProbeDef.Probe has unexpected type %T", x)
	}
	return nil
}

func _ProbeDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProbeDef)
	switch tag {
	case 20: // probe.ping_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_ping.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_PingProbe{msg}
		return true, err
	case 21: // probe.http_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_http.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_HttpProbe{msg}
		return true, err
	case 22: // probe.dns_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_dns.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_DnsProbe{msg}
		return true, err
	case 23: // probe.external_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_external.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_ExternalProbe{msg}
		return true, err
	case 24: // probe.udp_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cloudprober_probes_udp.ProbeConf)
		err := b.DecodeMessage(msg)
		m.Probe = &ProbeDef_UdpProbe{msg}
		return true, err
	case 99: // probe.user_defined_probe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Probe = &ProbeDef_UserDefinedProbe{x}
		return true, err
	default:
		return false, nil
	}
}

func _ProbeDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProbeDef)
	// probe
	switch x := m.Probe.(type) {
	case *ProbeDef_PingProbe:
		s := proto.Size(x.PingProbe)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_HttpProbe:
		s := proto.Size(x.HttpProbe)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_DnsProbe:
		s := proto.Size(x.DnsProbe)
		n += proto.SizeVarint(22<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_ExternalProbe:
		s := proto.Size(x.ExternalProbe)
		n += proto.SizeVarint(23<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_UdpProbe:
		s := proto.Size(x.UdpProbe)
		n += proto.SizeVarint(24<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProbeDef_UserDefinedProbe:
		n += proto.SizeVarint(99<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.UserDefinedProbe)))
		n += len(x.UserDefinedProbe)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ProbeDef)(nil), "cloudprober.probes.ProbeDef")
	proto.RegisterEnum("cloudprober.probes.ProbeDef_Type", ProbeDef_Type_name, ProbeDef_Type_value)
}

func init() { proto.RegisterFile("github.com/google/cloudprober/probes/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd7, 0x36, 0x59, 0xd3, 0xd7, 0x6e, 0x8a, 0x2c, 0x06, 0x51, 0x2f, 0x84, 0x1e, 0x20,
	0x5c, 0xd2, 0xae, 0x68, 0x48, 0xec, 0x04, 0x2c, 0x81, 0x4e, 0x1a, 0x5d, 0x95, 0x65, 0x12, 0xb7,
	0xa8, 0x8b, 0x9d, 0x2c, 0x52, 0x6b, 0x47, 0x89, 0x8d, 0xd8, 0xc7, 0xe5, 0x9b, 0x20, 0x3b, 0x09,
	0x22, 0xd0, 0xa1, 0x8a, 0x93, 0xa3, 0xe7, 0xdf, 0xfb, 0xf5, 0xfd, 0xdd, 0x07, 0xa7, 0x69, 0xc6,
	0xef, 0xc5, 0x9d, 0x1b, 0xb3, 0xed, 0x34, 0x65, 0x2c, 0xdd, 0x90, 0x69, 0xbc, 0x61, 0x02, 0xe7,
	0x05, 0xbb, 0x23, 0xc5, 0x54, 0x1d, 0xe5, 0x34, 0x66, 0x34, 0xc9, 0x52, 0x37, 0x2f, 0x18, 0x67,
	0x08, 0xfd, 0x06, 0xb8, 0x15, 0x30, 0x7e, 0xbb, 0x97, 0xe6, 0x9e, 0xf3, 0xbc, 0xe5, 0x1a, 0x9f,
	0xed, 0xd5, 0x87, 0x69, 0x7b, 0x84, 0xf1, 0xf9, 0x5e, 0x6d, 0xe4, 0x3b, 0x27, 0x05, 0x5d, 0x6f,
	0xda, 0xbd, 0xfb, 0x8d, 0x9a, 0x67, 0x34, 0xfd, 0x9f, 0x51, 0x05, 0xfe, 0x23, 0xe1, 0x9b, 0x7f,
	0xb7, 0xf1, 0x75, 0x91, 0x12, 0x5e, 0x36, 0x67, 0xd5, 0x34, 0xf9, 0xa1, 0x83, 0xb1, 0x92, 0x80,
	0x47, 0x12, 0x84, 0x40, 0xa3, 0xeb, 0x2d, 0xb1, 0x3a, 0x76, 0xd7, 0x19, 0x04, 0xea, 0x1b, 0x9d,
	0x81, 0xc6, 0x1f, 0x72, 0x62, 0x75, 0xed, 0xae, 0x73, 0x3c, 0x7f, 0xe1, 0xfe, 0xfd, 0x97, 0xb8,
	0x4d, 0xbf, 0x1b, 0x3e, 0xe4, 0x24, 0x50, 0x38, 0x3a, 0x81, 0xc3, 0x42, 0xd0, 0x88, 0x51, 0xab,
	0x67, 0x77, 0x9c, 0x41, 0xa0, 0x17, 0x82, 0x5e, 0x53, 0xf4, 0x1a, 0x8e, 0x32, 0xca, 0x49, 0xf1,
	0x6d, 0xbd, 0x89, 0xb6, 0x25, 0x89, 0x2d, 0xcd, 0xee, 0x38, 0xfa, 0xb9, 0x36, 0x9f, 0xcd, 0x66,
	0xc1, 0xa8, 0xb9, 0xfa, 0x52, 0x92, 0x18, 0xbd, 0x82, 0x11, 0xcf, 0xb6, 0x84, 0x09, 0x5e, 0x91,
	0x7a, 0x45, 0x9e, 0x4a, 0x72, 0x58, 0xdf, 0x28, 0xf0, 0x1d, 0xf4, 0xeb, 0x4c, 0xd6, 0xa1, 0xdd,
	0x75, 0x86, 0xf3, 0xe7, 0xad, 0x21, 0x9b, 0xbc, 0x61, 0x75, 0x7a, 0x24, 0x09, 0x1a, 0x1e, 0x5d,
	0x00, 0xc8, 0xe7, 0x8f, 0x14, 0x6a, 0x3d, 0xb1, 0x3b, 0xce, 0x70, 0x3e, 0xd9, 0x15, 0x51, 0x52,
	0x55, 0xce, 0x0b, 0x46, 0x93, 0xc5, 0x41, 0x30, 0x90, 0x15, 0x55, 0x90, 0x12, 0xb9, 0x6e, 0xb5,
	0xe4, 0xe4, 0x71, 0x89, 0xa4, 0xda, 0x12, 0x59, 0xa9, 0x24, 0xef, 0x61, 0x80, 0x69, 0x59, 0x3b,
	0x9e, 0x2a, 0xc7, 0xce, 0xb7, 0xc6, 0xb4, 0x6c, 0x29, 0x0c, 0x4c, 0xcb, 0xca, 0x70, 0x0d, 0xc7,
	0xcd, 0x1a, 0xd6, 0x9a, 0x67, 0x4a, 0xf3, 0x72, 0x97, 0xa6, 0x21, 0x5b, 0xae, 0xa3, 0xa6, 0xfa,
	0x6b, 0x24, 0x81, 0x9b, 0x58, 0xd6, 0xe3, 0x23, 0x09, 0xdc, 0x4e, 0x65, 0x08, 0x5c, 0x87, 0x72,
	0x01, 0x89, 0x92, 0x14, 0x11, 0x26, 0x49, 0x46, 0x09, 0xae, 0x55, 0xb1, 0x5c, 0x88, 0xc5, 0x41,
	0x60, 0xca, 0x3b, 0xaf, 0xba, 0x52, 0xfc, 0xe4, 0x0a, 0x34, 0xb9, 0x42, 0xc8, 0x00, 0x6d, 0x75,
	0xb9, 0xfc, 0x6c, 0x1e, 0xc8, 0xaf, 0x45, 0x18, 0xae, 0xcc, 0x0e, 0xea, 0x43, 0xcf, 0x5b, 0xde,
	0x98, 0x5d, 0x34, 0x02, 0xc3, 0xff, 0x1a, 0xfa, 0xc1, 0xf2, 0xc3, 0x95, 0xd9, 0x93, 0xe5, 0x5b,
	0x6f, 0x65, 0x6a, 0xc8, 0x84, 0xd1, 0xed, 0x8d, 0x1f, 0x44, 0x9e, 0xff, 0xe9, 0x72, 0xe9, 0x7b,
	0xa6, 0xfe, 0xb1, 0x0f, 0xba, 0xfa, 0xc1, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xa4, 0x46,
	0xfa, 0x7a, 0x04, 0x00, 0x00,
}
