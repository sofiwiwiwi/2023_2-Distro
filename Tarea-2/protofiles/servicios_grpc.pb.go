// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: protofiles/servicios.proto

package protofiles

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

// OMSClient is the client API for OMS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OMSClient interface {
	SendNombreEstado(ctx context.Context, in *InfoPersonaContinenteReq, opts ...grpc.CallOption) (*Empty, error)
	AskNombres(ctx context.Context, in *InfoPersonasCondicionReq, opts ...grpc.CallOption) (*InfoPersonasCondicionResp, error)
}

type oMSClient struct {
	cc grpc.ClientConnInterface
}

func NewOMSClient(cc grpc.ClientConnInterface) OMSClient {
	return &oMSClient{cc}
}

func (c *oMSClient) SendNombreEstado(ctx context.Context, in *InfoPersonaContinenteReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protofiles.OMS/sendNombreEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oMSClient) AskNombres(ctx context.Context, in *InfoPersonasCondicionReq, opts ...grpc.CallOption) (*InfoPersonasCondicionResp, error) {
	out := new(InfoPersonasCondicionResp)
	err := c.cc.Invoke(ctx, "/protofiles.OMS/askNombres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OMSServer is the server API for OMS service.
// All implementations must embed UnimplementedOMSServer
// for forward compatibility
type OMSServer interface {
	SendNombreEstado(context.Context, *InfoPersonaContinenteReq) (*Empty, error)
	AskNombres(context.Context, *InfoPersonasCondicionReq) (*InfoPersonasCondicionResp, error)
	mustEmbedUnimplementedOMSServer()
}

// UnimplementedOMSServer must be embedded to have forward compatible implementations.
type UnimplementedOMSServer struct {
}

func (UnimplementedOMSServer) SendNombreEstado(context.Context, *InfoPersonaContinenteReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNombreEstado not implemented")
}
func (UnimplementedOMSServer) AskNombres(context.Context, *InfoPersonasCondicionReq) (*InfoPersonasCondicionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskNombres not implemented")
}
func (UnimplementedOMSServer) mustEmbedUnimplementedOMSServer() {}

// UnsafeOMSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OMSServer will
// result in compilation errors.
type UnsafeOMSServer interface {
	mustEmbedUnimplementedOMSServer()
}

func RegisterOMSServer(s grpc.ServiceRegistrar, srv OMSServer) {
	s.RegisterService(&OMS_ServiceDesc, srv)
}

func _OMS_SendNombreEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPersonaContinenteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMSServer).SendNombreEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.OMS/sendNombreEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMSServer).SendNombreEstado(ctx, req.(*InfoPersonaContinenteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OMS_AskNombres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoPersonasCondicionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMSServer).AskNombres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.OMS/askNombres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMSServer).AskNombres(ctx, req.(*InfoPersonasCondicionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OMS_ServiceDesc is the grpc.ServiceDesc for OMS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OMS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.OMS",
	HandlerType: (*OMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendNombreEstado",
			Handler:    _OMS_SendNombreEstado_Handler,
		},
		{
			MethodName: "askNombres",
			Handler:    _OMS_AskNombres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/servicios.proto",
}

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	SendIdEstado(ctx context.Context, in *DatosIdNombreReq, opts ...grpc.CallOption) (*Empty, error)
	AskNombreId(ctx context.Context, in *NombrePersonaReq, opts ...grpc.CallOption) (*NombrePersonaResp, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) SendIdEstado(ctx context.Context, in *DatosIdNombreReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protofiles.dataNode/sendIdEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) AskNombreId(ctx context.Context, in *NombrePersonaReq, opts ...grpc.CallOption) (*NombrePersonaResp, error) {
	out := new(NombrePersonaResp)
	err := c.cc.Invoke(ctx, "/protofiles.dataNode/askNombreId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	SendIdEstado(context.Context, *DatosIdNombreReq) (*Empty, error)
	AskNombreId(context.Context, *NombrePersonaReq) (*NombrePersonaResp, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) SendIdEstado(context.Context, *DatosIdNombreReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIdEstado not implemented")
}
func (UnimplementedDataNodeServer) AskNombreId(context.Context, *NombrePersonaReq) (*NombrePersonaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskNombreId not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&DataNode_ServiceDesc, srv)
}

func _DataNode_SendIdEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatosIdNombreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).SendIdEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.dataNode/sendIdEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).SendIdEstado(ctx, req.(*DatosIdNombreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_AskNombreId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NombrePersonaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).AskNombreId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.dataNode/askNombreId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).AskNombreId(ctx, req.(*NombrePersonaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNode_ServiceDesc is the grpc.ServiceDesc for DataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.dataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendIdEstado",
			Handler:    _DataNode_SendIdEstado_Handler,
		},
		{
			MethodName: "askNombreId",
			Handler:    _DataNode_AskNombreId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/servicios.proto",
}
