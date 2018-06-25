// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/contester/runlib/contester_proto/Blobs.proto

/*
	Package contester_proto is a generated protocol buffer package.

	It is generated from these files:
		github.com/contester/runlib/contester_proto/Blobs.proto
		github.com/contester/runlib/contester_proto/Contester.proto
		github.com/contester/runlib/contester_proto/Execution.proto
		github.com/contester/runlib/contester_proto/Local.proto

	It has these top-level messages:
		Blob
		Module
		FileBlob
		Compilation
		RedirectParameters
		ExecutionResultFlags
		ExecutionResultTime
		LocalEnvironment
		LocalExecutionParameters
		LocalExecuteConnected
		LocalExecutionResult
		LocalExecuteConnectedResult
		LocalExecution
		BinaryTypeRequest
		BinaryTypeResponse
		ClearSandboxRequest
		IdentifyRequest
		SandboxLocations
		IdentifyResponse
		FileStat
		StatRequest
		FileStats
		GetRequest
		EmptyMessage
		CopyOperation
		CopyOperations
		NamePair
		RepeatedNamePairEntries
		RepeatedStringEntries
*/
package contester_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

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

type Blob_CompressionInfo_CompressionType int32

const (
	Blob_CompressionInfo_METHOD_NONE Blob_CompressionInfo_CompressionType = 0
	Blob_CompressionInfo_METHOD_ZLIB Blob_CompressionInfo_CompressionType = 1
)

var Blob_CompressionInfo_CompressionType_name = map[int32]string{
	0: "METHOD_NONE",
	1: "METHOD_ZLIB",
}
var Blob_CompressionInfo_CompressionType_value = map[string]int32{
	"METHOD_NONE": 0,
	"METHOD_ZLIB": 1,
}

func (x Blob_CompressionInfo_CompressionType) String() string {
	return proto.EnumName(Blob_CompressionInfo_CompressionType_name, int32(x))
}
func (Blob_CompressionInfo_CompressionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorBlobs, []int{0, 0, 0}
}

type Blob struct {
	Data        []byte                `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Compression *Blob_CompressionInfo `protobuf:"bytes,2,opt,name=compression" json:"compression,omitempty"`
	Sha1        []byte                `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptorBlobs, []int{0} }

func (m *Blob) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Blob) GetCompression() *Blob_CompressionInfo {
	if m != nil {
		return m.Compression
	}
	return nil
}

func (m *Blob) GetSha1() []byte {
	if m != nil {
		return m.Sha1
	}
	return nil
}

type Blob_CompressionInfo struct {
	Method       Blob_CompressionInfo_CompressionType `protobuf:"varint,1,opt,name=method,proto3,enum=contester.proto.Blob_CompressionInfo_CompressionType" json:"method,omitempty"`
	OriginalSize uint32                               `protobuf:"varint,2,opt,name=original_size,json=originalSize,proto3" json:"original_size,omitempty"`
}

func (m *Blob_CompressionInfo) Reset()                    { *m = Blob_CompressionInfo{} }
func (*Blob_CompressionInfo) ProtoMessage()               {}
func (*Blob_CompressionInfo) Descriptor() ([]byte, []int) { return fileDescriptorBlobs, []int{0, 0} }

func (m *Blob_CompressionInfo) GetMethod() Blob_CompressionInfo_CompressionType {
	if m != nil {
		return m.Method
	}
	return Blob_CompressionInfo_METHOD_NONE
}

func (m *Blob_CompressionInfo) GetOriginalSize() uint32 {
	if m != nil {
		return m.OriginalSize
	}
	return 0
}

type Module struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Data *Blob  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Module) Reset()                    { *m = Module{} }
func (*Module) ProtoMessage()               {}
func (*Module) Descriptor() ([]byte, []int) { return fileDescriptorBlobs, []int{1} }

func (m *Module) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Module) GetData() *Blob {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Module) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type FileBlob struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data *Blob  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *FileBlob) Reset()                    { *m = FileBlob{} }
func (*FileBlob) ProtoMessage()               {}
func (*FileBlob) Descriptor() ([]byte, []int) { return fileDescriptorBlobs, []int{2} }

