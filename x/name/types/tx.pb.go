// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/name/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgBindNameRequest defines an sdk.Msg type that is used to add an address/name binding under an optional parent name.
// The record may optionally be restricted to prevent additional names from being added under this one without the
// owner signing the request.
type MsgBindNameRequest struct {
	// The parent record to bind this name under.
	Parent NameRecord `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent"`
	// The name record to bind under the parent
	Record NameRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record"`
}

func (m *MsgBindNameRequest) Reset()         { *m = MsgBindNameRequest{} }
func (m *MsgBindNameRequest) String() string { return proto.CompactTextString(m) }
func (*MsgBindNameRequest) ProtoMessage()    {}
func (*MsgBindNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eacf6cd967218635, []int{0}
}
func (m *MsgBindNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBindNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBindNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBindNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBindNameRequest.Merge(m, src)
}
func (m *MsgBindNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgBindNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBindNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBindNameRequest proto.InternalMessageInfo

// MsgBindNameResponse defines the Msg/BindName response type.
type MsgBindNameResponse struct {
}

func (m *MsgBindNameResponse) Reset()         { *m = MsgBindNameResponse{} }
func (m *MsgBindNameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBindNameResponse) ProtoMessage()    {}
func (*MsgBindNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eacf6cd967218635, []int{1}
}
func (m *MsgBindNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBindNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBindNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBindNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBindNameResponse.Merge(m, src)
}
func (m *MsgBindNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBindNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBindNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBindNameResponse proto.InternalMessageInfo

// MsgDeleteNameRequest defines an sdk.Msg type that is used to remove an existing address/name binding.  The binding
// may not have any child names currently bound for this request to be successful.
type MsgDeleteNameRequest struct {
	// The parent record the record to remove is under.
	Parent NameRecord `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent"`
	// The record being removed
	Record NameRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record"`
}

func (m *MsgDeleteNameRequest) Reset()         { *m = MsgDeleteNameRequest{} }
func (m *MsgDeleteNameRequest) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteNameRequest) ProtoMessage()    {}
func (*MsgDeleteNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eacf6cd967218635, []int{2}
}
func (m *MsgDeleteNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteNameRequest.Merge(m, src)
}
func (m *MsgDeleteNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteNameRequest proto.InternalMessageInfo

// MsgDeleteNameResponse defines the Msg/DeleteName response type.
type MsgDeleteNameResponse struct {
}

func (m *MsgDeleteNameResponse) Reset()         { *m = MsgDeleteNameResponse{} }
func (m *MsgDeleteNameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteNameResponse) ProtoMessage()    {}
func (*MsgDeleteNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eacf6cd967218635, []int{3}
}
func (m *MsgDeleteNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteNameResponse.Merge(m, src)
}
func (m *MsgDeleteNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteNameResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgBindNameRequest)(nil), "provenance.name.v1.MsgBindNameRequest")
	proto.RegisterType((*MsgBindNameResponse)(nil), "provenance.name.v1.MsgBindNameResponse")
	proto.RegisterType((*MsgDeleteNameRequest)(nil), "provenance.name.v1.MsgDeleteNameRequest")
	proto.RegisterType((*MsgDeleteNameResponse)(nil), "provenance.name.v1.MsgDeleteNameResponse")
}

func init() { proto.RegisterFile("provenance/name/v1/tx.proto", fileDescriptor_eacf6cd967218635) }

