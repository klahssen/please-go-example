// Code generated by protoc-gen-go.
// source: maps.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of MapPokemon from map_pokemon.proto

// Ignoring public import of NearbyPokemon from map_pokemon.proto

// Ignoring public import of WildPokemon from map_pokemon.proto

type MapObjectsStatus int32

const (
	MapObjectsStatus_UNSET_STATUS   MapObjectsStatus = 0
	MapObjectsStatus_SUCCESS        MapObjectsStatus = 1
	MapObjectsStatus_LOCATION_UNSET MapObjectsStatus = 2
)

var MapObjectsStatus_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "SUCCESS",
	2: "LOCATION_UNSET",
}
var MapObjectsStatus_value = map[string]int32{
	"UNSET_STATUS":   0,
	"SUCCESS":        1,
	"LOCATION_UNSET": 2,
}

func (x MapObjectsStatus) String() string {
	return proto.EnumName(MapObjectsStatus_name, int32(x))
}
func (MapObjectsStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

type MapCell struct {
	// S2 geographic area that the cell covers (http://s2map.com/) (https://code.google.com/archive/p/s2-geometry-library/)
	S2CellId             uint64         `protobuf:"varint,1,opt,name=s2_cell_id,json=s2CellId" json:"s2_cell_id,omitempty"`
	CurrentTimestampMs   int64          `protobuf:"varint,2,opt,name=current_timestamp_ms,json=currentTimestampMs" json:"current_timestamp_ms,omitempty"`
	Forts                []*FortData    `protobuf:"bytes,3,rep,name=forts" json:"forts,omitempty"`
	SpawnPoints          []*SpawnPoint  `protobuf:"bytes,4,rep,name=spawn_points,json=spawnPoints" json:"spawn_points,omitempty"`
	DeletedObjects       []string       `protobuf:"bytes,6,rep,name=deleted_objects,json=deletedObjects" json:"deleted_objects,omitempty"`
	IsTruncatedList      bool           `protobuf:"varint,7,opt,name=is_truncated_list,json=isTruncatedList" json:"is_truncated_list,omitempty"`
	FortSummaries        []*FortSummary `protobuf:"bytes,8,rep,name=fort_summaries,json=fortSummaries" json:"fort_summaries,omitempty"`
	DecimatedSpawnPoints []*SpawnPoint  `protobuf:"bytes,9,rep,name=decimated_spawn_points,json=decimatedSpawnPoints" json:"decimated_spawn_points,omitempty"`
	// Pokemon within 2 steps or less.
	WildPokemons []*WildPokemon `protobuf:"bytes,5,rep,name=wild_pokemons,json=wildPokemons" json:"wild_pokemons,omitempty"`
	// Pokemon within 1 step or none.
	CatchablePokemons []*MapPokemon `protobuf:"bytes,10,rep,name=catchable_pokemons,json=catchablePokemons" json:"catchable_pokemons,omitempty"`
	// Pokemon farther away than 2 steps, but still in the area.
	NearbyPokemons []*NearbyPokemon `protobuf:"bytes,11,rep,name=nearby_pokemons,json=nearbyPokemons" json:"nearby_pokemons,omitempty"`
}

func (m *MapCell) Reset()                    { *m = MapCell{} }
func (m *MapCell) String() string            { return proto.CompactTextString(m) }
func (*MapCell) ProtoMessage()               {}
func (*MapCell) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *MapCell) GetS2CellId() uint64 {
	if m != nil {
		return m.S2CellId
	}
	return 0
}

func (m *MapCell) GetCurrentTimestampMs() int64 {
	if m != nil {
		return m.CurrentTimestampMs
	}
	return 0
}

func (m *MapCell) GetForts() []*FortData {
	if m != nil {
		return m.Forts
	}
	return nil
}

func (m *MapCell) GetSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.SpawnPoints
	}
	return nil
}

func (m *MapCell) GetDeletedObjects() []string {
	if m != nil {
		return m.DeletedObjects
	}
	return nil
}

func (m *MapCell) GetIsTruncatedList() bool {
	if m != nil {
		return m.IsTruncatedList
	}
	return false
}

