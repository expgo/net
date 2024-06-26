// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/ldapconfiguration.proto

package config

import (
	fmt "fmt"
	_ "github.com/expgo/net/proto/ext"
	proto "github.com/gogo/protobuf/proto"
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

type LDAPConfiguration struct {
	Address            string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address" xml:"address,omitempty"`
	BindDN             string        `protobuf:"bytes,2,opt,name=bind_dn,json=bindDn,proto3" json:"bindDN" xml:"bindDN,omitempty"`
	Transport          LDAPTransport `protobuf:"varint,3,opt,name=transport,proto3,enum=config.LDAPTransport" json:"transport" xml:"transport,omitempty"`
	InsecureSkipVerify bool          `protobuf:"varint,4,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecureSkipVerify" xml:"insecureSkipVerify,omitempty" default:"false"`
	SearchBaseDN       string        `protobuf:"bytes,5,opt,name=search_base_dn,json=searchBaseDn,proto3" json:"searchBaseDN" xml:"searchBaseDN,omitempty"`
	SearchFilter       string        `protobuf:"bytes,6,opt,name=search_filter,json=searchFilter,proto3" json:"searchFilter" xml:"searchFilter,omitempty"`
}

func (m *LDAPConfiguration) Reset()         { *m = LDAPConfiguration{} }
func (m *LDAPConfiguration) String() string { return proto.CompactTextString(m) }
func (*LDAPConfiguration) ProtoMessage()    {}
func (*LDAPConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b40e5bfc77016865, []int{0}
}
func (m *LDAPConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LDAPConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LDAPConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LDAPConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LDAPConfiguration.Merge(m, src)
}
func (m *LDAPConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *LDAPConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_LDAPConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_LDAPConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LDAPConfiguration)(nil), "config.LDAPConfiguration")
}

func init() { proto.RegisterFile("config/ldapconfiguration.proto", fileDescriptor_b40e5bfc77016865) }

var fileDescriptor_b40e5bfc77016865 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x81, 0xba, 0xc4, 0x2a, 0x15, 0x35, 0x50, 0x4c, 0xa8, 0x7c, 0x51, 0xe4, 0x21, 0x03,
	0x4a, 0xa4, 0xb2, 0x15, 0x31, 0xd4, 0x54, 0x0c, 0x08, 0x21, 0xe4, 0x42, 0x07, 0x96, 0xc8, 0x8e,
	0xcf, 0xe9, 0xa9, 0xce, 0x9d, 0x75, 0x77, 0xae, 0x52, 0x7e, 0x05, 0xf4, 0x17, 0x74, 0xe3, 0xaf,
	0x74, 0x8b, 0x47, 0xa6, 0x93, 0x9a, 0x6c, 0x1e, 0x3d, 0x32, 0xa1, 0x9c, 0x9d, 0xc6, 0x4e, 0xa3,
	0x6e, 0xef, 0x7d, 0xdf, 0x7b, 0xdf, 0xfb, 0xee, 0x3d, 0x5b, 0xb7, 0x06, 0x04, 0x87, 0x68, 0xd8,
	0x8b, 0x02, 0x2f, 0x2e, 0xc2, 0x84, 0x7a, 0x1c, 0x11, 0xdc, 0x8d, 0x29, 0xe1, 0xc4, 0xd0, 0x0a,
	0xb0, 0xd9, 0xac, 0xd4, 0x71, 0xea, 0x61, 0x16, 0x13, 0xca, 0x8b, 0x9a, 0x66, 0x03, 0x8e, 0xcb,
	0xb0, 0xfd, 0x5b, 0xd3, 0x77, 0x3e, 0x1f, 0x1d, 0x7e, 0xfd, 0x50, 0x95, 0x32, 0xbe, 0xeb, 0x9b,
	0x5e, 0x10, 0x50, 0xc8, 0x98, 0xa9, 0xb6, 0xd4, 0x4e, 0xc3, 0x79, 0x97, 0x09, 0xb0, 0x80, 0x72,
	0x01, 0x5e, 0x8e, 0x47, 0xd1, 0x41, 0xbb, 0xcc, 0xdf, 0x90, 0x11, 0xe2, 0x70, 0x14, 0xf3, 0x8b,
	0x76, 0x36, 0xb1, 0x77, 0xee, 0xa0, 0xee, 0xa2, 0xd1, 0x20, 0xfa, 0xa6, 0x8f, 0x70, 0xd0, 0x0f,
	0xb0, 0xf9, 0x40, 0xca, 0x9e, 0x4c, 0x05, 0xd0, 0x1c, 0x84, 0x83, 0xa3, 0x2f, 0x99, 0x00, 0x9a,
	0x2f, 0xa3, 0x5c, 0x80, 0x5d, 0xa9, 0x5f, 0xa4, 0x75, 0xf9, 0xa7, 0xab, 0x60, 0x3e, 0xb1, 0xcb,
	0xbe, 0xcb, 0xd4, 0x2e, 0xb5, 0xdc, 0x02, 0xc1, 0xc6, 0xb9, 0xde, 0xb8, 0x7d, 0xbb, 0xf9, 0xb0,
	0xa5, 0x76, 0xb6, 0xf7, 0x5f, 0x74, 0x8b, 0xc5, 0x74, 0xe7, 0xaf, 0xfe, 0xb6, 0x20, 0x9d, 0xc3,
	0x4c, 0x80, 0x65, 0x6d, 0x2e, 0xc0, 0x2b, 0x69, 0xe1, 0x16, 0xa9, 0xbb, 0x78, 0xb6, 0x06, 0x77,
	0x97, 0xed, 0xc6, 0x1f, 0x55, 0x7f, 0x8e, 0x30, 0x83, 0x83, 0x84, 0xc2, 0x3e, 0x3b, 0x43, 0x71,
	0xff, 0x1c, 0x52, 0x14, 0x5e, 0x98, 0x8f, 0x5a, 0x6a, 0xe7, 0xb1, 0x93, 0x64, 0x02, 0x18, 0x0b,
	0xfe, 0xf8, 0x0c, 0xc5, 0x27, 0x92, 0xcd, 0x05, 0xd8, 0x97, 0x53, 0xef, 0x52, 0x95, 0xf1, 0xad,
	0x00, 0x86, 0x5e, 0x12, 0xf1, 0x83, 0x76, 0xe8, 0x45, 0x0c, 0xce, 0xed, 0xec, 0xdd, 0xd7, 0xf0,
	0x6f, 0x62, 0x6f, 0xc8, 0x4a, 0x77, 0xcd, 0x48, 0xe3, 0x4a, 0xd5, 0xb7, 0x19, 0xf4, 0xe8, 0xe0,
	0xb4, 0xef, 0x7b, 0x0c, 0xce, 0x4f, 0xb3, 0x21, 0x4f, 0xf3, 0x73, 0x2a, 0xc0, 0xd6, 0xb1, 0x64,
	0x1c, 0x8f, 0x41, 0x79, 0xa0, 0x2d, 0x56, 0xc9, 0x73, 0x01, 0xf6, 0xa4, 0xdb, 0x2a, 0x58, 0x5f,
	0xd3, 0xee, 0x7a, 0x2a, 0x9f, 0xd8, 0x35, 0xa5, 0xcb, 0xd4, 0xae, 0x4d, 0x72, 0xab, 0x2c, 0x36,
	0x88, 0xfe, 0xa4, 0x74, 0x18, 0xa2, 0x88, 0x43, 0x6a, 0x6a, 0xd2, 0xe0, 0xa7, 0xa5, 0xa1, 0x8f,
	0x12, 0x5f, 0x31, 0x54, 0x80, 0x6b, 0x0d, 0xad, 0x52, 0x6e, 0x4d, 0xc7, 0x79, 0x7f, 0x7d, 0x63,
	0x29, 0xe9, 0x8d, 0xa5, 0x5c, 0x4f, 0x2d, 0x35, 0x9d, 0x5a, 0xea, 0xaf, 0x99, 0xa5, 0x5c, 0xcd,
	0x2c, 0x35, 0x9d, 0x59, 0xca, 0xdf, 0x99, 0xa5, 0xfc, 0x78, 0x3d, 0x44, 0xfc, 0x34, 0xf1, 0xbb,
	0x03, 0x32, 0xea, 0xc1, 0x71, 0x3c, 0x24, 0x3d, 0x0c, 0x79, 0xaf, 0xf8, 0xb0, 0x7c, 0x4d, 0xfe,
	0x59, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x04, 0x86, 0x79, 0xaa, 0x03, 0x00, 0x00,
}

