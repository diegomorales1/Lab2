// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.3
// source: NodoTai.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_NodoTai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{0}
}

type Ataque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dano int32 `protobuf:"varint,1,opt,name=dano,proto3" json:"dano,omitempty"`
}

func (x *Ataque) Reset() {
	*x = Ataque{}
	mi := &file_NodoTai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ataque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ataque) ProtoMessage() {}

func (x *Ataque) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ataque.ProtoReflect.Descriptor instead.
func (*Ataque) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{1}
}

func (x *Ataque) GetDano() int32 {
	if x != nil {
		return x.Dano
	}
	return 0
}

type Senal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor int32 `protobuf:"varint,1,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *Senal) Reset() {
	*x = Senal{}
	mi := &file_NodoTai_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Senal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Senal) ProtoMessage() {}

func (x *Senal) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Senal.ProtoReflect.Descriptor instead.
func (*Senal) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{2}
}

func (x *Senal) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

type NumeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numero float64 `protobuf:"fixed64,1,opt,name=numero,proto3" json:"numero,omitempty"` // Número flotante que envia el Primario a NodoTai
}

func (x *NumeroResponse) Reset() {
	*x = NumeroResponse{}
	mi := &file_NodoTai_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumeroResponse) ProtoMessage() {}

func (x *NumeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumeroResponse.ProtoReflect.Descriptor instead.
func (*NumeroResponse) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{3}
}

func (x *NumeroResponse) GetNumero() float64 {
	if x != nil {
		return x.Numero
	}
	return 0
}

type VidaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vida int32 `protobuf:"varint,1,opt,name=vida,proto3" json:"vida,omitempty"`
}

func (x *VidaRequest) Reset() {
	*x = VidaRequest{}
	mi := &file_NodoTai_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VidaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VidaRequest) ProtoMessage() {}

func (x *VidaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VidaRequest.ProtoReflect.Descriptor instead.
func (*VidaRequest) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{4}
}

func (x *VidaRequest) GetVida() int32 {
	if x != nil {
		return x.Vida
	}
	return 0
}

type AtaqueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje      string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	VidaRestante int32  `protobuf:"varint,2,opt,name=vida_restante,json=vidaRestante,proto3" json:"vida_restante,omitempty"`
}

func (x *AtaqueResponse) Reset() {
	*x = AtaqueResponse{}
	mi := &file_NodoTai_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AtaqueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtaqueResponse) ProtoMessage() {}

func (x *AtaqueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_NodoTai_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtaqueResponse.ProtoReflect.Descriptor instead.
func (*AtaqueResponse) Descriptor() ([]byte, []int) {
	return file_NodoTai_proto_rawDescGZIP(), []int{5}
}

func (x *AtaqueResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *AtaqueResponse) GetVidaRestante() int32 {
	if x != nil {
		return x.VidaRestante
	}
	return 0
}

var File_NodoTai_proto protoreflect.FileDescriptor

var file_NodoTai_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1c, 0x0a, 0x06, 0x41, 0x74, 0x61, 0x71, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x6e, 0x6f, 0x22,
	0x1d, 0x0a, 0x05, 0x53, 0x65, 0x6e, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x28,
	0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x64, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x69, 0x64, 0x61, 0x22, 0x4f, 0x0a, 0x0e, 0x41,
	0x74, 0x61, 0x71, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69, 0x64, 0x61, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x76, 0x69, 0x64, 0x61, 0x52, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x32, 0x79, 0x0a, 0x0f,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x12, 0x0e, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0e,
	0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x0e,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x44, 0x0a, 0x11, 0x44, 0x69, 0x61, 0x62, 0x6f,
	0x72, 0x6f, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0c,
	0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x41, 0x74, 0x61, 0x71, 0x75, 0x65, 0x12, 0x0f, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x61, 0x71, 0x75, 0x65, 0x1a, 0x0e, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x61, 0x6c, 0x42, 0x13, 0x5a,
	0x11, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NodoTai_proto_rawDescOnce sync.Once
	file_NodoTai_proto_rawDescData = file_NodoTai_proto_rawDesc
)

func file_NodoTai_proto_rawDescGZIP() []byte {
	file_NodoTai_proto_rawDescOnce.Do(func() {
		file_NodoTai_proto_rawDescData = protoimpl.X.CompressGZIP(file_NodoTai_proto_rawDescData)
	})
	return file_NodoTai_proto_rawDescData
}

var file_NodoTai_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_NodoTai_proto_goTypes = []any{
	(*Empty)(nil),          // 0: digimon.Empty
	(*Ataque)(nil),         // 1: digimon.Ataque
	(*Senal)(nil),          // 2: digimon.Senal
	(*NumeroResponse)(nil), // 3: digimon.NumeroResponse
	(*VidaRequest)(nil),    // 4: digimon.VidaRequest
	(*AtaqueResponse)(nil), // 5: digimon.AtaqueResponse
}
var file_NodoTai_proto_depIdxs = []int32{
	0, // 0: digimon.PrimarioService.Finalizar:input_type -> digimon.Empty
	0, // 1: digimon.PrimarioService.SolicitarDatos:input_type -> digimon.Empty
	1, // 2: digimon.DiaboromonService.EnviarAtaque:input_type -> digimon.Ataque
	2, // 3: digimon.PrimarioService.Finalizar:output_type -> digimon.Senal
	3, // 4: digimon.PrimarioService.SolicitarDatos:output_type -> digimon.NumeroResponse
	2, // 5: digimon.DiaboromonService.EnviarAtaque:output_type -> digimon.Senal
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NodoTai_proto_init() }
func file_NodoTai_proto_init() {
	if File_NodoTai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NodoTai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_NodoTai_proto_goTypes,
		DependencyIndexes: file_NodoTai_proto_depIdxs,
		MessageInfos:      file_NodoTai_proto_msgTypes,
	}.Build()
	File_NodoTai_proto = out.File
	file_NodoTai_proto_rawDesc = nil
	file_NodoTai_proto_goTypes = nil
	file_NodoTai_proto_depIdxs = nil
}
