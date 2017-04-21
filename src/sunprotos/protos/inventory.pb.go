// Code generated by protoc-gen-go.
// source: inventory.proto
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

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Costume from enums.proto

// Ignoring public import of EncounterType from enums.proto

// Ignoring public import of Filter from enums.proto

// Ignoring public import of Form from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of NotificationState from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of QuestType from enums.proto

// Ignoring public import of Slot from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of AvatarCustomization from data_avatar.proto

// Ignoring public import of AvatarItem from data_avatar.proto

// Ignoring public import of Label from data_avatar.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of BackgroundToken from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of ClientVersion from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of PokemonDisplay from data.proto

// Ignoring public import of RedeemPasscodeReward from data.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerAvatarType from data_player.proto

// Ignoring public import of DailyQuest from data_quests.proto

// Ignoring public import of Quest from data_quests.proto

type EggIncubatorType int32

const (
	EggIncubatorType_INCUBATOR_UNSET    EggIncubatorType = 0
	EggIncubatorType_INCUBATOR_DISTANCE EggIncubatorType = 1
)

var EggIncubatorType_name = map[int32]string{
	0: "INCUBATOR_UNSET",
	1: "INCUBATOR_DISTANCE",
}
var EggIncubatorType_value = map[string]int32{
	"INCUBATOR_UNSET":    0,
	"INCUBATOR_DISTANCE": 1,
}

