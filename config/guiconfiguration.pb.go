// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/guiconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/expgo/net/proto/ext"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GUIConfiguration struct {
	Enabled                   bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled" xml:"enabled,attr" default:"true"`
	RawAddress                string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address" xml:"address" default:"127.0.0.1:8384"`
	RawUnixSocketPermissions  string   `protobuf:"bytes,3,opt,name=unix_socket_permissions,json=unixSocketPermissions,proto3" json:"unixSocketPermissions" xml:"unixSocketPermissions,omitempty"`
	User                      string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user" xml:"user,omitempty"`
	Password                  string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password" xml:"password,omitempty"`
	AuthMode                  AuthMode `protobuf:"varint,6,opt,name=auth_mode,json=authMode,proto3,enum=config.AuthMode" json:"authMode" xml:"authMode,omitempty"`
	RawUseTLS                 bool     `protobuf:"varint,7,opt,name=use_tls,json=useTls,proto3" json:"useTLS" xml:"tls,attr"`
	APIKey                    string   `protobuf:"bytes,8,opt,name=api_key,json=apiKey,proto3" json:"apiKey" xml:"apikey,omitempty"`
	InsecureAdminAccess       bool     `protobuf:"varint,9,opt,name=insecure_admin_access,json=insecureAdminAccess,proto3" json:"insecureAdminAccess" xml:"insecureAdminAccess,omitempty"`
	Theme                     string   `protobuf:"bytes,10,opt,name=theme,proto3" json:"theme" xml:"theme" default:"default"`
	Debugging                 bool     `protobuf:"varint,11,opt,name=debugging,proto3" json:"debugging" xml:"debugging,attr"`
	InsecureSkipHostCheck     bool     `protobuf:"varint,12,opt,name=insecure_skip_host_check,json=insecureSkipHostCheck,proto3" json:"insecureSkipHostcheck" xml:"insecureSkipHostcheck,omitempty"`
	InsecureAllowFrameLoading bool     `protobuf:"varint,13,opt,name=insecure_allow_frame_loading,json=insecureAllowFrameLoading,proto3" json:"insecureAllowFrameLoading" xml:"insecureAllowFrameLoading,omitempty"`
}

func (m *GUIConfiguration) Reset()         { *m = GUIConfiguration{} }
func (m *GUIConfiguration) String() string { return proto.CompactTextString(m) }
func (*GUIConfiguration) ProtoMessage()    {}
func (*GUIConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9586d611855d64, []int{0}
}
func (m *GUIConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GUIConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GUIConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GUIConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GUIConfiguration.Merge(m, src)
}
func (m *GUIConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *GUIConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_GUIConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_GUIConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GUIConfiguration)(nil), "config.GUIConfiguration")
}

func init() { proto.RegisterFile("lib/config/guiconfiguration.proto", fileDescriptor_2a9586d611855d64) }

