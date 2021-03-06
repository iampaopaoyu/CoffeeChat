// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.Login.proto

package cim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 认证请求
type CIMAuthTokenReq struct {
	// cmd id:		0x0101
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// CoffeeChat不存储用户信息，消息推送时需要显示昵称
	// 基于流量考虑，昵称不放在每条消息中携带
	// 但是如果期间用户更新昵称后，消息推送显示昵称会有延迟，CoffeeChat认为是能接受的
	NickName   string        `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	UserToken  string        `protobuf:"bytes,3,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	ClientType CIMClientType `protobuf:"varint,4,opt,name=client_type,json=clientType,proto3,enum=CIM.Def.CIMClientType" json:"client_type,omitempty"`
	//optional
	ClientVersion        string   `protobuf:"bytes,5,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMAuthTokenReq) Reset()         { *m = CIMAuthTokenReq{} }
func (m *CIMAuthTokenReq) String() string { return proto.CompactTextString(m) }
func (*CIMAuthTokenReq) ProtoMessage()    {}
func (*CIMAuthTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6994aede821e63e, []int{0}
}

func (m *CIMAuthTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMAuthTokenReq.Unmarshal(m, b)
}
func (m *CIMAuthTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMAuthTokenReq.Marshal(b, m, deterministic)
}
func (m *CIMAuthTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMAuthTokenReq.Merge(m, src)
}
func (m *CIMAuthTokenReq) XXX_Size() int {
	return xxx_messageInfo_CIMAuthTokenReq.Size(m)
}
func (m *CIMAuthTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMAuthTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMAuthTokenReq proto.InternalMessageInfo

func (m *CIMAuthTokenReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMAuthTokenReq) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CIMAuthTokenReq) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *CIMAuthTokenReq) GetClientType() CIMClientType {
	if m != nil {
		return m.ClientType
	}
	return CIMClientType_kCIM_CLIENT_TYPE_DEFAULT
}

func (m *CIMAuthTokenReq) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type CIMAuthTokenRsp struct {
	// cmd id:		0x0102
	ServerTime uint32       `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	ResultCode CIMErrorCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=CIM.Def.CIMErrorCode" json:"result_code,omitempty"`
	//optional
	ResultString string `protobuf:"bytes,3,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"`
	//optional
	UserInfo             *CIMUserInfo `protobuf:"bytes,4,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CIMAuthTokenRsp) Reset()         { *m = CIMAuthTokenRsp{} }
func (m *CIMAuthTokenRsp) String() string { return proto.CompactTextString(m) }
func (*CIMAuthTokenRsp) ProtoMessage()    {}
func (*CIMAuthTokenRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6994aede821e63e, []int{1}
}

func (m *CIMAuthTokenRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMAuthTokenRsp.Unmarshal(m, b)
}
func (m *CIMAuthTokenRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMAuthTokenRsp.Marshal(b, m, deterministic)
}
func (m *CIMAuthTokenRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMAuthTokenRsp.Merge(m, src)
}
func (m *CIMAuthTokenRsp) XXX_Size() int {
	return xxx_messageInfo_CIMAuthTokenRsp.Size(m)
}
func (m *CIMAuthTokenRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMAuthTokenRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMAuthTokenRsp proto.InternalMessageInfo

func (m *CIMAuthTokenRsp) GetServerTime() uint32 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

func (m *CIMAuthTokenRsp) GetResultCode() CIMErrorCode {
	if m != nil {
		return m.ResultCode
	}
	return CIMErrorCode_kCIM_ERR_UNKNOWN
}

func (m *CIMAuthTokenRsp) GetResultString() string {
	if m != nil {
		return m.ResultString
	}
	return ""
}

func (m *CIMAuthTokenRsp) GetUserInfo() *CIMUserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

// 登出
type CIMLogoutReq struct {
	// cmd id:		0x0103
	UserId               uint64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientType           CIMClientType `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=CIM.Def.CIMClientType" json:"client_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CIMLogoutReq) Reset()         { *m = CIMLogoutReq{} }
func (m *CIMLogoutReq) String() string { return proto.CompactTextString(m) }
func (*CIMLogoutReq) ProtoMessage()    {}
func (*CIMLogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6994aede821e63e, []int{2}
}

