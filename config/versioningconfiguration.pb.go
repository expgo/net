// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/versioningconfiguration.proto

package config

import (
	fmt "fmt"
	fs "github.com/expgo/net/fs"
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

// VersioningConfiguration is used in the code and for JSON serialization
type VersioningConfiguration struct {
	Type             string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type" xml:"type,attr"`
	Params           map[string]string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"params" xml:"parameter" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CleanupIntervalS int               `protobuf:"varint,3,opt,name=cleanup_interval_s,json=cleanupIntervalS,proto3,casttype=int" json:"cleanupIntervalS" xml:"cleanupIntervalS" default:"3600"`
	FSPath           string            `protobuf:"bytes,4,opt,name=fs_path,json=fsPath,proto3" json:"fsPath" xml:"fsPath"`
	FSType           fs.FilesystemType `protobuf:"varint,5,opt,name=fs_type,json=fsType,proto3,enum=fs.FilesystemType" json:"fsType" xml:"fsType"`
}

func (m *VersioningConfiguration) Reset()         { *m = VersioningConfiguration{} }
func (m *VersioningConfiguration) String() string { return proto.CompactTextString(m) }
func (*VersioningConfiguration) ProtoMessage()    {}
func (*VersioningConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_63dc6e661d64fb3c, []int{0}
}
func (m *VersioningConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersioningConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersioningConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersioningConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersioningConfiguration.Merge(m, src)
}
func (m *VersioningConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *VersioningConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_VersioningConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_VersioningConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersioningConfiguration)(nil), "config.VersioningConfiguration")
	proto.RegisterMapType((map[string]string)(nil), "config.VersioningConfiguration.ParametersEntry")
}

func init() {
	proto.RegisterFile("config/versioningconfiguration.proto", fileDescriptor_63dc6e661d64fb3c)
}

