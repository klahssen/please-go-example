// Code generated by protoc-gen-go.
// source: data_battle.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

// Ignoring public import of GymMembership from data_gym.proto

// Ignoring public import of GymState from data_gym.proto

type BattleActionType int32

const (
	BattleActionType_ACTION_UNSET          BattleActionType = 0
	BattleActionType_ACTION_ATTACK         BattleActionType = 1
	BattleActionType_ACTION_DODGE          BattleActionType = 2
	BattleActionType_ACTION_SPECIAL_ATTACK BattleActionType = 3
	BattleActionType_ACTION_SWAP_POKEMON   BattleActionType = 4
	BattleActionType_ACTION_FAINT          BattleActionType = 5
	BattleActionType_ACTION_PLAYER_JOIN    BattleActionType = 6
	BattleActionType_ACTION_PLAYER_QUIT    BattleActionType = 7
	BattleActionType_ACTION_VICTORY        BattleActionType = 8
	BattleActionType_ACTION_DEFEAT         BattleActionType = 9
	BattleActionType_ACTION_TIMED_OUT      BattleActionType = 10
)

var BattleActionType_name = map[int32]string{
	0:  "ACTION_UNSET",
	1:  "ACTION_ATTACK",
	2:  "ACTION_DODGE",
	3:  "ACTION_SPECIAL_ATTACK",
	4:  "ACTION_SWAP_POKEMON",
	5:  "ACTION_FAINT",
	6:  "ACTION_PLAYER_JOIN",
	7:  "ACTION_PLAYER_QUIT",
	8:  "ACTION_VICTORY",
	9:  "ACTION_DEFEAT",
	10: "ACTION_TIMED_OUT",
}
var BattleActionType_value = map[string]int32{
	"ACTION_UNSET":          0,
	"ACTION_ATTACK":         1,
	"ACTION_DODGE":          2,
	"ACTION_SPECIAL_ATTACK": 3,
	"ACTION_SWAP_POKEMON":   4,
	"ACTION_FAINT":          5,
	"ACTION_PLAYER_JOIN":    6,
	"ACTION_PLAYER_QUIT":    7,
	"ACTION_VICTORY":        8,
	"ACTION_DEFEAT":         9,
	"ACTION_TIMED_OUT":      10,
}

