// Code generated by protoc-gen-go.
// source: data_badge.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BadgeCaptureReward struct {
	CaptureRewardMultiplier float32  `protobuf:"fixed32,1,opt,name=capture_reward_multiplier,json=captureRewardMultiplier" json:"capture_reward_multiplier,omitempty"`
	AvatarTemplateIds       []string `protobuf:"bytes,2,rep,name=avatar_template_ids,json=avatarTemplateIds" json:"avatar_template_ids,omitempty"`
}

func (m *BadgeCaptureReward) Reset()                    { *m = BadgeCaptureReward{} }
func (m *BadgeCaptureReward) String() string            { return proto.CompactTextString(m) }
func (*BadgeCaptureReward) ProtoMessage()               {}
func (*BadgeCaptureReward) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BadgeCaptureReward) GetCaptureRewardMultiplier() float32 {
	if m != nil {
		return m.CaptureRewardMultiplier
	}
	return 0
}

func (m *BadgeCaptureReward) GetAvatarTemplateIds() []string {
	if m != nil {
		return m.AvatarTemplateIds
	}
	return nil
}

func init() {
	proto.RegisterType((*BadgeCaptureReward)(nil), "SUNProtos.Data.Badge.BadgeCaptureReward")
}

func init() { proto.RegisterFile("data_badge.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x8c, 0x4f, 0x4a, 0x4c, 0x49, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x09, 0x0e,
	0xf5, 0x0b, 0x00, 0xb1, 0x8a, 0xf5, 0x5c, 0x12, 0x4b, 0x12, 0xf5, 0x9c, 0x40, 0x72, 0x4a, 0x0d,
	0x8c, 0x5c, 0x42, 0x60, 0x96, 0x73, 0x62, 0x41, 0x49, 0x69, 0x51, 0x6a, 0x50, 0x6a, 0x79, 0x62,
	0x51, 0x8a, 0x90, 0x15, 0x97, 0x64, 0x32, 0x44, 0x20, 0xbe, 0x08, 0x2c, 0x12, 0x9f, 0x5b, 0x9a,
	0x53, 0x92, 0x59, 0x90, 0x93, 0x99, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x14, 0x24, 0x9e,
	0x8c, 0xac, 0xc3, 0x17, 0x2e, 0x2d, 0xa4, 0xc7, 0x25, 0x9c, 0x58, 0x96, 0x58, 0x92, 0x58, 0x14,
	0x5f, 0x92, 0x9a, 0x5b, 0x90, 0x93, 0x58, 0x92, 0x1a, 0x9f, 0x99, 0x52, 0x2c, 0xc1, 0xa4, 0xc0,
	0xac, 0xc1, 0x19, 0x24, 0x08, 0x91, 0x0a, 0x81, 0xca, 0x78, 0xa6, 0x14, 0x3b, 0x71, 0x44, 0xb1,
	0x81, 0x5d, 0x58, 0x9c, 0x04, 0xa1, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x3e, 0xf0,
	0xff, 0xbd, 0x00, 0x00, 0x00,
}
