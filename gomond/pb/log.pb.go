// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Level int32

const (
	Level_info    Level = 0
	Level_debug   Level = 1
	Level_error   Level = 2
	Level_fatal   Level = 3
	Level_warning Level = 4
	Level_unknown Level = 5
)

var Level_name = map[int32]string{
	0: "info",
	1: "debug",
	2: "error",
	3: "fatal",
	4: "warning",
	5: "unknown",
}

var Level_value = map[string]int32{
	"info":    0,
	"debug":   1,
	"error":   2,
	"fatal":   3,
	"warning": 4,
	"unknown": 5,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

type Type int32

const (
	Type_file Type = 0
	Type_http Type = 1
)

var Type_name = map[int32]string{
	0: "file",
	1: "http",
}

var Type_value = map[string]int32{
	"file": 0,
	"http": 1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}

type Log struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Level                Level    `protobuf:"varint,2,opt,name=level,proto3,enum=pb.Level" json:"level,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	File                 string   `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	Line                 int32    `protobuf:"varint,7,opt,name=line,proto3" json:"line,omitempty"`
	Type                 Type     `protobuf:"varint,8,opt,name=type,proto3,enum=pb.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Log) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_info
}

func (m *Log) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Log) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Log) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Log) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_file
}

type AppLiveRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppLiveRequest) Reset()         { *m = AppLiveRequest{} }
func (m *AppLiveRequest) String() string { return proto.CompactTextString(m) }
func (*AppLiveRequest) ProtoMessage()    {}
func (*AppLiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{1}
}

func (m *AppLiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppLiveRequest.Unmarshal(m, b)
}
func (m *AppLiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppLiveRequest.Marshal(b, m, deterministic)
}
func (m *AppLiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppLiveRequest.Merge(m, src)
}
func (m *AppLiveRequest) XXX_Size() int {
	return xxx_messageInfo_AppLiveRequest.Size(m)
}
func (m *AppLiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppLiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppLiveRequest proto.InternalMessageInfo

func (m *AppLiveRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

type AppRangeRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	From                 int64    `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppRangeRequest) Reset()         { *m = AppRangeRequest{} }
func (m *AppRangeRequest) String() string { return proto.CompactTextString(m) }
func (*AppRangeRequest) ProtoMessage()    {}
func (*AppRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{2}
}

