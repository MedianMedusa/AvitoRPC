// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package phones

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PhonesClient is the client API for Phones service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhonesClient interface {
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type phonesClient struct {
	cc grpc.ClientConnInterface
}

func NewPhonesClient(cc grpc.ClientConnInterface) PhonesClient {
	return &phonesClient{cc}
}

func (c *phonesClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/phones.Phones/sendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhonesServer is the server API for Phones service.
// All implementations must embed UnimplementedPhonesServer
// for forward compatibility
type PhonesServer interface {
	SendRequest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPhonesServer()
}

// UnimplementedPhonesServer must be embedded to have forward compatible implementations.
type UnimplementedPhonesServer struct {
}

func (UnimplementedPhonesServer) SendRequest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedPhonesServer) mustEmbedUnimplementedPhonesServer() {}

// UnsafePhonesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhonesServer will
// result in compilation errors.
type UnsafePhonesServer interface {
	mustEmbedUnimplementedPhonesServer()
}

func RegisterPhonesServer(s grpc.ServiceRegistrar, srv PhonesServer) {
	s.RegisterService(&Phones_ServiceDesc, srv)
}

func _Phones_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhonesServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phones.Phones/sendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhonesServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Phones_ServiceDesc is the grpc.ServiceDesc for Phones service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Phones_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phones.Phones",
	HandlerType: (*PhonesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendRequest",
			Handler:    _Phones_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phones.proto",
}