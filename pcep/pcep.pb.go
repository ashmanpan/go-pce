// Code generated by protoc-gen-go.
// source: pcep.proto
// DO NOT EDIT!

/*
Package pcep is a generated protocol buffer package.

It is generated from these files:
	pcep.proto

It has these top-level messages:
	LspRro
	LspEro
	LspAttributes
	PbLsp
	LspEvent
	LspDumpRequest
	LspDumpResponse
	PccRouter
	UpdateLspRequest
	LspResponse
	CreateLspRequest
	DeleteLspRequest
	PccRouters
	Empty
*/
package pcep

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

type LspEventType int32

const (
	LspEventType_LSP_UP         LspEventType = 0
	LspEventType_LSP_DOWN       LspEventType = 1
	LspEventType_LSP_UPDATE     LspEventType = 2
	LspEventType_LSP_DELEGATION LspEventType = 3
	LspEventType_LSP_REVOCATION LspEventType = 4
)

var LspEventType_name = map[int32]string{
	0: "LSP_UP",
	1: "LSP_DOWN",
	2: "LSP_UPDATE",
	3: "LSP_DELEGATION",
	4: "LSP_REVOCATION",
}
var LspEventType_value = map[string]int32{
	"LSP_UP":         0,
	"LSP_DOWN":       1,
	"LSP_UPDATE":     2,
	"LSP_DELEGATION": 3,
	"LSP_REVOCATION": 4,
}

