// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator.proto

package validator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of integer strictly equal this value.
	LengthEq *int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	// Field value disable validation until is not empty
	Optional             *bool    `protobuf:"varint,17,opt,name=optional" json:"optional,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidator) Reset()         { *m = FieldValidator{} }
func (m *FieldValidator) String() string { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()    {}
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{0}
}

func (m *FieldValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValidator.Unmarshal(m, b)
}
func (m *FieldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValidator.Marshal(b, m, deterministic)
}
func (m *FieldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidator.Merge(m, src)
}
func (m *FieldValidator) XXX_Size() int {
	return xxx_messageInfo_FieldValidator.Size(m)
}
func (m *FieldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidator.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidator proto.InternalMessageInfo

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

func (m *FieldValidator) GetOptional() bool {
	if m != nil && m.Optional != nil {
		return *m.Optional
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_bf1c6ec7c0d80dd5) }

var fileDescriptor_bf1c6ec7c0d80dd5 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x86, 0x15, 0xda, 0x6d, 0x93, 0xd9, 0x76, 0xbb, 0x58, 0x20, 0x4d, 0x8b, 0x2a, 0x22, 0xb8,
	0xe4, 0x80, 0x52, 0x89, 0x23, 0x47, 0x50, 0xd8, 0xcb, 0xf2, 0xa1, 0x1c, 0x38, 0x70, 0x89, 0x42,
	0x77, 0xd6, 0xb5, 0xe4, 0xd8, 0xa9, 0x3d, 0x8b, 0x96, 0xbf, 0xc3, 0xdf, 0x84, 0x03, 0x8a, 0x43,
	0x36, 0xa9, 0xd4, 0xe3, 0x3c, 0xcf, 0x9b, 0x71, 0x6c, 0xbd, 0x70, 0xf1, 0xb3, 0xd6, 0x6a, 0x53,
	0xb3, 0x75, 0x79, 0xeb, 0x2c, 0x5b, 0x91, 0x1c, 0xc0, 0x55, 0x2a, 0xad, 0x95, 0x9a, 0x6e, 0x82,
	0xf8, 0xb1, 0xdb, 0xde, 0x6c, 0xc8, 0xdf, 0x3a, 0xd5, 0x1e, 0xc2, 0xaf, 0x7e, 0x1f, 0xc3, 0xe2,
	0xa3, 0x22, 0xbd, 0xf9, 0x36, 0x7c, 0x24, 0x9e, 0xc1, 0xcc, 0x91, 0xa4, 0x3d, 0x46, 0x69, 0x94,
	0x25, 0x65, 0x3f, 0x88, 0xe7, 0x70, 0xa2, 0x0c, 0x57, 0x92, 0xf1, 0x49, 0x1a, 0x65, 0x47, 0xe5,
	0x4c, 0x19, 0x5e, 0xf1, 0x80, 0x35, 0xe3, 0xd1, 0x01, 0xaf, 0x59, 0x5c, 0x03, 0x34, 0x5e, 0x56,
	0xb4, 0x57, 0x9e, 0x3d, 0x1e, 0xa7, 0x51, 0x16, 0x97, 0x49, 0xe3, 0x65, 0x11, 0x80, 0x78, 0x09,
	0xf3, 0xbb, 0x5d, 0x53, 0x9b, 0x8a, 0x9c, 0xb3, 0x0e, 0x67, 0xe1, 0x20, 0x08, 0xa8, 0xe8, 0x88,
	0xb8, 0x84, 0x78, 0xab, 0x6d, 0x1d, 0xce, 0x3b, 0x49, 0xa3, 0x2c, 0x2a, 0x4f, 0xc3, 0xbc, 0xe2,
	0x51, 0x69, 0xc6, 0xd3, 0x89, 0x5a, 0xb3, 0x78, 0x0d, 0xe7, 0xbd, 0xa2, 0xd6, 0x2b, 0x6d, 0x0d,
	0xc6, 0xc1, 0x9f, 0x05, 0x58, 0xf4, 0x4c, 0xbc, 0x80, 0x64, 0x58, 0x4d, 0x98, 0x84, 0x40, 0xfc,
	0x7f, 0x37, 0x8d, 0x52, 0x33, 0x21, 0x4c, 0xe4, 0x9a, 0x49, 0x64, 0xb0, 0xf4, 0xec, 0x94, 0x91,
	0x95, 0xb1, 0x5c, 0x51, 0xd3, 0xf2, 0x2f, 0x9c, 0x87, 0xab, 0x2d, 0x7a, 0xfe, 0xd9, 0x72, 0xd1,
	0x51, 0xf1, 0x06, 0x84, 0xa3, 0x96, 0x6a, 0xa6, 0x4d, 0x75, 0x6b, 0x77, 0x86, 0xab, 0x46, 0x19,
	0x3c, 0x0b, 0x2f, 0xb4, 0x1c, 0xcc, 0x87, 0x4e, 0x7c, 0x52, 0xe6, 0xb1, 0x74, 0xbd, 0xc7, 0xf3,
	0xc7, 0xd2, 0xf5, 0xbe, 0xfb, 0x45, 0x4d, 0x46, 0xf2, 0x5d, 0xf7, 0x36, 0x8b, 0x10, 0x8a, 0x7b,
	0xb0, 0xe2, 0x89, 0xd4, 0x8c, 0x17, 0x53, 0xb9, 0x9e, 0x4a, 0xba, 0xc7, 0xe5, 0x54, 0x16, 0xf7,
	0xe2, 0x0a, 0x62, 0xdb, 0xb2, 0xb2, 0xa6, 0xd6, 0xf8, 0x34, 0x5c, 0xea, 0x30, 0xbf, 0xfb, 0x0a,
	0xb3, 0x6d, 0xd7, 0x11, 0x71, 0x9d, 0xf7, 0x85, 0xca, 0x87, 0x42, 0xe5, 0xa1, 0x3b, 0x5f, 0x42,
	0xd0, 0xe3, 0xdf, 0x3f, 0x5d, 0x09, 0xe6, 0x6f, 0x2f, 0xf3, 0xb1, 0x93, 0x0f, 0xcb, 0x55, 0xf6,
	0x8b, 0xde, 0xcf, 0xbf, 0x8f, 0x2d, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x82, 0xc2, 0x25,
	0xc2, 0x02, 0x00, 0x00,
}
