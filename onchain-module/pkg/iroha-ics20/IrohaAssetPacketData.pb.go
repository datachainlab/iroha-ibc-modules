// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: IrohaAssetPacketData.proto

package iroha_ics20

import (
	fmt "fmt"
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

type IrohaAssetPacketData struct {
	SrcAccountId  string `protobuf:"bytes,1,opt,name=src_account_id,json=srcAccountId,proto3" json:"src_account_id,omitempty"`
	DestAccountId string `protobuf:"bytes,2,opt,name=dest_account_id,json=destAccountId,proto3" json:"dest_account_id,omitempty"`
	AssetId       string `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Amount        string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *IrohaAssetPacketData) Reset()         { *m = IrohaAssetPacketData{} }
func (m *IrohaAssetPacketData) String() string { return proto.CompactTextString(m) }
func (*IrohaAssetPacketData) ProtoMessage()    {}
func (*IrohaAssetPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_487496a1402bab07, []int{0}
}
func (m *IrohaAssetPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IrohaAssetPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IrohaAssetPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IrohaAssetPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IrohaAssetPacketData.Merge(m, src)
}
func (m *IrohaAssetPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IrohaAssetPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IrohaAssetPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IrohaAssetPacketData proto.InternalMessageInfo

func (m *IrohaAssetPacketData) GetSrcAccountId() string {
	if m != nil {
		return m.SrcAccountId
	}
	return ""
}

func (m *IrohaAssetPacketData) GetDestAccountId() string {
	if m != nil {
		return m.DestAccountId
	}
	return ""
}

func (m *IrohaAssetPacketData) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *IrohaAssetPacketData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IrohaAssetPacketData) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*IrohaAssetPacketData)(nil), "IrohaAssetPacketData")
}

func init() { proto.RegisterFile("IrohaAssetPacketData.proto", fileDescriptor_487496a1402bab07) }

var fileDescriptor_487496a1402bab07 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf2, 0x2c, 0xca, 0xcf,
	0x48, 0x74, 0x2c, 0x2e, 0x4e, 0x2d, 0x09, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x71, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xda, 0xc9, 0xc8, 0x25, 0x82, 0x4d, 0x5a, 0x48, 0x85,
	0x8b, 0xaf, 0xb8, 0x28, 0x39, 0x3e, 0x31, 0x39, 0x39, 0xbf, 0x34, 0xaf, 0x24, 0x3e, 0x33, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0xa7, 0xb8, 0x28, 0xd9, 0x11, 0x22, 0xe8, 0x99, 0x22,
	0xa4, 0xc6, 0xc5, 0x9f, 0x92, 0x5a, 0x5c, 0x82, 0xac, 0x8c, 0x09, 0xac, 0x8c, 0x17, 0x24, 0x8c,
	0x50, 0x27, 0xc9, 0xc5, 0x91, 0x08, 0xb2, 0x00, 0xa4, 0x80, 0x19, 0xac, 0x80, 0x1d, 0xcc, 0xf7,
	0x4c, 0x11, 0x52, 0xe0, 0xe2, 0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf,
	0x93, 0x60, 0x01, 0xcb, 0x22, 0x0b, 0x09, 0x89, 0x71, 0xb1, 0x25, 0xe6, 0x82, 0x0c, 0x92, 0x60,
	0x05, 0x4b, 0x42, 0x79, 0x4e, 0x49, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0xe5, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x92, 0x58, 0x92,
	0x98, 0x9c, 0x91, 0x98, 0x99, 0x97, 0x93, 0x98, 0xa4, 0x9f, 0x09, 0xf2, 0xaa, 0x6e, 0x66, 0x52,
	0xb2, 0x6e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0xb1, 0x7e, 0x7e, 0x1e, 0x58, 0x12, 0xca, 0xd7,
	0x2f, 0xc8, 0x4e, 0x87, 0x29, 0x4a, 0x2e, 0x36, 0x32, 0x48, 0x62, 0x03, 0x07, 0x93, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0x38, 0x1a, 0x07, 0x44, 0x01, 0x00, 0x00,
}

func (m *IrohaAssetPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IrohaAssetPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IrohaAssetPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintIrohaAssetPacketData(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIrohaAssetPacketData(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintIrohaAssetPacketData(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestAccountId) > 0 {
		i -= len(m.DestAccountId)
		copy(dAtA[i:], m.DestAccountId)
		i = encodeVarintIrohaAssetPacketData(dAtA, i, uint64(len(m.DestAccountId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SrcAccountId) > 0 {
		i -= len(m.SrcAccountId)
		copy(dAtA[i:], m.SrcAccountId)
		i = encodeVarintIrohaAssetPacketData(dAtA, i, uint64(len(m.SrcAccountId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIrohaAssetPacketData(dAtA []byte, offset int, v uint64) int {
	offset -= sovIrohaAssetPacketData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IrohaAssetPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcAccountId)
	if l > 0 {
		n += 1 + l + sovIrohaAssetPacketData(uint64(l))
	}
	l = len(m.DestAccountId)
	if l > 0 {
		n += 1 + l + sovIrohaAssetPacketData(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovIrohaAssetPacketData(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIrohaAssetPacketData(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovIrohaAssetPacketData(uint64(l))
	}
	return n
}

func sovIrohaAssetPacketData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIrohaAssetPacketData(x uint64) (n int) {
	return sovIrohaAssetPacketData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IrohaAssetPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIrohaAssetPacketData
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
			return fmt.Errorf("proto: IrohaAssetPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IrohaAssetPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcAccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrohaAssetPacketData
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
				return ErrInvalidLengthIrohaAssetPacketData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcAccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrohaAssetPacketData
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
				return ErrInvalidLengthIrohaAssetPacketData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrohaAssetPacketData
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
				return ErrInvalidLengthIrohaAssetPacketData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrohaAssetPacketData
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
				return ErrInvalidLengthIrohaAssetPacketData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrohaAssetPacketData
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
				return ErrInvalidLengthIrohaAssetPacketData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIrohaAssetPacketData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIrohaAssetPacketData
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
func skipIrohaAssetPacketData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIrohaAssetPacketData
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
					return 0, ErrIntOverflowIrohaAssetPacketData
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
					return 0, ErrIntOverflowIrohaAssetPacketData
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
				return 0, ErrInvalidLengthIrohaAssetPacketData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIrohaAssetPacketData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIrohaAssetPacketData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIrohaAssetPacketData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIrohaAssetPacketData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIrohaAssetPacketData = fmt.Errorf("proto: unexpected end of group")
)
