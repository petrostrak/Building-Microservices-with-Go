// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package currency

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

// Currencies is an enum which represents the allowed currencies for the API
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

var Currencies_name = map[int32]string{
	0:  "EUR",
	1:  "USD",
	2:  "JPY",
	3:  "BGN",
	4:  "CZK",
	5:  "DKK",
	6:  "GBP",
	7:  "HUF",
	8:  "PLN",
	9:  "RON",
	10: "SEK",
	11: "CHF",
	12: "ISK",
	13: "NOK",
	14: "HRK",
	15: "RUB",
	16: "TRY",
	17: "AUD",
	18: "BRL",
	19: "CAD",
	20: "CNY",
	21: "HKD",
	22: "IDR",
	23: "ILS",
	24: "INR",
	25: "KRW",
	26: "MXN",
	27: "MYR",
	28: "NZD",
	29: "PHP",
	30: "SGD",
	31: "THB",
	32: "ZAR",
}

var Currencies_value = map[string]int32{
	"EUR": 0,
	"USD": 1,
	"JPY": 2,
	"BGN": 3,
	"CZK": 4,
	"DKK": 5,
	"GBP": 6,
	"HUF": 7,
	"PLN": 8,
	"RON": 9,
	"SEK": 10,
	"CHF": 11,
	"ISK": 12,
	"NOK": 13,
	"HRK": 14,
	"RUB": 15,
	"TRY": 16,
	"AUD": 17,
	"BRL": 18,
	"CAD": 19,
	"CNY": 20,
	"HKD": 21,
	"IDR": 22,
	"ILS": 23,
	"INR": 24,
	"KRW": 25,
	"MXN": 26,
	"MYR": 27,
	"NZD": 28,
	"PHP": 29,
	"SGD": 30,
	"THB": 31,
	"ZAR": 32,
}

func (x Currencies) String() string {
	return proto.EnumName(Currencies_name, int32(x))
}

