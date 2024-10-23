// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.3
// source: Digimon.proto

package digimon

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

type Digimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Atributo string `protobuf:"bytes,2,opt,name=atributo,proto3" json:"atributo,omitempty"`
}

func (x *Digimon) Reset() {
	*x = Digimon{}
	mi := &file_Digimon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Digimon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Digimon) ProtoMessage() {}

func (x *Digimon) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Digimon.ProtoReflect.Descriptor instead.
func (*Digimon) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{0}
}

func (x *Digimon) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Digimon) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_Digimon_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CompletionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SacrificioPorcentaje string `protobuf:"bytes,1,opt,name=sacrificio_porcentaje,json=sacrificioPorcentaje,proto3" json:"sacrificio_porcentaje,omitempty"` // Porcentaje de sacrificio (PS)
	TiempoEspera         string `protobuf:"bytes,2,opt,name=tiempo_espera,json=tiempoEspera,proto3" json:"tiempo_espera,omitempty"`                         // tiempo de espera para enviar informacion (TE)
}

func (x *CompletionInfo) Reset() {
	*x = CompletionInfo{}
	mi := &file_Digimon_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionInfo) ProtoMessage() {}

func (x *CompletionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionInfo.ProtoReflect.Descriptor instead.
func (*CompletionInfo) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{2}
}

func (x *CompletionInfo) GetSacrificioPorcentaje() string {
	if x != nil {
		return x.SacrificioPorcentaje
	}
	return ""
}

func (x *CompletionInfo) GetTiempoEspera() string {
	if x != nil {
		return x.TiempoEspera
	}
	return ""
}

type Sacrificio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Atributo string `protobuf:"bytes,2,opt,name=atributo,proto3" json:"atributo,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"` // "Sacrificado" o "No sacrificado"
}

func (x *Sacrificio) Reset() {
	*x = Sacrificio{}
	mi := &file_Digimon_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sacrificio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sacrificio) ProtoMessage() {}

func (x *Sacrificio) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sacrificio.ProtoReflect.Descriptor instead.
func (*Sacrificio) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{3}
}

func (x *Sacrificio) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Sacrificio) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

func (x *Sacrificio) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type SacrificioResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodo              string        `protobuf:"bytes,1,opt,name=nodo,proto3" json:"nodo,omitempty"`         // Nombre del nodo regional
	Digimons          []*Sacrificio `protobuf:"bytes,2,rep,name=digimons,proto3" json:"digimons,omitempty"` // Lista de resultados de sacrificio
	MensajeEncriptado string        `protobuf:"bytes,3,opt,name=mensaje_encriptado,json=mensajeEncriptado,proto3" json:"mensaje_encriptado,omitempty"`
}

func (x *SacrificioResult) Reset() {
	*x = SacrificioResult{}
	mi := &file_Digimon_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SacrificioResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SacrificioResult) ProtoMessage() {}

func (x *SacrificioResult) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SacrificioResult.ProtoReflect.Descriptor instead.
func (*SacrificioResult) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{4}
}

func (x *SacrificioResult) GetNodo() string {
	if x != nil {
		return x.Nodo
	}
	return ""
}

func (x *SacrificioResult) GetDigimons() []*Sacrificio {
	if x != nil {
		return x.Digimons
	}
	return nil
}

func (x *SacrificioResult) GetMensajeEncriptado() string {
	if x != nil {
		return x.MensajeEncriptado
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	mi := &file_Digimon_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{5}
}

func (x *IdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AtributoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atributo string `protobuf:"bytes,1,opt,name=atributo,proto3" json:"atributo,omitempty"`
}

func (x *AtributoResponse) Reset() {
	*x = AtributoResponse{}
	mi := &file_Digimon_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AtributoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtributoResponse) ProtoMessage() {}

func (x *AtributoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtributoResponse.ProtoReflect.Descriptor instead.
func (*AtributoResponse) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{6}
}

func (x *AtributoResponse) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

type InfoNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Atributo string `protobuf:"bytes,2,opt,name=atributo,proto3" json:"atributo,omitempty"`
}

func (x *InfoNode) Reset() {
	*x = InfoNode{}
	mi := &file_Digimon_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoNode) ProtoMessage() {}

func (x *InfoNode) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoNode.ProtoReflect.Descriptor instead.
func (*InfoNode) Descriptor() ([]byte, []int) {
	return file_Digimon_proto_rawDescGZIP(), []int{7}
}

func (x *InfoNode) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InfoNode) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_Digimon_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[8]
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
	return file_Digimon_proto_rawDescGZIP(), []int{8}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_Digimon_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[9]
	if x != nil {
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
	return file_Digimon_proto_rawDescGZIP(), []int{9}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NumeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numero float64 `protobuf:"fixed64,1,opt,name=numero,proto3" json:"numero,omitempty"` // Número flotante que envía el cliente a NodoTai
}

func (x *NumeroResponse) Reset() {
	*x = NumeroResponse{}
	mi := &file_Digimon_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumeroResponse) ProtoMessage() {}

