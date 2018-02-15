// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/user.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	ID         string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Username   string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	Details    string `protobuf:"bytes,6,opt,name=Details" json:"Details,omitempty"`
	UserType   string `protobuf:"bytes,7,opt,name=UserType" json:"UserType,omitempty"`
	UserStatus string `protobuf:"bytes,8,opt,name=UserStatus" json:"UserStatus,omitempty"`
	CreatedAt  string `protobuf:"bytes,9,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt  string `protobuf:"bytes,10,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *User) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *User) GetUserStatus() string {
	if m != nil {
		return m.UserStatus
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ReqRegistration struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
	Details  string `protobuf:"bytes,5,opt,name=Details" json:"Details,omitempty"`
}

func (m *ReqRegistration) Reset()                    { *m = ReqRegistration{} }
func (m *ReqRegistration) String() string            { return proto.CompactTextString(m) }
func (*ReqRegistration) ProtoMessage()               {}
func (*ReqRegistration) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *ReqRegistration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqRegistration) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqRegistration) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReqRegistration) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ReqRegistration) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type ResRegistration struct {
	User   *User    `protobuf:"bytes,1,opt,name=User" json:"User,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResRegistration) Reset()                    { *m = ResRegistration{} }
func (m *ResRegistration) String() string            { return proto.CompactTextString(m) }
func (*ResRegistration) ProtoMessage()               {}
func (*ResRegistration) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *ResRegistration) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResRegistration) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqLogin struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *ReqLogin) Reset()                    { *m = ReqLogin{} }
func (m *ReqLogin) String() string            { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()               {}
func (*ReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *ReqLogin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ResLogin struct {
	Session *Session `protobuf:"bytes,1,opt,name=Session" json:"Session,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResLogin) Reset()                    { *m = ResLogin{} }
func (m *ResLogin) String() string            { return proto.CompactTextString(m) }
func (*ResLogin) ProtoMessage()               {}
func (*ResLogin) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *ResLogin) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ResLogin) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqProfile struct {
	UserID string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *ReqProfile) Reset()                    { *m = ReqProfile{} }
func (m *ReqProfile) String() string            { return proto.CompactTextString(m) }
func (*ReqProfile) ProtoMessage()               {}
func (*ReqProfile) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *ReqProfile) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type ResProfile struct {
	User   *User    `protobuf:"bytes,1,opt,name=User" json:"User,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResProfile) Reset()                    { *m = ResProfile{} }
func (m *ResProfile) String() string            { return proto.CompactTextString(m) }
func (*ResProfile) ProtoMessage()               {}
func (*ResProfile) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *ResProfile) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResProfile) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqUpdateUser struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=Details" json:"Details,omitempty"`
}

func (m *ReqUpdateUser) Reset()                    { *m = ReqUpdateUser{} }
func (m *ReqUpdateUser) String() string            { return proto.CompactTextString(m) }
func (*ReqUpdateUser) ProtoMessage()               {}
func (*ReqUpdateUser) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *ReqUpdateUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqUpdateUser) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type ResUpdateUser struct {
	User   *User    `protobuf:"bytes,1,opt,name=User" json:"User,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResUpdateUser) Reset()                    { *m = ResUpdateUser{} }
func (m *ResUpdateUser) String() string            { return proto.CompactTextString(m) }
func (*ResUpdateUser) ProtoMessage()               {}
func (*ResUpdateUser) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ResUpdateUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResUpdateUser) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqChangePassword struct {
	OldPassword string `protobuf:"bytes,1,opt,name=OldPassword" json:"OldPassword,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=NewPassword" json:"NewPassword,omitempty"`
}

func (m *ReqChangePassword) Reset()                    { *m = ReqChangePassword{} }
func (m *ReqChangePassword) String() string            { return proto.CompactTextString(m) }
func (*ReqChangePassword) ProtoMessage()               {}
func (*ReqChangePassword) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *ReqChangePassword) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ReqChangePassword) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ResChangePassword struct {
	Success bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResChangePassword) Reset()                    { *m = ResChangePassword{} }
func (m *ResChangePassword) String() string            { return proto.CompactTextString(m) }
func (*ResChangePassword) ProtoMessage()               {}
func (*ResChangePassword) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *ResChangePassword) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResChangePassword) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqChangeUserStatus struct {
	NewStatus string `protobuf:"bytes,1,opt,name=NewStatus" json:"NewStatus,omitempty"`
}

func (m *ReqChangeUserStatus) Reset()                    { *m = ReqChangeUserStatus{} }
func (m *ReqChangeUserStatus) String() string            { return proto.CompactTextString(m) }
func (*ReqChangeUserStatus) ProtoMessage()               {}
func (*ReqChangeUserStatus) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

func (m *ReqChangeUserStatus) GetNewStatus() string {
	if m != nil {
		return m.NewStatus
	}
	return ""
}

type ResChangeUserStatus struct {
	Success bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResChangeUserStatus) Reset()                    { *m = ResChangeUserStatus{} }
func (m *ResChangeUserStatus) String() string            { return proto.CompactTextString(m) }
func (*ResChangeUserStatus) ProtoMessage()               {}
func (*ResChangeUserStatus) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

func (m *ResChangeUserStatus) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResChangeUserStatus) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqChangeUserType struct {
	UserType string `protobuf:"bytes,1,opt,name=UserType" json:"UserType,omitempty"`
}

func (m *ReqChangeUserType) Reset()                    { *m = ReqChangeUserType{} }
func (m *ReqChangeUserType) String() string            { return proto.CompactTextString(m) }
func (*ReqChangeUserType) ProtoMessage()               {}
func (*ReqChangeUserType) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *ReqChangeUserType) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

type ResChangeUserType struct {
	Success bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResChangeUserType) Reset()                    { *m = ResChangeUserType{} }
func (m *ResChangeUserType) String() string            { return proto.CompactTextString(m) }
func (*ResChangeUserType) ProtoMessage()               {}
func (*ResChangeUserType) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{14} }

func (m *ResChangeUserType) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResChangeUserType) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqUserLogout struct {
	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken" json:"AccessToken,omitempty"`
}