func (x EggIncubatorType) String() string {
	return proto.EnumName(EggIncubatorType_name, int32(x))
}
func (EggIncubatorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type InventoryUpgradeType int32

const (
	InventoryUpgradeType_UPGRADE_UNSET            InventoryUpgradeType = 0
	InventoryUpgradeType_INCREASE_ITEM_STORAGE    InventoryUpgradeType = 1
	InventoryUpgradeType_INCREASE_POKEMON_STORAGE InventoryUpgradeType = 2
)

var InventoryUpgradeType_name = map[int32]string{
	0: "UPGRADE_UNSET",
	1: "INCREASE_ITEM_STORAGE",
	2: "INCREASE_POKEMON_STORAGE",
}
var InventoryUpgradeType_value = map[string]int32{
	"UPGRADE_UNSET":            0,
	"INCREASE_ITEM_STORAGE":    1,
	"INCREASE_POKEMON_STORAGE": 2,
}

func (x InventoryUpgradeType) String() string {
	return proto.EnumName(InventoryUpgradeType_name, int32(x))
}
func (InventoryUpgradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

type AppliedItem struct {
	ItemId    ItemId   `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=SUNProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemType  ItemType `protobuf:"varint,2,opt,name=item_type,json=itemType,enum=SUNProtos.Inventory.Item.ItemType" json:"item_type,omitempty"`
	ExpireMs  int64    `protobuf:"varint,3,opt,name=expire_ms,json=expireMs" json:"expire_ms,omitempty"`
	AppliedMs int64    `protobuf:"varint,4,opt,name=applied_ms,json=appliedMs" json:"applied_ms,omitempty"`
}

func (m *AppliedItem) Reset()                    { *m = AppliedItem{} }
func (m *AppliedItem) String() string            { return proto.CompactTextString(m) }
func (*AppliedItem) ProtoMessage()               {}
func (*AppliedItem) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *AppliedItem) GetItemId() ItemId {
	if m != nil {
		return m.ItemId
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *AppliedItem) GetItemType() ItemType {
	if m != nil {
		return m.ItemType
	}
	return ItemType_ITEM_TYPE_NONE
}

func (m *AppliedItem) GetExpireMs() int64 {
	if m != nil {
		return m.ExpireMs
	}
	return 0
}

func (m *AppliedItem) GetAppliedMs() int64 {
	if m != nil {
		return m.AppliedMs
	}
	return 0
}

type AppliedItems struct {
	Item []*AppliedItem `protobuf:"bytes,4,rep,name=item" json:"item,omitempty"`
}

func (m *AppliedItems) Reset()                    { *m = AppliedItems{} }
func (m *AppliedItems) String() string            { return proto.CompactTextString(m) }
func (*AppliedItems) ProtoMessage()               {}
func (*AppliedItems) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *AppliedItems) GetItem() []*AppliedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type Candy struct {
	FamilyId PokemonFamilyId `protobuf:"varint,1,opt,name=family_id,json=familyId,enum=SUNProtos.Enums.PokemonFamilyId" json:"family_id,omitempty"`
	Candy    int32           `protobuf:"varint,2,opt,name=candy" json:"candy,omitempty"`
}

func (m *Candy) Reset()                    { *m = Candy{} }
func (m *Candy) String() string            { return proto.CompactTextString(m) }
func (*Candy) ProtoMessage()               {}
func (*Candy) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *Candy) GetFamilyId() PokemonFamilyId {
	if m != nil {
		return m.FamilyId
	}
	return PokemonFamilyId_FAMILY_UNSET
}

func (m *Candy) GetCandy() int32 {
	if m != nil {
		return m.Candy
	}
	return 0
}

type EggIncubator struct {
	Id             string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ItemId         ItemId           `protobuf:"varint,2,opt,name=item_id,json=itemId,enum=SUNProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	IncubatorType  EggIncubatorType `protobuf:"varint,3,opt,name=incubator_type,json=incubatorType,enum=SUNProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	UsesRemaining  int32            `protobuf:"varint,4,opt,name=uses_remaining,json=usesRemaining" json:"uses_remaining,omitempty"`
	PokemonId      uint64           `protobuf:"varint,5,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
	StartKmWalked  float64          `protobuf:"fixed64,6,opt,name=start_km_walked,json=startKmWalked" json:"start_km_walked,omitempty"`
	TargetKmWalked float64          `protobuf:"fixed64,7,opt,name=target_km_walked,json=targetKmWalked" json:"target_km_walked,omitempty"`
}

func (m *EggIncubator) Reset()                    { *m = EggIncubator{} }
func (m *EggIncubator) String() string            { return proto.CompactTextString(m) }
func (*EggIncubator) ProtoMessage()               {}
func (*EggIncubator) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *EggIncubator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EggIncubator) GetItemId() ItemId {
	if m != nil {
		return m.ItemId
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *EggIncubator) GetIncubatorType() EggIncubatorType {
	if m != nil {
		return m.IncubatorType
	}
	return EggIncubatorType_INCUBATOR_UNSET
}

func (m *EggIncubator) GetUsesRemaining() int32 {
	if m != nil {
		return m.UsesRemaining
	}
	return 0
}

func (m *EggIncubator) GetPokemonId() uint64 {
	if m != nil {
		return m.PokemonId
	}
	return 0
}

func (m *EggIncubator) GetStartKmWalked() float64 {
	if m != nil {
		return m.StartKmWalked
	}
	return 0
}

func (m *EggIncubator) GetTargetKmWalked() float64 {
	if m != nil {
		return m.TargetKmWalked
	}
	return 0
}

type EggIncubators struct {
	EggIncubator []*EggIncubator `protobuf:"bytes,1,rep,name=egg_incubator,json=eggIncubator" json:"egg_incubator,omitempty"`
}

func (m *EggIncubators) Reset()                    { *m = EggIncubators{} }
func (m *EggIncubators) String() string            { return proto.CompactTextString(m) }
func (*EggIncubators) ProtoMessage()               {}
func (*EggIncubators) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *EggIncubators) GetEggIncubator() []*EggIncubator {
	if m != nil {
		return m.EggIncubator
	}
	return nil
}

type InventoryDelta struct {
	OriginalTimestampMs int64            `protobuf:"varint,1,opt,name=original_timestamp_ms,json=originalTimestampMs" json:"original_timestamp_ms,omitempty"`
	NewTimestampMs      int64            `protobuf:"varint,2,opt,name=new_timestamp_ms,json=newTimestampMs" json:"new_timestamp_ms,omitempty"`
	InventoryItems      []*InventoryItem `protobuf:"bytes,3,rep,name=inventory_items,json=inventoryItems" json:"inventory_items,omitempty"`
}

func (m *InventoryDelta) Reset()                    { *m = InventoryDelta{} }
func (m *InventoryDelta) String() string            { return proto.CompactTextString(m) }
func (*InventoryDelta) ProtoMessage()               {}
func (*InventoryDelta) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *InventoryDelta) GetOriginalTimestampMs() int64 {
	if m != nil {
		return m.OriginalTimestampMs
	}
	return 0
}

func (m *InventoryDelta) GetNewTimestampMs() int64 {
	if m != nil {
		return m.NewTimestampMs
	}
	return 0
}

func (m *InventoryDelta) GetInventoryItems() []*InventoryItem {
	if m != nil {
		return m.InventoryItems
	}
	return nil
}

type InventoryItem struct {
	ModifiedTimestampMs int64                      `protobuf:"varint,1,opt,name=modified_timestamp_ms,json=modifiedTimestampMs" json:"modified_timestamp_ms,omitempty"`
	DeletedItem         *InventoryItem_DeletedItem `protobuf:"bytes,2,opt,name=deleted_item,json=deletedItem" json:"deleted_item,omitempty"`
	InventoryItemData   *InventoryItemData         `protobuf:"bytes,3,opt,name=inventory_item_data,json=inventoryItemData" json:"inventory_item_data,omitempty"`
}

func (m *InventoryItem) Reset()                    { *m = InventoryItem{} }
func (m *InventoryItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem) ProtoMessage()               {}
func (*InventoryItem) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *InventoryItem) GetModifiedTimestampMs() int64 {
	if m != nil {
		return m.ModifiedTimestampMs
	}
	return 0
}

func (m *InventoryItem) GetDeletedItem() *InventoryItem_DeletedItem {
	if m != nil {
		return m.DeletedItem
	}
	return nil
}

func (m *InventoryItem) GetInventoryItemData() *InventoryItemData {
	if m != nil {
		return m.InventoryItemData
	}
	return nil
}

type InventoryItem_DeletedItem struct {
	PokemonId uint64 `protobuf:"fixed64,1,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
}

func (m *InventoryItem_DeletedItem) Reset()                    { *m = InventoryItem_DeletedItem{} }
func (m *InventoryItem_DeletedItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem_DeletedItem) ProtoMessage()               {}
func (*InventoryItem_DeletedItem) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6, 0} }