func (x BattleActionType) String() string {
	return proto.EnumName(BattleActionType_name, int32(x))
}
func (BattleActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type BattleState int32

const (
	BattleState_STATE_UNSET BattleState = 0
	BattleState_ACTIVE      BattleState = 1
	BattleState_VICTORY     BattleState = 2
	BattleState_DEFEATED    BattleState = 3
	BattleState_TIMED_OUT   BattleState = 4
)

var BattleState_name = map[int32]string{
	0: "STATE_UNSET",
	1: "ACTIVE",
	2: "VICTORY",
	3: "DEFEATED",
	4: "TIMED_OUT",
}
var BattleState_value = map[string]int32{
	"STATE_UNSET": 0,
	"ACTIVE":      1,
	"VICTORY":     2,
	"DEFEATED":    3,
	"TIMED_OUT":   4,
}

func (x BattleState) String() string {
	return proto.EnumName(BattleState_name, int32(x))
}
func (BattleState) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type BattleType int32

const (
	BattleType_BATTLE_TYPE_UNSET    BattleType = 0
	BattleType_BATTLE_TYPE_NORMAL   BattleType = 1
	BattleType_BATTLE_TYPE_TRAINING BattleType = 2
)

var BattleType_name = map[int32]string{
	0: "BATTLE_TYPE_UNSET",
	1: "BATTLE_TYPE_NORMAL",
	2: "BATTLE_TYPE_TRAINING",
}
var BattleType_value = map[string]int32{
	"BATTLE_TYPE_UNSET":    0,
	"BATTLE_TYPE_NORMAL":   1,
	"BATTLE_TYPE_TRAINING": 2,
}

func (x BattleType) String() string {
	return proto.EnumName(BattleType_name, int32(x))
}
func (BattleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type BattleAction struct {
	Type                          BattleActionType   `protobuf:"varint,1,opt,name=type,enum=SUNProtos.Data.Battle.BattleActionType" json:"type,omitempty"`
	ActionStartMs                 int64              `protobuf:"varint,2,opt,name=action_start_ms,json=actionStartMs" json:"action_start_ms,omitempty"`
	DurationMs                    int32              `protobuf:"varint,3,opt,name=duration_ms,json=durationMs" json:"duration_ms,omitempty"`
	EnergyDelta                   int32              `protobuf:"varint,5,opt,name=energy_delta,json=energyDelta" json:"energy_delta,omitempty"`
	AttackerIndex                 int32              `protobuf:"varint,6,opt,name=attacker_index,json=attackerIndex" json:"attacker_index,omitempty"`
	TargetIndex                   int32              `protobuf:"varint,7,opt,name=target_index,json=targetIndex" json:"target_index,omitempty"`
	ActivePokemonId               uint64             `protobuf:"fixed64,8,opt,name=active_pokemon_id,json=activePokemonId" json:"active_pokemon_id,omitempty"`
	PlayerJoined                  *BattleParticipant `protobuf:"bytes,9,opt,name=player_joined,json=playerJoined" json:"player_joined,omitempty"`
	BattleResults                 *BattleResults     `protobuf:"bytes,10,opt,name=battle_results,json=battleResults" json:"battle_results,omitempty"`
	DamageWindowsStartTimestampMs int64              `protobuf:"varint,11,opt,name=damage_windows_start_timestamp_ms,json=damageWindowsStartTimestampMs" json:"damage_windows_start_timestamp_ms,omitempty"`
	DamageWindowsEndTimestampMs   int64              `protobuf:"varint,12,opt,name=damage_windows_end_timestamp_ms,json=damageWindowsEndTimestampMs" json:"damage_windows_end_timestamp_ms,omitempty"`
	PlayerLeft                    *BattleParticipant `protobuf:"bytes,13,opt,name=player_left,json=playerLeft" json:"player_left,omitempty"`
	TargetPokemonId               uint64             `protobuf:"fixed64,14,opt,name=target_pokemon_id,json=targetPokemonId" json:"target_pokemon_id,omitempty"`
}

func (m *BattleAction) Reset()                    { *m = BattleAction{} }
func (m *BattleAction) String() string            { return proto.CompactTextString(m) }
func (*BattleAction) ProtoMessage()               {}
func (*BattleAction) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *BattleAction) GetType() BattleActionType {
	if m != nil {
		return m.Type
	}
	return BattleActionType_ACTION_UNSET
}

func (m *BattleAction) GetActionStartMs() int64 {
	if m != nil {
		return m.ActionStartMs
	}
	return 0
}

func (m *BattleAction) GetDurationMs() int32 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

func (m *BattleAction) GetEnergyDelta() int32 {
	if m != nil {
		return m.EnergyDelta
	}
	return 0
}

func (m *BattleAction) GetAttackerIndex() int32 {
	if m != nil {
		return m.AttackerIndex
	}
	return 0
}

func (m *BattleAction) GetTargetIndex() int32 {
	if m != nil {
		return m.TargetIndex
	}
	return 0
}

func (m *BattleAction) GetActivePokemonId() uint64 {
	if m != nil {
		return m.ActivePokemonId
	}
	return 0
}

func (m *BattleAction) GetPlayerJoined() *BattleParticipant {
	if m != nil {
		return m.PlayerJoined
	}
	return nil
}

func (m *BattleAction) GetBattleResults() *BattleResults {
	if m != nil {
		return m.BattleResults
	}
	return nil
}

func (m *BattleAction) GetDamageWindowsStartTimestampMs() int64 {
	if m != nil {
		return m.DamageWindowsStartTimestampMs
	}
	return 0
}

func (m *BattleAction) GetDamageWindowsEndTimestampMs() int64 {
	if m != nil {
		return m.DamageWindowsEndTimestampMs
	}
	return 0
}

func (m *BattleAction) GetPlayerLeft() *BattleParticipant {
	if m != nil {
		return m.PlayerLeft
	}
	return nil
}

func (m *BattleAction) GetTargetPokemonId() uint64 {
	if m != nil {
		return m.TargetPokemonId
	}
	return 0
}

type BattleLog struct {
	State                  BattleState     `protobuf:"varint,1,opt,name=state,enum=SUNProtos.Data.Battle.BattleState" json:"state,omitempty"`
	BattleType             BattleType      `protobuf:"varint,2,opt,name=battle_type,json=battleType,enum=SUNProtos.Data.Battle.BattleType" json:"battle_type,omitempty"`
	ServerMs               int64           `protobuf:"varint,3,opt,name=server_ms,json=serverMs" json:"server_ms,omitempty"`
	BattleActions          []*BattleAction `protobuf:"bytes,4,rep,name=battle_actions,json=battleActions" json:"battle_actions,omitempty"`
	BattleStartTimestampMs int64           `protobuf:"varint,5,opt,name=battle_start_timestamp_ms,json=battleStartTimestampMs" json:"battle_start_timestamp_ms,omitempty"`
	BattleEndTimestampMs   int64           `protobuf:"varint,6,opt,name=battle_end_timestamp_ms,json=battleEndTimestampMs" json:"battle_end_timestamp_ms,omitempty"`
}

func (m *BattleLog) Reset()                    { *m = BattleLog{} }
func (m *BattleLog) String() string            { return proto.CompactTextString(m) }
func (*BattleLog) ProtoMessage()               {}
func (*BattleLog) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *BattleLog) GetState() BattleState {
	if m != nil {
		return m.State
	}
	return BattleState_STATE_UNSET
}

