// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: runoschain/configstore/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgSetPort struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Dpid    uint64 `protobuf:"varint,2,opt,name=dpid,proto3" json:"dpid,omitempty"`
	Mac     string `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
	InPort  uint64 `protobuf:"varint,4,opt,name=inPort,proto3" json:"inPort,omitempty"`
}

func (m *MsgSetPort) Reset()         { *m = MsgSetPort{} }
func (m *MsgSetPort) String() string { return proto.CompactTextString(m) }
func (*MsgSetPort) ProtoMessage()    {}
func (*MsgSetPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_c674ca398ff72067, []int{0}
}
func (m *MsgSetPort) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPort.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPort.Merge(m, src)
}
func (m *MsgSetPort) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPort) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPort.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPort proto.InternalMessageInfo

func (m *MsgSetPort) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetPort) GetDpid() uint64 {
	if m != nil {
		return m.Dpid
	}
	return 0
}

func (m *MsgSetPort) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *MsgSetPort) GetInPort() uint64 {
	if m != nil {
		return m.InPort
	}
	return 0
}

type MsgSetPortResponse struct {
}

func (m *MsgSetPortResponse) Reset()         { *m = MsgSetPortResponse{} }
func (m *MsgSetPortResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetPortResponse) ProtoMessage()    {}
func (*MsgSetPortResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c674ca398ff72067, []int{1}
}
func (m *MsgSetPortResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPortResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPortResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPortResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPortResponse.Merge(m, src)
}
func (m *MsgSetPortResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPortResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPortResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPortResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetPort)(nil), "runos_chain.configstore.MsgSetPort")
	proto.RegisterType((*MsgSetPortResponse)(nil), "runos_chain.configstore.MsgSetPortResponse")
}

func init() { proto.RegisterFile("runoschain/configstore/tx.proto", fileDescriptor_c674ca398ff72067) }

var fileDescriptor_c674ca398ff72067 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2a, 0xcd, 0xcb,
	0x2f, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2e, 0xc9,
	0x2f, 0x4a, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x07, 0x2b, 0x88,
	0x07, 0xab, 0xd0, 0x43, 0x52, 0xa1, 0x94, 0xc2, 0xc5, 0xe5, 0x5b, 0x9c, 0x1e, 0x9c, 0x5a, 0x12,
	0x90, 0x5f, 0x54, 0x22, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x14, 0x64, 0xa6, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0x89, 0xc9, 0x12,
	0xcc, 0x60, 0x95, 0x20, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x66, 0x1e, 0xc8, 0x24, 0x09, 0x16, 0xb0,
	0x3a, 0x28, 0x4f, 0x49, 0x84, 0x4b, 0x08, 0x61, 0x4b, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x51, 0x12, 0x17, 0xb3, 0x6f, 0x71, 0xba, 0x50, 0x34, 0x17, 0x3b, 0xcc, 0x7e, 0x65, 0x3d,
	0x1c, 0xee, 0xd4, 0x43, 0x68, 0x97, 0xd2, 0x26, 0x42, 0x11, 0xcc, 0x0e, 0x27, 0xcb, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x47, 0x32, 0x45, 0xbf, 0x02, 0x35, 0xd8,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x41, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0xe9, 0x8a, 0x4f, 0x5d, 0x01, 0x00, 0x00,
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
	SetPort(ctx context.Context, in *MsgSetPort, opts ...grpc.CallOption) (*MsgSetPortResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetPort(ctx context.Context, in *MsgSetPort, opts ...grpc.CallOption) (*MsgSetPortResponse, error) {
	out := new(MsgSetPortResponse)
	err := c.cc.Invoke(ctx, "/runos_chain.configstore.Msg/SetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetPort(context.Context, *MsgSetPort) (*MsgSetPortResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetPort(ctx context.Context, req *MsgSetPort) (*MsgSetPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPort not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetPort)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runos_chain.configstore.Msg/SetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetPort(ctx, req.(*MsgSetPort))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runos_chain.configstore.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPort",
			Handler:    _Msg_SetPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runoschain/configstore/tx.proto",
}

func (m *MsgSetPort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPort) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPort) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InPort != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.InPort))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Mac) > 0 {
		i -= len(m.Mac)
		copy(dAtA[i:], m.Mac)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Mac)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Dpid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Dpid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetPortResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPortResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPortResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSetPort) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Dpid != 0 {
		n += 1 + sovTx(uint64(m.Dpid))
	}
	l = len(m.Mac)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InPort != 0 {
		n += 1 + sovTx(uint64(m.InPort))
	}
	return n
}

func (m *MsgSetPortResponse) Size() (n int) {
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
func (m *MsgSetPort) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetPort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dpid", wireType)
			}
			m.Dpid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dpid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mac", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mac = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InPort", wireType)
			}
			m.InPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InPort |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgSetPortResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetPortResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPortResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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