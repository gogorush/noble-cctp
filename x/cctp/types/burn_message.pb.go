// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/burn_message.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// *
// Message format for BurnMessages
// @param version the message body version
// @param burn_token the burn token address on source domain as bytes32
// @param mint_recipient the mint recipient address as bytes32
// @param amount the burn amount
// @param message_sender the message sender
type BurnMessage struct {
	Version       uint32                `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	BurnToken     []byte                `protobuf:"bytes,2,opt,name=burn_token,json=burnToken,proto3" json:"burn_token,omitempty"`
	MintRecipient []byte                `protobuf:"bytes,3,opt,name=mint_recipient,json=mintRecipient,proto3" json:"mint_recipient,omitempty"`
	Amount        cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	MessageSender []byte                `protobuf:"bytes,5,opt,name=message_sender,json=messageSender,proto3" json:"message_sender,omitempty"`
}

func (m *BurnMessage) Reset()         { *m = BurnMessage{} }
func (m *BurnMessage) String() string { return proto.CompactTextString(m) }
func (*BurnMessage) ProtoMessage()    {}
func (*BurnMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_770538a2f8e64d6c, []int{0}
}
func (m *BurnMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnMessage.Merge(m, src)
}
func (m *BurnMessage) XXX_Size() int {
	return m.Size()
}
func (m *BurnMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BurnMessage proto.InternalMessageInfo

func (m *BurnMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BurnMessage) GetBurnToken() []byte {
	if m != nil {
		return m.BurnToken
	}
	return nil
}

func (m *BurnMessage) GetMintRecipient() []byte {
	if m != nil {
		return m.MintRecipient
	}
	return nil
}

func (m *BurnMessage) GetMessageSender() []byte {
	if m != nil {
		return m.MessageSender
	}
	return nil
}

func init() {
	proto.RegisterType((*BurnMessage)(nil), "circle.cctp.v1.BurnMessage")
}

func init() { proto.RegisterFile("circle/cctp/v1/burn_message.proto", fileDescriptor_770538a2f8e64d6c) }

var fileDescriptor_770538a2f8e64d6c = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x41, 0x4a, 0x33, 0x31,
	0x18, 0x86, 0x27, 0xff, 0xaf, 0x95, 0x46, 0xdb, 0xc5, 0xa0, 0x10, 0x84, 0xa6, 0x55, 0x10, 0x0a,
	0x62, 0x42, 0x11, 0x2f, 0x50, 0x70, 0xe1, 0xc2, 0xcd, 0xe8, 0xca, 0x4d, 0xe9, 0xa4, 0x61, 0x1a,
	0xda, 0xe4, 0x1b, 0x92, 0x4c, 0xd1, 0x5b, 0x78, 0xac, 0x82, 0x9b, 0x2e, 0xc5, 0x45, 0x91, 0xce,
	0x45, 0x64, 0x92, 0x76, 0x97, 0x3c, 0x79, 0xf2, 0x26, 0xdf, 0x8b, 0xaf, 0x84, 0xb2, 0x62, 0x29,
	0xb9, 0x10, 0xbe, 0xe4, 0xab, 0x11, 0xcf, 0x2b, 0x6b, 0x26, 0x5a, 0x3a, 0x37, 0x2d, 0x24, 0x2b,
	0x2d, 0x78, 0x48, 0xbb, 0x51, 0x61, 0x8d, 0xc2, 0x56, 0xa3, 0xcb, 0xf3, 0x02, 0x0a, 0x08, 0x47,
	0xbc, 0x59, 0x45, 0xeb, 0xfa, 0x0b, 0xe1, 0xd3, 0x71, 0x65, 0xcd, 0x73, 0xbc, 0x9b, 0x12, 0x7c,
	0xb2, 0x92, 0xd6, 0x29, 0x30, 0x04, 0x0d, 0xd0, 0xb0, 0x93, 0x1d, 0xb6, 0x69, 0x0f, 0xe3, 0xf0,
	0x8a, 0x87, 0x85, 0x34, 0xe4, 0xdf, 0x00, 0x0d, 0xcf, 0xb2, 0x76, 0x43, 0x5e, 0x1b, 0x90, 0xde,
	0xe0, 0xae, 0x56, 0xc6, 0x4f, 0xac, 0x14, 0xaa, 0x54, 0xd2, 0x78, 0xf2, 0x3f, 0x28, 0x9d, 0x86,
	0x66, 0x07, 0x98, 0x3e, 0xe0, 0xd6, 0x54, 0x43, 0x65, 0x3c, 0x39, 0x1a, 0xa0, 0x61, 0x7b, 0xdc,
	0x5b, 0x6f, 0xfb, 0xc9, 0xcf, 0xb6, 0x7f, 0x21, 0xc0, 0x69, 0x70, 0x6e, 0xb6, 0x60, 0x0a, 0xb8,
	0x9e, 0xfa, 0x39, 0x7b, 0x32, 0x3e, 0xdb, 0xcb, 0x21, 0x3d, 0xfe, 0x70, 0xe2, 0xa4, 0x99, 0x49,
	0x4b, 0x8e, 0xf7, 0xe9, 0x91, 0xbe, 0x04, 0x38, 0x7e, 0x5c, 0xef, 0x28, 0xda, 0xec, 0x28, 0xfa,
	0xdd, 0x51, 0xf4, 0x59, 0xd3, 0x64, 0x53, 0xd3, 0xe4, 0xbb, 0xa6, 0xc9, 0xdb, 0x6d, 0xa1, 0xfc,
	0xbc, 0xca, 0x99, 0x00, 0x1d, 0xc6, 0xb7, 0x95, 0x9b, 0x73, 0x03, 0xf9, 0x52, 0xde, 0x85, 0x0e,
	0xdf, 0x63, 0x95, 0xfe, 0xa3, 0x94, 0x2e, 0x6f, 0x85, 0x6e, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x50, 0xcb, 0xab, 0x66, 0x01, 0x00, 0x00,
}

func (m *BurnMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageSender) > 0 {
		i -= len(m.MessageSender)
		copy(dAtA[i:], m.MessageSender)
		i = encodeVarintBurnMessage(dAtA, i, uint64(len(m.MessageSender)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBurnMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MintRecipient) > 0 {
		i -= len(m.MintRecipient)
		copy(dAtA[i:], m.MintRecipient)
		i = encodeVarintBurnMessage(dAtA, i, uint64(len(m.MintRecipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BurnToken) > 0 {
		i -= len(m.BurnToken)
		copy(dAtA[i:], m.BurnToken)
		i = encodeVarintBurnMessage(dAtA, i, uint64(len(m.BurnToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintBurnMessage(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBurnMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovBurnMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BurnMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovBurnMessage(uint64(m.Version))
	}
	l = len(m.BurnToken)
	if l > 0 {
		n += 1 + l + sovBurnMessage(uint64(l))
	}
	l = len(m.MintRecipient)
	if l > 0 {
		n += 1 + l + sovBurnMessage(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovBurnMessage(uint64(l))
	l = len(m.MessageSender)
	if l > 0 {
		n += 1 + l + sovBurnMessage(uint64(l))
	}
	return n
}

func sovBurnMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBurnMessage(x uint64) (n int) {
	return sovBurnMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BurnMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBurnMessage
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
			return fmt.Errorf("proto: BurnMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBurnMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBurnMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnToken = append(m.BurnToken[:0], dAtA[iNdEx:postIndex]...)
			if m.BurnToken == nil {
				m.BurnToken = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRecipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBurnMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBurnMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRecipient = append(m.MintRecipient[:0], dAtA[iNdEx:postIndex]...)
			if m.MintRecipient == nil {
				m.MintRecipient = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnMessage
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
				return ErrInvalidLengthBurnMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBurnMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageSender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBurnMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBurnMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageSender = append(m.MessageSender[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageSender == nil {
				m.MessageSender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBurnMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBurnMessage
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
func skipBurnMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBurnMessage
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
					return 0, ErrIntOverflowBurnMessage
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
					return 0, ErrIntOverflowBurnMessage
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
				return 0, ErrInvalidLengthBurnMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBurnMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBurnMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBurnMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBurnMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBurnMessage = fmt.Errorf("proto: unexpected end of group")
)
