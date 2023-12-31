// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: conversation/conversation.proto

package conversation

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

const (
	Conversation_GetConversation_FullMethodName                  = "/OpenIMServer.conversation.conversation/GetConversation"
	Conversation_GetAllConversations_FullMethodName              = "/OpenIMServer.conversation.conversation/GetAllConversations"
	Conversation_GetConversations_FullMethodName                 = "/OpenIMServer.conversation.conversation/GetConversations"
	Conversation_SetConversation_FullMethodName                  = "/OpenIMServer.conversation.conversation/SetConversation"
	Conversation_GetRecvMsgNotNotifyUserIDs_FullMethodName       = "/OpenIMServer.conversation.conversation/GetRecvMsgNotNotifyUserIDs"
	Conversation_CreateSingleChatConversations_FullMethodName    = "/OpenIMServer.conversation.conversation/CreateSingleChatConversations"
	Conversation_CreateGroupChatConversations_FullMethodName     = "/OpenIMServer.conversation.conversation/CreateGroupChatConversations"
	Conversation_SetConversationMaxSeq_FullMethodName            = "/OpenIMServer.conversation.conversation/SetConversationMaxSeq"
	Conversation_GetConversationIDs_FullMethodName               = "/OpenIMServer.conversation.conversation/GetConversationIDs"
	Conversation_SetConversations_FullMethodName                 = "/OpenIMServer.conversation.conversation/SetConversations"
	Conversation_GetUserConversationIDsHash_FullMethodName       = "/OpenIMServer.conversation.conversation/GetUserConversationIDsHash"
	Conversation_GetConversationsByConversationID_FullMethodName = "/OpenIMServer.conversation.conversation/GetConversationsByConversationID"
)

// ConversationClient is the client API for Conversation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationClient interface {
	GetConversation(ctx context.Context, in *GetConversationReq, opts ...grpc.CallOption) (*GetConversationResp, error)
	GetAllConversations(ctx context.Context, in *GetAllConversationsReq, opts ...grpc.CallOption) (*GetAllConversationsResp, error)
	GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
	SetConversation(ctx context.Context, in *SetConversationReq, opts ...grpc.CallOption) (*SetConversationResp, error)
	GetRecvMsgNotNotifyUserIDs(ctx context.Context, in *GetRecvMsgNotNotifyUserIDsReq, opts ...grpc.CallOption) (*GetRecvMsgNotNotifyUserIDsResp, error)
	CreateSingleChatConversations(ctx context.Context, in *CreateSingleChatConversationsReq, opts ...grpc.CallOption) (*CreateSingleChatConversationsResp, error)
	CreateGroupChatConversations(ctx context.Context, in *CreateGroupChatConversationsReq, opts ...grpc.CallOption) (*CreateGroupChatConversationsResp, error)
	SetConversationMaxSeq(ctx context.Context, in *SetConversationMaxSeqReq, opts ...grpc.CallOption) (*SetConversationMaxSeqResp, error)
	GetConversationIDs(ctx context.Context, in *GetConversationIDsReq, opts ...grpc.CallOption) (*GetConversationIDsResp, error)
	SetConversations(ctx context.Context, in *SetConversationsReq, opts ...grpc.CallOption) (*SetConversationsResp, error)
	GetUserConversationIDsHash(ctx context.Context, in *GetUserConversationIDsHashReq, opts ...grpc.CallOption) (*GetUserConversationIDsHashResp, error)
	GetConversationsByConversationID(ctx context.Context, in *GetConversationsByConversationIDReq, opts ...grpc.CallOption) (*GetConversationsByConversationIDResp, error)
}

type conversationClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationClient(cc grpc.ClientConnInterface) ConversationClient {
	return &conversationClient{cc}
}

