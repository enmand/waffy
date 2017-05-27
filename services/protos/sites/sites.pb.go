// Code generated by protoc-gen-gogo.
// source: protos/sites/sites.proto
// DO NOT EDIT!

/*
	Package sites is a generated protocol buffer package.

	It is generated from these files:
		protos/sites/sites.proto

	It has these top-level messages:
		Site
		Balancer
*/
package sites

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import nodes "protos/nodes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Site represents a Site that should be load balanced, and have Rules applied to it
type Site struct {
	Hostname    string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Alias       []string `protobuf:"bytes,2,rep,name=alias" json:"alias,omitempty"`
	Secure      bool     `protobuf:"varint,5,opt,name=secure,proto3" json:"secure,omitempty"`
	Autoencrypt bool     `protobuf:"varint,6,opt,name=autoencrypt,proto3" json:"autoencrypt,omitempty"`
}

func (m *Site) Reset()                    { *m = Site{} }
func (m *Site) String() string            { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()               {}
func (*Site) Descriptor() ([]byte, []int) { return fileDescriptorSites, []int{0} }

// Balancer represents a Site load balancer
type Balancer struct {
	Proto string        `protobuf:"bytes,3,opt,name=proto,proto3" json:"proto,omitempty"`
	Port  string        `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Sites []*Site       `protobuf:"bytes,1,rep,name=sites" json:"sites,omitempty"`
	Notes []*nodes.Node `protobuf:"bytes,2,rep,name=notes" json:"notes,omitempty"`
}

func (m *Balancer) Reset()                    { *m = Balancer{} }
func (m *Balancer) String() string            { return proto.CompactTextString(m) }
func (*Balancer) ProtoMessage()               {}
func (*Balancer) Descriptor() ([]byte, []int) { return fileDescriptorSites, []int{1} }

func (m *Balancer) GetSites() []*Site {
	if m != nil {
		return m.Sites
	}
	return nil
}

func (m *Balancer) GetNotes() []*nodes.Node {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*Site)(nil), "sites.Site")
	proto.RegisterType((*Balancer)(nil), "sites.Balancer")
}
func (m *Site) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Site) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hostname) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintSites(data, i, uint64(len(m.Hostname)))
		i += copy(data[i:], m.Hostname)
	}
	if len(m.Alias) > 0 {
		for _, s := range m.Alias {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.Secure {
		data[i] = 0x28
		i++
		if m.Secure {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Autoencrypt {
		data[i] = 0x30
		i++
		if m.Autoencrypt {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Balancer) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Balancer) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sites) > 0 {
		for _, msg := range m.Sites {
			data[i] = 0xa
			i++
			i = encodeVarintSites(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Notes) > 0 {
		for _, msg := range m.Notes {
			data[i] = 0x12
			i++
			i = encodeVarintSites(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Proto) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintSites(data, i, uint64(len(m.Proto)))
		i += copy(data[i:], m.Proto)
	}
	if len(m.Port) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintSites(data, i, uint64(len(m.Port)))
		i += copy(data[i:], m.Port)
	}
	return i, nil
}

func encodeFixed64Sites(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Sites(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSites(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Site) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	if len(m.Alias) > 0 {
		for _, s := range m.Alias {
			l = len(s)
			n += 1 + l + sovSites(uint64(l))
		}
	}
	if m.Secure {
		n += 2
	}
	if m.Autoencrypt {
		n += 2
	}
	return n
}

func (m *Balancer) Size() (n int) {
	var l int
	_ = l
	if len(m.Sites) > 0 {
		for _, e := range m.Sites {
			l = e.Size()
			n += 1 + l + sovSites(uint64(l))
		}
	}
	if len(m.Notes) > 0 {
		for _, e := range m.Notes {
			l = e.Size()
			n += 1 + l + sovSites(uint64(l))
		}
	}
	l = len(m.Proto)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovSites(uint64(l))
	}
	return n
}

func sovSites(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSites(x uint64) (n int) {
	return sovSites(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Site) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSites
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Site: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Site: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = append(m.Alias, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secure", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Secure = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Autoencrypt", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Autoencrypt = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSites(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSites
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
func (m *Balancer) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSites
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Balancer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balancer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sites", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sites = append(m.Sites, &Site{})
			if err := m.Sites[len(m.Sites)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notes = append(m.Notes, &nodes.Node{})
			if err := m.Notes[len(m.Notes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proto = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSites
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSites
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSites(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSites
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
func skipSites(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSites
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSites
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSites
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSites
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSites
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSites(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSites = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSites   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/sites/sites.proto", fileDescriptorSites) }

var fileDescriptorSites = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x85, 0xc9, 0xe6, 0x47, 0x61, 0xd2, 0xa0, 0x11, 0x42, 0xd6, 0x16, 0xab, 0xb0, 0xd5, 0x56,
	0x8b, 0x04, 0x37, 0xc8, 0x01, 0x28, 0xcc, 0x09, 0x4c, 0xb0, 0x44, 0xa4, 0x60, 0x47, 0xb6, 0x53,
	0xc0, 0x49, 0x38, 0x12, 0x25, 0x47, 0x40, 0x70, 0x11, 0xec, 0x19, 0x83, 0xb6, 0x98, 0x51, 0xde,
	0xbc, 0x4f, 0xca, 0x67, 0x10, 0x8b, 0xb3, 0xc1, 0xfa, 0x1b, 0x3f, 0x05, 0x9d, 0xf7, 0x91, 0x4e,
	0x58, 0x53, 0xd8, 0xfe, 0x01, 0xc6, 0x3e, 0xe9, 0xbc, 0x19, 0xd8, 0x3b, 0xa8, 0x1e, 0x22, 0x82,
	0x5b, 0x68, 0x9f, 0xad, 0x0f, 0x46, 0xbd, 0x68, 0x51, 0xf4, 0xc5, 0xe1, 0x5c, 0xfe, 0x67, 0xbc,
	0x84, 0x5a, 0xcd, 0x93, 0xf2, 0x62, 0xd3, 0x97, 0xb1, 0xe0, 0x80, 0x57, 0xd0, 0x78, 0x3d, 0xae,
	0x4e, 0x8b, 0x3a, 0xf2, 0xad, 0xcc, 0x09, 0x7b, 0xe8, 0xd4, 0x1a, 0xac, 0x36, 0xa3, 0x7b, 0x5d,
	0x82, 0x68, 0xa8, 0x3c, 0x3d, 0xed, 0xdf, 0xa0, 0x1d, 0xd4, 0xac, 0xcc, 0xa8, 0x1d, 0x5e, 0x03,
	0x2b, 0xc6, 0x9f, 0x96, 0x87, 0xee, 0xb6, 0x3b, 0xb2, 0x7d, 0x72, 0x92, 0xdc, 0x24, 0xc4, 0xd8,
	0x84, 0x6c, 0x32, 0xc2, 0xfe, 0xf7, 0x71, 0x4b, 0x6e, 0x92, 0x21, 0x3d, 0x47, 0x94, 0xa4, 0xce,
	0x01, 0x11, 0xaa, 0xc5, 0xba, 0x20, 0x2a, 0x3a, 0xd2, 0xf7, 0x70, 0xf1, 0xf1, 0xbd, 0x2b, 0x3e,
	0xe3, 0x7c, 0xc5, 0x79, 0xff, 0xd9, 0x9d, 0x3d, 0x36, 0x04, 0xdf, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x84, 0x78, 0x7d, 0x45, 0x01, 0x00, 0x00,
}
