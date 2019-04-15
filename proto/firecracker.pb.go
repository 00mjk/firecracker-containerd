// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firecracker.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FifoType int32

const (
	FifoType_LOG     FifoType = 0
	FifoType_METRICS FifoType = 1
)

var FifoType_name = map[int32]string{
	0: "LOG",
	1: "METRICS",
}
var FifoType_value = map[string]int32{
	"LOG":     0,
	"METRICS": 1,
}

func (x FifoType) String() string {
	return proto.EnumName(FifoType_name, int32(x))
}
func (FifoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{0}
}

// CreateVMRequest specifies creation parameters for a new FC instance
type CreateVMRequest struct {
	// Specifies the machine configuration for the VM
	MachineCfg *FirecrackerMachineConfiguration `protobuf:"bytes,1,opt,name=MachineCfg,proto3" json:"MachineCfg,omitempty"`
	// Specifies the file path where the kernel image is located
	KernelImagePath string `protobuf:"bytes,2,opt,name=KernelImagePath,proto3" json:"KernelImagePath,omitempty"`
	// Specifies the commandline arguments that should be passed to the kernel
	KernelArgs string `protobuf:"bytes,3,opt,name=KernelArgs,proto3" json:"KernelArgs,omitempty"`
	// Specifies the root block device for the VM
	RootDrive *FirecrackerDrive `protobuf:"bytes,4,opt,name=RootDrive,proto3" json:"RootDrive,omitempty"`
	// Specifies the additional block device config for the VM.
	AdditionalDrives []*FirecrackerDrive `protobuf:"bytes,5,rep,name=AdditionalDrives,proto3" json:"AdditionalDrives,omitempty"`
	// Specifies the networking configuration for a VM
	NetworkInterfaces []*FirecrackerNetworkInterface `protobuf:"bytes,6,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	// The number of dummy drives to reserve in advance before running FC instance.
	ContainerCount       int32    `protobuf:"varint,7,opt,name=ContainerCount,proto3" json:"ContainerCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMRequest) Reset()         { *m = CreateVMRequest{} }
func (m *CreateVMRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVMRequest) ProtoMessage()    {}
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{0}
}
func (m *CreateVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMRequest.Unmarshal(m, b)
}
func (m *CreateVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMRequest.Merge(dst, src)
}
func (m *CreateVMRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVMRequest.Size(m)
}
func (m *CreateVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMRequest proto.InternalMessageInfo

func (m *CreateVMRequest) GetMachineCfg() *FirecrackerMachineConfiguration {
	if m != nil {
		return m.MachineCfg
	}
	return nil
}

func (m *CreateVMRequest) GetKernelImagePath() string {
	if m != nil {
		return m.KernelImagePath
	}
	return ""
}

func (m *CreateVMRequest) GetKernelArgs() string {
	if m != nil {
		return m.KernelArgs
	}
	return ""
}

func (m *CreateVMRequest) GetRootDrive() *FirecrackerDrive {
	if m != nil {
		return m.RootDrive
	}
	return nil
}

func (m *CreateVMRequest) GetAdditionalDrives() []*FirecrackerDrive {
	if m != nil {
		return m.AdditionalDrives
	}
	return nil
}

func (m *CreateVMRequest) GetNetworkInterfaces() []*FirecrackerNetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *CreateVMRequest) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