func (Currencies) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination          Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RateRequest) Reset()         { *m = RateRequest{} }
func (m *RateRequest) String() string { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()    {}
func (*RateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *RateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateRequest.Unmarshal(m, b)
}
func (m *RateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateRequest.Marshal(b, m, deterministic)
}
func (m *RateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateRequest.Merge(m, src)
}
func (m *RateRequest) XXX_Size() int {
	return xxx_messageInfo_RateRequest.Size(m)
}
func (m *RateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateRequest proto.InternalMessageInfo

func (m *RateRequest) GetBase() Currencies {
	if m != nil {
		return m.Base
	}
	return Currencies_EUR
}

func (m *RateRequest) GetDestination() Currencies {
	if m != nil {
		return m.Destination
	}
	return Currencies_EUR
}

// RateResponse is the response from a GetRate call, it contains
// rate which is a floating point number and can be used to convert between the
// two currencies specified in the request.
type RateResponse struct {
	Rate                 float32  `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateResponse) Reset()         { *m = RateResponse{} }
func (m *RateResponse) String() string { return proto.CompactTextString(m) }
func (*RateResponse) ProtoMessage()    {}
func (*RateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{1}
}

func (m *RateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateResponse.Unmarshal(m, b)
}
func (m *RateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateResponse.Marshal(b, m, deterministic)
}
func (m *RateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateResponse.Merge(m, src)
}
func (m *RateResponse) XXX_Size() int {
	return xxx_messageInfo_RateResponse.Size(m)
}
func (m *RateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateResponse proto.InternalMessageInfo

func (m *RateResponse) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterEnum("Currencies", Currencies_name, Currencies_value)
	proto.RegisterType((*RateRequest)(nil), "RateRequest")
	proto.RegisterType((*RateResponse)(nil), "RateResponse")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd1, 0xcd, 0xee, 0xd2, 0x40,
	0x14, 0x05, 0x70, 0x29, 0x08, 0x38, 0x05, 0x3c, 0x8e, 0x5f, 0x88, 0x1f, 0x10, 0x16, 0xc6, 0x98,
	0xd8, 0x05, 0x3e, 0x01, 0x65, 0xa0, 0xd5, 0xc1, 0xd2, 0x4c, 0x6d, 0xb4, 0x24, 0x2e, 0x2a, 0x99,
	0x05, 0x9b, 0x16, 0xdb, 0xb2, 0xe0, 0xc1, 0x7c, 0x3f, 0x73, 0x2f, 0x26, 0x7f, 0x92, 0xff, 0xee,
	0x97, 0xcc, 0x9c, 0x9b, 0x93, 0x7b, 0xc5, 0xe8, 0x70, 0xae, 0x2a, 0x5b, 0x1c, 0x2e, 0xde, 0xa9,
	0x2a, 0x9b, 0x72, 0xfe, 0x4b, 0xb8, 0x26, 0x6f, 0xac, 0xb1, 0x7f, 0xce, 0xb6, 0x6e, 0xe4, 0x54,
	0x74, 0xfc, 0xbc, 0xb6, 0xe3, 0xd6, 0xac, 0xf5, 0x61, 0xb4, 0x70, 0xbd, 0xd5, 0xf5, 0xf7, 0xd1,
	0xd6, 0x86, 0x1f, 0xe4, 0x27, 0xe1, 0x2a, 0x5b, 0x37, 0xc7, 0x22, 0x6f, 0x8e, 0x65, 0x31, 0x76,
	0xee, 0xff, 0xbb, 0x7d, 0x9f, 0xcf, 0xc5, 0xe0, 0x3a, 0xbe, 0x3e, 0x95, 0x45, 0x6d, 0xa5, 0x14,
	0x9d, 0x2a, 0x6f, 0xae, 0xf3, 0x1d, 0xc3, 0xfe, 0xf8, 0xd7, 0x11, 0xe2, 0x2e, 0x2f, 0x7b, 0xa2,
	0xbd, 0x4e, 0x0d, 0x1e, 0x10, 0xd2, 0x44, 0xa1, 0x45, 0xf8, 0x1a, 0x67, 0x70, 0x08, 0x7e, 0x10,
	0xa1, 0x4d, 0x58, 0xed, 0x35, 0x3a, 0x04, 0xa5, 0x35, 0x1e, 0x12, 0x02, 0x3f, 0x46, 0x97, 0x10,
	0xa6, 0x1b, 0xf4, 0x08, 0xf1, 0x36, 0x42, 0x9f, 0x60, 0x76, 0x11, 0x1e, 0x11, 0x92, 0xb5, 0x86,
	0xe0, 0x78, 0xb8, 0x81, 0x4b, 0xf8, 0x92, 0x68, 0x0c, 0x08, 0xd1, 0x4e, 0x63, 0xc8, 0x71, 0xa3,
	0x31, 0xe2, 0x54, 0xea, 0xe3, 0x31, 0xe1, 0xbb, 0xc9, 0x00, 0xc2, 0x32, 0x55, 0x78, 0xc2, 0x35,
	0xcc, 0x16, 0x92, 0xe7, 0x2c, 0x15, 0x9e, 0x32, 0xa2, 0x0c, 0xcf, 0x38, 0xae, 0x15, 0x9e, 0xf3,
	0x64, 0x65, 0xf0, 0x82, 0xb1, 0x4d, 0xf0, 0x92, 0x11, 0x19, 0x8c, 0x09, 0xda, 0xfc, 0xc0, 0x2b,
	0xc2, 0xb7, 0x9f, 0x11, 0x26, 0x8c, 0xcc, 0xe0, 0x35, 0xd7, 0xd8, 0x2b, 0xbc, 0xe1, 0xf2, 0x61,
	0x8c, 0xb7, 0xdc, 0x39, 0x50, 0x78, 0xc7, 0x35, 0x42, 0x1f, 0x53, 0xc2, 0x7e, 0x69, 0x30, 0x5b,
	0x2c, 0x44, 0xff, 0xff, 0xda, 0x2e, 0xf2, 0xbd, 0xe8, 0x05, 0xb6, 0xa1, 0x55, 0xcb, 0x81, 0x77,
	0x73, 0xd0, 0xc9, 0xd0, 0xbb, 0xdd, 0xff, 0xef, 0x2e, 0x5f, 0xfd, 0xf3, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0xc9, 0x43, 0xb5, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	// GetRate returns the exchange rate for the two provided currency codes
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(ctx context.Context, req *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}