var fileDescriptor_2a9586d611855d64 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0x5b, 0x47, 0xb2, 0xb6, 0xae, 0x60, 0xb0, 0x4d, 0xcb, 0x04, 0x0d, 0xd7, 0x51, 0xd8,
	0xc2, 0x01, 0x02, 0x39, 0x71, 0x5a, 0x24, 0xf0, 0xa1, 0x80, 0x1c, 0x20, 0x4d, 0x60, 0x17, 0x08,
	0xe8, 0xfa, 0x92, 0x0b, 0xb1, 0x22, 0xd7, 0xd2, 0x42, 0xfc, 0x2b, 0x77, 0x09, 0x4b, 0x87, 0xf6,
	0x19, 0x0a, 0xf5, 0x5c, 0xa0, 0xcf, 0xd0, 0x4b, 0x5f, 0x21, 0x37, 0xe9, 0x54, 0xe4, 0xb4, 0x40,
	0xa4, 0x1b, 0x8f, 0x3c, 0xe6, 0x54, 0xec, 0xf2, 0x47, 0xa2, 0xac, 0xd4, 0xbd, 0xed, 0x7c, 0xf3,
	0xcd, 0x7c, 0x33, 0xc3, 0x19, 0x10, 0xdc, 0x75, 0x49, 0xef, 0xc0, 0x0e, 0xfc, 0x0b, 0xd2, 0x3f,
	0xe8, 0xc7, 0x24, 0x7b, 0xc5, 0x11, 0x62, 0x24, 0xf0, 0x3b, 0x61, 0x14, 0xb0, 0x40, 0xad, 0x67,
	0xe0, 0xed, 0x5b, 0x2b, 0x54, 0x14, 0xb3, 0x81, 0x17, 0x38, 0x38, 0xa3, 0xdc, 0x6e, 0xe2, 0x11,
	0xcb, 0x9e, 0xed, 0xb7, 0x3b, 0x60, 0xf7, 0x87, 0xf3, 0x97, 0xcf, 0x56, 0x13, 0xa9, 0x3d, 0xd0,
	0xc0, 0x3e, 0xea, 0xb9, 0xd8, 0xd1, 0x94, 0x3d, 0x65, 0x7f, 0xfb, 0xf8, 0x45, 0xc2, 0x61, 0x01,
	0xa5, 0x1c, 0xde, 0x1d, 0x79, 0xee, 0x51, 0x3b, 0xb7, 0x1f, 0x20, 0xc6, 0xa2, 0xf6, 0x9e, 0x83,
	0x2f, 0x50, 0xec, 0xb2, 0xa3, 0x36, 0x8b, 0x62, 0xdc, 0x4e, 0xa6, 0xc6, 0xce, 0xaa, 0xff, 0xfd,
	0xd4, 0xd8, 0x12, 0x0e, 0xb3, 0xc8, 0xa2, 0xfe, 0x02, 0x1a, 0xc8, 0x71, 0x22, 0x4c, 0xa9, 0xf6,
	0xd1, 0x9e, 0xb2, 0xdf, 0x3c, 0xb6, 0xe7, 0x1c, 0x02, 0x13, 0x5d, 0x76, 0x33, 0x54, 0x28, 0xe6,
	0x84, 0x94, 0xc3, 0x6f, 0xa4, 0x62, 0x6e, 0xaf, 0x88, 0x3d, 0x3a, 0x7c, 0xd2, 0x79, 0xd8, 0x79,
	0xd8, 0x79, 0x74, 0xf4, 0xf4, 0xf1, 0xd3, 0x6f, 0xdb, 0xef, 0xa7, 0x46, 0xab, 0x0a, 0x4d, 0x66,
	0xc6, 0x4a, 0x52, 0xb3, 0x48, 0xa9, 0xfe, 0xa3, 0x80, 0x2f, 0x63, 0x9f, 0x8c, 0x2c, 0x1a, 0xd8,
	0x43, 0xcc, 0xac, 0x10, 0x47, 0x1e, 0xa1, 0x94, 0x04, 0x3e, 0xd5, 0x3e, 0x96, 0xf5, 0xfc, 0xa1,
	0xcc, 0x39, 0xd4, 0x4c, 0x74, 0x79, 0xee, 0x93, 0xd1, 0x99, 0x64, 0xbd, 0x5a, 0x92, 0x12, 0x0e,
	0x6f, 0xc6, 0x9b, 0x1c, 0x29, 0x87, 0x5f, 0xcb, 0x62, 0x37, 0x7a, 0x1f, 0x04, 0x1e, 0x61, 0xd8,
	0x0b, 0xd9, 0x58, 0x8c, 0x08, 0x5e, 0xc3, 0x99, 0xcc, 0x8c, 0x0f, 0x16, 0x60, 0x6e, 0x96, 0x57,
	0x9f, 0x83, 0xad, 0x98, 0xe2, 0x48, 0xdb, 0x92, 0x4d, 0x1c, 0x26, 0x1c, 0x4a, 0x3b, 0xe5, 0xf0,
	0xf3, 0xac, 0x2c, 0x8a, 0xa3, 0x6a, 0x15, 0xad, 0x2a, 0x64, 0x4a, 0xbe, 0xfa, 0x1a, 0x6c, 0x87,
	0x88, 0xd2, 0xcb, 0x20, 0x72, 0xb4, 0x1b, 0x32, 0xd7, 0xf7, 0x09, 0x87, 0x25, 0x96, 0x72, 0xa8,
	0xc9, 0x7c, 0x05, 0x50, 0xcd, 0xa9, 0x5e, 0x85, 0xcd, 0x32, 0x56, 0xf5, 0x40, 0x53, 0x6c, 0xa4,
	0x25, 0x56, 0x52, 0xab, 0xef, 0x29, 0xfb, 0xad, 0xc3, 0xdd, 0x4e, 0xb6, 0xaa, 0x9d, 0x6e, 0xcc,
	0x06, 0x3f, 0x06, 0x0e, 0xce, 0xe4, 0x50, 0x6e, 0x95, 0x72, 0x05, 0xb0, 0x26, 0x77, 0x15, 0x36,
	0xcb, 0x58, 0x15, 0x83, 0x46, 0x4c, 0xb1, 0xc5, 0x5c, 0xaa, 0x35, 0xe4, 0x3a, 0x9f, 0xce, 0x39,
	0x6c, 0x8a, 0xc1, 0x52, 0xfc, 0xd3, 0xe9, 0x59, 0xc2, 0x61, 0x3d, 0x96, 0xaf, 0x94, 0xc3, 0x96,
	0x54, 0x61, 0x2e, 0xcd, 0xd6, 0x3a, 0x99, 0x1a, 0xdb, 0x85, 0x91, 0x4e, 0x8d, 0x9c, 0x37, 0x99,
	0x19, 0xcb, 0x70, 0x53, 0x82, 0x2e, 0x15, 0x32, 0x28, 0x24, 0xd6, 0x10, 0x8f, 0xb5, 0x6d, 0x39,
	0x30, 0x21, 0x53, 0xef, 0xbe, 0x7a, 0x79, 0x82, 0xc7, 0x42, 0x03, 0x85, 0xe4, 0x04, 0x8f, 0x53,
	0x0e, 0xbf, 0xc8, 0x3a, 0x09, 0xc9, 0x10, 0x8f, 0xab, 0x7d, 0xec, 0xae, 0x83, 0x93, 0x99, 0x91,
	0x67, 0x30, 0xf3, 0x78, 0xf5, 0x77, 0x05, 0xdc, 0x24, 0x3e, 0xc5, 0x76, 0x1c, 0x61, 0x0b, 0x39,
	0x1e, 0xf1, 0x2d, 0x64, 0xdb, 0xe2, 0x8e, 0x9a, 0xb2, 0x39, 0x2b, 0xe1, 0xf0, 0xb3, 0x82, 0xd0,
	0x15, 0xfe, 0xae, 0x74, 0xa7, 0x1c, 0xde, 0x93, 0xc2, 0x1b, 0x7c, 0xd5, 0x2a, 0xee, 0xfc, 0x27,
	0xc3, 0xdc, 0x94, 0x5c, 0x3d, 0x01, 0x37, 0xd8, 0x00, 0x7b, 0x58, 0x03, 0xb2, 0xf5, 0xef, 0x12,
	0x0e, 0x33, 0x20, 0xe5, 0xf0, 0x4e, 0x36, 0x53, 0x61, 0xad, 0x9c, 0x6e, 0xfe, 0x10, 0x37, 0xdb,
	0xc8, 0xdf, 0x66, 0x16, 0xa2, 0x9e, 0x83, 0xa6, 0x83, 0x7b, 0x71, 0xbf, 0x4f, 0xfc, 0xbe, 0xf6,
	0x89, 0xec, 0xea, 0x49, 0xc2, 0xe1, 0x12, 0x2c, 0xb7, 0xb9, 0x44, 0xca, 0xcf, 0xd5, 0xaa, 0x42,
	0xe6, 0x32, 0x48, 0xfd, 0x5b, 0x01, 0x5a, 0x39, 0x39, 0x3a, 0x24, 0xa1, 0x35, 0x08, 0x28, 0xb3,
	0xec, 0x01, 0xb6, 0x87, 0xda, 0x8e, 0x94, 0xf9, 0x55, 0xdc, 0x75, 0xc1, 0x39, 0x1b, 0x92, 0xf0,
	0x45, 0x40, 0x99, 0x24, 0x94, 0x77, 0xbd, 0xd1, 0xbb, 0x76, 0xd7, 0xd7, 0x70, 0xd2, 0xa9, 0xb1,
	0x59, 0xc4, 0xbc, 0x02, 0x3f, 0x13, 0xb0, 0xfa, 0x97, 0x02, 0xbe, 0x5a, 0x7e, 0x73, 0xd7, 0x0d,
	0x2e, 0xad, 0x8b, 0x08, 0x79, 0xd8, 0x72, 0x03, 0xe4, 0x88, 0x21, 0x7d, 0x2a, 0xab, 0xff, 0x39,
	0xe1, 0xf0, 0x56, 0xf9, 0x75, 0x04, 0xed, 0xb9, 0x60, 0x9d, 0x66, 0xa4, 0x94, 0xc3, 0xfb, 0xd5,
	0x05, 0x58, 0x67, 0x54, 0xbb, 0xb8, 0xf7, 0x3f, 0x78, 0xe6, 0x87, 0xe5, 0x8e, 0x4f, 0xde, 0xbc,
	0xd3, 0x6b, 0xb3, 0x77, 0x7a, 0xed, 0xcd, 0x5c, 0x57, 0x66, 0x73, 0x5d, 0xf9, 0x6d, 0xa1, 0xd7,
	0xfe, 0x5c, 0xe8, 0xca, 0x6c, 0xa1, 0xd7, 0xde, 0x2e, 0xf4, 0xda, 0xeb, 0xfb, 0x7d, 0xc2, 0x06,
	0x71, 0xaf, 0x63, 0x07, 0xde, 0x01, 0x1d, 0xfb, 0x36, 0x1b, 0x10, 0xbf, 0xbf, 0xf2, 0x5a, 0xfe,
	0xc1, 0x7a, 0x75, 0xf9, 0xbb, 0x7a, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x8c, 0xef,
	0xc0, 0x01, 0x07, 0x00, 0x00,
}

