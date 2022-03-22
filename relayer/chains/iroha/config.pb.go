// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/chains/iroha/config.proto

package iroha

import (
	fmt "fmt"
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

type ChainConfig struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	RpcAddr string `protobuf:"bytes,2,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	// use for relayer
	AccountId                 string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	IbcHostAddress            string `protobuf:"bytes,5,opt,name=ibc_host_address,json=ibcHostAddress,proto3" json:"ibc_host_address,omitempty"`
	IbcHandlerAddress         string `protobuf:"bytes,6,opt,name=ibc_handler_address,json=ibcHandlerAddress,proto3" json:"ibc_handler_address,omitempty"`
	IrohaIcs20BankAddress     string `protobuf:"bytes,7,opt,name=iroha_ics20_bank_address,json=irohaIcs20BankAddress,proto3" json:"iroha_ics20_bank_address,omitempty"`
	IrohaIcs20TransferAddress string `protobuf:"bytes,8,opt,name=iroha_ics20_transfer_address,json=irohaIcs20TransferAddress,proto3" json:"iroha_ics20_transfer_address,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd3ec72750e2d68, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.iroha.config.ChainConfig")
}

func init() { proto.RegisterFile("relayer/chains/iroha/config.proto", fileDescriptor_3cd3ec72750e2d68) }

var fileDescriptor_3cd3ec72750e2d68 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x40, 0x93, 0x22, 0xfa, 0x30, 0x12, 0x82, 0x00, 0x52, 0xca, 0xc3, 0x02, 0xa6, 0x2e, 0x4d,
	0x10, 0x20, 0x31, 0xa2, 0xb6, 0x0b, 0x5d, 0x2b, 0x26, 0x96, 0xc8, 0xb1, 0xdd, 0xc4, 0x6a, 0xb0,
	0x23, 0x3b, 0x1d, 0xfa, 0x17, 0x7c, 0x0d, 0xdf, 0xd0, 0xb1, 0x23, 0x23, 0x34, 0x3f, 0x82, 0x72,
	0x9d, 0x96, 0x0e, 0x6c, 0xc9, 0x3d, 0xe7, 0xd8, 0xb2, 0x2e, 0xba, 0xd1, 0x3c, 0x23, 0x0b, 0xae,
	0x43, 0x9a, 0x12, 0x21, 0x4d, 0x28, 0xb4, 0x4a, 0x49, 0x48, 0x95, 0x9c, 0x8a, 0x24, 0xc8, 0xb5,
	0x2a, 0x94, 0x77, 0x51, 0x2b, 0x81, 0x55, 0x02, 0x50, 0x02, 0xab, 0x9c, 0x9f, 0x26, 0x2a, 0x51,
	0xe0, 0x85, 0xd5, 0x97, 0x4d, 0x6e, 0x3f, 0x1b, 0xe8, 0x60, 0x54, 0xd9, 0x23, 0xb0, 0xbc, 0x2e,
	0x6a, 0x43, 0x1c, 0x09, 0xe6, 0xbb, 0xd7, 0x6e, 0xaf, 0x33, 0x69, 0xc1, 0xff, 0x98, 0x55, 0x48,
	0xe7, 0x34, 0x22, 0x8c, 0x69, 0xbf, 0x61, 0x91, 0xce, 0xe9, 0x80, 0x31, 0xed, 0x5d, 0x21, 0x44,
	0x28, 0x55, 0x73, 0x59, 0x54, 0xdd, 0x1e, 0xc0, 0x4e, 0x3d, 0x19, 0x33, 0xaf, 0x87, 0x8e, 0x44,
	0x4c, 0xa3, 0x54, 0x99, 0x02, 0x72, 0x6e, 0x8c, 0xbf, 0x0f, 0xd2, 0xa1, 0x88, 0xe9, 0x8b, 0x32,
	0xc5, 0xc0, 0x4e, 0xbd, 0x00, 0x9d, 0x80, 0x49, 0x24, 0xcb, 0xb8, 0xde, 0xca, 0x4d, 0x90, 0x8f,
	0x2b, 0xd9, 0x92, 0x8d, 0xff, 0x84, 0x7c, 0x78, 0x64, 0x24, 0xa8, 0xb9, 0xbf, 0x8b, 0x62, 0x22,
	0x67, 0xdb, 0xa8, 0x05, 0xd1, 0x19, 0xf0, 0x71, 0x85, 0x87, 0x44, 0xce, 0x36, 0xe1, 0x33, 0xba,
	0xdc, 0x0d, 0x0b, 0x4d, 0xa4, 0x99, 0xee, 0xdc, 0xd8, 0x86, 0xb8, 0xfb, 0x17, 0xbf, 0xd6, 0x46,
	0x7d, 0xc0, 0x70, 0xb2, 0xfc, 0xc1, 0xce, 0x72, 0x8d, 0xdd, 0xd5, 0x1a, 0xbb, 0xdf, 0x6b, 0xec,
	0x7e, 0x94, 0xd8, 0x59, 0x95, 0xd8, 0xf9, 0x2a, 0xb1, 0xf3, 0xf6, 0x98, 0x88, 0x22, 0x9d, 0xc7,
	0x01, 0x55, 0xef, 0x61, 0xba, 0xc8, 0xb9, 0xce, 0x38, 0x4b, 0xb8, 0xee, 0x67, 0x24, 0x36, 0xe1,
	0x62, 0x2e, 0xfa, 0xff, 0x2d, 0x33, 0x6e, 0xc2, 0x4e, 0x1e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0x5e, 0xeb, 0x64, 0xeb, 0x01, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IrohaIcs20TransferAddress) > 0 {
		i -= len(m.IrohaIcs20TransferAddress)
		copy(dAtA[i:], m.IrohaIcs20TransferAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IrohaIcs20TransferAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IrohaIcs20BankAddress) > 0 {
		i -= len(m.IrohaIcs20BankAddress)
		copy(dAtA[i:], m.IrohaIcs20BankAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IrohaIcs20BankAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.IbcHandlerAddress) > 0 {
		i -= len(m.IbcHandlerAddress)
		copy(dAtA[i:], m.IbcHandlerAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHandlerAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.IbcHostAddress) > 0 {
		i -= len(m.IbcHostAddress)
		copy(dAtA[i:], m.IbcHostAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHostAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AccountId) > 0 {
		i -= len(m.AccountId)
		copy(dAtA[i:], m.AccountId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.AccountId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RpcAddr) > 0 {
		i -= len(m.RpcAddr)
		copy(dAtA[i:], m.RpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RpcAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.RpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.AccountId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHostAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHandlerAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IrohaIcs20BankAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IrohaIcs20TransferAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHostAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHandlerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHandlerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IrohaIcs20BankAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IrohaIcs20BankAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IrohaIcs20TransferAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IrohaIcs20TransferAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