var fileDescriptor_63dc6e661d64fb3c = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x95, 0xfc, 0x55, 0xbc, 0x29, 0x49, 0x59, 0x0a, 0x15, 0x2e, 0x68, 0xcd, 0xe2, 0x16, 0x1f,
	0x8a, 0x14, 0x12, 0x28, 0xc5, 0xb4, 0x14, 0x5c, 0x9a, 0x52, 0xe8, 0x21, 0x28, 0xa1, 0x87, 0xf6,
	0x60, 0x36, 0xee, 0xca, 0x16, 0x91, 0x25, 0xa1, 0x5d, 0x1b, 0xab, 0xbf, 0x22, 0xf4, 0x17, 0xf4,
	0xe7, 0xf8, 0x66, 0x1f, 0x7b, 0x5a, 0x88, 0x7d, 0xd3, 0x51, 0xc7, 0xf4, 0x52, 0x76, 0x57, 0x11,
	0x21, 0x25, 0xb7, 0x79, 0x6f, 0xde, 0xbe, 0x99, 0xd9, 0x19, 0xd0, 0x1b, 0xc7, 0x91, 0x1f, 0x4c,
	0xdc, 0x05, 0x4d, 0x59, 0x10, 0x47, 0x41, 0x34, 0xd1, 0xc4, 0x3c, 0x25, 0x3c, 0x88, 0x23, 0x27,
	0x49, 0x63, 0x1e, 0xc3, 0x96, 0x26, 0x3b, 0xfb, 0x3e, 0x73, 0x79, 0x96, 0x50, 0xa6, 0xf9, 0x4e,
	0x9b, 0x2e, 0xb9, 0x0e, 0xf1, 0xdf, 0x06, 0x78, 0xf6, 0xb5, 0x32, 0xf9, 0x70, 0xd7, 0x04, 0xbe,
	0x05, 0x0d, 0xf9, 0xca, 0x32, 0xbb, 0x66, 0xbf, 0x3d, 0xec, 0xe7, 0x02, 0x29, 0x5c, 0x08, 0x74,
	0xb0, 0x9c, 0x85, 0x03, 0x2c, 0xc1, 0x2b, 0xc2, 0x79, 0x8a, 0xf3, 0x75, 0xaf, 0x5d, 0x21, 0x4f,
	0xa9, 0xe0, 0x95, 0x09, 0x40, 0x42, 0x52, 0x32, 0xa3, 0x9c, 0xa6, 0xcc, 0xaa, 0x75, 0xeb, 0xfd,
	0xbd, 0x23, 0xd7, 0xd1, 0x2d, 0x39, 0x0f, 0xd4, 0x74, 0x4e, 0xab, 0x17, 0x1f, 0x23, 0x9e, 0x66,
	0xc3, 0xf7, 0x2b, 0x81, 0x8c, 0xad, 0x40, 0x2d, 0x95, 0x60, 0xb9, 0x40, 0x2d, 0x65, 0xca, 0xaa,
	0x2e, 0xaa, 0x1a, 0xb8, 0x58, 0xf7, 0xca, 0xe4, 0xaf, 0x4d, 0xaf, 0x7c, 0xe0, 0xdd, 0xe9, 0x01,
	0xfe, 0x04, 0x70, 0x1c, 0x52, 0x12, 0xcd, 0x93, 0x51, 0x10, 0x71, 0x9a, 0x2e, 0x48, 0x38, 0x62,
	0x56, 0xbd, 0x6b, 0xf6, 0x9b, 0xc3, 0x2f, 0xb9, 0x40, 0x4f, 0xca, 0xec, 0xe7, 0x32, 0x79, 0x56,
	0x08, 0xf4, 0x42, 0x15, 0xb9, 0x9f, 0xc0, 0xdd, 0x1f, 0xd4, 0x27, 0xf3, 0x90, 0x0f, 0xf0, 0xf1,
	0xeb, 0xc3, 0x43, 0x7c, 0x23, 0x50, 0x3d, 0x88, 0xf8, 0xcd, 0xba, 0xd7, 0x90, 0xd8, 0xfb, 0xcf,
	0x09, 0x7e, 0x02, 0x8f, 0x7c, 0x36, 0x4a, 0x08, 0x9f, 0x5a, 0x0d, 0xf5, 0x9f, 0x8e, 0x9c, 0xea,
	0xe4, 0xec, 0x94, 0xf0, 0xa9, 0x9c, 0xca, 0x67, 0x32, 0x2a, 0x04, 0x7a, 0xac, 0x0a, 0x6a, 0x88,
	0xe5, 0x20, 0x5a, 0xe3, 0x95, 0x0a, 0xf8, 0x5d, 0x19, 0xa9, 0xc5, 0x34, 0xbb, 0x66, 0x7f, 0xff,
	0x08, 0x3a, 0x3e, 0x73, 0x4e, 0x82, 0x90, 0xb2, 0x8c, 0x71, 0x3a, 0x3b, 0xcf, 0x12, 0x7a, 0x6b,
	0x2e, 0x63, 0x6d, 0x7e, 0xae, 0x17, 0x77, 0x6b, 0x2e, 0x61, 0x69, 0x2e, 0x43, 0xaf, 0x54, 0x74,
	0x66, 0xe0, 0xe0, 0xde, 0x06, 0xe0, 0x4b, 0x50, 0xbf, 0xa4, 0x59, 0x79, 0x04, 0x4f, 0x73, 0x81,
	0x24, 0x2c, 0x04, 0x6a, 0x2b, 0xab, 0x4b, 0x9a, 0x61, 0x4f, 0x32, 0xd0, 0x01, 0xcd, 0x05, 0x09,
	0xe7, 0xd4, 0xaa, 0x29, 0xa5, 0x95, 0x0b, 0xa4, 0x89, 0x42, 0xa0, 0x3d, 0xa5, 0x55, 0x08, 0x7b,
	0x9a, 0x1d, 0xd4, 0xde, 0x98, 0xc3, 0x77, 0xab, 0x6b, 0xdb, 0xd8, 0x5c, 0xdb, 0xc6, 0x6a, 0x6b,
	0x9b, 0x9b, 0xad, 0x6d, 0x5e, 0xed, 0x6c, 0xe3, 0xf7, 0xce, 0x36, 0x37, 0x3b, 0xdb, 0xf8, 0xb3,
	0xb3, 0x8d, 0x6f, 0xcf, 0x27, 0x01, 0x9f, 0xce, 0x2f, 0x9c, 0x71, 0x3c, 0x73, 0xe9, 0x32, 0x99,
	0xc4, 0x6e, 0x44, 0xb9, 0xab, 0x8f, 0xe8, 0xa2, 0xa5, 0x6e, 0xf8, 0xf8, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x7f, 0x13, 0x91, 0x0e, 0x03, 0x00, 0x00,
}

