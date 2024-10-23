// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/globalfee/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryGasPrices struct {
}

func (m *QueryGasPrices) Reset()         { *m = QueryGasPrices{} }
func (m *QueryGasPrices) String() string { return proto.CompactTextString(m) }
func (*QueryGasPrices) ProtoMessage()    {}
func (*QueryGasPrices) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1fc751239d71db, []int{0}
}
func (m *QueryGasPrices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGasPrices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGasPrices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGasPrices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGasPrices.Merge(m, src)
}
func (m *QueryGasPrices) XXX_Size() int {
	return m.Size()
}
func (m *QueryGasPrices) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGasPrices.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGasPrices proto.InternalMessageInfo

type QueryGasPricesResponse struct {
	GasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=gas_prices,json=gasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"gas_prices"`
}

func (m *QueryGasPricesResponse) Reset()         { *m = QueryGasPricesResponse{} }
func (m *QueryGasPricesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGasPricesResponse) ProtoMessage()    {}
func (*QueryGasPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1fc751239d71db, []int{1}
}
func (m *QueryGasPricesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGasPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGasPricesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGasPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGasPricesResponse.Merge(m, src)
}
func (m *QueryGasPricesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGasPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGasPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGasPricesResponse proto.InternalMessageInfo

func (m *QueryGasPricesResponse) GetGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.GasPrices
	}
	return nil
}

type QueryBypassMessages struct {
}

func (m *QueryBypassMessages) Reset()         { *m = QueryBypassMessages{} }
func (m *QueryBypassMessages) String() string { return proto.CompactTextString(m) }
func (*QueryBypassMessages) ProtoMessage()    {}
func (*QueryBypassMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1fc751239d71db, []int{2}
}
func (m *QueryBypassMessages) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBypassMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBypassMessages.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBypassMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBypassMessages.Merge(m, src)
}
func (m *QueryBypassMessages) XXX_Size() int {
	return m.Size()
}
func (m *QueryBypassMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBypassMessages.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBypassMessages proto.InternalMessageInfo

type QueryBypassMessagesResponse struct {
	BypassMessages []string `protobuf:"bytes,1,rep,name=bypass_messages,json=bypassMessages,proto3" json:"bypass_messages,omitempty"`
}

func (m *QueryBypassMessagesResponse) Reset()         { *m = QueryBypassMessagesResponse{} }
func (m *QueryBypassMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBypassMessagesResponse) ProtoMessage()    {}
func (*QueryBypassMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1fc751239d71db, []int{3}
}
func (m *QueryBypassMessagesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBypassMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBypassMessagesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBypassMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBypassMessagesResponse.Merge(m, src)
}
func (m *QueryBypassMessagesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBypassMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBypassMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBypassMessagesResponse proto.InternalMessageInfo

func (m *QueryBypassMessagesResponse) GetBypassMessages() []string {
	if m != nil {
		return m.BypassMessages
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGasPrices)(nil), "noble.globalfee.v1.QueryGasPrices")
	proto.RegisterType((*QueryGasPricesResponse)(nil), "noble.globalfee.v1.QueryGasPricesResponse")
	proto.RegisterType((*QueryBypassMessages)(nil), "noble.globalfee.v1.QueryBypassMessages")
	proto.RegisterType((*QueryBypassMessagesResponse)(nil), "noble.globalfee.v1.QueryBypassMessagesResponse")
}

func init() { proto.RegisterFile("noble/globalfee/v1/query.proto", fileDescriptor_8d1fc751239d71db) }

var fileDescriptor_8d1fc751239d71db = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x95, 0x0a, 0x3b, 0xc2, 0xaa, 0xe3, 0x0f, 0x24, 0x2d, 0xd3, 0x12, 0x91, 0x4a,
	0x4b, 0x67, 0xdc, 0x7a, 0xf1, 0x6a, 0x14, 0x3c, 0x15, 0xb4, 0x47, 0x2f, 0x65, 0x26, 0x8e, 0x63,
	0x30, 0x99, 0x17, 0xf7, 0x65, 0x0b, 0x7b, 0xf5, 0xd4, 0x9b, 0x82, 0x57, 0xff, 0x00, 0xf1, 0xa4,
	0xff, 0x45, 0x8f, 0x05, 0x2f, 0x9e, 0x54, 0x76, 0x05, 0xff, 0x0d, 0xc9, 0x24, 0xcd, 0x6e, 0xd6,
	0x15, 0xbd, 0x24, 0xc3, 0xfb, 0xbe, 0xf9, 0xbe, 0xcf, 0x7b, 0x6f, 0x28, 0x77, 0xa0, 0x33, 0x23,
	0x6d, 0x06, 0x5a, 0x65, 0xcf, 0x8d, 0x91, 0x47, 0x03, 0xf9, 0x6a, 0x64, 0x86, 0x63, 0x51, 0x0c,
	0xa1, 0x04, 0xc6, 0xbc, 0x2e, 0x5a, 0x5d, 0x1c, 0x0d, 0xc2, 0xcb, 0x2a, 0x4f, 0x1d, 0x48, 0xff,
	0xad, 0xd3, 0x42, 0x9e, 0x00, 0xe6, 0x80, 0x52, 0x2b, 0xac, 0x2c, 0xb4, 0x29, 0xd5, 0x40, 0x26,
	0x90, 0xba, 0x46, 0x5f, 0x6b, 0x74, 0x6f, 0xbd, 0x50, 0x23, 0xbc, 0x6a, 0xc1, 0x82, 0x3f, 0xca,
	0xea, 0xd4, 0x44, 0xd7, 0x2d, 0x80, 0xcd, 0x8c, 0x54, 0x45, 0x2a, 0x95, 0x73, 0x50, 0xaa, 0x32,
	0x05, 0x87, 0xb5, 0x1a, 0x5d, 0xa2, 0xfd, 0x27, 0x95, 0xc5, 0x23, 0x85, 0x8f, 0x87, 0x69, 0x62,
	0x30, 0x7a, 0x43, 0xe8, 0xf5, 0x6e, 0xe8, 0xc0, 0x60, 0x01, 0x0e, 0x0d, 0x1b, 0x51, 0x6a, 0x15,
	0x1e, 0x16, 0x3e, 0x7a, 0x83, 0x6c, 0x9e, 0xbb, 0x7d, 0x61, 0x6f, 0x5d, 0xd4, 0x48, 0xa2, 0x42,
	0x16, 0x0d, 0xb2, 0x78, 0x68, 0x92, 0x07, 0x90, 0xba, 0xf8, 0xde, 0xc9, 0xb7, 0x8d, 0xe0, 0xe3,
	0xf7, 0x8d, 0x1d, 0x9b, 0x96, 0x2f, 0x46, 0x5a, 0x24, 0x90, 0xcb, 0xa6, 0x85, 0xfa, 0xb7, 0x8b,
	0xcf, 0x5e, 0xca, 0x72, 0x5c, 0x18, 0x3c, 0xbb, 0x83, 0x1f, 0x7e, 0x7d, 0xda, 0x26, 0x07, 0x3d,
	0xdb, 0x12, 0x5d, 0xa3, 0x57, 0x3c, 0x50, 0x3c, 0x2e, 0x14, 0xe2, 0xbe, 0x41, 0x54, 0xd6, 0x60,
	0xb4, 0x4f, 0xd7, 0x96, 0x84, 0x5b, 0x58, 0x41, 0x2f, 0x6a, 0xaf, 0x1c, 0xe6, 0x8d, 0xe4, 0x89,
	0x7b, 0xf1, 0x6a, 0x5d, 0xa0, 0xaf, 0x3b, 0xf7, 0xf6, 0x3e, 0xaf, 0xd0, 0x55, 0xef, 0xc7, 0x8e,
	0x09, 0xed, 0xb5, 0xcd, 0xb3, 0x48, 0xfc, 0xb9, 0x3a, 0xd1, 0x1d, 0x50, 0xb8, 0xfd, 0xef, 0x9c,
	0x33, 0xae, 0x68, 0xe7, 0xb8, 0x2a, 0xff, 0xfa, 0xcb, 0xcf, 0x77, 0x2b, 0x9b, 0x8c, 0xcb, 0x25,
	0xef, 0x66, 0x36, 0x63, 0xf6, 0x9e, 0xd0, 0x7e, 0xb7, 0x3f, 0xb6, 0xf5, 0xd7, 0x5a, 0xdd, 0xc4,
	0x50, 0xfe, 0x67, 0x62, 0x4b, 0x76, 0x67, 0x46, 0x76, 0x8b, 0xdd, 0x5c, 0x46, 0xb6, 0x30, 0xd0,
	0xf8, 0xfe, 0xc9, 0x84, 0x93, 0xd3, 0x09, 0x27, 0x3f, 0x26, 0x9c, 0xbc, 0x9d, 0xf2, 0xe0, 0x74,
	0xca, 0x83, 0xaf, 0x53, 0x1e, 0x3c, 0xdd, 0x9a, 0x5b, 0xb8, 0x37, 0xda, 0x55, 0x88, 0xa6, 0xc4,
	0x39, 0x3f, 0xbf, 0x75, 0x7d, 0xde, 0xbf, 0xc3, 0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0xc9, 0x62, 0x6c, 0x41, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	GasPrices(ctx context.Context, in *QueryGasPrices, opts ...grpc.CallOption) (*QueryGasPricesResponse, error)
	BypassMessages(ctx context.Context, in *QueryBypassMessages, opts ...grpc.CallOption) (*QueryBypassMessagesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GasPrices(ctx context.Context, in *QueryGasPrices, opts ...grpc.CallOption) (*QueryGasPricesResponse, error) {
	out := new(QueryGasPricesResponse)
	err := c.cc.Invoke(ctx, "/noble.globalfee.v1.Query/GasPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BypassMessages(ctx context.Context, in *QueryBypassMessages, opts ...grpc.CallOption) (*QueryBypassMessagesResponse, error) {
	out := new(QueryBypassMessagesResponse)
	err := c.cc.Invoke(ctx, "/noble.globalfee.v1.Query/BypassMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GasPrices(context.Context, *QueryGasPrices) (*QueryGasPricesResponse, error)
	BypassMessages(context.Context, *QueryBypassMessages) (*QueryBypassMessagesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GasPrices(ctx context.Context, req *QueryGasPrices) (*QueryGasPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GasPrices not implemented")
}
func (*UnimplementedQueryServer) BypassMessages(ctx context.Context, req *QueryBypassMessages) (*QueryBypassMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BypassMessages not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GasPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGasPrices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GasPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noble.globalfee.v1.Query/GasPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GasPrices(ctx, req.(*QueryGasPrices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BypassMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBypassMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BypassMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noble.globalfee.v1.Query/BypassMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BypassMessages(ctx, req.(*QueryBypassMessages))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noble.globalfee.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GasPrices",
			Handler:    _Query_GasPrices_Handler,
		},
		{
			MethodName: "BypassMessages",
			Handler:    _Query_BypassMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/globalfee/v1/query.proto",
}

func (m *QueryGasPrices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGasPrices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGasPrices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGasPricesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGasPricesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGasPricesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for iNdEx := len(m.GasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryBypassMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBypassMessages) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBypassMessages) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBypassMessagesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBypassMessagesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBypassMessagesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BypassMessages) > 0 {
		for iNdEx := len(m.BypassMessages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BypassMessages[iNdEx])
			copy(dAtA[i:], m.BypassMessages[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.BypassMessages[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGasPrices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGasPricesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for _, e := range m.GasPrices {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryBypassMessages) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBypassMessagesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BypassMessages) > 0 {
		for _, s := range m.BypassMessages {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGasPrices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGasPrices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGasPrices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGasPricesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGasPricesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGasPricesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = append(m.GasPrices, types.DecCoin{})
			if err := m.GasPrices[len(m.GasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBypassMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBypassMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBypassMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBypassMessagesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBypassMessagesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBypassMessagesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BypassMessages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BypassMessages = append(m.BypassMessages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