func (m *MapCell) GetFortSummaries() []*FortSummary {
	if m != nil {
		return m.FortSummaries
	}
	return nil
}

func (m *MapCell) GetDecimatedSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.DecimatedSpawnPoints
	}
	return nil
}

func (m *MapCell) GetWildPokemons() []*WildPokemon {
	if m != nil {
		return m.WildPokemons
	}
	return nil
}

func (m *MapCell) GetCatchablePokemons() []*MapPokemon {
	if m != nil {
		return m.CatchablePokemons
	}
	return nil
}

func (m *MapCell) GetNearbyPokemons() []*NearbyPokemon {
	if m != nil {
		return m.NearbyPokemons
	}
	return nil
}

type SpawnPoint struct {
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *SpawnPoint) Reset()                    { *m = SpawnPoint{} }
func (m *SpawnPoint) String() string            { return proto.CompactTextString(m) }
func (*SpawnPoint) ProtoMessage()               {}
func (*SpawnPoint) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *SpawnPoint) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *SpawnPoint) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*MapCell)(nil), "SUNProtos.Map.MapCell")
	proto.RegisterType((*SpawnPoint)(nil), "SUNProtos.Map.SpawnPoint")
	proto.RegisterEnum("SUNProtos.Map.MapObjectsStatus", MapObjectsStatus_name, MapObjectsStatus_value)
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6b, 0xdb, 0x4c,
	0x10, 0xc6, 0xa3, 0x38, 0x8e, 0xed, 0xf1, 0xff, 0x25, 0xbc, 0xe8, 0x35, 0x86, 0xba, 0xa6, 0x50,
	0x93, 0x83, 0x29, 0xee, 0xb5, 0x97, 0xd4, 0xb5, 0x4b, 0x20, 0xb2, 0x85, 0x24, 0x53, 0xe8, 0x65,
	0x59, 0x4b, 0x9b, 0x76, 0xdb, 0x95, 0xb4, 0x68, 0xd6, 0x98, 0x7c, 0x9e, 0x7e, 0xd1, 0xa2, 0x95,
	0x2d, 0x27, 0x01, 0xd3, 0x8b, 0x98, 0x99, 0xe7, 0x99, 0xdf, 0xce, 0xb0, 0x2b, 0x80, 0x98, 0x29,
	0x9c, 0xaa, 0x2c, 0xd5, 0x29, 0x69, 0xfb, 0x9b, 0x95, 0x9b, 0x47, 0x38, 0x75, 0x98, 0x1a, 0x74,
	0x62, 0xa6, 0xe8, 0x63, 0x9a, 0xe9, 0x42, 0x1e, 0xf4, 0xf3, 0x5c, 0xa5, 0xbf, 0x79, 0x9c, 0x26,
	0x45, 0x69, 0xfc, 0xa7, 0x0a, 0x35, 0x87, 0xa9, 0x39, 0x97, 0x92, 0x0c, 0x01, 0x70, 0x46, 0x43,
	0x2e, 0x25, 0x15, 0x91, 0x6d, 0x8d, 0xac, 0xc9, 0x95, 0x57, 0xc7, 0x59, 0xae, 0xdd, 0x47, 0xe4,
	0x03, 0xdc, 0x84, 0xbb, 0x2c, 0xe3, 0x89, 0xa6, 0x5a, 0xc4, 0x1c, 0x35, 0x8b, 0x15, 0x8d, 0xd1,
	0xbe, 0x1c, 0x59, 0x93, 0x8a, 0x47, 0x0e, 0x5a, 0x70, 0x94, 0x1c, 0x24, 0x33, 0xa8, 0xe6, 0x87,
	0xa3, 0x5d, 0x19, 0x55, 0x26, 0xcd, 0xd9, 0x70, 0xfa, 0x62, 0xba, 0xe9, 0x32, 0x1f, 0x2c, 0xff,
	0x7c, 0x61, 0x9a, 0x79, 0x85, 0x95, 0x7c, 0x82, 0x16, 0x2a, 0xb6, 0x4f, 0xa8, 0x4a, 0x45, 0xa2,
	0xd1, 0xbe, 0x32, 0xad, 0xff, 0xbf, 0x6a, 0xf5, 0x73, 0x8b, 0x9b, 0x3b, 0xbc, 0x26, 0x96, 0x31,
	0x92, 0xf7, 0xd0, 0x8d, 0xb8, 0xe4, 0x9a, 0x47, 0x34, 0xdd, 0xfe, 0xe2, 0xa1, 0x46, 0xfb, 0x7a,
	0x54, 0x99, 0x34, 0xbc, 0xce, 0xa1, 0xbc, 0x2e, 0xaa, 0xe4, 0x16, 0xfa, 0x02, 0xa9, 0xce, 0x76,
	0x49, 0xc8, 0x72, 0xb7, 0x14, 0xa8, 0xed, 0xda, 0xc8, 0x9a, 0xd4, 0xbd, 0xae, 0xc0, 0xe0, 0x58,
	0x7f, 0x10, 0xa8, 0xc9, 0x12, 0x3a, 0xf9, 0x6c, 0x14, 0x77, 0x71, 0xcc, 0x32, 0xc1, 0xd1, 0xae,
	0x9b, 0xa1, 0xde, 0x9c, 0xdb, 0xc7, 0x37, 0xc6, 0x27, 0xaf, 0xfd, 0x58, 0x26, 0x82, 0x23, 0x59,
	0xc3, 0x7f, 0x11, 0x0f, 0x45, 0x6c, 0x0e, 0x7c, 0xb1, 0x64, 0xe3, 0x5f, 0x4b, 0xde, 0x94, 0x8d,
	0xfe, 0xb3, 0x6d, 0xbf, 0x42, 0x7b, 0x2f, 0x64, 0x74, 0xbc, 0x51, 0xb4, 0xab, 0x86, 0x33, 0x7e,
	0xc5, 0x71, 0x0f, 0x17, 0xfe, 0x4d, 0xc8, 0xe8, 0x10, 0x7b, 0xad, 0xfd, 0x29, 0x41, 0xe2, 0x02,
	0x09, 0x99, 0x0e, 0x7f, 0xb2, 0xad, 0xe4, 0x27, 0x1a, 0x18, 0xda, 0xdb, 0x33, 0x34, 0x87, 0xa9,
	0x23, 0xac, 0x5f, 0x36, 0x97, 0x44, 0x07, 0xba, 0x09, 0x67, 0xd9, 0xf6, 0xe9, 0x84, 0x6b, 0x1a,
	0xdc, 0xbb, 0x33, 0xb8, 0x95, 0x71, 0x1f, 0x89, 0x9d, 0xe4, 0x79, 0x8a, 0xe3, 0x25, 0xc0, 0x69,
	0x71, 0x32, 0x80, 0xba, 0x64, 0x5a, 0xe8, 0x5d, 0xc4, 0xcd, 0xeb, 0xb3, 0xbc, 0x32, 0x27, 0x43,
	0x68, 0xc8, 0x34, 0xf9, 0x51, 0x88, 0x15, 0x23, 0x9e, 0x0a, 0xb7, 0x0b, 0xe8, 0x39, 0x4c, 0x1d,
	0x1e, 0x81, 0xaf, 0x99, 0xde, 0x21, 0xe9, 0x41, 0x6b, 0xb3, 0xf2, 0x17, 0x01, 0xf5, 0x83, 0xbb,
	0x60, 0xe3, 0xf7, 0x2e, 0x48, 0x13, 0x6a, 0xfe, 0x66, 0x3e, 0x5f, 0xf8, 0x7e, 0xcf, 0x22, 0x04,
	0x3a, 0x0f, 0xeb, 0xf9, 0x5d, 0x70, 0xbf, 0x5e, 0x51, 0xe3, 0xeb, 0x5d, 0x7e, 0xae, 0x7f, 0xbf,
	0x36, 0x7f, 0x0f, 0xba, 0x17, 0xae, 0xb5, 0x2d, 0xe2, 0x8f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xb2, 0xee, 0x6f, 0x89, 0x03, 0x00, 0x00,
}