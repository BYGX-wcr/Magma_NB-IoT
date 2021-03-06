//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: eap.proto

package protos

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EapType int32

const (
	// Mandatory EAP Method types
	EapType_Reserved      EapType = 0
	EapType_Identity      EapType = 1
	EapType_Notification  EapType = 2
	EapType_Legacy_Nak    EapType = 3
	EapType_MD5_Challenge EapType = 4
	EapType_Expanded      EapType = 254
	EapType_Experimental  EapType = 255
	// EAP Method Authenticator types
	EapType_TLS      EapType = 13
	EapType_SIM      EapType = 18
	EapType_AKA      EapType = 23
	EapType_AKAPrime EapType = 50
)

// Enum value maps for EapType.
var (
	EapType_name = map[int32]string{
		0:   "Reserved",
		1:   "Identity",
		2:   "Notification",
		3:   "Legacy_Nak",
		4:   "MD5_Challenge",
		254: "Expanded",
		255: "Experimental",
		13:  "TLS",
		18:  "SIM",
		23:  "AKA",
		50:  "AKAPrime",
	}
	EapType_value = map[string]int32{
		"Reserved":      0,
		"Identity":      1,
		"Notification":  2,
		"Legacy_Nak":    3,
		"MD5_Challenge": 4,
		"Expanded":      254,
		"Experimental":  255,
		"TLS":           13,
		"SIM":           18,
		"AKA":           23,
		"AKAPrime":      50,
	}
)

func (x EapType) Enum() *EapType {
	p := new(EapType)
	*p = x
	return p
}

func (x EapType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EapType) Descriptor() protoreflect.EnumDescriptor {
	return file_eap_proto_enumTypes[0].Descriptor()
}

func (EapType) Type() protoreflect.EnumType {
	return &file_eap_proto_enumTypes[0]
}

