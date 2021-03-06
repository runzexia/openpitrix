// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/types/etcd.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/metadata/types"

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

type EtcdConfig struct {
	User                 string          `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	Password             string          `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	MaxTxnOps            int32           `protobuf:"varint,3,opt,name=max_txn_ops,json=maxTxnOps,proto3" json:"max_txn_ops"`
	TimeoutSeconds       int32           `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
	NodeList             []*EtcdEndpoint `protobuf:"bytes,5,rep,name=node_list,json=nodeList,proto3" json:"node_list"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EtcdConfig) Reset()         { *m = EtcdConfig{} }
func (m *EtcdConfig) String() string { return proto.CompactTextString(m) }
func (*EtcdConfig) ProtoMessage()    {}
func (*EtcdConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_etcd_afa84e248da3d15f, []int{0}
}
func (m *EtcdConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtcdConfig.Unmarshal(m, b)
}
func (m *EtcdConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtcdConfig.Marshal(b, m, deterministic)
}
func (dst *EtcdConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtcdConfig.Merge(dst, src)
}
func (m *EtcdConfig) XXX_Size() int {
	return xxx_messageInfo_EtcdConfig.Size(m)
}
func (m *EtcdConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EtcdConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EtcdConfig proto.InternalMessageInfo

func (m *EtcdConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *EtcdConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EtcdConfig) GetMaxTxnOps() int32 {
	if m != nil {
		return m.MaxTxnOps
	}
	return 0
}

func (m *EtcdConfig) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *EtcdConfig) GetNodeList() []*EtcdEndpoint {
	if m != nil {
		return m.NodeList
	}
	return nil
}

type EtcdEndpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtcdEndpoint) Reset()         { *m = EtcdEndpoint{} }
func (m *EtcdEndpoint) String() string { return proto.CompactTextString(m) }
func (*EtcdEndpoint) ProtoMessage()    {}
func (*EtcdEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_etcd_afa84e248da3d15f, []int{1}
}
func (m *EtcdEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtcdEndpoint.Unmarshal(m, b)
}
func (m *EtcdEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtcdEndpoint.Marshal(b, m, deterministic)
}
func (dst *EtcdEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtcdEndpoint.Merge(dst, src)
}
func (m *EtcdEndpoint) XXX_Size() int {
	return xxx_messageInfo_EtcdEndpoint.Size(m)
}
func (m *EtcdEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_EtcdEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_EtcdEndpoint proto.InternalMessageInfo

func (m *EtcdEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *EtcdEndpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*EtcdConfig)(nil), "metadata.types.EtcdConfig")
	proto.RegisterType((*EtcdEndpoint)(nil), "metadata.types.EtcdEndpoint")
}

func init() { proto.RegisterFile("metadata/types/etcd.proto", fileDescriptor_etcd_afa84e248da3d15f) }

var fileDescriptor_etcd_afa84e248da3d15f = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x6d, 0xa5, 0xd9, 0x4a, 0x85, 0x3d, 0x45, 0x11, 0x09, 0xbd, 0x98, 0x53, 0x16,
	0x14, 0x8a, 0xe2, 0x4d, 0xe9, 0x4d, 0x10, 0xa2, 0x27, 0x2f, 0x61, 0x93, 0x5d, 0xeb, 0xa2, 0xd9,
	0x19, 0x32, 0x53, 0x8c, 0xff, 0xce, 0x9f, 0x26, 0xbb, 0x5a, 0xb5, 0xa7, 0x79, 0xf3, 0xbd, 0x39,
	0xbc, 0x79, 0xe2, 0xa8, 0xb3, 0xac, 0x8d, 0x66, 0xad, 0xf8, 0x03, 0x2d, 0x29, 0xcb, 0xad, 0x29,
	0xb1, 0x07, 0x06, 0x39, 0xdf, 0x5a, 0x65, 0xb4, 0x16, 0x9f, 0x89, 0x10, 0x2b, 0x6e, 0xcd, 0x2d,
	0xf8, 0x67, 0xb7, 0x96, 0x52, 0x8c, 0x37, 0x64, 0xfb, 0x2c, 0xc9, 0x93, 0x22, 0xad, 0xa2, 0x96,
	0xc7, 0x62, 0x8a, 0x9a, 0xe8, 0x1d, 0x7a, 0x93, 0xed, 0x45, 0xfe, 0xbb, 0xcb, 0x53, 0x31, 0xeb,
	0xf4, 0x50, 0xf3, 0xe0, 0x6b, 0x40, 0xca, 0x46, 0x79, 0x52, 0x4c, 0xaa, 0xb4, 0xd3, 0xc3, 0xe3,
	0xe0, 0xef, 0x91, 0xe4, 0x99, 0x38, 0x64, 0xd7, 0x59, 0xd8, 0x70, 0x4d, 0xb6, 0x05, 0x6f, 0x28,
	0x1b, 0xc7, 0x9b, 0xf9, 0x0f, 0x7e, 0xf8, 0xa6, 0xf2, 0x4a, 0xa4, 0x1e, 0x8c, 0xad, 0xdf, 0x1c,
	0x71, 0x36, 0xc9, 0x47, 0xc5, 0xec, 0xfc, 0xa4, 0xdc, 0xcd, 0x5a, 0x86, 0x9c, 0x2b, 0x6f, 0x10,
	0x9c, 0xe7, 0x6a, 0x1a, 0xce, 0xef, 0x1c, 0xf1, 0x62, 0x29, 0x0e, 0xfe, 0x3b, 0xe1, 0x87, 0x17,
	0x20, 0xde, 0xfe, 0x10, 0x74, 0x60, 0x08, 0x3d, 0xc7, 0xfc, 0x93, 0x2a, 0xea, 0x9b, 0xcb, 0xa7,
	0x25, 0xa0, 0xf5, 0xe8, 0xb8, 0x77, 0x43, 0xe9, 0x40, 0xfd, 0x6d, 0x0a, 0x5f, 0xd7, 0x0a, 0x1b,
	0xb5, 0xdb, 0xe3, 0x35, 0x36, 0x71, 0x36, 0xfb, 0xb1, 0xcb, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0xaa, 0x7e, 0x5c, 0x68, 0x01, 0x00, 0x00,
}