func (m *CIMLogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMLogoutReq.Unmarshal(m, b)
}
func (m *CIMLogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMLogoutReq.Marshal(b, m, deterministic)
}
func (m *CIMLogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMLogoutReq.Merge(m, src)
}
func (m *CIMLogoutReq) XXX_Size() int {
	return xxx_messageInfo_CIMLogoutReq.Size(m)
}
func (m *CIMLogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMLogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMLogoutReq proto.InternalMessageInfo

func (m *CIMLogoutReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMLogoutReq) GetClientType() CIMClientType {
	if m != nil {
		return m.ClientType
	}
	return CIMClientType_kCIM_CLIENT_TYPE_DEFAULT
}

type CIMLogoutRsp struct {
	// cmd id:		0x0104
	ResultCode           uint32   `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMLogoutRsp) Reset()         { *m = CIMLogoutRsp{} }
func (m *CIMLogoutRsp) String() string { return proto.CompactTextString(m) }
func (*CIMLogoutRsp) ProtoMessage()    {}
func (*CIMLogoutRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6994aede821e63e, []int{3}
}

func (m *CIMLogoutRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMLogoutRsp.Unmarshal(m, b)
}
func (m *CIMLogoutRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMLogoutRsp.Marshal(b, m, deterministic)
}
func (m *CIMLogoutRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMLogoutRsp.Merge(m, src)
}
func (m *CIMLogoutRsp) XXX_Size() int {
	return xxx_messageInfo_CIMLogoutRsp.Size(m)
}
func (m *CIMLogoutRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMLogoutRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMLogoutRsp proto.InternalMessageInfo

func (m *CIMLogoutRsp) GetResultCode() uint32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

// 心跳
type CIMHeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMHeartBeat) Reset()         { *m = CIMHeartBeat{} }
func (m *CIMHeartBeat) String() string { return proto.CompactTextString(m) }
func (*CIMHeartBeat) ProtoMessage()    {}
func (*CIMHeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6994aede821e63e, []int{4}
}

func (m *CIMHeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMHeartBeat.Unmarshal(m, b)
}
func (m *CIMHeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMHeartBeat.Marshal(b, m, deterministic)
}
func (m *CIMHeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMHeartBeat.Merge(m, src)
}
func (m *CIMHeartBeat) XXX_Size() int {
	return xxx_messageInfo_CIMHeartBeat.Size(m)
}
func (m *CIMHeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMHeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_CIMHeartBeat proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CIMAuthTokenReq)(nil), "CIM.Login.CIMAuthTokenReq")
	proto.RegisterType((*CIMAuthTokenRsp)(nil), "CIM.Login.CIMAuthTokenRsp")
	proto.RegisterType((*CIMLogoutReq)(nil), "CIM.Login.CIMLogoutReq")
	proto.RegisterType((*CIMLogoutRsp)(nil), "CIM.Login.CIMLogoutRsp")
	proto.RegisterType((*CIMHeartBeat)(nil), "CIM.Login.CIMHeartBeat")
}

func init() { proto.RegisterFile("CIM.Login.proto", fileDescriptor_f6994aede821e63e) }

var fileDescriptor_f6994aede821e63e = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8f, 0x93, 0x40,
	0x14, 0xc6, 0xc3, 0x76, 0x5d, 0x97, 0xc7, 0x82, 0xc9, 0x44, 0x5d, 0xa2, 0x31, 0x12, 0x8c, 0x49,
	0x4f, 0x18, 0x6b, 0xa2, 0x67, 0x17, 0x4d, 0x4a, 0xd2, 0x7a, 0xc0, 0xea, 0xc1, 0x0b, 0xd2, 0xe1,
	0xd1, 0x4e, 0x5a, 0x66, 0x70, 0x18, 0x9a, 0xf4, 0xdf, 0xf3, 0xe0, 0xdf, 0x65, 0x66, 0x06, 0x2b,
	0xe9, 0xc1, 0x78, 0x7b, 0xfc, 0xde, 0x7b, 0x1f, 0xef, 0xfb, 0x00, 0x1e, 0xa4, 0xd9, 0x32, 0x59,
	0x88, 0x0d, 0xe3, 0x49, 0x2b, 0x85, 0x12, 0xc4, 0x3d, 0x81, 0x27, 0xbe, 0x2e, 0x3f, 0x60, 0x6d,
	0x3b, 0xf1, 0x2f, 0xc7, 0x4c, 0xbf, 0xef, 0xd5, 0x76, 0x25, 0x76, 0xc8, 0x73, 0xfc, 0x41, 0x6e,
	0xe1, 0x7e, 0xdf, 0xa1, 0x2c, 0x58, 0x15, 0x3a, 0x91, 0x33, 0xbd, 0xcc, 0xaf, 0xf4, 0x63, 0x56,
	0x91, 0xa7, 0xe0, 0x72, 0x46, 0x77, 0x05, 0x2f, 0x1b, 0x0c, 0x2f, 0x22, 0x67, 0xea, 0xe6, 0xd7,
	0x1a, 0x7c, 0x2a, 0x1b, 0x24, 0xcf, 0x00, 0xcc, 0x96, 0xd2, 0x32, 0xe1, 0xc4, 0x74, 0x5d, 0x4d,
	0x8c, 0x2e, 0x79, 0x07, 0x1e, 0xdd, 0x33, 0xe4, 0xaa, 0x50, 0xc7, 0x16, 0xc3, 0xcb, 0xc8, 0x99,
	0x06, 0xb3, 0xc7, 0xc9, 0x9f, 0x6b, 0xd2, 0x6c, 0x99, 0x9a, 0xf6, 0xea, 0xd8, 0x62, 0x0e, 0xf4,
	0x54, 0x93, 0x97, 0x10, 0x0c, 0x8b, 0x07, 0x94, 0x1d, 0x13, 0x3c, 0xbc, 0x67, 0xb4, 0x7d, 0x4b,
	0xbf, 0x5a, 0x18, 0xff, 0x3c, 0x37, 0xd2, 0xb5, 0xe4, 0x39, 0x78, 0x1d, 0xca, 0x83, 0x3e, 0x8a,
	0x35, 0x68, 0xcc, 0xf8, 0x39, 0x58, 0xb4, 0x62, 0x0d, 0x92, 0xb7, 0xe0, 0x49, 0xec, 0xfa, 0xbd,
	0x2a, 0xa8, 0xa8, 0xac, 0xa5, 0x60, 0xf6, 0x68, 0x7c, 0xd4, 0x47, 0x29, 0x85, 0x4c, 0x45, 0x85,
	0x39, 0xd8, 0x49, 0x5d, 0x93, 0x17, 0xe0, 0x0f, 0x7b, 0x9d, 0x92, 0x8c, 0x6f, 0x06, 0xbb, 0x37,
	0x16, 0x7e, 0x36, 0x8c, 0xbc, 0x06, 0xd7, 0xc6, 0xc8, 0x6b, 0x61, 0xfc, 0x7a, 0xb3, 0x87, 0x63,
	0xe9, 0x2f, 0x3a, 0x54, 0x5e, 0x8b, 0xfc, 0xba, 0x1f, 0xaa, 0xf8, 0x3b, 0xdc, 0xa4, 0xd9, 0x72,
	0x21, 0x36, 0xa2, 0x57, 0xff, 0xfc, 0x12, 0x67, 0x69, 0x5e, 0xfc, 0x6f, 0x9a, 0xf1, 0xab, 0xf1,
	0x1b, 0x6c, 0x44, 0xe3, 0x04, 0x86, 0x88, 0xfe, 0x5a, 0x8d, 0x03, 0xb3, 0x30, 0xc7, 0x52, 0xaa,
	0x3b, 0x2c, 0xd5, 0x5d, 0x04, 0xb7, 0x54, 0x34, 0x09, 0x15, 0x75, 0x8d, 0x48, 0xb7, 0xa5, 0xb2,
	0x3f, 0xd2, 0xba, 0xaf, 0xe7, 0x93, 0x6f, 0x13, 0xca, 0x9a, 0xf5, 0x95, 0x01, 0x6f, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x22, 0xf5, 0xed, 0x86, 0x02, 0x00, 0x00,
}
