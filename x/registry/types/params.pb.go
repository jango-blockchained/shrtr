// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/params.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	StakingInflationRatio      float32 `protobuf:"fixed32,1,opt,name=stakingInflationRatio,proto3" json:"stakingInflationRatio,omitempty"`
	TopLevelDomainBaseFeeInUsd int64   `protobuf:"varint,2,opt,name=topLevelDomainBaseFeeInUsd,proto3" json:"topLevelDomainBaseFeeInUsd,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_094deb43283078b0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetStakingInflationRatio() float32 {
	if m != nil {
		return m.StakingInflationRatio
	}
	return 0
}

func (m *Params) GetTopLevelDomainBaseFeeInUsd() int64 {
	if m != nil {
		return m.TopLevelDomainBaseFeeInUsd
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mycel.registry.Params")
}

func init() { proto.RegisterFile("mycel/registry/params.proto", fileDescriptor_094deb43283078b0) }

var fileDescriptor_094deb43283078b0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xad, 0x4c, 0x4e,
	0xcd, 0xd1, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x4b, 0xea, 0xc1, 0x24, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x52, 0xfa, 0x20, 0x16, 0x44, 0x95, 0x52, 0x0b, 0x23, 0x17, 0x5b,
	0x00, 0x58, 0x9b, 0x90, 0x09, 0x97, 0x68, 0x71, 0x49, 0x62, 0x76, 0x66, 0x5e, 0xba, 0x67, 0x5e,
	0x5a, 0x4e, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x10, 0x88, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0a,
	0xc2, 0x2e, 0x29, 0x64, 0xc7, 0x25, 0x55, 0x92, 0x5f, 0xe0, 0x93, 0x5a, 0x96, 0x9a, 0xe3, 0x92,
	0x9f, 0x9b, 0x98, 0x99, 0xe7, 0x94, 0x58, 0x9c, 0xea, 0x96, 0x9a, 0xea, 0x99, 0x17, 0x5a, 0x9c,
	0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x47, 0x85, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x0c,
	0x4e, 0x1e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x97, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf6, 0x91, 0x6e, 0x0a, 0xd8, 0x10, 0x08,
	0x47, 0xbf, 0x02, 0xe1, 0xfb, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xbf, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x2e, 0x01, 0x96, 0x1c, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TopLevelDomainBaseFeeInUsd != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TopLevelDomainBaseFeeInUsd))
		i--
		dAtA[i] = 0x10
	}
	if m.StakingInflationRatio != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.StakingInflationRatio))))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StakingInflationRatio != 0 {
		n += 5
	}
	if m.TopLevelDomainBaseFeeInUsd != 0 {
		n += 1 + sovParams(uint64(m.TopLevelDomainBaseFeeInUsd))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingInflationRatio", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.StakingInflationRatio = float32(math.Float32frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopLevelDomainBaseFeeInUsd", wireType)
			}
			m.TopLevelDomainBaseFeeInUsd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopLevelDomainBaseFeeInUsd |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
