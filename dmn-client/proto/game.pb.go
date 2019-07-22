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
	EnemyHp              int32    `protobuf:"varint,4,opt,name=enemyHp,proto3" json:"enemyHp,omitempty"`
	HeroHp               int32    `protobuf:"varint,5,opt,name=heroHp,proto3" json:"heroHp,omitempty"`
	Interactions         int32    `protobuf:"varint,6,opt,name=interactions,proto3" json:"interactions,omitempty"`
	IsHeroTurn           bool     `protobuf:"varint,7,opt,name=isHeroTurn,proto3" json:"isHeroTurn,omitempty"`
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
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x8f, 0x7b, 0xcb, 0xc9, 0xa8, 0xe2, 0xe2, 0x05, 0xb2, 0x10, 0x42, 0x55, 0x16, 0xd0,
	0x55, 0x84, 0x60, 0xcf, 0xa2, 0x52, 0xd5, 0x74, 0x03, 0x55, 0xc4, 0x0b, 0x98, 0x30, 0x0a, 0x11,
	0xf8, 0x22, 0xdb, 0x55, 0xc5, 0xeb, 0xf0, 0x2c, 0xf0, 0x5e, 0xc8, 0x76, 0xda, 0x04, 0x56, 0xc9,
	0xff, 0xcd, 0x38, 0xce, 0x37, 0x03, 0x50, 0x73, 0x81, 0xb9, 0x36, 0xca, 0x29, 0x3a, 0x0e, 0x8f,
	0xec, 0x9b, 0x40, 0xba, 0xe2, 0x02, 0xd7, 0x52, 0x6f, 0x1d, 0x3d, 0x83, 0x09, 0xaf, 0x5c, 0xa3,
	0x24, 0x23, 0x33, 0x32, 0x4f, 0xcb, 0x36, 0x51, 0x0a, 0xa3, 0xc6, 0xa1, 0x60, 0x83, 0x40, 0xc3,
	0xbb, 0xef, 0x75, 0xdc, 0xd4, 0xe8, 0xd8, 0x30, 0xf6, 0xc6, 0x44, 0x19, 0x24, 0x28, 0x51, 0x7c,
	0x14, 0x9a, 0x8d, 0x66, 0x64, 0x3e, 0x2e, 0xf7, 0xd1, 0x9f, 0x78, 0x45, 0xa3, 0x0a, 0xcd, 0xc6,
	0xa1, 0xd0, 0x26, 0x9a, 0xc1, 0xb4, 0x91, 0x0e, 0x4d, 0xbc, 0xcc, 0xb2, 0x49, 0xa8, 0xfe, 0x62,
	0xf4, 0x12, 0xa0, 0xb1, 0x05, 0x1a, 0xf5, 0xb4, 0x35, 0x92, 0x25, 0x33, 0x32, 0xff, 0x5f, 0xf6,
	0x48, 0xf6, 0x45, 0x00, 0xbc, 0xc7, 0xe3, 0xd6, 0x79, 0x11, 0x06, 0x89, 0x40, 0x6b, 0x79, 0x8d,
	0xad, 0xc9, 0x3e, 0xfa, 0x8a, 0xc4, 0xdd, 0xba, 0xb3, 0xd9, 0x47, 0x7a, 0x01, 0xa9, 0x50, 0x2f,
	0xf8, 0xfe, 0xc0, 0x05, 0xb6, 0x4e, 0x1d, 0xa0, 0x57, 0x70, 0x74, 0x08, 0x56, 0xf3, 0x0a, 0x83,
	0x5d, 0x5a, 0xfe, 0xa1, 0xfe, 0x47, 0x25, 0xee, 0x96, 0xed, 0x04, 0xa2, 0x68, 0x8f, 0xf8, 0x5b,
	0x24, 0xee, 0x8a, 0x38, 0x87, 0x68, 0xda, 0x81, 0xdb, 0xfb, 0x68, 0xb1, 0x94, 0x75, 0x23, 0x91,
	0xde, 0x40, 0xa2, 0x8d, 0xaa, 0xd0, 0x5a, 0x7a, 0x12, 0xd7, 0x96, 0x1f, 0x76, 0x75, 0x7e, 0xda,
	0x23, 0xd1, 0x3a, 0xfb, 0xb7, 0xb8, 0x86, 0xa9, 0x32, 0x75, 0xfe, 0xd6, 0x60, 0x5e, 0x1b, 0x5d,
	0x2d, 0x8e, 0xbb, 0xaf, 0x6d, 0x7c, 0xf7, 0x86, 0x7c, 0x0e, 0x86, 0xab, 0xe5, 0xe6, 0x79, 0x12,
	0x0e, 0xdf, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x4b, 0xb5, 0x58, 0x13, 0x02, 0x00, 0x00,
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