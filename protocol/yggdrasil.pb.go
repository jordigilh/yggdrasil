// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: protocol/yggdrasil.proto

package protocol

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

type Event int32

const (
	// UNKNOWN is an undeclared event; protobuf requires the presence of a zero value. It shouldn't be used.
	Event_UNKNOWN Event = 0
	// RECEIVED_DISCONNECT is sent by the dispatcher to workers when the service received the "disconnect" command.
	Event_RECEIVED_DISCONNECT Event = 1
	// UNEXPECTED_DISCONNECT is sent by the dispatcher when the service unexpectedly disconnects from the network.
	Event_UNEXPECTED_DISCONNECT Event = 2
	// CONNECTION_RESTORED is sent by the dispatcher when the service re-connects to the network.
	Event_CONNECTION_RESTORED Event = 3
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "UNKNOWN",
		1: "RECEIVED_DISCONNECT",
		2: "UNEXPECTED_DISCONNECT",
		3: "CONNECTION_RESTORED",
	}
	Event_value = map[string]int32{
		"UNKNOWN":               0,
		"RECEIVED_DISCONNECT":   1,
		"UNEXPECTED_DISCONNECT": 2,
		"CONNECTION_RESTORED":   3,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_yggdrasil_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_protocol_yggdrasil_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{0}
}

// An Empty message.
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{0}
}

// A RegistrationRequest message contains information necessary for a client to
// request registration with the dispatcher for a specified work type.
type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of work the worker is capable of handling.
	Handler string `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	// The PID of the worker.
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	// Whether or not the worker requires detached content processing.
	DetachedContent bool `protobuf:"varint,3,opt,name=detached_content,json=detachedContent,proto3" json:"detached_content,omitempty"`
	// A set of features a worker can announce during registration.
	Features map[string]string `protobuf:"bytes,4,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationRequest) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *RegistrationRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *RegistrationRequest) GetDetachedContent() bool {
	if x != nil {
		return x.DetachedContent
	}
	return false
}

func (x *RegistrationRequest) GetFeatures() map[string]string {
	if x != nil {
		return x.Features
	}
	return nil
}

// A RegistrationResponse message contains the result of a registration request.
type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not the dispatcher accepted the registration request.
	Registered bool `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	// The address on which the worker can be dialed to assign work.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *RegistrationResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// A Data message contains data and metadata suitable to exchange data between
// the dispatcher and a worker.
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The MQTT message ID that generated this message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// Optional key-value pairs to be included in the data message.
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The data payload.
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// The MQTT message ID this message is in response to.
	ResponseTo string `protobuf:"bytes,4,opt,name=response_to,json=responseTo,proto3" json:"response_to,omitempty"`
	// The destination of the message.
	Directive string `protobuf:"bytes,5,opt,name=directive,proto3" json:"directive,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Data) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Data) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Data) GetResponseTo() string {
	if x != nil {
		return x.ResponseTo
	}
	return ""
}

func (x *Data) GetDirective() string {
	if x != nil {
		return x.Directive
	}
	return ""
}

// A Receipt message is sent as a successful response to a Send method.
type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{4}
}

// An EventReceipt message is sent as a successful response to the NotifyEven method.
type EventReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventReceipt) Reset() {
	*x = EventReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReceipt) ProtoMessage() {}

func (x *EventReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReceipt.ProtoReflect.Descriptor instead.
func (*EventReceipt) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{5}
}

// An EventNotification message contains data and metadata describing event occurring in the dispatcher.
type EventNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of an event.
	Name Event `protobuf:"varint,1,opt,name=name,proto3,enum=yggdrasil.Event" json:"name,omitempty"`
	// Details of the event.
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	// Optional key-value pairs.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventNotification) Reset() {
	*x = EventNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventNotification) ProtoMessage() {}

func (x *EventNotification) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventNotification.ProtoReflect.Descriptor instead.
func (*EventNotification) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{6}
}

func (x *EventNotification) GetName() Event {
	if x != nil {
		return x.Name
	}
	return Event_UNKNOWN
}

func (x *EventNotification) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *EventNotification) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_protocol_yggdrasil_proto protoreflect.FileDescriptor

