// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/subdomain_registration_config.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type SubdomainRegistrationFees struct {
	FeeByLength map[uint32]*Fee `protobuf:"bytes,1,rep,name=feeByLength,proto3" json:"feeByLength,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FeeByName   map[string]*Fee `protobuf:"bytes,2,rep,name=feeByName,proto3" json:"feeByName,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DefaultFee  *types.Coin     `protobuf:"bytes,3,opt,name=defaultFee,proto3" json:"defaultFee,omitempty"`
}

func (m *SubdomainRegistrationFees) Reset()         { *m = SubdomainRegistrationFees{} }
func (m *SubdomainRegistrationFees) String() string { return proto.CompactTextString(m) }
func (*SubdomainRegistrationFees) ProtoMessage()    {}
func (*SubdomainRegistrationFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_93c8ca252d69a2bc, []int{0}
}
func (m *SubdomainRegistrationFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubdomainRegistrationFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubdomainRegistrationFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubdomainRegistrationFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubdomainRegistrationFees.Merge(m, src)
}
func (m *SubdomainRegistrationFees) XXX_Size() int {
	return m.Size()
}
func (m *SubdomainRegistrationFees) XXX_DiscardUnknown() {
	xxx_messageInfo_SubdomainRegistrationFees.DiscardUnknown(m)
}

var xxx_messageInfo_SubdomainRegistrationFees proto.InternalMessageInfo

func (m *SubdomainRegistrationFees) GetFeeByLength() map[uint32]*Fee {
	if m != nil {
		return m.FeeByLength
	}
	return nil
}

func (m *SubdomainRegistrationFees) GetFeeByName() map[string]*Fee {
	if m != nil {
		return m.FeeByName
	}
	return nil
}

func (m *SubdomainRegistrationFees) GetDefaultFee() *types.Coin {
	if m != nil {
		return m.DefaultFee
	}
	return nil
}

