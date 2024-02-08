// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/incentive.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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

// VotingPowerDistCache is the cache for voting power distribution of finality providers
// and their BTC delegations at a height
type VotingPowerDistCache struct {
	TotalVotingPower uint64 `protobuf:"varint,1,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	// finality_providers is a list of finality providers' voting power information
	FinalityProviders []*FinalityProviderDistInfo `protobuf:"bytes,2,rep,name=finality_providers,json=finalityProviders,proto3" json:"finality_providers,omitempty"`
}

func (m *VotingPowerDistCache) Reset()         { *m = VotingPowerDistCache{} }
func (m *VotingPowerDistCache) String() string { return proto.CompactTextString(m) }
func (*VotingPowerDistCache) ProtoMessage()    {}
func (*VotingPowerDistCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{0}
}
func (m *VotingPowerDistCache) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VotingPowerDistCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VotingPowerDistCache.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VotingPowerDistCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingPowerDistCache.Merge(m, src)
}
func (m *VotingPowerDistCache) XXX_Size() int {
	return m.Size()
}
func (m *VotingPowerDistCache) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingPowerDistCache.DiscardUnknown(m)
}

var xxx_messageInfo_VotingPowerDistCache proto.InternalMessageInfo

func (m *VotingPowerDistCache) GetTotalVotingPower() uint64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *VotingPowerDistCache) GetFinalityProviders() []*FinalityProviderDistInfo {
	if m != nil {
		return m.FinalityProviders
	}
	return nil
}

// FinalityProviderDistInfo is the reward distribution of a finality provider and its BTC delegations
type FinalityProviderDistInfo struct {
	// btc_pk is the Bitcoin secp256k1 PK of this finality provider
	// the PK follows encoding in BIP-340 spec
	BtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=btc_pk,json=btcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"btc_pk,omitempty"`
	// babylon_pk is the Babylon public key of the finality provider
	BabylonPk *secp256k1.PubKey `protobuf:"bytes,2,opt,name=babylon_pk,json=babylonPk,proto3" json:"babylon_pk,omitempty"`
	// commission defines the commission rate of finality provider
	Commission *cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=commission,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"commission,omitempty"`
	// total_voting_power is the total voting power of the finality provider
	TotalVotingPower uint64 `protobuf:"varint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	// btc_dels is a list of BTC delegations' voting power information under this finality provider
	BtcDels []*BTCDelDistInfo `protobuf:"bytes,5,rep,name=btc_dels,json=btcDels,proto3" json:"btc_dels,omitempty"`
}

func (m *FinalityProviderDistInfo) Reset()         { *m = FinalityProviderDistInfo{} }
func (m *FinalityProviderDistInfo) String() string { return proto.CompactTextString(m) }
func (*FinalityProviderDistInfo) ProtoMessage()    {}
func (*FinalityProviderDistInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{1}
}
func (m *FinalityProviderDistInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalityProviderDistInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalityProviderDistInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalityProviderDistInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalityProviderDistInfo.Merge(m, src)
}
func (m *FinalityProviderDistInfo) XXX_Size() int {
	return m.Size()
}
func (m *FinalityProviderDistInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalityProviderDistInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FinalityProviderDistInfo proto.InternalMessageInfo

func (m *FinalityProviderDistInfo) GetBabylonPk() *secp256k1.PubKey {
	if m != nil {
		return m.BabylonPk
	}
	return nil
}

func (m *FinalityProviderDistInfo) GetTotalVotingPower() uint64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *FinalityProviderDistInfo) GetBtcDels() []*BTCDelDistInfo {
	if m != nil {
		return m.BtcDels
	}
	return nil
}

// BTCDelDistInfo contains the information related to reward distribution for a BTC delegations
type BTCDelDistInfo struct {
	// btc_pk is the Bitcoin secp256k1 PK of this BTC delegation
	// the PK follows encoding in BIP-340 spec
	BtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=btc_pk,json=btcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"btc_pk,omitempty"`
	// babylon_pk is the Babylon public key of the BTC delegations
	BabylonPk *secp256k1.PubKey `protobuf:"bytes,2,opt,name=babylon_pk,json=babylonPk,proto3" json:"babylon_pk,omitempty"`
	// voting_power is the voting power of the BTC delegation
	VotingPower uint64 `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *BTCDelDistInfo) Reset()         { *m = BTCDelDistInfo{} }
