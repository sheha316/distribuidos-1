// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comms.proto

package comms

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

type Request_CrearOrdenPyme struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto             string   `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor                int32    `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda               string   `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino              string   `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioritario          int32    `protobuf:"varint,6,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_CrearOrdenPyme) Reset()         { *m = Request_CrearOrdenPyme{} }
func (m *Request_CrearOrdenPyme) String() string { return proto.CompactTextString(m) }
func (*Request_CrearOrdenPyme) ProtoMessage()    {}
func (*Request_CrearOrdenPyme) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{0}
}

func (m *Request_CrearOrdenPyme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_CrearOrdenPyme.Unmarshal(m, b)
}
func (m *Request_CrearOrdenPyme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_CrearOrdenPyme.Marshal(b, m, deterministic)
}
func (m *Request_CrearOrdenPyme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_CrearOrdenPyme.Merge(m, src)
}
func (m *Request_CrearOrdenPyme) XXX_Size() int {
	return xxx_messageInfo_Request_CrearOrdenPyme.Size(m)
}
func (m *Request_CrearOrdenPyme) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_CrearOrdenPyme.DiscardUnknown(m)
}

var xxx_messageInfo_Request_CrearOrdenPyme proto.InternalMessageInfo

func (m *Request_CrearOrdenPyme) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_CrearOrdenPyme) GetProducto() string {
	if m != nil {
		return m.Producto
	}
	return ""
}

func (m *Request_CrearOrdenPyme) GetValor() int32 {
	if m != nil {
		return m.Valor
	}
	return 0
}

func (m *Request_CrearOrdenPyme) GetTienda() string {
	if m != nil {
		return m.Tienda
	}
	return ""
}

func (m *Request_CrearOrdenPyme) GetDestino() string {
	if m != nil {
		return m.Destino
	}
	return ""
}

func (m *Request_CrearOrdenPyme) GetPrioritario() int32 {
	if m != nil {
		return m.Prioritario
	}
	return 0
}

type Request_CrearOrdenRetail struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto             string   `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor                int32    `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda               string   `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino              string   `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_CrearOrdenRetail) Reset()         { *m = Request_CrearOrdenRetail{} }
func (m *Request_CrearOrdenRetail) String() string { return proto.CompactTextString(m) }
func (*Request_CrearOrdenRetail) ProtoMessage()    {}
func (*Request_CrearOrdenRetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{1}
}

func (m *Request_CrearOrdenRetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_CrearOrdenRetail.Unmarshal(m, b)
}
func (m *Request_CrearOrdenRetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_CrearOrdenRetail.Marshal(b, m, deterministic)
}
func (m *Request_CrearOrdenRetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_CrearOrdenRetail.Merge(m, src)
}
func (m *Request_CrearOrdenRetail) XXX_Size() int {
	return xxx_messageInfo_Request_CrearOrdenRetail.Size(m)
}
func (m *Request_CrearOrdenRetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_CrearOrdenRetail.DiscardUnknown(m)
}

var xxx_messageInfo_Request_CrearOrdenRetail proto.InternalMessageInfo

func (m *Request_CrearOrdenRetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_CrearOrdenRetail) GetProducto() string {
	if m != nil {
		return m.Producto
	}
	return ""
}

func (m *Request_CrearOrdenRetail) GetValor() int32 {
	if m != nil {
		return m.Valor
	}
	return 0
}

func (m *Request_CrearOrdenRetail) GetTienda() string {
	if m != nil {
		return m.Tienda
	}
	return ""
}

func (m *Request_CrearOrdenRetail) GetDestino() string {
	if m != nil {
		return m.Destino
	}
	return ""
}

type Response_CrearOrden struct {
	Seguimiento          int32    `protobuf:"varint,1,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_CrearOrden) Reset()         { *m = Response_CrearOrden{} }
