// Code generated by protoc-gen-go.
// source: vault.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	vault.proto

It has these top-level messages:
	InitStatusRequest
	InitStatusResponse
	InitRequest
	InitResponse
	SealStatusRequest
	SealStatusResponse
	UnsealRequest
	UnsealResponse
	Status
	SealStatus
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message is currently empty, as this request is empty on Vault.
type InitStatusRequest struct {
}

func (m *InitStatusRequest) Reset()                    { *m = InitStatusRequest{} }
func (m *InitStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*InitStatusRequest) ProtoMessage()               {}
func (*InitStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// InitStatusResponse is the output from a GET to /sys/init
type InitStatusResponse struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Err    string  `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *InitStatusResponse) Reset()                    { *m = InitStatusResponse{} }
func (m *InitStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*InitStatusResponse) ProtoMessage()               {}
func (*InitStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InitStatusResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type InitRequest struct {
	SecretShares      uint32   `protobuf:"varint,1,opt,name=secret_shares,json=secretShares" json:"secret_shares,omitempty"`
	SecretThreshold   uint32   `protobuf:"varint,2,opt,name=secret_threshold,json=secretThreshold" json:"secret_threshold,omitempty"`
	StoredShares      uint32   `protobuf:"varint,3,opt,name=stored_shares,json=storedShares" json:"stored_shares,omitempty"`
	PgpKeys           []string `protobuf:"bytes,4,rep,name=pgp_keys,json=pgpKeys" json:"pgp_keys,omitempty"`
	RecoveryShares    uint32   `protobuf:"varint,5,opt,name=recovery_shares,json=recoveryShares" json:"recovery_shares,omitempty"`
	RecoveryThreshold uint32   `protobuf:"varint,6,opt,name=recovery_threshold,json=recoveryThreshold" json:"recovery_threshold,omitempty"`
	RecoveryPgpKeys   []string `protobuf:"bytes,7,rep,name=recovery_pgp_keys,json=recoveryPgpKeys" json:"recovery_pgp_keys,omitempty"`
	RootTokenPgpKey   string   `protobuf:"bytes,8,opt,name=root_token_pgp_key,json=rootTokenPgpKey" json:"root_token_pgp_key,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InitResponse struct {
	Keys               []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	KeysBase64         []string `protobuf:"bytes,2,rep,name=keys_base64,json=keysBase64" json:"keys_base64,omitempty"`
	RecoveryKeys       []string `protobuf:"bytes,3,rep,name=recovery_keys,json=recoveryKeys" json:"recovery_keys,omitempty"`
	RecoveryKeysBase64 []string `protobuf:"bytes,4,rep,name=recovery_keys_base64,json=recoveryKeysBase64" json:"recovery_keys_base64,omitempty"`
	RootToken          string   `protobuf:"bytes,5,opt,name=root_token,json=rootToken" json:"root_token,omitempty"`
	Err                string   `protobuf:"bytes,6,opt,name=err" json:"err,omitempty"`
}

