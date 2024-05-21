// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
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

// GenesisState defines the finality module's genesis state.
type GenesisState struct {
	// params the current params of the state.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// indexed_blocks all the btc blocks and if their status are finalized.
	IndexedBlocks []*IndexedBlock `protobuf:"bytes,2,rep,name=indexed_blocks,json=indexedBlocks,proto3" json:"indexed_blocks,omitempty"`
	// evidences all the evidences ever registered.
	Evidences []*Evidence `protobuf:"bytes,3,rep,name=evidences,proto3" json:"evidences,omitempty"`
	// votes_sigs contains all the votes of finality providers ever registered.
	VoteSigs []*VoteSig `protobuf:"bytes,4,rep,name=vote_sigs,json=voteSigs,proto3" json:"vote_sigs,omitempty"`
	// public_randomness contains all the public randomness ever commited from the finality providers.
	PublicRandomness []*PublicRandomness `protobuf:"bytes,5,rep,name=public_randomness,json=publicRandomness,proto3" json:"public_randomness,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dc577f74d797d1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetIndexedBlocks() []*IndexedBlock {
	if m != nil {
		return m.IndexedBlocks
	}
	return nil
}

func (m *GenesisState) GetEvidences() []*Evidence {
	if m != nil {
		return m.Evidences
	}
	return nil
}

func (m *GenesisState) GetVoteSigs() []*VoteSig {
	if m != nil {
		return m.VoteSigs
	}
	return nil
}

func (m *GenesisState) GetPublicRandomness() []*PublicRandomness {
	if m != nil {
		return m.PublicRandomness
	}
	return nil
}

// VoteSig the vote of an finality provider
// with the block of the vote, the finality provider btc public key and the vote signature.
type VoteSig struct {
	// block_height is the height of the voted block.
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,2,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// finality_sig is the finality signature to this block
	// where finality signature is an EOTS signature, i.e.
	FinalitySig *github_com_babylonchain_babylon_types.SchnorrEOTSSig `protobuf:"bytes,3,opt,name=finality_sig,json=finalitySig,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrEOTSSig" json:"finality_sig,omitempty"`
}

func (m *VoteSig) Reset()         { *m = VoteSig{} }
func (m *VoteSig) String() string { return proto.CompactTextString(m) }
func (*VoteSig) ProtoMessage()    {}
func (*VoteSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dc577f74d797d1, []int{1}
}
func (m *VoteSig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteSig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteSig.Merge(m, src)
}
func (m *VoteSig) XXX_Size() int {
	return m.Size()
}
func (m *VoteSig) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteSig.DiscardUnknown(m)
}

var xxx_messageInfo_VoteSig proto.InternalMessageInfo

func (m *VoteSig) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// PublicRandomness the block height and public randomness that the finality provider has submitted.
type PublicRandomness struct {
	// block_height is the height of block which the finality provider submited public randomness.
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// fp_btc_pk is the BTC PK of the finality provider that casts this vote.
	FpBtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,2,opt,name=fp_btc_pk,json=fpBtcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"fp_btc_pk,omitempty"`
	// pub_rand is the public randomness the finality provider has committed to.
	PubRand *github_com_babylonchain_babylon_types.SchnorrPubRand `protobuf:"bytes,3,opt,name=pub_rand,json=pubRand,proto3,customtype=github.com/babylonchain/babylon/types.SchnorrPubRand" json:"pub_rand,omitempty"`
}

func (m *PublicRandomness) Reset()         { *m = PublicRandomness{} }
func (m *PublicRandomness) String() string { return proto.CompactTextString(m) }
func (*PublicRandomness) ProtoMessage()    {}
func (*PublicRandomness) Descriptor() ([]byte, []int) {
	return fileDescriptor_52dc577f74d797d1, []int{2}
}
func (m *PublicRandomness) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicRandomness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicRandomness.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicRandomness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicRandomness.Merge(m, src)
}
func (m *PublicRandomness) XXX_Size() int {
	return m.Size()
}
func (m *PublicRandomness) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicRandomness.DiscardUnknown(m)
}

var xxx_messageInfo_PublicRandomness proto.InternalMessageInfo

func (m *PublicRandomness) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "babylon.finality.v1.GenesisState")
	proto.RegisterType((*VoteSig)(nil), "babylon.finality.v1.VoteSig")
	proto.RegisterType((*PublicRandomness)(nil), "babylon.finality.v1.PublicRandomness")
}

