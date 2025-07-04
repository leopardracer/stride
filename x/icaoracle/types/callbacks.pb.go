// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/icaoracle/callbacks.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Callback data for instantiating an oracle
type InstantiateOracleCallback struct {
	OracleChainId string `protobuf:"bytes,1,opt,name=oracle_chain_id,json=oracleChainId,proto3" json:"oracle_chain_id,omitempty"`
}

func (m *InstantiateOracleCallback) Reset()         { *m = InstantiateOracleCallback{} }
func (m *InstantiateOracleCallback) String() string { return proto.CompactTextString(m) }
func (*InstantiateOracleCallback) ProtoMessage()    {}
func (*InstantiateOracleCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b4c39df2554f0a2, []int{0}
}
func (m *InstantiateOracleCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstantiateOracleCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstantiateOracleCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstantiateOracleCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstantiateOracleCallback.Merge(m, src)
}
func (m *InstantiateOracleCallback) XXX_Size() int {
	return m.Size()
}
func (m *InstantiateOracleCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_InstantiateOracleCallback.DiscardUnknown(m)
}

var xxx_messageInfo_InstantiateOracleCallback proto.InternalMessageInfo

func (m *InstantiateOracleCallback) GetOracleChainId() string {
	if m != nil {
		return m.OracleChainId
	}
	return ""
}

// Callback data for updating a value in the oracle
type UpdateOracleCallback struct {
	OracleChainId string  `protobuf:"bytes,1,opt,name=oracle_chain_id,json=oracleChainId,proto3" json:"oracle_chain_id,omitempty"`
	Metric        *Metric `protobuf:"bytes,2,opt,name=metric,proto3" json:"metric,omitempty"`
}

func (m *UpdateOracleCallback) Reset()         { *m = UpdateOracleCallback{} }
func (m *UpdateOracleCallback) String() string { return proto.CompactTextString(m) }
func (*UpdateOracleCallback) ProtoMessage()    {}
func (*UpdateOracleCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b4c39df2554f0a2, []int{1}
}
func (m *UpdateOracleCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateOracleCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateOracleCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateOracleCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOracleCallback.Merge(m, src)
}
func (m *UpdateOracleCallback) XXX_Size() int {
	return m.Size()
}
func (m *UpdateOracleCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOracleCallback.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOracleCallback proto.InternalMessageInfo

func (m *UpdateOracleCallback) GetOracleChainId() string {
	if m != nil {
		return m.OracleChainId
	}
	return ""
}

func (m *UpdateOracleCallback) GetMetric() *Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func init() {
	proto.RegisterType((*InstantiateOracleCallback)(nil), "stride.icaoracle.InstantiateOracleCallback")
	proto.RegisterType((*UpdateOracleCallback)(nil), "stride.icaoracle.UpdateOracleCallback")
}

func init() { proto.RegisterFile("stride/icaoracle/callbacks.proto", fileDescriptor_7b4c39df2554f0a2) }

var fileDescriptor_7b4c39df2554f0a2 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0xcf, 0x4c, 0x4e, 0xcc, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x4f, 0x4e, 0xcc,
	0xc9, 0x49, 0x4a, 0x4c, 0xce, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xa8,
	0xd0, 0x83, 0xab, 0x90, 0xc2, 0xd4, 0x03, 0x67, 0x41, 0xf4, 0x28, 0x39, 0x73, 0x49, 0x7a, 0xe6,
	0x15, 0x97, 0x24, 0xe6, 0x95, 0x64, 0x26, 0x96, 0xa4, 0xfa, 0x83, 0xa5, 0x9c, 0xa1, 0xe6, 0x0a,
	0xa9, 0x71, 0xf1, 0x43, 0x14, 0xc7, 0x27, 0x67, 0x24, 0x66, 0xe6, 0xc5, 0x67, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0x42, 0x84, 0x9d, 0x41, 0xa2, 0x9e, 0x29, 0x4a, 0x05, 0x5c,
	0x22, 0xa1, 0x05, 0x29, 0x64, 0xeb, 0x17, 0x32, 0xe0, 0x62, 0xcb, 0x4d, 0x2d, 0x29, 0xca, 0x4c,
	0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd0, 0x43, 0xf7, 0x89, 0x9e, 0x2f, 0x58, 0x3e,
	0x08, 0xaa, 0xce, 0xc9, 0xf7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x83, 0xc1, 0xa6, 0xe8, 0xfa,
	0x24, 0x26, 0x15, 0xeb, 0x43, 0x43, 0xa2, 0xcc, 0xc8, 0x5c, 0xbf, 0x02, 0x29, 0x3c, 0x4a, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x81, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xba,
	0xca, 0x90, 0x64, 0x01, 0x00, 0x00,
}

func (m *InstantiateOracleCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstantiateOracleCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstantiateOracleCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OracleChainId) > 0 {
		i -= len(m.OracleChainId)
		copy(dAtA[i:], m.OracleChainId)
		i = encodeVarintCallbacks(dAtA, i, uint64(len(m.OracleChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateOracleCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateOracleCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateOracleCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metric != nil {
		{
			size, err := m.Metric.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallbacks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.OracleChainId) > 0 {
		i -= len(m.OracleChainId)
		copy(dAtA[i:], m.OracleChainId)
		i = encodeVarintCallbacks(dAtA, i, uint64(len(m.OracleChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCallbacks(dAtA []byte, offset int, v uint64) int {
	offset -= sovCallbacks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InstantiateOracleCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OracleChainId)
	if l > 0 {
		n += 1 + l + sovCallbacks(uint64(l))
	}
	return n
}

func (m *UpdateOracleCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OracleChainId)
	if l > 0 {
		n += 1 + l + sovCallbacks(uint64(l))
	}
	if m.Metric != nil {
		l = m.Metric.Size()
		n += 1 + l + sovCallbacks(uint64(l))
	}
	return n
}

func sovCallbacks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallbacks(x uint64) (n int) {
	return sovCallbacks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InstantiateOracleCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbacks
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
			return fmt.Errorf("proto: InstantiateOracleCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstantiateOracleCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallbacks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbacks
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
func (m *UpdateOracleCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbacks
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
			return fmt.Errorf("proto: UpdateOracleCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateOracleCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metric == nil {
				m.Metric = &Metric{}
			}
			if err := m.Metric.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallbacks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbacks
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
func skipCallbacks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
				return 0, ErrInvalidLengthCallbacks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCallbacks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCallbacks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCallbacks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCallbacks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCallbacks = fmt.Errorf("proto: unexpected end of group")
)
