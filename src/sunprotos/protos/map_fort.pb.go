// Code generated by protoc-gen-go.
// source: map_fort.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type FortRenderingType int32

const (
	FortRenderingType_DEFAULT       FortRenderingType = 0
	FortRenderingType_INTERNAL_TEST FortRenderingType = 1
)

var FortRenderingType_name = map[int32]string{
	0: "DEFAULT",
	1: "INTERNAL_TEST",
}
var FortRenderingType_value = map[string]int32{
	"DEFAULT":       0,
	"INTERNAL_TEST": 1,
}

func (x FortRenderingType) String() string {
	return proto.EnumName(FortRenderingType_name, int32(x))
}
func (FortRenderingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type FortSponsor int32

const (
	FortSponsor_UNSET_SPONSOR FortSponsor = 0
	FortSponsor_MCDONALDS     FortSponsor = 1
	FortSponsor_POKEMON_STORE FortSponsor = 2
	FortSponsor_TOHO          FortSponsor = 3
	FortSponsor_SOFTBANK      FortSponsor = 4
	FortSponsor_GLOBE         FortSponsor = 5
	FortSponsor_SPATULA       FortSponsor = 6
	FortSponsor_THERMOMETER   FortSponsor = 7
	FortSponsor_KNIFE         FortSponsor = 8
	FortSponsor_GRILL         FortSponsor = 9
	FortSponsor_SMOKER        FortSponsor = 10
	FortSponsor_PAN           FortSponsor = 11
	FortSponsor_BBQ           FortSponsor = 12
	FortSponsor_FRYER         FortSponsor = 13
	FortSponsor_STEAMER       FortSponsor = 14
	FortSponsor_HOOD          FortSponsor = 15
	FortSponsor_SLOWCOOKER    FortSponsor = 16
	FortSponsor_MIXER         FortSponsor = 17
	FortSponsor_SCOOPER       FortSponsor = 18
	FortSponsor_MUFFINTIN     FortSponsor = 19
	FortSponsor_SALAMANDER    FortSponsor = 20
	FortSponsor_PLANCHA       FortSponsor = 21
)

var FortSponsor_name = map[int32]string{
	0:  "UNSET_SPONSOR",
	1:  "MCDONALDS",
	2:  "POKEMON_STORE",
	3:  "TOHO",
	4:  "SOFTBANK",
	5:  "GLOBE",
	6:  "SPATULA",
	7:  "THERMOMETER",
	8:  "KNIFE",
	9:  "GRILL",
	10: "SMOKER",
	11: "PAN",
	12: "BBQ",
	13: "FRYER",
	14: "STEAMER",
	15: "HOOD",
	16: "SLOWCOOKER",
	17: "MIXER",
	18: "SCOOPER",
	19: "MUFFINTIN",
	20: "SALAMANDER",
	21: "PLANCHA",
}
var FortSponsor_value = map[string]int32{
	"UNSET_SPONSOR": 0,
	"MCDONALDS":     1,
	"POKEMON_STORE": 2,
	"TOHO":          3,
	"SOFTBANK":      4,
	"GLOBE":         5,
	"SPATULA":       6,
	"THERMOMETER":   7,
	"KNIFE":         8,
	"GRILL":         9,
	"SMOKER":        10,
	"PAN":           11,
	"BBQ":           12,
	"FRYER":         13,
	"STEAMER":       14,
	"HOOD":          15,
	"SLOWCOOKER":    16,
	"MIXER":         17,
	"SCOOPER":       18,
	"MUFFINTIN":     19,
	"SALAMANDER":    20,
	"PLANCHA":       21,
}

func (x FortSponsor) String() string {
	return proto.EnumName(FortSponsor_name, int32(x))
}
func (FortSponsor) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

type FortType int32

const (
	FortType_GYM        FortType = 0
	FortType_CHECKPOINT FortType = 1
)

var FortType_name = map[int32]string{
	0: "GYM",
	1: "CHECKPOINT",
}
var FortType_value = map[string]int32{
	"GYM":        0,
	"CHECKPOINT": 1,
}

func (x FortType) String() string {
	return proto.EnumName(FortType_name, int32(x))
}
func (FortType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

type FortData struct {
	Id                          string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastModifiedTimestampMs     int64             `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                    float64           `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude                   float64           `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	OwnedByTeam                 TeamColor         `protobuf:"varint,5,opt,name=owned_by_team,json=ownedByTeam,enum=SUNProtos.Enums.TeamColor" json:"owned_by_team,omitempty"`
	GuardPokemonId              PokemonId         `protobuf:"varint,6,opt,name=guard_pokemon_id,json=guardPokemonId,enum=SUNProtos.Enums.PokemonId" json:"guard_pokemon_id,omitempty"`
	GuardPokemonCp              int32             `protobuf:"varint,7,opt,name=guard_pokemon_cp,json=guardPokemonCp" json:"guard_pokemon_cp,omitempty"`
	Enabled                     bool              `protobuf:"varint,8,opt,name=enabled" json:"enabled,omitempty"`
	Type                        FortType          `protobuf:"varint,9,opt,name=type,enum=SUNProtos.Map.Fort.FortType" json:"type,omitempty"`
	GymPoints                   int64             `protobuf:"varint,10,opt,name=gym_points,json=gymPoints" json:"gym_points,omitempty"`
	IsInBattle                  bool              `protobuf:"varint,11,opt,name=is_in_battle,json=isInBattle" json:"is_in_battle,omitempty"`
	ActiveFortModifier          []ItemId          `protobuf:"varint,12,rep,packed,name=active_fort_modifier,json=activeFortModifier,enum=SUNProtos.Inventory.Item.ItemId" json:"active_fort_modifier,omitempty"`
	LureInfo                    *FortLureInfo     `protobuf:"bytes,13,opt,name=lure_info,json=lureInfo" json:"lure_info,omitempty"`
	CooldownCompleteTimestampMs int64             `protobuf:"varint,14,opt,name=cooldown_complete_timestamp_ms,json=cooldownCompleteTimestampMs" json:"cooldown_complete_timestamp_ms,omitempty"`
	Sponsor                     FortSponsor       `protobuf:"varint,15,opt,name=sponsor,enum=SUNProtos.Map.Fort.FortSponsor" json:"sponsor,omitempty"`
	RenderingType               FortRenderingType `protobuf:"varint,16,opt,name=rendering_type,json=renderingType,enum=SUNProtos.Map.Fort.FortRenderingType" json:"rendering_type,omitempty"`
	DeployLockoutEndMs          int64             `protobuf:"varint,17,opt,name=deploy_lockout_end_ms,json=deployLockoutEndMs" json:"deploy_lockout_end_ms,omitempty"`
	GuardPokemonDisplay         *PokemonDisplay   `protobuf:"bytes,18,opt,name=guard_pokemon_display,json=guardPokemonDisplay" json:"guard_pokemon_display,omitempty"`
	Closed                      bool              `protobuf:"varint,19,opt,name=closed" json:"closed,omitempty"`
}

func (m *FortData) Reset()                    { *m = FortData{} }
func (m *FortData) String() string            { return proto.CompactTextString(m) }
func (*FortData) ProtoMessage()               {}
func (*FortData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *FortData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FortData) GetLastModifiedTimestampMs() int64 {
	if m != nil {
		return m.LastModifiedTimestampMs
	}
	return 0
}

func (m *FortData) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *FortData) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *FortData) GetOwnedByTeam() TeamColor {
	if m != nil {
		return m.OwnedByTeam
	}
	return TeamColor_NEUTRAL
}

func (m *FortData) GetGuardPokemonId() PokemonId {
	if m != nil {
		return m.GuardPokemonId
	}
	return PokemonId_MISSINGNO
}

func (m *FortData) GetGuardPokemonCp() int32 {
	if m != nil {
		return m.GuardPokemonCp
	}
	return 0
}

func (m *FortData) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *FortData) GetType() FortType {
	if m != nil {
		return m.Type
	}
	return FortType_GYM
}

func (m *FortData) GetGymPoints() int64 {
	if m != nil {
		return m.GymPoints
	}
	return 0
}

func (m *FortData) GetIsInBattle() bool {
	if m != nil {
		return m.IsInBattle
	}
	return false
}

func (m *FortData) GetActiveFortModifier() []ItemId {
	if m != nil {
		return m.ActiveFortModifier
	}
	return nil
}

func (m *FortData) GetLureInfo() *FortLureInfo {
	if m != nil {
		return m.LureInfo
	}
	return nil
}

func (m *FortData) GetCooldownCompleteTimestampMs() int64 {
	if m != nil {
		return m.CooldownCompleteTimestampMs
	}
	return 0
}

func (m *FortData) GetSponsor() FortSponsor {
	if m != nil {
		return m.Sponsor
	}
	return FortSponsor_UNSET_SPONSOR
}

func (m *FortData) GetRenderingType() FortRenderingType {
	if m != nil {
		return m.RenderingType
	}
	return FortRenderingType_DEFAULT
}

func (m *FortData) GetDeployLockoutEndMs() int64 {
	if m != nil {
		return m.DeployLockoutEndMs
	}
	return 0
}

func (m *FortData) GetGuardPokemonDisplay() *PokemonDisplay {
	if m != nil {
		return m.GuardPokemonDisplay
	}
	return nil
}

func (m *FortData) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

type FortLureInfo struct {
	FortId                 string    `protobuf:"bytes,1,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	EncounterId            uint64    `protobuf:"fixed64,2,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	ActivePokemonId        PokemonId `protobuf:"varint,3,opt,name=active_pokemon_id,json=activePokemonId,enum=SUNProtos.Enums.PokemonId" json:"active_pokemon_id,omitempty"`
	LureExpiresTimestampMs int64     `protobuf:"varint,4,opt,name=lure_expires_timestamp_ms,json=lureExpiresTimestampMs" json:"lure_expires_timestamp_ms,omitempty"`
}

func (m *FortLureInfo) Reset()                    { *m = FortLureInfo{} }
func (m *FortLureInfo) String() string            { return proto.CompactTextString(m) }
func (*FortLureInfo) ProtoMessage()               {}
func (*FortLureInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *FortLureInfo) GetFortId() string {
	if m != nil {
		return m.FortId
	}
	return ""
}

func (m *FortLureInfo) GetEncounterId() uint64 {
	if m != nil {
		return m.EncounterId
	}
	return 0
}

func (m *FortLureInfo) GetActivePokemonId() PokemonId {
	if m != nil {
		return m.ActivePokemonId
	}
	return PokemonId_MISSINGNO
}

func (m *FortLureInfo) GetLureExpiresTimestampMs() int64 {
	if m != nil {
		return m.LureExpiresTimestampMs
	}
	return 0
}

type FortModifier struct {
	ItemId                 ItemId `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=SUNProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ExpirationTimestampMs  int64  `protobuf:"varint,2,opt,name=expiration_timestamp_ms,json=expirationTimestampMs" json:"expiration_timestamp_ms,omitempty"`
	DeployerPlayerCodename string `protobuf:"bytes,3,opt,name=deployer_player_codename,json=deployerPlayerCodename" json:"deployer_player_codename,omitempty"`
}

func (m *FortModifier) Reset()                    { *m = FortModifier{} }
func (m *FortModifier) String() string            { return proto.CompactTextString(m) }
func (*FortModifier) ProtoMessage()               {}
func (*FortModifier) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *FortModifier) GetItemId() ItemId {
	if m != nil {
		return m.ItemId
	}
	return ItemId_ITEM_UNKNOWN
}

func (m *FortModifier) GetExpirationTimestampMs() int64 {
	if m != nil {
		return m.ExpirationTimestampMs
	}
	return 0
}

func (m *FortModifier) GetDeployerPlayerCodename() string {
	if m != nil {
		return m.DeployerPlayerCodename
	}
	return ""
}

type FortSummary struct {
	FortSummaryId           string  `protobuf:"bytes,1,opt,name=fort_summary_id,json=fortSummaryId" json:"fort_summary_id,omitempty"`
	LastModifiedTimestampMs int64   `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *FortSummary) Reset()                    { *m = FortSummary{} }
func (m *FortSummary) String() string            { return proto.CompactTextString(m) }
func (*FortSummary) ProtoMessage()               {}
func (*FortSummary) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *FortSummary) GetFortSummaryId() string {
	if m != nil {
		return m.FortSummaryId
	}
	return ""
}

func (m *FortSummary) GetLastModifiedTimestampMs() int64 {
	if m != nil {
		return m.LastModifiedTimestampMs
	}
	return 0
}

func (m *FortSummary) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *FortSummary) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*FortData)(nil), "SUNProtos.Map.Fort.FortData")
	proto.RegisterType((*FortLureInfo)(nil), "SUNProtos.Map.Fort.FortLureInfo")
	proto.RegisterType((*FortModifier)(nil), "SUNProtos.Map.Fort.FortModifier")
	proto.RegisterType((*FortSummary)(nil), "SUNProtos.Map.Fort.FortSummary")
	proto.RegisterEnum("SUNProtos.Map.Fort.FortRenderingType", FortRenderingType_name, FortRenderingType_value)
	proto.RegisterEnum("SUNProtos.Map.Fort.FortSponsor", FortSponsor_name, FortSponsor_value)
	proto.RegisterEnum("SUNProtos.Map.Fort.FortType", FortType_name, FortType_value)
}

