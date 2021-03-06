// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GameInput struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Item                 string   `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	ModelName            string   `protobuf:"bytes,4,opt,name=modelName,proto3" json:"modelName,omitempty"`
	ModelNamespace       string   `protobuf:"bytes,5,opt,name=modelNamespace,proto3" json:"modelNamespace,omitempty"`
	EnemyHp              int32    `protobuf:"varint,6,opt,name=enemyHp,proto3" json:"enemyHp,omitempty"`
	HeroHp               int32    `protobuf:"varint,7,opt,name=heroHp,proto3" json:"heroHp,omitempty"`
	Interactions         int32    `protobuf:"varint,8,opt,name=interactions,proto3" json:"interactions,omitempty"`
	IsHeroTurn           bool     `protobuf:"varint,9,opt,name=isHeroTurn,proto3" json:"isHeroTurn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameInput) Reset()         { *m = GameInput{} }
func (m *GameInput) String() string { return proto.CompactTextString(m) }
func (*GameInput) ProtoMessage()    {}
func (*GameInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *GameInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameInput.Unmarshal(m, b)
}
func (m *GameInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameInput.Marshal(b, m, deterministic)
}
func (m *GameInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameInput.Merge(m, src)
}
func (m *GameInput) XXX_Size() int {
	return xxx_messageInfo_GameInput.Size(m)
}
func (m *GameInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GameInput.DiscardUnknown(m)
}

var xxx_messageInfo_GameInput proto.InternalMessageInfo

func (m *GameInput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GameInput) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

func (m *GameInput) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *GameInput) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *GameInput) GetModelNamespace() string {
	if m != nil {
		return m.ModelNamespace
	}
	return ""
}

func (m *GameInput) GetEnemyHp() int32 {
	if m != nil {
		return m.EnemyHp
	}
	return 0
}

func (m *GameInput) GetHeroHp() int32 {
	if m != nil {
		return m.HeroHp
	}
	return 0
}

func (m *GameInput) GetInteractions() int32 {
	if m != nil {
		return m.Interactions
	}
	return 0
}

func (m *GameInput) GetIsHeroTurn() bool {
	if m != nil {
		return m.IsHeroTurn
	}
	return false
}

type GameOutput struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	NewItem              string   `protobuf:"bytes,2,opt,name=newItem,proto3" json:"newItem,omitempty"`
	ModelName            string   `protobuf:"bytes,3,opt,name=modelName,proto3" json:"modelName,omitempty"`
	ModelNamespace       string   `protobuf:"bytes,4,opt,name=modelNamespace,proto3" json:"modelNamespace,omitempty"`
	NewEnemyHp           int32    `protobuf:"varint,5,opt,name=newEnemyHp,proto3" json:"newEnemyHp,omitempty"`
	NewHeroHp            int32    `protobuf:"varint,6,opt,name=newHeroHp,proto3" json:"newHeroHp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameOutput) Reset()         { *m = GameOutput{} }
func (m *GameOutput) String() string { return proto.CompactTextString(m) }
func (*GameOutput) ProtoMessage()    {}
func (*GameOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *GameOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameOutput.Unmarshal(m, b)
}
func (m *GameOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameOutput.Marshal(b, m, deterministic)
}
func (m *GameOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameOutput.Merge(m, src)
}
func (m *GameOutput) XXX_Size() int {
	return xxx_messageInfo_GameOutput.Size(m)
}
func (m *GameOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GameOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GameOutput proto.InternalMessageInfo

func (m *GameOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GameOutput) GetNewItem() string {
	if m != nil {
		return m.NewItem
	}
	return ""
}

func (m *GameOutput) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *GameOutput) GetModelNamespace() string {
	if m != nil {
		return m.ModelNamespace
	}
	return ""
}

func (m *GameOutput) GetNewEnemyHp() int32 {
	if m != nil {
		return m.NewEnemyHp
	}
	return 0
}

func (m *GameOutput) GetNewHeroHp() int32 {
	if m != nil {
		return m.NewHeroHp
	}
	return 0
}

