// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/proto/config_service/config_service.proto

package config_service

import (
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

type EnableGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *EnableGamesRequest) Reset() {
	*x = EnableGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableGamesRequest) ProtoMessage() {}

func (x *EnableGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableGamesRequest.ProtoReflect.Descriptor instead.
func (*EnableGamesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnableGamesRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type SetFEHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *SetFEHashRequest) Reset() {
	*x = SetFEHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFEHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFEHashRequest) ProtoMessage() {}

func (x *SetFEHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFEHashRequest.ProtoReflect.Descriptor instead.
func (*SetFEHashRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetFEHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type PermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Director bool   `protobuf:"varint,2,opt,name=director,proto3" json:"director,omitempty"`
	Admin    bool   `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	Mod      bool   `protobuf:"varint,4,opt,name=mod,proto3" json:"mod,omitempty"`
	Bot      bool   `protobuf:"varint,5,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (x *PermissionsRequest) Reset() {
	*x = PermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionsRequest) ProtoMessage() {}

func (x *PermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionsRequest.ProtoReflect.Descriptor instead.
func (*PermissionsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PermissionsRequest) GetDirector() bool {
	if x != nil {
		return x.Director
	}
	return false
}

func (x *PermissionsRequest) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *PermissionsRequest) GetMod() bool {
	if x != nil {
		return x.Mod
	}
	return false
}

func (x *PermissionsRequest) GetBot() bool {
	if x != nil {
		return x.Bot
	}
	return false
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Uuid       string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	IsBot      bool   `protobuf:"varint,4,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	IsDirector bool   `protobuf:"varint,5,opt,name=is_director,json=isDirector,proto3" json:"is_director,omitempty"`
	IsMod      bool   `protobuf:"varint,6,opt,name=is_mod,json=isMod,proto3" json:"is_mod,omitempty"`
	IsAdmin    bool   `protobuf:"varint,7,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetIsBot() bool {
	if x != nil {
		return x.IsBot
	}
	return false
}

func (x *UserResponse) GetIsDirector() bool {
	if x != nil {
		return x.IsDirector
	}
	return false
}

func (x *UserResponse) GetIsMod() bool {
	if x != nil {
		return x.IsMod
	}
	return false
}

func (x *UserResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{5}
}

type Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link  string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Announcement) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Announcement) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type SetAnnouncementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Announcements []*Announcement `protobuf:"bytes,1,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *SetAnnouncementsRequest) Reset() {
	*x = SetAnnouncementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAnnouncementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAnnouncementsRequest) ProtoMessage() {}

func (x *SetAnnouncementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAnnouncementsRequest.ProtoReflect.Descriptor instead.
func (*SetAnnouncementsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{7}
}

func (x *SetAnnouncementsRequest) GetAnnouncements() []*Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

type GetAnnouncementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAnnouncementsRequest) Reset() {
	*x = GetAnnouncementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnouncementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnouncementsRequest) ProtoMessage() {}

func (x *GetAnnouncementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnouncementsRequest.ProtoReflect.Descriptor instead.
func (*GetAnnouncementsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{8}
}

type AnnouncementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Announcements []*Announcement `protobuf:"bytes,1,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *AnnouncementsResponse) Reset() {
	*x = AnnouncementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_config_service_config_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementsResponse) ProtoMessage() {}

func (x *AnnouncementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_config_service_config_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementsResponse.ProtoReflect.Descriptor instead.
func (*AnnouncementsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_config_service_config_service_proto_rawDescGZIP(), []int{9}
}

func (x *AnnouncementsResponse) GetAnnouncements() []*Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

var File_api_proto_config_service_config_service_proto protoreflect.FileDescriptor

var file_api_proto_config_service_config_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x2e, 0x0a, 0x12, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0x26, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x46, 0x45, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6d, 0x6f, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x6f, 0x74,
	0x22, 0x29, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73,
	0x5f, 0x6d, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4d, 0x6f,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x10, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c,
	0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x5d, 0x0a, 0x17,
	0x53, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x15, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0x9d, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x46, 0x45, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x45,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x12, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x62, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x6f, 0x31, 0x34, 0x2f, 0x6c, 0x69, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_config_service_config_service_proto_rawDescOnce sync.Once
	file_api_proto_config_service_config_service_proto_rawDescData = file_api_proto_config_service_config_service_proto_rawDesc
)