var fileDescriptor_eacf6cd967218635 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0xcf, 0x4b, 0xcc, 0x4d, 0xd5, 0x2f, 0x33, 0xd4, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x48, 0xea, 0x81, 0x24, 0xf5, 0xca,
	0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x94, 0x2c,
	0x16, 0x63, 0xc0, 0x3a, 0xc0, 0xd2, 0x4a, 0xb3, 0x18, 0xb9, 0x84, 0x7c, 0x8b, 0xd3, 0x9d, 0x32,
	0xf3, 0x52, 0xfc, 0x12, 0x73, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x6c, 0xb8,
	0xd8, 0x0a, 0x12, 0x8b, 0x52, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xe4, 0xf4,
	0x30, 0x2d, 0xd4, 0x83, 0x68, 0x48, 0xce, 0x2f, 0x4a, 0x71, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21,
	0x08, 0xaa, 0x07, 0xa4, 0xbb, 0x08, 0x2c, 0x2e, 0xc1, 0x44, 0x8a, 0x6e, 0x88, 0x1e, 0x2b, 0x8e,
	0x8e, 0x05, 0xf2, 0x0c, 0x2f, 0x16, 0xc8, 0x33, 0x28, 0x89, 0x72, 0x09, 0xa3, 0xb8, 0xad, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x55, 0x69, 0x0e, 0x23, 0x97, 0x88, 0x6f, 0x71, 0xba, 0x4b, 0x6a, 0x4e,
	0x6a, 0x49, 0xea, 0xe0, 0x73, 0xb5, 0x38, 0x97, 0x28, 0x9a, 0xeb, 0x20, 0xee, 0x36, 0x3a, 0xca,
	0xc8, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x14, 0xcd, 0xc5, 0x01, 0xf3, 0x93, 0x90, 0x1a, 0x36, 0x4b,
	0x30, 0x23, 0x44, 0x4a, 0x9d, 0xa0, 0x3a, 0x88, 0x25, 0x42, 0x89, 0x5c, 0x5c, 0x08, 0xab, 0x85,
	0x34, 0x70, 0x68, 0xc3, 0x08, 0x3b, 0x29, 0x4d, 0x22, 0x54, 0x42, 0xac, 0x70, 0x2a, 0x3e, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xd1, 0xcc, 0x7c, 0x2c, 0xc6, 0x04, 0x30,
	0x46, 0x39, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xbb, 0x65, 0xa6,
	0x97, 0x16, 0xa5, 0x86, 0xa4, 0x26, 0x67, 0xe4, 0xe5, 0xe7, 0xe4, 0xa7, 0x67, 0xa6, 0x16, 0xeb,
	0x23, 0xf4, 0xe8, 0x26, 0xe5, 0xe4, 0x27, 0x67, 0x27, 0x67, 0x24, 0x66, 0xe6, 0xe9, 0x57, 0x40,
	0x12, 0x6c, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xbd, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x44, 0x5a, 0xad, 0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// BindName binds a name to an address under a root name.
	BindName(ctx context.Context, in *MsgBindNameRequest, opts ...grpc.CallOption) (*MsgBindNameResponse, error)
	// DeleteName defines a method to verify a particular invariance.
	DeleteName(ctx context.Context, in *MsgDeleteNameRequest, opts ...grpc.CallOption) (*MsgDeleteNameResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) BindName(ctx context.Context, in *MsgBindNameRequest, opts ...grpc.CallOption) (*MsgBindNameResponse, error) {
	out := new(MsgBindNameResponse)
	err := c.cc.Invoke(ctx, "/provenance.name.v1.Msg/BindName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteName(ctx context.Context, in *MsgDeleteNameRequest, opts ...grpc.CallOption) (*MsgDeleteNameResponse, error) {
	out := new(MsgDeleteNameResponse)
	err := c.cc.Invoke(ctx, "/provenance.name.v1.Msg/DeleteName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// BindName binds a name to an address under a root name.
	BindName(context.Context, *MsgBindNameRequest) (*MsgBindNameResponse, error)
	// DeleteName defines a method to verify a particular invariance.
	DeleteName(context.Context, *MsgDeleteNameRequest) (*MsgDeleteNameResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) BindName(ctx context.Context, req *MsgBindNameRequest) (*MsgBindNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindName not implemented")
}
func (*UnimplementedMsgServer) DeleteName(ctx context.Context, req *MsgDeleteNameRequest) (*MsgDeleteNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteName not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_BindName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBindNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BindName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.name.v1.Msg/BindName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BindName(ctx, req.(*MsgBindNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.name.v1.Msg/DeleteName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteName(ctx, req.(*MsgDeleteNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.name.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BindName",
			Handler:    _Msg_BindName_Handler,
		},
		{
			MethodName: "DeleteName",
			Handler:    _Msg_DeleteName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/name/v1/tx.proto",
}

func (m *MsgBindNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBindNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBindNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Parent.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgBindNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBindNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBindNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Parent.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgDeleteNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBindNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Parent.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Record.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBindNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Parent.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Record.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgDeleteNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBindNameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBindNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBindNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Parent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBindNameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBindNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBindNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteNameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Parent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeleteNameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeleteNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