func init() {
	proto.RegisterType((*GameInput)(nil), "proto.GameInput")
	proto.RegisterType((*GameOutput)(nil), "proto.GameOutput")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x4d, 0xdb, 0x34, 0xcd, 0xa5, 0xf8, 0x33, 0x0b, 0x19, 0x44, 0x4a, 0xc9, 0x42, 0xbb,
	0x0a, 0xa2, 0x7b, 0x17, 0x85, 0xd2, 0x74, 0xa3, 0xa5, 0xf8, 0x02, 0x63, 0xbc, 0xc4, 0xa0, 0xf3,
	0xc3, 0xcc, 0x94, 0xe2, 0x33, 0xf8, 0x16, 0x3e, 0x8b, 0x0f, 0x26, 0x33, 0x93, 0x36, 0xad, 0x08,
	0xae, 0x92, 0xf3, 0xdd, 0x93, 0xcc, 0x9c, 0x73, 0x01, 0x2a, 0xc6, 0x31, 0x57, 0x5a, 0x5a, 0x49,
	0x62, 0xff, 0xc8, 0x3e, 0x3b, 0x90, 0xce, 0x19, 0xc7, 0x85, 0x50, 0x6b, 0x4b, 0xce, 0xa1, 0xcf,
	0x4a, 0x5b, 0x4b, 0x41, 0xa3, 0x71, 0x34, 0x49, 0x57, 0x8d, 0x22, 0x04, 0x7a, 0xb5, 0x45, 0x4e,
	0x3b, 0x9e, 0xfa, 0x77, 0xe7, 0xb5, 0x4c, 0x57, 0x68, 0x69, 0x37, 0x78, 0x83, 0x22, 0x97, 0x90,
	0x72, 0xf9, 0x82, 0xef, 0x0f, 0x8c, 0x23, 0xed, 0xf9, 0x51, 0x0b, 0xc8, 0x15, 0x1c, 0xef, 0x84,
	0x51, 0xac, 0x44, 0x1a, 0x7b, 0xcb, 0x2f, 0x4a, 0x28, 0x24, 0x28, 0x90, 0x7f, 0x14, 0x8a, 0xf6,
	0xc7, 0xd1, 0x24, 0x5e, 0x6d, 0xa5, 0x3b, 0xf7, 0x15, 0xb5, 0x2c, 0x14, 0x4d, 0xfc, 0xa0, 0x51,
	0x24, 0x83, 0x61, 0x2d, 0x2c, 0xea, 0x70, 0x65, 0x43, 0x07, 0x7e, 0x7a, 0xc0, 0xc8, 0x08, 0xa0,
	0x36, 0x05, 0x6a, 0xf9, 0xb4, 0xd6, 0x82, 0xa6, 0xe3, 0x68, 0x32, 0x58, 0xed, 0x91, 0xec, 0x3b,
	0x02, 0x70, 0x6d, 0x3c, 0xae, 0xad, 0xab, 0x83, 0x42, 0xc2, 0xd1, 0x18, 0x56, 0x61, 0xd3, 0xc7,
	0x56, 0xba, 0x89, 0xc0, 0xcd, 0xa2, 0xed, 0x64, 0x2b, 0x0f, 0xe3, 0x77, 0xff, 0x8f, 0xdf, 0xfb,
	0x33, 0xfe, 0x08, 0x40, 0xe0, 0x66, 0xd6, 0x34, 0x10, 0xfb, 0x28, 0x7b, 0xc4, 0x9d, 0x22, 0x70,
	0x53, 0x84, 0x1e, 0x42, 0x41, 0x2d, 0xb8, 0xbd, 0x0f, 0x29, 0x66, 0xa2, 0xaa, 0x05, 0x92, 0x1b,
	0x48, 0x94, 0x96, 0x25, 0x1a, 0x43, 0x4e, 0xc3, 0xf2, 0xf3, 0xdd, 0xc6, 0x2f, 0xce, 0xf6, 0x48,
	0x48, 0x9d, 0x1d, 0x4d, 0xaf, 0x61, 0x28, 0x75, 0x95, 0xbf, 0xd5, 0x98, 0x57, 0x5a, 0x95, 0xd3,
	0x93, 0xf6, 0x6f, 0x4b, 0xe7, 0x5e, 0x46, 0x5f, 0x9d, 0xee, 0x7c, 0xb6, 0x7c, 0xee, 0xfb, 0x8f,
	0xef, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xef, 0x1b, 0xb5, 0x59, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameEngineClient is the client API for GameEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameEngineClient interface {
	Process(ctx context.Context, in *GameInput, opts ...grpc.CallOption) (*GameOutput, error)
}

type gameEngineClient struct {
	cc *grpc.ClientConn
}

func NewGameEngineClient(cc *grpc.ClientConn) GameEngineClient {
	return &gameEngineClient{cc}
}

func (c *gameEngineClient) Process(ctx context.Context, in *GameInput, opts ...grpc.CallOption) (*GameOutput, error) {
	out := new(GameOutput)
	err := c.cc.Invoke(ctx, "/proto.GameEngine/process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameEngineServer is the server API for GameEngine service.
type GameEngineServer interface {
	Process(context.Context, *GameInput) (*GameOutput, error)
}

// UnimplementedGameEngineServer can be embedded to have forward compatible implementations.
type UnimplementedGameEngineServer struct {
}

func (*UnimplementedGameEngineServer) Process(ctx context.Context, req *GameInput) (*GameOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterGameEngineServer(s *grpc.Server, srv GameEngineServer) {
	s.RegisterService(&_GameEngine_serviceDesc, srv)
}

func _GameEngine_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameEngineServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameEngine/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameEngineServer).Process(ctx, req.(*GameInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GameEngine",
	HandlerType: (*GameEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "process",
			Handler:    _GameEngine_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