func (x *NumeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Digimon_proto_msgTypes[10]
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
	return file_Digimon_proto_rawDescGZIP(), []int{10}
}

func (x *NumeroResponse) GetNumero() float64 {
	if x != nil {
		return x.Numero
	}
	return 0
}

var File_Digimon_proto protoreflect.FileDescriptor

var file_Digimon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x07, 0x44, 0x69, 0x67, 0x69,
	0x6d, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x6a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x15, 0x73, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x69, 0x6f, 0x5f, 0x70, 0x6f, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x50,
	0x6f, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x6a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x65,
	0x6d, 0x70, 0x6f, 0x5f, 0x65, 0x73, 0x70, 0x65, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x45, 0x73, 0x70, 0x65, 0x72, 0x61, 0x22, 0x58,
	0x0a, 0x0a, 0x53, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x53, 0x61, 0x63,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x61,
	0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x52, 0x08, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f,
	0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x5f, 0x65, 0x6e,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x69, 0x70, 0x74, 0x61, 0x64,
	0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e,
	0x0a, 0x10, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x32,
	0x7e, 0x0a, 0x0e, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e,
	0x12, 0x10, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x6d,
	0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x0c, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x32,
	0x52, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x69, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x69, 0x6f, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63,
	0x6b, 0x22, 0x00, 0x32, 0x7f, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x11, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x4e,
	0x6f, 0x64, 0x65, 0x1a, 0x11, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x69, 0x67, 0x69,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x79, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x69, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x72, 0x12, 0x0e, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x0e, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1b, 0x5a, 0x19, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x64, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Digimon_proto_rawDescOnce sync.Once
	file_Digimon_proto_rawDescData = file_Digimon_proto_rawDesc
)

func file_Digimon_proto_rawDescGZIP() []byte {
	file_Digimon_proto_rawDescOnce.Do(func() {
		file_Digimon_proto_rawDescData = protoimpl.X.CompressGZIP(file_Digimon_proto_rawDescData)
	})
	return file_Digimon_proto_rawDescData
}

var file_Digimon_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_Digimon_proto_goTypes = []any{
	(*Digimon)(nil),          // 0: digimon.Digimon
	(*Ack)(nil),              // 1: digimon.Ack
	(*CompletionInfo)(nil),   // 2: digimon.CompletionInfo
	(*Sacrificio)(nil),       // 3: digimon.Sacrificio
	(*SacrificioResult)(nil), // 4: digimon.SacrificioResult
	(*IdRequest)(nil),        // 5: digimon.IdRequest
	(*AtributoResponse)(nil), // 6: digimon.AtributoResponse
	(*InfoNode)(nil),         // 7: digimon.InfoNode
	(*Empty)(nil),            // 8: digimon.Empty
	(*Response)(nil),         // 9: digimon.Response
	(*NumeroResponse)(nil),   // 10: digimon.NumeroResponse
}
var file_Digimon_proto_depIdxs = []int32{
	3,  // 0: digimon.SacrificioResult.digimons:type_name -> digimon.Sacrificio
	0,  // 1: digimon.DigimonService.SendDigimon:input_type -> digimon.Digimon
	2,  // 2: digimon.DigimonService.NotifyCompletion:input_type -> digimon.CompletionInfo
	4,  // 3: digimon.ResultService.SendSacrificioResult:input_type -> digimon.SacrificioResult
	7,  // 4: digimon.InfoService.MandarInfo:input_type -> digimon.InfoNode
	5,  // 5: digimon.InfoService.ObtenerInfo:input_type -> digimon.IdRequest
	8,  // 6: digimon.PrimarioService.Finalizar:input_type -> digimon.Empty
	8,  // 7: digimon.PrimarioService.SolicitarDatos:input_type -> digimon.Empty
	1,  // 8: digimon.DigimonService.SendDigimon:output_type -> digimon.Ack
	1,  // 9: digimon.DigimonService.NotifyCompletion:output_type -> digimon.Ack
	1,  // 10: digimon.ResultService.SendSacrificioResult:output_type -> digimon.Ack
	9,  // 11: digimon.InfoService.MandarInfo:output_type -> digimon.Response
	6,  // 12: digimon.InfoService.ObtenerInfo:output_type -> digimon.AtributoResponse
	8,  // 13: digimon.PrimarioService.Finalizar:output_type -> digimon.Empty
	10, // 14: digimon.PrimarioService.SolicitarDatos:output_type -> digimon.NumeroResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_Digimon_proto_init() }
func file_Digimon_proto_init() {
	if File_Digimon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Digimon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_Digimon_proto_goTypes,
		DependencyIndexes: file_Digimon_proto_depIdxs,
		MessageInfos:      file_Digimon_proto_msgTypes,
	}.Build()
	File_Digimon_proto = out.File
	file_Digimon_proto_rawDesc = nil
	file_Digimon_proto_goTypes = nil
	file_Digimon_proto_depIdxs = nil
}