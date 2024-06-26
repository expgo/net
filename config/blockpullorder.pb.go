// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/blockpullorder.proto

package config

import (
	fmt "fmt"
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

type BlockPullOrder int32

const (
	BlockPullOrderStandard BlockPullOrder = 0
	BlockPullOrderRandom   BlockPullOrder = 1
	BlockPullOrderInOrder  BlockPullOrder = 2
)

var BlockPullOrder_name = map[int32]string{
	0: "BLOCK_PULL_ORDER_STANDARD",
	1: "BLOCK_PULL_ORDER_RANDOM",
	2: "BLOCK_PULL_ORDER_IN_ORDER",
}

var BlockPullOrder_value = map[string]int32{
	"BLOCK_PULL_ORDER_STANDARD": 0,
	"BLOCK_PULL_ORDER_RANDOM":   1,
	"BLOCK_PULL_ORDER_IN_ORDER": 2,
}

func (BlockPullOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_592de28274533a6f, []int{0}
}

func init() {
	proto.RegisterEnum("config.BlockPullOrder", BlockPullOrder_name, BlockPullOrder_value)
}

func init() { proto.RegisterFile("config/blockpullorder.proto", fileDescriptor_592de28274533a6f) }

var fileDescriptor_592de28274533a6f = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0x28, 0xcd, 0xc9, 0xc9, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x4a, 0x29, 0x17, 0xa5, 0x16,
	0xe4, 0x17, 0xeb, 0x83, 0x05, 0x93, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x58, 0xeb, 0x10, 0x23, 0x17, 0x9f, 0x13, 0xc8, 0x94, 0x80, 0xd2, 0x9c, 0x1c, 0x7f,
	0x90, 0x29, 0x42, 0x96, 0x5c, 0x92, 0x4e, 0x3e, 0xfe, 0xce, 0xde, 0xf1, 0x01, 0xa1, 0x3e, 0x3e,
	0xf1, 0xfe, 0x41, 0x2e, 0xae, 0x41, 0xf1, 0xc1, 0x21, 0x8e, 0x7e, 0x2e, 0x8e, 0x41, 0x2e, 0x02,
	0x0c, 0x52, 0x52, 0x5d, 0x73, 0x15, 0xc4, 0x50, 0xb5, 0x04, 0x97, 0x24, 0xe6, 0xa5, 0x24, 0x16,
	0xa5, 0x08, 0x99, 0x72, 0x89, 0x63, 0x68, 0x0d, 0x72, 0xf4, 0x73, 0xf1, 0xf7, 0x15, 0x60, 0x94,
	0x92, 0xe8, 0x9a, 0xab, 0x20, 0x82, 0xaa, 0x31, 0x28, 0x31, 0x2f, 0x25, 0x3f, 0x57, 0xc8, 0x02,
	0x8b, 0x8d, 0x9e, 0x7e, 0x10, 0x86, 0x00, 0x93, 0x94, 0x64, 0xd7, 0x5c, 0x05, 0x51, 0x54, 0x8d,
	0x9e, 0x79, 0x60, 0x4a, 0x8a, 0x65, 0xc5, 0x12, 0x39, 0x06, 0x27, 0xdb, 0x13, 0x0f, 0xe5, 0x18,
	0x2e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x16, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xe9,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xd4, 0x8a, 0x82, 0xf4, 0x7c,
	0xfd, 0xbc, 0xd4, 0x12, 0x7d, 0x48, 0x40, 0x25, 0xb1, 0x81, 0x83, 0xc2, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x00, 0x88, 0x4b, 0x56, 0x01, 0x00, 0x00,
}