func (m *LDAPConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LDAPConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LDAPConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SearchFilter) > 0 {
		i -= len(m.SearchFilter)
		copy(dAtA[i:], m.SearchFilter)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.SearchFilter)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SearchBaseDN) > 0 {
		i -= len(m.SearchBaseDN)
		copy(dAtA[i:], m.SearchBaseDN)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.SearchBaseDN)))
		i--
		dAtA[i] = 0x2a
	}
	if m.InsecureSkipVerify {
		i--
		if m.InsecureSkipVerify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Transport != 0 {
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(m.Transport))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BindDN) > 0 {
		i -= len(m.BindDN)
		copy(dAtA[i:], m.BindDN)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.BindDN)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLdapconfiguration(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLdapconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovLdapconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LDAPConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	l = len(m.BindDN)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	if m.Transport != 0 {
		n += 1 + sovLdapconfiguration(uint64(m.Transport))
	}
	if m.InsecureSkipVerify {
		n += 2
	}
	l = len(m.SearchBaseDN)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	l = len(m.SearchFilter)
	if l > 0 {
		n += 1 + l + sovLdapconfiguration(uint64(l))
	}
	return n
}

func sovLdapconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLdapconfiguration(x uint64) (n int) {
	return sovLdapconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LDAPConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLdapconfiguration
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
			return fmt.Errorf("proto: LDAPConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LDAPConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindDN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BindDN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			m.Transport = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Transport |= LDAPTransport(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureSkipVerify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
			m.InsecureSkipVerify = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchBaseDN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchBaseDN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchFilter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLdapconfiguration
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
				return ErrInvalidLengthLdapconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLdapconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchFilter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLdapconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLdapconfiguration
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
func skipLdapconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLdapconfiguration
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
					return 0, ErrIntOverflowLdapconfiguration
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
					return 0, ErrIntOverflowLdapconfiguration
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
				return 0, ErrInvalidLengthLdapconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLdapconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLdapconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLdapconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLdapconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLdapconfiguration = fmt.Errorf("proto: unexpected end of group")
)
