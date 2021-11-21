//*
// Copyright Soramitsu Co., Ltd. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: block.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload    *BlockV1_Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signatures []*Signature     `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *BlockV1) Reset() {
	*x = BlockV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockV1) ProtoMessage() {}

func (x *BlockV1) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockV1.ProtoReflect.Descriptor instead.
func (*BlockV1) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockV1) GetPayload() *BlockV1_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *BlockV1) GetSignatures() []*Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BlockVersion:
	//	*Block_BlockV1
	BlockVersion isBlock_BlockVersion `protobuf_oneof:"block_version"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (m *Block) GetBlockVersion() isBlock_BlockVersion {
	if m != nil {
		return m.BlockVersion
	}
	return nil
}

func (x *Block) GetBlockV1() *BlockV1 {
	if x, ok := x.GetBlockVersion().(*Block_BlockV1); ok {
		return x.BlockV1
	}
	return nil
}

type isBlock_BlockVersion interface {
	isBlock_BlockVersion()
}

type Block_BlockV1 struct {
	BlockV1 *BlockV1 `protobuf:"bytes,1,opt,name=block_v1,json=blockV1,proto3,oneof"`
}

func (*Block_BlockV1) isBlock_BlockVersion() {}

// everything that should be signed:
type BlockV1_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TxNumber     uint32         `protobuf:"varint,2,opt,name=tx_number,json=txNumber,proto3" json:"tx_number,omitempty"` ///< The number of accepted transactions inside.
	///< Maximum 16384 or 2^14.
	Height        uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`                                     ///< The current block number in a ledger.
	PrevBlockHash string `protobuf:"bytes,4,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"` ///< Previous block hash.
	CreatedTime   uint64 `protobuf:"varint,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	/// Hashes of the transactions that did not pass stateful validation.
	/// Needed here to be able to guarantee the client that this transaction
	/// was not and will never be executed.
	RejectedTransactionsHashes []string `protobuf:"bytes,6,rep,name=rejected_transactions_hashes,json=rejectedTransactionsHashes,proto3" json:"rejected_transactions_hashes,omitempty"`
}

func (x *BlockV1_Payload) Reset() {
	*x = BlockV1_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockV1_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockV1_Payload) ProtoMessage() {}

func (x *BlockV1_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockV1_Payload.ProtoReflect.Descriptor instead.
func (*BlockV1_Payload) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BlockV1_Payload) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockV1_Payload) GetTxNumber() uint32 {
	if x != nil {
		return x.TxNumber
	}
	return 0
}

func (x *BlockV1_Payload) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockV1_Payload) GetPrevBlockHash() string {
	if x != nil {
		return x.PrevBlockHash
	}
	return ""
}

func (x *BlockV1_Payload) GetCreatedTime() uint64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *BlockV1_Payload) GetRejectedTransactionsHashes() []string {
	if x != nil {
		return x.RejectedTransactionsHashes
	}
	return nil
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69,
	0x72, 0x6f, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x0f, 0x70,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x90, 0x03, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x12, 0x3a,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x69, 0x72, 0x6f, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x69, 0x72, 0x6f, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x8c, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x72, 0x6f, 0x68, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x1c, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1a, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x35, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x69, 0x72, 0x6f, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x56, 0x31, 0x42, 0x0f, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x1a, 0x5a, 0x18, 0x69, 0x72, 0x6f, 0x68, 0x61, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_block_proto_goTypes = []interface{}{
	(*BlockV1)(nil),         // 0: iroha.protocol.Block_v1
	(*Block)(nil),           // 1: iroha.protocol.Block
	(*BlockV1_Payload)(nil), // 2: iroha.protocol.Block_v1.Payload
	(*Signature)(nil),       // 3: iroha.protocol.Signature
	(*Transaction)(nil),     // 4: iroha.protocol.Transaction
}
var file_block_proto_depIdxs = []int32{
	2, // 0: iroha.protocol.Block_v1.payload:type_name -> iroha.protocol.Block_v1.Payload
	3, // 1: iroha.protocol.Block_v1.signatures:type_name -> iroha.protocol.Signature
	0, // 2: iroha.protocol.Block.block_v1:type_name -> iroha.protocol.Block_v1
	4, // 3: iroha.protocol.Block_v1.Payload.transactions:type_name -> iroha.protocol.Transaction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_primitive_proto_init()
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockV1_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_block_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Block_BlockV1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}