func (x EapType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EapType.Descriptor instead.
func (EapType) EnumDescriptor() ([]byte, []int) {
	return file_eap_proto_rawDescGZIP(), []int{0}
}

type EapCode int32

const (
	EapCode_Undefined EapCode = 0
	EapCode_Request   EapCode = 1
	EapCode_Response  EapCode = 2
	EapCode_Success   EapCode = 3
	EapCode_Failure   EapCode = 4
)

// Enum value maps for EapCode.
var (
	EapCode_name = map[int32]string{
		0: "Undefined",
		1: "Request",
		2: "Response",
		3: "Success",
		4: "Failure",
	}
	EapCode_value = map[string]int32{
		"Undefined": 0,
		"Request":   1,
		"Response":  2,
		"Success":   3,
		"Failure":   4,
	}
)

func (x EapCode) Enum() *EapCode {
	p := new(EapCode)
	*p = x
	return p
}

func (x EapCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EapCode) Descriptor() protoreflect.EnumDescriptor {
	return file_eap_proto_enumTypes[1].Descriptor()
}

func (EapCode) Type() protoreflect.EnumType {
	return &file_eap_proto_enumTypes[1]
}

func (x EapCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EapCode.Descriptor instead.
func (EapCode) EnumDescriptor() ([]byte, []int) {
	return file_eap_proto_rawDescGZIP(), []int{1}
}

type Eap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Ctx     *Context `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
}

func (x *Eap) Reset() {
	*x = Eap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Eap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Eap) ProtoMessage() {}

func (x *Eap) ProtoReflect() protoreflect.Message {
	mi := &file_eap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Eap.ProtoReflect.Descriptor instead.
func (*Eap) Descriptor() ([]byte, []int) {
	return file_eap_proto_rawDescGZIP(), []int{0}
}

func (x *Eap) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Eap) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

type EapIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Ctx     *Context `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Method  uint32   `protobuf:"varint,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *EapIdentity) Reset() {
	*x = EapIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EapIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EapIdentity) ProtoMessage() {}

func (x *EapIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_eap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EapIdentity.ProtoReflect.Descriptor instead.
func (*EapIdentity) Descriptor() ([]byte, []int) {
	return file_eap_proto_rawDescGZIP(), []int{1}
}

func (x *EapIdentity) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *EapIdentity) GetCtx() *Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *EapIdentity) GetMethod() uint32 {
	if x != nil {
		return x.Method
	}
	return 0
}

type EapMethodList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []byte `protobuf:"bytes,1,opt,name=methods,proto3" json:"methods,omitempty"`
}

func (x *EapMethodList) Reset() {
	*x = EapMethodList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EapMethodList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EapMethodList) ProtoMessage() {}

func (x *EapMethodList) ProtoReflect() protoreflect.Message {
	mi := &file_eap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EapMethodList.ProtoReflect.Descriptor instead.
func (*EapMethodList) Descriptor() ([]byte, []int) {
	return file_eap_proto_rawDescGZIP(), []int{2}
}

func (x *EapMethodList) GetMethods() []byte {
	if x != nil {
		return x.Methods
	}
	return nil
}

var File_eap_proto protoreflect.FileDescriptor

var file_eap_proto_rawDesc = []byte{
	0x0a, 0x09, 0x65, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x61, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x03, 0x65, 0x61, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x22, 0x67,
	0x0a, 0x0c, 0x65, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x65, 0x61, 0x70, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2a, 0xa6, 0x01, 0x0a, 0x08, 0x65, 0x61, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x4e, 0x61, 0x6b, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x4d, 0x44, 0x35, 0x5f, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x10, 0xfe,
	0x01, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x10, 0xff, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x10, 0x0d, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x49, 0x4d, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4b, 0x41, 0x10, 0x17, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x4b, 0x41, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x10, 0x32, 0x2a, 0x4e, 0x0a,
	0x08, 0x65, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x04, 0x32, 0xc3, 0x01,
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x3e, 0x0a, 0x0f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x65, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x61,
	0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x61, 0x61, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x1a, 0x0f, 0x2e, 0x61, 0x61, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x10, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1b, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x65, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x32, 0xc0, 0x01, 0x0a, 0x0a, 0x65, 0x61, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x0f, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70,
	0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x61,
	0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x1a, 0x0f, 0x2e,
	0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1b, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x65, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f,
	0x66, 0x65, 0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eap_proto_rawDescOnce sync.Once
	file_eap_proto_rawDescData = file_eap_proto_rawDesc
)

func file_eap_proto_rawDescGZIP() []byte {
	file_eap_proto_rawDescOnce.Do(func() {
		file_eap_proto_rawDescData = protoimpl.X.CompressGZIP(file_eap_proto_rawDescData)
	})
	return file_eap_proto_rawDescData
}

var file_eap_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eap_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eap_proto_goTypes = []interface{}{
	(EapType)(0),          // 0: aaa.protos.eap_type
	(EapCode)(0),          // 1: aaa.protos.eap_code
	(*Eap)(nil),           // 2: aaa.protos.eap
	(*EapIdentity)(nil),   // 3: aaa.protos.eap_identity
	(*EapMethodList)(nil), // 4: aaa.protos.eap_method_list
	(*Context)(nil),       // 5: aaa.protos.context
	(*Void)(nil),          // 6: aaa.protos.Void
}
var file_eap_proto_depIdxs = []int32{
	5, // 0: aaa.protos.eap.ctx:type_name -> aaa.protos.context
	5, // 1: aaa.protos.eap_identity.ctx:type_name -> aaa.protos.context
	3, // 2: aaa.protos.authenticator.handle_identity:input_type -> aaa.protos.eap_identity
	2, // 3: aaa.protos.authenticator.handle:input_type -> aaa.protos.eap
	6, // 4: aaa.protos.authenticator.supported_methods:input_type -> aaa.protos.Void
	3, // 5: aaa.protos.eap_router.handle_identity:input_type -> aaa.protos.eap_identity
	2, // 6: aaa.protos.eap_router.handle:input_type -> aaa.protos.eap
	6, // 7: aaa.protos.eap_router.supported_methods:input_type -> aaa.protos.Void
	2, // 8: aaa.protos.authenticator.handle_identity:output_type -> aaa.protos.eap
	2, // 9: aaa.protos.authenticator.handle:output_type -> aaa.protos.eap
	4, // 10: aaa.protos.authenticator.supported_methods:output_type -> aaa.protos.eap_method_list
	2, // 11: aaa.protos.eap_router.handle_identity:output_type -> aaa.protos.eap
	2, // 12: aaa.protos.eap_router.handle:output_type -> aaa.protos.eap
	4, // 13: aaa.protos.eap_router.supported_methods:output_type -> aaa.protos.eap_method_list
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eap_proto_init() }
func file_eap_proto_init() {
	if File_eap_proto != nil {
		return
	}
	file_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Eap); i {
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
		file_eap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EapIdentity); i {
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
		file_eap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EapMethodList); i {
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
			RawDescriptor: file_eap_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_eap_proto_goTypes,
		DependencyIndexes: file_eap_proto_depIdxs,
		EnumInfos:         file_eap_proto_enumTypes,
		MessageInfos:      file_eap_proto_msgTypes,
	}.Build()
	File_eap_proto = out.File
	file_eap_proto_rawDesc = nil
	file_eap_proto_goTypes = nil
	file_eap_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticatorClient interface {
	// handle_identity passes Identity EAP payload to corresponding method provider & returns corresponding
	// EAP result
	// NOTE: Identity Request is handled by APs & does not involve EAP Authenticator's support
	HandleIdentity(ctx context.Context, in *EapIdentity, opts ...grpc.CallOption) (*Eap, error)
	// handle handles passed EAP payload & returns corresponding EAP result
	Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error)
	// supported_methods returns sorted list (ascending, by type) of registered EAP Provider Methods
	SupportedMethods(ctx context.Context, in *Void, opts ...grpc.CallOption) (*EapMethodList, error)
}

type authenticatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorClient(cc grpc.ClientConnInterface) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) HandleIdentity(ctx context.Context, in *EapIdentity, opts ...grpc.CallOption) (*Eap, error) {
	out := new(Eap)
	err := c.cc.Invoke(ctx, "/aaa.protos.authenticator/handle_identity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error) {
	out := new(Eap)
	err := c.cc.Invoke(ctx, "/aaa.protos.authenticator/handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) SupportedMethods(ctx context.Context, in *Void, opts ...grpc.CallOption) (*EapMethodList, error) {
	out := new(EapMethodList)
	err := c.cc.Invoke(ctx, "/aaa.protos.authenticator/supported_methods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
type AuthenticatorServer interface {
	// handle_identity passes Identity EAP payload to corresponding method provider & returns corresponding
	// EAP result
	// NOTE: Identity Request is handled by APs & does not involve EAP Authenticator's support
	HandleIdentity(context.Context, *EapIdentity) (*Eap, error)
	// handle handles passed EAP payload & returns corresponding EAP result
	Handle(context.Context, *Eap) (*Eap, error)
	// supported_methods returns sorted list (ascending, by type) of registered EAP Provider Methods
	SupportedMethods(context.Context, *Void) (*EapMethodList, error)
}

// UnimplementedAuthenticatorServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServer struct {
}

func (*UnimplementedAuthenticatorServer) HandleIdentity(context.Context, *EapIdentity) (*Eap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleIdentity not implemented")
}
func (*UnimplementedAuthenticatorServer) Handle(context.Context, *Eap) (*Eap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (*UnimplementedAuthenticatorServer) SupportedMethods(context.Context, *Void) (*EapMethodList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupportedMethods not implemented")
}

func RegisterAuthenticatorServer(s *grpc.Server, srv AuthenticatorServer) {
	s.RegisterService(&_Authenticator_serviceDesc, srv)
}

func _Authenticator_HandleIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EapIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).HandleIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.authenticator/HandleIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).HandleIdentity(ctx, req.(*EapIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Eap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.authenticator/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Handle(ctx, req.(*Eap))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_SupportedMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).SupportedMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.authenticator/SupportedMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).SupportedMethods(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aaa.protos.authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handle_identity",
			Handler:    _Authenticator_HandleIdentity_Handler,
		},
		{
			MethodName: "handle",
			Handler:    _Authenticator_Handle_Handler,
		},
		{
			MethodName: "supported_methods",
			Handler:    _Authenticator_SupportedMethods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eap.proto",
}

// EapRouterClient is the client API for EapRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EapRouterClient interface {
	// handle_identity passes Identity EAP payload to corresponding method provider & returns corresponding
	// EAP result
	// NOTE: Identity Request is handled by APs & does not involve EAP Authenticator's support
	HandleIdentity(ctx context.Context, in *EapIdentity, opts ...grpc.CallOption) (*Eap, error)
	// handle handles passed EAP payload & returns corresponding EAP result
	Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error)
	// supported_methods returns sorted list (ascending, by type) of registered EAP Provider Methods
	SupportedMethods(ctx context.Context, in *Void, opts ...grpc.CallOption) (*EapMethodList, error)
}

type eapRouterClient struct {
	cc grpc.ClientConnInterface
}

func NewEapRouterClient(cc grpc.ClientConnInterface) EapRouterClient {
	return &eapRouterClient{cc}
}

func (c *eapRouterClient) HandleIdentity(ctx context.Context, in *EapIdentity, opts ...grpc.CallOption) (*Eap, error) {
	out := new(Eap)
	err := c.cc.Invoke(ctx, "/aaa.protos.eap_router/handle_identity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eapRouterClient) Handle(ctx context.Context, in *Eap, opts ...grpc.CallOption) (*Eap, error) {
	out := new(Eap)
	err := c.cc.Invoke(ctx, "/aaa.protos.eap_router/handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eapRouterClient) SupportedMethods(ctx context.Context, in *Void, opts ...grpc.CallOption) (*EapMethodList, error) {
	out := new(EapMethodList)
	err := c.cc.Invoke(ctx, "/aaa.protos.eap_router/supported_methods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EapRouterServer is the server API for EapRouter service.
type EapRouterServer interface {
	// handle_identity passes Identity EAP payload to corresponding method provider & returns corresponding
	// EAP result
	// NOTE: Identity Request is handled by APs & does not involve EAP Authenticator's support
	HandleIdentity(context.Context, *EapIdentity) (*Eap, error)
	// handle handles passed EAP payload & returns corresponding EAP result
	Handle(context.Context, *Eap) (*Eap, error)
	// supported_methods returns sorted list (ascending, by type) of registered EAP Provider Methods
	SupportedMethods(context.Context, *Void) (*EapMethodList, error)
}

// UnimplementedEapRouterServer can be embedded to have forward compatible implementations.
type UnimplementedEapRouterServer struct {
}

func (*UnimplementedEapRouterServer) HandleIdentity(context.Context, *EapIdentity) (*Eap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleIdentity not implemented")
}
func (*UnimplementedEapRouterServer) Handle(context.Context, *Eap) (*Eap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (*UnimplementedEapRouterServer) SupportedMethods(context.Context, *Void) (*EapMethodList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupportedMethods not implemented")
}

func RegisterEapRouterServer(s *grpc.Server, srv EapRouterServer) {
	s.RegisterService(&_EapRouter_serviceDesc, srv)
}

func _EapRouter_HandleIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EapIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EapRouterServer).HandleIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.eap_router/HandleIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EapRouterServer).HandleIdentity(ctx, req.(*EapIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _EapRouter_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Eap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EapRouterServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.eap_router/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EapRouterServer).Handle(ctx, req.(*Eap))
	}
	return interceptor(ctx, in, info, handler)
}

func _EapRouter_SupportedMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EapRouterServer).SupportedMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aaa.protos.eap_router/SupportedMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EapRouterServer).SupportedMethods(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _EapRouter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aaa.protos.eap_router",
	HandlerType: (*EapRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handle_identity",
			Handler:    _EapRouter_HandleIdentity_Handler,
		},
		{
			MethodName: "handle",
			Handler:    _EapRouter_Handle_Handler,
		},
		{
			MethodName: "supported_methods",
			Handler:    _EapRouter_SupportedMethods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eap.proto",
}
