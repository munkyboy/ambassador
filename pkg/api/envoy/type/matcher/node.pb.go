// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/node.proto

package envoy_type_matcher

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Specifies the way to match a Node.
// The match follows AND semantics.
type NodeMatcher struct {
	// Specifies match criteria on the node id.
	NodeId *StringMatcher `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Specifies match criteria on the node metadata.
	NodeMetadatas        []*StructMatcher `protobuf:"bytes,2,rep,name=node_metadatas,json=nodeMetadatas,proto3" json:"node_metadatas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeMatcher) Reset()         { *m = NodeMatcher{} }
func (m *NodeMatcher) String() string { return proto.CompactTextString(m) }
func (*NodeMatcher) ProtoMessage()    {}
func (*NodeMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20314fb2f725fb2, []int{0}
}
func (m *NodeMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMatcher.Merge(m, src)
}
func (m *NodeMatcher) XXX_Size() int {
	return m.Size()
}
func (m *NodeMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMatcher proto.InternalMessageInfo

func (m *NodeMatcher) GetNodeId() *StringMatcher {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeMatcher) GetNodeMetadatas() []*StructMatcher {
	if m != nil {
		return m.NodeMetadatas
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeMatcher)(nil), "envoy.type.matcher.NodeMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/node.proto", fileDescriptor_c20314fb2f725fb2) }

var fileDescriptor_c20314fb2f725fb2 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x4b, 0xeb, 0x81, 0xa4,
	0xf5, 0xa0, 0xd2, 0x52, 0xf2, 0x58, 0xb4, 0x14, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x43, 0x34, 0xe1,
	0x52, 0x50, 0x9a, 0x5c, 0x02, 0x55, 0x20, 0x5b, 0x9a, 0x52, 0x90, 0xa8, 0x9f, 0x98, 0x97, 0x97,
	0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0xac, 0x5f, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x0c, 0x91,
	0x56, 0x9a, 0xcc, 0xc8, 0xc5, 0xed, 0x97, 0x9f, 0x92, 0xea, 0x0b, 0xd1, 0x2b, 0x64, 0xc5, 0xc5,
	0x0e, 0x72, 0x52, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa2, 0x1e, 0xa6,
	0xb3, 0xf4, 0x82, 0xc1, 0x4e, 0x80, 0xea, 0x09, 0x62, 0x03, 0xe9, 0xf0, 0x4c, 0x11, 0xf2, 0xe0,
	0xe2, 0x03, 0xeb, 0xcd, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x2c, 0x96, 0x60, 0x52, 0x60,
	0xc6, 0x63, 0x44, 0x69, 0x72, 0x09, 0xcc, 0x08, 0x5e, 0x90, 0x46, 0x5f, 0x98, 0x3e, 0x27, 0xe7,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0x57, 0xc3, 0x89,
	0x8b, 0x6c, 0x4c, 0x02, 0x0c, 0x5c, 0x0a, 0x99, 0xf9, 0x10, 0x93, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a,
	0xb1, 0x18, 0xea, 0xc4, 0x09, 0xf2, 0x4a, 0x00, 0xc8, 0x63, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x1f,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xe6, 0x2b, 0x13, 0x77, 0x01, 0x00, 0x00,
}

func (m *NodeMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NodeMetadatas) > 0 {
		for iNdEx := len(m.NodeMetadatas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeMetadatas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.NodeId != nil {
		{
			size, err := m.NodeId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNode(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodeMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeId != nil {
		l = m.NodeId.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	if len(m.NodeMetadatas) > 0 {
		for _, e := range m.NodeMetadatas {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: NodeMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeId == nil {
				m.NodeId = &StringMatcher{}
			}
			if err := m.NodeId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeMetadatas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeMetadatas = append(m.NodeMetadatas, &StructMatcher{})
			if err := m.NodeMetadatas[len(m.NodeMetadatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
