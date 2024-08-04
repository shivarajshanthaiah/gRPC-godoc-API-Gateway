// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: admin.proto

package __

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

type AdminNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminNoParam) Reset() {
	*x = AdminNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminNoParam) ProtoMessage() {}

func (x *AdminNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminNoParam.ProtoReflect.Descriptor instead.
func (*AdminNoParam) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Dob       string `protobuf:"bytes,4,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender    string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Address   string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *UserList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type AdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AdminResponse) Reset() {
	*x = AdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResponse) ProtoMessage() {}

func (x *AdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResponse.ProtoReflect.Descriptor instead.
func (*AdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdminResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AdminResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AdminResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *AdminRequest) Reset() {
	*x = AdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRequest) ProtoMessage() {}

func (x *AdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRequest.ProtoReflect.Descriptor instead.
func (*AdminRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
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
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AUserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AUserID) Reset() {
	*x = AUserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AUserID) ProtoMessage() {}

func (x *AUserID) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AUserID.ProtoReflect.Descriptor instead.
func (*AUserID) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{6}
}

func (x *AUserID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Dob       string `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Password  string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{7}
}

func (x *UserModel) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserModel) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserModel) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *UserModel) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserModel) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DoctorModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorName string `protobuf:"bytes,1,opt,name=doctor_name,json=doctorName,proto3" json:"doctor_name,omitempty"`
	Age        uint64 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Speciality string `protobuf:"bytes,3,opt,name=speciality,proto3" json:"speciality,omitempty"`
	Gender     string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DoctorModel) Reset() {
	*x = DoctorModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorModel) ProtoMessage() {}

func (x *DoctorModel) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorModel.ProtoReflect.Descriptor instead.
func (*DoctorModel) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{8}
}

func (x *DoctorModel) GetDoctorName() string {
	if x != nil {
		return x.DoctorName
	}
	return ""
}

func (x *DoctorModel) GetAge() uint64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *DoctorModel) GetSpeciality() string {
	if x != nil {
		return x.Speciality
	}
	return ""
}

func (x *DoctorModel) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *DoctorModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DoctorID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DoctorID) Reset() {
	*x = DoctorID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorID) ProtoMessage() {}

func (x *DoctorID) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorID.ProtoReflect.Descriptor instead.
func (*DoctorID) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{9}
}

func (x *DoctorID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DoctorName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DoctorName) Reset() {
	*x = DoctorName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorName) ProtoMessage() {}

func (x *DoctorName) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorName.ProtoReflect.Descriptor instead.
func (*DoctorName) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{10}
}

func (x *DoctorName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DoctorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctors []*DoctorModel `protobuf:"bytes,1,rep,name=doctors,proto3" json:"doctors,omitempty"`
}

func (x *DoctorList) Reset() {
	*x = DoctorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorList) ProtoMessage() {}

func (x *DoctorList) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorList.ProtoReflect.Descriptor instead.
func (*DoctorList) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{11}
}

func (x *DoctorList) GetDoctors() []*DoctorModel {
	if x != nil {
		return x.Doctors
	}
	return nil
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0xcc, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x5a, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x19, 0x0a, 0x07, 0x41, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x20, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x32, 0xe9, 0x04, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x46, 0x6e, 0x12, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0e, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x46, 0x6e, 0x12, 0x10, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x46, 0x6e, 0x12,
	0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x45, 0x64, 0x69,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x46, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x13, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x15, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x3d, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_admin_proto_goTypes = []interface{}{
	(*AdminNoParam)(nil),  // 0: adminpb.AdminNoParam
	(*User)(nil),          // 1: adminpb.User
	(*UserList)(nil),      // 2: adminpb.UserList
	(*AdminResponse)(nil), // 3: adminpb.AdminResponse
	(*AdminRequest)(nil),  // 4: adminpb.AdminRequest
	(*UserRequest)(nil),   // 5: adminpb.UserRequest
	(*AUserID)(nil),       // 6: adminpb.AUserID
	(*UserModel)(nil),     // 7: adminpb.UserModel
	(*DoctorModel)(nil),   // 8: adminpb.DoctorModel
	(*DoctorID)(nil),      // 9: adminpb.DoctorID
	(*DoctorName)(nil),    // 10: adminpb.DoctorName
	(*DoctorList)(nil),    // 11: adminpb.DoctorList
}
var file_admin_proto_depIdxs = []int32{
	1,  // 0: adminpb.UserList.users:type_name -> adminpb.User
	8,  // 1: adminpb.DoctorList.doctors:type_name -> adminpb.DoctorModel
	4,  // 2: adminpb.AdminService.AdminLogin:input_type -> adminpb.AdminRequest
	5,  // 3: adminpb.AdminService.FindUserByEmailFn:input_type -> adminpb.UserRequest
	6,  // 4: adminpb.AdminService.FindUserByIDFn:input_type -> adminpb.AUserID
	0,  // 5: adminpb.AdminService.FindAllUsersFn:input_type -> adminpb.AdminNoParam
	7,  // 6: adminpb.AdminService.EditUserFn:input_type -> adminpb.UserModel
	6,  // 7: adminpb.AdminService.DeleteUserFn:input_type -> adminpb.AUserID
	8,  // 8: adminpb.AdminService.CreateDoctor:input_type -> adminpb.DoctorModel
	9,  // 9: adminpb.AdminService.UserFetchDoctorByID:input_type -> adminpb.DoctorID
	10, // 10: adminpb.AdminService.UserFetchDoctorByName:input_type -> adminpb.DoctorName
	0,  // 11: adminpb.AdminService.FetchAllDoctors:input_type -> adminpb.AdminNoParam
	3,  // 12: adminpb.AdminService.AdminLogin:output_type -> adminpb.AdminResponse
	1,  // 13: adminpb.AdminService.FindUserByEmailFn:output_type -> adminpb.User
	1,  // 14: adminpb.AdminService.FindUserByIDFn:output_type -> adminpb.User
	2,  // 15: adminpb.AdminService.FindAllUsersFn:output_type -> adminpb.UserList
	3,  // 16: adminpb.AdminService.EditUserFn:output_type -> adminpb.AdminResponse
	3,  // 17: adminpb.AdminService.DeleteUserFn:output_type -> adminpb.AdminResponse
	3,  // 18: adminpb.AdminService.CreateDoctor:output_type -> adminpb.AdminResponse
	8,  // 19: adminpb.AdminService.UserFetchDoctorByID:output_type -> adminpb.DoctorModel
	8,  // 20: adminpb.AdminService.UserFetchDoctorByName:output_type -> adminpb.DoctorModel
	11, // 21: adminpb.AdminService.FetchAllDoctors:output_type -> adminpb.DoctorList
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminNoParam); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminResponse); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRequest); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AUserID); i {
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
		file_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorModel); i {
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
		file_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorID); i {
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
		file_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorName); i {
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
		file_admin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorList); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