func (m *GUIConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GUIConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GUIConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InsecureAllowFrameLoading {
		i--
		if m.InsecureAllowFrameLoading {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.InsecureSkipHostCheck {
		i--
		if m.InsecureSkipHostCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if m.Debugging {
		i--
		if m.Debugging {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.Theme) > 0 {
		i -= len(m.Theme)
		copy(dAtA[i:], m.Theme)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.Theme)))
		i--
		dAtA[i] = 0x52
	}
	if m.InsecureAdminAccess {
		i--
		if m.InsecureAdminAccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.APIKey) > 0 {
		i -= len(m.APIKey)
		copy(dAtA[i:], m.APIKey)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.APIKey)))
		i--
		dAtA[i] = 0x42
	}
	if m.RawUseTLS {
		i--
		if m.RawUseTLS {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.AuthMode != 0 {
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(m.AuthMode))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RawUnixSocketPermissions) > 0 {
		i -= len(m.RawUnixSocketPermissions)
		copy(dAtA[i:], m.RawUnixSocketPermissions)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.RawUnixSocketPermissions)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RawAddress) > 0 {
		i -= len(m.RawAddress)
		copy(dAtA[i:], m.RawAddress)
		i = encodeVarintGuiconfiguration(dAtA, i, uint64(len(m.RawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuiconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuiconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GUIConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.RawAddress)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.RawUnixSocketPermissions)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.AuthMode != 0 {
		n += 1 + sovGuiconfiguration(uint64(m.AuthMode))
	}
	if m.RawUseTLS {
		n += 2
	}
	l = len(m.APIKey)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.InsecureAdminAccess {
		n += 2
	}
	l = len(m.Theme)
	if l > 0 {
		n += 1 + l + sovGuiconfiguration(uint64(l))
	}
	if m.Debugging {
		n += 2
	}
	if m.InsecureSkipHostCheck {
		n += 2
	}
	if m.InsecureAllowFrameLoading {
		n += 2
	}
	return n
}

func sovGuiconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuiconfiguration(x uint64) (n int) {
	return sovGuiconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GUIConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuiconfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GUIConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GUIConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawUnixSocketPermissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawUnixSocketPermissions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthMode", wireType)
			}
			m.AuthMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthMode |= AuthMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawUseTLS", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RawUseTLS = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureAdminAccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InsecureAdminAccess = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Theme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Theme = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debugging", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Debugging = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureSkipHostCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InsecureSkipHostCheck = bool(v != 0)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureAllowFrameLoading", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InsecureAllowFrameLoading = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGuiconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuiconfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGuiconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuiconfiguration
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGuiconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGuiconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuiconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuiconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuiconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuiconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuiconfiguration = fmt.Errorf("proto: unexpected end of group")
)