func (m *AppRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppRangeRequest.Unmarshal(m, b)
}
func (m *AppRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppRangeRequest.Marshal(b, m, deterministic)
}
func (m *AppRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppRangeRequest.Merge(m, src)
}
func (m *AppRangeRequest) XXX_Size() int {
	return xxx_messageInfo_AppRangeRequest.Size(m)
}
func (m *AppRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppRangeRequest proto.InternalMessageInfo

func (m *AppRangeRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *AppRangeRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *AppRangeRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type AppLatestRequest struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppLatestRequest) Reset()         { *m = AppLatestRequest{} }
func (m *AppLatestRequest) String() string { return proto.CompactTextString(m) }
func (*AppLatestRequest) ProtoMessage()    {}
func (*AppLatestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{3}
}

func (m *AppLatestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppLatestRequest.Unmarshal(m, b)
}
func (m *AppLatestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppLatestRequest.Marshal(b, m, deterministic)
}
func (m *AppLatestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppLatestRequest.Merge(m, src)
}
func (m *AppLatestRequest) XXX_Size() int {
	return xxx_messageInfo_AppLatestRequest.Size(m)
}
func (m *AppLatestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppLatestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppLatestRequest proto.InternalMessageInfo

func (m *AppLatestRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *AppLatestRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type AppResponse struct {
	Logs                 []*Log   `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	App                  string   `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppResponse) Reset()         { *m = AppResponse{} }
func (m *AppResponse) String() string { return proto.CompactTextString(m) }
func (*AppResponse) ProtoMessage()    {}
func (*AppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{4}
}

func (m *AppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppResponse.Unmarshal(m, b)
}
func (m *AppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppResponse.Marshal(b, m, deterministic)
}
func (m *AppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppResponse.Merge(m, src)
}
func (m *AppResponse) XXX_Size() int {
	return xxx_messageInfo_AppResponse.Size(m)
}
func (m *AppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppResponse proto.InternalMessageInfo

func (m *AppResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *AppResponse) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *AppResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.Level", Level_name, Level_value)
	proto.RegisterEnum("pb.Type", Type_name, Type_value)
	proto.RegisterType((*Log)(nil), "pb.Log")
	proto.RegisterType((*AppLiveRequest)(nil), "pb.AppLiveRequest")
	proto.RegisterType((*AppRangeRequest)(nil), "pb.AppRangeRequest")
	proto.RegisterType((*AppLatestRequest)(nil), "pb.AppLatestRequest")
	proto.RegisterType((*AppResponse)(nil), "pb.AppResponse")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0x1b, 0x31,
	0x10, 0xf5, 0x7e, 0xf9, 0x63, 0x5c, 0x1c, 0x31, 0xe4, 0x20, 0xdc, 0x42, 0xcb, 0x5e, 0x1a, 0x72,
	0x30, 0xad, 0x7b, 0xeb, 0x2d, 0xa7, 0x52, 0xf0, 0x21, 0x88, 0xfe, 0x01, 0x19, 0xcb, 0x6b, 0xd1,
	0x5d, 0x49, 0xdd, 0x95, 0x13, 0xf2, 0x07, 0xfb, 0xbb, 0x3a, 0xa3, 0x64, 0xa1, 0x21, 0xf8, 0xf6,
	0xe6, 0xf1, 0xf4, 0xf4, 0xe6, 0x49, 0xb0, 0x68, 0x7d, 0xb3, 0x09, 0xbd, 0x8f, 0x1e, 0xf3, 0xb0,
	0xaf, 0xff, 0x66, 0x50, 0xec, 0x7c, 0x83, 0x02, 0x0a, 0x1d, 0x82, 0xcc, 0x3e, 0x65, 0x37, 0x0b,
	0xc5, 0x10, 0x3f, 0x42, 0xd5, 0x9a, 0x07, 0xd3, 0xca, 0x9c, 0xb8, 0xd5, 0x76, 0xb1, 0x09, 0xfb,
	0xcd, 0x8e, 0x09, 0xf5, 0xcc, 0xa3, 0x84, 0x59, 0xd0, 0x4f, 0xad, 0xd7, 0x07, 0x59, 0x90, 0xe4,
	0x9d, 0x1a, 0x47, 0x44, 0x28, 0xa3, 0xed, 0x8c, 0x2c, 0x89, 0x2e, 0x54, 0xc2, 0xac, 0xee, 0xcc,
	0x30, 0xe8, 0xc6, 0xc8, 0x2a, 0x5d, 0x32, 0x8e, 0xac, 0x3e, 0xda, 0xd6, 0xc8, 0x69, 0xa2, 0x13,
	0x66, 0xae, 0xb5, 0xce, 0xc8, 0x19, 0x71, 0x95, 0x4a, 0x18, 0x3f, 0x90, 0xeb, 0x53, 0x30, 0x72,
	0x9e, 0xf2, 0xcc, 0x39, 0xcf, 0x2f, 0x9a, 0x55, 0x62, 0xeb, 0x1a, 0x56, 0x77, 0x21, 0xec, 0xec,
	0x83, 0x51, 0xe6, 0xcf, 0xd9, 0x0c, 0xf1, 0xed, 0x4a, 0xf5, 0x0f, 0xb8, 0x22, 0x8d, 0xd2, 0xae,
	0xb9, 0x2c, 0x4a, 0x71, 0x7a, 0xdf, 0xa5, 0xb5, 0x29, 0x3c, 0x63, 0x5c, 0x41, 0x1e, 0x7d, 0xda,
	0xb2, 0x50, 0x84, 0xea, 0xef, 0x20, 0xf8, 0x32, 0x1d, 0xc9, 0xe2, 0xb2, 0xd3, 0x35, 0x35, 0x68,
	0x3b, 0x1b, 0x5f, 0xac, 0x9e, 0x87, 0xfa, 0x1e, 0x96, 0x1c, 0xc2, 0x0c, 0xc1, 0xbb, 0xc1, 0xe0,
	0x7b, 0xda, 0xd4, 0x37, 0x03, 0x9d, 0x2b, 0x6e, 0x96, 0xdb, 0x59, 0x6a, 0xd9, 0x37, 0x2a, 0x91,
	0xa3, 0x67, 0xfe, 0x2a, 0xdd, 0x49, 0x0f, 0xa7, 0x94, 0x85, 0xca, 0x62, 0x7c, 0xbb, 0x83, 0x2a,
	0x3d, 0x0c, 0xce, 0xa1, 0xb4, 0xee, 0xe8, 0xc5, 0x04, 0x17, 0x50, 0x1d, 0xcc, 0xfe, 0xdc, 0x88,
	0x8c, 0xa1, 0xe9, 0x7b, 0xdf, 0x8b, 0x9c, 0xe1, 0x51, 0x47, 0xdd, 0x8a, 0x02, 0x97, 0x30, 0x7b,
	0xd4, 0xbd, 0xb3, 0xae, 0x11, 0x25, 0x0f, 0x67, 0xf7, 0xdb, 0xf9, 0x47, 0x27, 0xaa, 0xdb, 0x35,
	0x94, 0x5c, 0x2b, 0x9b, 0xf1, 0x53, 0x90, 0x19, 0xa1, 0x53, 0x8c, 0x41, 0x64, 0xdb, 0x03, 0x4c,
	0x29, 0xdc, 0xdd, 0xfd, 0x4f, 0xfc, 0x0c, 0x25, 0x77, 0x8d, 0xc8, 0x81, 0x5f, 0x17, 0xbf, 0x1e,
	0x97, 0xa8, 0x27, 0x5f, 0x32, 0xfc, 0x4a, 0x47, 0x52, 0x4f, 0x78, 0x3d, 0x4a, 0xff, 0xaf, 0x6d,
	0x7d, 0xf5, 0xc2, 0x8e, 0x85, 0xd4, 0x93, 0xfd, 0x34, 0x7d, 0xcf, 0x6f, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x35, 0x16, 0xa7, 0xa1, 0xab, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogAPIClient is the client API for LogAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogAPIClient interface {
	Live(ctx context.Context, in *AppLiveRequest, opts ...grpc.CallOption) (LogAPI_LiveClient, error)
	Latest(ctx context.Context, in *AppLatestRequest, opts ...grpc.CallOption) (*AppResponse, error)
}

type logAPIClient struct {
	cc *grpc.ClientConn
}

func NewLogAPIClient(cc *grpc.ClientConn) LogAPIClient {
	return &logAPIClient{cc}
}

func (c *logAPIClient) Live(ctx context.Context, in *AppLiveRequest, opts ...grpc.CallOption) (LogAPI_LiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogAPI_serviceDesc.Streams[0], "/pb.LogAPI/Live", opts...)
	if err != nil {
		return nil, err
	}
	x := &logAPILiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogAPI_LiveClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logAPILiveClient struct {
	grpc.ClientStream
}

func (x *logAPILiveClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logAPIClient) Latest(ctx context.Context, in *AppLatestRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/pb.LogAPI/Latest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogAPIServer is the server API for LogAPI service.
type LogAPIServer interface {
	Live(*AppLiveRequest, LogAPI_LiveServer) error
	Latest(context.Context, *AppLatestRequest) (*AppResponse, error)
}

// UnimplementedLogAPIServer can be embedded to have forward compatible implementations.
type UnimplementedLogAPIServer struct {
}

func (*UnimplementedLogAPIServer) Live(req *AppLiveRequest, srv LogAPI_LiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Live not implemented")
}
func (*UnimplementedLogAPIServer) Latest(ctx context.Context, req *AppLatestRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Latest not implemented")
}

func RegisterLogAPIServer(s *grpc.Server, srv LogAPIServer) {
	s.RegisterService(&_LogAPI_serviceDesc, srv)
}

func _LogAPI_Live_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AppLiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogAPIServer).Live(m, &logAPILiveServer{stream})
}

type LogAPI_LiveServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logAPILiveServer struct {
	grpc.ServerStream
}

func (x *logAPILiveServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _LogAPI_Latest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppLatestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).Latest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogAPI/Latest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).Latest(ctx, req.(*AppLatestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogAPI",
	HandlerType: (*LogAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Latest",
			Handler:    _LogAPI_Latest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Live",
			Handler:       _LogAPI_Live_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "log.proto",
}
