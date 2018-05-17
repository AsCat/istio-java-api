// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/list/config/config.proto

/*
	Package config is a generated protocol buffer package.

	The `list` adapter makes it possible to perform simple whitelist or blacklist
	checks. You can configure the adapter with the list to check, or you can point
	it to a URL from where the list should be fetched. Lists can be simple strings,
	IP addresses, or regex patterns.

	It is generated from these files:
		mixer/adapter/list/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import strconv "strconv"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Determines the type of list that the adapter is consulting.
type Params_ListEntryType int32

const (
	// List entries are treated as plain strings.
	STRINGS Params_ListEntryType = 0
	// List entries are treated as case-insensitive strings.
	CASE_INSENSITIVE_STRINGS Params_ListEntryType = 1
	// List entries are treated as IP addresses and ranges.
	IP_ADDRESSES Params_ListEntryType = 2
	// List entries are treated as re2 regexp. See [here](https://github.com/google/re2/wiki/Syntax) for the supported syntax.
	REGEX Params_ListEntryType = 3
)

var Params_ListEntryType_name = map[int32]string{
	0: "STRINGS",
	1: "CASE_INSENSITIVE_STRINGS",
	2: "IP_ADDRESSES",
	3: "REGEX",
}
var Params_ListEntryType_value = map[string]int32{
	"STRINGS":                  0,
	"CASE_INSENSITIVE_STRINGS": 1,
	"IP_ADDRESSES":             2,
	"REGEX":                    3,
}

func (Params_ListEntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

// Configuration format for the `list` adapter.
type Params struct {
	// Where to find the list to check against. This may be ommited for a completely local list.
	ProviderUrl string `protobuf:"bytes,1,opt,name=provider_url,json=providerUrl,proto3" json:"provider_url,omitempty"`
	// Determines how often the provider is polled for
	// an updated list
	RefreshInterval time.Duration `protobuf:"bytes,2,opt,name=refresh_interval,json=refreshInterval,stdduration" json:"refresh_interval"`
	// Indicates how long to keep a list before discarding it.
	// Typically, the TTL value should be set to noticeably longer (> 2x) than the
	// refresh interval to ensure continued operation in the face of transient
	// server outages.
	Ttl time.Duration `protobuf:"bytes,3,opt,name=ttl,stdduration" json:"ttl"`
	// Indicates the amount of time a caller of this adapter can cache an answer
	// before it should ask the adapter again.
	CachingInterval time.Duration `protobuf:"bytes,4,opt,name=caching_interval,json=cachingInterval,stdduration" json:"caching_interval"`
	// Indicates the number of times a caller of this adapter can use a cached answer
	// before it should ask the adapter again.
	CachingUseCount int32 `protobuf:"varint,5,opt,name=caching_use_count,json=cachingUseCount,proto3" json:"caching_use_count,omitempty"`
	// List entries that are consulted first, before the list from the server
	Overrides []string `protobuf:"bytes,6,rep,name=overrides" json:"overrides,omitempty"`
	// Determines the kind of list entry and overrides.
	EntryType Params_ListEntryType `protobuf:"varint,7,opt,name=entry_type,json=entryType,proto3,enum=adapter.list.config.Params_ListEntryType" json:"entry_type,omitempty"`
	// Whether the list operates as a blacklist or a whitelist.
	Blacklist bool `protobuf:"varint,8,opt,name=blacklist,proto3" json:"blacklist,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.list.config.Params")
	proto.RegisterEnum("adapter.list.config.Params_ListEntryType", Params_ListEntryType_name, Params_ListEntryType_value)
}
func (x Params_ListEntryType) String() string {
	s, ok := Params_ListEntryType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProviderUrl) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ProviderUrl)))
		i += copy(dAtA[i:], m.ProviderUrl)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.RefreshInterval)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RefreshInterval, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.Ttl)))
	n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Ttl, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.CachingInterval)))
	n3, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.CachingInterval, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.CachingUseCount != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.CachingUseCount))
	}
	if len(m.Overrides) > 0 {
		for _, s := range m.Overrides {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.EntryType != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.EntryType))
	}
	if m.Blacklist {
		dAtA[i] = 0x40
		i++
		if m.Blacklist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProviderUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RefreshInterval)
	n += 1 + l + sovConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Ttl)
	n += 1 + l + sovConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.CachingInterval)
	n += 1 + l + sovConfig(uint64(l))
	if m.CachingUseCount != 0 {
		n += 1 + sovConfig(uint64(m.CachingUseCount))
	}
	if len(m.Overrides) > 0 {
		for _, s := range m.Overrides {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if m.EntryType != 0 {
		n += 1 + sovConfig(uint64(m.EntryType))
	}
	if m.Blacklist {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`ProviderUrl:` + fmt.Sprintf("%v", this.ProviderUrl) + `,`,
		`RefreshInterval:` + strings.Replace(strings.Replace(this.RefreshInterval.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`Ttl:` + strings.Replace(strings.Replace(this.Ttl.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`CachingInterval:` + strings.Replace(strings.Replace(this.CachingInterval.String(), "Duration", "google_protobuf.Duration", 1), `&`, ``, 1) + `,`,
		`CachingUseCount:` + fmt.Sprintf("%v", this.CachingUseCount) + `,`,
		`Overrides:` + fmt.Sprintf("%v", this.Overrides) + `,`,
		`EntryType:` + fmt.Sprintf("%v", this.EntryType) + `,`,
		`Blacklist:` + fmt.Sprintf("%v", this.Blacklist) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RefreshInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Ttl, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CachingInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.CachingInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CachingUseCount", wireType)
			}
			m.CachingUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CachingUseCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Overrides = append(m.Overrides, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryType", wireType)
			}
			m.EntryType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntryType |= (Params_ListEntryType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Blacklist = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowConfig
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/list/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0xef, 0x9a, 0x36, 0x8d, 0x2f, 0x05, 0xcc, 0xc1, 0x60, 0xaa, 0xea, 0x6a, 0x3a, 0x20,
	0xc3, 0x70, 0x96, 0x8a, 0x90, 0x58, 0xdb, 0xc6, 0x2a, 0x96, 0x50, 0x54, 0xd9, 0x29, 0x20, 0x16,
	0xcb, 0x71, 0x2e, 0xee, 0x09, 0xd7, 0x67, 0x9d, 0xcf, 0x11, 0xd9, 0x78, 0x01, 0x24, 0x46, 0x1e,
	0x81, 0x47, 0xc9, 0xd8, 0x91, 0x09, 0x88, 0x59, 0x18, 0xfb, 0x08, 0xc8, 0xb1, 0xdd, 0x08, 0x89,
	0x01, 0x26, 0xff, 0xf5, 0xdd, 0xf7, 0xbb, 0xef, 0xf3, 0xff, 0xd0, 0xa3, 0x4b, 0xfe, 0x9e, 0x49,
	0x3b, 0x9c, 0x84, 0x99, 0x62, 0xd2, 0x4e, 0x78, 0xae, 0xec, 0x48, 0xa4, 0x53, 0x1e, 0x37, 0x1f,
	0x9a, 0x49, 0xa1, 0x04, 0xbe, 0xd7, 0x38, 0x68, 0xe5, 0xa0, 0xf5, 0xd1, 0x2e, 0x89, 0x85, 0x88,
	0x13, 0x66, 0xaf, 0x2c, 0xe3, 0x62, 0x6a, 0x4f, 0x0a, 0x19, 0x2a, 0x2e, 0xd2, 0x1a, 0xda, 0xbd,
	0x1f, 0x8b, 0x58, 0xac, 0x46, 0xbb, 0x9a, 0x6a, 0xf5, 0xe0, 0xe3, 0x26, 0xea, 0x9e, 0x85, 0x32,
	0xbc, 0xcc, 0xf1, 0x43, 0xb4, 0x93, 0x49, 0x31, 0xe3, 0x13, 0x26, 0x83, 0x42, 0x26, 0x06, 0x34,
	0xa1, 0xa5, 0x79, 0xfd, 0x56, 0x3b, 0x97, 0x09, 0x1e, 0x22, 0x5d, 0xb2, 0xa9, 0x64, 0xf9, 0x45,
	0xc0, 0x53, 0xc5, 0xe4, 0x2c, 0x4c, 0x8c, 0x0d, 0x13, 0x5a, 0xfd, 0xc3, 0x07, 0xb4, 0x8e, 0xa7,
	0x6d, 0x3c, 0x1d, 0x34, 0xf1, 0xc7, 0xbd, 0xc5, 0xb7, 0x7d, 0xf0, 0xf9, 0xfb, 0x3e, 0xf4, 0xee,
	0x34, 0xb0, 0xdb, 0xb0, 0xf8, 0x19, 0xea, 0x28, 0x95, 0x18, 0x9d, 0x7f, 0xbf, 0xa2, 0xf2, 0x57,
	0x35, 0xa2, 0x30, 0xba, 0xe0, 0x69, 0xbc, 0xae, 0xb1, 0xf9, 0x1f, 0x35, 0x1a, 0xf8, 0xa6, 0xc6,
	0x13, 0x74, 0xb7, 0xbd, 0xaf, 0xc8, 0x59, 0x10, 0x89, 0x22, 0x55, 0xc6, 0x96, 0x09, 0xad, 0xad,
	0x1b, 0xef, 0x79, 0xce, 0x4e, 0x2a, 0x19, 0xef, 0x21, 0x4d, 0xcc, 0x98, 0x94, 0x7c, 0xc2, 0x72,
	0xa3, 0x6b, 0x76, 0x2c, 0xcd, 0x5b, 0x0b, 0xf8, 0x05, 0x42, 0x2c, 0x55, 0x72, 0x1e, 0xa8, 0x79,
	0xc6, 0x8c, 0x6d, 0x13, 0x5a, 0xb7, 0x0f, 0x1f, 0xd3, 0xbf, 0x3c, 0x17, 0xad, 0x97, 0x4e, 0x5f,
	0xf2, 0x5c, 0x39, 0x15, 0x31, 0x9a, 0x67, 0xcc, 0xd3, 0x58, 0x3b, 0x56, 0x39, 0xe3, 0x24, 0x8c,
	0xde, 0x55, 0x8c, 0xd1, 0x33, 0xa1, 0xd5, 0xf3, 0xd6, 0xc2, 0xc1, 0x6b, 0x74, 0xeb, 0x0f, 0x12,
	0xf7, 0xd1, 0xb6, 0x3f, 0xf2, 0xdc, 0xe1, 0xa9, 0xaf, 0x03, 0xbc, 0x87, 0x8c, 0x93, 0x23, 0xdf,
	0x09, 0xdc, 0xa1, 0xef, 0x0c, 0x7d, 0x77, 0xe4, 0xbe, 0x72, 0x82, 0xf6, 0x14, 0x62, 0x1d, 0xed,
	0xb8, 0x67, 0xc1, 0xd1, 0x60, 0xe0, 0x39, 0xbe, 0xef, 0xf8, 0xfa, 0x06, 0xd6, 0xd0, 0x96, 0xe7,
	0x9c, 0x3a, 0x6f, 0xf4, 0xce, 0xf1, 0xf3, 0xc5, 0x92, 0x80, 0xab, 0x25, 0x01, 0x5f, 0x97, 0x04,
	0x5c, 0x2f, 0x09, 0xf8, 0x50, 0x12, 0xf8, 0xa5, 0x24, 0x60, 0x51, 0x12, 0x78, 0x55, 0x12, 0xf8,
	0xa3, 0x24, 0xf0, 0x57, 0x49, 0xc0, 0x75, 0x49, 0xe0, 0xa7, 0x9f, 0x04, 0xbc, 0xed, 0xd6, 0x3f,
	0x34, 0xee, 0xae, 0x56, 0xfe, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x09, 0x6d, 0xa2,
	0xc5, 0x02, 0x00, 0x00,
}