func (m *Response_CrearOrden) String() string { return proto.CompactTextString(m) }
func (*Response_CrearOrden) ProtoMessage()    {}
func (*Response_CrearOrden) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{2}
}

func (m *Response_CrearOrden) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_CrearOrden.Unmarshal(m, b)
}
func (m *Response_CrearOrden) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_CrearOrden.Marshal(b, m, deterministic)
}
func (m *Response_CrearOrden) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_CrearOrden.Merge(m, src)
}
func (m *Response_CrearOrden) XXX_Size() int {
	return xxx_messageInfo_Response_CrearOrden.Size(m)
}
func (m *Response_CrearOrden) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_CrearOrden.DiscardUnknown(m)
}

var xxx_messageInfo_Response_CrearOrden proto.InternalMessageInfo

func (m *Response_CrearOrden) GetSeguimiento() int32 {
	if m != nil {
		return m.Seguimiento
	}
	return 0
}

type Request_Seguimiento struct {
	Seguimiento          int32    `protobuf:"varint,1,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Seguimiento) Reset()         { *m = Request_Seguimiento{} }
func (m *Request_Seguimiento) String() string { return proto.CompactTextString(m) }
func (*Request_Seguimiento) ProtoMessage()    {}
func (*Request_Seguimiento) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{3}
}

func (m *Request_Seguimiento) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Seguimiento.Unmarshal(m, b)
}
func (m *Request_Seguimiento) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Seguimiento.Marshal(b, m, deterministic)
}
func (m *Request_Seguimiento) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Seguimiento.Merge(m, src)
}
func (m *Request_Seguimiento) XXX_Size() int {
	return xxx_messageInfo_Request_Seguimiento.Size(m)
}
func (m *Request_Seguimiento) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Seguimiento.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Seguimiento proto.InternalMessageInfo

func (m *Request_Seguimiento) GetSeguimiento() int32 {
	if m != nil {
		return m.Seguimiento
	}
	return 0
}

type Response_Seguimiento struct {
	Estado               string   `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Seguimiento) Reset()         { *m = Response_Seguimiento{} }
func (m *Response_Seguimiento) String() string { return proto.CompactTextString(m) }
func (*Response_Seguimiento) ProtoMessage()    {}
func (*Response_Seguimiento) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{4}
}

func (m *Response_Seguimiento) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Seguimiento.Unmarshal(m, b)
}
func (m *Response_Seguimiento) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Seguimiento.Marshal(b, m, deterministic)
}
func (m *Response_Seguimiento) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Seguimiento.Merge(m, src)
}
func (m *Response_Seguimiento) XXX_Size() int {
	return xxx_messageInfo_Response_Seguimiento.Size(m)
}
func (m *Response_Seguimiento) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Seguimiento.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Seguimiento proto.InternalMessageInfo

func (m *Response_Seguimiento) GetEstado() string {
	if m != nil {
		return m.Estado
	}
	return ""
}