func (m *VersioningConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersioningConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersioningConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FSType != 0 {
		i = encodeVarintVersioningconfiguration(dAtA, i, uint64(m.FSType))
		i--
		dAtA[i] = 0x28
	}
	if len(m.FSPath) > 0 {
		i -= len(m.FSPath)
		copy(dAtA[i:], m.FSPath)
		i = encodeVarintVersioningconfiguration(dAtA, i, uint64(len(m.FSPath)))
		i--
		dAtA[i] = 0x22
	}
	if m.CleanupIntervalS != 0 {
		i = encodeVarintVersioningconfiguration(dAtA, i, uint64(m.CleanupIntervalS))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Params) > 0 {
		for k := range m.Params {
			v := m.Params[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintVersioningconfiguration(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintVersioningconfiguration(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintVersioningconfiguration(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintVersioningconfiguration(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVersioningconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovVersioningconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VersioningConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovVersioningconfiguration(uint64(l))
	}
	if len(m.Params) > 0 {
		for k, v := range m.Params {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovVersioningconfiguration(uint64(len(k))) + 1 + len(v) + sovVersioningconfiguration(uint64(len(v)))
			n += mapEntrySize + 1 + sovVersioningconfiguration(uint64(mapEntrySize))
		}
	}
	if m.CleanupIntervalS != 0 {
		n += 1 + sovVersioningconfiguration(uint64(m.CleanupIntervalS))
	}
	l = len(m.FSPath)
	if l > 0 {
		n += 1 + l + sovVersioningconfiguration(uint64(l))
	}
	if m.FSType != 0 {
		n += 1 + sovVersioningconfiguration(uint64(m.FSType))
	}
	return n
}

func sovVersioningconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVersioningconfiguration(x uint64) (n int) {
	return sovVersioningconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VersioningConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVersioningconfiguration
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
			return fmt.Errorf("proto: VersioningConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersioningConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersioningconfiguration
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
				return ErrInvalidLengthVersioningconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersioningconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersioningconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVersioningconfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVersioningconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVersioningconfiguration
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVersioningconfiguration
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthVersioningconfiguration
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthVersioningconfiguration
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVersioningconfiguration
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthVersioningconfiguration
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthVersioningconfiguration
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipVersioningconfiguration(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthVersioningconfiguration
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Params[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CleanupIntervalS", wireType)
			}
			m.CleanupIntervalS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersioningconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CleanupIntervalS |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FSPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersioningconfiguration
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
				return ErrInvalidLengthVersioningconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVersioningconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FSPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FSType", wireType)
			}
			m.FSType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersioningconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FSType |= fs.FilesystemType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVersioningconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVersioningconfiguration
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
func skipVersioningconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVersioningconfiguration
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
					return 0, ErrIntOverflowVersioningconfiguration
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
					return 0, ErrIntOverflowVersioningconfiguration
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
				return 0, ErrInvalidLengthVersioningconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVersioningconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVersioningconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVersioningconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersioningconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVersioningconfiguration = fmt.Errorf("proto: unexpected end of group")
)