func (m *BTCDelDistInfo) String() string { return proto.CompactTextString(m) }
func (*BTCDelDistInfo) ProtoMessage()    {}
func (*BTCDelDistInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{2}
}
func (m *BTCDelDistInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCDelDistInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCDelDistInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCDelDistInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCDelDistInfo.Merge(m, src)
}
func (m *BTCDelDistInfo) XXX_Size() int {
	return m.Size()
}
func (m *BTCDelDistInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCDelDistInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BTCDelDistInfo proto.InternalMessageInfo

func (m *BTCDelDistInfo) GetBabylonPk() *secp256k1.PubKey {
	if m != nil {
		return m.BabylonPk
	}
	return nil
}

func (m *BTCDelDistInfo) GetVotingPower() uint64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*VotingPowerDistCache)(nil), "babylon.btcstaking.v1.VotingPowerDistCache")
	proto.RegisterType((*FinalityProviderDistInfo)(nil), "babylon.btcstaking.v1.FinalityProviderDistInfo")
	proto.RegisterType((*BTCDelDistInfo)(nil), "babylon.btcstaking.v1.BTCDelDistInfo")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/incentive.proto", fileDescriptor_ac354c3bd6d7a66b)
}

var fileDescriptor_ac354c3bd6d7a66b = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x76, 0x1b, 0xcc, 0x9d, 0x10, 0x44, 0x43, 0x0a, 0x43, 0x4a, 0x4b, 0xa5, 0x49,
	0x3d, 0x30, 0x9b, 0x76, 0xb0, 0x23, 0x42, 0x59, 0x84, 0x34, 0xb1, 0x49, 0x51, 0x85, 0x38, 0x70,
	0xa0, 0x8a, 0x5d, 0x37, 0xb5, 0x92, 0xd8, 0x51, 0xed, 0x05, 0xf2, 0x16, 0x3c, 0x04, 0x8f, 0xc0,
	0x13, 0x70, 0xe2, 0x38, 0x71, 0x42, 0x3b, 0x4c, 0xa8, 0xbd, 0xf0, 0x18, 0x28, 0x89, 0xd9, 0x5a,
	0x44, 0xc4, 0x95, 0x9b, 0x3f, 0xfd, 0xff, 0xff, 0xef, 0xfb, 0xfc, 0xb3, 0x0c, 0xf7, 0x49, 0x40,
	0xf2, 0x58, 0x0a, 0x4c, 0x34, 0x55, 0x3a, 0x88, 0xb8, 0x08, 0x71, 0x36, 0xc0, 0x5c, 0x50, 0x26,
	0x34, 0xcf, 0x18, 0x4a, 0xe7, 0x52, 0x4b, 0xeb, 0xbe, 0xb1, 0xa1, 0x1b, 0x1b, 0xca, 0x06, 0x7b,
	0xbb, 0xa1, 0x0c, 0x65, 0xe9, 0xc0, 0xc5, 0xa9, 0x32, 0xef, 0x3d, 0xa0, 0x52, 0x25, 0x52, 0x8d,
	0x2b, 0xa1, 0x2a, 0x8c, 0xd4, 0xab, 0x2a, 0x4c, 0xe7, 0x79, 0xaa, 0x25, 0x56, 0x8c, 0xa6, 0xc3,
	0x67, 0x47, 0xd1, 0x00, 0x47, 0x2c, 0x37, 0x9e, 0xde, 0x27, 0x00, 0x77, 0xdf, 0x48, 0xcd, 0x45,
	0xe8, 0xcb, 0xf7, 0x6c, 0xee, 0x71, 0xa5, 0x8f, 0x03, 0x3a, 0x63, 0xd6, 0x63, 0x68, 0x69, 0xa9,
	0x83, 0x78, 0x9c, 0x95, 0xea, 0x38, 0x2d, 0x64, 0x1b, 0x74, 0x41, 0x7f, 0x63, 0x74, 0xb7, 0x54,
	0x56, 0x62, 0xd6, 0x3b, 0x68, 0x4d, 0xb9, 0x08, 0x62, 0xae, 0xf3, 0x62, 0x93, 0x8c, 0x4f, 0xd8,
	0x5c, 0xd9, 0xcd, 0x6e, 0xab, 0xdf, 0x1e, 0x62, 0xf4, 0xd7, 0xfb, 0xa0, 0x97, 0x26, 0xe0, 0x1b,
	0x7f, 0x31, 0xfb, 0x44, 0x4c, 0xe5, 0xe8, 0xde, 0xf4, 0x0f, 0x45, 0xf5, 0x7e, 0x36, 0xa1, 0x5d,
	0xe7, 0xb7, 0xce, 0xe0, 0x16, 0xd1, 0x74, 0x9c, 0x46, 0xe5, 0x7a, 0x3b, 0xee, 0xd1, 0xe5, 0x55,
	0x67, 0x18, 0x72, 0x3d, 0x3b, 0x27, 0x88, 0xca, 0x04, 0x9b, 0xf1, 0x74, 0x16, 0x70, 0xf1, 0xbb,
	0xc0, 0x3a, 0x4f, 0x99, 0x42, 0xee, 0x89, 0x7f, 0xf8, 0xf4, 0x89, 0x7f, 0x4e, 0x5e, 0xb1, 0x7c,
	0xb4, 0x49, 0x34, 0xf5, 0x23, 0xeb, 0x39, 0x84, 0xc6, 0x54, 0xb4, 0x6c, 0x76, 0x41, 0xbf, 0x3d,
	0xec, 0x20, 0x43, 0xb6, 0x62, 0x89, 0xae, 0x59, 0x22, 0x93, 0xdd, 0x36, 0x11, 0x3f, 0xb2, 0xce,
	0x20, 0xa4, 0x32, 0x49, 0xb8, 0x52, 0x5c, 0x0a, 0xbb, 0xd5, 0x05, 0xfd, 0x6d, 0xf7, 0xe0, 0xf2,
	0xaa, 0xf3, 0xb0, 0x6a, 0xa1, 0x26, 0x11, 0xe2, 0x12, 0x27, 0x81, 0x9e, 0xa1, 0x53, 0x16, 0x06,
	0x34, 0xf7, 0x18, 0xfd, 0xf6, 0xf9, 0x00, 0x9a, 0x09, 0x1e, 0xa3, 0xa3, 0x95, 0x06, 0x35, 0x0f,
	0xb1, 0x51, 0xf3, 0x10, 0x2f, 0xe0, 0xed, 0x82, 0xc5, 0x84, 0xc5, 0xca, 0xde, 0x2c, 0xf1, 0xef,
	0xd7, 0xe0, 0x77, 0x5f, 0x1f, 0x7b, 0x2c, 0xbe, 0x86, 0x7e, 0x8b, 0x68, 0xea, 0xb1, 0x58, 0xf5,
	0xbe, 0x00, 0x78, 0x67, 0x5d, 0xfb, 0xdf, 0x00, 0x3f, 0x82, 0x3b, 0x6b, 0x2c, 0x5a, 0x25, 0x8b,
	0x76, 0x76, 0x83, 0xc1, 0x3d, 0xfd, 0xba, 0x70, 0xc0, 0xc5, 0xc2, 0x01, 0x3f, 0x16, 0x0e, 0xf8,
	0xb8, 0x74, 0x1a, 0x17, 0x4b, 0xa7, 0xf1, 0x7d, 0xe9, 0x34, 0xde, 0xfe, 0x73, 0xef, 0x0f, 0xab,
	0xbf, 0xb3, 0xbc, 0x04, 0xd9, 0x2a, 0xff, 0xca, 0xe1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69,
	0x5e, 0x2e, 0x45, 0xc0, 0x03, 0x00, 0x00,
}