type Request_SolicitarPaquete struct {
	Tipo                 string   `protobuf:"bytes,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_SolicitarPaquete) Reset()         { *m = Request_SolicitarPaquete{} }
func (m *Request_SolicitarPaquete) String() string { return proto.CompactTextString(m) }
func (*Request_SolicitarPaquete) ProtoMessage()    {}
func (*Request_SolicitarPaquete) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{5}
}

func (m *Request_SolicitarPaquete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_SolicitarPaquete.Unmarshal(m, b)
}
func (m *Request_SolicitarPaquete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_SolicitarPaquete.Marshal(b, m, deterministic)
}
func (m *Request_SolicitarPaquete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_SolicitarPaquete.Merge(m, src)
}
func (m *Request_SolicitarPaquete) XXX_Size() int {
	return xxx_messageInfo_Request_SolicitarPaquete.Size(m)
}
func (m *Request_SolicitarPaquete) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_SolicitarPaquete.DiscardUnknown(m)
}

var xxx_messageInfo_Request_SolicitarPaquete proto.InternalMessageInfo

func (m *Request_SolicitarPaquete) GetTipo() string {
	if m != nil {
		return m.Tipo
	}
	return ""
}

func (m *Request_SolicitarPaquete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response_SolicitarPaquete struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Seguimiento          int32    `protobuf:"varint,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo                 string   `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor                int32    `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda               string   `protobuf:"bytes,5,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino              string   `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_SolicitarPaquete) Reset()         { *m = Response_SolicitarPaquete{} }
func (m *Response_SolicitarPaquete) String() string { return proto.CompactTextString(m) }
func (*Response_SolicitarPaquete) ProtoMessage()    {}
func (*Response_SolicitarPaquete) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{6}
}

func (m *Response_SolicitarPaquete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_SolicitarPaquete.Unmarshal(m, b)
}
func (m *Response_SolicitarPaquete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_SolicitarPaquete.Marshal(b, m, deterministic)
}
func (m *Response_SolicitarPaquete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_SolicitarPaquete.Merge(m, src)
}
func (m *Response_SolicitarPaquete) XXX_Size() int {
	return xxx_messageInfo_Response_SolicitarPaquete.Size(m)
}
func (m *Response_SolicitarPaquete) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_SolicitarPaquete.DiscardUnknown(m)
}

var xxx_messageInfo_Response_SolicitarPaquete proto.InternalMessageInfo

func (m *Response_SolicitarPaquete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response_SolicitarPaquete) GetSeguimiento() int32 {
	if m != nil {
		return m.Seguimiento
	}
	return 0
}

func (m *Response_SolicitarPaquete) GetTipo() string {
	if m != nil {
		return m.Tipo
	}
	return ""
}

func (m *Response_SolicitarPaquete) GetValor() int32 {
	if m != nil {
		return m.Valor
	}
	return 0
}

func (m *Response_SolicitarPaquete) GetTienda() string {
	if m != nil {
		return m.Tienda
	}
	return ""
}

func (m *Response_SolicitarPaquete) GetDestino() string {
	if m != nil {
		return m.Destino
	}
	return ""
}

type Request_Estado struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Intentos             int32    `protobuf:"varint,2,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Fecha                string   `protobuf:"bytes,3,opt,name=fecha,proto3" json:"fecha,omitempty"`
	Estado               string   `protobuf:"bytes,4,opt,name=estado,proto3" json:"estado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Estado) Reset()         { *m = Request_Estado{} }
func (m *Request_Estado) String() string { return proto.CompactTextString(m) }
func (*Request_Estado) ProtoMessage()    {}
func (*Request_Estado) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{7}
}

func (m *Request_Estado) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Estado.Unmarshal(m, b)
}
func (m *Request_Estado) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Estado.Marshal(b, m, deterministic)
}
func (m *Request_Estado) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Estado.Merge(m, src)
}
func (m *Request_Estado) XXX_Size() int {
	return xxx_messageInfo_Request_Estado.Size(m)
}
func (m *Request_Estado) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Estado.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Estado proto.InternalMessageInfo

func (m *Request_Estado) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_Estado) GetIntentos() int32 {
	if m != nil {
		return m.Intentos
	}
	return 0
}

func (m *Request_Estado) GetFecha() string {
	if m != nil {
		return m.Fecha
	}
	return ""
}

func (m *Request_Estado) GetEstado() string {
	if m != nil {
		return m.Estado
	}
	return ""
}

type Response_Estado struct {
	Recibido             string   `protobuf:"bytes,1,opt,name=recibido,proto3" json:"recibido,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Estado) Reset()         { *m = Response_Estado{} }
func (m *Response_Estado) String() string { return proto.CompactTextString(m) }
func (*Response_Estado) ProtoMessage()    {}
func (*Response_Estado) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{8}
}