func file_api_proto_config_service_config_service_proto_rawDescGZIP() []byte {
	file_api_proto_config_service_config_service_proto_rawDescOnce.Do(func() {
		file_api_proto_config_service_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_config_service_config_service_proto_rawDescData)
	})
	return file_api_proto_config_service_config_service_proto_rawDescData
}

var file_api_proto_config_service_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_config_service_config_service_proto_goTypes = []interface{}{
	(*EnableGamesRequest)(nil),      // 0: config_service.EnableGamesRequest
	(*SetFEHashRequest)(nil),        // 1: config_service.SetFEHashRequest
	(*PermissionsRequest)(nil),      // 2: config_service.PermissionsRequest
	(*UserRequest)(nil),             // 3: config_service.UserRequest
	(*UserResponse)(nil),            // 4: config_service.UserResponse
	(*ConfigResponse)(nil),          // 5: config_service.ConfigResponse
	(*Announcement)(nil),            // 6: config_service.Announcement
	(*SetAnnouncementsRequest)(nil), // 7: config_service.SetAnnouncementsRequest
	(*GetAnnouncementsRequest)(nil), // 8: config_service.GetAnnouncementsRequest
	(*AnnouncementsResponse)(nil),   // 9: config_service.AnnouncementsResponse
}
var file_api_proto_config_service_config_service_proto_depIdxs = []int32{
	6, // 0: config_service.SetAnnouncementsRequest.announcements:type_name -> config_service.Announcement
	6, // 1: config_service.AnnouncementsResponse.announcements:type_name -> config_service.Announcement
	0, // 2: config_service.ConfigService.SetGamesEnabled:input_type -> config_service.EnableGamesRequest
	1, // 3: config_service.ConfigService.SetFEHash:input_type -> config_service.SetFEHashRequest
	2, // 4: config_service.ConfigService.SetUserPermissions:input_type -> config_service.PermissionsRequest
	3, // 5: config_service.ConfigService.GetUserDetails:input_type -> config_service.UserRequest
	7, // 6: config_service.ConfigService.SetAnnouncements:input_type -> config_service.SetAnnouncementsRequest
	8, // 7: config_service.ConfigService.GetAnnouncements:input_type -> config_service.GetAnnouncementsRequest
	5, // 8: config_service.ConfigService.SetGamesEnabled:output_type -> config_service.ConfigResponse
	5, // 9: config_service.ConfigService.SetFEHash:output_type -> config_service.ConfigResponse
	5, // 10: config_service.ConfigService.SetUserPermissions:output_type -> config_service.ConfigResponse
	4, // 11: config_service.ConfigService.GetUserDetails:output_type -> config_service.UserResponse
	5, // 12: config_service.ConfigService.SetAnnouncements:output_type -> config_service.ConfigResponse
	9, // 13: config_service.ConfigService.GetAnnouncements:output_type -> config_service.AnnouncementsResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_config_service_config_service_proto_init() }
func file_api_proto_config_service_config_service_proto_init() {
	if File_api_proto_config_service_config_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_config_service_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableGamesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFEHashRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAnnouncementsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnouncementsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_config_service_config_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_config_service_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_config_service_config_service_proto_goTypes,
		DependencyIndexes: file_api_proto_config_service_config_service_proto_depIdxs,
		MessageInfos:      file_api_proto_config_service_config_service_proto_msgTypes,
	}.Build()
	File_api_proto_config_service_config_service_proto = out.File
	file_api_proto_config_service_config_service_proto_rawDesc = nil
	file_api_proto_config_service_config_service_proto_goTypes = nil
	file_api_proto_config_service_config_service_proto_depIdxs = nil
}