type Fee struct {
	IsRegistrable bool        `protobuf:"varint,1,opt,name=isRegistrable,proto3" json:"isRegistrable,omitempty"`
	Fee           *types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_93c8ca252d69a2bc, []int{1}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetIsRegistrable() bool {
	if m != nil {
		return m.IsRegistrable
	}
	return false
}

func (m *Fee) GetFee() *types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

type SubdomainRegistrationConfig struct {
	MaxSubdomainRegistrations uint64                     `protobuf:"varint,1,opt,name=maxSubdomainRegistrations,proto3" json:"maxSubdomainRegistrations,omitempty"`
	SubdomainRegistrationFees *SubdomainRegistrationFees `protobuf:"bytes,2,opt,name=subdomainRegistrationFees,proto3" json:"subdomainRegistrationFees,omitempty"`
}

func (m *SubdomainRegistrationConfig) Reset()         { *m = SubdomainRegistrationConfig{} }
func (m *SubdomainRegistrationConfig) String() string { return proto.CompactTextString(m) }
func (*SubdomainRegistrationConfig) ProtoMessage()    {}
func (*SubdomainRegistrationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_93c8ca252d69a2bc, []int{2}
}
func (m *SubdomainRegistrationConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubdomainRegistrationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubdomainRegistrationConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubdomainRegistrationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubdomainRegistrationConfig.Merge(m, src)
}
func (m *SubdomainRegistrationConfig) XXX_Size() int {
	return m.Size()
}
func (m *SubdomainRegistrationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SubdomainRegistrationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SubdomainRegistrationConfig proto.InternalMessageInfo

func (m *SubdomainRegistrationConfig) GetMaxSubdomainRegistrations() uint64 {
	if m != nil {
		return m.MaxSubdomainRegistrations
	}
	return 0
}

func (m *SubdomainRegistrationConfig) GetSubdomainRegistrationFees() *SubdomainRegistrationFees {
	if m != nil {
		return m.SubdomainRegistrationFees
	}
	return nil
}

func init() {
	proto.RegisterType((*SubdomainRegistrationFees)(nil), "mycel.registry.SubdomainRegistrationFees")
	proto.RegisterMapType((map[uint32]*Fee)(nil), "mycel.registry.SubdomainRegistrationFees.FeeByLengthEntry")
	proto.RegisterMapType((map[string]*Fee)(nil), "mycel.registry.SubdomainRegistrationFees.FeeByNameEntry")
	proto.RegisterType((*Fee)(nil), "mycel.registry.Fee")
	proto.RegisterType((*SubdomainRegistrationConfig)(nil), "mycel.registry.SubdomainRegistrationConfig")
}

func init() {
	proto.RegisterFile("mycel/registry/subdomain_registration_config.proto", fileDescriptor_93c8ca252d69a2bc)
}

var fileDescriptor_93c8ca252d69a2bc = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x86, 0x77, 0x1a, 0x15, 0xfb, 0x2d, 0x2d, 0x65, 0xbc, 0x6c, 0x56, 0x08, 0xcb, 0xe2, 0x61,
	0x8b, 0x38, 0xa1, 0xeb, 0xa5, 0x16, 0x4f, 0x2d, 0x06, 0x0f, 0x22, 0x38, 0x05, 0x11, 0x11, 0xca,
	0x24, 0xfd, 0x92, 0x0e, 0x26, 0x99, 0x92, 0x99, 0x94, 0xe6, 0x5f, 0xf8, 0x63, 0xfc, 0x09, 0x1e,
	0x3c, 0xf6, 0xe8, 0x51, 0x76, 0xff, 0x88, 0x64, 0xb2, 0x61, 0xb3, 0xa5, 0x11, 0xed, 0x2d, 0x33,
	0xbc, 0xef, 0xf3, 0xbd, 0xbc, 0x99, 0x0f, 0xe6, 0x59, 0x15, 0x61, 0xea, 0x17, 0x98, 0x48, 0x6d,
	0x8a, 0xca, 0xd7, 0x65, 0x78, 0xae, 0x32, 0x21, 0xf3, 0xb3, 0xd5, 0x95, 0x30, 0x52, 0xe5, 0x67,
	0x91, 0xca, 0x63, 0x99, 0xb0, 0xcb, 0x42, 0x19, 0x45, 0x77, 0xad, 0x87, 0xb5, 0x9e, 0xb1, 0x17,
	0x29, 0x9d, 0x29, 0xed, 0x87, 0x42, 0xa3, 0x7f, 0x75, 0x10, 0xa2, 0x11, 0x07, 0x7e, 0xa4, 0x64,
	0xde, 0xe8, 0xa7, 0xdf, 0x1d, 0x70, 0x4f, 0x5b, 0x2e, 0xef, 0x60, 0x03, 0x44, 0x4d, 0xbf, 0xc0,
	0x30, 0x46, 0x3c, 0xae, 0xde, 0x61, 0x9e, 0x98, 0x8b, 0x11, 0x99, 0x38, 0xb3, 0xe1, 0xfc, 0x88,
	0x6d, 0xce, 0x60, 0xbd, 0x7e, 0x16, 0xac, 0xcd, 0x6f, 0x72, 0x53, 0x54, 0xbc, 0x8b, 0xa3, 0x1f,
	0x61, 0xdb, 0x1e, 0xdf, 0x8b, 0x0c, 0x47, 0x5b, 0x96, 0x7d, 0xf8, 0x9f, 0xec, 0xda, 0xda, 0x90,
	0xd7, 0x28, 0xfa, 0x0a, 0xe0, 0x1c, 0x63, 0x51, 0xa6, 0x26, 0x40, 0x1c, 0x39, 0x13, 0x32, 0x1b,
	0xce, 0x5d, 0xd6, 0x14, 0xc1, 0xea, 0x22, 0xd8, 0xaa, 0x08, 0x76, 0xa2, 0x64, 0xce, 0x3b, 0xe2,
	0xf1, 0x29, 0xec, 0xdd, 0xce, 0x4c, 0xf7, 0xc0, 0xf9, 0x8a, 0xd5, 0x88, 0x4c, 0xc8, 0x6c, 0x87,
	0xd7, 0x9f, 0x74, 0x1f, 0x1e, 0x5e, 0x89, 0xb4, 0xac, 0x43, 0xd7, 0xec, 0x27, 0xb7, 0x43, 0x07,
	0x88, 0xbc, 0x51, 0x1c, 0x6d, 0x1d, 0x92, 0xf1, 0x07, 0xd8, 0xdd, 0x0c, 0xdb, 0x45, 0x6e, 0xdf,
	0x07, 0x39, 0xfd, 0x04, 0x4e, 0x80, 0x48, 0x9f, 0xc1, 0x8e, 0xd4, 0x6d, 0x33, 0x61, 0x8a, 0x96,
	0xf8, 0x98, 0x6f, 0x5e, 0xd2, 0xe7, 0xe0, 0xc4, 0xd8, 0x92, 0xff, 0x52, 0x44, 0xad, 0x9a, 0xfe,
	0x20, 0xf0, 0xf4, 0xce, 0xd2, 0x4f, 0xec, 0x33, 0xa3, 0xaf, 0xc1, 0xcd, 0xc4, 0xf5, 0x9d, 0x0a,
	0x6d, 0xc7, 0x3f, 0xe0, 0xfd, 0x02, 0x9a, 0x80, 0xab, 0xfb, 0xfe, 0xe8, 0x2a, 0xe0, 0xfe, 0x3f,
	0x3f, 0x01, 0xde, 0xcf, 0x3a, 0x7e, 0xfb, 0x73, 0xe1, 0x91, 0x9b, 0x85, 0x47, 0x7e, 0x2f, 0x3c,
	0xf2, 0x6d, 0xe9, 0x0d, 0x6e, 0x96, 0xde, 0xe0, 0xd7, 0xd2, 0x1b, 0x7c, 0x66, 0x89, 0x34, 0x17,
	0x65, 0xc8, 0x22, 0x95, 0xf9, 0x76, 0xd2, 0x8b, 0x06, 0xd1, 0x1c, 0xfc, 0xeb, 0xf5, 0xbe, 0x99,
	0xea, 0x12, 0x75, 0xf8, 0xc8, 0x2e, 0xca, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xae,
	0x3a, 0x8a, 0x8e, 0x03, 0x00, 0x00,
}

func (m *SubdomainRegistrationFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubdomainRegistrationFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubdomainRegistrationFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultFee != nil {
		{
			size, err := m.DefaultFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FeeByName) > 0 {
		for k := range m.FeeByName {
			v := m.FeeByName[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.FeeByLength) > 0 {
		for k := range m.FeeByLength {
			v := m.FeeByLength[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.IsRegistrable {
		i--
		if m.IsRegistrable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubdomainRegistrationConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubdomainRegistrationConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubdomainRegistrationConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubdomainRegistrationFees != nil {
		{
			size, err := m.SubdomainRegistrationFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxSubdomainRegistrations != 0 {
		i = encodeVarintSubdomainRegistrationConfig(dAtA, i, uint64(m.MaxSubdomainRegistrations))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubdomainRegistrationConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubdomainRegistrationConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SubdomainRegistrationFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeByLength) > 0 {
		for k, v := range m.FeeByLength {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovSubdomainRegistrationConfig(uint64(l))
			}
			mapEntrySize := 1 + sovSubdomainRegistrationConfig(uint64(k)) + l
			n += mapEntrySize + 1 + sovSubdomainRegistrationConfig(uint64(mapEntrySize))
		}
	}
	if len(m.FeeByName) > 0 {
		for k, v := range m.FeeByName {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovSubdomainRegistrationConfig(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovSubdomainRegistrationConfig(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovSubdomainRegistrationConfig(uint64(mapEntrySize))
		}
	}
	if m.DefaultFee != nil {
		l = m.DefaultFee.Size()
		n += 1 + l + sovSubdomainRegistrationConfig(uint64(l))
	}
	return n
}

func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsRegistrable {
		n += 2
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovSubdomainRegistrationConfig(uint64(l))
	}
	return n
}

func (m *SubdomainRegistrationConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxSubdomainRegistrations != 0 {
		n += 1 + sovSubdomainRegistrationConfig(uint64(m.MaxSubdomainRegistrations))
	}
	if m.SubdomainRegistrationFees != nil {
		l = m.SubdomainRegistrationFees.Size()
		n += 1 + l + sovSubdomainRegistrationConfig(uint64(l))
	}
	return n
}

func sovSubdomainRegistrationConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubdomainRegistrationConfig(x uint64) (n int) {
	return sovSubdomainRegistrationConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SubdomainRegistrationFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainRegistrationConfig
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
			return fmt.Errorf("proto: SubdomainRegistrationFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubdomainRegistrationFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeByLength", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeByLength == nil {
				m.FeeByLength = make(map[uint32]*Fee)
			}
			var mapkey uint32
			var mapvalue *Fee
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSubdomainRegistrationConfig
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
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSubdomainRegistrationConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSubdomainRegistrationConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Fee{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSubdomainRegistrationConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.FeeByLength[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeByName", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeByName == nil {
				m.FeeByName = make(map[string]*Fee)
			}
			var mapkey string
			var mapvalue *Fee
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSubdomainRegistrationConfig
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
							return ErrIntOverflowSubdomainRegistrationConfig
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
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSubdomainRegistrationConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Fee{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSubdomainRegistrationConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthSubdomainRegistrationConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.FeeByName[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultFee == nil {
				m.DefaultFee = &types.Coin{}
			}
			if err := m.DefaultFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainRegistrationConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
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
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainRegistrationConfig
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRegistrable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
			m.IsRegistrable = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &types.Coin{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainRegistrationConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
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
func (m *SubdomainRegistrationConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainRegistrationConfig
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
			return fmt.Errorf("proto: SubdomainRegistrationConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubdomainRegistrationConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubdomainRegistrations", wireType)
			}
			m.MaxSubdomainRegistrations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubdomainRegistrations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainRegistrationFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainRegistrationConfig
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
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubdomainRegistrationFees == nil {
				m.SubdomainRegistrationFees = &SubdomainRegistrationFees{}
			}
			if err := m.SubdomainRegistrationFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainRegistrationConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainRegistrationConfig
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
func skipSubdomainRegistrationConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubdomainRegistrationConfig
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
					return 0, ErrIntOverflowSubdomainRegistrationConfig
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
					return 0, ErrIntOverflowSubdomainRegistrationConfig
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
				return 0, ErrInvalidLengthSubdomainRegistrationConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubdomainRegistrationConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubdomainRegistrationConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubdomainRegistrationConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubdomainRegistrationConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubdomainRegistrationConfig = fmt.Errorf("proto: unexpected end of group")
)
