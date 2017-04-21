// Code generated by protoc-gen-go.
// source: data_redeem.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type PokeCandy struct {
	PokemonId  uint64 `protobuf:"fixed64,1,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
	CandyCount int32  `protobuf:"varint,2,opt,name=candy_count,json=candyCount" json:"candy_count,omitempty"`
}

func (m *PokeCandy) Reset()                    { *m = PokeCandy{} }
func (m *PokeCandy) String() string            { return proto.CompactTextString(m) }
func (*PokeCandy) ProtoMessage()               {}
func (*PokeCandy) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *PokeCandy) GetPokemonId() uint64 {
	if m != nil {
		return m.PokemonId
	}
	return 0
}

func (m *PokeCandy) GetCandyCount() int32 {
	if m != nil {
		return m.CandyCount
	}
	return 0
}

type RedeemedAvatarItem struct {
	AvatarTemplateId string `protobuf:"bytes,1,opt,name=avatar_template_id,json=avatarTemplateId" json:"avatar_template_id,omitempty"`
	ItemCount        int32  `protobuf:"varint,2,opt,name=item_count,json=itemCount" json:"item_count,omitempty"`
}

func (m *RedeemedAvatarItem) Reset()                    { *m = RedeemedAvatarItem{} }
func (m *RedeemedAvatarItem) String() string            { return proto.CompactTextString(m) }
func (*RedeemedAvatarItem) ProtoMessage()               {}
func (*RedeemedAvatarItem) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *RedeemedAvatarItem) GetAvatarTemplateId() string {
	if m != nil {
		return m.AvatarTemplateId
	}
	return ""
}

func (m *RedeemedAvatarItem) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

type RedeemedItem struct {
	Item      ItemId `protobuf:"varint,1,opt,name=item,enum=SUNProtos.Inventory.Item.ItemId" json:"item,omitempty"`
	ItemCount int32  `protobuf:"varint,2,opt,name=item_count,json=itemCount" json:"item_count,omitempty"`
}

func (m *RedeemedItem) Reset()                    { *m = RedeemedItem{} }
func (m *RedeemedItem) String() string            { return proto.CompactTextString(m) }
func (*RedeemedItem) ProtoMessage()               {}
func (*RedeemedItem) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *RedeemedItem) GetItem() ItemId {
	if m != nil {
		return m.Item
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *RedeemedItem) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PokeCandy)(nil), "SUNProtos.Data.Redeem.PokeCandy")
	proto.RegisterType((*RedeemedAvatarItem)(nil), "SUNProtos.Data.Redeem.RedeemedAvatarItem")
	proto.RegisterType((*RedeemedItem)(nil), "SUNProtos.Data.Redeem.RedeemedItem")
}

func init() { proto.RegisterFile("data_redeem.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x68, 0x31, 0xa3, 0x88, 0x2e, 0x0a, 0x45, 0x10, 0x43, 0x4f, 0x3d, 0xc8, 0x1e,
	0xd4, 0x3f, 0xa0, 0xf5, 0xb2, 0x08, 0x12, 0x56, 0xbd, 0x78, 0x59, 0xc6, 0xec, 0x1c, 0x42, 0xcd,
	0x6e, 0x48, 0xc7, 0x42, 0xff, 0xbd, 0xec, 0xa4, 0x8b, 0x78, 0xea, 0x25, 0x4c, 0xde, 0x7b, 0xbc,
	0x6f, 0x66, 0xe1, 0xdc, 0x23, 0xa3, 0x1b, 0xc8, 0x13, 0x75, 0xba, 0x1f, 0x22, 0x47, 0x75, 0xf9,
	0xf6, 0xf1, 0x5a, 0xa7, 0x69, 0xa5, 0x9f, 0x91, 0x51, 0x5b, 0x31, 0xaf, 0x2e, 0xda, 0xb0, 0xa6,
	0xc0, 0x71, 0xd8, 0xb8, 0x96, 0x73, 0x78, 0xf6, 0x02, 0x65, 0x1d, 0x97, 0xb4, 0xc0, 0xe0, 0x37,
	0xea, 0x1a, 0xa0, 0x8f, 0x4b, 0xea, 0x62, 0x70, 0xad, 0x9f, 0x16, 0x55, 0x31, 0x9f, 0xd8, 0x72,
	0xab, 0x18, 0xaf, 0x6e, 0xe0, 0xb8, 0x49, 0x39, 0xd7, 0xc4, 0x9f, 0xc0, 0xd3, 0xfd, 0xaa, 0x98,
	0x1f, 0x5a, 0x10, 0x69, 0x91, 0x94, 0x19, 0x82, 0x1a, 0x61, 0xe4, 0x1f, 0xd7, 0xc8, 0x38, 0x18,
	0xa6, 0x4e, 0xdd, 0x82, 0x42, 0xf9, 0x73, 0x4c, 0x5d, 0xff, 0x8d, 0x4c, 0xb9, 0xbd, 0xb4, 0x67,
	0xa3, 0xf3, 0xbe, 0x35, 0x8c, 0x4f, 0x3b, 0xa4, 0xf5, 0xfe, 0x31, 0xca, 0xa4, 0x8c, 0x88, 0x06,
	0x4e, 0x32, 0x42, 0xca, 0x1f, 0xe0, 0x20, 0x99, 0x52, 0x77, 0x7a, 0x57, 0xe9, 0xbf, 0xdb, 0x4d,
	0x3e, 0x57, 0xa7, 0xa0, 0x7c, 0x8c, 0xb7, 0x92, 0xde, 0x01, 0x79, 0x3a, 0xfa, 0x9c, 0xc8, 0xeb,
	0xac, 0xea, 0xbd, 0xaf, 0x71, 0xba, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xdd, 0x79, 0x5d,
	0x69, 0x01, 0x00, 0x00,
}