var file_protocol_yggdrasil_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x79, 0x67, 0x67, 0x64, 0x72,
	0x61, 0x73, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xf3,
	0x01, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x65,
	0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x09, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x11, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x46,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x61, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x43,
	0x45, 0x49, 0x56, 0x45, 0x44, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x44, 0x10, 0x03, 0x32, 0x8a, 0x01, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x79,
	0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e,
	0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x22, 0x00, 0x32, 0x7f, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x2d, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69,
	0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73,
	0x69, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x79, 0x67,
	0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x2f, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_yggdrasil_proto_rawDescOnce sync.Once
	file_protocol_yggdrasil_proto_rawDescData = file_protocol_yggdrasil_proto_rawDesc
)

func file_protocol_yggdrasil_proto_rawDescGZIP() []byte {
	file_protocol_yggdrasil_proto_rawDescOnce.Do(func() {
		file_protocol_yggdrasil_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_yggdrasil_proto_rawDescData)
	})
	return file_protocol_yggdrasil_proto_rawDescData
}

var file_protocol_yggdrasil_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_yggdrasil_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protocol_yggdrasil_proto_goTypes = []interface{}{
	(Event)(0),                   // 0: yggdrasil.Event
	(*Empty)(nil),                // 1: yggdrasil.Empty
	(*RegistrationRequest)(nil),  // 2: yggdrasil.RegistrationRequest
	(*RegistrationResponse)(nil), // 3: yggdrasil.RegistrationResponse
	(*Data)(nil),                 // 4: yggdrasil.Data
	(*Receipt)(nil),              // 5: yggdrasil.Receipt
	(*EventReceipt)(nil),         // 6: yggdrasil.EventReceipt
	(*EventNotification)(nil),    // 7: yggdrasil.EventNotification
	nil,                          // 8: yggdrasil.RegistrationRequest.FeaturesEntry
	nil,                          // 9: yggdrasil.Data.MetadataEntry
	nil,                          // 10: yggdrasil.EventNotification.MetadataEntry
}
var file_protocol_yggdrasil_proto_depIdxs = []int32{
	8,  // 0: yggdrasil.RegistrationRequest.features:type_name -> yggdrasil.RegistrationRequest.FeaturesEntry
	9,  // 1: yggdrasil.Data.metadata:type_name -> yggdrasil.Data.MetadataEntry
	0,  // 2: yggdrasil.EventNotification.name:type_name -> yggdrasil.Event
	10, // 3: yggdrasil.EventNotification.metadata:type_name -> yggdrasil.EventNotification.MetadataEntry
	2,  // 4: yggdrasil.Dispatcher.Register:input_type -> yggdrasil.RegistrationRequest
	4,  // 5: yggdrasil.Dispatcher.Send:input_type -> yggdrasil.Data
	4,  // 6: yggdrasil.Worker.Send:input_type -> yggdrasil.Data
	7,  // 7: yggdrasil.Worker.NotifyEvent:input_type -> yggdrasil.EventNotification
	3,  // 8: yggdrasil.Dispatcher.Register:output_type -> yggdrasil.RegistrationResponse
	5,  // 9: yggdrasil.Dispatcher.Send:output_type -> yggdrasil.Receipt
	5,  // 10: yggdrasil.Worker.Send:output_type -> yggdrasil.Receipt
	6,  // 11: yggdrasil.Worker.NotifyEvent:output_type -> yggdrasil.EventReceipt
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protocol_yggdrasil_proto_init() }
func file_protocol_yggdrasil_proto_init() {
	if File_protocol_yggdrasil_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_yggdrasil_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protocol_yggdrasil_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_protocol_yggdrasil_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationResponse); i {
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
		file_protocol_yggdrasil_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_protocol_yggdrasil_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_protocol_yggdrasil_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReceipt); i {
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
		file_protocol_yggdrasil_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventNotification); i {
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
			RawDescriptor: file_protocol_yggdrasil_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protocol_yggdrasil_proto_goTypes,
		DependencyIndexes: file_protocol_yggdrasil_proto_depIdxs,
		EnumInfos:         file_protocol_yggdrasil_proto_enumTypes,
		MessageInfos:      file_protocol_yggdrasil_proto_msgTypes,
	}.Build()
	File_protocol_yggdrasil_proto = out.File
	file_protocol_yggdrasil_proto_rawDesc = nil
	file_protocol_yggdrasil_proto_goTypes = nil
	file_protocol_yggdrasil_proto_depIdxs = nil
}