func (c *conversationClient) GetConversation(ctx context.Context, in *GetConversationReq, opts ...grpc.CallOption) (*GetConversationResp, error) {
	out := new(GetConversationResp)
	err := c.cc.Invoke(ctx, Conversation_GetConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetAllConversations(ctx context.Context, in *GetAllConversationsReq, opts ...grpc.CallOption) (*GetAllConversationsResp, error) {
	out := new(GetAllConversationsResp)
	err := c.cc.Invoke(ctx, Conversation_GetAllConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	out := new(GetConversationsResp)
	err := c.cc.Invoke(ctx, Conversation_GetConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversation(ctx context.Context, in *SetConversationReq, opts ...grpc.CallOption) (*SetConversationResp, error) {
	out := new(SetConversationResp)
	err := c.cc.Invoke(ctx, Conversation_SetConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetRecvMsgNotNotifyUserIDs(ctx context.Context, in *GetRecvMsgNotNotifyUserIDsReq, opts ...grpc.CallOption) (*GetRecvMsgNotNotifyUserIDsResp, error) {
	out := new(GetRecvMsgNotNotifyUserIDsResp)
	err := c.cc.Invoke(ctx, Conversation_GetRecvMsgNotNotifyUserIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) CreateSingleChatConversations(ctx context.Context, in *CreateSingleChatConversationsReq, opts ...grpc.CallOption) (*CreateSingleChatConversationsResp, error) {
	out := new(CreateSingleChatConversationsResp)
	err := c.cc.Invoke(ctx, Conversation_CreateSingleChatConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) CreateGroupChatConversations(ctx context.Context, in *CreateGroupChatConversationsReq, opts ...grpc.CallOption) (*CreateGroupChatConversationsResp, error) {
	out := new(CreateGroupChatConversationsResp)
	err := c.cc.Invoke(ctx, Conversation_CreateGroupChatConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversationMaxSeq(ctx context.Context, in *SetConversationMaxSeqReq, opts ...grpc.CallOption) (*SetConversationMaxSeqResp, error) {
	out := new(SetConversationMaxSeqResp)
	err := c.cc.Invoke(ctx, Conversation_SetConversationMaxSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationIDs(ctx context.Context, in *GetConversationIDsReq, opts ...grpc.CallOption) (*GetConversationIDsResp, error) {
	out := new(GetConversationIDsResp)
	err := c.cc.Invoke(ctx, Conversation_GetConversationIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversations(ctx context.Context, in *SetConversationsReq, opts ...grpc.CallOption) (*SetConversationsResp, error) {
	out := new(SetConversationsResp)
	err := c.cc.Invoke(ctx, Conversation_SetConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetUserConversationIDsHash(ctx context.Context, in *GetUserConversationIDsHashReq, opts ...grpc.CallOption) (*GetUserConversationIDsHashResp, error) {
	out := new(GetUserConversationIDsHashResp)
	err := c.cc.Invoke(ctx, Conversation_GetUserConversationIDsHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationsByConversationID(ctx context.Context, in *GetConversationsByConversationIDReq, opts ...grpc.CallOption) (*GetConversationsByConversationIDResp, error) {
	out := new(GetConversationsByConversationIDResp)
	err := c.cc.Invoke(ctx, Conversation_GetConversationsByConversationID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServer is the server API for Conversation service.
// All implementations must embed UnimplementedConversationServer
// for forward compatibility
type ConversationServer interface {
	GetConversation(context.Context, *GetConversationReq) (*GetConversationResp, error)
	GetAllConversations(context.Context, *GetAllConversationsReq) (*GetAllConversationsResp, error)
	GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error)
	SetConversation(context.Context, *SetConversationReq) (*SetConversationResp, error)
	GetRecvMsgNotNotifyUserIDs(context.Context, *GetRecvMsgNotNotifyUserIDsReq) (*GetRecvMsgNotNotifyUserIDsResp, error)
	CreateSingleChatConversations(context.Context, *CreateSingleChatConversationsReq) (*CreateSingleChatConversationsResp, error)
	CreateGroupChatConversations(context.Context, *CreateGroupChatConversationsReq) (*CreateGroupChatConversationsResp, error)
	SetConversationMaxSeq(context.Context, *SetConversationMaxSeqReq) (*SetConversationMaxSeqResp, error)
	GetConversationIDs(context.Context, *GetConversationIDsReq) (*GetConversationIDsResp, error)
	SetConversations(context.Context, *SetConversationsReq) (*SetConversationsResp, error)
	GetUserConversationIDsHash(context.Context, *GetUserConversationIDsHashReq) (*GetUserConversationIDsHashResp, error)
	GetConversationsByConversationID(context.Context, *GetConversationsByConversationIDReq) (*GetConversationsByConversationIDResp, error)
	mustEmbedUnimplementedConversationServer()
}

// UnimplementedConversationServer must be embedded to have forward compatible implementations.
type UnimplementedConversationServer struct {
}

func (UnimplementedConversationServer) GetConversation(context.Context, *GetConversationReq) (*GetConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversation not implemented")
}
func (UnimplementedConversationServer) GetAllConversations(context.Context, *GetAllConversationsReq) (*GetAllConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConversations not implemented")
}
func (UnimplementedConversationServer) GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedConversationServer) SetConversation(context.Context, *SetConversationReq) (*SetConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversation not implemented")
}
func (UnimplementedConversationServer) GetRecvMsgNotNotifyUserIDs(context.Context, *GetRecvMsgNotNotifyUserIDsReq) (*GetRecvMsgNotNotifyUserIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecvMsgNotNotifyUserIDs not implemented")
}
func (UnimplementedConversationServer) CreateSingleChatConversations(context.Context, *CreateSingleChatConversationsReq) (*CreateSingleChatConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleChatConversations not implemented")
}
func (UnimplementedConversationServer) CreateGroupChatConversations(context.Context, *CreateGroupChatConversationsReq) (*CreateGroupChatConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChatConversations not implemented")
}
func (UnimplementedConversationServer) SetConversationMaxSeq(context.Context, *SetConversationMaxSeqReq) (*SetConversationMaxSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversationMaxSeq not implemented")
}
func (UnimplementedConversationServer) GetConversationIDs(context.Context, *GetConversationIDsReq) (*GetConversationIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationIDs not implemented")
}
func (UnimplementedConversationServer) SetConversations(context.Context, *SetConversationsReq) (*SetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversations not implemented")
}
func (UnimplementedConversationServer) GetUserConversationIDsHash(context.Context, *GetUserConversationIDsHashReq) (*GetUserConversationIDsHashResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConversationIDsHash not implemented")
}
func (UnimplementedConversationServer) GetConversationsByConversationID(context.Context, *GetConversationsByConversationIDReq) (*GetConversationsByConversationIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationsByConversationID not implemented")
}
func (UnimplementedConversationServer) mustEmbedUnimplementedConversationServer() {}

// UnsafeConversationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversationServer will
// result in compilation errors.
type UnsafeConversationServer interface {
	mustEmbedUnimplementedConversationServer()
}

func RegisterConversationServer(s grpc.ServiceRegistrar, srv ConversationServer) {
	s.RegisterService(&Conversation_ServiceDesc, srv)
}

func _Conversation_GetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversation(ctx, req.(*GetConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetAllConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetAllConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetAllConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetAllConversations(ctx, req.(*GetAllConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversations(ctx, req.(*GetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_SetConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversation(ctx, req.(*SetConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetRecvMsgNotNotifyUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecvMsgNotNotifyUserIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetRecvMsgNotNotifyUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetRecvMsgNotNotifyUserIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetRecvMsgNotNotifyUserIDs(ctx, req.(*GetRecvMsgNotNotifyUserIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_CreateSingleChatConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSingleChatConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).CreateSingleChatConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_CreateSingleChatConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).CreateSingleChatConversations(ctx, req.(*CreateSingleChatConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_CreateGroupChatConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupChatConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).CreateGroupChatConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_CreateGroupChatConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).CreateGroupChatConversations(ctx, req.(*CreateGroupChatConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversationMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationMaxSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversationMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_SetConversationMaxSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversationMaxSeq(ctx, req.(*SetConversationMaxSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetConversationIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationIDs(ctx, req.(*GetConversationIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_SetConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversations(ctx, req.(*SetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetUserConversationIDsHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserConversationIDsHashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetUserConversationIDsHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetUserConversationIDsHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetUserConversationIDsHash(ctx, req.(*GetUserConversationIDsHashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationsByConversationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsByConversationIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationsByConversationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_GetConversationsByConversationID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationsByConversationID(ctx, req.(*GetConversationsByConversationIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Conversation_ServiceDesc is the grpc.ServiceDesc for Conversation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conversation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OpenIMServer.conversation.conversation",
	HandlerType: (*ConversationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversation",
			Handler:    _Conversation_GetConversation_Handler,
		},
		{
			MethodName: "GetAllConversations",
			Handler:    _Conversation_GetAllConversations_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _Conversation_GetConversations_Handler,
		},
		{
			MethodName: "SetConversation",
			Handler:    _Conversation_SetConversation_Handler,
		},
		{
			MethodName: "GetRecvMsgNotNotifyUserIDs",
			Handler:    _Conversation_GetRecvMsgNotNotifyUserIDs_Handler,
		},
		{
			MethodName: "CreateSingleChatConversations",
			Handler:    _Conversation_CreateSingleChatConversations_Handler,
		},
		{
			MethodName: "CreateGroupChatConversations",
			Handler:    _Conversation_CreateGroupChatConversations_Handler,
		},
		{
			MethodName: "SetConversationMaxSeq",
			Handler:    _Conversation_SetConversationMaxSeq_Handler,
		},
		{
			MethodName: "GetConversationIDs",
			Handler:    _Conversation_GetConversationIDs_Handler,
		},
		{
			MethodName: "SetConversations",
			Handler:    _Conversation_SetConversations_Handler,
		},
		{
			MethodName: "GetUserConversationIDsHash",
			Handler:    _Conversation_GetUserConversationIDsHash_Handler,
		},
		{
			MethodName: "GetConversationsByConversationID",
			Handler:    _Conversation_GetConversationsByConversationID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation/conversation.proto",
}