type CreateVMResponse struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMResponse) Reset()         { *m = CreateVMResponse{} }
func (m *CreateVMResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVMResponse) ProtoMessage()    {}
func (*CreateVMResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{1}
}
func (m *CreateVMResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMResponse.Unmarshal(m, b)
}
func (m *CreateVMResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMResponse.Marshal(b, m, deterministic)
}
func (dst *CreateVMResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMResponse.Merge(dst, src)
}
func (m *CreateVMResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVMResponse.Size(m)
}
func (m *CreateVMResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMResponse proto.InternalMessageInfo

func (m *CreateVMResponse) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type StopVMRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopVMRequest) Reset()         { *m = StopVMRequest{} }
func (m *StopVMRequest) String() string { return proto.CompactTextString(m) }
func (*StopVMRequest) ProtoMessage()    {}
func (*StopVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{2}
}
func (m *StopVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopVMRequest.Unmarshal(m, b)
}
func (m *StopVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopVMRequest.Marshal(b, m, deterministic)
}
func (dst *StopVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopVMRequest.Merge(dst, src)
}
func (m *StopVMRequest) XXX_Size() int {
	return xxx_messageInfo_StopVMRequest.Size(m)
}
func (m *StopVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopVMRequest proto.InternalMessageInfo

func (m *StopVMRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMAddressRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMAddressRequest) Reset()         { *m = GetVMAddressRequest{} }
func (m *GetVMAddressRequest) String() string { return proto.CompactTextString(m) }
func (*GetVMAddressRequest) ProtoMessage()    {}
func (*GetVMAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{3}
}
func (m *GetVMAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMAddressRequest.Unmarshal(m, b)
}
func (m *GetVMAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMAddressRequest.Marshal(b, m, deterministic)
}
func (dst *GetVMAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMAddressRequest.Merge(dst, src)
}
func (m *GetVMAddressRequest) XXX_Size() int {
	return xxx_messageInfo_GetVMAddressRequest.Size(m)
}
func (m *GetVMAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMAddressRequest proto.InternalMessageInfo

func (m *GetVMAddressRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type GetVMAddressResponse struct {
	SocketPath           string   `protobuf:"bytes,1,opt,name=SocketPath,proto3" json:"SocketPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMAddressResponse) Reset()         { *m = GetVMAddressResponse{} }
func (m *GetVMAddressResponse) String() string { return proto.CompactTextString(m) }
func (*GetVMAddressResponse) ProtoMessage()    {}
func (*GetVMAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{4}
}
func (m *GetVMAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMAddressResponse.Unmarshal(m, b)
}
func (m *GetVMAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMAddressResponse.Marshal(b, m, deterministic)
}
func (dst *GetVMAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMAddressResponse.Merge(dst, src)
}
func (m *GetVMAddressResponse) XXX_Size() int {
	return xxx_messageInfo_GetVMAddressResponse.Size(m)
}
func (m *GetVMAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMAddressResponse proto.InternalMessageInfo

func (m *GetVMAddressResponse) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

type GetFifoPathRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	FifoType             FifoType `protobuf:"varint,2,opt,name=FifoType,proto3,enum=FifoType" json:"FifoType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFifoPathRequest) Reset()         { *m = GetFifoPathRequest{} }
func (m *GetFifoPathRequest) String() string { return proto.CompactTextString(m) }
func (*GetFifoPathRequest) ProtoMessage()    {}
func (*GetFifoPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{5}
}
func (m *GetFifoPathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFifoPathRequest.Unmarshal(m, b)
}
func (m *GetFifoPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFifoPathRequest.Marshal(b, m, deterministic)
}
func (dst *GetFifoPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFifoPathRequest.Merge(dst, src)
}
func (m *GetFifoPathRequest) XXX_Size() int {
	return xxx_messageInfo_GetFifoPathRequest.Size(m)
}
func (m *GetFifoPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFifoPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFifoPathRequest proto.InternalMessageInfo

func (m *GetFifoPathRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *GetFifoPathRequest) GetFifoType() FifoType {
	if m != nil {
		return m.FifoType
	}
	return FifoType_LOG
}

type GetFifoPathResponse struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFifoPathResponse) Reset()         { *m = GetFifoPathResponse{} }
func (m *GetFifoPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetFifoPathResponse) ProtoMessage()    {}
func (*GetFifoPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{6}
}
func (m *GetFifoPathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFifoPathResponse.Unmarshal(m, b)
}
func (m *GetFifoPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFifoPathResponse.Marshal(b, m, deterministic)
}
func (dst *GetFifoPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFifoPathResponse.Merge(dst, src)
}
func (m *GetFifoPathResponse) XXX_Size() int {
	return xxx_messageInfo_GetFifoPathResponse.Size(m)
}
func (m *GetFifoPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFifoPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFifoPathResponse proto.InternalMessageInfo

func (m *GetFifoPathResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SetVMMetadataRequest struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,proto3" json:"VMID,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetVMMetadataRequest) Reset()         { *m = SetVMMetadataRequest{} }
func (m *SetVMMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*SetVMMetadataRequest) ProtoMessage()    {}
func (*SetVMMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firecracker_f0b40e152181d1b0, []int{7}
}
func (m *SetVMMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetVMMetadataRequest.Unmarshal(m, b)
}
func (m *SetVMMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetVMMetadataRequest.Marshal(b, m, deterministic)
}
func (dst *SetVMMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetVMMetadataRequest.Merge(dst, src)
}
func (m *SetVMMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_SetVMMetadataRequest.Size(m)
}
func (m *SetVMMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetVMMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetVMMetadataRequest proto.InternalMessageInfo

func (m *SetVMMetadataRequest) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func (m *SetVMMetadataRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVMRequest)(nil), "CreateVMRequest")
	proto.RegisterType((*CreateVMResponse)(nil), "CreateVMResponse")
	proto.RegisterType((*StopVMRequest)(nil), "StopVMRequest")
	proto.RegisterType((*GetVMAddressRequest)(nil), "GetVMAddressRequest")
	proto.RegisterType((*GetVMAddressResponse)(nil), "GetVMAddressResponse")
	proto.RegisterType((*GetFifoPathRequest)(nil), "GetFifoPathRequest")
	proto.RegisterType((*GetFifoPathResponse)(nil), "GetFifoPathResponse")
	proto.RegisterType((*SetVMMetadataRequest)(nil), "SetVMMetadataRequest")
	proto.RegisterEnum("FifoType", FifoType_name, FifoType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirecrackerClient is the client API for Firecracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirecrackerClient interface {
	// Runs new Firecracker VM instance
	CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*CreateVMResponse, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets VM's instance socket file location
	GetVMAddress(ctx context.Context, in *GetVMAddressRequest, opts ...grpc.CallOption) (*GetVMAddressResponse, error)
	// Gets VM's instance FIFO file location
	GetFifoPath(ctx context.Context, in *GetFifoPathRequest, opts ...grpc.CallOption) (*GetFifoPathResponse, error)
	// Sets VM's instance metadata
	SetVMMetadata(ctx context.Context, in *SetVMMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type firecrackerClient struct {
	cc *grpc.ClientConn
}

func NewFirecrackerClient(cc *grpc.ClientConn) FirecrackerClient {
	return &firecrackerClient{cc}
}

func (c *firecrackerClient) CreateVM(ctx context.Context, in *CreateVMRequest, opts ...grpc.CallOption) (*CreateVMResponse, error) {
	out := new(CreateVMResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) StopVM(ctx context.Context, in *StopVMRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/StopVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) GetVMAddress(ctx context.Context, in *GetVMAddressRequest, opts ...grpc.CallOption) (*GetVMAddressResponse, error) {
	out := new(GetVMAddressResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/GetVMAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) GetFifoPath(ctx context.Context, in *GetFifoPathRequest, opts ...grpc.CallOption) (*GetFifoPathResponse, error) {
	out := new(GetFifoPathResponse)
	err := c.cc.Invoke(ctx, "/Firecracker/GetFifoPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firecrackerClient) SetVMMetadata(ctx context.Context, in *SetVMMetadataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Firecracker/SetVMMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirecrackerServer is the server API for Firecracker service.
type FirecrackerServer interface {
	// Runs new Firecracker VM instance
	CreateVM(context.Context, *CreateVMRequest) (*CreateVMResponse, error)
	// Stops existing Firecracker instance by VM ID
	StopVM(context.Context, *StopVMRequest) (*empty.Empty, error)
	// Gets VM's instance socket file location
	GetVMAddress(context.Context, *GetVMAddressRequest) (*GetVMAddressResponse, error)
	// Gets VM's instance FIFO file location
	GetFifoPath(context.Context, *GetFifoPathRequest) (*GetFifoPathResponse, error)
	// Sets VM's instance metadata
	SetVMMetadata(context.Context, *SetVMMetadataRequest) (*empty.Empty, error)
}

func RegisterFirecrackerServer(s *grpc.Server, srv FirecrackerServer) {
	s.RegisterService(&_Firecracker_serviceDesc, srv)
}

func _Firecracker_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).CreateVM(ctx, req.(*CreateVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_StopVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).StopVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/StopVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).StopVM(ctx, req.(*StopVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_GetVMAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVMAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).GetVMAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/GetVMAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).GetVMAddress(ctx, req.(*GetVMAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_GetFifoPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFifoPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).GetFifoPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/GetFifoPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).GetFifoPath(ctx, req.(*GetFifoPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Firecracker_SetVMMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVMMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirecrackerServer).SetVMMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Firecracker/SetVMMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirecrackerServer).SetVMMetadata(ctx, req.(*SetVMMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Firecracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Firecracker",
	HandlerType: (*FirecrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVM",
			Handler:    _Firecracker_CreateVM_Handler,
		},
		{
			MethodName: "StopVM",
			Handler:    _Firecracker_StopVM_Handler,
		},
		{
			MethodName: "GetVMAddress",
			Handler:    _Firecracker_GetVMAddress_Handler,
		},
		{
			MethodName: "GetFifoPath",
			Handler:    _Firecracker_GetFifoPath_Handler,
		},
		{
			MethodName: "SetVMMetadata",
			Handler:    _Firecracker_SetVMMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firecracker.proto",
}

func init() { proto.RegisterFile("firecracker.proto", fileDescriptor_firecracker_f0b40e152181d1b0) }

var fileDescriptor_firecracker_f0b40e152181d1b0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0x12, 0x41,
	0x14, 0x95, 0xd2, 0xf2, 0x71, 0xb1, 0x14, 0xa6, 0xd4, 0x6c, 0x56, 0xd3, 0x6c, 0x30, 0x36, 0xe8,
	0xc3, 0x60, 0x30, 0x31, 0x26, 0x46, 0x23, 0xd2, 0x42, 0x50, 0xd7, 0x9a, 0xa1, 0xe1, 0xc1, 0xb7,
	0x11, 0xee, 0x6e, 0x37, 0xd0, 0x9d, 0x75, 0x76, 0xd0, 0xf0, 0xe8, 0xaf, 0xf1, 0x6f, 0x1a, 0x66,
	0xd9, 0xee, 0x02, 0x2b, 0x4f, 0xec, 0x3d, 0xf7, 0x5c, 0xee, 0xc9, 0x39, 0x73, 0xa1, 0xee, 0x78,
	0x12, 0x27, 0x92, 0x4f, 0x66, 0x28, 0x69, 0x20, 0x85, 0x12, 0xe6, 0x63, 0x57, 0x08, 0x77, 0x8e,
	0x6d, 0x5d, 0xfd, 0x58, 0x38, 0x6d, 0xbc, 0x0b, 0xd4, 0x72, 0xdd, 0xac, 0xa8, 0x65, 0x80, 0x61,
	0x54, 0x34, 0xff, 0xe4, 0xe1, 0xa4, 0x27, 0x91, 0x2b, 0x1c, 0xdb, 0x0c, 0x7f, 0x2e, 0x30, 0x54,
	0xe4, 0x03, 0x80, 0xcd, 0x27, 0xb7, 0x9e, 0x8f, 0x3d, 0xc7, 0x35, 0x72, 0x56, 0xae, 0x55, 0xe9,
	0x58, 0xb4, 0x9f, 0x6c, 0x89, 0xbb, 0xc2, 0x77, 0x3c, 0x77, 0x21, 0xb9, 0xf2, 0x84, 0xcf, 0x52,
	0x33, 0xa4, 0x05, 0x27, 0x9f, 0x51, 0xfa, 0x38, 0x1f, 0xde, 0x71, 0x17, 0xbf, 0x71, 0x75, 0x6b,
	0x1c, 0x58, 0xb9, 0x56, 0x99, 0x6d, 0xc3, 0xe4, 0x1c, 0x20, 0x82, 0xba, 0xd2, 0x0d, 0x8d, 0xbc,
	0x26, 0xa5, 0x10, 0xd2, 0x86, 0x32, 0x13, 0x42, 0x5d, 0x4a, 0xef, 0x17, 0x1a, 0x87, 0x5a, 0x4a,
	0x3d, 0x2d, 0x45, 0x37, 0x58, 0xc2, 0x21, 0xef, 0xa0, 0xd6, 0x9d, 0x4e, 0xbd, 0x95, 0x24, 0x3e,
	0xd7, 0x50, 0x68, 0x1c, 0x59, 0xf9, 0xec, 0xb9, 0x1d, 0x2a, 0xf9, 0x04, 0xf5, 0xaf, 0xa8, 0x7e,
	0x0b, 0x39, 0x1b, 0xfa, 0x0a, 0xa5, 0xc3, 0x27, 0x18, 0x1a, 0x05, 0x3d, 0xff, 0x24, 0x3d, 0xbf,
	0x4d, 0x62, 0xbb, 0x63, 0xe4, 0x02, 0xaa, 0x3d, 0xe1, 0x2b, 0xee, 0xf9, 0x28, 0x7b, 0x62, 0xe1,
	0x2b, 0xa3, 0x68, 0xe5, 0x5a, 0x47, 0x6c, 0x0b, 0x6d, 0x5e, 0x40, 0x2d, 0x89, 0x20, 0x0c, 0x84,
	0x1f, 0x22, 0x21, 0x70, 0x38, 0xb6, 0x87, 0x97, 0xda, 0xfd, 0x32, 0xd3, 0xdf, 0xcd, 0xa7, 0x70,
	0x3c, 0x52, 0x22, 0x48, 0x82, 0xca, 0x22, 0x3d, 0x87, 0xd3, 0x01, 0xaa, 0xb1, 0xdd, 0x9d, 0x4e,
	0x25, 0x86, 0xe1, 0x3e, 0xea, 0x6b, 0x68, 0x6c, 0x52, 0xd7, 0xbb, 0xcf, 0x01, 0x46, 0x62, 0x32,
	0x43, 0xa5, 0x83, 0x8b, 0x26, 0x52, 0x48, 0xf3, 0x1a, 0xc8, 0x00, 0x55, 0xdf, 0x73, 0xc4, 0xaa,
	0xdc, 0xb3, 0x81, 0x3c, 0x83, 0xd2, 0x8a, 0x76, 0xb3, 0x0c, 0x50, 0x3f, 0x80, 0x6a, 0xa7, 0x4c,
	0x63, 0x80, 0xdd, 0xb7, 0xd6, 0x9a, 0x93, 0x3f, 0x4c, 0x3c, 0x48, 0x29, 0xd0, 0xdf, 0xcd, 0x3e,
	0x34, 0x46, 0x2b, 0xcd, 0x36, 0x2a, 0x3e, 0xe5, 0x8a, 0xef, 0xdb, 0x6e, 0x42, 0x29, 0xa6, 0xad,
	0x9f, 0xdf, 0x7d, 0xfd, 0xc2, 0x4a, 0x94, 0x91, 0x22, 0xe4, 0xbf, 0x5c, 0x0f, 0x6a, 0x0f, 0x48,
	0x05, 0x8a, 0xf6, 0xd5, 0x0d, 0x1b, 0xf6, 0x46, 0xb5, 0x5c, 0xe7, 0xef, 0x01, 0x54, 0x52, 0x81,
	0x93, 0x36, 0x94, 0xe2, 0x94, 0x48, 0x8d, 0x6e, 0xdd, 0x8c, 0x59, 0xa7, 0x3b, 0x11, 0xbe, 0x84,
	0x42, 0x14, 0x17, 0xa9, 0xd2, 0x8d, 0xdc, 0xcc, 0x47, 0x34, 0xba, 0x4f, 0x1a, 0xdf, 0x27, 0xbd,
	0x5a, 0xdd, 0x27, 0x79, 0x0b, 0x0f, 0xd3, 0x81, 0x90, 0x06, 0xcd, 0x88, 0xd2, 0x3c, 0xa3, 0x99,
	0xa9, 0xbd, 0x81, 0x4a, 0xca, 0x44, 0x72, 0x4a, 0x77, 0x33, 0x32, 0x1b, 0x34, 0xcb, 0xe7, 0xf7,
	0x70, 0xbc, 0xe1, 0x29, 0x39, 0xa3, 0x59, 0x1e, 0xff, 0x4f, 0xf6, 0xc7, 0xe2, 0xf7, 0xa3, 0x08,
	0x29, 0xe8, 0x9f, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xd7, 0x4f, 0x65, 0x92, 0x04,
	0x00, 0x00,
}