func (m *ReqUserLogout) Reset()                    { *m = ReqUserLogout{} }
func (m *ReqUserLogout) String() string            { return proto.CompactTextString(m) }
func (*ReqUserLogout) ProtoMessage()               {}
func (*ReqUserLogout) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{15} }

func (m *ReqUserLogout) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ResUserLogout struct {
	Success bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResUserLogout) Reset()                    { *m = ResUserLogout{} }
func (m *ResUserLogout) String() string            { return proto.CompactTextString(m) }
func (*ResUserLogout) ProtoMessage()               {}
func (*ResUserLogout) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{16} }

func (m *ResUserLogout) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResUserLogout) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqResetPasswordRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *ReqResetPasswordRequest) Reset()                    { *m = ReqResetPasswordRequest{} }
func (m *ReqResetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*ReqResetPasswordRequest) ProtoMessage()               {}
func (*ReqResetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{17} }

func (m *ReqResetPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ResResetPasswordRequest struct {
	Success bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResResetPasswordRequest) Reset()                    { *m = ResResetPasswordRequest{} }
func (m *ResResetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*ResResetPasswordRequest) ProtoMessage()               {}
func (*ResResetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{18} }

func (m *ResResetPasswordRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResResetPasswordRequest) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ReqResetPassword struct {
	PasswordResetToken string `protobuf:"bytes,1,opt,name=PasswordResetToken" json:"PasswordResetToken,omitempty"`
	NewPassword        string `protobuf:"bytes,2,opt,name=NewPassword" json:"NewPassword,omitempty"`
}

func (m *ReqResetPassword) Reset()                    { *m = ReqResetPassword{} }
func (m *ReqResetPassword) String() string            { return proto.CompactTextString(m) }
func (*ReqResetPassword) ProtoMessage()               {}
func (*ReqResetPassword) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{19} }

func (m *ReqResetPassword) GetPasswordResetToken() string {
	if m != nil {
		return m.PasswordResetToken
	}
	return ""
}

func (m *ReqResetPassword) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ResResetPassword struct {
	Success bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Errors  []*Error `protobuf:"bytes,2,rep,name=Errors" json:"Errors,omitempty"`
}

func (m *ResResetPassword) Reset()                    { *m = ResResetPassword{} }
func (m *ResResetPassword) String() string            { return proto.CompactTextString(m) }
func (*ResResetPassword) ProtoMessage()               {}
func (*ResResetPassword) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{20} }