func (m *InventoryItem_DeletedItem) GetPokemonId() uint64 {
	if m != nil {
		return m.PokemonId
	}
	return 0
}

type InventoryItemData struct {
	PokemonData       *PokemonData       `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	Item              *ItemData          `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	PokedexEntry      *PokedexEntry      `protobuf:"bytes,3,opt,name=pokedex_entry,json=pokedexEntry" json:"pokedex_entry,omitempty"`
	PlayerStats       *PlayerStats       `protobuf:"bytes,4,opt,name=player_stats,json=playerStats" json:"player_stats,omitempty"`
	PlayerCurrency    *PlayerCurrency    `protobuf:"bytes,5,opt,name=player_currency,json=playerCurrency" json:"player_currency,omitempty"`
	PlayerCamera      *PlayerCamera      `protobuf:"bytes,6,opt,name=player_camera,json=playerCamera" json:"player_camera,omitempty"`
	InventoryUpgrades *InventoryUpgrades `protobuf:"bytes,7,opt,name=inventory_upgrades,json=inventoryUpgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      *AppliedItems      `protobuf:"bytes,8,opt,name=applied_items,json=appliedItems" json:"applied_items,omitempty"`
	EggIncubators     *EggIncubators     `protobuf:"bytes,9,opt,name=egg_incubators,json=eggIncubators" json:"egg_incubators,omitempty"`
	Candy             *Candy             `protobuf:"bytes,10,opt,name=candy" json:"candy,omitempty"`
	Quest             *Quest             `protobuf:"bytes,11,opt,name=quest" json:"quest,omitempty"`
	AvatarItem        *AvatarItem        `protobuf:"bytes,12,opt,name=avatar_item,json=avatarItem" json:"avatar_item,omitempty"`
}

func (m *InventoryItemData) Reset()                    { *m = InventoryItemData{} }
func (m *InventoryItemData) String() string            { return proto.CompactTextString(m) }
func (*InventoryItemData) ProtoMessage()               {}
func (*InventoryItemData) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *InventoryItemData) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *InventoryItemData) GetItem() *ItemData {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InventoryItemData) GetPokedexEntry() *PokedexEntry {
	if m != nil {
		return m.PokedexEntry
	}
	return nil
}

func (m *InventoryItemData) GetPlayerStats() *PlayerStats {
	if m != nil {
		return m.PlayerStats
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCurrency() *PlayerCurrency {
	if m != nil {
		return m.PlayerCurrency
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCamera() *PlayerCamera {
	if m != nil {
		return m.PlayerCamera
	}
	return nil
}

func (m *InventoryItemData) GetInventoryUpgrades() *InventoryUpgrades {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func (m *InventoryItemData) GetAppliedItems() *AppliedItems {
	if m != nil {
		return m.AppliedItems
	}
	return nil
}

func (m *InventoryItemData) GetEggIncubators() *EggIncubators {
	if m != nil {
		return m.EggIncubators
	}
	return nil
}

func (m *InventoryItemData) GetCandy() *Candy {
	if m != nil {
		return m.Candy
	}
	return nil
}

func (m *InventoryItemData) GetQuest() *Quest {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *InventoryItemData) GetAvatarItem() *AvatarItem {
	if m != nil {
		return m.AvatarItem
	}
	return nil
}

type InventoryKey struct {
	PokemonId         uint64          `protobuf:"fixed64,1,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
	Item              ItemId          `protobuf:"varint,2,opt,name=item,enum=SUNProtos.Inventory.Item.ItemId" json:"item,omitempty"`
	PokedexEntryId    int32           `protobuf:"varint,3,opt,name=pokedex_entry_id,json=pokedexEntryId" json:"pokedex_entry_id,omitempty"`
	PlayerStats       bool            `protobuf:"varint,4,opt,name=player_stats,json=playerStats" json:"player_stats,omitempty"`
	PlayerCurrency    bool            `protobuf:"varint,5,opt,name=player_currency,json=playerCurrency" json:"player_currency,omitempty"`
	PlayerCamera      bool            `protobuf:"varint,6,opt,name=player_camera,json=playerCamera" json:"player_camera,omitempty"`
	InventoryUpgrades bool            `protobuf:"varint,7,opt,name=inventory_upgrades,json=inventoryUpgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      bool            `protobuf:"varint,8,opt,name=applied_items,json=appliedItems" json:"applied_items,omitempty"`
	EggIncubators     bool            `protobuf:"varint,9,opt,name=egg_incubators,json=eggIncubators" json:"egg_incubators,omitempty"`
	PokemonFamilyId   PokemonFamilyId `protobuf:"varint,10,opt,name=pokemon_family_id,json=pokemonFamilyId,enum=SUNProtos.Enums.PokemonFamilyId" json:"pokemon_family_id,omitempty"`
	QuestType         QuestType       `protobuf:"varint,11,opt,name=quest_type,json=questType,enum=SUNProtos.Enums.QuestType" json:"quest_type,omitempty"`
	AvatarTemplateId  string          `protobuf:"bytes,12,opt,name=avatar_template_id,json=avatarTemplateId" json:"avatar_template_id,omitempty"`
}

func (m *InventoryKey) Reset()                    { *m = InventoryKey{} }
func (m *InventoryKey) String() string            { return proto.CompactTextString(m) }
func (*InventoryKey) ProtoMessage()               {}
func (*InventoryKey) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *InventoryKey) GetPokemonId() uint64 {
	if m != nil {
		return m.PokemonId
	}
	return 0
}

func (m *InventoryKey) GetItem() ItemId {
	if m != nil {
		return m.Item
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *InventoryKey) GetPokedexEntryId() int32 {
	if m != nil {
		return m.PokedexEntryId
	}
	return 0
}

func (m *InventoryKey) GetPlayerStats() bool {
	if m != nil {
		return m.PlayerStats
	}
	return false
}

func (m *InventoryKey) GetPlayerCurrency() bool {
	if m != nil {
		return m.PlayerCurrency
	}
	return false
}

func (m *InventoryKey) GetPlayerCamera() bool {
	if m != nil {
		return m.PlayerCamera
	}
	return false
}

func (m *InventoryKey) GetInventoryUpgrades() bool {
	if m != nil {
		return m.InventoryUpgrades
	}
	return false
}

func (m *InventoryKey) GetAppliedItems() bool {
	if m != nil {
		return m.AppliedItems
	}
	return false
}

func (m *InventoryKey) GetEggIncubators() bool {
	if m != nil {
		return m.EggIncubators
	}
	return false
}

func (m *InventoryKey) GetPokemonFamilyId() PokemonFamilyId {
	if m != nil {
		return m.PokemonFamilyId
	}
	return PokemonFamilyId_FAMILY_UNSET
}

func (m *InventoryKey) GetQuestType() QuestType {
	if m != nil {
		return m.QuestType
	}
	return QuestType_QUEST_UNKNOWN_TYPE
}

func (m *InventoryKey) GetAvatarTemplateId() string {
	if m != nil {
		return m.AvatarTemplateId
	}
	return ""
}

type InventoryUpgrade struct {
	ItemId            ItemId               `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=SUNProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,json=upgradeType,enum=SUNProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
	AdditionalStorage int32                `protobuf:"varint,3,opt,name=additional_storage,json=additionalStorage" json:"additional_storage,omitempty"`
}