func (m *FileBlob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileBlob) GetData() *Blob {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Blob)(nil), "contester.proto.Blob")
	proto.RegisterType((*Blob_CompressionInfo)(nil), "contester.proto.Blob.CompressionInfo")
	proto.RegisterType((*Module)(nil), "contester.proto.Module")
	proto.RegisterType((*FileBlob)(nil), "contester.proto.FileBlob")
	proto.RegisterEnum("contester.proto.Blob_CompressionInfo_CompressionType", Blob_CompressionInfo_CompressionType_name, Blob_CompressionInfo_CompressionType_value)
}
func (this *Blob) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Blob)
	if !ok {
		that2, ok := that.(Blob)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !this.Compression.Equal(that1.Compression) {
		return false
	}
	if !bytes.Equal(this.Sha1, that1.Sha1) {
		return false
	}
	return true
}
func (this *Blob_CompressionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Blob_CompressionInfo)
	if !ok {
		that2, ok := that.(Blob_CompressionInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if this.OriginalSize != that1.OriginalSize {
		return false
	}
	return true
}
func (this *Module) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Module)
	if !ok {
		that2, ok := that.(Module)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *FileBlob) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileBlob)
	if !ok {
		that2, ok := that.(FileBlob)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	return true
}
func (m *Blob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blob) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.Compression != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(m.Compression.Size()))
		n1, err := m.Compression.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Sha1) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(len(m.Sha1)))
		i += copy(dAtA[i:], m.Sha1)
	}
	return i, nil
}

func (m *Blob_CompressionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blob_CompressionInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Method != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(m.Method))
	}
	if m.OriginalSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(m.OriginalSize))
	}
	return i, nil
}

func (m *Module) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Module) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Data != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *FileBlob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileBlob) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Data != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBlobs(dAtA, i, uint64(m.Data.Size()))
		n3, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintBlobs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Blob) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovBlobs(uint64(l))
	}
	if m.Compression != nil {
		l = m.Compression.Size()
		n += 1 + l + sovBlobs(uint64(l))
	}
	l = len(m.Sha1)
	if l > 0 {
		n += 1 + l + sovBlobs(uint64(l))
	}
	return n
}

func (m *Blob_CompressionInfo) Size() (n int) {
	var l int
	_ = l
	if m.Method != 0 {
		n += 1 + sovBlobs(uint64(m.Method))
	}
	if m.OriginalSize != 0 {
		n += 1 + sovBlobs(uint64(m.OriginalSize))
	}
	return n
}

func (m *Module) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovBlobs(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBlobs(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBlobs(uint64(l))
	}
	return n
}

func (m *FileBlob) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBlobs(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBlobs(uint64(l))
	}
	return n
}