func (m *VotingPowerDistCache) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VotingPowerDistCache) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VotingPowerDistCache) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalityProviders) > 0 {
		for iNdEx := len(m.FinalityProviders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FinalityProviders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FinalityProviderDistInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalityProviderDistInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalityProviderDistInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BtcDels) > 0 {
		for iNdEx := len(m.BtcDels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BtcDels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x20
	}
	if m.Commission != nil {
		{
			size := m.Commission.Size()
			i -= size
			if _, err := m.Commission.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BabylonPk != nil {
		{
			size, err := m.BabylonPk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BtcPk != nil {
		{
			size := m.BtcPk.Size()
			i -= size
			if _, err := m.BtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BTCDelDistInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCDelDistInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCDelDistInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.BabylonPk != nil {
		{
			size, err := m.BabylonPk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BtcPk != nil {
		{
			size := m.BtcPk.Size()
			i -= size
			if _, err := m.BtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VotingPowerDistCache) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalVotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.TotalVotingPower))
	}
	if len(m.FinalityProviders) > 0 {
		for _, e := range m.FinalityProviders {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *FinalityProviderDistInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcPk != nil {
		l = m.BtcPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.BabylonPk != nil {
		l = m.BabylonPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.Commission != nil {
		l = m.Commission.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.TotalVotingPower))
	}
	if len(m.BtcDels) > 0 {
		for _, e := range m.BtcDels {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *BTCDelDistInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcPk != nil {
		l = m.BtcPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.BabylonPk != nil {
		l = m.BabylonPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.VotingPower))
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VotingPowerDistCache) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: VotingPowerDistCache: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VotingPowerDistCache: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalityProviders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalityProviders = append(m.FinalityProviders, &FinalityProviderDistInfo{})
			if err := m.FinalityProviders[len(m.FinalityProviders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *FinalityProviderDistInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: FinalityProviderDistInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalityProviderDistInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.BtcPk = &v
			if err := m.BtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonPk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BabylonPk == nil {
				m.BabylonPk = &secp256k1.PubKey{}
			}
			if err := m.BabylonPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.Commission = &v
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcDels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcDels = append(m.BtcDels, &BTCDelDistInfo{})
			if err := m.BtcDels[len(m.BtcDels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *BTCDelDistInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: BTCDelDistInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCDelDistInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.BtcPk = &v
			if err := m.BtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonPk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BabylonPk == nil {
				m.BabylonPk = &secp256k1.PubKey{}
			}
			if err := m.BabylonPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
