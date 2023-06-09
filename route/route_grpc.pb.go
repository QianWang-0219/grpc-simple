// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: route.proto

package route

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LocalGuideClient is the client API for LocalGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalGuideClient interface {
	GetLocation(ctx context.Context, opts ...grpc.CallOption) (LocalGuide_GetLocationClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (LocalGuide_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *MetaData, opts ...grpc.CallOption) (LocalGuide_DownloadFileClient, error)
}

type localGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalGuideClient(cc grpc.ClientConnInterface) LocalGuideClient {
	return &localGuideClient{cc}
}

func (c *localGuideClient) GetLocation(ctx context.Context, opts ...grpc.CallOption) (LocalGuide_GetLocationClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocalGuide_ServiceDesc.Streams[0], "/route.localGuide/GetLocation", opts...)
	if err != nil {
		return nil, err
	}
	x := &localGuideGetLocationClient{stream}
	return x, nil
}

type LocalGuide_GetLocationClient interface {
	Send(*IniLoc) error
	Recv() (*FinLoc, error)
	grpc.ClientStream
}

type localGuideGetLocationClient struct {
	grpc.ClientStream
}

func (x *localGuideGetLocationClient) Send(m *IniLoc) error {
	return x.ClientStream.SendMsg(m)
}

func (x *localGuideGetLocationClient) Recv() (*FinLoc, error) {
	m := new(FinLoc)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *localGuideClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (LocalGuide_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocalGuide_ServiceDesc.Streams[1], "/route.localGuide/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &localGuideUploadFileClient{stream}
	return x, nil
}

type LocalGuide_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*StringResponse, error)
	grpc.ClientStream
}

type localGuideUploadFileClient struct {
	grpc.ClientStream
}

func (x *localGuideUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *localGuideUploadFileClient) CloseAndRecv() (*StringResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *localGuideClient) DownloadFile(ctx context.Context, in *MetaData, opts ...grpc.CallOption) (LocalGuide_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocalGuide_ServiceDesc.Streams[2], "/route.localGuide/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &localGuideDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LocalGuide_DownloadFileClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type localGuideDownloadFileClient struct {
	grpc.ClientStream
}

func (x *localGuideDownloadFileClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocalGuideServer is the server API for LocalGuide service.
// All implementations must embed UnimplementedLocalGuideServer
// for forward compatibility
type LocalGuideServer interface {
	GetLocation(LocalGuide_GetLocationServer) error
	UploadFile(LocalGuide_UploadFileServer) error
	DownloadFile(*MetaData, LocalGuide_DownloadFileServer) error
	mustEmbedUnimplementedLocalGuideServer()
}

// UnimplementedLocalGuideServer must be embedded to have forward compatible implementations.
type UnimplementedLocalGuideServer struct {
}

func (UnimplementedLocalGuideServer) GetLocation(LocalGuide_GetLocationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocalGuideServer) UploadFile(LocalGuide_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedLocalGuideServer) DownloadFile(*MetaData, LocalGuide_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedLocalGuideServer) mustEmbedUnimplementedLocalGuideServer() {}

// UnsafeLocalGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalGuideServer will
// result in compilation errors.
type UnsafeLocalGuideServer interface {
	mustEmbedUnimplementedLocalGuideServer()
}

func RegisterLocalGuideServer(s grpc.ServiceRegistrar, srv LocalGuideServer) {
	s.RegisterService(&LocalGuide_ServiceDesc, srv)
}

func _LocalGuide_GetLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocalGuideServer).GetLocation(&localGuideGetLocationServer{stream})
}

type LocalGuide_GetLocationServer interface {
	Send(*FinLoc) error
	Recv() (*IniLoc, error)
	grpc.ServerStream
}

type localGuideGetLocationServer struct {
	grpc.ServerStream
}

func (x *localGuideGetLocationServer) Send(m *FinLoc) error {
	return x.ServerStream.SendMsg(m)
}

func (x *localGuideGetLocationServer) Recv() (*IniLoc, error) {
	m := new(IniLoc)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LocalGuide_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocalGuideServer).UploadFile(&localGuideUploadFileServer{stream})
}

type LocalGuide_UploadFileServer interface {
	SendAndClose(*StringResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type localGuideUploadFileServer struct {
	grpc.ServerStream
}

func (x *localGuideUploadFileServer) SendAndClose(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *localGuideUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LocalGuide_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetaData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocalGuideServer).DownloadFile(m, &localGuideDownloadFileServer{stream})
}

type LocalGuide_DownloadFileServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type localGuideDownloadFileServer struct {
	grpc.ServerStream
}

func (x *localGuideDownloadFileServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LocalGuide_ServiceDesc is the grpc.ServiceDesc for LocalGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.localGuide",
	HandlerType: (*LocalGuideServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLocation",
			Handler:       _LocalGuide_GetLocation_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _LocalGuide_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _LocalGuide_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "route.proto",
}