func (m *ResResetPassword) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResResetPassword) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "protos.User")
	proto.RegisterType((*ReqRegistration)(nil), "protos.ReqRegistration")
	proto.RegisterType((*ResRegistration)(nil), "protos.ResRegistration")
	proto.RegisterType((*ReqLogin)(nil), "protos.ReqLogin")
	proto.RegisterType((*ResLogin)(nil), "protos.ResLogin")
	proto.RegisterType((*ReqProfile)(nil), "protos.ReqProfile")
	proto.RegisterType((*ResProfile)(nil), "protos.ResProfile")
	proto.RegisterType((*ReqUpdateUser)(nil), "protos.ReqUpdateUser")
	proto.RegisterType((*ResUpdateUser)(nil), "protos.ResUpdateUser")
	proto.RegisterType((*ReqChangePassword)(nil), "protos.ReqChangePassword")
	proto.RegisterType((*ResChangePassword)(nil), "protos.ResChangePassword")
	proto.RegisterType((*ReqChangeUserStatus)(nil), "protos.ReqChangeUserStatus")
	proto.RegisterType((*ResChangeUserStatus)(nil), "protos.ResChangeUserStatus")
	proto.RegisterType((*ReqChangeUserType)(nil), "protos.ReqChangeUserType")
	proto.RegisterType((*ResChangeUserType)(nil), "protos.ResChangeUserType")
	proto.RegisterType((*ReqUserLogout)(nil), "protos.ReqUserLogout")
	proto.RegisterType((*ResUserLogout)(nil), "protos.ResUserLogout")
	proto.RegisterType((*ReqResetPasswordRequest)(nil), "protos.ReqResetPasswordRequest")
	proto.RegisterType((*ResResetPasswordRequest)(nil), "protos.ResResetPasswordRequest")
	proto.RegisterType((*ReqResetPassword)(nil), "protos.ReqResetPassword")
	proto.RegisterType((*ResResetPassword)(nil), "protos.ResResetPassword")
}