func (m *BattleLog) GetBattleType() BattleType {
	if m != nil {
		return m.BattleType
	}
	return BattleType_BATTLE_TYPE_UNSET
}

func (m *BattleLog) GetServerMs() int64 {
	if m != nil {
		return m.ServerMs
	}
	return 0
}

func (m *BattleLog) GetBattleActions() []*BattleAction {
	if m != nil {
		return m.BattleActions
	}
	return nil
}

func (m *BattleLog) GetBattleStartTimestampMs() int64 {
	if m != nil {
		return m.BattleStartTimestampMs
	}
	return 0
}

func (m *BattleLog) GetBattleEndTimestampMs() int64 {
	if m != nil {
		return m.BattleEndTimestampMs
	}
	return 0
}

type BattleParticipant struct {
	ActivePokemon        *BattlePokemonInfo   `protobuf:"bytes,1,opt,name=active_pokemon,json=activePokemon" json:"active_pokemon,omitempty"`
	TrainerPublicProfile *PlayerPublicProfile `protobuf:"bytes,2,opt,name=trainer_public_profile,json=trainerPublicProfile" json:"trainer_public_profile,omitempty"`
	ReversePokemon       []*BattlePokemonInfo `protobuf:"bytes,3,rep,name=reverse_pokemon,json=reversePokemon" json:"reverse_pokemon,omitempty"`
	DefeatedPokemon      []*BattlePokemonInfo `protobuf:"bytes,4,rep,name=defeated_pokemon,json=defeatedPokemon" json:"defeated_pokemon,omitempty"`
}

func (m *BattleParticipant) Reset()                    { *m = BattleParticipant{} }
func (m *BattleParticipant) String() string            { return proto.CompactTextString(m) }
func (*BattleParticipant) ProtoMessage()               {}
func (*BattleParticipant) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *BattleParticipant) GetActivePokemon() *BattlePokemonInfo {
	if m != nil {
		return m.ActivePokemon
	}
	return nil
}

