// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: panacea/aol/v2/record.proto

package types

import (
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

// Record defines a record type.
type Record struct {
	Key           []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	NanoTimestamp int64  `protobuf:"varint,3,opt,name=nano_timestamp,json=nanoTimestamp,proto3" json:"nano_timestamp,omitempty"`
	WriterAddress string `protobuf:"bytes,4,opt,name=writer_address,json=writerAddress,proto3" json:"writer_address,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a5bacdb910a2ff, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetNanoTimestamp() int64 {
	if m != nil {
		return m.NanoTimestamp
	}
	return 0
}

func (m *Record) GetWriterAddress() string {
	if m != nil {
		return m.WriterAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "panacea.aol.v2.Record")
}

func init() { proto.RegisterFile("panacea/aol/v2/record.proto", fileDescriptor_e7a5bacdb910a2ff) }

var fileDescriptor_e7a5bacdb910a2ff = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xb3, 0x46, 0x0f, 0x5c, 0xbc, 0x43, 0xc2, 0x15, 0x41, 0x61, 0x09, 0x82, 0x90, 0xc6,
	0x2c, 0x9c, 0x4f, 0xa0, 0xad, 0x8d, 0x04, 0x2b, 0x9b, 0x63, 0xb2, 0x19, 0x62, 0x30, 0xc9, 0x84,
	0xcd, 0x5e, 0xf4, 0x1a, 0x9f, 0xc1, 0xc7, 0xb2, 0xbc, 0xd2, 0x52, 0x92, 0x17, 0x91, 0xdd, 0xe4,
	0xba, 0x7f, 0xbe, 0xff, 0x9f, 0xe2, 0xe3, 0xd7, 0x2d, 0x34, 0xa0, 0x10, 0x24, 0x50, 0x25, 0xfb,
	0x8d, 0xd4, 0xa8, 0x48, 0xe7, 0x49, 0xab, 0xc9, 0x50, 0xb0, 0x9a, 0xcb, 0x04, 0xa8, 0x4a, 0xfa,
	0xcd, 0xd5, 0xba, 0xa0, 0x82, 0x5c, 0x25, 0x6d, 0x9a, 0x56, 0x37, 0x5f, 0x7c, 0x91, 0xba, 0xaf,
	0xe0, 0x92, 0xfb, 0xef, 0xb8, 0x0f, 0x59, 0xc4, 0xe2, 0x8b, 0xd4, 0xc6, 0x60, 0xcd, 0xcf, 0x7a,
	0xa8, 0x76, 0x18, 0x9e, 0x38, 0x36, 0x1d, 0xc1, 0x2d, 0x5f, 0x35, 0xd0, 0xd0, 0xd6, 0x94, 0x35,
	0x76, 0x06, 0xea, 0x36, 0xf4, 0x23, 0x16, 0xfb, 0xe9, 0xd2, 0xd2, 0x97, 0x23, 0xb4, 0xb3, 0x0f,
	0x5d, 0x1a, 0xd4, 0x5b, 0xc8, 0x73, 0x8d, 0x5d, 0x17, 0x9e, 0x46, 0x2c, 0x3e, 0x4f, 0x97, 0x13,
	0x7d, 0x98, 0xe0, 0xe3, 0xd3, 0xcf, 0x20, 0xd8, 0x61, 0x10, 0xec, 0x6f, 0x10, 0xec, 0x7b, 0x14,
	0xde, 0x61, 0x14, 0xde, 0xef, 0x28, 0xbc, 0x67, 0xf6, 0x2a, 0x8b, 0xd2, 0xbc, 0xed, 0xb2, 0x44,
	0x51, 0x2d, 0x6b, 0xcc, 0xcb, 0xac, 0x22, 0x25, 0x67, 0xab, 0x3b, 0x45, 0x1a, 0xad, 0xf3, 0xa7,
	0x93, 0x37, 0xfb, 0x16, 0xbb, 0x6c, 0xe1, 0x9c, 0xee, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0x82, 0x56, 0x05, 0x18, 0x01, 0x00, 0x00,
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WriterAddress) > 0 {
		i -= len(m.WriterAddress)
		copy(dAtA[i:], m.WriterAddress)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.WriterAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.NanoTimestamp != 0 {
		i = encodeVarintRecord(dAtA, i, uint64(m.NanoTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRecord(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	if m.NanoTimestamp != 0 {
		n += 1 + sovRecord(uint64(m.NanoTimestamp))
	}
	l = len(m.WriterAddress)
	if l > 0 {
		n += 1 + l + sovRecord(uint64(l))
	}
	return n
}

func sovRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecord(x uint64) (n int) {
	return sovRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecord
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NanoTimestamp", wireType)
			}
			m.NanoTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NanoTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WriterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecord
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
				return ErrInvalidLengthRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WriterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecord
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
func skipRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
					return 0, ErrIntOverflowRecord
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
				return 0, ErrInvalidLengthRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecord = fmt.Errorf("proto: unexpected end of group")
)