func init() { proto.RegisterFile("map_fort.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 1059 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0xae, 0x93, 0x36, 0x2f, 0x93, 0x97, 0x6e, 0xb6, 0x6f, 0xa6, 0x1c, 0x47, 0x28, 0x02, 0x45,
	0xfd, 0x10, 0xc1, 0x9d, 0x84, 0xa8, 0x10, 0x48, 0x4e, 0xe2, 0x5c, 0xad, 0xfa, 0x8d, 0xb5, 0x2b,
	0x38, 0xbe, 0x58, 0x6e, 0xbc, 0xad, 0xac, 0xb3, 0xbd, 0x96, 0xed, 0xdc, 0x91, 0xff, 0x04, 0xbf,
	0x80, 0xdf, 0x00, 0xbf, 0x09, 0xed, 0xda, 0x69, 0x12, 0x8e, 0x0a, 0xbe, 0xdd, 0x97, 0xc4, 0x33,
	0xcf, 0x33, 0xe3, 0x99, 0x9d, 0x67, 0xd6, 0xd0, 0x8f, 0xfd, 0xd4, 0xbb, 0x67, 0x59, 0x31, 0x4e,
	0x33, 0x56, 0x30, 0x8c, 0x9d, 0x5b, 0xd3, 0xe6, 0x4f, 0xf9, 0xd8, 0xf0, 0xd3, 0xf1, 0x9c, 0x65,
	0xc5, 0x39, 0x04, 0x7e, 0xe1, 0x97, 0xf8, 0x79, 0x87, 0x26, 0xcb, 0x38, 0xaf, 0x8c, 0xe3, 0x30,
	0x79, 0x4b, 0x93, 0x82, 0x65, 0x2b, 0x2f, 0x2c, 0x68, 0x5c, 0x7a, 0x2f, 0xfe, 0x6a, 0x42, 0x8b,
	0xc7, 0xcd, 0xfc, 0xc2, 0xc7, 0x7d, 0xa8, 0x85, 0x81, 0x2c, 0x0d, 0xa5, 0x51, 0x9b, 0xd4, 0xc2,
	0x00, 0x7f, 0x07, 0xe7, 0x91, 0x9f, 0x17, 0x5e, 0xcc, 0x82, 0xf0, 0x3e, 0xa4, 0x81, 0x57, 0x84,
	0x31, 0xcd, 0x0b, 0x3f, 0x4e, 0xbd, 0x38, 0x97, 0x6b, 0x43, 0x69, 0x54, 0x27, 0x67, 0x9c, 0x61,
	0x54, 0x04, 0x77, 0x8d, 0x1b, 0x39, 0x3e, 0x87, 0x56, 0xe4, 0x17, 0x61, 0xb1, 0x0c, 0xa8, 0x5c,
	0x1f, 0x4a, 0x23, 0x89, 0x3c, 0xda, 0xf8, 0x19, 0xb4, 0x23, 0x96, 0x3c, 0x94, 0xe0, 0xbe, 0x00,
	0x37, 0x0e, 0xfc, 0x03, 0xf4, 0xd8, 0xbb, 0x84, 0x06, 0xde, 0xdd, 0xca, 0x2b, 0xa8, 0x1f, 0xcb,
	0x07, 0x43, 0x69, 0xd4, 0x7f, 0x71, 0x3e, 0xde, 0xb4, 0xab, 0x8a, 0xc6, 0x5c, 0xea, 0xc7, 0x53,
	0x16, 0xb1, 0x8c, 0x74, 0x44, 0xc0, 0x64, 0xc5, 0x3d, 0x78, 0x06, 0xe8, 0x61, 0xe9, 0x67, 0x81,
	0x97, 0xb2, 0x37, 0x34, 0x66, 0x89, 0x17, 0x06, 0x72, 0xe3, 0x89, 0x14, 0x76, 0x49, 0xd1, 0x02,
	0xd2, 0x17, 0x31, 0x8f, 0x36, 0x1e, 0xfd, 0x33, 0xcb, 0x22, 0x95, 0x9b, 0x43, 0x69, 0x74, 0xb0,
	0xcb, 0x9c, 0xa6, 0x58, 0x86, 0x26, 0x4d, 0xfc, 0xbb, 0x88, 0x06, 0x72, 0x6b, 0x28, 0x8d, 0x5a,
	0x64, 0x6d, 0xe2, 0xaf, 0x60, 0xbf, 0x58, 0xa5, 0x54, 0x6e, 0x8b, 0xb7, 0x3f, 0x1b, 0xbf, 0x3f,
	0x2f, 0xf1, 0xe3, 0xae, 0x52, 0x4a, 0x04, 0x13, 0x7f, 0x02, 0xf0, 0xb0, 0x8a, 0xbd, 0x94, 0x85,
	0x49, 0x91, 0xcb, 0x20, 0x8e, 0xb8, 0xfd, 0xb0, 0x8a, 0x6d, 0xe1, 0xc0, 0x43, 0xe8, 0x86, 0xb9,
	0x17, 0x26, 0xde, 0x9d, 0x5f, 0x14, 0x11, 0x95, 0x3b, 0xe2, 0x7d, 0x10, 0xe6, 0x5a, 0x32, 0x11,
	0x1e, 0x4c, 0xe0, 0xd8, 0x5f, 0x14, 0xe1, 0x5b, 0x2a, 0x84, 0xb2, 0x1e, 0x5d, 0x26, 0x77, 0x87,
	0xf5, 0x51, 0xff, 0xc5, 0x70, 0xab, 0x04, 0x6d, 0xad, 0x87, 0xb1, 0xc6, 0xf5, 0xc0, 0x7f, 0xb4,
	0x80, 0xe0, 0x32, 0x9a, 0x97, 0x55, 0x4d, 0x35, 0xc3, 0xdf, 0x43, 0x3b, 0x5a, 0x66, 0xd4, 0x0b,
	0x93, 0x7b, 0x26, 0xf7, 0x86, 0xd2, 0xa8, 0xb3, 0x93, 0x68, 0xa7, 0x17, 0x7d, 0x99, 0x51, 0x2d,
	0xb9, 0x67, 0xa4, 0x15, 0x55, 0x4f, 0x78, 0x0a, 0xcf, 0x17, 0x8c, 0x45, 0x01, 0x7b, 0x97, 0x78,
	0x0b, 0x16, 0xa7, 0x11, 0x2d, 0xe8, 0xae, 0x94, 0xfa, 0xa2, 0xcf, 0x8f, 0xd7, 0xac, 0x69, 0x45,
	0xda, 0x96, 0xd3, 0x15, 0x34, 0xf3, 0x94, 0x25, 0x39, 0xcb, 0xe4, 0x43, 0x71, 0x9a, 0x9f, 0x3e,
	0x55, 0x81, 0x53, 0xd2, 0xc8, 0x9a, 0x8f, 0x75, 0xe8, 0x67, 0x34, 0x09, 0x68, 0x16, 0x26, 0x0f,
	0x9e, 0x98, 0x07, 0x12, 0x19, 0xbe, 0x78, 0x2a, 0x03, 0x59, 0xb3, 0xc5, 0x60, 0x7a, 0xd9, 0xb6,
	0x89, 0xbf, 0x86, 0x93, 0x80, 0xa6, 0x11, 0x5b, 0x79, 0x11, 0x5b, 0xbc, 0x61, 0xcb, 0xc2, 0xa3,
	0x49, 0xc0, 0x9b, 0x18, 0x88, 0x26, 0x70, 0x09, 0xea, 0x25, 0xa6, 0x26, 0x81, 0x91, 0x63, 0x02,
	0x27, 0xbb, 0x52, 0x0a, 0xc2, 0x3c, 0x8d, 0xfc, 0x95, 0x8c, 0xc5, 0x59, 0x3e, 0xdf, 0xaa, 0x83,
	0xef, 0xe1, 0x5a, 0x94, 0xb3, 0x92, 0x45, 0x8e, 0xb6, 0xf5, 0x56, 0x39, 0xf1, 0x29, 0x34, 0x16,
	0x11, 0xcb, 0x69, 0x20, 0x1f, 0x09, 0x0d, 0x54, 0xd6, 0xc5, 0x9f, 0x12, 0x74, 0xb7, 0xe7, 0x80,
	0xcf, 0xa0, 0x29, 0x94, 0xf0, 0xb8, 0xd9, 0x0d, 0x6e, 0x6a, 0x01, 0xfe, 0x0c, 0xba, 0x34, 0x59,
	0xb0, 0x65, 0x52, 0xd0, 0x8c, 0xa3, 0x7c, 0x9f, 0x1b, 0xa4, 0xf3, 0xe8, 0xd3, 0x02, 0x3c, 0x87,
	0x41, 0x25, 0xa6, 0xad, 0x55, 0xaa, 0xff, 0xe7, 0x2a, 0x1d, 0x96, 0x41, 0x9b, 0x5d, 0xba, 0x82,
	0x8f, 0x84, 0x80, 0xe8, 0xaf, 0x69, 0x98, 0xd1, 0x7c, 0x77, 0xf8, 0xfb, 0xe2, 0xdc, 0x4e, 0x39,
	0x41, 0x2d, 0xf1, 0xad, 0xb9, 0x5f, 0xfc, 0x51, 0xf5, 0xf3, 0x28, 0xc6, 0x2b, 0x68, 0xf2, 0xfb,
	0x6b, 0xdd, 0xcf, 0xff, 0xd1, 0x74, 0x23, 0x14, 0xff, 0xf8, 0x1b, 0x38, 0x13, 0x15, 0xf8, 0x45,
	0xc8, 0x92, 0x7f, 0xbb, 0xcc, 0x4e, 0x36, 0xf0, 0xb6, 0xf6, 0xbe, 0x05, 0xb9, 0x9c, 0x2a, 0xcd,
	0x3c, 0x7e, 0xf8, 0x34, 0xf3, 0x16, 0x2c, 0xa0, 0x89, 0x1f, 0x97, 0x57, 0x5b, 0x9b, 0x9c, 0xae,
	0x71, 0x5b, 0xc0, 0xd3, 0x0a, 0xbd, 0xf8, 0x4d, 0x82, 0x8e, 0xd0, 0xe4, 0x32, 0x8e, 0xfd, 0x6c,
	0x85, 0xbf, 0x84, 0x43, 0x31, 0x8c, 0xbc, 0xb4, 0x37, 0x43, 0xe9, 0xdd, 0x6f, 0x58, 0xda, 0x87,
	0xba, 0x79, 0x2f, 0x5f, 0xc2, 0xe0, 0x3d, 0xfd, 0xe3, 0x0e, 0x34, 0x67, 0xea, 0x5c, 0xb9, 0xd5,
	0x5d, 0xb4, 0x87, 0x07, 0xd0, 0xd3, 0x4c, 0x57, 0x25, 0xa6, 0xa2, 0x7b, 0xae, 0xea, 0xb8, 0x48,
	0xba, 0xfc, 0xbd, 0x56, 0xf5, 0x58, 0xad, 0xdb, 0x00, 0x7a, 0xb7, 0xa6, 0xa3, 0xba, 0x9e, 0x63,
	0x5b, 0xa6, 0x63, 0x11, 0xb4, 0x87, 0x7b, 0xd0, 0x36, 0xa6, 0x33, 0xcb, 0x54, 0xf4, 0x99, 0x83,
	0x24, 0xce, 0xb0, 0xad, 0x1b, 0xd5, 0xb0, 0x4c, 0xcf, 0x71, 0x2d, 0xa2, 0xa2, 0x1a, 0x6e, 0xc1,
	0xbe, 0x6b, 0x5d, 0x5b, 0xa8, 0x8e, 0xbb, 0xd0, 0x72, 0xac, 0xb9, 0x3b, 0x51, 0xcc, 0x1b, 0xb4,
	0x8f, 0xdb, 0x70, 0xf0, 0x4a, 0xb7, 0x26, 0x2a, 0x3a, 0xe0, 0x75, 0x38, 0xb6, 0xe2, 0xde, 0xea,
	0x0a, 0x6a, 0xe0, 0x43, 0xe8, 0xb8, 0xd7, 0x2a, 0x31, 0x2c, 0x43, 0x75, 0x55, 0x82, 0x9a, 0x9c,
	0x78, 0x63, 0x6a, 0x73, 0x15, 0xb5, 0x44, 0x0c, 0xd1, 0x74, 0x1d, 0xb5, 0x31, 0x40, 0xc3, 0x31,
	0xac, 0x1b, 0x95, 0x20, 0xc0, 0x4d, 0xa8, 0xdb, 0x8a, 0x89, 0x3a, 0xfc, 0x61, 0x32, 0xf9, 0x11,
	0x75, 0x39, 0x71, 0x4e, 0x5e, 0xab, 0x04, 0xf5, 0x44, 0x72, 0x57, 0x55, 0x0c, 0x95, 0xa0, 0x3e,
	0x2f, 0xe6, 0xda, 0xb2, 0x66, 0xe8, 0x10, 0xf7, 0x01, 0x1c, 0xdd, 0xfa, 0x69, 0x6a, 0x89, 0x1c,
	0x88, 0x47, 0x18, 0xda, 0xcf, 0x2a, 0x41, 0x03, 0x11, 0x31, 0xb5, 0x2c, 0x5b, 0x25, 0x08, 0x8b,
	0x06, 0x6f, 0xe7, 0x73, 0xcd, 0x74, 0x35, 0x13, 0x1d, 0x89, 0x30, 0x45, 0x57, 0x0c, 0xc5, 0x9c,
	0xa9, 0x04, 0x1d, 0x73, 0xae, 0xad, 0x2b, 0xe6, 0xf4, 0x5a, 0x41, 0x27, 0x97, 0x9f, 0x97, 0x5f,
	0x5c, 0x71, 0xb6, 0x4d, 0xa8, 0xbf, 0x7a, 0x6d, 0xa0, 0x3d, 0x1e, 0x31, 0xbd, 0x56, 0xa7, 0x37,
	0xb6, 0xa5, 0x99, 0x2e, 0x92, 0x26, 0xad, 0x5f, 0x1a, 0xe2, 0x03, 0x9d, 0xdb, 0x7b, 0xb6, 0x64,
	0xd7, 0xee, 0x4a, 0xeb, 0xe5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x83, 0x73, 0xdd, 0x03,
	0x08, 0x00, 0x00,
}