func (x LspEventType) String() string {
	return proto.EnumName(LspEventType_name, int32(x))
}
func (LspEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PccState int32

const (
	PccState_IDLE PccState = 0
	PccState_UP   PccState = 1
)

var PccState_name = map[int32]string{
	0: "IDLE",
	1: "UP",
}
var PccState_value = map[string]int32{
	"IDLE": 0,
	"UP":   1,
}

func (x PccState) String() string {
	return proto.EnumName(PccState_name, int32(x))
}
func (PccState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LspRro struct {
	HopIP                string `protobuf:"bytes,1,opt,name=hopIP" json:"hopIP,omitempty"`
	RecordedLabel        int32  `protobuf:"varint,2,opt,name=recordedLabel" json:"recordedLabel,omitempty"`
	LocalProtectionInUse bool   `protobuf:"varint,3,opt,name=localProtectionInUse" json:"localProtectionInUse,omitempty"`
}

func (m *LspRro) Reset()                    { *m = LspRro{} }
func (m *LspRro) String() string            { return proto.CompactTextString(m) }
func (*LspRro) ProtoMessage()               {}
func (*LspRro) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LspEro struct {
	Loose bool   `protobuf:"varint,1,opt,name=loose" json:"loose,omitempty"`
	HopIP string `protobuf:"bytes,2,opt,name=hopIP" json:"hopIP,omitempty"`
}

func (m *LspEro) Reset()                    { *m = LspEro{} }
func (m *LspEro) String() string            { return proto.CompactTextString(m) }
func (*LspEro) ProtoMessage()               {}
func (*LspEro) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LspAttributes struct {
	SetupPriority     int32     `protobuf:"varint,1,opt,name=setupPriority" json:"setupPriority,omitempty"`
	HoldPriority      int32     `protobuf:"varint,2,opt,name=holdPriority" json:"holdPriority,omitempty"`
	ReservedBandwidth int64     `protobuf:"varint,3,opt,name=reservedBandwidth" json:"reservedBandwidth,omitempty"`
	MaxAvgBandwidth   int64     `protobuf:"varint,4,opt,name=maxAvgBandwidth" json:"maxAvgBandwidth,omitempty"`
	PathMetric        int32     `protobuf:"varint,5,opt,name=pathMetric" json:"pathMetric,omitempty"`
	ReportedPath      []*LspRro `protobuf:"bytes,6,rep,name=reportedPath" json:"reportedPath,omitempty"`
	ExplicitPath      []*LspEro `protobuf:"bytes,7,rep,name=explicitPath" json:"explicitPath,omitempty"`
	ExcludeAny        []int32   `protobuf:"varint,8,rep,packed,name=excludeAny" json:"excludeAny,omitempty"`
}

func (m *LspAttributes) Reset()                    { *m = LspAttributes{} }
func (m *LspAttributes) String() string            { return proto.CompactTextString(m) }
func (*LspAttributes) ProtoMessage()               {}
func (*LspAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LspAttributes) GetReportedPath() []*LspRro {
	if m != nil {
		return m.ReportedPath
	}
	return nil
}

func (m *LspAttributes) GetExplicitPath() []*LspEro {
	if m != nil {
		return m.ExplicitPath
	}
	return nil
}

type PbLsp struct {
	LspID         int32          `protobuf:"varint,1,opt,name=lspID" json:"lspID,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SourceIP      string         `protobuf:"bytes,3,opt,name=sourceIP" json:"sourceIP,omitempty"`
	DestinationIP string         `protobuf:"bytes,4,opt,name=destinationIP" json:"destinationIP,omitempty"`
	Active        bool           `protobuf:"varint,5,opt,name=active" json:"active,omitempty"`
	Attrs         *LspAttributes `protobuf:"bytes,6,opt,name=attrs" json:"attrs,omitempty"`
}

func (m *PbLsp) Reset()                    { *m = PbLsp{} }
func (m *PbLsp) String() string            { return proto.CompactTextString(m) }
func (*PbLsp) ProtoMessage()               {}
func (*PbLsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PbLsp) GetAttrs() *LspAttributes {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type LspEvent struct {
	Type     LspEventType `protobuf:"varint,1,opt,name=type,enum=pcep.LspEventType" json:"type,omitempty"`
	RouterIP string       `protobuf:"bytes,2,opt,name=routerIP" json:"routerIP,omitempty"`
	Lsp      *PbLsp       `protobuf:"bytes,3,opt,name=lsp" json:"lsp,omitempty"`
}

func (m *LspEvent) Reset()                    { *m = LspEvent{} }
func (m *LspEvent) String() string            { return proto.CompactTextString(m) }
func (*LspEvent) ProtoMessage()               {}
func (*LspEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LspEvent) GetLsp() *PbLsp {
	if m != nil {
		return m.Lsp
	}
	return nil
}

type LspDumpRequest struct {
	RouterIP string `protobuf:"bytes,1,opt,name=routerIP" json:"routerIP,omitempty"`
	Filters  int32  `protobuf:"varint,2,opt,name=filters" json:"filters,omitempty"`
}

func (m *LspDumpRequest) Reset()                    { *m = LspDumpRequest{} }
func (m *LspDumpRequest) String() string            { return proto.CompactTextString(m) }
func (*LspDumpRequest) ProtoMessage()               {}
func (*LspDumpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type LspDumpResponse struct {
	Lsps []*LspEvent `protobuf:"bytes,1,rep,name=lsps" json:"lsps,omitempty"`
}

func (m *LspDumpResponse) Reset()                    { *m = LspDumpResponse{} }
func (m *LspDumpResponse) String() string            { return proto.CompactTextString(m) }
func (*LspDumpResponse) ProtoMessage()               {}
func (*LspDumpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LspDumpResponse) GetLsps() []*LspEvent {
	if m != nil {
		return m.Lsps
	}
	return nil
}

type PccRouter struct {
	Name           string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	State          PccState `protobuf:"varint,2,opt,name=state,enum=pcep.PccState" json:"state,omitempty"`
	IsStateful     bool     `protobuf:"varint,3,opt,name=isStateful" json:"isStateful,omitempty"`
	SupportsCreate bool     `protobuf:"varint,4,opt,name=supportsCreate" json:"supportsCreate,omitempty"`
	SessionID      int32    `protobuf:"varint,5,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *PccRouter) Reset()                    { *m = PccRouter{} }
func (m *PccRouter) String() string            { return proto.CompactTextString(m) }
func (*PccRouter) ProtoMessage()               {}
func (*PccRouter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type UpdateLspRequest struct {
	Router string `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
	// map of LspID -> changed atttributes for resignaling.
	// Caller needs to include both changed and unchanged(original)
	// attribute values.
	LspChanges map[int32]*LspAttributes `protobuf:"bytes,2,rep,name=lspChanges" json:"lspChanges,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UpdateLspRequest) Reset()                    { *m = UpdateLspRequest{} }
func (m *UpdateLspRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLspRequest) ProtoMessage()               {}
func (*UpdateLspRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateLspRequest) GetLspChanges() map[int32]*LspAttributes {
	if m != nil {
		return m.LspChanges
	}
	return nil
}

type LspResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *LspResponse) Reset()                    { *m = LspResponse{} }
func (m *LspResponse) String() string            { return proto.CompactTextString(m) }
func (*LspResponse) ProtoMessage()               {}
func (*LspResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type CreateLspRequest struct {
	Router       string   `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
	LspsToCreate []*PbLsp `protobuf:"bytes,2,rep,name=lspsToCreate" json:"lspsToCreate,omitempty"`
}

func (m *CreateLspRequest) Reset()                    { *m = CreateLspRequest{} }
func (m *CreateLspRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLspRequest) ProtoMessage()               {}
func (*CreateLspRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateLspRequest) GetLspsToCreate() []*PbLsp {
	if m != nil {
		return m.LspsToCreate
	}
	return nil
}

type DeleteLspRequest struct {
	Router string  `protobuf:"bytes,1,opt,name=router" json:"router,omitempty"`
	LspIDs []int32 `protobuf:"varint,3,rep,packed,name=lspIDs" json:"lspIDs,omitempty"`
}

func (m *DeleteLspRequest) Reset()                    { *m = DeleteLspRequest{} }
func (m *DeleteLspRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLspRequest) ProtoMessage()               {}
func (*DeleteLspRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type PccRouters struct {
	Routers []*PccRouter `protobuf:"bytes,1,rep,name=routers" json:"routers,omitempty"`
}

func (m *PccRouters) Reset()                    { *m = PccRouters{} }
func (m *PccRouters) String() string            { return proto.CompactTextString(m) }
func (*PccRouters) ProtoMessage()               {}
func (*PccRouters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PccRouters) GetRouters() []*PccRouter {
	if m != nil {
		return m.Routers
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*LspRro)(nil), "pcep.LspRro")
	proto.RegisterType((*LspEro)(nil), "pcep.LspEro")
	proto.RegisterType((*LspAttributes)(nil), "pcep.LspAttributes")
	proto.RegisterType((*PbLsp)(nil), "pcep.PbLsp")
	proto.RegisterType((*LspEvent)(nil), "pcep.LspEvent")
	proto.RegisterType((*LspDumpRequest)(nil), "pcep.LspDumpRequest")
	proto.RegisterType((*LspDumpResponse)(nil), "pcep.LspDumpResponse")
	proto.RegisterType((*PccRouter)(nil), "pcep.PccRouter")
	proto.RegisterType((*UpdateLspRequest)(nil), "pcep.UpdateLspRequest")
	proto.RegisterType((*LspResponse)(nil), "pcep.LspResponse")
	proto.RegisterType((*CreateLspRequest)(nil), "pcep.CreateLspRequest")
	proto.RegisterType((*DeleteLspRequest)(nil), "pcep.DeleteLspRequest")
	proto.RegisterType((*PccRouters)(nil), "pcep.PccRouters")
	proto.RegisterType((*Empty)(nil), "pcep.Empty")
	proto.RegisterEnum("pcep.LspEventType", LspEventType_name, LspEventType_value)
	proto.RegisterEnum("pcep.PccState", PccState_name, PccState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PceService service

type PceServiceClient interface {
	// get a full current lsp dump for a given router
	GetLspDump(ctx context.Context, in *LspDumpRequest, opts ...grpc.CallOption) (*LspDumpResponse, error)
	// get a list of active PCC routers
	GetPccRouters(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PccRouters, error)
	UpdateLsps(ctx context.Context, in *UpdateLspRequest, opts ...grpc.CallOption) (*LspResponse, error)
	// create new LSPs on a given router
	CreateLsps(ctx context.Context, in *CreateLspRequest, opts ...grpc.CallOption) (*LspResponse, error)
	// delete previously created LSPs on a given router
	DeleteLsps(ctx context.Context, in *DeleteLspRequest, opts ...grpc.CallOption) (*LspResponse, error)
}

type pceServiceClient struct {
	cc *grpc.ClientConn
}

func NewPceServiceClient(cc *grpc.ClientConn) PceServiceClient {
	return &pceServiceClient{cc}
}

func (c *pceServiceClient) GetLspDump(ctx context.Context, in *LspDumpRequest, opts ...grpc.CallOption) (*LspDumpResponse, error) {
	out := new(LspDumpResponse)
	err := grpc.Invoke(ctx, "/pcep.PceService/GetLspDump", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pceServiceClient) GetPccRouters(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PccRouters, error) {
	out := new(PccRouters)
	err := grpc.Invoke(ctx, "/pcep.PceService/GetPccRouters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pceServiceClient) UpdateLsps(ctx context.Context, in *UpdateLspRequest, opts ...grpc.CallOption) (*LspResponse, error) {
	out := new(LspResponse)
	err := grpc.Invoke(ctx, "/pcep.PceService/UpdateLsps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pceServiceClient) CreateLsps(ctx context.Context, in *CreateLspRequest, opts ...grpc.CallOption) (*LspResponse, error) {
	out := new(LspResponse)
	err := grpc.Invoke(ctx, "/pcep.PceService/CreateLsps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pceServiceClient) DeleteLsps(ctx context.Context, in *DeleteLspRequest, opts ...grpc.CallOption) (*LspResponse, error) {
	out := new(LspResponse)
	err := grpc.Invoke(ctx, "/pcep.PceService/DeleteLsps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PceService service

type PceServiceServer interface {
	// get a full current lsp dump for a given router
	GetLspDump(context.Context, *LspDumpRequest) (*LspDumpResponse, error)
	// get a list of active PCC routers
	GetPccRouters(context.Context, *Empty) (*PccRouters, error)
	UpdateLsps(context.Context, *UpdateLspRequest) (*LspResponse, error)
	// create new LSPs on a given router
	CreateLsps(context.Context, *CreateLspRequest) (*LspResponse, error)
	// delete previously created LSPs on a given router
	DeleteLsps(context.Context, *DeleteLspRequest) (*LspResponse, error)
}

func RegisterPceServiceServer(s *grpc.Server, srv PceServiceServer) {
	s.RegisterService(&_PceService_serviceDesc, srv)
}

func _PceService_GetLspDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LspDumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PceServiceServer).GetLspDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcep.PceService/GetLspDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PceServiceServer).GetLspDump(ctx, req.(*LspDumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PceService_GetPccRouters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PceServiceServer).GetPccRouters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcep.PceService/GetPccRouters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PceServiceServer).GetPccRouters(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PceService_UpdateLsps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLspRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PceServiceServer).UpdateLsps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcep.PceService/UpdateLsps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PceServiceServer).UpdateLsps(ctx, req.(*UpdateLspRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PceService_CreateLsps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLspRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PceServiceServer).CreateLsps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcep.PceService/CreateLsps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PceServiceServer).CreateLsps(ctx, req.(*CreateLspRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PceService_DeleteLsps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLspRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PceServiceServer).DeleteLsps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcep.PceService/DeleteLsps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PceServiceServer).DeleteLsps(ctx, req.(*DeleteLspRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pcep.PceService",
	HandlerType: (*PceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLspDump",
			Handler:    _PceService_GetLspDump_Handler,
		},
		{
			MethodName: "GetPccRouters",
			Handler:    _PceService_GetPccRouters_Handler,
		},
		{
			MethodName: "UpdateLsps",
			Handler:    _PceService_UpdateLsps_Handler,
		},
		{
			MethodName: "CreateLsps",
			Handler:    _PceService_CreateLsps_Handler,
		},
		{
			MethodName: "DeleteLsps",
			Handler:    _PceService_DeleteLsps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("pcep.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 950 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0xe7, 0x3f, 0x27, 0xd9, 0xac, 0x3b, 0x2c, 0x2b, 0x2b, 0x2a, 0x28, 0xb2, 0xaa, 0x55,
	0x5a, 0xa1, 0x52, 0x05, 0x10, 0x14, 0xc4, 0x45, 0xba, 0x49, 0x57, 0x91, 0x42, 0xd7, 0x9a, 0xdd,
	0x85, 0x0b, 0x2e, 0xc0, 0x6b, 0x9f, 0x36, 0x16, 0x8e, 0x3d, 0xcc, 0x8c, 0xc3, 0xe6, 0x99, 0xe0,
	0x9a, 0x47, 0xe0, 0x0d, 0x78, 0x1f, 0x34, 0x33, 0x8e, 0xe3, 0x98, 0xad, 0xd4, 0x3b, 0x9f, 0xef,
	0x9c, 0x39, 0x73, 0xe6, 0xfb, 0xbe, 0x19, 0x19, 0x80, 0x05, 0xc8, 0x9e, 0x33, 0x9e, 0xca, 0x94,
	0x34, 0xd4, 0xb7, 0x7b, 0x0f, 0xad, 0xa5, 0x60, 0x94, 0xa7, 0xe4, 0x14, 0x9a, 0xab, 0x94, 0x2d,
	0x3c, 0xc7, 0x1a, 0x59, 0xe3, 0x2e, 0x35, 0x01, 0x79, 0x02, 0xc7, 0x1c, 0x83, 0x94, 0x87, 0x18,
	0x2e, 0xfd, 0x3b, 0x8c, 0x9d, 0xda, 0xc8, 0x1a, 0x37, 0xe9, 0x21, 0x48, 0x26, 0x70, 0x1a, 0xa7,
	0x81, 0x1f, 0x7b, 0x3c, 0x95, 0x18, 0xc8, 0x28, 0x4d, 0x16, 0xc9, 0xad, 0x40, 0xa7, 0x3e, 0xb2,
	0xc6, 0x1d, 0xfa, 0x60, 0xce, 0xfd, 0x52, 0xef, 0x3c, 0x37, 0x3b, 0xc7, 0x69, 0x2a, 0x50, 0xef,
	0xdc, 0xa1, 0x26, 0xd8, 0xcf, 0x53, 0x2b, 0xcd, 0xe3, 0xfe, 0x5b, 0x83, 0xe3, 0xa5, 0x60, 0x53,
	0x29, 0x79, 0x74, 0x97, 0x49, 0x14, 0x6a, 0x42, 0x81, 0x32, 0x63, 0x1e, 0x8f, 0x52, 0x1e, 0xc9,
	0xad, 0xee, 0xd2, 0xa4, 0x87, 0x20, 0x71, 0xa1, 0xbf, 0x4a, 0xe3, 0xb0, 0x28, 0x32, 0xc7, 0x38,
	0xc0, 0xc8, 0x67, 0xf0, 0x88, 0xa3, 0x40, 0xbe, 0xc1, 0xf0, 0x95, 0x9f, 0x84, 0x7f, 0x44, 0xa1,
	0x5c, 0xe9, 0x23, 0xd4, 0xe9, 0xff, 0x13, 0x64, 0x0c, 0x27, 0x6b, 0xff, 0x7e, 0xba, 0x79, 0xb7,
	0xaf, 0x6d, 0xe8, 0xda, 0x2a, 0x4c, 0x3e, 0x05, 0x60, 0xbe, 0x5c, 0xfd, 0x80, 0x92, 0x47, 0x81,
	0xd3, 0xd4, 0x3b, 0x97, 0x10, 0xf2, 0x02, 0xfa, 0x1c, 0x59, 0xca, 0x25, 0x86, 0x9e, 0x2f, 0x57,
	0x4e, 0x6b, 0x54, 0x1f, 0xf7, 0x26, 0xfd, 0xe7, 0x5a, 0x2c, 0xa3, 0x0e, 0x3d, 0xa8, 0x50, 0x2b,
	0xf0, 0x9e, 0xc5, 0x51, 0x10, 0x49, 0xbd, 0xa2, 0x5d, 0x59, 0x31, 0x57, 0x2b, 0xca, 0x15, 0x6a,
	0x06, 0xbc, 0x0f, 0xe2, 0x2c, 0xc4, 0x69, 0xb2, 0x75, 0x3a, 0xa3, 0xba, 0x9a, 0x61, 0x8f, 0xb8,
	0x7f, 0x5b, 0xd0, 0xf4, 0xee, 0x96, 0x82, 0x69, 0x35, 0x04, 0x5b, 0xcc, 0x72, 0x1e, 0x4d, 0x40,
	0x08, 0x34, 0x12, 0x7f, 0x8d, 0xb9, 0x18, 0xfa, 0x9b, 0x0c, 0xa1, 0x23, 0xd2, 0x8c, 0x07, 0xb8,
	0xf0, 0x34, 0x4d, 0x5d, 0x5a, 0xc4, 0x4a, 0x95, 0x10, 0x85, 0x8c, 0x12, 0x5f, 0x2b, 0xee, 0x69,
	0x6e, 0xba, 0xf4, 0x10, 0x24, 0x67, 0xd0, 0xf2, 0x03, 0x19, 0x6d, 0x50, 0xb3, 0xd2, 0xa1, 0x79,
	0x44, 0x9e, 0x42, 0xd3, 0x97, 0x92, 0x0b, 0xa7, 0x35, 0xb2, 0xc6, 0xbd, 0xc9, 0x47, 0xc5, 0xc1,
	0xf6, 0xba, 0x53, 0x53, 0xe1, 0xae, 0xa1, 0xa3, 0x0e, 0xbc, 0xc1, 0x44, 0x92, 0x73, 0x68, 0xc8,
	0x2d, 0x33, 0x3e, 0x1a, 0x4c, 0xc8, 0x9e, 0x0e, 0x95, 0xbd, 0xd9, 0x32, 0xa4, 0x3a, 0xaf, 0x06,
	0xe7, 0x69, 0x26, 0x91, 0x17, 0xee, 0x2a, 0x62, 0xf2, 0x09, 0xd4, 0x63, 0xc1, 0xf4, 0x79, 0x7a,
	0x93, 0x9e, 0x69, 0xa1, 0x89, 0xa1, 0x0a, 0x77, 0x5f, 0xc3, 0x60, 0x29, 0xd8, 0x2c, 0x5b, 0x33,
	0x8a, 0xbf, 0x67, 0x28, 0xe4, 0x41, 0x33, 0xab, 0xd2, 0xcc, 0x81, 0xf6, 0xdb, 0x28, 0x96, 0xc8,
	0x45, 0x6e, 0xb8, 0x5d, 0xe8, 0x7e, 0x05, 0x27, 0x45, 0x1f, 0xc1, 0xd2, 0x44, 0x20, 0x71, 0xa1,
	0x11, 0x0b, 0x26, 0x1c, 0x4b, 0x8b, 0x39, 0x38, 0x9c, 0x9e, 0xea, 0x9c, 0xfb, 0xa7, 0x05, 0x5d,
	0x2f, 0x08, 0xa8, 0xde, 0xa0, 0x10, 0xc5, 0x2a, 0x89, 0xf2, 0x04, 0x9a, 0x42, 0xfa, 0xd2, 0x28,
	0x35, 0xd8, 0xb5, 0xf1, 0x82, 0xe0, 0x5a, 0xa1, 0xd4, 0x24, 0x95, 0x1d, 0x22, 0xa1, 0x91, 0xb7,
	0x59, 0x9c, 0x5f, 0xd3, 0x12, 0x42, 0xce, 0x61, 0x20, 0x32, 0xa6, 0x1c, 0x27, 0x2e, 0x38, 0xaa,
	0x76, 0x0d, 0x5d, 0x53, 0x41, 0xc9, 0x63, 0xe8, 0x0a, 0x14, 0x42, 0xa9, 0x39, 0xcb, 0x9d, 0xbd,
	0x07, 0xdc, 0x7f, 0x2c, 0xb0, 0x6f, 0x59, 0xe8, 0x4b, 0x54, 0xfc, 0xe5, 0x7c, 0x9d, 0x41, 0xcb,
	0xf0, 0x93, 0x8f, 0x9d, 0x47, 0xe4, 0x35, 0x40, 0x2c, 0xd8, 0xc5, 0xca, 0x4f, 0xde, 0xa1, 0xa2,
	0x4b, 0x91, 0x70, 0x6e, 0xa6, 0xaf, 0xf6, 0x50, 0xac, 0xe4, 0x85, 0xf3, 0x44, 0xf2, 0x2d, 0x2d,
	0xad, 0x1c, 0x52, 0xcd, 0x6c, 0x39, 0x4d, 0x6c, 0xa8, 0xff, 0x86, 0xbb, 0x87, 0x41, 0x7d, 0x2a,
	0x83, 0x6d, 0xfc, 0x38, 0x33, 0x2c, 0xbd, 0xcf, 0x60, 0xba, 0xe2, 0xdb, 0xda, 0x37, 0x96, 0xfb,
	0x3d, 0xf4, 0xf4, 0xee, 0xb9, 0x52, 0x0e, 0xb4, 0x45, 0x16, 0x04, 0x28, 0x44, 0xfe, 0x64, 0xed,
	0x42, 0x75, 0x79, 0x90, 0xf3, 0x94, 0xef, 0x1e, 0x2d, 0x1d, 0xb8, 0x3f, 0x83, 0x6d, 0xf8, 0xfa,
	0x00, 0x1a, 0x3e, 0x87, 0xbe, 0x52, 0xfa, 0x26, 0xcd, 0x79, 0x37, 0x44, 0x1c, 0x18, 0xf1, 0xa0,
	0xc0, 0x7d, 0x05, 0xf6, 0x0c, 0x63, 0xfc, 0xa0, 0xe6, 0x67, 0xd0, 0xd2, 0xd7, 0x59, 0x38, 0x75,
	0xfd, 0x02, 0xe4, 0x91, 0xfb, 0x35, 0x40, 0xe1, 0x2a, 0x41, 0x9e, 0x42, 0xdb, 0xd4, 0xef, 0xbc,
	0x78, 0x52, 0x98, 0xc8, 0x94, 0xd0, 0x5d, 0xde, 0x6d, 0x43, 0x73, 0xbe, 0x66, 0x72, 0xfb, 0xec,
	0x57, 0xe8, 0x97, 0x2f, 0x1a, 0x01, 0x68, 0x2d, 0xaf, 0xbd, 0x5f, 0x6e, 0x3d, 0xfb, 0x88, 0xf4,
	0xa1, 0xa3, 0xbe, 0x67, 0x57, 0x3f, 0xbd, 0xb1, 0x2d, 0x32, 0x00, 0x30, 0x99, 0xd9, 0xf4, 0x66,
	0x6e, 0xd7, 0x08, 0x81, 0x81, 0xce, 0xce, 0x97, 0xf3, 0xcb, 0xe9, 0xcd, 0xe2, 0xea, 0x8d, 0x5d,
	0xdf, 0x61, 0x74, 0xfe, 0xe3, 0xd5, 0x85, 0xc1, 0x1a, 0xcf, 0x1e, 0x43, 0x67, 0xe7, 0x62, 0xd2,
	0x81, 0xc6, 0x62, 0xb6, 0x9c, 0xdb, 0x47, 0xa4, 0x05, 0xb5, 0x5b, 0xcf, 0xb6, 0x26, 0x7f, 0xd5,
	0xd4, 0x11, 0xf0, 0x1a, 0xf9, 0x26, 0x0a, 0x90, 0x7c, 0x07, 0x70, 0x89, 0x32, 0xbf, 0x61, 0xe4,
	0xb4, 0x90, 0xb7, 0x74, 0x71, 0x87, 0x1f, 0x57, 0x50, 0x23, 0xae, 0x7b, 0x44, 0x5e, 0xc0, 0xf1,
	0x25, 0xca, 0x12, 0x21, 0x39, 0xfb, 0xfa, 0xa4, 0x43, 0xbb, 0x42, 0x86, 0x70, 0x8f, 0xc8, 0x4b,
	0x80, 0xc2, 0xa3, 0x82, 0x9c, 0x3d, 0xec, 0xda, 0xe1, 0xa3, 0xfd, 0x8b, 0xbe, 0xdf, 0xec, 0x25,
	0x40, 0xe1, 0x8d, 0x62, 0x69, 0xd5, 0x2d, 0xef, 0x5d, 0x5a, 0x28, 0x5f, 0x2c, 0xad, 0x7a, 0xe1,
	0xc1, 0xa5, 0x77, 0x2d, 0xfd, 0x0f, 0xf0, 0xc5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x89,
	0x48, 0xb0, 0x11, 0x08, 0x00, 0x00,
}
