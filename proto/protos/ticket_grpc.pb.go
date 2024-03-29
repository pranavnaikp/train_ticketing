// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: ticket.proto

package protos

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

// TicketBookingServiceClient is the client API for TicketBookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketBookingServiceClient interface {
	BookTicket(ctx context.Context, in *BookTicketRequest, opts ...grpc.CallOption) (*BookTicketResponse, error)
	GetUserDetail(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SectionResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifyUserSeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
	GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*ReceiptResponse, error)
}

type ticketBookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketBookingServiceClient(cc grpc.ClientConnInterface) TicketBookingServiceClient {
	return &ticketBookingServiceClient{cc}
}

func (c *ticketBookingServiceClient) BookTicket(ctx context.Context, in *BookTicketRequest, opts ...grpc.CallOption) (*BookTicketResponse, error) {
	out := new(BookTicketResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/BookTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketBookingServiceClient) GetUserDetail(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/GetUserDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketBookingServiceClient) GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SectionResponse, error) {
	out := new(SectionResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/GetUsersBySection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketBookingServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketBookingServiceClient) ModifyUserSeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/ModifyUserSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketBookingServiceClient) GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*ReceiptResponse, error) {
	out := new(ReceiptResponse)
	err := c.cc.Invoke(ctx, "/ticketbooking.TicketBookingService/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketBookingServiceServer is the server API for TicketBookingService service.
// All implementations must embed UnimplementedTicketBookingServiceServer
// for forward compatibility
type TicketBookingServiceServer interface {
	BookTicket(context.Context, *BookTicketRequest) (*BookTicketResponse, error)
	GetUserDetail(context.Context, *UserRequest) (*UserResponse, error)
	GetUsersBySection(context.Context, *SectionRequest) (*SectionResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifyUserSeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	GetReceipt(context.Context, *ReceiptRequest) (*ReceiptResponse, error)
	mustEmbedUnimplementedTicketBookingServiceServer()
}

// UnimplementedTicketBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketBookingServiceServer struct {
}

func (UnimplementedTicketBookingServiceServer) BookTicket(context.Context, *BookTicketRequest) (*BookTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedTicketBookingServiceServer) GetUserDetail(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetail not implemented")
}
func (UnimplementedTicketBookingServiceServer) GetUsersBySection(context.Context, *SectionRequest) (*SectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersBySection not implemented")
}
func (UnimplementedTicketBookingServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketBookingServiceServer) ModifyUserSeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserSeat not implemented")
}
func (UnimplementedTicketBookingServiceServer) GetReceipt(context.Context, *ReceiptRequest) (*ReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTicketBookingServiceServer) mustEmbedUnimplementedTicketBookingServiceServer() {}

// UnsafeTicketBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketBookingServiceServer will
// result in compilation errors.
type UnsafeTicketBookingServiceServer interface {
	mustEmbedUnimplementedTicketBookingServiceServer()
}

func RegisterTicketBookingServiceServer(s grpc.ServiceRegistrar, srv TicketBookingServiceServer) {
	s.RegisterService(&TicketBookingService_ServiceDesc, srv)
}

func _TicketBookingService_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/BookTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).BookTicket(ctx, req.(*BookTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketBookingService_GetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).GetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/GetUserDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).GetUserDetail(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketBookingService_GetUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).GetUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/GetUsersBySection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).GetUsersBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketBookingService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketBookingService_ModifyUserSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).ModifyUserSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/ModifyUserSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).ModifyUserSeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketBookingService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketBookingServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticketbooking.TicketBookingService/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketBookingServiceServer).GetReceipt(ctx, req.(*ReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketBookingService_ServiceDesc is the grpc.ServiceDesc for TicketBookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketBookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketbooking.TicketBookingService",
	HandlerType: (*TicketBookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTicket",
			Handler:    _TicketBookingService_BookTicket_Handler,
		},
		{
			MethodName: "GetUserDetail",
			Handler:    _TicketBookingService_GetUserDetail_Handler,
		},
		{
			MethodName: "GetUsersBySection",
			Handler:    _TicketBookingService_GetUsersBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketBookingService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifyUserSeat",
			Handler:    _TicketBookingService_ModifyUserSeat_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TicketBookingService_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}
