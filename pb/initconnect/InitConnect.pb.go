// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: InitConnect.proto

package initconnect

import (
	_ "github.com/Lawrence93Ho/go-futu-api/pb/common"
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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientVer  *int32  `protobuf:"varint,1,req,name=clientVer" json:"clientVer,omitempty"`   //客户端版本号，clientVer = "."以前的数 * 100 + "."以后的，举例：1.1版本的clientVer为1 * 100 + 1 = 101，2.21版本为2 * 100 + 21 = 221
	ClientID   *string `protobuf:"bytes,2,req,name=clientID" json:"clientID,omitempty"`      //客户端唯一标识，无生具体生成规则，客户端自己保证唯一性即可
	RecvNotify *bool   `protobuf:"varint,3,opt,name=recvNotify" json:"recvNotify,omitempty"` //此连接是否接收市场状态、交易需要重新解锁等等事件通知，true代表接收，FutuOpenD就会向此连接推送这些通知，反之false代表不接收不推送
	//如果通信要加密，首先得在FutuOpenD和客户端都配置RSA密钥，不配置始终不加密
	//如果配置了RSA密钥且指定的加密算法不为PacketEncAlgo_None则加密(即便这里不设置，配置了RSA密钥，也会采用默认加密方式)，默认采用FTAES_ECB算法
	PacketEncAlgo       *int32  `protobuf:"varint,4,opt,name=packetEncAlgo" json:"packetEncAlgo,omitempty"`            //指定包加密算法，参见Common.PacketEncAlgo的枚举定义
	PushProtoFmt        *int32  `protobuf:"varint,5,opt,name=pushProtoFmt" json:"pushProtoFmt,omitempty"`              //指定这条连接上的推送协议格式，若不指定则使用push_proto_type配置项
	ProgrammingLanguage *string `protobuf:"bytes,6,opt,name=programmingLanguage" json:"programmingLanguage,omitempty"` //接口编程语言，用于统计语言偏好
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InitConnect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_InitConnect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_InitConnect_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetClientVer() int32 {
	if x != nil && x.ClientVer != nil {
		return *x.ClientVer
	}
	return 0
}

func (x *C2S) GetClientID() string {
	if x != nil && x.ClientID != nil {
		return *x.ClientID
	}
	return ""
}

func (x *C2S) GetRecvNotify() bool {
	if x != nil && x.RecvNotify != nil {
		return *x.RecvNotify
	}
	return false
}

func (x *C2S) GetPacketEncAlgo() int32 {
	if x != nil && x.PacketEncAlgo != nil {
		return *x.PacketEncAlgo
	}
	return 0
}

func (x *C2S) GetPushProtoFmt() int32 {
	if x != nil && x.PushProtoFmt != nil {
		return *x.PushProtoFmt
	}
	return 0
}

func (x *C2S) GetProgrammingLanguage() string {
	if x != nil && x.ProgrammingLanguage != nil {
		return *x.ProgrammingLanguage
	}
	return ""
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerVer         *int32  `protobuf:"varint,1,req,name=serverVer" json:"serverVer,omitempty"`                 //FutuOpenD的版本号
	LoginUserID       *uint64 `protobuf:"varint,2,req,name=loginUserID" json:"loginUserID,omitempty"`             //FutuOpenD登陆的牛牛用户ID
	ConnID            *uint64 `protobuf:"varint,3,req,name=connID" json:"connID,omitempty"`                       //此连接的连接ID，连接的唯一标识
	ConnAESKey        *string `protobuf:"bytes,4,req,name=connAESKey" json:"connAESKey,omitempty"`                //此连接后续AES加密通信的Key，固定为16字节长字符串
	KeepAliveInterval *int32  `protobuf:"varint,5,req,name=keepAliveInterval" json:"keepAliveInterval,omitempty"` //心跳保活间隔
	AesCBCiv          *string `protobuf:"bytes,6,opt,name=aesCBCiv" json:"aesCBCiv,omitempty"`                    //AES加密通信CBC加密模式的iv，固定为16字节长字符串
	UserAttribution   *int32  `protobuf:"varint,7,opt,name=userAttribution" json:"userAttribution,omitempty"`     //用户类型，牛牛用户或MooMoo用户
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InitConnect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_InitConnect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_InitConnect_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetServerVer() int32 {
	if x != nil && x.ServerVer != nil {
		return *x.ServerVer
	}
	return 0
}

func (x *S2C) GetLoginUserID() uint64 {
	if x != nil && x.LoginUserID != nil {
		return *x.LoginUserID
	}
	return 0
}

func (x *S2C) GetConnID() uint64 {
	if x != nil && x.ConnID != nil {
		return *x.ConnID
	}
	return 0
}

func (x *S2C) GetConnAESKey() string {
	if x != nil && x.ConnAESKey != nil {
		return *x.ConnAESKey
	}
	return ""
}

func (x *S2C) GetKeepAliveInterval() int32 {
	if x != nil && x.KeepAliveInterval != nil {
		return *x.KeepAliveInterval
	}
	return 0
}

func (x *S2C) GetAesCBCiv() string {
	if x != nil && x.AesCBCiv != nil {
		return *x.AesCBCiv
	}
	return ""
}

func (x *S2C) GetUserAttribution() int32 {
	if x != nil && x.UserAttribution != nil {
		return *x.UserAttribution
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InitConnect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_InitConnect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_InitConnect_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //返回结果，参见Common.RetType的枚举定义
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`             //返回结果描述
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`          //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InitConnect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_InitConnect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_InitConnect_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_InitConnect_proto protoreflect.FileDescriptor

var file_InitConnect_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb,
	0x01, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c, 0x67,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45,
	0x6e, 0x63, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x46, 0x6d, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x75,
	0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x6d, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xf1, 0x01, 0x0a,
	0x03, 0x53, 0x32, 0x43, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x6e, 0x41, 0x45, 0x53, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x41, 0x45, 0x53, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x11,
	0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x02, 0x28, 0x05, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x65,
	0x73, 0x43, 0x42, 0x43, 0x69, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x65,
	0x73, 0x43, 0x42, 0x43, 0x69, 0x76, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x63,
	0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22,
	0x80, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d,
	0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73,
	0x32, 0x63, 0x42, 0x46, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f,
	0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x69,
	0x6e, 0x69, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
}

var (
	file_InitConnect_proto_rawDescOnce sync.Once
	file_InitConnect_proto_rawDescData = file_InitConnect_proto_rawDesc
)

func file_InitConnect_proto_rawDescGZIP() []byte {
	file_InitConnect_proto_rawDescOnce.Do(func() {
		file_InitConnect_proto_rawDescData = protoimpl.X.CompressGZIP(file_InitConnect_proto_rawDescData)
	})
	return file_InitConnect_proto_rawDescData
}

var file_InitConnect_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_InitConnect_proto_goTypes = []interface{}{
	(*C2S)(nil),      // 0: InitConnect.C2S
	(*S2C)(nil),      // 1: InitConnect.S2C
	(*Request)(nil),  // 2: InitConnect.Request
	(*Response)(nil), // 3: InitConnect.Response
}
var file_InitConnect_proto_depIdxs = []int32{
	0, // 0: InitConnect.Request.c2s:type_name -> InitConnect.C2S
	1, // 1: InitConnect.Response.s2c:type_name -> InitConnect.S2C
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_InitConnect_proto_init() }
func file_InitConnect_proto_init() {
	if File_InitConnect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_InitConnect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_InitConnect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_InitConnect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_InitConnect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_InitConnect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InitConnect_proto_goTypes,
		DependencyIndexes: file_InitConnect_proto_depIdxs,
		MessageInfos:      file_InitConnect_proto_msgTypes,
	}.Build()
	File_InitConnect_proto = out.File
	file_InitConnect_proto_rawDesc = nil
	file_InitConnect_proto_goTypes = nil
	file_InitConnect_proto_depIdxs = nil
}