func (m *Response_Estado) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Estado.Unmarshal(m, b)
}
func (m *Response_Estado) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Estado.Marshal(b, m, deterministic)
}
func (m *Response_Estado) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Estado.Merge(m, src)
}
func (m *Response_Estado) XXX_Size() int {
	return xxx_messageInfo_Response_Estado.Size(m)
}
func (m *Response_Estado) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Estado.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Estado proto.InternalMessageInfo

func (m *Response_Estado) GetRecibido() string {
	if m != nil {
		return m.Recibido
	}
	return ""
}

type Dummy struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dummy) Reset()         { *m = Dummy{} }
func (m *Dummy) String() string { return proto.CompactTextString(m) }
func (*Dummy) ProtoMessage()    {}
func (*Dummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_db39efb7717b7d47, []int{9}
}

func (m *Dummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dummy.Unmarshal(m, b)
}
func (m *Dummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dummy.Marshal(b, m, deterministic)
}
func (m *Dummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dummy.Merge(m, src)
}
func (m *Dummy) XXX_Size() int {
	return xxx_messageInfo_Dummy.Size(m)
}
func (m *Dummy) XXX_DiscardUnknown() {
	xxx_messageInfo_Dummy.DiscardUnknown(m)
}

var xxx_messageInfo_Dummy proto.InternalMessageInfo

func (m *Dummy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Request_CrearOrdenPyme)(nil), "comms.Request_CrearOrdenPyme")
	proto.RegisterType((*Request_CrearOrdenRetail)(nil), "comms.Request_CrearOrdenRetail")
	proto.RegisterType((*Response_CrearOrden)(nil), "comms.Response_CrearOrden")
	proto.RegisterType((*Request_Seguimiento)(nil), "comms.Request_Seguimiento")
	proto.RegisterType((*Response_Seguimiento)(nil), "comms.Response_Seguimiento")
	proto.RegisterType((*Request_SolicitarPaquete)(nil), "comms.Request_SolicitarPaquete")
	proto.RegisterType((*Response_SolicitarPaquete)(nil), "comms.Response_SolicitarPaquete")
	proto.RegisterType((*Request_Estado)(nil), "comms.Request_Estado")
	proto.RegisterType((*Response_Estado)(nil), "comms.Response_Estado")
	proto.RegisterType((*Dummy)(nil), "comms.Dummy")
}

func init() {
	proto.RegisterFile("comms.proto", fileDescriptor_db39efb7717b7d47)
}

