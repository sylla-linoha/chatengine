//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Teamgram Authors.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: session.tl.proto

package session

import (
	context "context"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCSession_SessionQueryAuthKey_FullMethodName           = "/session.RPCSession/session_queryAuthKey"
	RPCSession_SessionSetAuthKey_FullMethodName             = "/session.RPCSession/session_setAuthKey"
	RPCSession_SessionCreateSession_FullMethodName          = "/session.RPCSession/session_createSession"
	RPCSession_SessionSendDataToSession_FullMethodName      = "/session.RPCSession/session_sendDataToSession"
	RPCSession_SessionSendHttpDataToSession_FullMethodName  = "/session.RPCSession/session_sendHttpDataToSession"
	RPCSession_SessionCloseSession_FullMethodName           = "/session.RPCSession/session_closeSession"
	RPCSession_SessionPushUpdatesData_FullMethodName        = "/session.RPCSession/session_pushUpdatesData"
	RPCSession_SessionPushSessionUpdatesData_FullMethodName = "/session.RPCSession/session_pushSessionUpdatesData"
	RPCSession_SessionPushRpcResultData_FullMethodName      = "/session.RPCSession/session_pushRpcResultData"
)

// RPCSessionClient is the client API for RPCSession service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCSessionClient interface {
	SessionQueryAuthKey(ctx context.Context, in *TLSessionQueryAuthKey, opts ...grpc.CallOption) (*mtproto.AuthKeyInfo, error)
	SessionSetAuthKey(ctx context.Context, in *TLSessionSetAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionCreateSession(ctx context.Context, in *TLSessionCreateSession, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionSendDataToSession(ctx context.Context, in *TLSessionSendDataToSession, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionSendHttpDataToSession(ctx context.Context, in *TLSessionSendHttpDataToSession, opts ...grpc.CallOption) (*HttpSessionData, error)
	SessionCloseSession(ctx context.Context, in *TLSessionCloseSession, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionPushUpdatesData(ctx context.Context, in *TLSessionPushUpdatesData, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionPushSessionUpdatesData(ctx context.Context, in *TLSessionPushSessionUpdatesData, opts ...grpc.CallOption) (*mtproto.Bool, error)
	SessionPushRpcResultData(ctx context.Context, in *TLSessionPushRpcResultData, opts ...grpc.CallOption) (*mtproto.Bool, error)
}

type rPCSessionClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCSessionClient(cc grpc.ClientConnInterface) RPCSessionClient {
	return &rPCSessionClient{cc}
}

func (c *rPCSessionClient) SessionQueryAuthKey(ctx context.Context, in *TLSessionQueryAuthKey, opts ...grpc.CallOption) (*mtproto.AuthKeyInfo, error) {
	out := new(mtproto.AuthKeyInfo)
	err := c.cc.Invoke(ctx, RPCSession_SessionQueryAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionSetAuthKey(ctx context.Context, in *TLSessionSetAuthKey, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionSetAuthKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionCreateSession(ctx context.Context, in *TLSessionCreateSession, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionCreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionSendDataToSession(ctx context.Context, in *TLSessionSendDataToSession, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionSendDataToSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionSendHttpDataToSession(ctx context.Context, in *TLSessionSendHttpDataToSession, opts ...grpc.CallOption) (*HttpSessionData, error) {
	out := new(HttpSessionData)
	err := c.cc.Invoke(ctx, RPCSession_SessionSendHttpDataToSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionCloseSession(ctx context.Context, in *TLSessionCloseSession, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionCloseSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionPushUpdatesData(ctx context.Context, in *TLSessionPushUpdatesData, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionPushUpdatesData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionPushSessionUpdatesData(ctx context.Context, in *TLSessionPushSessionUpdatesData, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionPushSessionUpdatesData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSessionClient) SessionPushRpcResultData(ctx context.Context, in *TLSessionPushRpcResultData, opts ...grpc.CallOption) (*mtproto.Bool, error) {
	out := new(mtproto.Bool)
	err := c.cc.Invoke(ctx, RPCSession_SessionPushRpcResultData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCSessionServer is the server API for RPCSession service.
// All implementations should embed UnimplementedRPCSessionServer
// for forward compatibility
type RPCSessionServer interface {
	SessionQueryAuthKey(context.Context, *TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error)
	SessionSetAuthKey(context.Context, *TLSessionSetAuthKey) (*mtproto.Bool, error)
	SessionCreateSession(context.Context, *TLSessionCreateSession) (*mtproto.Bool, error)
	SessionSendDataToSession(context.Context, *TLSessionSendDataToSession) (*mtproto.Bool, error)
	SessionSendHttpDataToSession(context.Context, *TLSessionSendHttpDataToSession) (*HttpSessionData, error)
	SessionCloseSession(context.Context, *TLSessionCloseSession) (*mtproto.Bool, error)
	SessionPushUpdatesData(context.Context, *TLSessionPushUpdatesData) (*mtproto.Bool, error)
	SessionPushSessionUpdatesData(context.Context, *TLSessionPushSessionUpdatesData) (*mtproto.Bool, error)
	SessionPushRpcResultData(context.Context, *TLSessionPushRpcResultData) (*mtproto.Bool, error)
}

// UnimplementedRPCSessionServer should be embedded to have forward compatible implementations.
type UnimplementedRPCSessionServer struct {
}

func (UnimplementedRPCSessionServer) SessionQueryAuthKey(context.Context, *TLSessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionQueryAuthKey not implemented")
}
func (UnimplementedRPCSessionServer) SessionSetAuthKey(context.Context, *TLSessionSetAuthKey) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionSetAuthKey not implemented")
}
func (UnimplementedRPCSessionServer) SessionCreateSession(context.Context, *TLSessionCreateSession) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionCreateSession not implemented")
}
func (UnimplementedRPCSessionServer) SessionSendDataToSession(context.Context, *TLSessionSendDataToSession) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionSendDataToSession not implemented")
}
func (UnimplementedRPCSessionServer) SessionSendHttpDataToSession(context.Context, *TLSessionSendHttpDataToSession) (*HttpSessionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionSendHttpDataToSession not implemented")
}
func (UnimplementedRPCSessionServer) SessionCloseSession(context.Context, *TLSessionCloseSession) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionCloseSession not implemented")
}
func (UnimplementedRPCSessionServer) SessionPushUpdatesData(context.Context, *TLSessionPushUpdatesData) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionPushUpdatesData not implemented")
}
func (UnimplementedRPCSessionServer) SessionPushSessionUpdatesData(context.Context, *TLSessionPushSessionUpdatesData) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionPushSessionUpdatesData not implemented")
}
func (UnimplementedRPCSessionServer) SessionPushRpcResultData(context.Context, *TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionPushRpcResultData not implemented")
}

// UnsafeRPCSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCSessionServer will
// result in compilation errors.
type UnsafeRPCSessionServer interface {
	mustEmbedUnimplementedRPCSessionServer()
}

func RegisterRPCSessionServer(s grpc.ServiceRegistrar, srv RPCSessionServer) {
	s.RegisterService(&RPCSession_ServiceDesc, srv)
}

func _RPCSession_SessionQueryAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionQueryAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionQueryAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionQueryAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionQueryAuthKey(ctx, req.(*TLSessionQueryAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionSetAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionSetAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionSetAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionSetAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionSetAuthKey(ctx, req.(*TLSessionSetAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionCreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionCreateSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionCreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionCreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionCreateSession(ctx, req.(*TLSessionCreateSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionSendDataToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionSendDataToSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionSendDataToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionSendDataToSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionSendDataToSession(ctx, req.(*TLSessionSendDataToSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionSendHttpDataToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionSendHttpDataToSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionSendHttpDataToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionSendHttpDataToSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionSendHttpDataToSession(ctx, req.(*TLSessionSendHttpDataToSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionCloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionCloseSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionCloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionCloseSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionCloseSession(ctx, req.(*TLSessionCloseSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionPushUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionPushUpdatesData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionPushUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionPushUpdatesData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionPushUpdatesData(ctx, req.(*TLSessionPushUpdatesData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionPushSessionUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionPushSessionUpdatesData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionPushSessionUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionPushSessionUpdatesData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionPushSessionUpdatesData(ctx, req.(*TLSessionPushSessionUpdatesData))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSession_SessionPushRpcResultData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSessionPushRpcResultData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSessionServer).SessionPushRpcResultData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCSession_SessionPushRpcResultData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSessionServer).SessionPushRpcResultData(ctx, req.(*TLSessionPushRpcResultData))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCSession_ServiceDesc is the grpc.ServiceDesc for RPCSession service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCSession_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.RPCSession",
	HandlerType: (*RPCSessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "session_queryAuthKey",
			Handler:    _RPCSession_SessionQueryAuthKey_Handler,
		},
		{
			MethodName: "session_setAuthKey",
			Handler:    _RPCSession_SessionSetAuthKey_Handler,
		},
		{
			MethodName: "session_createSession",
			Handler:    _RPCSession_SessionCreateSession_Handler,
		},
		{
			MethodName: "session_sendDataToSession",
			Handler:    _RPCSession_SessionSendDataToSession_Handler,
		},
		{
			MethodName: "session_sendHttpDataToSession",
			Handler:    _RPCSession_SessionSendHttpDataToSession_Handler,
		},
		{
			MethodName: "session_closeSession",
			Handler:    _RPCSession_SessionCloseSession_Handler,
		},
		{
			MethodName: "session_pushUpdatesData",
			Handler:    _RPCSession_SessionPushUpdatesData_Handler,
		},
		{
			MethodName: "session_pushSessionUpdatesData",
			Handler:    _RPCSession_SessionPushSessionUpdatesData_Handler,
		},
		{
			MethodName: "session_pushRpcResultData",
			Handler:    _RPCSession_SessionPushRpcResultData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.tl.proto",
}