func sovBlobs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlobs(x uint64) (n int) {
	return sovBlobs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Blob) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Blob{`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`Compression:` + strings.Replace(fmt.Sprintf("%v", this.Compression), "Blob_CompressionInfo", "Blob_CompressionInfo", 1) + `,`,
		`Sha1:` + fmt.Sprintf("%v", this.Sha1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Blob_CompressionInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Blob_CompressionInfo{`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`OriginalSize:` + fmt.Sprintf("%v", this.OriginalSize) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Module) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Module{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Data:` + strings.Replace(fmt.Sprintf("%v", this.Data), "Blob", "Blob", 1) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FileBlob) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FileBlob{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Data:` + strings.Replace(fmt.Sprintf("%v", this.Data), "Blob", "Blob", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBlobs(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Blob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlobs
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
			return fmt.Errorf("proto: Blob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compression == nil {
				m.Compression = &Blob_CompressionInfo{}
			}
			if err := m.Compression.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha1 = append(m.Sha1[:0], dAtA[iNdEx:postIndex]...)
			if m.Sha1 == nil {
				m.Sha1 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlobs
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
func (m *Blob_CompressionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlobs
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
			return fmt.Errorf("proto: CompressionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompressionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			m.Method = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Method |= (Blob_CompressionInfo_CompressionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSize", wireType)
			}
			m.OriginalSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OriginalSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlobs
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
func (m *Module) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlobs
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
			return fmt.Errorf("proto: Module: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Module: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &Blob{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlobs
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
func (m *FileBlob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlobs
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
			return fmt.Errorf("proto: FileBlob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileBlob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlobs
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
				return ErrInvalidLengthBlobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &Blob{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlobs
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
func skipBlobs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlobs
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
					return 0, ErrIntOverflowBlobs
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
					return 0, ErrIntOverflowBlobs
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
				return 0, ErrInvalidLengthBlobs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlobs
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
				next, err := skipBlobs(dAtA[start:])
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
	ErrInvalidLengthBlobs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlobs   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/contester/runlib/contester_proto/Blobs.proto", fileDescriptorBlobs)
}

var fileDescriptorBlobs = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x2f, 0x2a, 0xcd, 0xcb, 0xc9, 0x4c, 0x42, 0x08, 0xc4, 0x17, 0x14, 0xe5, 0x97, 0xe4, 0xeb,
	0x3b, 0xe5, 0xe4, 0x27, 0x15, 0xeb, 0x81, 0xd9, 0x42, 0xfc, 0x70, 0x49, 0x88, 0x80, 0x94, 0x2e,
	0x92, 0x49, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0x62, 0x00,
	0x88, 0x05, 0x51, 0xae, 0xb4, 0x85, 0x89, 0x8b, 0x05, 0x64, 0x9e, 0x90, 0x10, 0x17, 0x4b, 0x4a,
	0x62, 0x49, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0x2d, 0xe4, 0xce, 0xc5, 0x9d,
	0x9c, 0x9f, 0x5b, 0x50, 0x94, 0x5a, 0x5c, 0x9c, 0x99, 0x9f, 0x27, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x6d, 0xa4, 0xaa, 0x87, 0x66, 0xa5, 0x1e, 0x48, 0xbf, 0x9e, 0x33, 0x42, 0xa1, 0x67, 0x5e, 0x5a,
	0x7e, 0x10, 0xb2, 0x4e, 0x90, 0xe1, 0xc5, 0x19, 0x89, 0x86, 0x12, 0xcc, 0x10, 0xc3, 0x41, 0x6c,
	0xa9, 0x5d, 0x8c, 0x5c, 0xfc, 0x68, 0x9a, 0x84, 0x7c, 0xb9, 0xd8, 0x72, 0x53, 0x4b, 0x32, 0xf2,
	0x53, 0xc0, 0xce, 0xe0, 0x33, 0x32, 0x25, 0xca, 0x2e, 0x64, 0x7e, 0x48, 0x65, 0x41, 0x6a, 0x10,
	0xd4, 0x10, 0x21, 0x65, 0x2e, 0xde, 0xfc, 0xa2, 0xcc, 0xf4, 0xcc, 0xbc, 0xc4, 0x9c, 0xf8, 0xe2,
	0xcc, 0xaa, 0x54, 0xb0, 0x0f, 0x78, 0x83, 0x78, 0x60, 0x82, 0xc1, 0x99, 0x55, 0xa9, 0x4a, 0xc6,
	0x28, 0xce, 0x00, 0xe9, 0x17, 0xe2, 0xe7, 0xe2, 0xf6, 0x75, 0x0d, 0xf1, 0xf0, 0x77, 0x89, 0xf7,
	0xf3, 0xf7, 0x73, 0x15, 0x60, 0x40, 0x12, 0x88, 0xf2, 0xf1, 0x74, 0x12, 0x60, 0x54, 0x8a, 0xe6,
	0x62, 0xf3, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x05, 0x79, 0xad, 0xa4, 0xb2, 0x20, 0x15, 0xec, 0x60,
	0xce, 0x20, 0x30, 0x5b, 0x48, 0x13, 0x1a, 0x96, 0x90, 0x00, 0x13, 0xc5, 0xea, 0x09, 0x68, 0x10,
	0x0b, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x82, 0x43, 0x86, 0x33, 0x08, 0xcc, 0x56, 0xf2, 0xe4,
	0xe2, 0x70, 0xcb, 0xcc, 0x49, 0x85, 0x45, 0x0b, 0x58, 0x9e, 0x11, 0x21, 0x4f, 0x82, 0xf1, 0x4e,
	0x56, 0x37, 0x1e, 0xca, 0x31, 0x34, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x81, 0x4b,
	0x26, 0xbf, 0x28, 0x5d, 0xaf, 0xb8, 0x24, 0x33, 0x2f, 0xbd, 0x28, 0xb1, 0x12, 0xdd, 0x98, 0x24,
	0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xfa, 0x08, 0x44, 0x9c, 0x02, 0x00,
	0x00,
}