func (m *InventoryUpgrade) Reset()                    { *m = InventoryUpgrade{} }
func (m *InventoryUpgrade) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrade) ProtoMessage()               {}
func (*InventoryUpgrade) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *InventoryUpgrade) GetItemId() ItemId {
	if m != nil {
		return m.ItemId
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *InventoryUpgrade) GetUpgradeType() InventoryUpgradeType {
	if m != nil {
		return m.UpgradeType
	}
	return InventoryUpgradeType_UPGRADE_UNSET
}

func (m *InventoryUpgrade) GetAdditionalStorage() int32 {
	if m != nil {
		return m.AdditionalStorage
	}
	return 0
}

type InventoryUpgrades struct {
	InventoryUpgrades []*InventoryUpgrade `protobuf:"bytes,1,rep,name=inventory_upgrades,json=inventoryUpgrades" json:"inventory_upgrades,omitempty"`
}

func (m *InventoryUpgrades) Reset()                    { *m = InventoryUpgrades{} }
func (m *InventoryUpgrades) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrades) ProtoMessage()               {}
func (*InventoryUpgrades) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *InventoryUpgrades) GetInventoryUpgrades() []*InventoryUpgrade {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func init() {
	proto.RegisterType((*AppliedItem)(nil), "SUNProtos.Inventory.AppliedItem")
	proto.RegisterType((*AppliedItems)(nil), "SUNProtos.Inventory.AppliedItems")
	proto.RegisterType((*Candy)(nil), "SUNProtos.Inventory.Candy")
	proto.RegisterType((*EggIncubator)(nil), "SUNProtos.Inventory.EggIncubator")
	proto.RegisterType((*EggIncubators)(nil), "SUNProtos.Inventory.EggIncubators")
	proto.RegisterType((*InventoryDelta)(nil), "SUNProtos.Inventory.InventoryDelta")
	proto.RegisterType((*InventoryItem)(nil), "SUNProtos.Inventory.InventoryItem")
	proto.RegisterType((*InventoryItem_DeletedItem)(nil), "SUNProtos.Inventory.InventoryItem.DeletedItem")
	proto.RegisterType((*InventoryItemData)(nil), "SUNProtos.Inventory.InventoryItemData")
	proto.RegisterType((*InventoryKey)(nil), "SUNProtos.Inventory.InventoryKey")
	proto.RegisterType((*InventoryUpgrade)(nil), "SUNProtos.Inventory.InventoryUpgrade")
	proto.RegisterType((*InventoryUpgrades)(nil), "SUNProtos.Inventory.InventoryUpgrades")
	proto.RegisterEnum("SUNProtos.Inventory.EggIncubatorType", EggIncubatorType_name, EggIncubatorType_value)
	proto.RegisterEnum("SUNProtos.Inventory.InventoryUpgradeType", InventoryUpgradeType_name, InventoryUpgradeType_value)
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 1197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0x9c, 0x38, 0xb5, 0x8f, 0x64, 0xc7, 0x61, 0xda, 0xc1, 0x4b, 0x3b, 0xc0, 0x55, 0x91,
	0xce, 0x2b, 0x3a, 0x63, 0xf0, 0x86, 0x01, 0xbd, 0xd8, 0x0a, 0x37, 0x56, 0x3b, 0x21, 0x8d, 0xe3,
	0xd2, 0xf6, 0x0a, 0x0c, 0x03, 0x04, 0x36, 0x62, 0x0d, 0xa1, 0x96, 0xac, 0x4a, 0x74, 0x5b, 0xbf,
	0xcb, 0x1e, 0x62, 0xf7, 0xbb, 0xd8, 0xc5, 0xb0, 0x9b, 0x3d, 0xd5, 0xc0, 0x43, 0xc9, 0x96, 0x1c,
	0x35, 0xce, 0xb0, 0x9b, 0x88, 0xe7, 0x87, 0xc7, 0x1f, 0xc9, 0xef, 0xfc, 0x04, 0xf6, 0xbd, 0xe0,
	0x3d, 0x0f, 0xc4, 0x3c, 0x5a, 0x76, 0xc2, 0x68, 0x2e, 0xe6, 0xe4, 0x70, 0x34, 0x19, 0x0c, 0xe5,
	0x2a, 0xee, 0xd8, 0xa9, 0xe9, 0xe8, 0xd6, 0xca, 0xcb, 0xf1, 0x04, 0xf7, 0x95, 0xeb, 0x91, 0xce,
	0x83, 0x85, 0x1f, 0x27, 0xc2, 0x81, 0xcb, 0x04, 0x73, 0xd8, 0x7b, 0x26, 0x58, 0x94, 0xa8, 0x40,
	0xaa, 0x72, 0xe6, 0x70, 0xc6, 0x96, 0x3c, 0xca, 0xa9, 0xde, 0x2d, 0x78, 0x2c, 0x92, 0x20, 0xe6,
	0xdf, 0x1a, 0xe8, 0xbd, 0x30, 0x9c, 0x79, 0xdc, 0xb5, 0x05, 0xf7, 0xc9, 0x63, 0xb8, 0x29, 0x7f,
	0xcf, 0xf1, 0xdc, 0xa6, 0xd6, 0xd2, 0xda, 0xf5, 0x6e, 0xab, 0x53, 0x00, 0xaf, 0x23, 0x7d, 0xf1,
	0x8f, 0xed, 0xd2, 0x3d, 0x0f, 0xbf, 0xe4, 0x09, 0x54, 0x71, 0xab, 0x58, 0x86, 0xbc, 0x59, 0xc2,
	0xcd, 0xe6, 0xd5, 0x9b, 0xc7, 0xcb, 0x90, 0xd3, 0x8a, 0x97, 0xac, 0xc8, 0x1d, 0xa8, 0xf2, 0x8f,
	0xa1, 0x17, 0x71, 0xc7, 0x8f, 0x9b, 0x3b, 0x2d, 0xad, 0xbd, 0x43, 0x2b, 0x4a, 0x71, 0x16, 0x93,
	0x2f, 0x00, 0x98, 0xc2, 0x29, 0xad, 0xbb, 0x68, 0xad, 0x26, 0x9a, 0xb3, 0xd8, 0xec, 0x83, 0x91,
	0x39, 0x46, 0x4c, 0xbe, 0x83, 0x5d, 0x19, 0xb7, 0xb9, 0xdb, 0xda, 0x69, 0xeb, 0x9f, 0x38, 0x44,
	0x66, 0x03, 0x45, 0x6f, 0xf3, 0x57, 0x28, 0x9f, 0xb0, 0xc0, 0x5d, 0x92, 0x1f, 0xa0, 0xfa, 0x86,
	0xf9, 0xde, 0x6c, 0x59, 0x7c, 0x11, 0x16, 0x3e, 0xc3, 0x70, 0xfe, 0x96, 0xfb, 0xf3, 0xe0, 0x19,
	0x3a, 0xda, 0x2e, 0xad, 0xbc, 0x49, 0x56, 0xe4, 0x16, 0x94, 0x2f, 0x64, 0x1c, 0xbc, 0x86, 0x32,
	0x55, 0x82, 0xf9, 0x67, 0x09, 0x0c, 0x6b, 0x3a, 0xb5, 0x83, 0x8b, 0xc5, 0x6b, 0x26, 0xe6, 0x11,
	0xa9, 0x43, 0x29, 0x09, 0x5f, 0xa5, 0x25, 0xcf, 0xcd, 0x5e, 0x7e, 0xe9, 0x3f, 0x5e, 0xfe, 0x0b,
	0xa8, 0x7b, 0x69, 0x5c, 0xf5, 0x02, 0x3b, 0x18, 0xe1, 0xb8, 0x30, 0x42, 0x16, 0x05, 0x3e, 0x42,
	0xcd, 0xcb, 0x8a, 0xe4, 0x18, 0xea, 0x8b, 0x98, 0xc7, 0x4e, 0xc4, 0x7d, 0xe6, 0x05, 0x5e, 0x30,
	0xc5, 0x0b, 0x2f, 0xd3, 0x9a, 0xd4, 0xd2, 0x54, 0x29, 0xdf, 0x24, 0x54, 0x77, 0x20, 0x21, 0x97,
	0x5b, 0x5a, 0x7b, 0x97, 0x56, 0x13, 0x8d, 0xed, 0x92, 0x07, 0xb0, 0x1f, 0x0b, 0x16, 0x09, 0xe7,
	0xad, 0xef, 0x7c, 0x60, 0xb3, 0xb7, 0xdc, 0x6d, 0xee, 0xb5, 0xb4, 0xb6, 0x46, 0x6b, 0xa8, 0x3e,
	0xf5, 0x5f, 0xa1, 0x92, 0xb4, 0xa1, 0x21, 0x58, 0x34, 0xe5, 0x59, 0xc7, 0x9b, 0xe8, 0x58, 0x57,
	0xfa, 0xd4, 0xd3, 0x7c, 0x05, 0xb5, 0x2c, 0xf4, 0x98, 0x3c, 0x83, 0x1a, 0x9f, 0x4e, 0x9d, 0x15,
	0xfa, 0xa6, 0x86, 0xef, 0x7d, 0x6f, 0xeb, 0xa9, 0xa9, 0xc1, 0x33, 0x92, 0xf9, 0x87, 0x06, 0xf5,
	0x95, 0x63, 0x9f, 0xcf, 0x04, 0x23, 0x5d, 0xb8, 0x3d, 0x8f, 0xbc, 0xa9, 0x17, 0xb0, 0x99, 0x23,
	0x3c, 0x9f, 0xc7, 0x82, 0xf9, 0xa1, 0xe4, 0x9e, 0x86, 0xdc, 0x3b, 0x4c, 0x8d, 0xe3, 0xd4, 0x76,
	0x16, 0xcb, 0x93, 0x04, 0xfc, 0x43, 0xde, 0xbd, 0x84, 0xee, 0xf5, 0x80, 0x7f, 0xc8, 0x7a, 0x9e,
	0x66, 0xea, 0x00, 0x66, 0xb8, 0x64, 0xbc, 0x84, 0xfe, 0x89, 0x94, 0x49, 0x57, 0x48, 0xd6, 0xba,
	0x97, 0x15, 0x63, 0xf3, 0xb7, 0x12, 0xd4, 0x72, 0x1e, 0x12, 0xbc, 0x3f, 0x77, 0xbd, 0x37, 0x32,
	0x5d, 0x8a, 0xc0, 0xa7, 0xc6, 0x2c, 0xa4, 0x97, 0x60, 0xb8, 0x7c, 0xc6, 0x05, 0x77, 0x11, 0x10,
	0x02, 0xd7, 0xbb, 0x9d, 0xed, 0x78, 0x3a, 0x7d, 0xb5, 0x0d, 0xb1, 0xe9, 0xee, 0x5a, 0x20, 0x3f,
	0xc3, 0x61, 0xfe, 0x94, 0x8e, 0xac, 0x40, 0x48, 0x4d, 0xbd, 0xfb, 0x60, 0x7b, 0xe4, 0x3e, 0x13,
	0x8c, 0x1e, 0x78, 0x9b, 0xaa, 0xa3, 0x47, 0xa0, 0x67, 0x7e, 0x73, 0x83, 0x87, 0xf2, 0x88, 0x7b,
	0x19, 0x1e, 0x9a, 0x7f, 0xed, 0xc1, 0xc1, 0xa5, 0xb0, 0xe4, 0x47, 0x30, 0xd2, 0x4d, 0x08, 0x4a,
	0x43, 0x50, 0x77, 0x32, 0xa0, 0xa4, 0x5b, 0x9a, 0xe4, 0x88, 0x44, 0x0f, 0xd7, 0x02, 0xf9, 0x3e,
	0xa9, 0x30, 0xea, 0x9a, 0xb6, 0x54, 0x3a, 0xdc, 0x8e, 0xfe, 0xa4, 0x07, 0x35, 0x19, 0xc6, 0xe5,
	0x1f, 0x1d, 0x1e, 0x88, 0x68, 0x99, 0xdc, 0xc6, 0xdd, 0xa2, 0x1f, 0x76, 0xf9, 0x47, 0x4b, 0xfa,
	0x50, 0x23, 0xcc, 0x48, 0xc4, 0x02, 0x43, 0xd5, 0x75, 0x27, 0x16, 0x4c, 0xa8, 0x6a, 0x98, 0x87,
	0xa0, 0x22, 0xa8, 0xda, 0xaf, 0x3e, 0x23, 0xe9, 0x49, 0xf5, 0x70, 0x2d, 0x90, 0x01, 0xec, 0x27,
	0x61, 0x2e, 0x16, 0x51, 0xc4, 0x83, 0x8b, 0x25, 0xe6, 0xb0, 0x9e, 0x2b, 0x1a, 0x97, 0x23, 0x9d,
	0x24, 0xce, 0xb4, 0x1e, 0xe6, 0x64, 0xf2, 0x13, 0xd4, 0xd2, 0x78, 0xcc, 0xe7, 0x11, 0xc3, 0x6c,
	0xd7, 0xbb, 0xf7, 0xaf, 0x8e, 0x86, 0xae, 0x34, 0x39, 0x90, 0x92, 0xc8, 0x04, 0xc8, 0x9a, 0x37,
	0x8b, 0x70, 0x1a, 0x31, 0x97, 0xc7, 0x58, 0x13, 0xb6, 0xd2, 0x66, 0x92, 0x78, 0x67, 0x68, 0x93,
	0xaa, 0x64, 0xb5, 0x48, 0x7b, 0x88, 0x4a, 0xb9, 0x0a, 0x46, 0xbc, 0xb7, 0xad, 0x3b, 0xc4, 0xd4,
	0x60, 0xd9, 0xe6, 0x62, 0x43, 0x3d, 0x57, 0x75, 0xe2, 0x66, 0xf5, 0x0a, 0x12, 0xe4, 0x2a, 0x16,
	0xad, 0xf1, 0x5c, 0x01, 0xfb, 0x26, 0xed, 0x14, 0x80, 0x11, 0x8e, 0x0a, 0x23, 0x60, 0x4f, 0x4a,
	0xba, 0x08, 0xe9, 0x42, 0x19, 0x3b, 0x78, 0x53, 0x2f, 0xe6, 0xcd, 0x4b, 0xd5, 0xde, 0xf1, 0x43,
	0x95, 0x2b, 0x79, 0x0a, 0xba, 0x9a, 0x13, 0x54, 0x66, 0x1b, 0x97, 0x8e, 0x8d, 0x3b, 0x7b, 0x6a,
	0x94, 0x50, 0x1f, 0x4c, 0x66, 0x60, 0xab, 0xb5, 0xf9, 0xfb, 0x2e, 0x18, 0x2b, 0x48, 0xa7, 0x7c,
	0xb9, 0x25, 0xeb, 0x56, 0x1d, 0xf8, 0xba, 0x9d, 0x4c, 0x65, 0x47, 0x1b, 0x1a, 0xb9, 0xec, 0x90,
	0xa1, 0x77, 0xb0, 0xf7, 0xd4, 0xb3, 0x29, 0x60, 0xbb, 0xe4, 0x5e, 0x41, 0x12, 0x54, 0xf2, 0x04,
	0xff, 0xb2, 0x98, 0xe0, 0x95, 0x4b, 0xcc, 0xbd, 0x5f, 0xc4, 0xdc, 0xca, 0x06, 0x29, 0xbf, 0xfe,
	0x24, 0x29, 0x2b, 0x45, 0x64, 0xbb, 0x5f, 0x44, 0xb6, 0xca, 0x06, 0x93, 0x8e, 0x0b, 0x99, 0x54,
	0xd9, 0x64, 0xc9, 0x0b, 0x38, 0x48, 0xaf, 0x7a, 0x3d, 0x96, 0xc0, 0x35, 0xc7, 0x92, 0xfd, 0x30,
	0xaf, 0x20, 0x8f, 0x01, 0x90, 0x16, 0x6a, 0x4e, 0xd0, 0x31, 0xcc, 0xd1, 0xa5, 0x30, 0x48, 0x20,
	0x1c, 0x0e, 0xaa, 0xef, 0xd2, 0x25, 0x79, 0x04, 0x24, 0x21, 0x92, 0xe0, 0x7e, 0x38, 0x63, 0x82,
	0x4b, 0x24, 0x06, 0x4e, 0x30, 0x0d, 0x65, 0x19, 0x27, 0x06, 0xdb, 0x35, 0xff, 0xd1, 0xa0, 0xb1,
	0x99, 0x98, 0xff, 0x67, 0xc2, 0x7c, 0x01, 0x46, 0x72, 0xef, 0xd9, 0x21, 0xf3, 0xab, 0x6b, 0x15,
	0x04, 0x3c, 0x89, 0xbe, 0x58, 0x0b, 0xf2, 0x3d, 0x99, 0xeb, 0x7a, 0xc2, 0x9b, 0xcb, 0x16, 0x1f,
	0x8b, 0x79, 0xc4, 0xa6, 0x3c, 0x21, 0xdb, 0xc1, 0xda, 0x32, 0x52, 0x06, 0xd3, 0xcb, 0x34, 0x91,
	0xd5, 0x23, 0x8f, 0x0b, 0x39, 0xa1, 0x86, 0x90, 0xe3, 0x6b, 0xe1, 0x2a, 0xa0, 0xce, 0xc3, 0x27,
	0xd0, 0xd8, 0x9c, 0xd0, 0xc8, 0x21, 0xec, 0xdb, 0x83, 0x93, 0xc9, 0xd3, 0xde, 0xf8, 0x9c, 0x3a,
	0x93, 0xc1, 0xc8, 0x1a, 0x37, 0x6e, 0x90, 0xcf, 0x80, 0xac, 0x95, 0x7d, 0x7b, 0x34, 0xee, 0x0d,
	0x4e, 0xac, 0x86, 0xf6, 0xf0, 0x35, 0xdc, 0x2a, 0x3a, 0x3f, 0x39, 0x80, 0xda, 0x64, 0xf8, 0x9c,
	0xf6, 0xfa, 0xd6, 0x2a, 0xc4, 0xe7, 0x70, 0xdb, 0x1e, 0x9c, 0x50, 0xab, 0x37, 0xb2, 0x1c, 0x7b,
	0x6c, 0x9d, 0x39, 0xa3, 0xf1, 0x39, 0xed, 0x3d, 0xb7, 0x1a, 0x1a, 0xb9, 0x0b, 0xcd, 0x95, 0x69,
	0x78, 0x7e, 0x6a, 0x9d, 0x9d, 0x0f, 0x56, 0xd6, 0xd2, 0xd3, 0xca, 0x2f, 0x7b, 0xf8, 0x2f, 0x44,
	0x3c, 0xbc, 0x31, 0xd4, 0x86, 0xa5, 0xe1, 0xce, 0x70, 0x77, 0x58, 0x7e, 0xad, 0x74, 0xdf, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x70, 0x64, 0x21, 0xa5, 0xe6, 0x0c, 0x00, 0x00,
}