func init() { proto.RegisterFile("protos/user.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0x34, 0x1f, 0x13, 0x4a, 0xe9, 0x12, 0x81, 0x15, 0xa1, 0x2a, 0xb2, 0x40, 0x2a,
	0x97, 0x54, 0xb4, 0xe2, 0xc8, 0xa1, 0x6d, 0x7a, 0xa8, 0x54, 0x85, 0x68, 0x93, 0x00, 0x8a, 0xb8,
	0xb8, 0xcd, 0x10, 0x4c, 0x53, 0x3b, 0xf6, 0xac, 0x15, 0xc1, 0x7f, 0xe0, 0xaf, 0xf2, 0x1b, 0xd0,
	0xee, 0xda, 0xce, 0x3a, 0xa4, 0x28, 0x55, 0x7a, 0x8a, 0xe7, 0xbd, 0x9d, 0xf1, 0x9b, 0x79, 0xe3,
	0x0d, 0xec, 0xcf, 0xe3, 0x50, 0x84, 0x74, 0x94, 0x10, 0xc6, 0x1d, 0xf5, 0xcc, 0x2a, 0x1a, 0x6a,
	0xb1, 0x94, 0xc2, 0x38, 0x0e, 0x53, 0xae, 0xd5, 0x4c, 0x31, 0x42, 0x22, 0x3f, 0x0c, 0x34, 0xea,
	0xfe, 0xb1, 0xa0, 0x3c, 0x22, 0x8c, 0xd9, 0x53, 0xb0, 0x2f, 0xbb, 0x8e, 0xd5, 0xb6, 0x0e, 0xeb,
	0xdc, 0xbe, 0xec, 0x32, 0x06, 0xe5, 0x9e, 0x77, 0x87, 0x8e, 0xad, 0x10, 0xf5, 0xcc, 0x5a, 0x50,
	0x93, 0x67, 0x03, 0x89, 0x97, 0x14, 0x9e, 0xc7, 0xac, 0x09, 0x3b, 0x17, 0x77, 0x9e, 0x3f, 0x73,
	0xca, 0x8a, 0xd0, 0x01, 0x73, 0xa0, 0xda, 0x45, 0xe1, 0xf9, 0x33, 0x72, 0x2a, 0x0a, 0xcf, 0xc2,
	0xac, 0xd6, 0xf0, 0xe7, 0x1c, 0x9d, 0xea, 0xb2, 0x96, 0x8c, 0xd9, 0x01, 0x80, 0x7c, 0x1e, 0x08,
	0x4f, 0x24, 0xe4, 0xd4, 0x14, 0x6b, 0x20, 0xec, 0x15, 0xd4, 0xcf, 0x63, 0xf4, 0x04, 0x4e, 0x4e,
	0x85, 0x53, 0x57, 0xf4, 0x12, 0x90, 0xec, 0x68, 0x3e, 0x49, 0x59, 0xd0, 0x6c, 0x0e, 0xb8, 0xbf,
	0x2d, 0xd8, 0xe3, 0x18, 0x71, 0x9c, 0xfa, 0x24, 0x62, 0x4f, 0xf8, 0x61, 0x90, 0xf7, 0x6a, 0xdd,
	0xd3, 0xab, 0x7d, 0x5f, 0xaf, 0x25, 0xb3, 0xd7, 0x16, 0xd4, 0xfa, 0x1e, 0xd1, 0x22, 0x8c, 0x27,
	0xe9, 0x10, 0xf2, 0xd8, 0x9c, 0xc3, 0x4e, 0x61, 0x0e, 0xee, 0x58, 0xca, 0xa1, 0x82, 0x9c, 0xb6,
	0xb6, 0x44, 0xc9, 0x69, 0x1c, 0x3f, 0xd1, 0x4e, 0x51, 0x47, 0x62, 0x5c, 0x9b, 0xf5, 0x06, 0x2a,
	0x17, 0xd2, 0x5a, 0x72, 0xec, 0x76, 0xe9, 0xb0, 0x71, 0xbc, 0x9b, 0x9d, 0x51, 0x28, 0x4f, 0x49,
	0xf7, 0x0c, 0x6a, 0x1c, 0xa3, 0xab, 0x70, 0xea, 0x07, 0x85, 0x7e, 0xac, 0x95, 0x7e, 0x4c, 0xe5,
	0x76, 0x51, 0xb9, 0xfb, 0x55, 0xd6, 0x20, 0x5d, 0xe3, 0x2d, 0x54, 0x07, 0x7a, 0x7b, 0x52, 0x6d,
	0x7b, 0xd9, 0x7b, 0x53, 0x98, 0x67, 0xfc, 0xa6, 0x0a, 0x5f, 0x03, 0x70, 0x8c, 0xfa, 0x71, 0xf8,
	0xcd, 0x9f, 0x21, 0x7b, 0x01, 0x15, 0xa9, 0x29, 0xdf, 0xc3, 0x34, 0x72, 0x47, 0xf2, 0x14, 0x65,
	0xa7, 0x1e, 0x6d, 0x3c, 0x1f, 0x60, 0x97, 0x63, 0xa4, 0x57, 0x43, 0xe5, 0xad, 0xdb, 0x03, 0xc3,
	0x39, 0xbb, 0xe8, 0xdc, 0x17, 0x99, 0x4e, 0x46, 0xfa, 0xa3, 0x09, 0xfb, 0x0c, 0xfb, 0x1c, 0xa3,
	0xf3, 0xef, 0x5e, 0x30, 0xc5, 0x7c, 0x85, 0xda, 0xd0, 0xf8, 0x38, 0x9b, 0xe4, 0x3e, 0x69, 0x8d,
	0x26, 0x24, 0x4f, 0xf4, 0x70, 0xb1, 0xe2, 0xa4, 0x09, 0xb9, 0x43, 0x59, 0x98, 0x56, 0x0a, 0x3b,
	0x50, 0x1d, 0x24, 0x37, 0x37, 0x48, 0xa4, 0x8a, 0xd6, 0x78, 0x16, 0x6e, 0x2a, 0xf7, 0x04, 0x9e,
	0xe7, 0x72, 0x8b, 0x5f, 0x69, 0x0f, 0x17, 0xe9, 0x47, 0xac, 0xe5, 0x2e, 0x01, 0xf7, 0x93, 0x4c,
	0xa2, 0x7f, 0x92, 0xb6, 0x16, 0x73, 0x64, 0xcc, 0x2e, 0xbf, 0x50, 0xcc, 0xcb, 0xc6, 0x2a, 0x5e,
	0x36, 0x85, 0x99, 0xe4, 0x09, 0x5b, 0xcb, 0x78, 0xa7, 0x77, 0x8b, 0x30, 0xbe, 0x0a, 0xa7, 0x61,
	0x22, 0xa4, 0x39, 0xa7, 0xaa, 0xc2, 0x30, 0xbc, 0xc5, 0x20, 0xb3, 0xcf, 0x80, 0xdc, 0xbe, 0xde,
	0xa7, 0x65, 0xca, 0xd6, 0x22, 0xde, 0xc3, 0x4b, 0x75, 0xd5, 0x11, 0x8a, 0xcc, 0x6d, 0x8e, 0x51,
	0x82, 0x24, 0xfe, 0x77, 0x1d, 0xb8, 0x63, 0x99, 0x46, 0x6b, 0xd3, 0xb6, 0x96, 0x34, 0x81, 0x67,
	0xab, 0x92, 0x58, 0x07, 0xd8, 0xf2, 0x3d, 0x84, 0xc2, 0x9c, 0xd0, 0x1a, 0x66, 0x83, 0x3d, 0x1f,
	0xc8, 0xb7, 0x14, 0x3b, 0x90, 0xd2, 0xa9, 0x28, 0x9d, 0x1e, 0x24, 0xfd, 0xec, 0x00, 0x9a, 0x81,
	0x1f, 0xfc, 0xf0, 0x3a, 0xe4, 0xdd, 0xfa, 0xd7, 0x9d, 0xd8, 0x9b, 0x27, 0xc1, 0x2f, 0x9c, 0x8d,
	0xd3, 0x3f, 0xdd, 0x6b, 0xfd, 0x7b, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x18, 0x35, 0xa9, 0x0c,
	0x98, 0x07, 0x00, 0x00,
}