var fileDescriptor_db39efb7717b7d47 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xde, 0xb4, 0x4d, 0x81, 0x29, 0x2a, 0x95, 0x59, 0x4a, 0x08, 0x42, 0x54, 0x3e, 0x71, 0xa1,
	0x42, 0x70, 0xe0, 0x86, 0x84, 0x16, 0x24, 0x90, 0x40, 0x2c, 0xd9, 0x03, 0x47, 0xe4, 0x4d, 0xbc,
	0x8b, 0x51, 0x1d, 0x07, 0xc7, 0x45, 0xda, 0x97, 0xe0, 0x3d, 0xb8, 0xf0, 0x0e, 0xbc, 0x19, 0x8e,
	0xed, 0xba, 0xae, 0xd3, 0x0a, 0x6e, 0x9c, 0x9a, 0x6f, 0x3c, 0xf3, 0xcd, 0x37, 0x7f, 0x85, 0x49,
	0x29, 0x38, 0x6f, 0x97, 0x8d, 0x14, 0x4a, 0xa0, 0xd4, 0x00, 0xfc, 0x2b, 0x81, 0x79, 0x41, 0xbf,
	0xad, 0x69, 0xab, 0x3e, 0x9f, 0x48, 0x4a, 0xe4, 0x07, 0x59, 0xd1, 0xfa, 0xf4, 0x8a, 0x53, 0x34,
	0x85, 0x01, 0xab, 0xb2, 0x64, 0x91, 0x3c, 0xba, 0x51, 0xe8, 0x2f, 0x94, 0xc3, 0x75, 0x1d, 0x5a,
	0xad, 0x4b, 0x25, 0xb2, 0x81, 0xb1, 0x7a, 0x8c, 0x8e, 0x21, 0xfd, 0x4e, 0x56, 0x42, 0x66, 0x43,
	0xfd, 0x90, 0x16, 0x16, 0xa0, 0x39, 0x8c, 0x15, 0xa3, 0x75, 0x45, 0xb2, 0x91, 0xf1, 0x77, 0x08,
	0x65, 0x70, 0xad, 0xd2, 0x09, 0x59, 0x2d, 0xb2, 0xd4, 0x3c, 0x6c, 0x20, 0x5a, 0xc0, 0xa4, 0x91,
	0x4c, 0x48, 0xa6, 0x88, 0xfe, 0xc9, 0xc6, 0x86, 0x2d, 0x34, 0xe1, 0x1f, 0x09, 0x64, 0x7d, 0xc1,
	0x05, 0x55, 0x84, 0xad, 0xfe, 0x87, 0x64, 0xfc, 0x1c, 0x6e, 0x17, 0xb4, 0x6d, 0x44, 0xdd, 0xd2,
	0x40, 0x50, 0x57, 0x49, 0x4b, 0x2f, 0xd7, 0x8c, 0xeb, 0x78, 0x9d, 0x3d, 0xb1, 0x95, 0x04, 0x26,
	0x1b, 0x68, 0x0b, 0x39, 0xdb, 0x9a, 0xff, 0x21, 0x70, 0x09, 0xc7, 0x3e, 0x63, 0x18, 0xa9, 0xb5,
	0x6b, 0x32, 0x52, 0x09, 0xd7, 0x01, 0x87, 0xf0, 0x8b, 0x6d, 0xc7, 0xce, 0xc4, 0x8a, 0x95, 0x5d,
	0x27, 0x4f, 0x89, 0x36, 0x28, 0x8a, 0x10, 0x8c, 0x14, 0x6b, 0x36, 0x11, 0xe6, 0xdb, 0x75, 0x71,
	0xb0, 0xe9, 0x22, 0xfe, 0x99, 0xc0, 0xbd, 0x6d, 0xc2, 0x98, 0x21, 0xee, 0x79, 0xa4, 0x7f, 0xd0,
	0xd3, 0xef, 0x73, 0x0e, 0x83, 0x9c, 0x7e, 0x1a, 0xa3, 0xfd, 0xd3, 0x48, 0x0f, 0x4d, 0x63, 0xbc,
	0x3b, 0x8d, 0xaf, 0x30, 0xdd, 0xd4, 0xfa, 0xda, 0x54, 0xbf, 0x6f, 0x27, 0x58, 0xad, 0x3a, 0x21,
	0xad, 0x13, 0xe7, 0x71, 0xa7, 0xe2, 0x82, 0x96, 0x5f, 0x88, 0x93, 0x66, 0x41, 0xd0, 0xd7, 0xd1,
	0x4e, 0x5f, 0x1f, 0xc3, 0x2d, 0xdf, 0x16, 0x97, 0x4c, 0x93, 0x4b, 0x5a, 0xb2, 0x73, 0xe6, 0x87,
	0xe0, 0x31, 0xbe, 0x0b, 0xe9, 0xab, 0x35, 0xe7, 0x57, 0xb1, 0xa2, 0xa7, 0xbf, 0x87, 0x90, 0x9e,
	0x74, 0xd7, 0x88, 0xde, 0xc3, 0x34, 0x3a, 0xc2, 0x07, 0x4b, 0x7b, 0xb4, 0xfb, 0x6f, 0x34, 0xcf,
	0xfd, 0x73, 0x6f, 0x03, 0xf1, 0x11, 0xfa, 0x08, 0xb3, 0xde, 0x89, 0x3c, 0x3c, 0x48, 0x68, 0x1d,
	0xfe, 0x42, 0xf9, 0x06, 0x26, 0xe1, 0xca, 0xe5, 0x11, 0x5b, 0xf0, 0x96, 0xdf, 0x8f, 0x89, 0x82,
	0x47, 0xcd, 0xf4, 0x09, 0x66, 0xbd, 0x5d, 0x8a, 0xc5, 0xc5, 0x0e, 0xf9, 0xa2, 0xc7, 0x19, 0x79,
	0x68, 0xe2, 0x97, 0x30, 0x7d, 0x5b, 0x5f, 0x08, 0xc9, 0x89, 0x74, 0x53, 0xb9, 0x13, 0xd1, 0x5a,
	0x73, 0x3e, 0x8f, 0xc9, 0xac, 0x5d, 0x53, 0x3c, 0x81, 0xd9, 0x3b, 0xc6, 0x1b, 0x46, 0x64, 0x41,
	0x2f, 0x59, 0xab, 0xa4, 0xde, 0x8d, 0x9b, 0xce, 0xdb, 0xcc, 0x30, 0xdf, 0x41, 0xf8, 0xe8, 0x7c,
	0x6c, 0xfe, 0x55, 0x9f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xa8, 0x83, 0x58, 0x64, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommsClient is the client API for Comms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommsClient interface {
	CrearOrdenPyme(ctx context.Context, in *Request_CrearOrdenPyme, opts ...grpc.CallOption) (*Response_CrearOrden, error)
	CrearOrdenRetail(ctx context.Context, in *Request_CrearOrdenRetail, opts ...grpc.CallOption) (*Response_CrearOrden, error)
	Seguimiento(ctx context.Context, in *Request_Seguimiento, opts ...grpc.CallOption) (*Response_Seguimiento, error)
	SolicitarPaquete(ctx context.Context, in *Request_SolicitarPaquete, opts ...grpc.CallOption) (*Response_SolicitarPaquete, error)
	InformarEstado(ctx context.Context, in *Request_Estado, opts ...grpc.CallOption) (*Response_Estado, error)
	LimpiarRegistros(ctx context.Context, in *Dummy, opts ...grpc.CallOption) (*Dummy, error)
}

type commsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommsClient(cc grpc.ClientConnInterface) CommsClient {
	return &commsClient{cc}
}

func (c *commsClient) CrearOrdenPyme(ctx context.Context, in *Request_CrearOrdenPyme, opts ...grpc.CallOption) (*Response_CrearOrden, error) {
	out := new(Response_CrearOrden)
	err := c.cc.Invoke(ctx, "/comms.Comms/CrearOrdenPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) CrearOrdenRetail(ctx context.Context, in *Request_CrearOrdenRetail, opts ...grpc.CallOption) (*Response_CrearOrden, error) {
	out := new(Response_CrearOrden)
	err := c.cc.Invoke(ctx, "/comms.Comms/CrearOrdenRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) Seguimiento(ctx context.Context, in *Request_Seguimiento, opts ...grpc.CallOption) (*Response_Seguimiento, error) {
	out := new(Response_Seguimiento)
	err := c.cc.Invoke(ctx, "/comms.Comms/Seguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) SolicitarPaquete(ctx context.Context, in *Request_SolicitarPaquete, opts ...grpc.CallOption) (*Response_SolicitarPaquete, error) {
	out := new(Response_SolicitarPaquete)
	err := c.cc.Invoke(ctx, "/comms.Comms/SolicitarPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) InformarEstado(ctx context.Context, in *Request_Estado, opts ...grpc.CallOption) (*Response_Estado, error) {
	out := new(Response_Estado)
	err := c.cc.Invoke(ctx, "/comms.Comms/InformarEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsClient) LimpiarRegistros(ctx context.Context, in *Dummy, opts ...grpc.CallOption) (*Dummy, error) {
	out := new(Dummy)
	err := c.cc.Invoke(ctx, "/comms.Comms/LimpiarRegistros", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommsServer is the server API for Comms service.
type CommsServer interface {
	CrearOrdenPyme(context.Context, *Request_CrearOrdenPyme) (*Response_CrearOrden, error)
	CrearOrdenRetail(context.Context, *Request_CrearOrdenRetail) (*Response_CrearOrden, error)
	Seguimiento(context.Context, *Request_Seguimiento) (*Response_Seguimiento, error)
	SolicitarPaquete(context.Context, *Request_SolicitarPaquete) (*Response_SolicitarPaquete, error)
	InformarEstado(context.Context, *Request_Estado) (*Response_Estado, error)
	LimpiarRegistros(context.Context, *Dummy) (*Dummy, error)
}

// UnimplementedCommsServer can be embedded to have forward compatible implementations.
type UnimplementedCommsServer struct {
}

func (*UnimplementedCommsServer) CrearOrdenPyme(ctx context.Context, req *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrearOrdenPyme not implemented")
}
func (*UnimplementedCommsServer) CrearOrdenRetail(ctx context.Context, req *Request_CrearOrdenRetail) (*Response_CrearOrden, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrearOrdenRetail not implemented")
}
func (*UnimplementedCommsServer) Seguimiento(ctx context.Context, req *Request_Seguimiento) (*Response_Seguimiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seguimiento not implemented")
}
func (*UnimplementedCommsServer) SolicitarPaquete(ctx context.Context, req *Request_SolicitarPaquete) (*Response_SolicitarPaquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarPaquete not implemented")
}
func (*UnimplementedCommsServer) InformarEstado(ctx context.Context, req *Request_Estado) (*Response_Estado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InformarEstado not implemented")
}
func (*UnimplementedCommsServer) LimpiarRegistros(ctx context.Context, req *Dummy) (*Dummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimpiarRegistros not implemented")
}

func RegisterCommsServer(s *grpc.Server, srv CommsServer) {
	s.RegisterService(&_Comms_serviceDesc, srv)
}

func _Comms_CrearOrdenPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_CrearOrdenPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).CrearOrdenPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/CrearOrdenPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).CrearOrdenPyme(ctx, req.(*Request_CrearOrdenPyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_CrearOrdenRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_CrearOrdenRetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).CrearOrdenRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/CrearOrdenRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).CrearOrdenRetail(ctx, req.(*Request_CrearOrdenRetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_Seguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Seguimiento)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).Seguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/Seguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).Seguimiento(ctx, req.(*Request_Seguimiento))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_SolicitarPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_SolicitarPaquete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).SolicitarPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/SolicitarPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).SolicitarPaquete(ctx, req.(*Request_SolicitarPaquete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_InformarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Estado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).InformarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/InformarEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).InformarEstado(ctx, req.(*Request_Estado))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms_LimpiarRegistros_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dummy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsServer).LimpiarRegistros(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Comms/LimpiarRegistros",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsServer).LimpiarRegistros(ctx, req.(*Dummy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comms.Comms",
	HandlerType: (*CommsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrearOrdenPyme",
			Handler:    _Comms_CrearOrdenPyme_Handler,
		},
		{
			MethodName: "CrearOrdenRetail",
			Handler:    _Comms_CrearOrdenRetail_Handler,
		},
		{
			MethodName: "Seguimiento",
			Handler:    _Comms_Seguimiento_Handler,
		},
		{
			MethodName: "SolicitarPaquete",
			Handler:    _Comms_SolicitarPaquete_Handler,
		},
		{
			MethodName: "InformarEstado",
			Handler:    _Comms_InformarEstado_Handler,
		},
		{
			MethodName: "LimpiarRegistros",
			Handler:    _Comms_LimpiarRegistros_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms.proto",
}
