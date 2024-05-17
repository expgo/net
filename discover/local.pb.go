// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: discover/local.proto

package discover

import (
	fmt "fmt"
	_ "github.com/expgo/net/proto/ext"
	github_com_expgo_net_protocol "github.com/expgo/net/protocol"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Announce struct {
	ID         github_com_expgo_net_protocol.DeviceID `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/expgo/net/protocol.DeviceID" json:"id" xml:"id"`
	Addresses  []string                               `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses" xml:"address"`
	InstanceID int64                                  `protobuf:"varint,3,opt,name=instance_id,json=instanceId,proto3" json:"instanceId" xml:"instanceId"`
}

func (m *Announce) Reset()         { *m = Announce{} }
func (m *Announce) String() string { return proto.CompactTextString(m) }
func (*Announce) ProtoMessage()    {}
func (*Announce) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb113366fcb5a76, []int{0}
}
func (m *Announce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Announce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Announce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Announce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announce.Merge(m, src)
}
func (m *Announce) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Announce) XXX_DiscardUnknown() {
	xxx_messageInfo_Announce.DiscardUnknown(m)
}

var xxx_messageInfo_Announce proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Announce)(nil), "discover.Announce")
}

func init() { proto.RegisterFile("discover/local.proto", fileDescriptor_dcb113366fcb5a76) }

var fileDescriptor_dcb113366fcb5a76 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0x03, 0x31,
	0x00, 0x86, 0x2f, 0x29, 0x48, 0x1b, 0x15, 0xe4, 0x70, 0x28, 0x05, 0x93, 0x72, 0x16, 0xe9, 0xd4,
	0x1b, 0xdc, 0x44, 0x14, 0x8f, 0x5b, 0x6e, 0x93, 0x8e, 0x2e, 0xb5, 0x4d, 0xe2, 0x19, 0xb8, 0x5e,
	0xca, 0xe5, 0x5a, 0xfa, 0x08, 0x82, 0x8b, 0xf4, 0x09, 0x7c, 0x9c, 0x8e, 0x19, 0xc5, 0x21, 0xd0,
	0xbb, 0xad, 0x63, 0x9f, 0x40, 0xcc, 0xb5, 0xb6, 0x83, 0xdb, 0x9f, 0x2f, 0x5f, 0xc2, 0xcf, 0x8f,
	0xce, 0x99, 0x50, 0x54, 0xce, 0x78, 0xe6, 0x27, 0x92, 0x0e, 0x93, 0xde, 0x24, 0x93, 0xb9, 0x74,
	0xeb, 0x3b, 0xda, 0xba, 0xcc, 0xf8, 0x44, 0x2a, 0xdf, 0xe2, 0xd1, 0xf4, 0xc5, 0x8f, 0x65, 0x2c,
	0xed, 0xc1, 0xa6, 0x4a, 0x6f, 0x35, 0xf8, 0x3c, 0xaf, 0xa2, 0xf7, 0x0e, 0x51, 0xfd, 0x21, 0x4d,
	0xe5, 0x34, 0xa5, 0xdc, 0x7d, 0x46, 0x50, 0xb0, 0x26, 0x68, 0x83, 0xee, 0x49, 0xf0, 0xb8, 0x34,
	0xc4, 0xf9, 0x36, 0xe4, 0x2a, 0x16, 0xf9, 0xeb, 0x74, 0xd4, 0xa3, 0x72, 0xec, 0xf3, 0xf9, 0x24,
	0x96, 0x7e, 0xca, 0xf3, 0xea, 0x7f, 0x2a, 0x93, 0x5e, 0xc8, 0x67, 0x82, 0xf2, 0x28, 0x2c, 0x0c,
	0x81, 0x51, 0xb8, 0x36, 0x04, 0x0a, 0xb6, 0x31, 0xa4, 0x3e, 0x1f, 0x27, 0x37, 0x9e, 0x60, 0xde,
	0x9b, 0xee, 0x80, 0x85, 0xee, 0xc0, 0x28, 0xec, 0x43, 0xc1, 0xdc, 0x5b, 0xd4, 0x18, 0x32, 0x96,
	0x71, 0xa5, 0xb8, 0x6a, 0xc2, 0x76, 0xad, 0xdb, 0x08, 0xf0, 0xda, 0x90, 0x3d, 0xdc, 0x18, 0x72,
	0x6a, 0xdf, 0x6e, 0x89, 0xd7, 0xdf, 0xdf, 0xb9, 0x03, 0x74, 0x2c, 0x52, 0x95, 0x0f, 0x53, 0xca,
	0x07, 0x82, 0x35, 0x6b, 0x6d, 0xd0, 0xad, 0x05, 0x77, 0x85, 0x21, 0x28, 0xda, 0x62, 0x5b, 0x01,
	0xed, 0xa4, 0xe8, 0xb7, 0xca, 0x59, 0x55, 0xe5, 0x0f, 0x79, 0x0b, 0xdd, 0x39, 0xf0, 0xfb, 0x07,
	0x76, 0x70, 0xbf, 0x5c, 0x61, 0x47, 0xaf, 0xb0, 0xb3, 0x2c, 0x30, 0xd0, 0x05, 0x06, 0x1f, 0x25,
	0x76, 0x3e, 0x4b, 0x0c, 0x74, 0x89, 0x9d, 0xaf, 0x12, 0x3b, 0x4f, 0x17, 0xff, 0xce, 0xb1, 0x9b,
	0x7f, 0x74, 0x64, 0x87, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x51, 0x88, 0x33, 0x84, 0xa7,
	0x01, 0x00, 0x00,
}

func (m *Announce) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Announce) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Announce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InstanceID != 0 {
		i = encodeVarintLocal(dAtA, i, uint64(m.InstanceID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintLocal(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.ID.ProtoSize()
		i -= size
		if _, err := m.ID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLocal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLocal(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Announce) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.ProtoSize()
	n += 1 + l + sovLocal(uint64(l))
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovLocal(uint64(l))
		}
	}
	if m.InstanceID != 0 {
		n += 1 + sovLocal(uint64(m.InstanceID))
	}
	return n
}

func sovLocal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocal(x uint64) (n int) {
	return sovLocal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Announce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocal
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
			return fmt.Errorf("proto: Announce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Announce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
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
				return ErrInvalidLengthLocal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceID", wireType)
			}
			m.InstanceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InstanceID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLocal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocal
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
func skipLocal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocal
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
					return 0, ErrIntOverflowLocal
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
					return 0, ErrIntOverflowLocal
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
				return 0, ErrInvalidLengthLocal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocal = fmt.Errorf("proto: unexpected end of group")
)
