// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/external/config.proto

/*
Package external is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/external/config.proto

It has these top-level messages:
	ProbeConf
*/
package external

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// External probes support two mode: ONCE and SERVER. In ONCE mode, external
// command is re-executed for each probe run, while in SERVER mode, command
// is run in server mode, re-executed only if not running already.
type ProbeConf_Mode int32

const (
	ProbeConf_ONCE   ProbeConf_Mode = 0
	ProbeConf_SERVER ProbeConf_Mode = 1
)

var ProbeConf_Mode_name = map[int32]string{
	0: "ONCE",
	1: "SERVER",
}
var ProbeConf_Mode_value = map[string]int32{
	"ONCE":   0,
	"SERVER": 1,
}

func (x ProbeConf_Mode) Enum() *ProbeConf_Mode {
	p := new(ProbeConf_Mode)
	*p = x
	return p
}
func (x ProbeConf_Mode) String() string {
	return proto.EnumName(ProbeConf_Mode_name, int32(x))
}
func (x *ProbeConf_Mode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_Mode_value, data, "ProbeConf_Mode")
	if err != nil {
		return err
	}
	*x = ProbeConf_Mode(value)
	return nil
}
func (ProbeConf_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ProbeConf struct {
	Mode *ProbeConf_Mode `protobuf:"varint,1,opt,name=mode,enum=cloudprober.probes.external.ProbeConf_Mode,def=1" json:"mode,omitempty"`
	// Command.  For ONCE probes, arguments are processed for the following field
	// substitutions:
	// @probe@    Name of the probe
	// @target@   Hostname of the target
	// @address@  IP address of the target
	//
	// For example, for target ig-us-central1-a, /tools/recreate_vm -vm @target@
	// will get converted to: /tools/recreate_vm -vm ig-us-central1-a
	Command *string `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	// Options for the ProbeRequest for SERVER mode probes.
	Options []*ProbeConf_Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// Additional labels, e.g. "region=us-east1".
	Labels *string `protobuf:"bytes,4,opt,name=labels" json:"labels,omitempty"`
	// Stats export interval in milliseconds.
	StatsExportIntervalMsec *int32 `protobuf:"varint,5,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	// IP version: For target resolution
	IpVersion        *int32 `protobuf:"varint,6,opt,name=ip_version,json=ipVersion,def=4" json:"ip_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_Mode ProbeConf_Mode = ProbeConf_SERVER
const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000
const Default_ProbeConf_IpVersion int32 = 4

func (m *ProbeConf) GetMode() ProbeConf_Mode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_ProbeConf_Mode
}

func (m *ProbeConf) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

func (m *ProbeConf) GetOptions() []*ProbeConf_Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ProbeConf) GetLabels() string {
	if m != nil && m.Labels != nil {
		return *m.Labels
	}
	return ""
}

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func (m *ProbeConf) GetIpVersion() int32 {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return Default_ProbeConf_IpVersion
}

type ProbeConf_Option struct {
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Values are substituted similar to command arguments for ONCE mode probes.
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProbeConf_Option) Reset()                    { *m = ProbeConf_Option{} }
func (m *ProbeConf_Option) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf_Option) ProtoMessage()               {}
func (*ProbeConf_Option) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ProbeConf_Option) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeConf_Option) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.external.ProbeConf")
	proto.RegisterType((*ProbeConf_Option)(nil), "cloudprober.probes.external.ProbeConf.Option")
	proto.RegisterEnum("cloudprober.probes.external.ProbeConf_Mode", ProbeConf_Mode_name, ProbeConf_Mode_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/external/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xad, 0x6b, 0x3b, 0xfb, 0x04, 0x19, 0x41, 0xb4, 0x4c, 0x0f, 0x65, 0xa7, 0x82, 0xd8,
	0xd6, 0xe1, 0x69, 0x47, 0x47, 0x19, 0x1e, 0xe6, 0x24, 0xc2, 0xae, 0xa5, 0x6b, 0xb3, 0x5a, 0x48,
	0xf3, 0x4a, 0x92, 0x8d, 0x7d, 0x03, 0xbf, 0xb6, 0x2c, 0xed, 0xc4, 0x93, 0x78, 0x6a, 0xff, 0x8f,
	0xdf, 0x2f, 0x2f, 0xf9, 0xc3, 0xac, 0xaa, 0xf5, 0xe7, 0x6e, 0x13, 0x15, 0xd8, 0xc4, 0x15, 0x62,
	0xc5, 0x59, 0x5c, 0x70, 0xdc, 0x95, 0xad, 0xc4, 0x0d, 0x93, 0xb1, 0xf9, 0xa8, 0x98, 0x1d, 0x34,
	0x93, 0x22, 0xe7, 0x71, 0x81, 0x62, 0x5b, 0x57, 0x51, 0x2b, 0x51, 0x23, 0xb9, 0xfb, 0x45, 0x46,
	0x1d, 0x19, 0x9d, 0xc8, 0xc9, 0xd7, 0x00, 0xbc, 0xf7, 0xe3, 0x6c, 0x8e, 0x62, 0x4b, 0x16, 0x60,
	0x37, 0x58, 0x32, 0xdf, 0x0a, 0xac, 0xf0, 0x6a, 0xfa, 0x10, 0xfd, 0x61, 0x46, 0x3f, 0x56, 0xb4,
	0xc4, 0x92, 0xcd, 0xdc, 0x8f, 0x94, 0xae, 0x53, 0x4a, 0xcd, 0x01, 0xc4, 0x87, 0x61, 0x81, 0x4d,
	0x93, 0x8b, 0xd2, 0x3f, 0x0f, 0xac, 0xd0, 0xa3, 0xa7, 0x48, 0x16, 0x30, 0xc4, 0x56, 0xd7, 0x28,
	0x94, 0x3f, 0x08, 0x06, 0xe1, 0xe5, 0xf4, 0xf1, 0x9f, 0x5b, 0x56, 0xc6, 0xa2, 0x27, 0x9b, 0xdc,
	0x80, 0xcb, 0xf3, 0x0d, 0xe3, 0xca, 0xb7, 0xcd, 0x86, 0x3e, 0x91, 0x17, 0x18, 0x2b, 0x9d, 0x6b,
	0x95, 0xb1, 0x43, 0x8b, 0x52, 0x67, 0xb5, 0xd0, 0x4c, 0xee, 0x73, 0x9e, 0x35, 0x8a, 0x15, 0xbe,
	0x13, 0x58, 0xa1, 0x33, 0x73, 0x9e, 0x92, 0x24, 0x49, 0xe8, 0xad, 0x01, 0x53, 0xc3, 0xbd, 0xf6,
	0xd8, 0x52, 0xb1, 0x82, 0x04, 0x00, 0x75, 0x9b, 0xed, 0x99, 0x54, 0x35, 0x0a, 0xdf, 0x35, 0x8e,
	0xf5, 0x4c, 0xbd, 0xba, 0x5d, 0x77, 0xb3, 0xf1, 0x14, 0xdc, 0xee, 0x42, 0x84, 0x80, 0x2d, 0xf2,
	0xa6, 0xeb, 0xcc, 0xa3, 0xe6, 0x9f, 0x5c, 0x83, 0xb3, 0xcf, 0xf9, 0x8e, 0xf5, 0x8f, 0xef, 0xc2,
	0xe4, 0x1e, 0xec, 0x63, 0x55, 0xe4, 0x02, 0xec, 0xd5, 0xdb, 0x3c, 0x1d, 0x9d, 0x11, 0x80, 0xbe,
	0xb6, 0x91, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x09, 0x96, 0x93, 0x29, 0xe3, 0x01, 0x00, 0x00,
}
