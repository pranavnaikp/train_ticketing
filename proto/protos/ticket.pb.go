// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: ticket.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request and response messages
type BookTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	SeatNumber int32  `protobuf:"varint,4,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
}

func (x *BookTicketRequest) Reset() {
	*x = BookTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookTicketRequest) ProtoMessage() {}

func (x *BookTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookTicketRequest.ProtoReflect.Descriptor instead.
func (*BookTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *BookTicketRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *BookTicketRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *BookTicketRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BookTicketRequest) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

type BookTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId string  `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	Message   string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FullName  string  `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email     string  `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	From      string  `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To        string  `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	Price     float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *BookTicketResponse) Reset() {
	*x = BookTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookTicketResponse) ProtoMessage() {}

func (x *BookTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookTicketResponse.ProtoReflect.Descriptor instead.
func (*BookTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *BookTicketResponse) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *BookTicketResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BookTicketResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *BookTicketResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BookTicketResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *BookTicketResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *BookTicketResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type TicketReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId  string  `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	From       string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Price      float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	SeatNumber int32   `protobuf:"varint,5,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	UserName   string  `protobuf:"bytes,6,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserEmail  string  `protobuf:"bytes,7,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *TicketReceipt) Reset() {
	*x = TicketReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketReceipt) ProtoMessage() {}

func (x *TicketReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketReceipt.ProtoReflect.Descriptor instead.
func (*TicketReceipt) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *TicketReceipt) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *TicketReceipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TicketReceipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TicketReceipt) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TicketReceipt) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

func (x *TicketReceipt) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *TicketReceipt) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	SeatNumber int32  `protobuf:"varint,4,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetail) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserDetail) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDetail) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserDetail `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *UserResponse) GetUser() *UserDetail {
	if x != nil {
		return x.User
	}
	return nil
}

type SectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"` // "A" or "B"
}

func (x *SectionRequest) Reset() {
	*x = SectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionRequest) ProtoMessage() {}

func (x *SectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionRequest.ProtoReflect.Descriptor instead.
func (*SectionRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *SectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type SectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserDetail `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SectionResponse) Reset() {
	*x = SectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionResponse) ProtoMessage() {}

func (x *SectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionResponse.ProtoReflect.Descriptor instead.
func (*SectionResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *SectionResponse) GetUsers() []*UserDetail {
	if x != nil {
		return x.Users
	}
	return nil
}

type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	NewSeatNumber int32  `protobuf:"varint,2,opt,name=new_seat_number,json=newSeatNumber,proto3" json:"new_seat_number,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *ModifySeatRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeatNumber() int32 {
	if x != nil {
		return x.NewSeatNumber
	}
	return 0
}

type ModifySeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ModifySeatResponse) Reset() {
	*x = ModifySeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatResponse) ProtoMessage() {}

func (x *ModifySeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatResponse.ProtoReflect.Descriptor instead.
func (*ModifySeatResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{11}
}

func (x *ModifySeatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId string `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
}

func (x *ReceiptRequest) Reset() {
	*x = ReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptRequest) ProtoMessage() {}

func (x *ReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptRequest.ProtoReflect.Descriptor instead.
func (*ReceiptRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{12}
}

func (x *ReceiptRequest) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

type ReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt *TicketReceipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *ReceiptResponse) Reset() {
	*x = ReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticket_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptResponse) ProtoMessage() {}

func (x *ReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptResponse.ProtoReflect.Descriptor instead.
func (*ReceiptResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{13}
}

func (x *ReceiptResponse) GetReceipt() *TicketReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x86, 0x01,
	0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x6b, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x7f, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x3d, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x0f,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x29, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2e, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x11, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65,
	0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2e,
	0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f,
	0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x22,
	0x49, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x32, 0xfe, 0x03, 0x0a, 0x14, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1d, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_ticket_proto_goTypes = []interface{}{
	(*BookTicketRequest)(nil),  // 0: ticketbooking.BookTicketRequest
	(*BookTicketResponse)(nil), // 1: ticketbooking.BookTicketResponse
	(*TicketReceipt)(nil),      // 2: ticketbooking.TicketReceipt
	(*UserDetail)(nil),         // 3: ticketbooking.UserDetail
	(*UserRequest)(nil),        // 4: ticketbooking.UserRequest
	(*UserResponse)(nil),       // 5: ticketbooking.UserResponse
	(*SectionRequest)(nil),     // 6: ticketbooking.SectionRequest
	(*SectionResponse)(nil),    // 7: ticketbooking.SectionResponse
	(*RemoveUserRequest)(nil),  // 8: ticketbooking.RemoveUserRequest
	(*RemoveUserResponse)(nil), // 9: ticketbooking.RemoveUserResponse
	(*ModifySeatRequest)(nil),  // 10: ticketbooking.ModifySeatRequest
	(*ModifySeatResponse)(nil), // 11: ticketbooking.ModifySeatResponse
	(*ReceiptRequest)(nil),     // 12: ticketbooking.ReceiptRequest
	(*ReceiptResponse)(nil),    // 13: ticketbooking.ReceiptResponse
}
var file_ticket_proto_depIdxs = []int32{
	3,  // 0: ticketbooking.UserResponse.user:type_name -> ticketbooking.UserDetail
	3,  // 1: ticketbooking.SectionResponse.users:type_name -> ticketbooking.UserDetail
	2,  // 2: ticketbooking.ReceiptResponse.receipt:type_name -> ticketbooking.TicketReceipt
	0,  // 3: ticketbooking.TicketBookingService.BookTicket:input_type -> ticketbooking.BookTicketRequest
	4,  // 4: ticketbooking.TicketBookingService.GetUserDetail:input_type -> ticketbooking.UserRequest
	6,  // 5: ticketbooking.TicketBookingService.GetUsersBySection:input_type -> ticketbooking.SectionRequest
	8,  // 6: ticketbooking.TicketBookingService.RemoveUser:input_type -> ticketbooking.RemoveUserRequest
	10, // 7: ticketbooking.TicketBookingService.ModifyUserSeat:input_type -> ticketbooking.ModifySeatRequest
	12, // 8: ticketbooking.TicketBookingService.GetReceipt:input_type -> ticketbooking.ReceiptRequest
	1,  // 9: ticketbooking.TicketBookingService.BookTicket:output_type -> ticketbooking.BookTicketResponse
	5,  // 10: ticketbooking.TicketBookingService.GetUserDetail:output_type -> ticketbooking.UserResponse
	7,  // 11: ticketbooking.TicketBookingService.GetUsersBySection:output_type -> ticketbooking.SectionResponse
	9,  // 12: ticketbooking.TicketBookingService.RemoveUser:output_type -> ticketbooking.RemoveUserResponse
	11, // 13: ticketbooking.TicketBookingService.ModifyUserSeat:output_type -> ticketbooking.ModifySeatResponse
	13, // 14: ticketbooking.TicketBookingService.GetReceipt:output_type -> ticketbooking.ReceiptResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookTicketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookTicketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketReceipt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ticket_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