func init() { proto.RegisterFile("babylon/finality/v1/genesis.proto", fileDescriptor_52dc577f74d797d1) }

var fileDescriptor_52dc577f74d797d1 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0x87, 0x9b, 0xb6, 0xac, 0xab, 0x5b, 0xd0, 0x08, 0x1c, 0xa2, 0x02, 0xe9, 0x1f, 0x09, 0xa9,
	0xa7, 0x64, 0xeb, 0x26, 0xc4, 0xc4, 0x2d, 0xd2, 0xc4, 0x06, 0x07, 0x22, 0x07, 0x71, 0x80, 0x43,
	0x14, 0xa7, 0x6e, 0x62, 0xb5, 0xb5, 0xad, 0xd8, 0x89, 0xd6, 0x6f, 0xc1, 0xc7, 0xda, 0x71, 0x47,
	0x34, 0x89, 0x82, 0xda, 0x2f, 0x82, 0xe2, 0xa4, 0x1b, 0x9a, 0x22, 0x81, 0xb8, 0x70, 0xb3, 0x9d,
	0xe7, 0x7d, 0xfc, 0xbe, 0xbf, 0xc8, 0x60, 0x88, 0x02, 0xb4, 0x5a, 0x30, 0x6a, 0xcf, 0x08, 0x0d,
	0x16, 0x44, 0xae, 0xec, 0xec, 0xc8, 0x8e, 0x30, 0xc5, 0x82, 0x08, 0x8b, 0x27, 0x4c, 0x32, 0xfd,
	0x49, 0x89, 0x58, 0x3b, 0xc4, 0xca, 0x8e, 0x7a, 0x4f, 0x23, 0x16, 0x31, 0xf5, 0xdd, 0xce, 0x57,
	0x05, 0xda, 0x1b, 0x54, 0xd9, 0x78, 0x90, 0x04, 0xcb, 0x52, 0xd6, 0x1b, 0x55, 0x11, 0xb7, 0x62,
	0xc5, 0x8c, 0x7e, 0xd4, 0x41, 0xf7, 0x6d, 0xd1, 0x82, 0x27, 0x03, 0x89, 0xf5, 0x53, 0xb0, 0x57,
	0x48, 0x0c, 0x6d, 0xa0, 0x8d, 0x3b, 0x93, 0x67, 0x56, 0x45, 0x4b, 0x96, 0xab, 0x10, 0xa7, 0x79,
	0xb5, 0xee, 0xd7, 0x60, 0x59, 0xa0, 0x9f, 0x83, 0x47, 0x84, 0x4e, 0xf1, 0x25, 0x9e, 0xfa, 0x68,
	0xc1, 0xc2, 0xb9, 0x30, 0xea, 0x83, 0xc6, 0xb8, 0x33, 0x19, 0x56, 0x2a, 0x2e, 0x0a, 0xd4, 0xc9,
	0x49, 0xf8, 0x90, 0xfc, 0xb6, 0x13, 0xfa, 0x1b, 0xd0, 0xc6, 0x19, 0x99, 0x62, 0x1a, 0x62, 0x61,
	0x34, 0x94, 0xe4, 0x45, 0xa5, 0xe4, 0xac, 0xa4, 0xe0, 0x1d, 0xaf, 0x9f, 0x82, 0x76, 0xc6, 0x24,
	0xf6, 0x05, 0x89, 0x84, 0xd1, 0x54, 0xc5, 0xcf, 0x2b, 0x8b, 0x3f, 0x31, 0x89, 0x3d, 0x12, 0xc1,
	0xfd, 0xac, 0x58, 0x08, 0x1d, 0x82, 0xc7, 0x3c, 0x45, 0x0b, 0x12, 0xfa, 0x49, 0x40, 0xa7, 0x6c,
	0x49, 0xb1, 0x10, 0xc6, 0x03, 0xa5, 0x78, 0x59, 0x9d, 0x83, 0xa2, 0xe1, 0x2d, 0x0c, 0x0f, 0xf8,
	0xbd, 0x93, 0xd1, 0x77, 0x0d, 0xb4, 0xca, 0x9b, 0xf4, 0x21, 0xe8, 0xaa, 0x64, 0xfc, 0x18, 0x93,
	0x28, 0x96, 0x2a, 0xe2, 0x26, 0xec, 0xa8, 0xb3, 0x73, 0x75, 0xa4, 0x43, 0xd0, 0x9e, 0x71, 0x1f,
	0xc9, 0xd0, 0xe7, 0x73, 0xa3, 0x3e, 0xd0, 0xc6, 0x5d, 0xe7, 0xd5, 0xcd, 0xba, 0x3f, 0x89, 0x88,
	0x8c, 0x53, 0x64, 0x85, 0x6c, 0x69, 0x97, 0x8d, 0x84, 0x71, 0x40, 0xe8, 0x6e, 0x63, 0xcb, 0x15,
	0xc7, 0xc2, 0x72, 0x2e, 0xdc, 0xe3, 0x93, 0x43, 0x37, 0x45, 0xef, 0xf1, 0x0a, 0xb6, 0x66, 0xdc,
	0x91, 0xa1, 0x3b, 0xd7, 0xbf, 0x80, 0xee, 0xae, 0xe9, 0x3c, 0x15, 0xa3, 0xa1, 0xb4, 0xaf, 0x6f,
	0xd6, 0xfd, 0x93, 0xbf, 0xd3, 0x7a, 0x61, 0x4c, 0x59, 0x92, 0x9c, 0x7d, 0xf8, 0xe8, 0xe5, 0x81,
	0x75, 0x76, 0x36, 0x8f, 0x44, 0xa3, 0xb5, 0x06, 0x0e, 0xee, 0xc7, 0xf0, 0xbf, 0x06, 0xf5, 0xc0,
	0x3e, 0x4f, 0x91, 0xfa, 0x79, 0xff, 0x3c, 0xa4, 0x9b, 0xa2, 0x7c, 0x10, 0xd8, 0xe2, 0xc5, 0xc2,
	0x79, 0x77, 0xb5, 0x31, 0xb5, 0xeb, 0x8d, 0xa9, 0xfd, 0xdc, 0x98, 0xda, 0xd7, 0xad, 0x59, 0xbb,
	0xde, 0x9a, 0xb5, 0x6f, 0x5b, 0xb3, 0xf6, 0xf9, 0xf0, 0x4f, 0xe2, 0xcb, 0xbb, 0xa7, 0xa7, 0xee,
	0x40, 0x7b, 0xea, 0xd5, 0x1d, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x48, 0x67, 0xec, 0x5e, 0x0b,
	0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicRandomness) > 0 {
		for iNdEx := len(m.PublicRandomness) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PublicRandomness[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VoteSigs) > 0 {
		for iNdEx := len(m.VoteSigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VoteSigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Evidences) > 0 {
		for iNdEx := len(m.Evidences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.IndexedBlocks) > 0 {
		for iNdEx := len(m.IndexedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IndexedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VoteSig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteSig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteSig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalitySig != nil {
		{
			size := m.FinalitySig.Size()
			i -= size
			if _, err := m.FinalitySig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PublicRandomness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicRandomness) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicRandomness) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PubRand != nil {
		{
			size := m.PubRand.Size()
			i -= size
			if _, err := m.PubRand.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FpBtcPk != nil {
		{
			size := m.FpBtcPk.Size()
			i -= size
			if _, err := m.FpBtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.IndexedBlocks) > 0 {
		for _, e := range m.IndexedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Evidences) > 0 {
		for _, e := range m.Evidences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VoteSigs) > 0 {
		for _, e := range m.VoteSigs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PublicRandomness) > 0 {
		for _, e := range m.PublicRandomness {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *VoteSig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.BlockHeight))
	}
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.FinalitySig != nil {
		l = m.FinalitySig.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *PublicRandomness) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.BlockHeight))
	}
	if m.FpBtcPk != nil {
		l = m.FpBtcPk.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PubRand != nil {
		l = m.PubRand.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndexedBlocks = append(m.IndexedBlocks, &IndexedBlock{})
			if err := m.IndexedBlocks[len(m.IndexedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidences = append(m.Evidences, &Evidence{})
			if err := m.Evidences[len(m.Evidences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteSigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteSigs = append(m.VoteSigs, &VoteSig{})
			if err := m.VoteSigs[len(m.VoteSigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicRandomness", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicRandomness = append(m.PublicRandomness, &PublicRandomness{})
			if err := m.PublicRandomness[len(m.PublicRandomness)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *VoteSig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: VoteSig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteSig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalitySig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrEOTSSig
			m.FinalitySig = &v
			if err := m.FinalitySig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PublicRandomness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PublicRandomness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicRandomness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FpBtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.FpBtcPk = &v
			if err := m.FpBtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubRand", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.SchnorrPubRand
			m.PubRand = &v
			if err := m.PubRand.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
