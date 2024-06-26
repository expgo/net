// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/authmode.proto

package config

import (
	fmt "fmt"
	_ "github.com/expgo/net/proto/ext"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AuthMode int32

const (
	AuthModeStatic AuthMode = 0
	AuthModeLDAP   AuthMode = 1
)

var AuthMode_name = map[int32]string{
	0: "AUTH_MODE_STATIC",
	1: "AUTH_MODE_LDAP",
}

var AuthMode_value = map[string]int32{
	"AUTH_MODE_STATIC": 0,
	"AUTH_MODE_LDAP":   1,
}

func (AuthMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_304f551a9ba2716f, []int{0}
}

func init() {
	proto.RegisterEnum("config.AuthMode", AuthMode_name, AuthMode_value)
}

func init() { proto.RegisterFile("config/authmode.proto", fileDescriptor_304f551a9ba2716f) }

var fileDescriptor_304f551a9ba2716f = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x4f, 0x2c, 0x2d, 0xc9, 0xc8, 0xcd, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x83, 0x08, 0x4b, 0x29, 0x17, 0xa5, 0x16, 0xe4, 0x17, 0xeb, 0x83, 0x05, 0x93,
	0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30, 0x0b, 0xa2, 0x58, 0x8a, 0x33, 0xb5,
	0xa2, 0x04, 0xc2, 0xd4, 0x2a, 0xe0, 0xe2, 0x70, 0x2c, 0x2d, 0xc9, 0xf0, 0xcd, 0x4f, 0x49, 0x15,
	0xd2, 0xe0, 0x12, 0x70, 0x0c, 0x0d, 0xf1, 0x88, 0xf7, 0xf5, 0x77, 0x71, 0x8d, 0x0f, 0x0e, 0x71,
	0x0c, 0xf1, 0x74, 0x16, 0x60, 0x90, 0x12, 0xea, 0x9a, 0xab, 0xc0, 0x07, 0x53, 0x13, 0x5c, 0x92,
	0x58, 0x92, 0x99, 0x2c, 0x64, 0xc2, 0xc5, 0x87, 0x50, 0xe9, 0xe3, 0xe2, 0x18, 0x20, 0xc0, 0x28,
	0xa5, 0xd0, 0x35, 0x57, 0x81, 0x07, 0xa6, 0x0e, 0x24, 0x76, 0xa9, 0x4f, 0x15, 0x85, 0x2f, 0xc5,
	0xb2, 0x62, 0x89, 0x1c, 0x83, 0x93, 0xed, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x1e, 0xca, 0x31, 0x9c,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31,
	0x5e, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x74, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x6a, 0x45, 0x41, 0x7a, 0xbe, 0x7e, 0x5e, 0x6a, 0x89, 0x3e,
	0xc4, 0x83, 0x49, 0x6c, 0x60, 0x77, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x94, 0x6d,
	0xcb, 0x08, 0x01, 0x00, 0x00,
}
