// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Client_ClientType int32

const (
	Client_UNKNOWN Client_ClientType = 0
	Client_SDK     Client_ClientType = 1
	Client_AGENT   Client_ClientType = 2
)

var Client_ClientType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SDK",
	2: "AGENT",
}
var Client_ClientType_value = map[string]int32{
	"UNKNOWN": 0,
	"SDK":     1,
	"AGENT":   2,
}

func (x Client_ClientType) String() string {
	return proto.EnumName(Client_ClientType_name, int32(x))
}
func (Client_ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_7fce320934922288, []int{0, 0}
}

type Client struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type                 Client_ClientType     `protobuf:"varint,2,opt,name=type,proto3,enum=v1.Client_ClientType" json:"type,omitempty"`
	Version              *wrappers.StringValue `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Location             *Location             `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_7fce320934922288, []int{0}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Client) GetType() Client_ClientType {
	if m != nil {
		return m.Type
	}
	return Client_UNKNOWN
}

func (m *Client) GetVersion() *wrappers.StringValue {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Client) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterEnum("v1.Client_ClientType", Client_ClientType_name, Client_ClientType_value)
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_client_7fce320934922288) }

var fileDescriptor_client_7fce320934922288 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x4b, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x17, 0x25, 0x16, 0x14,
	0xa4, 0x16, 0x15, 0x43, 0xd4, 0x48, 0x71, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0xdf,
	0x18, 0xb9, 0xd8, 0x9c, 0xc1, 0x26, 0x08, 0x19, 0x70, 0xb1, 0x64, 0xe4, 0x17, 0x97, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xc9, 0xe8, 0x41, 0x8c, 0xd1, 0x83, 0x19, 0xa3, 0x17, 0x5c, 0x52,
	0x94, 0x99, 0x97, 0x1e, 0x96, 0x98, 0x53, 0x9a, 0x1a, 0x04, 0x56, 0x29, 0xa4, 0xc9, 0xc5, 0x52,
	0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xaa, 0x57, 0x66, 0xa8, 0x07,
	0x31, 0x0b, 0x4a, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x81, 0x95, 0x08, 0x99, 0x71, 0xb1, 0x97, 0xa5,
	0x16, 0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x13, 0x61, 0x3e, 0x4c, 0xb1, 0x90, 0x06, 0x17, 0x47,
	0x4e, 0x7e, 0x72, 0x62, 0x09, 0x48, 0x23, 0x0b, 0x58, 0x23, 0x0f, 0xc8, 0x1a, 0x1f, 0xa8, 0x58,
	0x10, 0x5c, 0x56, 0x49, 0x97, 0x8b, 0x0b, 0x61, 0xab, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7,
	0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x10, 0x3b, 0x17, 0x73, 0xb0, 0x8b, 0xb7, 0x00, 0xa3, 0x10,
	0x27, 0x17, 0xab, 0xa3, 0xbb, 0xab, 0x5f, 0x88, 0x00, 0x53, 0x12, 0x1b, 0xd8, 0x5e, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x03, 0xfc, 0x88, 0x40, 0x01, 0x00, 0x00,
}