func (m *BattleParticipant) GetTrainerPublicProfile() *PlayerPublicProfile {
	if m != nil {
		return m.TrainerPublicProfile
	}
	return nil
}

func (m *BattleParticipant) GetReversePokemon() []*BattlePokemonInfo {
	if m != nil {
		return m.ReversePokemon
	}
	return nil
}

func (m *BattleParticipant) GetDefeatedPokemon() []*BattlePokemonInfo {
	if m != nil {
		return m.DefeatedPokemon
	}
	return nil
}

type BattlePokemonInfo struct {
	PokemonData   *PokemonData `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	CurrentHealth int32        `protobuf:"varint,2,opt,name=current_health,json=currentHealth" json:"current_health,omitempty"`
	CurrentEnergy int32        `protobuf:"varint,3,opt,name=current_energy,json=currentEnergy" json:"current_energy,omitempty"`
}

func (m *BattlePokemonInfo) Reset()                    { *m = BattlePokemonInfo{} }
func (m *BattlePokemonInfo) String() string            { return proto.CompactTextString(m) }
func (*BattlePokemonInfo) ProtoMessage()               {}
func (*BattlePokemonInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *BattlePokemonInfo) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *BattlePokemonInfo) GetCurrentHealth() int32 {
	if m != nil {
		return m.CurrentHealth
	}
	return 0
}

func (m *BattlePokemonInfo) GetCurrentEnergy() int32 {
	if m != nil {
		return m.CurrentEnergy
	}
	return 0
}

type BattleResults struct {
	GymState                *GymState            `protobuf:"bytes,1,opt,name=gym_state,json=gymState" json:"gym_state,omitempty"`
	Attackers               []*BattleParticipant `protobuf:"bytes,2,rep,name=attackers" json:"attackers,omitempty"`
	PlayerExperienceAwarded []int32              `protobuf:"varint,3,rep,packed,name=player_experience_awarded,json=playerExperienceAwarded" json:"player_experience_awarded,omitempty"`
	NextDefenderPokemonId   int64                `protobuf:"varint,4,opt,name=next_defender_pokemon_id,json=nextDefenderPokemonId" json:"next_defender_pokemon_id,omitempty"`
	GymPointsDelta          int32                `protobuf:"varint,5,opt,name=gym_points_delta,json=gymPointsDelta" json:"gym_points_delta,omitempty"`
}

func (m *BattleResults) Reset()                    { *m = BattleResults{} }
func (m *BattleResults) String() string            { return proto.CompactTextString(m) }
func (*BattleResults) ProtoMessage()               {}
func (*BattleResults) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *BattleResults) GetGymState() *GymState {
	if m != nil {
		return m.GymState
	}
	return nil
}

func (m *BattleResults) GetAttackers() []*BattleParticipant {
	if m != nil {
		return m.Attackers
	}
	return nil
}

func (m *BattleResults) GetPlayerExperienceAwarded() []int32 {
	if m != nil {
		return m.PlayerExperienceAwarded
	}
	return nil
}

func (m *BattleResults) GetNextDefenderPokemonId() int64 {
	if m != nil {
		return m.NextDefenderPokemonId
	}
	return 0
}

func (m *BattleResults) GetGymPointsDelta() int32 {
	if m != nil {
		return m.GymPointsDelta
	}
	return 0
}

func init() {
	proto.RegisterType((*BattleAction)(nil), "SUNProtos.Data.Battle.BattleAction")
	proto.RegisterType((*BattleLog)(nil), "SUNProtos.Data.Battle.BattleLog")
	proto.RegisterType((*BattleParticipant)(nil), "SUNProtos.Data.Battle.BattleParticipant")
	proto.RegisterType((*BattlePokemonInfo)(nil), "SUNProtos.Data.Battle.BattlePokemonInfo")
	proto.RegisterType((*BattleResults)(nil), "SUNProtos.Data.Battle.BattleResults")
	proto.RegisterEnum("SUNProtos.Data.Battle.BattleActionType", BattleActionType_name, BattleActionType_value)
	proto.RegisterEnum("SUNProtos.Data.Battle.BattleState", BattleState_name, BattleState_value)
	proto.RegisterEnum("SUNProtos.Data.Battle.BattleType", BattleType_name, BattleType_value)
}

func init() { proto.RegisterFile("data_battle.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1080 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x0d, 0xf5, 0x8a, 0x74, 0x29, 0xc9, 0xf4, 0xd4, 0x0f, 0x3a, 0x6e, 0x11, 0x59, 0x7d, 0x09,
	0x5e, 0x78, 0xe1, 0xa2, 0x68, 0xd3, 0x02, 0x05, 0x68, 0x8b, 0x76, 0x68, 0xeb, 0xc1, 0x50, 0x74,
	0x02, 0x77, 0x33, 0x1d, 0x89, 0x63, 0x85, 0x8d, 0x44, 0x11, 0xe4, 0xf8, 0xa1, 0xcf, 0xe9, 0x27,
	0xf4, 0x4f, 0xfa, 0x11, 0x5d, 0xf7, 0x0b, 0xba, 0x28, 0x38, 0x33, 0x94, 0x28, 0x3b, 0x70, 0xe3,
	0x85, 0xe1, 0x99, 0x73, 0xcf, 0x3d, 0x1c, 0xde, 0xc3, 0x7b, 0x47, 0xb0, 0xee, 0x11, 0x46, 0xf0,
	0x90, 0x30, 0x36, 0xa1, 0x07, 0x61, 0x34, 0x63, 0x33, 0xb4, 0x39, 0xb8, 0xe8, 0xd9, 0xc9, 0x2a,
	0x3e, 0x68, 0x13, 0x46, 0x0e, 0x8e, 0x78, 0xf0, 0x85, 0x60, 0x86, 0x13, 0x32, 0xa7, 0x91, 0x60,
	0xbe, 0x80, 0x04, 0x92, 0xeb, 0x3a, 0x0f, 0x8f, 0xe7, 0x53, 0xb1, 0x6f, 0xfe, 0x55, 0x84, 0xaa,
	0xc8, 0x34, 0x46, 0xcc, 0x9f, 0x05, 0xe8, 0x67, 0x28, 0xb0, 0x79, 0x48, 0x75, 0xa5, 0xa1, 0xb4,
	0xea, 0x87, 0xdf, 0x1e, 0x7c, 0xf4, 0x29, 0x07, 0xd9, 0x14, 0x77, 0x1e, 0x52, 0x87, 0x27, 0xa1,
	0x6f, 0x60, 0x8d, 0x70, 0x0c, 0xc7, 0x8c, 0x44, 0x0c, 0x4f, 0x63, 0x3d, 0xd7, 0x50, 0x5a, 0x79,
	0xa7, 0x26, 0xe0, 0x41, 0x82, 0x76, 0x63, 0xf4, 0x12, 0x54, 0xef, 0x3a, 0x22, 0x9c, 0x39, 0x8d,
	0xf5, 0x7c, 0x43, 0x69, 0x15, 0x1d, 0x48, 0xa1, 0x6e, 0x8c, 0xf6, 0xa0, 0x4a, 0x03, 0x1a, 0x8d,
	0xe7, 0xd8, 0xa3, 0x13, 0x46, 0xf4, 0x22, 0x67, 0xa8, 0x02, 0x6b, 0x27, 0x10, 0xfa, 0x1a, 0xea,
	0x84, 0x31, 0x32, 0xfa, 0x40, 0x23, 0xec, 0x07, 0x1e, 0xbd, 0xd3, 0x4b, 0x9c, 0x54, 0x4b, 0x51,
	0x2b, 0x01, 0x13, 0x25, 0x46, 0xa2, 0x31, 0x65, 0x92, 0xf4, 0x5c, 0x28, 0x09, 0x4c, 0x50, 0xf6,
	0x61, 0x3d, 0x39, 0xde, 0x0d, 0xc5, 0xe1, 0xec, 0x03, 0x9d, 0xce, 0x02, 0xec, 0x7b, 0x7a, 0xb9,
	0xa1, 0xb4, 0x4a, 0xce, 0x9a, 0x08, 0xd8, 0x02, 0xb7, 0x3c, 0xd4, 0x85, 0x9a, 0xa8, 0x2d, 0xfe,
	0x7d, 0xe6, 0x07, 0xd4, 0xd3, 0x2b, 0x0d, 0xa5, 0xa5, 0x1e, 0xb6, 0x1e, 0xad, 0x93, 0x4d, 0x22,
	0xe6, 0x8f, 0xfc, 0x90, 0x04, 0xcc, 0xa9, 0x8a, 0xf4, 0x33, 0x9e, 0x8d, 0xce, 0xa1, 0x2e, 0x4c,
	0xc5, 0x11, 0x8d, 0xaf, 0x27, 0x2c, 0xd6, 0x81, 0xeb, 0x7d, 0xf5, 0xa8, 0x9e, 0x23, 0xb8, 0x4e,
	0x6d, 0x98, 0xdd, 0xa2, 0xd7, 0xb0, 0xe7, 0x91, 0x29, 0x19, 0x53, 0x7c, 0xeb, 0x07, 0xde, 0xec,
	0x36, 0x96, 0x2e, 0x30, 0x7f, 0x4a, 0x63, 0x46, 0xa6, 0x61, 0x52, 0x6b, 0x95, 0xfb, 0xf1, 0x85,
	0x20, 0xbe, 0x13, 0x3c, 0x6e, 0x8b, 0x9b, 0xb2, 0xba, 0x31, 0x6a, 0xc3, 0xcb, 0x7b, 0x4a, 0x34,
	0xf0, 0x56, 0x75, 0xaa, 0x5c, 0x67, 0x77, 0x45, 0xc7, 0x0c, 0xbc, 0xac, 0x8a, 0x05, 0xaa, 0xac,
	0xd5, 0x84, 0x5e, 0x31, 0xbd, 0xf6, 0xc4, 0x4a, 0x81, 0x48, 0xee, 0xd0, 0x2b, 0x96, 0x58, 0x24,
	0x5d, 0xcc, 0x58, 0x54, 0x17, 0x16, 0x89, 0xc0, 0xc2, 0xa2, 0xe6, 0xdf, 0x39, 0xa8, 0x08, 0xb5,
	0xce, 0x6c, 0x8c, 0x7e, 0x84, 0x62, 0xcc, 0x08, 0x4b, 0x3f, 0xe8, 0xe6, 0xa3, 0x8f, 0x1f, 0x24,
	0x4c, 0x47, 0x24, 0xa0, 0x23, 0x50, 0xa5, 0x37, 0xbc, 0x21, 0x72, 0x3c, 0x7f, 0xef, 0xd1, 0x7c,
	0xde, 0x0a, 0x30, 0x5c, 0xac, 0xd1, 0x2e, 0x54, 0x62, 0x1a, 0xdd, 0xd0, 0x28, 0xfd, 0xcc, 0xf3,
	0x4e, 0x59, 0x00, 0xdd, 0x18, 0x9d, 0x2d, 0xcc, 0x17, 0xdd, 0x11, 0xeb, 0x85, 0x46, 0xbe, 0xa5,
	0x1e, 0x7e, 0xf9, 0x09, 0x4d, 0x97, 0x7a, 0x2f, 0x76, 0x31, 0x7a, 0x05, 0x3b, 0x52, 0xeb, 0x23,
	0x9e, 0x17, 0xf9, 0x83, 0xb7, 0x86, 0xe9, 0x3b, 0xae, 0x9a, 0xfd, 0x3d, 0x6c, 0xcb, 0xd4, 0x07,
	0x26, 0x97, 0x78, 0xe2, 0x86, 0x08, 0xaf, 0xba, 0xdb, 0xfc, 0x27, 0x07, 0xeb, 0x0f, 0x4c, 0x43,
	0x7d, 0xa8, 0xaf, 0xf6, 0x12, 0xaf, 0xfb, 0xff, 0xda, 0x2e, 0xcd, 0x0b, 0xae, 0x66, 0x62, 0x54,
	0x2c, 0x5a, 0x0e, 0xfd, 0x06, 0x5b, 0x2c, 0x22, 0x7e, 0x40, 0x23, 0x1c, 0x5e, 0x0f, 0x27, 0xfe,
	0x08, 0x87, 0xd1, 0xec, 0xca, 0x9f, 0x08, 0x43, 0xd4, 0xc3, 0xfd, 0xfb, 0xc2, 0xb6, 0x18, 0x7d,
	0xe2, 0x9f, 0xcd, 0x53, 0x6c, 0x91, 0xe1, 0x6c, 0x48, 0xa5, 0x15, 0x14, 0xbd, 0x81, 0xb5, 0x88,
	0xde, 0xd0, 0x28, 0x5e, 0x9e, 0x39, 0xcf, 0x7d, 0xf8, 0xf4, 0x33, 0xd7, 0xa5, 0x40, 0x7a, 0xe8,
	0x01, 0x68, 0x1e, 0xbd, 0xa2, 0x84, 0x51, 0x6f, 0xa1, 0x59, 0x78, 0xa2, 0xe6, 0x5a, 0xaa, 0x20,
	0xc1, 0xe6, 0x1f, 0xca, 0xa2, 0xe0, 0x4b, 0x1a, 0xfa, 0x05, 0xaa, 0x69, 0x4b, 0x24, 0xa3, 0x5d,
	0x96, 0x7b, 0xf7, 0x41, 0x55, 0x04, 0x27, 0x59, 0x3b, 0x6a, 0xb8, 0xdc, 0x24, 0x63, 0x74, 0x74,
	0x1d, 0x45, 0x34, 0x60, 0xf8, 0x3d, 0x25, 0x13, 0xf6, 0x9e, 0xd7, 0xb5, 0xe8, 0xd4, 0x24, 0xfa,
	0x9a, 0x83, 0x59, 0x9a, 0x18, 0xc2, 0x72, 0x68, 0xa7, 0x34, 0x93, 0x83, 0xcd, 0x3f, 0x73, 0x50,
	0x5b, 0x99, 0x51, 0xe8, 0x15, 0x54, 0xc6, 0xf3, 0x29, 0x5e, 0xf6, 0xa0, 0x7a, 0xf8, 0xf9, 0xfd,
	0xc3, 0x9d, 0xce, 0xa7, 0xc9, 0x9f, 0xe8, 0xbe, 0xf2, 0x58, 0xae, 0xd0, 0x09, 0x54, 0xd2, 0x59,
	0x9e, 0xdc, 0x23, 0xf9, 0x27, 0x4d, 0x8f, 0x65, 0x2a, 0xfa, 0x09, 0x76, 0xe4, 0x1c, 0xa2, 0x77,
	0x21, 0x8d, 0x7c, 0x1a, 0x8c, 0x28, 0x26, 0xb7, 0x24, 0xf2, 0xa8, 0xc7, 0xad, 0x2e, 0x3a, 0xdb,
	0x82, 0x60, 0x2e, 0xe2, 0x86, 0x08, 0xa3, 0x1f, 0x40, 0x0f, 0xe8, 0x1d, 0xc3, 0x89, 0x19, 0x81,
	0x97, 0x7c, 0x84, 0xcb, 0xf9, 0x53, 0xe0, 0xdd, 0xb1, 0x99, 0xc4, 0xdb, 0x32, 0xbc, 0xbc, 0x28,
	0x5a, 0xa0, 0x25, 0xef, 0x1d, 0xce, 0xfc, 0x80, 0xc5, 0x2b, 0xb7, 0x58, 0x7d, 0x3c, 0x9f, 0xda,
	0x1c, 0xe6, 0x17, 0xd9, 0xfe, 0xbf, 0x0a, 0x68, 0xf7, 0xef, 0x53, 0xa4, 0x41, 0xd5, 0x38, 0x76,
	0xad, 0x7e, 0x0f, 0x5f, 0xf4, 0x06, 0xa6, 0xab, 0x3d, 0x43, 0xeb, 0x50, 0x93, 0x88, 0xe1, 0xba,
	0xc6, 0xf1, 0xb9, 0xa6, 0x64, 0x48, 0xed, 0x7e, 0xfb, 0xd4, 0xd4, 0x72, 0x68, 0x07, 0x36, 0x25,
	0x32, 0xb0, 0xcd, 0x63, 0xcb, 0xe8, 0xa4, 0xe4, 0x3c, 0xda, 0x86, 0xcf, 0xd2, 0xd0, 0x3b, 0xc3,
	0xc6, 0x76, 0xff, 0xdc, 0xec, 0xf6, 0x7b, 0x5a, 0x21, 0xa3, 0x72, 0x62, 0x58, 0x3d, 0x57, 0x2b,
	0xa2, 0x2d, 0x40, 0x12, 0xb1, 0x3b, 0xc6, 0xa5, 0xe9, 0xe0, 0xb3, 0xbe, 0xd5, 0xd3, 0x4a, 0x0f,
	0xf1, 0x37, 0x17, 0x96, 0xab, 0x3d, 0x47, 0x08, 0xea, 0x12, 0x7f, 0x6b, 0x1d, 0xbb, 0x7d, 0xe7,
	0x52, 0x2b, 0x67, 0x8e, 0xdb, 0x36, 0x4f, 0x4c, 0xc3, 0xd5, 0x2a, 0x68, 0x03, 0x34, 0x09, 0xb9,
	0x56, 0xd7, 0x6c, 0xe3, 0xfe, 0x85, 0xab, 0xc1, 0xbe, 0x0b, 0x6a, 0x66, 0xf8, 0xa2, 0x35, 0x50,
	0x07, 0xae, 0xe1, 0x9a, 0x8b, 0xf7, 0x06, 0x28, 0x25, 0x59, 0x6f, 0x4d, 0x4d, 0x41, 0x2a, 0x3c,
	0x4f, 0x9f, 0x90, 0x43, 0x55, 0x28, 0x0b, 0x69, 0xb3, 0xad, 0xe5, 0x51, 0x0d, 0x2a, 0x4b, 0xd5,
	0xc2, 0xfe, 0x05, 0xc0, 0x72, 0x24, 0xa3, 0x4d, 0x58, 0x3f, 0x32, 0x5c, 0xb7, 0x63, 0x62, 0xf7,
	0xd2, 0x5e, 0x4a, 0x6f, 0x01, 0xca, 0xc2, 0xbd, 0xbe, 0xd3, 0x35, 0x3a, 0x9a, 0x82, 0x74, 0xd8,
	0xc8, 0xe2, 0xae, 0x63, 0x58, 0x3d, 0xab, 0x77, 0xaa, 0xe5, 0x8e, 0xca, 0xbf, 0x96, 0xf8, 0xef,
	0xa6, 0xd8, 0x7e, 0x66, 0x2b, 0x76, 0x6e, 0x28, 0x76, 0xdf, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x28, 0x64, 0xd3, 0xa0, 0x09, 0x00, 0x00,
}
