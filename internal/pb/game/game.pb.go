// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: game.proto

package game

import (
	models "github.com/tsunakit99/ankylo-cup-backend/internal/pb/models"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGameListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGameListRequest) Reset() {
	*x = GetGameListRequest{}
	mi := &file_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameListRequest) ProtoMessage() {}

func (x *GetGameListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameListRequest.ProtoReflect.Descriptor instead.
func (*GetGameListRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type GetGameListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*models.Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *GetGameListResponse) Reset() {
	*x = GetGameListResponse{}
	mi := &file_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameListResponse) ProtoMessage() {}

func (x *GetGameListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameListResponse.ProtoReflect.Descriptor instead.
func (*GetGameListResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *GetGameListResponse) GetGames() []*models.Game {
	if x != nil {
		return x.Games
	}
	return nil
}

type GetGameByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId int32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *GetGameByIdRequest) Reset() {
	*x = GetGameByIdRequest{}
	mi := &file_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameByIdRequest) ProtoMessage() {}

func (x *GetGameByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameByIdRequest.ProtoReflect.Descriptor instead.
func (*GetGameByIdRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *GetGameByIdRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type GetGameByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *models.Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *GetGameByIdResponse) Reset() {
	*x = GetGameByIdResponse{}
	mi := &file_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameByIdResponse) ProtoMessage() {}

func (x *GetGameByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameByIdResponse.ProtoReflect.Descriptor instead.
func (*GetGameByIdResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *GetGameByIdResponse) GetGame() *models.Game {
	if x != nil {
		return x.Game
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x73, 0x75, 0x6e, 0x61, 0x6b, 0x69, 0x74, 0x39, 0x39, 0x2f, 0x61, 0x6e, 0x6b, 0x79, 0x6c,
	0x6f, 0x2d, 0x63, 0x75, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_proto_goTypes = []any{
	(*GetGameListRequest)(nil),  // 0: game.GetGameListRequest
	(*GetGameListResponse)(nil), // 1: game.GetGameListResponse
	(*GetGameByIdRequest)(nil),  // 2: game.GetGameByIdRequest
	(*GetGameByIdResponse)(nil), // 3: game.GetGameByIdResponse
	(*models.Game)(nil),         // 4: models.Game
}
var file_game_proto_depIdxs = []int32{
	4, // 0: game.GetGameListResponse.games:type_name -> models.Game
	4, // 1: game.GetGameByIdResponse.game:type_name -> models.Game
	0, // 2: game.GameService.GetGameList:input_type -> game.GetGameListRequest
	2, // 3: game.GameService.GetGameById:input_type -> game.GetGameByIdRequest
	1, // 4: game.GameService.GetGameList:output_type -> game.GetGameListResponse
	3, // 5: game.GameService.GetGameById:output_type -> game.GetGameByIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