func (m *InitResponse) Reset()                    { *m = InitResponse{} }
func (m *InitResponse) String() string            { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()               {}
func (*InitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// The request message containing request name (Not really used but grpc requires it).
type SealStatusRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SealStatusRequest) Reset()                    { *m = SealStatusRequest{} }
func (m *SealStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*SealStatusRequest) ProtoMessage()               {}
func (*SealStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SealStatusResponse struct {
	SealStatus *SealStatus `protobuf:"bytes,1,opt,name=seal_status,json=sealStatus" json:"seal_status,omitempty"`
	Err        string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *SealStatusResponse) Reset()                    { *m = SealStatusResponse{} }
func (m *SealStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*SealStatusResponse) ProtoMessage()               {}
func (*SealStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SealStatusResponse) GetSealStatus() *SealStatus {
	if m != nil {
		return m.SealStatus
	}
	return nil
}

type UnsealRequest struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Reset_ bool   `protobuf:"varint,2,opt,name=reset" json:"reset,omitempty"`
}

func (m *UnsealRequest) Reset()                    { *m = UnsealRequest{} }
func (m *UnsealRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsealRequest) ProtoMessage()               {}
func (*UnsealRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type UnsealResponse struct {
	SealStatus *SealStatus `protobuf:"bytes,1,opt,name=seal_status,json=sealStatus" json:"seal_status,omitempty"`
	Err        string      `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *UnsealResponse) Reset()                    { *m = UnsealResponse{} }
func (m *UnsealResponse) String() string            { return proto.CompactTextString(m) }
func (*UnsealResponse) ProtoMessage()               {}
func (*UnsealResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UnsealResponse) GetSealStatus() *SealStatus {
	if m != nil {
		return m.SealStatus
	}
	return nil
}

//       Iniitialization status of Vault
type Status struct {
	Initialized bool `protobuf:"varint,1,opt,name=initialized" json:"initialized,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

//       Seal status of Vault
type SealStatus struct {
	Sealed      bool   `protobuf:"varint,1,opt,name=sealed" json:"sealed,omitempty"`
	T           uint32 `protobuf:"varint,2,opt,name=t" json:"t,omitempty"`
	N           uint32 `protobuf:"varint,3,opt,name=n" json:"n,omitempty"`
	Progress    uint32 `protobuf:"varint,4,opt,name=progress" json:"progress,omitempty"`
	Version     string `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	ClusterName string `protobuf:"bytes,6,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	ClusterId   string `protobuf:"bytes,7,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
}

func (m *SealStatus) Reset()                    { *m = SealStatus{} }
func (m *SealStatus) String() string            { return proto.CompactTextString(m) }
func (*SealStatus) ProtoMessage()               {}
func (*SealStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*InitStatusRequest)(nil), "pb.InitStatusRequest")
	proto.RegisterType((*InitStatusResponse)(nil), "pb.InitStatusResponse")
	proto.RegisterType((*InitRequest)(nil), "pb.InitRequest")
	proto.RegisterType((*InitResponse)(nil), "pb.InitResponse")
	proto.RegisterType((*SealStatusRequest)(nil), "pb.SealStatusRequest")
	proto.RegisterType((*SealStatusResponse)(nil), "pb.SealStatusResponse")
	proto.RegisterType((*UnsealRequest)(nil), "pb.UnsealRequest")
	proto.RegisterType((*UnsealResponse)(nil), "pb.UnsealResponse")
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*SealStatus)(nil), "pb.SealStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Vault service

type VaultClient interface {
	// InitStatus retrieves the output from a GET to /sys/init
	// This returns the initialization status of Vault.
	InitStatus(ctx context.Context, in *InitStatusRequest, opts ...grpc.CallOption) (*InitStatusResponse, error)
}

type vaultClient struct {
	cc *grpc.ClientConn
}

func NewVaultClient(cc *grpc.ClientConn) VaultClient {
	return &vaultClient{cc}
}

func (c *vaultClient) InitStatus(ctx context.Context, in *InitStatusRequest, opts ...grpc.CallOption) (*InitStatusResponse, error) {
	out := new(InitStatusResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/InitStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Vault service

type VaultServer interface {
	// InitStatus retrieves the output from a GET to /sys/init
	// This returns the initialization status of Vault.
	InitStatus(context.Context, *InitStatusRequest) (*InitStatusResponse, error)
}

func RegisterVaultServer(s *grpc.Server, srv VaultServer) {
	s.RegisterService(&_Vault_serviceDesc, srv)
}

func _Vault_InitStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).InitStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/InitStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).InitStatus(ctx, req.(*InitStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vault_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Vault",
	HandlerType: (*VaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitStatus",
			Handler:    _Vault_InitStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("vault.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0xfd, 0xa5, 0xdd, 0x4d, 0xdb, 0x49, 0xff, 0xfa, 0xb7, 0xac, 0xc2, 0x4a, 0x88, 0x12, 0x0e,
	0x5b, 0x16, 0x51, 0xd0, 0x82, 0xe0, 0xc4, 0x85, 0x03, 0xd2, 0x82, 0x84, 0x50, 0xba, 0xc0, 0x31,
	0x4a, 0xdb, 0x51, 0x1b, 0x6d, 0x36, 0x0e, 0xb6, 0x5b, 0xa9, 0x7c, 0x32, 0xbe, 0x03, 0x12, 0x9f,
	0x09, 0x79, 0x6c, 0x27, 0x2d, 0x70, 0xe4, 0x14, 0xcf, 0x9b, 0x97, 0x37, 0xe3, 0xf1, 0xb3, 0x21,
	0xd8, 0xa6, 0x9b, 0x5c, 0x4d, 0x4b, 0xc1, 0x15, 0x67, 0x8d, 0x72, 0x1e, 0xfd, 0x0f, 0xa3, 0xab,
	0x22, 0x53, 0x33, 0x95, 0xaa, 0x8d, 0x8c, 0xf1, 0xeb, 0x06, 0xa5, 0x8a, 0xde, 0x01, 0xdb, 0x07,
	0x65, 0xc9, 0x0b, 0x89, 0x2c, 0x02, 0x5f, 0x12, 0x12, 0x7a, 0x63, 0x6f, 0x12, 0x5c, 0xc2, 0xb4,
	0x9c, 0x4f, 0x2d, 0xc7, 0x66, 0xd8, 0x10, 0x9a, 0x28, 0x44, 0xd8, 0x18, 0x7b, 0x93, 0x4e, 0xac,
	0x97, 0xd1, 0xcf, 0x06, 0x04, 0x5a, 0xcc, 0x6a, 0xb3, 0x87, 0xd0, 0x93, 0xb8, 0x10, 0xa8, 0x12,
	0xb9, 0x4e, 0x05, 0x1a, 0xb1, 0x5e, 0xdc, 0x35, 0xe0, 0x8c, 0x30, 0xf6, 0x08, 0x86, 0x96, 0xa4,
	0xd6, 0x02, 0xe5, 0x9a, 0xe7, 0x4b, 0xd2, 0xec, 0xc5, 0x03, 0x83, 0x5f, 0x3b, 0x98, 0xf4, 0x14,
	0x17, 0xb8, 0x74, 0x7a, 0x4d, 0xab, 0x47, 0xa0, 0xd5, 0xbb, 0x0b, 0xed, 0x72, 0x55, 0x26, 0x37,
	0xb8, 0x93, 0xe1, 0xd1, 0xb8, 0x39, 0xe9, 0xc4, 0xad, 0x72, 0x55, 0xbe, 0xc7, 0x9d, 0x64, 0xe7,
	0x30, 0x10, 0xb8, 0xe0, 0x5b, 0x14, 0x3b, 0xa7, 0x70, 0x4c, 0x0a, 0x7d, 0x07, 0x5b, 0x8d, 0x27,
	0xc0, 0x2a, 0x62, 0xdd, 0x95, 0x4f, 0xdc, 0x91, 0xcb, 0xd4, 0x7d, 0x5d, 0x40, 0x05, 0x26, 0x55,
	0xed, 0x16, 0xd5, 0xae, 0x0a, 0x7e, 0xb4, 0x3d, 0x3c, 0x06, 0x26, 0x38, 0x57, 0x89, 0xe2, 0x37,
	0x58, 0x38, 0x76, 0xd8, 0xa6, 0x21, 0x0e, 0x74, 0xe6, 0x5a, 0x27, 0x0c, 0x3b, 0xfa, 0xe1, 0x41,
	0xd7, 0x0c, 0xd4, 0x9e, 0x0b, 0x83, 0x23, 0x12, 0xf7, 0x48, 0x9c, 0xd6, 0xec, 0x3e, 0x04, 0xfa,
	0x9b, 0xcc, 0x53, 0x89, 0x2f, 0x5f, 0x84, 0x0d, 0x4a, 0x81, 0x86, 0xde, 0x10, 0xa2, 0xc7, 0x56,
	0xb5, 0x47, 0x7f, 0x37, 0x89, 0xd2, 0x75, 0x20, 0xf5, 0xf5, 0x0c, 0x4e, 0x0e, 0x48, 0x4e, 0xce,
	0x8c, 0x90, 0xed, 0x73, 0xad, 0xec, 0x3d, 0x80, 0x7a, 0x27, 0x34, 0xc8, 0x4e, 0xdc, 0xa9, 0x76,
	0xe0, 0xec, 0xe1, 0xd7, 0xf6, 0x38, 0x87, 0xd1, 0x0c, 0xd3, 0xfc, 0xc0, 0x7f, 0x7a, 0x47, 0x45,
	0x7a, 0x8b, 0x64, 0x8d, 0x4e, 0x4c, 0xeb, 0xe8, 0x0b, 0xb0, 0x7d, 0xa2, 0xdd, 0xfb, 0x53, 0x08,
	0x24, 0xa6, 0x79, 0x72, 0x60, 0xcc, 0x3e, 0x19, 0xb3, 0x26, 0x83, 0xac, 0xd6, 0x7f, 0x31, 0xe8,
	0x2b, 0xe8, 0x7d, 0x2a, 0x34, 0xc3, 0x55, 0x1f, 0x42, 0x53, 0x8f, 0xdf, 0x14, 0xd7, 0x4b, 0x76,
	0x02, 0xc7, 0x02, 0x25, 0x2a, 0xfa, 0xad, 0x1d, 0x9b, 0x20, 0x9a, 0x41, 0xdf, 0xfd, 0xf8, 0xef,
	0xba, 0xb9, 0x00, 0xdf, 0xe6, 0xc6, 0x10, 0x64, 0x45, 0xa6, 0xb2, 0x34, 0xcf, 0xbe, 0xe1, 0x92,
	0xc4, 0xda, 0xf1, 0x3e, 0x14, 0x7d, 0xf7, 0x00, 0x6a, 0x61, 0x76, 0x0a, 0xbe, 0x96, 0xae, 0xb8,
	0x36, 0x62, 0x5d, 0xf0, 0x94, 0xbd, 0x3d, 0x9e, 0xd2, 0x51, 0x61, 0xef, 0x88, 0x57, 0xb0, 0x33,
	0x68, 0x97, 0x82, 0xaf, 0x04, 0x4a, 0x7d, 0x31, 0x34, 0x58, 0xc5, 0x2c, 0x84, 0xd6, 0x16, 0x85,
	0xcc, 0xb8, 0x3b, 0x48, 0x17, 0xb2, 0x07, 0xd0, 0x5d, 0xe4, 0x1b, 0xa9, 0x50, 0x24, 0x74, 0x4e,
	0xe6, 0x3c, 0x03, 0x8b, 0x7d, 0x48, 0x6f, 0x51, 0x1b, 0xc1, 0x51, 0xb2, 0x65, 0xd8, 0x32, 0x46,
	0xb0, 0xc8, 0xd5, 0xf2, 0xf2, 0x2d, 0x1c, 0x7f, 0xd6, 0x2f, 0x11, 0x7b, 0x0d, 0x50, 0x3f, 0x35,
	0xec, 0x8e, 0x9e, 0xd5, 0x1f, 0xef, 0xd1, 0xd9, 0xe9, 0xef, 0xb0, 0x99, 0x77, 0xf4, 0xdf, 0xdc,
	0xa7, 0x97, 0xec, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x87, 0x15, 0xf5, 0xd8, 0x04,
	0x00, 0x00,
}