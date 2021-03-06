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
	ConfigureRequest
	ConfigureResponse
	ConfigStatus
	MountOutput
	MountConfigOutput
	AuthMountOutput
	AuthConfigOutput
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

// The request message is currently empty, as this request is empty on Vault.
type SealStatusRequest struct {
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

type ConfigureRequest struct {
	Url   string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ConfigureResponse struct {
	ConfigStatus *ConfigStatus `protobuf:"bytes,1,opt,name=config_status,json=configStatus" json:"config_status,omitempty"`
	Err          string        `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ConfigureResponse) Reset()                    { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()               {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ConfigureResponse) GetConfigStatus() *ConfigStatus {
	if m != nil {
		return m.ConfigStatus
	}
	return nil
}

type ConfigStatus struct {
	ConfigId string                      `protobuf:"bytes,1,opt,name=config_id,json=configId" json:"config_id,omitempty"`
	Mounts   map[string]*MountOutput     `protobuf:"bytes,2,rep,name=mounts" json:"mounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Auths    map[string]*AuthMountOutput `protobuf:"bytes,3,rep,name=auths" json:"auths,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies []string                    `protobuf:"bytes,4,rep,name=policies" json:"policies,omitempty"`
}

func (m *ConfigStatus) Reset()                    { *m = ConfigStatus{} }
func (m *ConfigStatus) String() string            { return proto.CompactTextString(m) }
func (*ConfigStatus) ProtoMessage()               {}
func (*ConfigStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ConfigStatus) GetMounts() map[string]*MountOutput {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *ConfigStatus) GetAuths() map[string]*AuthMountOutput {
	if m != nil {
		return m.Auths
	}
	return nil
}

type MountOutput struct {
	Type        string             `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Config      *MountConfigOutput `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *MountOutput) Reset()                    { *m = MountOutput{} }
func (m *MountOutput) String() string            { return proto.CompactTextString(m) }
func (*MountOutput) ProtoMessage()               {}
func (*MountOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *MountOutput) GetConfig() *MountConfigOutput {
	if m != nil {
		return m.Config
	}
	return nil
}

type MountConfigOutput struct {
	DefaultLeaseTtl uint32 `protobuf:"varint,1,opt,name=default_lease_ttl,json=defaultLeaseTtl" json:"default_lease_ttl,omitempty"`
	MaxLeaseTtl     uint32 `protobuf:"varint,2,opt,name=max_lease_ttl,json=maxLeaseTtl" json:"max_lease_ttl,omitempty"`
}

func (m *MountConfigOutput) Reset()                    { *m = MountConfigOutput{} }
func (m *MountConfigOutput) String() string            { return proto.CompactTextString(m) }
func (*MountConfigOutput) ProtoMessage()               {}
func (*MountConfigOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type AuthMountOutput struct {
	Type        string            `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Config      *AuthConfigOutput `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *AuthMountOutput) Reset()                    { *m = AuthMountOutput{} }
func (m *AuthMountOutput) String() string            { return proto.CompactTextString(m) }
func (*AuthMountOutput) ProtoMessage()               {}
func (*AuthMountOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *AuthMountOutput) GetConfig() *AuthConfigOutput {
	if m != nil {
		return m.Config
	}
	return nil
}

type AuthConfigOutput struct {
	DefaultLeaseTtl uint32 `protobuf:"varint,1,opt,name=default_lease_ttl,json=defaultLeaseTtl" json:"default_lease_ttl,omitempty"`
	MaxLeaseTtl     uint32 `protobuf:"varint,2,opt,name=max_lease_ttl,json=maxLeaseTtl" json:"max_lease_ttl,omitempty"`
}

func (m *AuthConfigOutput) Reset()                    { *m = AuthConfigOutput{} }
func (m *AuthConfigOutput) String() string            { return proto.CompactTextString(m) }
func (*AuthConfigOutput) ProtoMessage()               {}
func (*AuthConfigOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

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
	proto.RegisterType((*ConfigureRequest)(nil), "pb.ConfigureRequest")
	proto.RegisterType((*ConfigureResponse)(nil), "pb.ConfigureResponse")
	proto.RegisterType((*ConfigStatus)(nil), "pb.ConfigStatus")
	proto.RegisterType((*MountOutput)(nil), "pb.MountOutput")
	proto.RegisterType((*MountConfigOutput)(nil), "pb.MountConfigOutput")
	proto.RegisterType((*AuthMountOutput)(nil), "pb.AuthMountOutput")
	proto.RegisterType((*AuthConfigOutput)(nil), "pb.AuthConfigOutput")
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
	// Init retrieves the output from a PUT to /sys/init
	// This initializes a new Vault.  Vault must not have been previously initialized.
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	// SealStatus retrieves the output from a GET to /sys/seal-status
	// Returns the seal status of the Vault.
	SealStatus(ctx context.Context, in *SealStatusRequest, opts ...grpc.CallOption) (*SealStatusResponse, error)
	// Unseal retrieves the output from a PUT to /sys/unseal
	// Enter a single master key share to progress the unsealing of the Vault.
	// If the threshold number of master key shares is reached, Vault will attempt to
	// unseal the Vault. Otherwise, this API must be called multiple times until that
	// threshold is met.
	Unseal(ctx context.Context, in *UnsealRequest, opts ...grpc.CallOption) (*UnsealResponse, error)
	// Configure applies a set of configuration files to Vault. By
	// convention, these are json files located at some URL (e.g. git or
	// aws s3).
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
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

func (c *vaultClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) SealStatus(ctx context.Context, in *SealStatusRequest, opts ...grpc.CallOption) (*SealStatusResponse, error) {
	out := new(SealStatusResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/SealStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) Unseal(ctx context.Context, in *UnsealRequest, opts ...grpc.CallOption) (*UnsealResponse, error) {
	out := new(UnsealResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/Unseal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vaultClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := grpc.Invoke(ctx, "/pb.Vault/Configure", in, out, c.cc, opts...)
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
	// Init retrieves the output from a PUT to /sys/init
	// This initializes a new Vault.  Vault must not have been previously initialized.
	Init(context.Context, *InitRequest) (*InitResponse, error)
	// SealStatus retrieves the output from a GET to /sys/seal-status
	// Returns the seal status of the Vault.
	SealStatus(context.Context, *SealStatusRequest) (*SealStatusResponse, error)
	// Unseal retrieves the output from a PUT to /sys/unseal
	// Enter a single master key share to progress the unsealing of the Vault.
	// If the threshold number of master key shares is reached, Vault will attempt to
	// unseal the Vault. Otherwise, this API must be called multiple times until that
	// threshold is met.
	Unseal(context.Context, *UnsealRequest) (*UnsealResponse, error)
	// Configure applies a set of configuration files to Vault. By
	// convention, these are json files located at some URL (e.g. git or
	// aws s3).
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
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

func _Vault_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_SealStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SealStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).SealStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/SealStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).SealStatus(ctx, req.(*SealStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_Unseal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).Unseal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/Unseal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).Unseal(ctx, req.(*UnsealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vault_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Vault/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultServer).Configure(ctx, req.(*ConfigureRequest))
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
		{
			MethodName: "Init",
			Handler:    _Vault_Init_Handler,
		},
		{
			MethodName: "SealStatus",
			Handler:    _Vault_SealStatus_Handler,
		},
		{
			MethodName: "Unseal",
			Handler:    _Vault_Unseal_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Vault_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("vault.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x27, 0x8d, 0x9b, 0x1c, 0x27, 0x4d, 0x72, 0xda, 0xae, 0x42, 0x16, 0x44, 0x31, 0x42,
	0x74, 0x77, 0xd9, 0xc2, 0x86, 0xe5, 0x47, 0x95, 0xf6, 0x02, 0x10, 0x17, 0x5d, 0x58, 0x40, 0x6e,
	0x81, 0x1b, 0x24, 0xcb, 0x89, 0xcf, 0xb6, 0x56, 0x1d, 0xdb, 0xcc, 0x8c, 0xab, 0x0d, 0x6f, 0xc3,
	0x5b, 0xf0, 0x0e, 0x2b, 0xf1, 0x4c, 0x68, 0xfe, 0x6c, 0xe7, 0xe7, 0x0a, 0xed, 0x55, 0x3c, 0xdf,
	0xf9, 0xce, 0x37, 0x33, 0xdf, 0x9c, 0x39, 0x13, 0xf0, 0xee, 0xa2, 0x32, 0x15, 0x67, 0x05, 0xcb,
	0x45, 0x8e, 0xad, 0x62, 0xee, 0x1f, 0xc2, 0xf8, 0x22, 0x4b, 0xc4, 0xa5, 0x88, 0x44, 0xc9, 0x03,
	0xfa, 0xb3, 0x24, 0x2e, 0xfc, 0x17, 0x80, 0x4d, 0x90, 0x17, 0x79, 0xc6, 0x09, 0x7d, 0x70, 0xb9,
	0x42, 0x26, 0xce, 0x89, 0x73, 0xea, 0xcd, 0xe0, 0xac, 0x98, 0x9f, 0x19, 0x8e, 0x89, 0xe0, 0x08,
	0xda, 0xc4, 0xd8, 0xa4, 0x75, 0xe2, 0x9c, 0xf6, 0x02, 0xf9, 0xe9, 0xff, 0xdb, 0x02, 0x4f, 0x8a,
	0x19, 0x6d, 0xfc, 0x10, 0x06, 0x9c, 0x16, 0x8c, 0x44, 0xc8, 0x6f, 0x22, 0x46, 0x5a, 0x6c, 0x10,
	0xf4, 0x35, 0x78, 0xa9, 0x30, 0x7c, 0x08, 0x23, 0x43, 0x12, 0x37, 0x8c, 0xf8, 0x4d, 0x9e, 0xc6,
	0x4a, 0x73, 0x10, 0x0c, 0x35, 0x7e, 0x65, 0x61, 0xa5, 0x27, 0x72, 0x46, 0xb1, 0xd5, 0x6b, 0x1b,
	0x3d, 0x05, 0x1a, 0xbd, 0x77, 0xa0, 0x5b, 0x5c, 0x17, 0xe1, 0x2d, 0xad, 0xf8, 0x64, 0xef, 0xa4,
	0x7d, 0xda, 0x0b, 0xf6, 0x8b, 0xeb, 0xe2, 0x07, 0x5a, 0x71, 0xfc, 0x18, 0x86, 0x8c, 0x16, 0xf9,
	0x1d, 0xb1, 0x95, 0x55, 0xe8, 0x28, 0x85, 0x03, 0x0b, 0x1b, 0x8d, 0x27, 0x80, 0x15, 0xb1, 0x5e,
	0x95, 0xab, 0xb8, 0x63, 0x1b, 0xa9, 0xd7, 0xf5, 0x08, 0x2a, 0x30, 0xac, 0xe6, 0xde, 0x57, 0x73,
	0x57, 0x13, 0xfe, 0x62, 0xd6, 0xf0, 0x18, 0x90, 0xe5, 0xb9, 0x08, 0x45, 0x7e, 0x4b, 0x99, 0x65,
	0x4f, 0xba, 0xca, 0xc4, 0xa1, 0x8c, 0x5c, 0xc9, 0x80, 0x66, 0xfb, 0x6f, 0x1c, 0xe8, 0x6b, 0x43,
	0xcd, 0xb9, 0x20, 0xec, 0x29, 0x71, 0x47, 0x89, 0xab, 0x6f, 0x7c, 0x1f, 0x3c, 0xf9, 0x1b, 0xce,
	0x23, 0x4e, 0x5f, 0x3e, 0x9b, 0xb4, 0x54, 0x08, 0x24, 0xf4, 0xad, 0x42, 0xa4, 0x6d, 0xd5, 0xf2,
	0x54, 0x76, 0x5b, 0x51, 0xfa, 0x16, 0x54, 0xeb, 0xfa, 0x0c, 0x8e, 0xd6, 0x48, 0x56, 0x4e, 0x5b,
	0x88, 0x4d, 0xae, 0x91, 0x7d, 0x0f, 0xa0, 0xde, 0x89, 0x32, 0xb2, 0x17, 0xf4, 0xaa, 0x1d, 0xd8,
	0xf2, 0x70, 0xeb, 0xf2, 0x38, 0x84, 0xf1, 0x25, 0x45, 0xe9, 0x7a, 0xfd, 0xfd, 0x0e, 0xd8, 0x04,
	0xcd, 0x3e, 0x3f, 0x05, 0x8f, 0x53, 0x94, 0x86, 0x6b, 0x45, 0x78, 0xa0, 0x8a, 0xb0, 0x26, 0x03,
	0xaf, 0xbe, 0x77, 0x14, 0xe3, 0x57, 0x30, 0xf8, 0x35, 0x93, 0x0c, 0x5b, 0x8d, 0x23, 0x68, 0x4b,
	0xab, 0x1d, 0x4d, 0xb9, 0xa5, 0x15, 0x1e, 0x41, 0x87, 0x11, 0x27, 0xa1, 0xd2, 0xba, 0x81, 0x1e,
	0xf8, 0x97, 0x70, 0x60, 0x13, 0xdf, 0xde, 0x6a, 0x1e, 0x81, 0x6b, 0x62, 0x27, 0xe0, 0x25, 0x59,
	0x22, 0x92, 0x28, 0x4d, 0xfe, 0xa2, 0x58, 0x89, 0x75, 0x83, 0x26, 0xe4, 0xff, 0xe3, 0x00, 0xd4,
	0xc2, 0x78, 0x1f, 0x5c, 0x29, 0x5d, 0x71, 0xcd, 0x08, 0xfb, 0xe0, 0x08, 0x73, 0x53, 0x1c, 0x21,
	0x47, 0x99, 0xb9, 0x0f, 0x4e, 0x86, 0x53, 0xe8, 0x16, 0x2c, 0xbf, 0x66, 0xc4, 0xe5, 0x25, 0x90,
	0x60, 0x35, 0xc6, 0x09, 0xec, 0xdf, 0x11, 0xe3, 0x49, 0x6e, 0x0f, 0xcd, 0x0e, 0xf1, 0x03, 0xe8,
	0x2f, 0xd2, 0x92, 0x0b, 0x62, 0x61, 0x16, 0x2d, 0xc9, 0x9c, 0x9d, 0x67, 0xb0, 0x9f, 0xa2, 0x25,
	0xc9, 0x43, 0xb7, 0x94, 0x24, 0x9e, 0xec, 0xeb, 0x43, 0x37, 0xc8, 0x45, 0xec, 0x9f, 0xc3, 0xe8,
	0xbb, 0x3c, 0x7b, 0x95, 0x5c, 0x97, 0x8c, 0x1a, 0xbe, 0x97, 0x2c, 0xb5, 0xbe, 0x97, 0x2c, 0x95,
	0xbe, 0xeb, 0xa2, 0xd1, 0x06, 0xe9, 0x81, 0xff, 0x07, 0x8c, 0x1b, 0xb9, 0xc6, 0xfa, 0x2f, 0x60,
	0xb0, 0x50, 0xe0, 0xba, 0xf9, 0x23, 0x69, 0xbe, 0x66, 0x1b, 0xfb, 0xfb, 0x8b, 0xc6, 0x68, 0xc7,
	0x01, 0xbc, 0x69, 0x41, 0xbf, 0x99, 0x80, 0x0f, 0xa0, 0x67, 0x94, 0x93, 0xd8, 0x2c, 0xae, 0xab,
	0x81, 0x8b, 0x18, 0x9f, 0x81, 0xbb, 0xcc, 0xcb, 0x4c, 0x70, 0x75, 0x9d, 0xbc, 0xd9, 0xbb, 0x9b,
	0xf3, 0x9d, 0xbd, 0x54, 0xe1, 0xef, 0x33, 0xc1, 0x56, 0x81, 0xe1, 0xe2, 0x53, 0xe8, 0x44, 0xa5,
	0xb8, 0xd1, 0x17, 0xcc, 0x9b, 0x3d, 0xd8, 0x4a, 0xfa, 0x46, 0x46, 0x75, 0x8e, 0x66, 0xaa, 0x83,
	0xca, 0xd3, 0x64, 0x91, 0x90, 0xed, 0x56, 0xd5, 0x78, 0xfa, 0x02, 0xbc, 0xc6, 0x2c, 0x3b, 0xea,
	0xf7, 0x23, 0xe8, 0xdc, 0x45, 0x69, 0x49, 0x6a, 0x9f, 0xde, 0x6c, 0x28, 0xe7, 0x53, 0x19, 0x3f,
	0x97, 0xa2, 0x28, 0x45, 0xa0, 0xa3, 0xe7, 0xad, 0xaf, 0x9d, 0xe9, 0x4b, 0x80, 0x7a, 0xf2, 0x1d,
	0x52, 0x0f, 0xd7, 0xa5, 0x0e, 0xa5, 0x94, 0x4c, 0xd8, 0x2d, 0xe7, 0x33, 0xb3, 0x34, 0x1d, 0x91,
	0x6d, 0x49, 0xac, 0x0a, 0x32, 0x82, 0xea, 0x5b, 0xd6, 0x79, 0x4c, 0x7c, 0xc1, 0x92, 0x42, 0xc8,
	0x52, 0xd3, 0x47, 0xd1, 0x84, 0xf0, 0x09, 0xb8, 0xda, 0x70, 0x55, 0xb7, 0xde, 0xec, 0xb8, 0x5a,
	0xbf, 0x36, 0xcd, 0x4c, 0x6b, 0x48, 0xfe, 0x02, 0xc6, 0x5b, 0x41, 0xd9, 0x7a, 0x63, 0x7a, 0x25,
	0x1f, 0xba, 0x30, 0xa5, 0x88, 0x53, 0x28, 0x44, 0x6a, 0x9e, 0x99, 0xa1, 0x09, 0xfc, 0x28, 0xf1,
	0x2b, 0x91, 0xa2, 0x0f, 0x83, 0x65, 0xf4, 0xba, 0xc1, 0xd3, 0x97, 0xc7, 0x5b, 0x46, 0xaf, 0x2d,
	0xc7, 0x2f, 0x61, 0xb8, 0xb1, 0xed, 0xff, 0xb9, 0xb9, 0x4f, 0x36, 0x36, 0x77, 0x64, 0x1d, 0xdd,
	0xb9, 0xb7, 0x39, 0x8c, 0x36, 0x63, 0x6f, 0x7b, 0x6b, 0xb3, 0xbf, 0x5b, 0xd0, 0xf9, 0x4d, 0x66,
	0xe1, 0x73, 0x80, 0xfa, 0xcd, 0x47, 0x65, 0xfb, 0xd6, 0x1f, 0x83, 0xe9, 0xfd, 0x4d, 0x58, 0xdf,
	0x48, 0xff, 0x1e, 0x3e, 0x86, 0x3d, 0x89, 0xe3, 0xd0, 0x32, 0x6c, 0xca, 0xa8, 0x06, 0x2a, 0xf2,
	0xf3, 0xb5, 0x5e, 0x76, 0xbc, 0xd1, 0x34, 0x9b, 0x73, 0x6d, 0x3f, 0x03, 0xfe, 0x3d, 0x7c, 0x0a,
	0xae, 0x6e, 0xc6, 0x38, 0x96, 0x9c, 0xb5, 0x8e, 0x3e, 0xc5, 0x26, 0x54, 0xa5, 0x9c, 0x43, 0xaf,
	0xea, 0x23, 0x78, 0x54, 0xdf, 0xc1, 0xba, 0x25, 0x4d, 0x8f, 0x37, 0x50, 0x9b, 0x3b, 0x77, 0xd5,
	0xbf, 0xa5, 0xcf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x17, 0xba, 0x0a, 0xdd, 0x3c, 0x09, 0x00,
	0x00,
}
