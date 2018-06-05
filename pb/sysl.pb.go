// Code writeerated by protoc-gen-go. DO NOT EDIT.
// source: pb/sysl.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this writeerated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Deltas are diff annotations set for each major element in a model.
// TODO: Is this really needed? Investigate protobuf diff tools.
type Delta int32

const (
	Delta_NO_Delta     Delta = 0
	Delta_DELTA_SAME   Delta = 1
	Delta_DELTA_CHANGE Delta = 2
	Delta_DELTA_ADD    Delta = 3
	Delta_DELTA_REMOVE Delta = 4
)

var Delta_name = map[int32]string{
	0: "NO_Delta",
	1: "DELTA_SAME",
	2: "DELTA_CHANGE",
	3: "DELTA_ADD",
	4: "DELTA_REMOVE",
}
var Delta_value = map[string]int32{
	"NO_Delta":     0,
	"DELTA_SAME":   1,
	"DELTA_CHANGE": 2,
	"DELTA_ADD":    3,
	"DELTA_REMOVE": 4,
}

func (x Delta) String() string {
	return proto.EnumName(Delta_name, int32(x))
}
func (Delta) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{0}
}

type Endpoint_RestParams_Method int32

const (
	Endpoint_RestParams_NO_Method        Endpoint_RestParams_Method = 0
	Endpoint_RestParams_GET              Endpoint_RestParams_Method = 1
	Endpoint_RestParams_PUT              Endpoint_RestParams_Method = 3
	Endpoint_RestParams_POST             Endpoint_RestParams_Method = 4
	Endpoint_RestParams_DELETE           Endpoint_RestParams_Method = 5
	Endpoint_RestParams_PATCH            Endpoint_RestParams_Method = 6
	Endpoint_RestParams_DONOTUSE_OPTIONS Endpoint_RestParams_Method = 7
	Endpoint_RestParams_DONOTUSE_HEAD    Endpoint_RestParams_Method = 2
)

var Endpoint_RestParams_Method_name = map[int32]string{
	0: "NO_Method",
	1: "GET",
	3: "PUT",
	4: "POST",
	5: "DELETE",
	6: "PATCH",
	7: "DONOTUSE_OPTIONS",
	2: "DONOTUSE_HEAD",
}
var Endpoint_RestParams_Method_value = map[string]int32{
	"NO_Method":        0,
	"GET":              1,
	"PUT":              3,
	"POST":             4,
	"DELETE":           5,
	"PATCH":            6,
	"DONOTUSE_OPTIONS": 7,
	"DONOTUSE_HEAD":    2,
}

func (x Endpoint_RestParams_Method) String() string {
	return proto.EnumName(Endpoint_RestParams_Method_name, int32(x))
}
func (Endpoint_RestParams_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{5, 1, 0}
}

type Endpoint_RestParams_QueryParam_Loc int32

const (
	Endpoint_RestParams_QueryParam_NO_In Endpoint_RestParams_QueryParam_Loc = 0
	Endpoint_RestParams_QueryParam_QUERY Endpoint_RestParams_QueryParam_Loc = 1
	Endpoint_RestParams_QueryParam_PATH  Endpoint_RestParams_QueryParam_Loc = 2
	Endpoint_RestParams_QueryParam_BODY  Endpoint_RestParams_QueryParam_Loc = 3
)

var Endpoint_RestParams_QueryParam_Loc_name = map[int32]string{
	0: "NO_In",
	1: "QUERY",
	2: "PATH",
	3: "BODY",
}
var Endpoint_RestParams_QueryParam_Loc_value = map[string]int32{
	"NO_In": 0,
	"QUERY": 1,
	"PATH":  2,
	"BODY":  3,
}

func (x Endpoint_RestParams_QueryParam_Loc) String() string {
	return proto.EnumName(Endpoint_RestParams_QueryParam_Loc_name, int32(x))
}
func (Endpoint_RestParams_QueryParam_Loc) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{5, 1, 0, 0}
}

type Loop_Mode int32

const (
	Loop_NO_Mode Loop_Mode = 0
	Loop_WHILE   Loop_Mode = 1
	Loop_UNTIL   Loop_Mode = 2
)

var Loop_Mode_name = map[int32]string{
	0: "NO_Mode",
	1: "WHILE",
	2: "UNTIL",
}
var Loop_Mode_value = map[string]int32{
	"NO_Mode": 0,
	"WHILE":   1,
	"UNTIL":   2,
}

func (x Loop_Mode) String() string {
	return proto.EnumName(Loop_Mode_name, int32(x))
}
func (Loop_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{11, 0}
}

type Type_Primitive int32

const (
	Type_NO_Primitive Type_Primitive = 0
	Type_EMPTY        Type_Primitive = 1
	Type_ANY          Type_Primitive = 2
	Type_BOOL         Type_Primitive = 3
	Type_INT          Type_Primitive = 4
	Type_FLOAT        Type_Primitive = 5
	Type_DECIMAL      Type_Primitive = 12
	// Unicode string (Python 2 unicode, Python 3 str, and SQL nvarchar).
	Type_STRING Type_Primitive = 6
	// Octet sequence, like Python 3 bytes and SQL varbinary.
	Type_BYTES Type_Primitive = 7
	// = STRING, but only 8-bit, like Python 2 str or SQL varchar.
	Type_STRING_8 Type_Primitive = 8
	Type_DATE     Type_Primitive = 9
	Type_DATETIME Type_Primitive = 10
	Type_XML      Type_Primitive = 11
	Type_UUID     Type_Primitive = 13
)

var Type_Primitive_name = map[int32]string{
	0:  "NO_Primitive",
	1:  "EMPTY",
	2:  "ANY",
	3:  "BOOL",
	4:  "INT",
	5:  "FLOAT",
	12: "DECIMAL",
	6:  "STRING",
	7:  "BYTES",
	8:  "STRING_8",
	9:  "DATE",
	10: "DATETIME",
	11: "XML",
	13: "UUID",
}
var Type_Primitive_value = map[string]int32{
	"NO_Primitive": 0,
	"EMPTY":        1,
	"ANY":          2,
	"BOOL":         3,
	"INT":          4,
	"FLOAT":        5,
	"DECIMAL":      12,
	"STRING":       6,
	"BYTES":        7,
	"STRING_8":     8,
	"DATE":         9,
	"DATETIME":     10,
	"XML":          11,
	"UUID":         13,
}

func (x Type_Primitive) String() string {
	return proto.EnumName(Type_Primitive_name, int32(x))
}
func (Type_Primitive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 0}
}

type Expr_UnExpr_Op int32

const (
	Expr_UnExpr_NO_Op          Expr_UnExpr_Op = 0
	Expr_UnExpr_NEG            Expr_UnExpr_Op = 1
	Expr_UnExpr_POS            Expr_UnExpr_Op = 2
	Expr_UnExpr_NOT            Expr_UnExpr_Op = 3
	Expr_UnExpr_INV            Expr_UnExpr_Op = 4
	Expr_UnExpr_SINGLE         Expr_UnExpr_Op = 5
	Expr_UnExpr_SINGLE_OR_NULL Expr_UnExpr_Op = 6
)

var Expr_UnExpr_Op_name = map[int32]string{
	0: "NO_Op",
	1: "NEG",
	2: "POS",
	3: "NOT",
	4: "INV",
	5: "SINGLE",
	6: "SINGLE_OR_NULL",
}
var Expr_UnExpr_Op_value = map[string]int32{
	"NO_Op":          0,
	"NEG":            1,
	"POS":            2,
	"NOT":            3,
	"INV":            4,
	"SINGLE":         5,
	"SINGLE_OR_NULL": 6,
}

func (x Expr_UnExpr_Op) String() string {
	return proto.EnumName(Expr_UnExpr_Op_name, int32(x))
}
func (Expr_UnExpr_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 6, 0}
}

type Expr_BinExpr_Op int32

const (
	Expr_BinExpr_NO_Op           Expr_BinExpr_Op = 0
	Expr_BinExpr_EQ              Expr_BinExpr_Op = 1
	Expr_BinExpr_NE              Expr_BinExpr_Op = 2
	Expr_BinExpr_LT              Expr_BinExpr_Op = 3
	Expr_BinExpr_LE              Expr_BinExpr_Op = 4
	Expr_BinExpr_GT              Expr_BinExpr_Op = 5
	Expr_BinExpr_GE              Expr_BinExpr_Op = 6
	Expr_BinExpr_IN              Expr_BinExpr_Op = 24
	Expr_BinExpr_CONTAINS        Expr_BinExpr_Op = 25
	Expr_BinExpr_NOT_IN          Expr_BinExpr_Op = 26
	Expr_BinExpr_NOT_CONTAINS    Expr_BinExpr_Op = 27
	Expr_BinExpr_ADD             Expr_BinExpr_Op = 7
	Expr_BinExpr_SUB             Expr_BinExpr_Op = 8
	Expr_BinExpr_MUL             Expr_BinExpr_Op = 9
	Expr_BinExpr_DIV             Expr_BinExpr_Op = 10
	Expr_BinExpr_MOD             Expr_BinExpr_Op = 11
	Expr_BinExpr_POW             Expr_BinExpr_Op = 12
	Expr_BinExpr_AND             Expr_BinExpr_Op = 13
	Expr_BinExpr_OR              Expr_BinExpr_Op = 14
	Expr_BinExpr_BUTNOT          Expr_BinExpr_Op = 21
	Expr_BinExpr_BITAND          Expr_BinExpr_Op = 15
	Expr_BinExpr_BITOR           Expr_BinExpr_Op = 16
	Expr_BinExpr_BITXOR          Expr_BinExpr_Op = 17
	Expr_BinExpr_COALESCE        Expr_BinExpr_Op = 18
	Expr_BinExpr_WHERE           Expr_BinExpr_Op = 19
	Expr_BinExpr_TO_MATCHING     Expr_BinExpr_Op = 20
	Expr_BinExpr_TO_NOT_MATCHING Expr_BinExpr_Op = 23
	Expr_BinExpr_FLATTEN         Expr_BinExpr_Op = 22
)

var Expr_BinExpr_Op_name = map[int32]string{
	0:  "NO_Op",
	1:  "EQ",
	2:  "NE",
	3:  "LT",
	4:  "LE",
	5:  "GT",
	6:  "GE",
	24: "IN",
	25: "CONTAINS",
	26: "NOT_IN",
	27: "NOT_CONTAINS",
	7:  "ADD",
	8:  "SUB",
	9:  "MUL",
	10: "DIV",
	11: "MOD",
	12: "POW",
	13: "AND",
	14: "OR",
	21: "BUTNOT",
	15: "BITAND",
	16: "BITOR",
	17: "BITXOR",
	18: "COALESCE",
	19: "WHERE",
	20: "TO_MATCHING",
	23: "TO_NOT_MATCHING",
	22: "FLATTEN",
}
var Expr_BinExpr_Op_value = map[string]int32{
	"NO_Op":           0,
	"EQ":              1,
	"NE":              2,
	"LT":              3,
	"LE":              4,
	"GT":              5,
	"GE":              6,
	"IN":              24,
	"CONTAINS":        25,
	"NOT_IN":          26,
	"NOT_CONTAINS":    27,
	"ADD":             7,
	"SUB":             8,
	"MUL":             9,
	"DIV":             10,
	"MOD":             11,
	"POW":             12,
	"AND":             13,
	"OR":              14,
	"BUTNOT":          21,
	"BITAND":          15,
	"BITOR":           16,
	"BITXOR":          17,
	"COALESCE":        18,
	"WHERE":           19,
	"TO_MATCHING":     20,
	"TO_NOT_MATCHING": 23,
	"FLATTEN":         22,
}

func (x Expr_BinExpr_Op) String() string {
	return proto.EnumName(Expr_BinExpr_Op_name, int32(x))
}
func (Expr_BinExpr_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 7, 0}
}

type Expr_RelExpr_Op int32

const (
	Expr_RelExpr_NO_Op          Expr_RelExpr_Op = 0
	Expr_RelExpr_MIN            Expr_RelExpr_Op = 1
	Expr_RelExpr_MAX            Expr_RelExpr_Op = 2
	Expr_RelExpr_SUM            Expr_RelExpr_Op = 3
	Expr_RelExpr_AVERAGE        Expr_RelExpr_Op = 4
	Expr_RelExpr_FUTURE_WHERE   Expr_RelExpr_Op = 5
	Expr_RelExpr_FUTURE_FLATTEN Expr_RelExpr_Op = 6
	Expr_RelExpr_RANK           Expr_RelExpr_Op = 7
	Expr_RelExpr_SNAPSHOT       Expr_RelExpr_Op = 8
	Expr_RelExpr_FIRST_BY       Expr_RelExpr_Op = 9
)

var Expr_RelExpr_Op_name = map[int32]string{
	0: "NO_Op",
	1: "MIN",
	2: "MAX",
	3: "SUM",
	4: "AVERAGE",
	5: "FUTURE_WHERE",
	6: "FUTURE_FLATTEN",
	7: "RANK",
	8: "SNAPSHOT",
	9: "FIRST_BY",
}
var Expr_RelExpr_Op_value = map[string]int32{
	"NO_Op":          0,
	"MIN":            1,
	"MAX":            2,
	"SUM":            3,
	"AVERAGE":        4,
	"FUTURE_WHERE":   5,
	"FUTURE_FLATTEN": 6,
	"RANK":           7,
	"SNAPSHOT":       8,
	"FIRST_BY":       9,
}

func (x Expr_RelExpr_Op) String() string {
	return proto.EnumName(Expr_RelExpr_Op_name, int32(x))
}
func (Expr_RelExpr_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 8, 0}
}

type SourceContext struct {
	File                 string                  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Start                *SourceContext_Location `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End                  *SourceContext_Location `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	Delta                Delta                   `protobuf:"varint,6,opt,name=delta,enum=sysl.Delta" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SourceContext) Reset()         { *m = SourceContext{} }
func (m *SourceContext) String() string { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()    {}
func (*SourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{0}
}
func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext.Unmarshal(m, b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext.Marshal(b, m, deterministic)
}
func (dst *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(dst, src)
}
func (m *SourceContext) XXX_Size() int {
	return xxx_messageInfo_SourceContext.Size(m)
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

func (m *SourceContext) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *SourceContext) GetStart() *SourceContext_Location {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SourceContext) GetEnd() *SourceContext_Location {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *SourceContext) GetDelta() Delta {
	if m != nil {
		return m.Delta
	}
	return Delta_NO_Delta
}

type SourceContext_Location struct {
	Line                 int32    `protobuf:"varint,1,opt,name=line" json:"line,omitempty"`
	Col                  int32    `protobuf:"varint,2,opt,name=col" json:"col,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceContext_Location) Reset()         { *m = SourceContext_Location{} }
func (m *SourceContext_Location) String() string { return proto.CompactTextString(m) }
func (*SourceContext_Location) ProtoMessage()    {}
func (*SourceContext_Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{0, 0}
}
func (m *SourceContext_Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext_Location.Unmarshal(m, b)
}
func (m *SourceContext_Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext_Location.Marshal(b, m, deterministic)
}
func (dst *SourceContext_Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext_Location.Merge(dst, src)
}
func (m *SourceContext_Location) XXX_Size() int {
	return xxx_messageInfo_SourceContext_Location.Size(m)
}
func (m *SourceContext_Location) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext_Location.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext_Location proto.InternalMessageInfo

func (m *SourceContext_Location) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *SourceContext_Location) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

type Module struct {
	Apps                 map[string]*Application `protobuf:"bytes,2,rep,name=apps" json:"apps,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Types                map[string]*Type        `protobuf:"bytes,3,rep,name=types" json:"types,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SourceContext        *SourceContext          `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{1}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Module.Unmarshal(m, b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Module.Marshal(b, m, deterministic)
}
func (dst *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(dst, src)
}
func (m *Module) XXX_Size() int {
	return xxx_messageInfo_Module.Size(m)
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetApps() map[string]*Application {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *Module) GetTypes() map[string]*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Module) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

type Attribute struct {
	// Types that are valid to be assigned to Attribute:
	//	*Attribute_S
	//	*Attribute_I
	//	*Attribute_N
	//	*Attribute_A
	Attribute            isAttribute_Attribute `protobuf_oneof:"attribute"`
	SourceContext        *SourceContext        `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{2}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (dst *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(dst, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

type isAttribute_Attribute interface {
	isAttribute_Attribute()
}

type Attribute_S struct {
	S string `protobuf:"bytes,4,opt,name=s,oneof"`
}
type Attribute_I struct {
	I int64 `protobuf:"varint,5,opt,name=i,oneof"`
}
type Attribute_N struct {
	N float64 `protobuf:"fixed64,6,opt,name=n,oneof"`
}
type Attribute_A struct {
	A *Attribute_Array `protobuf:"bytes,7,opt,name=a,oneof"`
}

func (*Attribute_S) isAttribute_Attribute() {}
func (*Attribute_I) isAttribute_Attribute() {}
func (*Attribute_N) isAttribute_Attribute() {}
func (*Attribute_A) isAttribute_Attribute() {}

func (m *Attribute) GetAttribute() isAttribute_Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Attribute) GetS() string {
	if x, ok := m.GetAttribute().(*Attribute_S); ok {
		return x.S
	}
	return ""
}

func (m *Attribute) GetI() int64 {
	if x, ok := m.GetAttribute().(*Attribute_I); ok {
		return x.I
	}
	return 0
}

func (m *Attribute) GetN() float64 {
	if x, ok := m.GetAttribute().(*Attribute_N); ok {
		return x.N
	}
	return 0
}

func (m *Attribute) GetA() *Attribute_Array {
	if x, ok := m.GetAttribute().(*Attribute_A); ok {
		return x.A
	}
	return nil
}

func (m *Attribute) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Attribute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Attribute_OneofMarshaler, _Attribute_OneofUnmarshaler, _Attribute_OneofSizer, []interface{}{
		(*Attribute_S)(nil),
		(*Attribute_I)(nil),
		(*Attribute_N)(nil),
		(*Attribute_A)(nil),
	}
}

func _Attribute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Attribute)
	// attribute
	switch x := m.Attribute.(type) {
	case *Attribute_S:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.S)
	case *Attribute_I:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.I))
	case *Attribute_N:
		b.EncodeVarint(6<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.N))
	case *Attribute_A:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.A); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Attribute.Attribute has unexpected type %T", x)
	}
	return nil
}

func _Attribute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Attribute)
	switch tag {
	case 4: // attribute.s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Attribute = &Attribute_S{x}
		return true, err
	case 5: // attribute.i
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Attribute = &Attribute_I{int64(x)}
		return true, err
	case 6: // attribute.n
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Attribute = &Attribute_N{math.Float64frombits(x)}
		return true, err
	case 7: // attribute.a
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Attribute_Array)
		err := b.DecodeMessage(msg)
		m.Attribute = &Attribute_A{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Attribute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Attribute)
	// attribute
	switch x := m.Attribute.(type) {
	case *Attribute_S:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.S)))
		n += len(x.S)
	case *Attribute_I:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.I))
	case *Attribute_N:
		n += 1 // tag and wire
		n += 8
	case *Attribute_A:
		s := proto.Size(x.A)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Attribute_Array struct {
	Elt                  []*Attribute `protobuf:"bytes,1,rep,name=elt" json:"elt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Attribute_Array) Reset()         { *m = Attribute_Array{} }
func (m *Attribute_Array) String() string { return proto.CompactTextString(m) }
func (*Attribute_Array) ProtoMessage()    {}
func (*Attribute_Array) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{2, 0}
}
func (m *Attribute_Array) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute_Array.Unmarshal(m, b)
}
func (m *Attribute_Array) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute_Array.Marshal(b, m, deterministic)
}
func (dst *Attribute_Array) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute_Array.Merge(dst, src)
}
func (m *Attribute_Array) XXX_Size() int {
	return xxx_messageInfo_Attribute_Array.Size(m)
}
func (m *Attribute_Array) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute_Array.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute_Array proto.InternalMessageInfo

func (m *Attribute_Array) GetElt() []*Attribute {
	if m != nil {
		return m.Elt
	}
	return nil
}

type AppName struct {
	Part                 []string `protobuf:"bytes,1,rep,name=part" json:"part,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppName) Reset()         { *m = AppName{} }
func (m *AppName) String() string { return proto.CompactTextString(m) }
func (*AppName) ProtoMessage()    {}
func (*AppName) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{3}
}
func (m *AppName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppName.Unmarshal(m, b)
}
func (m *AppName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppName.Marshal(b, m, deterministic)
}
func (dst *AppName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppName.Merge(dst, src)
}
func (m *AppName) XXX_Size() int {
	return xxx_messageInfo_AppName.Size(m)
}
func (m *AppName) XXX_DiscardUnknown() {
	xxx_messageInfo_AppName.DiscardUnknown(m)
}

var xxx_messageInfo_AppName proto.InternalMessageInfo

func (m *AppName) GetPart() []string {
	if m != nil {
		return m.Part
	}
	return nil
}

type Application struct {
	Name                 *AppName              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	LongName             string                `protobuf:"bytes,2,opt,name=long_name,json=longName" json:"long_name,omitempty"`
	Docstring            string                `protobuf:"bytes,3,opt,name=docstring" json:"docstring,omitempty"`
	Attrs                map[string]*Attribute `protobuf:"bytes,4,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Endpoints            map[string]*Endpoint  `protobuf:"bytes,5,rep,name=endpoints" json:"endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Types                map[string]*Type      `protobuf:"bytes,6,rep,name=types" json:"types,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Views                map[string]*View      `protobuf:"bytes,10,rep,name=views" json:"views,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Mixin2               []*Application        `protobuf:"bytes,8,rep,name=mixin2" json:"mixin2,omitempty"`
	Wrapped              *Application          `protobuf:"bytes,9,opt,name=wrapped" json:"wrapped,omitempty"`
	SourceContext        *SourceContext        `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	DONOTUSEMixin        []*AppName            `protobuf:"bytes,7,rep,name=DONOTUSE_mixin,json=DONOTUSEMixin" json:"DONOTUSE_mixin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{4}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetName() *AppName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Application) GetLongName() string {
	if m != nil {
		return m.LongName
	}
	return ""
}

func (m *Application) GetDocstring() string {
	if m != nil {
		return m.Docstring
	}
	return ""
}

func (m *Application) GetAttrs() map[string]*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (m *Application) GetEndpoints() map[string]*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Application) GetTypes() map[string]*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Application) GetViews() map[string]*View {
	if m != nil {
		return m.Views
	}
	return nil
}

func (m *Application) GetMixin2() []*Application {
	if m != nil {
		return m.Mixin2
	}
	return nil
}

func (m *Application) GetWrapped() *Application {
	if m != nil {
		return m.Wrapped
	}
	return nil
}

func (m *Application) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Application) GetDONOTUSEMixin() []*AppName {
	if m != nil {
		return m.DONOTUSEMixin
	}
	return nil
}

type Endpoint struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	LongName             string                `protobuf:"bytes,2,opt,name=long_name,json=longName" json:"long_name,omitempty"`
	Docstring            string                `protobuf:"bytes,3,opt,name=docstring" json:"docstring,omitempty"`
	Attrs                map[string]*Attribute `protobuf:"bytes,4,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Flag                 []string              `protobuf:"bytes,10,rep,name=flag" json:"flag,omitempty"`
	Source               *AppName              `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	IsPubsub             bool                  `protobuf:"varint,6,opt,name=is_pubsub,json=isPubsub" json:"is_pubsub,omitempty"`
	Param                []*Param              `protobuf:"bytes,9,rep,name=param" json:"param,omitempty"`
	Stmt                 []*Statement          `protobuf:"bytes,7,rep,name=stmt" json:"stmt,omitempty"`
	RestParams           *Endpoint_RestParams  `protobuf:"bytes,8,opt,name=rest_params,json=restParams" json:"rest_params,omitempty"`
	SourceContext        *SourceContext        `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{5}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetLongName() string {
	if m != nil {
		return m.LongName
	}
	return ""
}

func (m *Endpoint) GetDocstring() string {
	if m != nil {
		return m.Docstring
	}
	return ""
}

func (m *Endpoint) GetAttrs() map[string]*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (m *Endpoint) GetFlag() []string {
	if m != nil {
		return m.Flag
	}
	return nil
}

func (m *Endpoint) GetSource() *AppName {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Endpoint) GetIsPubsub() bool {
	if m != nil {
		return m.IsPubsub
	}
	return false
}

func (m *Endpoint) GetParam() []*Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *Endpoint) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *Endpoint) GetRestParams() *Endpoint_RestParams {
	if m != nil {
		return m.RestParams
	}
	return nil
}

func (m *Endpoint) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

type Endpoint_RestParams struct {
	Method               Endpoint_RestParams_Method        `protobuf:"varint,1,opt,name=method,enum=sysl.Endpoint_RestParams_Method" json:"method,omitempty"`
	Path                 string                            `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	QueryParam           []*Endpoint_RestParams_QueryParam `protobuf:"bytes,3,rep,name=query_param,json=queryParam" json:"query_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Endpoint_RestParams) Reset()         { *m = Endpoint_RestParams{} }
func (m *Endpoint_RestParams) String() string { return proto.CompactTextString(m) }
func (*Endpoint_RestParams) ProtoMessage()    {}
func (*Endpoint_RestParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{5, 1}
}
func (m *Endpoint_RestParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_RestParams.Unmarshal(m, b)
}
func (m *Endpoint_RestParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_RestParams.Marshal(b, m, deterministic)
}
func (dst *Endpoint_RestParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_RestParams.Merge(dst, src)
}
func (m *Endpoint_RestParams) XXX_Size() int {
	return xxx_messageInfo_Endpoint_RestParams.Size(m)
}
func (m *Endpoint_RestParams) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_RestParams.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_RestParams proto.InternalMessageInfo

func (m *Endpoint_RestParams) GetMethod() Endpoint_RestParams_Method {
	if m != nil {
		return m.Method
	}
	return Endpoint_RestParams_NO_Method
}

func (m *Endpoint_RestParams) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Endpoint_RestParams) GetQueryParam() []*Endpoint_RestParams_QueryParam {
	if m != nil {
		return m.QueryParam
	}
	return nil
}

type Endpoint_RestParams_QueryParam struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type                 *Type    `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Loc                  bool     `protobuf:"varint,4,opt,name=loc" json:"loc,omitempty"`
	DONOTUSEParam        string   `protobuf:"bytes,3,opt,name=DONOTUSE_param,json=DONOTUSEParam" json:"DONOTUSE_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_RestParams_QueryParam) Reset()         { *m = Endpoint_RestParams_QueryParam{} }
func (m *Endpoint_RestParams_QueryParam) String() string { return proto.CompactTextString(m) }
func (*Endpoint_RestParams_QueryParam) ProtoMessage()    {}
func (*Endpoint_RestParams_QueryParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{5, 1, 0}
}
func (m *Endpoint_RestParams_QueryParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_RestParams_QueryParam.Unmarshal(m, b)
}
func (m *Endpoint_RestParams_QueryParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_RestParams_QueryParam.Marshal(b, m, deterministic)
}
func (dst *Endpoint_RestParams_QueryParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_RestParams_QueryParam.Merge(dst, src)
}
func (m *Endpoint_RestParams_QueryParam) XXX_Size() int {
	return xxx_messageInfo_Endpoint_RestParams_QueryParam.Size(m)
}
func (m *Endpoint_RestParams_QueryParam) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_RestParams_QueryParam.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_RestParams_QueryParam proto.InternalMessageInfo

func (m *Endpoint_RestParams_QueryParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint_RestParams_QueryParam) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Endpoint_RestParams_QueryParam) GetLoc() bool {
	if m != nil {
		return m.Loc
	}
	return false
}

func (m *Endpoint_RestParams_QueryParam) GetDONOTUSEParam() string {
	if m != nil {
		return m.DONOTUSEParam
	}
	return ""
}

type Param struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type                 *Type    `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Param) Reset()         { *m = Param{} }
func (m *Param) String() string { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()    {}
func (*Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{6}
}
func (m *Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Param.Unmarshal(m, b)
}
func (m *Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Param.Marshal(b, m, deterministic)
}
func (dst *Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Param.Merge(dst, src)
}
func (m *Param) XXX_Size() int {
	return xxx_messageInfo_Param.Size(m)
}
func (m *Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Param proto.InternalMessageInfo

func (m *Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Param) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type Statement struct {
	// Types that are valid to be assigned to Stmt:
	//	*Statement_Action
	//	*Statement_Call
	//	*Statement_Cond
	//	*Statement_Loop
	//	*Statement_LoopN
	//	*Statement_Foreach
	//	*Statement_Alt
	//	*Statement_Group
	//	*Statement_Ret
	Stmt                 isStatement_Stmt      `protobuf_oneof:"stmt"`
	Attrs                map[string]*Attribute `protobuf:"bytes,10,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SourceContext        *SourceContext        `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{7}
}
func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (dst *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(dst, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

type isStatement_Stmt interface {
	isStatement_Stmt()
}

type Statement_Action struct {
	Action *Action `protobuf:"bytes,1,opt,name=action,oneof"`
}
type Statement_Call struct {
	Call *Call `protobuf:"bytes,2,opt,name=call,oneof"`
}
type Statement_Cond struct {
	Cond *Cond `protobuf:"bytes,3,opt,name=cond,oneof"`
}
type Statement_Loop struct {
	Loop *Loop `protobuf:"bytes,4,opt,name=loop,oneof"`
}
type Statement_LoopN struct {
	LoopN *LoopN `protobuf:"bytes,5,opt,name=loop_n,json=loopN,oneof"`
}
type Statement_Foreach struct {
	Foreach *Foreach `protobuf:"bytes,9,opt,name=foreach,oneof"`
}
type Statement_Alt struct {
	Alt *Alt `protobuf:"bytes,6,opt,name=alt,oneof"`
}
type Statement_Group struct {
	Group *Group `protobuf:"bytes,7,opt,name=group,oneof"`
}
type Statement_Ret struct {
	Ret *Return `protobuf:"bytes,8,opt,name=ret,oneof"`
}

func (*Statement_Action) isStatement_Stmt()  {}
func (*Statement_Call) isStatement_Stmt()    {}
func (*Statement_Cond) isStatement_Stmt()    {}
func (*Statement_Loop) isStatement_Stmt()    {}
func (*Statement_LoopN) isStatement_Stmt()   {}
func (*Statement_Foreach) isStatement_Stmt() {}
func (*Statement_Alt) isStatement_Stmt()     {}
func (*Statement_Group) isStatement_Stmt()   {}
func (*Statement_Ret) isStatement_Stmt()     {}

func (m *Statement) GetStmt() isStatement_Stmt {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *Statement) GetAction() *Action {
	if x, ok := m.GetStmt().(*Statement_Action); ok {
		return x.Action
	}
	return nil
}

func (m *Statement) GetCall() *Call {
	if x, ok := m.GetStmt().(*Statement_Call); ok {
		return x.Call
	}
	return nil
}

func (m *Statement) GetCond() *Cond {
	if x, ok := m.GetStmt().(*Statement_Cond); ok {
		return x.Cond
	}
	return nil
}

func (m *Statement) GetLoop() *Loop {
	if x, ok := m.GetStmt().(*Statement_Loop); ok {
		return x.Loop
	}
	return nil
}

func (m *Statement) GetLoopN() *LoopN {
	if x, ok := m.GetStmt().(*Statement_LoopN); ok {
		return x.LoopN
	}
	return nil
}

func (m *Statement) GetForeach() *Foreach {
	if x, ok := m.GetStmt().(*Statement_Foreach); ok {
		return x.Foreach
	}
	return nil
}

func (m *Statement) GetAlt() *Alt {
	if x, ok := m.GetStmt().(*Statement_Alt); ok {
		return x.Alt
	}
	return nil
}

func (m *Statement) GetGroup() *Group {
	if x, ok := m.GetStmt().(*Statement_Group); ok {
		return x.Group
	}
	return nil
}

func (m *Statement) GetRet() *Return {
	if x, ok := m.GetStmt().(*Statement_Ret); ok {
		return x.Ret
	}
	return nil
}

func (m *Statement) GetAttrs() map[string]*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (m *Statement) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Statement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Statement_OneofMarshaler, _Statement_OneofUnmarshaler, _Statement_OneofSizer, []interface{}{
		(*Statement_Action)(nil),
		(*Statement_Call)(nil),
		(*Statement_Cond)(nil),
		(*Statement_Loop)(nil),
		(*Statement_LoopN)(nil),
		(*Statement_Foreach)(nil),
		(*Statement_Alt)(nil),
		(*Statement_Group)(nil),
		(*Statement_Ret)(nil),
	}
}

func _Statement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Statement)
	// stmt
	switch x := m.Stmt.(type) {
	case *Statement_Action:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Action); err != nil {
			return err
		}
	case *Statement_Call:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Call); err != nil {
			return err
		}
	case *Statement_Cond:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cond); err != nil {
			return err
		}
	case *Statement_Loop:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Loop); err != nil {
			return err
		}
	case *Statement_LoopN:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoopN); err != nil {
			return err
		}
	case *Statement_Foreach:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Foreach); err != nil {
			return err
		}
	case *Statement_Alt:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Alt); err != nil {
			return err
		}
	case *Statement_Group:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Group); err != nil {
			return err
		}
	case *Statement_Ret:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ret); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Statement.Stmt has unexpected type %T", x)
	}
	return nil
}

func _Statement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Statement)
	switch tag {
	case 1: // stmt.action
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Action)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Action{msg}
		return true, err
	case 2: // stmt.call
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Call)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Call{msg}
		return true, err
	case 3: // stmt.cond
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cond)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Cond{msg}
		return true, err
	case 4: // stmt.loop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Loop)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Loop{msg}
		return true, err
	case 5: // stmt.loop_n
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoopN)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_LoopN{msg}
		return true, err
	case 9: // stmt.foreach
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Foreach)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Foreach{msg}
		return true, err
	case 6: // stmt.alt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Alt)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Alt{msg}
		return true, err
	case 7: // stmt.group
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Group)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Group{msg}
		return true, err
	case 8: // stmt.ret
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Return)
		err := b.DecodeMessage(msg)
		m.Stmt = &Statement_Ret{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Statement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Statement)
	// stmt
	switch x := m.Stmt.(type) {
	case *Statement_Action:
		s := proto.Size(x.Action)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Call:
		s := proto.Size(x.Call)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Cond:
		s := proto.Size(x.Cond)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Loop:
		s := proto.Size(x.Loop)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_LoopN:
		s := proto.Size(x.LoopN)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Foreach:
		s := proto.Size(x.Foreach)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Alt:
		s := proto.Size(x.Alt)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Group:
		s := proto.Size(x.Group)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Statement_Ret:
		s := proto.Size(x.Ret)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Action struct {
	Action               string   `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{8}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (dst *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(dst, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type Call struct {
	Target               *AppName              `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Endpoint             string                `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	Arg                  []*Call_Arg           `protobuf:"bytes,4,rep,name=arg" json:"arg,omitempty"`
	DONOTUSEAttrs        map[string]*Attribute `protobuf:"bytes,3,rep,name=DONOTUSE_attrs,json=DONOTUSEAttrs" json:"DONOTUSE_attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{9}
}
func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (dst *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(dst, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetTarget() *AppName {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Call) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Call) GetArg() []*Call_Arg {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Call) GetDONOTUSEAttrs() map[string]*Attribute {
	if m != nil {
		return m.DONOTUSEAttrs
	}
	return nil
}

type Call_Arg struct {
	Value                *Value   `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DONOTUSEType         *Type    `protobuf:"bytes,2,opt,name=DONOTUSE_type,json=DONOTUSEType" json:"DONOTUSE_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call_Arg) Reset()         { *m = Call_Arg{} }
func (m *Call_Arg) String() string { return proto.CompactTextString(m) }
func (*Call_Arg) ProtoMessage()    {}
func (*Call_Arg) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{9, 0}
}
func (m *Call_Arg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call_Arg.Unmarshal(m, b)
}
func (m *Call_Arg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call_Arg.Marshal(b, m, deterministic)
}
func (dst *Call_Arg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call_Arg.Merge(dst, src)
}
func (m *Call_Arg) XXX_Size() int {
	return xxx_messageInfo_Call_Arg.Size(m)
}
func (m *Call_Arg) XXX_DiscardUnknown() {
	xxx_messageInfo_Call_Arg.DiscardUnknown(m)
}

var xxx_messageInfo_Call_Arg proto.InternalMessageInfo

func (m *Call_Arg) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Call_Arg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Call_Arg) GetDONOTUSEType() *Type {
	if m != nil {
		return m.DONOTUSEType
	}
	return nil
}

type Cond struct {
	Test                 string       `protobuf:"bytes,1,opt,name=test" json:"test,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,2,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Cond) Reset()         { *m = Cond{} }
func (m *Cond) String() string { return proto.CompactTextString(m) }
func (*Cond) ProtoMessage()    {}
func (*Cond) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{10}
}
func (m *Cond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cond.Unmarshal(m, b)
}
func (m *Cond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cond.Marshal(b, m, deterministic)
}
func (dst *Cond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cond.Merge(dst, src)
}
func (m *Cond) XXX_Size() int {
	return xxx_messageInfo_Cond.Size(m)
}
func (m *Cond) XXX_DiscardUnknown() {
	xxx_messageInfo_Cond.DiscardUnknown(m)
}

var xxx_messageInfo_Cond proto.InternalMessageInfo

func (m *Cond) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

func (m *Cond) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Loop struct {
	Mode                 Loop_Mode    `protobuf:"varint,1,opt,name=mode,enum=sysl.Loop_Mode" json:"mode,omitempty"`
	Criterion            string       `protobuf:"bytes,2,opt,name=criterion" json:"criterion,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,3,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Loop) Reset()         { *m = Loop{} }
func (m *Loop) String() string { return proto.CompactTextString(m) }
func (*Loop) ProtoMessage()    {}
func (*Loop) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{11}
}
func (m *Loop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Loop.Unmarshal(m, b)
}
func (m *Loop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Loop.Marshal(b, m, deterministic)
}
func (dst *Loop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loop.Merge(dst, src)
}
func (m *Loop) XXX_Size() int {
	return xxx_messageInfo_Loop.Size(m)
}
func (m *Loop) XXX_DiscardUnknown() {
	xxx_messageInfo_Loop.DiscardUnknown(m)
}

var xxx_messageInfo_Loop proto.InternalMessageInfo

func (m *Loop) GetMode() Loop_Mode {
	if m != nil {
		return m.Mode
	}
	return Loop_NO_Mode
}

func (m *Loop) GetCriterion() string {
	if m != nil {
		return m.Criterion
	}
	return ""
}

func (m *Loop) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type LoopN struct {
	Count                int32        `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,3,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoopN) Reset()         { *m = LoopN{} }
func (m *LoopN) String() string { return proto.CompactTextString(m) }
func (*LoopN) ProtoMessage()    {}
func (*LoopN) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{12}
}
func (m *LoopN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoopN.Unmarshal(m, b)
}
func (m *LoopN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoopN.Marshal(b, m, deterministic)
}
func (dst *LoopN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoopN.Merge(dst, src)
}
func (m *LoopN) XXX_Size() int {
	return xxx_messageInfo_LoopN.Size(m)
}
func (m *LoopN) XXX_DiscardUnknown() {
	xxx_messageInfo_LoopN.DiscardUnknown(m)
}

var xxx_messageInfo_LoopN proto.InternalMessageInfo

func (m *LoopN) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LoopN) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Foreach struct {
	Collection           string       `protobuf:"bytes,1,opt,name=collection" json:"collection,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,3,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Foreach) Reset()         { *m = Foreach{} }
func (m *Foreach) String() string { return proto.CompactTextString(m) }
func (*Foreach) ProtoMessage()    {}
func (*Foreach) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{13}
}
func (m *Foreach) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foreach.Unmarshal(m, b)
}
func (m *Foreach) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foreach.Marshal(b, m, deterministic)
}
func (dst *Foreach) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foreach.Merge(dst, src)
}
func (m *Foreach) XXX_Size() int {
	return xxx_messageInfo_Foreach.Size(m)
}
func (m *Foreach) XXX_DiscardUnknown() {
	xxx_messageInfo_Foreach.DiscardUnknown(m)
}

var xxx_messageInfo_Foreach proto.InternalMessageInfo

func (m *Foreach) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Foreach) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Alt struct {
	Choice               []*Alt_Choice `protobuf:"bytes,1,rep,name=choice" json:"choice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Alt) Reset()         { *m = Alt{} }
func (m *Alt) String() string { return proto.CompactTextString(m) }
func (*Alt) ProtoMessage()    {}
func (*Alt) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{14}
}
func (m *Alt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alt.Unmarshal(m, b)
}
func (m *Alt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alt.Marshal(b, m, deterministic)
}
func (dst *Alt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alt.Merge(dst, src)
}
func (m *Alt) XXX_Size() int {
	return xxx_messageInfo_Alt.Size(m)
}
func (m *Alt) XXX_DiscardUnknown() {
	xxx_messageInfo_Alt.DiscardUnknown(m)
}

var xxx_messageInfo_Alt proto.InternalMessageInfo

func (m *Alt) GetChoice() []*Alt_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

type Alt_Choice struct {
	Cond                 string       `protobuf:"bytes,1,opt,name=cond" json:"cond,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,2,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Alt_Choice) Reset()         { *m = Alt_Choice{} }
func (m *Alt_Choice) String() string { return proto.CompactTextString(m) }
func (*Alt_Choice) ProtoMessage()    {}
func (*Alt_Choice) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{14, 0}
}
func (m *Alt_Choice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alt_Choice.Unmarshal(m, b)
}
func (m *Alt_Choice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alt_Choice.Marshal(b, m, deterministic)
}
func (dst *Alt_Choice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alt_Choice.Merge(dst, src)
}
func (m *Alt_Choice) XXX_Size() int {
	return xxx_messageInfo_Alt_Choice.Size(m)
}
func (m *Alt_Choice) XXX_DiscardUnknown() {
	xxx_messageInfo_Alt_Choice.DiscardUnknown(m)
}

var xxx_messageInfo_Alt_Choice proto.InternalMessageInfo

func (m *Alt_Choice) GetCond() string {
	if m != nil {
		return m.Cond
	}
	return ""
}

func (m *Alt_Choice) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Group struct {
	Title                string       `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Stmt                 []*Statement `protobuf:"bytes,2,rep,name=stmt" json:"stmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{15}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (dst *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(dst, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Group) GetStmt() []*Statement {
	if m != nil {
		return m.Stmt
	}
	return nil
}

type Return struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Return) Reset()         { *m = Return{} }
func (m *Return) String() string { return proto.CompactTextString(m) }
func (*Return) ProtoMessage()    {}
func (*Return) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{16}
}
func (m *Return) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Return.Unmarshal(m, b)
}
func (m *Return) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Return.Marshal(b, m, deterministic)
}
func (dst *Return) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Return.Merge(dst, src)
}
func (m *Return) XXX_Size() int {
	return xxx_messageInfo_Return.Size(m)
}
func (m *Return) XXX_DiscardUnknown() {
	xxx_messageInfo_Return.DiscardUnknown(m)
}

var xxx_messageInfo_Return proto.InternalMessageInfo

func (m *Return) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type Type struct {
	// Types that are valid to be assigned to Type:
	//	*Type_Primitive_
	//	*Type_Enum_
	//	*Type_Tuple_
	//	*Type_List_
	//	*Type_Map_
	//	*Type_OneOf_
	//	*Type_Relation_
	//	*Type_TypeRef
	//	*Type_Set
	//	*Type_NoType_
	Type                 isType_Type           `protobuf_oneof:"type"`
	Attrs                map[string]*Attribute `protobuf:"bytes,8,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Constraint           []*Type_Constraint    `protobuf:"bytes,10,rep,name=constraint" json:"constraint,omitempty"`
	Docstring            string                `protobuf:"bytes,11,opt,name=docstring" json:"docstring,omitempty"`
	Opt                  bool                  `protobuf:"varint,12,opt,name=opt" json:"opt,omitempty"`
	SourceContext        *SourceContext        `protobuf:"bytes,99,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (dst *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(dst, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

type isType_Type interface {
	isType_Type()
}

type Type_Primitive_ struct {
	Primitive Type_Primitive `protobuf:"varint,1,opt,name=primitive,enum=sysl.Type_Primitive,oneof"`
}
type Type_Enum_ struct {
	Enum *Type_Enum `protobuf:"bytes,2,opt,name=enum,oneof"`
}
type Type_Tuple_ struct {
	Tuple *Type_Tuple `protobuf:"bytes,3,opt,name=tuple,oneof"`
}
type Type_List_ struct {
	List *Type_List `protobuf:"bytes,4,opt,name=list,oneof"`
}
type Type_Map_ struct {
	Map *Type_Map `protobuf:"bytes,5,opt,name=map,oneof"`
}
type Type_OneOf_ struct {
	OneOf *Type_OneOf `protobuf:"bytes,6,opt,name=one_of,json=oneOf,oneof"`
}
type Type_Relation_ struct {
	Relation *Type_Relation `protobuf:"bytes,7,opt,name=relation,oneof"`
}
type Type_TypeRef struct {
	TypeRef *ScopedRef `protobuf:"bytes,9,opt,name=type_ref,json=typeRef,oneof"`
}
type Type_Set struct {
	Set *Type `protobuf:"bytes,13,opt,name=set,oneof"`
}
type Type_NoType_ struct {
	NoType *Type_NoType `protobuf:"bytes,14,opt,name=no_type,json=noType,oneof"`
}

func (*Type_Primitive_) isType_Type() {}
func (*Type_Enum_) isType_Type()      {}
func (*Type_Tuple_) isType_Type()     {}
func (*Type_List_) isType_Type()      {}
func (*Type_Map_) isType_Type()       {}
func (*Type_OneOf_) isType_Type()     {}
func (*Type_Relation_) isType_Type()  {}
func (*Type_TypeRef) isType_Type()    {}
func (*Type_Set) isType_Type()        {}
func (*Type_NoType_) isType_Type()    {}

func (m *Type) GetType() isType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type) GetPrimitive() Type_Primitive {
	if x, ok := m.GetType().(*Type_Primitive_); ok {
		return x.Primitive
	}
	return Type_NO_Primitive
}

func (m *Type) GetEnum() *Type_Enum {
	if x, ok := m.GetType().(*Type_Enum_); ok {
		return x.Enum
	}
	return nil
}

func (m *Type) GetTuple() *Type_Tuple {
	if x, ok := m.GetType().(*Type_Tuple_); ok {
		return x.Tuple
	}
	return nil
}

func (m *Type) GetList() *Type_List {
	if x, ok := m.GetType().(*Type_List_); ok {
		return x.List
	}
	return nil
}

func (m *Type) GetMap() *Type_Map {
	if x, ok := m.GetType().(*Type_Map_); ok {
		return x.Map
	}
	return nil
}

func (m *Type) GetOneOf() *Type_OneOf {
	if x, ok := m.GetType().(*Type_OneOf_); ok {
		return x.OneOf
	}
	return nil
}

func (m *Type) GetRelation() *Type_Relation {
	if x, ok := m.GetType().(*Type_Relation_); ok {
		return x.Relation
	}
	return nil
}

func (m *Type) GetTypeRef() *ScopedRef {
	if x, ok := m.GetType().(*Type_TypeRef); ok {
		return x.TypeRef
	}
	return nil
}

func (m *Type) GetSet() *Type {
	if x, ok := m.GetType().(*Type_Set); ok {
		return x.Set
	}
	return nil
}

func (m *Type) GetNoType() *Type_NoType {
	if x, ok := m.GetType().(*Type_NoType_); ok {
		return x.NoType
	}
	return nil
}

func (m *Type) GetAttrs() map[string]*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (m *Type) GetConstraint() []*Type_Constraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *Type) GetDocstring() string {
	if m != nil {
		return m.Docstring
	}
	return ""
}

func (m *Type) GetOpt() bool {
	if m != nil {
		return m.Opt
	}
	return false
}

func (m *Type) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Type) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Type_OneofMarshaler, _Type_OneofUnmarshaler, _Type_OneofSizer, []interface{}{
		(*Type_Primitive_)(nil),
		(*Type_Enum_)(nil),
		(*Type_Tuple_)(nil),
		(*Type_List_)(nil),
		(*Type_Map_)(nil),
		(*Type_OneOf_)(nil),
		(*Type_Relation_)(nil),
		(*Type_TypeRef)(nil),
		(*Type_Set)(nil),
		(*Type_NoType_)(nil),
	}
}

func _Type_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Type)
	// type
	switch x := m.Type.(type) {
	case *Type_Primitive_:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Primitive))
	case *Type_Enum_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Enum); err != nil {
			return err
		}
	case *Type_Tuple_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tuple); err != nil {
			return err
		}
	case *Type_List_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *Type_Map_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Map); err != nil {
			return err
		}
	case *Type_OneOf_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneOf); err != nil {
			return err
		}
	case *Type_Relation_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Relation); err != nil {
			return err
		}
	case *Type_TypeRef:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeRef); err != nil {
			return err
		}
	case *Type_Set:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Set); err != nil {
			return err
		}
	case *Type_NoType_:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NoType); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Type.Type has unexpected type %T", x)
	}
	return nil
}

func _Type_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Type)
	switch tag {
	case 1: // type.primitive
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &Type_Primitive_{Type_Primitive(x)}
		return true, err
	case 2: // type.enum
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_Enum)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Enum_{msg}
		return true, err
	case 3: // type.tuple
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_Tuple)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Tuple_{msg}
		return true, err
	case 4: // type.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_List)
		err := b.DecodeMessage(msg)
		m.Type = &Type_List_{msg}
		return true, err
	case 5: // type.map
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_Map)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Map_{msg}
		return true, err
	case 6: // type.one_of
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_OneOf)
		err := b.DecodeMessage(msg)
		m.Type = &Type_OneOf_{msg}
		return true, err
	case 7: // type.relation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_Relation)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Relation_{msg}
		return true, err
	case 9: // type.type_ref
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ScopedRef)
		err := b.DecodeMessage(msg)
		m.Type = &Type_TypeRef{msg}
		return true, err
	case 13: // type.set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type)
		err := b.DecodeMessage(msg)
		m.Type = &Type_Set{msg}
		return true, err
	case 14: // type.no_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Type_NoType)
		err := b.DecodeMessage(msg)
		m.Type = &Type_NoType_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Type_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Type)
	// type
	switch x := m.Type.(type) {
	case *Type_Primitive_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Primitive))
	case *Type_Enum_:
		s := proto.Size(x.Enum)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Tuple_:
		s := proto.Size(x.Tuple)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_List_:
		s := proto.Size(x.List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Map_:
		s := proto.Size(x.Map)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_OneOf_:
		s := proto.Size(x.OneOf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Relation_:
		s := proto.Size(x.Relation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_TypeRef:
		s := proto.Size(x.TypeRef)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_Set:
		s := proto.Size(x.Set)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Type_NoType_:
		s := proto.Size(x.NoType)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Type_Enum struct {
	Items                map[string]int64 `protobuf:"bytes,1,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Type_Enum) Reset()         { *m = Type_Enum{} }
func (m *Type_Enum) String() string { return proto.CompactTextString(m) }
func (*Type_Enum) ProtoMessage()    {}
func (*Type_Enum) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 1}
}
func (m *Type_Enum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Enum.Unmarshal(m, b)
}
func (m *Type_Enum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Enum.Marshal(b, m, deterministic)
}
func (dst *Type_Enum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Enum.Merge(dst, src)
}
func (m *Type_Enum) XXX_Size() int {
	return xxx_messageInfo_Type_Enum.Size(m)
}
func (m *Type_Enum) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Enum.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Enum proto.InternalMessageInfo

func (m *Type_Enum) GetItems() map[string]int64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Type_Tuple struct {
	AttrDefs             map[string]*Type             `protobuf:"bytes,1,rep,name=attr_defs,json=attrDefs" json:"attr_defs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FUTUREFields         map[string]*Type_Tuple_Field `protobuf:"bytes,2,rep,name=FUTURE_fields,json=FUTUREFields" json:"FUTURE_fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Type_Tuple) Reset()         { *m = Type_Tuple{} }
func (m *Type_Tuple) String() string { return proto.CompactTextString(m) }
func (*Type_Tuple) ProtoMessage()    {}
func (*Type_Tuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 2}
}
func (m *Type_Tuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Tuple.Unmarshal(m, b)
}
func (m *Type_Tuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Tuple.Marshal(b, m, deterministic)
}
func (dst *Type_Tuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Tuple.Merge(dst, src)
}
func (m *Type_Tuple) XXX_Size() int {
	return xxx_messageInfo_Type_Tuple.Size(m)
}
func (m *Type_Tuple) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Tuple.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Tuple proto.InternalMessageInfo

func (m *Type_Tuple) GetAttrDefs() map[string]*Type {
	if m != nil {
		return m.AttrDefs
	}
	return nil
}

func (m *Type_Tuple) GetFUTUREFields() map[string]*Type_Tuple_Field {
	if m != nil {
		return m.FUTUREFields
	}
	return nil
}

type Type_Tuple_Field struct {
	Type                 *Type    `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	MinRepeats           int64    `protobuf:"varint,2,opt,name=min_repeats,json=minRepeats" json:"min_repeats,omitempty"`
	MaxRepeats           int64    `protobuf:"varint,3,opt,name=max_repeats,json=maxRepeats" json:"max_repeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Tuple_Field) Reset()         { *m = Type_Tuple_Field{} }
func (m *Type_Tuple_Field) String() string { return proto.CompactTextString(m) }
func (*Type_Tuple_Field) ProtoMessage()    {}
func (*Type_Tuple_Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 2, 2}
}
func (m *Type_Tuple_Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Tuple_Field.Unmarshal(m, b)
}
func (m *Type_Tuple_Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Tuple_Field.Marshal(b, m, deterministic)
}
func (dst *Type_Tuple_Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Tuple_Field.Merge(dst, src)
}
func (m *Type_Tuple_Field) XXX_Size() int {
	return xxx_messageInfo_Type_Tuple_Field.Size(m)
}
func (m *Type_Tuple_Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Tuple_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Tuple_Field proto.InternalMessageInfo

func (m *Type_Tuple_Field) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type_Tuple_Field) GetMinRepeats() int64 {
	if m != nil {
		return m.MinRepeats
	}
	return 0
}

func (m *Type_Tuple_Field) GetMaxRepeats() int64 {
	if m != nil {
		return m.MaxRepeats
	}
	return 0
}

type Type_List struct {
	Type                 *Type    `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_List) Reset()         { *m = Type_List{} }
func (m *Type_List) String() string { return proto.CompactTextString(m) }
func (*Type_List) ProtoMessage()    {}
func (*Type_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 3}
}
func (m *Type_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_List.Unmarshal(m, b)
}
func (m *Type_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_List.Marshal(b, m, deterministic)
}
func (dst *Type_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_List.Merge(dst, src)
}
func (m *Type_List) XXX_Size() int {
	return xxx_messageInfo_Type_List.Size(m)
}
func (m *Type_List) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_List.DiscardUnknown(m)
}

var xxx_messageInfo_Type_List proto.InternalMessageInfo

func (m *Type_List) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type Type_Map struct {
	Key                  *Type    `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                *Type    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Map) Reset()         { *m = Type_Map{} }
func (m *Type_Map) String() string { return proto.CompactTextString(m) }
func (*Type_Map) ProtoMessage()    {}
func (*Type_Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 4}
}
func (m *Type_Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Map.Unmarshal(m, b)
}
func (m *Type_Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Map.Marshal(b, m, deterministic)
}
func (dst *Type_Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Map.Merge(dst, src)
}
func (m *Type_Map) XXX_Size() int {
	return xxx_messageInfo_Type_Map.Size(m)
}
func (m *Type_Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Map proto.InternalMessageInfo

func (m *Type_Map) GetKey() *Type {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Type_Map) GetValue() *Type {
	if m != nil {
		return m.Value
	}
	return nil
}

type Type_OneOf struct {
	Type                 []*Type  `protobuf:"bytes,1,rep,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_OneOf) Reset()         { *m = Type_OneOf{} }
func (m *Type_OneOf) String() string { return proto.CompactTextString(m) }
func (*Type_OneOf) ProtoMessage()    {}
func (*Type_OneOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 5}
}
func (m *Type_OneOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_OneOf.Unmarshal(m, b)
}
func (m *Type_OneOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_OneOf.Marshal(b, m, deterministic)
}
func (dst *Type_OneOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_OneOf.Merge(dst, src)
}
func (m *Type_OneOf) XXX_Size() int {
	return xxx_messageInfo_Type_OneOf.Size(m)
}
func (m *Type_OneOf) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_OneOf.DiscardUnknown(m)
}

var xxx_messageInfo_Type_OneOf proto.InternalMessageInfo

func (m *Type_OneOf) GetType() []*Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type Type_Relation struct {
	AttrDefs             map[string]*Type     `protobuf:"bytes,1,rep,name=attr_defs,json=attrDefs" json:"attr_defs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PrimaryKey           *Type_Relation_Key   `protobuf:"bytes,2,opt,name=primary_key,json=primaryKey" json:"primary_key,omitempty"`
	Key                  []*Type_Relation_Key `protobuf:"bytes,3,rep,name=key" json:"key,omitempty"`
	Inject               []string             `protobuf:"bytes,4,rep,name=inject" json:"inject,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Type_Relation) Reset()         { *m = Type_Relation{} }
func (m *Type_Relation) String() string { return proto.CompactTextString(m) }
func (*Type_Relation) ProtoMessage()    {}
func (*Type_Relation) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 6}
}
func (m *Type_Relation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Relation.Unmarshal(m, b)
}
func (m *Type_Relation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Relation.Marshal(b, m, deterministic)
}
func (dst *Type_Relation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Relation.Merge(dst, src)
}
func (m *Type_Relation) XXX_Size() int {
	return xxx_messageInfo_Type_Relation.Size(m)
}
func (m *Type_Relation) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Relation.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Relation proto.InternalMessageInfo

func (m *Type_Relation) GetAttrDefs() map[string]*Type {
	if m != nil {
		return m.AttrDefs
	}
	return nil
}

func (m *Type_Relation) GetPrimaryKey() *Type_Relation_Key {
	if m != nil {
		return m.PrimaryKey
	}
	return nil
}

func (m *Type_Relation) GetKey() []*Type_Relation_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Type_Relation) GetInject() []string {
	if m != nil {
		return m.Inject
	}
	return nil
}

type Type_Relation_Key struct {
	AttrName             []string `protobuf:"bytes,1,rep,name=attr_name,json=attrName" json:"attr_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Relation_Key) Reset()         { *m = Type_Relation_Key{} }
func (m *Type_Relation_Key) String() string { return proto.CompactTextString(m) }
func (*Type_Relation_Key) ProtoMessage()    {}
func (*Type_Relation_Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 6, 1}
}
func (m *Type_Relation_Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Relation_Key.Unmarshal(m, b)
}
func (m *Type_Relation_Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Relation_Key.Marshal(b, m, deterministic)
}
func (dst *Type_Relation_Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Relation_Key.Merge(dst, src)
}
func (m *Type_Relation_Key) XXX_Size() int {
	return xxx_messageInfo_Type_Relation_Key.Size(m)
}
func (m *Type_Relation_Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Relation_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Relation_Key proto.InternalMessageInfo

func (m *Type_Relation_Key) GetAttrName() []string {
	if m != nil {
		return m.AttrName
	}
	return nil
}

type Type_Foreign struct {
	App                  *AppName `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Relation             string   `protobuf:"bytes,2,opt,name=relation" json:"relation,omitempty"`
	AttrName             string   `protobuf:"bytes,3,opt,name=attr_name,json=attrName" json:"attr_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Foreign) Reset()         { *m = Type_Foreign{} }
func (m *Type_Foreign) String() string { return proto.CompactTextString(m) }
func (*Type_Foreign) ProtoMessage()    {}
func (*Type_Foreign) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 7}
}
func (m *Type_Foreign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Foreign.Unmarshal(m, b)
}
func (m *Type_Foreign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Foreign.Marshal(b, m, deterministic)
}
func (dst *Type_Foreign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Foreign.Merge(dst, src)
}
func (m *Type_Foreign) XXX_Size() int {
	return xxx_messageInfo_Type_Foreign.Size(m)
}
func (m *Type_Foreign) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Foreign.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Foreign proto.InternalMessageInfo

func (m *Type_Foreign) GetApp() *AppName {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Type_Foreign) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *Type_Foreign) GetAttrName() string {
	if m != nil {
		return m.AttrName
	}
	return ""
}

type Type_Constraint struct {
	Range                *Type_Constraint_Range      `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Length               *Type_Constraint_Length     `protobuf:"bytes,2,opt,name=length" json:"length,omitempty"`
	Resolution           *Type_Constraint_Resolution `protobuf:"bytes,3,opt,name=resolution" json:"resolution,omitempty"`
	Precision            int32                       `protobuf:"varint,4,opt,name=precision" json:"precision,omitempty"`
	Scale                int32                       `protobuf:"varint,5,opt,name=scale" json:"scale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Type_Constraint) Reset()         { *m = Type_Constraint{} }
func (m *Type_Constraint) String() string { return proto.CompactTextString(m) }
func (*Type_Constraint) ProtoMessage()    {}
func (*Type_Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 8}
}
func (m *Type_Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Constraint.Unmarshal(m, b)
}
func (m *Type_Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Constraint.Marshal(b, m, deterministic)
}
func (dst *Type_Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Constraint.Merge(dst, src)
}
func (m *Type_Constraint) XXX_Size() int {
	return xxx_messageInfo_Type_Constraint.Size(m)
}
func (m *Type_Constraint) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Constraint.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Constraint proto.InternalMessageInfo

func (m *Type_Constraint) GetRange() *Type_Constraint_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Type_Constraint) GetLength() *Type_Constraint_Length {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Type_Constraint) GetResolution() *Type_Constraint_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Type_Constraint) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *Type_Constraint) GetScale() int32 {
	if m != nil {
		return m.Scale
	}
	return 0
}

type Type_Constraint_Range struct {
	Min                  *Value   `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max                  *Value   `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Constraint_Range) Reset()         { *m = Type_Constraint_Range{} }
func (m *Type_Constraint_Range) String() string { return proto.CompactTextString(m) }
func (*Type_Constraint_Range) ProtoMessage()    {}
func (*Type_Constraint_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 8, 0}
}
func (m *Type_Constraint_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Constraint_Range.Unmarshal(m, b)
}
func (m *Type_Constraint_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Constraint_Range.Marshal(b, m, deterministic)
}
func (dst *Type_Constraint_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Constraint_Range.Merge(dst, src)
}
func (m *Type_Constraint_Range) XXX_Size() int {
	return xxx_messageInfo_Type_Constraint_Range.Size(m)
}
func (m *Type_Constraint_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Constraint_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Constraint_Range proto.InternalMessageInfo

func (m *Type_Constraint_Range) GetMin() *Value {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Type_Constraint_Range) GetMax() *Value {
	if m != nil {
		return m.Max
	}
	return nil
}

type Type_Constraint_Length struct {
	Min                  int64    `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max                  int64    `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Constraint_Length) Reset()         { *m = Type_Constraint_Length{} }
func (m *Type_Constraint_Length) String() string { return proto.CompactTextString(m) }
func (*Type_Constraint_Length) ProtoMessage()    {}
func (*Type_Constraint_Length) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 8, 1}
}
func (m *Type_Constraint_Length) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Constraint_Length.Unmarshal(m, b)
}
func (m *Type_Constraint_Length) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Constraint_Length.Marshal(b, m, deterministic)
}
func (dst *Type_Constraint_Length) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Constraint_Length.Merge(dst, src)
}
func (m *Type_Constraint_Length) XXX_Size() int {
	return xxx_messageInfo_Type_Constraint_Length.Size(m)
}
func (m *Type_Constraint_Length) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Constraint_Length.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Constraint_Length proto.InternalMessageInfo

func (m *Type_Constraint_Length) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Type_Constraint_Length) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

// e.g.: 3 decimal places = {base = 10, index = -3}
type Type_Constraint_Resolution struct {
	Base                 int32    `protobuf:"varint,1,opt,name=base" json:"base,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_Constraint_Resolution) Reset()         { *m = Type_Constraint_Resolution{} }
func (m *Type_Constraint_Resolution) String() string { return proto.CompactTextString(m) }
func (*Type_Constraint_Resolution) ProtoMessage()    {}
func (*Type_Constraint_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 8, 2}
}
func (m *Type_Constraint_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_Constraint_Resolution.Unmarshal(m, b)
}
func (m *Type_Constraint_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_Constraint_Resolution.Marshal(b, m, deterministic)
}
func (dst *Type_Constraint_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_Constraint_Resolution.Merge(dst, src)
}
func (m *Type_Constraint_Resolution) XXX_Size() int {
	return xxx_messageInfo_Type_Constraint_Resolution.Size(m)
}
func (m *Type_Constraint_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_Constraint_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Type_Constraint_Resolution proto.InternalMessageInfo

func (m *Type_Constraint_Resolution) GetBase() int32 {
	if m != nil {
		return m.Base
	}
	return 0
}

func (m *Type_Constraint_Resolution) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Type_NoType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type_NoType) Reset()         { *m = Type_NoType{} }
func (m *Type_NoType) String() string { return proto.CompactTextString(m) }
func (*Type_NoType) ProtoMessage()    {}
func (*Type_NoType) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{17, 9}
}
func (m *Type_NoType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type_NoType.Unmarshal(m, b)
}
func (m *Type_NoType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type_NoType.Marshal(b, m, deterministic)
}
func (dst *Type_NoType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type_NoType.Merge(dst, src)
}
func (m *Type_NoType) XXX_Size() int {
	return xxx_messageInfo_Type_NoType.Size(m)
}
func (m *Type_NoType) XXX_DiscardUnknown() {
	xxx_messageInfo_Type_NoType.DiscardUnknown(m)
}

var xxx_messageInfo_Type_NoType proto.InternalMessageInfo

type View struct {
	Param                []*Param              `protobuf:"bytes,1,rep,name=param" json:"param,omitempty"`
	RetType              *Type                 `protobuf:"bytes,2,opt,name=ret_type,json=retType" json:"ret_type,omitempty"`
	Expr                 *Expr                 `protobuf:"bytes,3,opt,name=expr" json:"expr,omitempty"`
	Views                map[string]*View      `protobuf:"bytes,4,rep,name=views" json:"views,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Attrs                map[string]*Attribute `protobuf:"bytes,8,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *View) Reset()         { *m = View{} }
func (m *View) String() string { return proto.CompactTextString(m) }
func (*View) ProtoMessage()    {}
func (*View) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{18}
}
func (m *View) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_View.Unmarshal(m, b)
}
func (m *View) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_View.Marshal(b, m, deterministic)
}
func (dst *View) XXX_Merge(src proto.Message) {
	xxx_messageInfo_View.Merge(dst, src)
}
func (m *View) XXX_Size() int {
	return xxx_messageInfo_View.Size(m)
}
func (m *View) XXX_DiscardUnknown() {
	xxx_messageInfo_View.DiscardUnknown(m)
}

var xxx_messageInfo_View proto.InternalMessageInfo

func (m *View) GetParam() []*Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *View) GetRetType() *Type {
	if m != nil {
		return m.RetType
	}
	return nil
}

func (m *View) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *View) GetViews() map[string]*View {
	if m != nil {
		return m.Views
	}
	return nil
}

func (m *View) GetAttrs() map[string]*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Expr struct {
	// Types that are valid to be assigned to Expr:
	//	*Expr_Name
	//	*Expr_Literal
	//	*Expr_GetAttr_
	//	*Expr_Transform_
	//	*Expr_Ifelse
	//	*Expr_Call_
	//	*Expr_Unexpr
	//	*Expr_Binexpr
	//	*Expr_Relexpr
	//	*Expr_Navigate_
	//	*Expr_List_
	//	*Expr_Set
	//	*Expr_Tuple_
	Expr                 isExpr_Expr `protobuf_oneof:"expr"`
	Type                 *Type       `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Expr) Reset()         { *m = Expr{} }
func (m *Expr) String() string { return proto.CompactTextString(m) }
func (*Expr) ProtoMessage()    {}
func (*Expr) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19}
}
func (m *Expr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr.Unmarshal(m, b)
}
func (m *Expr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr.Marshal(b, m, deterministic)
}
func (dst *Expr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr.Merge(dst, src)
}
func (m *Expr) XXX_Size() int {
	return xxx_messageInfo_Expr.Size(m)
}
func (m *Expr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr proto.InternalMessageInfo

type isExpr_Expr interface {
	isExpr_Expr()
}

type Expr_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,oneof"`
}
type Expr_Literal struct {
	Literal *Value `protobuf:"bytes,2,opt,name=literal,oneof"`
}
type Expr_GetAttr_ struct {
	GetAttr *Expr_GetAttr `protobuf:"bytes,3,opt,name=get_attr,json=getAttr,oneof"`
}
type Expr_Transform_ struct {
	Transform *Expr_Transform `protobuf:"bytes,4,opt,name=transform,oneof"`
}
type Expr_Ifelse struct {
	Ifelse *Expr_IfElse `protobuf:"bytes,5,opt,name=ifelse,oneof"`
}
type Expr_Call_ struct {
	Call *Expr_Call `protobuf:"bytes,6,opt,name=call,oneof"`
}
type Expr_Unexpr struct {
	Unexpr *Expr_UnExpr `protobuf:"bytes,7,opt,name=unexpr,oneof"`
}
type Expr_Binexpr struct {
	Binexpr *Expr_BinExpr `protobuf:"bytes,8,opt,name=binexpr,oneof"`
}
type Expr_Relexpr struct {
	Relexpr *Expr_RelExpr `protobuf:"bytes,12,opt,name=relexpr,oneof"`
}
type Expr_Navigate_ struct {
	Navigate *Expr_Navigate `protobuf:"bytes,9,opt,name=navigate,oneof"`
}
type Expr_List_ struct {
	List *Expr_List `protobuf:"bytes,10,opt,name=list,oneof"`
}
type Expr_Set struct {
	Set *Expr_List `protobuf:"bytes,11,opt,name=set,oneof"`
}
type Expr_Tuple_ struct {
	Tuple *Expr_Tuple `protobuf:"bytes,14,opt,name=tuple,oneof"`
}

func (*Expr_Name) isExpr_Expr()       {}
func (*Expr_Literal) isExpr_Expr()    {}
func (*Expr_GetAttr_) isExpr_Expr()   {}
func (*Expr_Transform_) isExpr_Expr() {}
func (*Expr_Ifelse) isExpr_Expr()     {}
func (*Expr_Call_) isExpr_Expr()      {}
func (*Expr_Unexpr) isExpr_Expr()     {}
func (*Expr_Binexpr) isExpr_Expr()    {}
func (*Expr_Relexpr) isExpr_Expr()    {}
func (*Expr_Navigate_) isExpr_Expr()  {}
func (*Expr_List_) isExpr_Expr()      {}
func (*Expr_Set) isExpr_Expr()        {}
func (*Expr_Tuple_) isExpr_Expr()     {}

func (m *Expr) GetExpr() isExpr_Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *Expr) GetName() string {
	if x, ok := m.GetExpr().(*Expr_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Expr) GetLiteral() *Value {
	if x, ok := m.GetExpr().(*Expr_Literal); ok {
		return x.Literal
	}
	return nil
}

func (m *Expr) GetGetAttr() *Expr_GetAttr {
	if x, ok := m.GetExpr().(*Expr_GetAttr_); ok {
		return x.GetAttr
	}
	return nil
}

func (m *Expr) GetTransform() *Expr_Transform {
	if x, ok := m.GetExpr().(*Expr_Transform_); ok {
		return x.Transform
	}
	return nil
}

func (m *Expr) GetIfelse() *Expr_IfElse {
	if x, ok := m.GetExpr().(*Expr_Ifelse); ok {
		return x.Ifelse
	}
	return nil
}

func (m *Expr) GetCall() *Expr_Call {
	if x, ok := m.GetExpr().(*Expr_Call_); ok {
		return x.Call
	}
	return nil
}

func (m *Expr) GetUnexpr() *Expr_UnExpr {
	if x, ok := m.GetExpr().(*Expr_Unexpr); ok {
		return x.Unexpr
	}
	return nil
}

func (m *Expr) GetBinexpr() *Expr_BinExpr {
	if x, ok := m.GetExpr().(*Expr_Binexpr); ok {
		return x.Binexpr
	}
	return nil
}

func (m *Expr) GetRelexpr() *Expr_RelExpr {
	if x, ok := m.GetExpr().(*Expr_Relexpr); ok {
		return x.Relexpr
	}
	return nil
}

func (m *Expr) GetNavigate() *Expr_Navigate {
	if x, ok := m.GetExpr().(*Expr_Navigate_); ok {
		return x.Navigate
	}
	return nil
}

func (m *Expr) GetList() *Expr_List {
	if x, ok := m.GetExpr().(*Expr_List_); ok {
		return x.List
	}
	return nil
}

func (m *Expr) GetSet() *Expr_List {
	if x, ok := m.GetExpr().(*Expr_Set); ok {
		return x.Set
	}
	return nil
}

func (m *Expr) GetTuple() *Expr_Tuple {
	if x, ok := m.GetExpr().(*Expr_Tuple_); ok {
		return x.Tuple
	}
	return nil
}

func (m *Expr) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Expr) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Expr_OneofMarshaler, _Expr_OneofUnmarshaler, _Expr_OneofSizer, []interface{}{
		(*Expr_Name)(nil),
		(*Expr_Literal)(nil),
		(*Expr_GetAttr_)(nil),
		(*Expr_Transform_)(nil),
		(*Expr_Ifelse)(nil),
		(*Expr_Call_)(nil),
		(*Expr_Unexpr)(nil),
		(*Expr_Binexpr)(nil),
		(*Expr_Relexpr)(nil),
		(*Expr_Navigate_)(nil),
		(*Expr_List_)(nil),
		(*Expr_Set)(nil),
		(*Expr_Tuple_)(nil),
	}
}

func _Expr_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Expr)
	// expr
	switch x := m.Expr.(type) {
	case *Expr_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *Expr_Literal:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Literal); err != nil {
			return err
		}
	case *Expr_GetAttr_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetAttr); err != nil {
			return err
		}
	case *Expr_Transform_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transform); err != nil {
			return err
		}
	case *Expr_Ifelse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ifelse); err != nil {
			return err
		}
	case *Expr_Call_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Call); err != nil {
			return err
		}
	case *Expr_Unexpr:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Unexpr); err != nil {
			return err
		}
	case *Expr_Binexpr:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Binexpr); err != nil {
			return err
		}
	case *Expr_Relexpr:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Relexpr); err != nil {
			return err
		}
	case *Expr_Navigate_:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Navigate); err != nil {
			return err
		}
	case *Expr_List_:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *Expr_Set:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Set); err != nil {
			return err
		}
	case *Expr_Tuple_:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tuple); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Expr.Expr has unexpected type %T", x)
	}
	return nil
}

func _Expr_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Expr)
	switch tag {
	case 1: // expr.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Expr = &Expr_Name{x}
		return true, err
	case 2: // expr.literal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Literal{msg}
		return true, err
	case 3: // expr.get_attr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_GetAttr)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_GetAttr_{msg}
		return true, err
	case 4: // expr.transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Transform)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Transform_{msg}
		return true, err
	case 5: // expr.ifelse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_IfElse)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Ifelse{msg}
		return true, err
	case 6: // expr.call
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Call)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Call_{msg}
		return true, err
	case 7: // expr.unexpr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_UnExpr)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Unexpr{msg}
		return true, err
	case 8: // expr.binexpr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_BinExpr)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Binexpr{msg}
		return true, err
	case 12: // expr.relexpr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_RelExpr)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Relexpr{msg}
		return true, err
	case 9: // expr.navigate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Navigate)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Navigate_{msg}
		return true, err
	case 10: // expr.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_List)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_List_{msg}
		return true, err
	case 11: // expr.set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_List)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Set{msg}
		return true, err
	case 14: // expr.tuple
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Tuple)
		err := b.DecodeMessage(msg)
		m.Expr = &Expr_Tuple_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Expr_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Expr)
	// expr
	switch x := m.Expr.(type) {
	case *Expr_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *Expr_Literal:
		s := proto.Size(x.Literal)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_GetAttr_:
		s := proto.Size(x.GetAttr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Transform_:
		s := proto.Size(x.Transform)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Ifelse:
		s := proto.Size(x.Ifelse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Call_:
		s := proto.Size(x.Call)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Unexpr:
		s := proto.Size(x.Unexpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Binexpr:
		s := proto.Size(x.Binexpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Relexpr:
		s := proto.Size(x.Relexpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Navigate_:
		s := proto.Size(x.Navigate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_List_:
		s := proto.Size(x.List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Set:
		s := proto.Size(x.Set)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Tuple_:
		s := proto.Size(x.Tuple)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Expr_GetAttr struct {
	Arg                  *Expr    `protobuf:"bytes,1,opt,name=arg" json:"arg,omitempty"`
	Attr                 string   `protobuf:"bytes,2,opt,name=attr" json:"attr,omitempty"`
	Nullsafe             bool     `protobuf:"varint,3,opt,name=nullsafe" json:"nullsafe,omitempty"`
	Setof                bool     `protobuf:"varint,4,opt,name=setof" json:"setof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_GetAttr) Reset()         { *m = Expr_GetAttr{} }
func (m *Expr_GetAttr) String() string { return proto.CompactTextString(m) }
func (*Expr_GetAttr) ProtoMessage()    {}
func (*Expr_GetAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 0}
}
func (m *Expr_GetAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_GetAttr.Unmarshal(m, b)
}
func (m *Expr_GetAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_GetAttr.Marshal(b, m, deterministic)
}
func (dst *Expr_GetAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_GetAttr.Merge(dst, src)
}
func (m *Expr_GetAttr) XXX_Size() int {
	return xxx_messageInfo_Expr_GetAttr.Size(m)
}
func (m *Expr_GetAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_GetAttr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_GetAttr proto.InternalMessageInfo

func (m *Expr_GetAttr) GetArg() *Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Expr_GetAttr) GetAttr() string {
	if m != nil {
		return m.Attr
	}
	return ""
}

func (m *Expr_GetAttr) GetNullsafe() bool {
	if m != nil {
		return m.Nullsafe
	}
	return false
}

func (m *Expr_GetAttr) GetSetof() bool {
	if m != nil {
		return m.Setof
	}
	return false
}

type Expr_Navigate struct {
	Arg                  *Expr    `protobuf:"bytes,1,opt,name=arg" json:"arg,omitempty"`
	Attr                 string   `protobuf:"bytes,2,opt,name=attr" json:"attr,omitempty"`
	Nullsafe             bool     `protobuf:"varint,3,opt,name=nullsafe" json:"nullsafe,omitempty"`
	Setof                bool     `protobuf:"varint,4,opt,name=setof" json:"setof,omitempty"`
	Via                  string   `protobuf:"bytes,5,opt,name=via" json:"via,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Navigate) Reset()         { *m = Expr_Navigate{} }
func (m *Expr_Navigate) String() string { return proto.CompactTextString(m) }
func (*Expr_Navigate) ProtoMessage()    {}
func (*Expr_Navigate) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 1}
}
func (m *Expr_Navigate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Navigate.Unmarshal(m, b)
}
func (m *Expr_Navigate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Navigate.Marshal(b, m, deterministic)
}
func (dst *Expr_Navigate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Navigate.Merge(dst, src)
}
func (m *Expr_Navigate) XXX_Size() int {
	return xxx_messageInfo_Expr_Navigate.Size(m)
}
func (m *Expr_Navigate) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Navigate.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Navigate proto.InternalMessageInfo

func (m *Expr_Navigate) GetArg() *Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Expr_Navigate) GetAttr() string {
	if m != nil {
		return m.Attr
	}
	return ""
}

func (m *Expr_Navigate) GetNullsafe() bool {
	if m != nil {
		return m.Nullsafe
	}
	return false
}

func (m *Expr_Navigate) GetSetof() bool {
	if m != nil {
		return m.Setof
	}
	return false
}

func (m *Expr_Navigate) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

type Expr_List struct {
	Expr                 []*Expr  `protobuf:"bytes,1,rep,name=expr" json:"expr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_List) Reset()         { *m = Expr_List{} }
func (m *Expr_List) String() string { return proto.CompactTextString(m) }
func (*Expr_List) ProtoMessage()    {}
func (*Expr_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 2}
}
func (m *Expr_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_List.Unmarshal(m, b)
}
func (m *Expr_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_List.Marshal(b, m, deterministic)
}
func (dst *Expr_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_List.Merge(dst, src)
}
func (m *Expr_List) XXX_Size() int {
	return xxx_messageInfo_Expr_List.Size(m)
}
func (m *Expr_List) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_List.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_List proto.InternalMessageInfo

func (m *Expr_List) GetExpr() []*Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

type Expr_Transform struct {
	Arg                  *Expr                  `protobuf:"bytes,1,opt,name=arg" json:"arg,omitempty"`
	Scopevar             string                 `protobuf:"bytes,2,opt,name=scopevar" json:"scopevar,omitempty"`
	Stmt                 []*Expr_Transform_Stmt `protobuf:"bytes,3,rep,name=stmt" json:"stmt,omitempty"`
	AllAttrs             bool                   `protobuf:"varint,4,opt,name=all_attrs,json=allAttrs" json:"all_attrs,omitempty"`
	ExceptAttrs          []string               `protobuf:"bytes,5,rep,name=except_attrs,json=exceptAttrs" json:"except_attrs,omitempty"`
	Nullsafe             bool                   `protobuf:"varint,6,opt,name=nullsafe" json:"nullsafe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Expr_Transform) Reset()         { *m = Expr_Transform{} }
func (m *Expr_Transform) String() string { return proto.CompactTextString(m) }
func (*Expr_Transform) ProtoMessage()    {}
func (*Expr_Transform) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 3}
}
func (m *Expr_Transform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Transform.Unmarshal(m, b)
}
func (m *Expr_Transform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Transform.Marshal(b, m, deterministic)
}
func (dst *Expr_Transform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Transform.Merge(dst, src)
}
func (m *Expr_Transform) XXX_Size() int {
	return xxx_messageInfo_Expr_Transform.Size(m)
}
func (m *Expr_Transform) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Transform.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Transform proto.InternalMessageInfo

func (m *Expr_Transform) GetArg() *Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Expr_Transform) GetScopevar() string {
	if m != nil {
		return m.Scopevar
	}
	return ""
}

func (m *Expr_Transform) GetStmt() []*Expr_Transform_Stmt {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *Expr_Transform) GetAllAttrs() bool {
	if m != nil {
		return m.AllAttrs
	}
	return false
}

func (m *Expr_Transform) GetExceptAttrs() []string {
	if m != nil {
		return m.ExceptAttrs
	}
	return nil
}

func (m *Expr_Transform) GetNullsafe() bool {
	if m != nil {
		return m.Nullsafe
	}
	return false
}

type Expr_Transform_Stmt struct {
	// Types that are valid to be assigned to Stmt:
	//	*Expr_Transform_Stmt_Assign_
	//	*Expr_Transform_Stmt_Let
	//	*Expr_Transform_Stmt_Inject
	Stmt                 isExpr_Transform_Stmt_Stmt `protobuf_oneof:"stmt"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Expr_Transform_Stmt) Reset()         { *m = Expr_Transform_Stmt{} }
func (m *Expr_Transform_Stmt) String() string { return proto.CompactTextString(m) }
func (*Expr_Transform_Stmt) ProtoMessage()    {}
func (*Expr_Transform_Stmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 3, 0}
}
func (m *Expr_Transform_Stmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Transform_Stmt.Unmarshal(m, b)
}
func (m *Expr_Transform_Stmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Transform_Stmt.Marshal(b, m, deterministic)
}
func (dst *Expr_Transform_Stmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Transform_Stmt.Merge(dst, src)
}
func (m *Expr_Transform_Stmt) XXX_Size() int {
	return xxx_messageInfo_Expr_Transform_Stmt.Size(m)
}
func (m *Expr_Transform_Stmt) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Transform_Stmt.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Transform_Stmt proto.InternalMessageInfo

type isExpr_Transform_Stmt_Stmt interface {
	isExpr_Transform_Stmt_Stmt()
}

type Expr_Transform_Stmt_Assign_ struct {
	Assign *Expr_Transform_Stmt_Assign `protobuf:"bytes,1,opt,name=assign,oneof"`
}
type Expr_Transform_Stmt_Let struct {
	Let *Expr_Transform_Stmt_Assign `protobuf:"bytes,2,opt,name=let,oneof"`
}
type Expr_Transform_Stmt_Inject struct {
	Inject *Expr `protobuf:"bytes,4,opt,name=inject,oneof"`
}

func (*Expr_Transform_Stmt_Assign_) isExpr_Transform_Stmt_Stmt() {}
func (*Expr_Transform_Stmt_Let) isExpr_Transform_Stmt_Stmt()     {}
func (*Expr_Transform_Stmt_Inject) isExpr_Transform_Stmt_Stmt()  {}

func (m *Expr_Transform_Stmt) GetStmt() isExpr_Transform_Stmt_Stmt {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *Expr_Transform_Stmt) GetAssign() *Expr_Transform_Stmt_Assign {
	if x, ok := m.GetStmt().(*Expr_Transform_Stmt_Assign_); ok {
		return x.Assign
	}
	return nil
}

func (m *Expr_Transform_Stmt) GetLet() *Expr_Transform_Stmt_Assign {
	if x, ok := m.GetStmt().(*Expr_Transform_Stmt_Let); ok {
		return x.Let
	}
	return nil
}

func (m *Expr_Transform_Stmt) GetInject() *Expr {
	if x, ok := m.GetStmt().(*Expr_Transform_Stmt_Inject); ok {
		return x.Inject
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Expr_Transform_Stmt) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Expr_Transform_Stmt_OneofMarshaler, _Expr_Transform_Stmt_OneofUnmarshaler, _Expr_Transform_Stmt_OneofSizer, []interface{}{
		(*Expr_Transform_Stmt_Assign_)(nil),
		(*Expr_Transform_Stmt_Let)(nil),
		(*Expr_Transform_Stmt_Inject)(nil),
	}
}

func _Expr_Transform_Stmt_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Expr_Transform_Stmt)
	// stmt
	switch x := m.Stmt.(type) {
	case *Expr_Transform_Stmt_Assign_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Assign); err != nil {
			return err
		}
	case *Expr_Transform_Stmt_Let:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Let); err != nil {
			return err
		}
	case *Expr_Transform_Stmt_Inject:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Inject); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Expr_Transform_Stmt.Stmt has unexpected type %T", x)
	}
	return nil
}

func _Expr_Transform_Stmt_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Expr_Transform_Stmt)
	switch tag {
	case 1: // stmt.assign
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Transform_Stmt_Assign)
		err := b.DecodeMessage(msg)
		m.Stmt = &Expr_Transform_Stmt_Assign_{msg}
		return true, err
	case 2: // stmt.let
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr_Transform_Stmt_Assign)
		err := b.DecodeMessage(msg)
		m.Stmt = &Expr_Transform_Stmt_Let{msg}
		return true, err
	case 4: // stmt.inject
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Expr)
		err := b.DecodeMessage(msg)
		m.Stmt = &Expr_Transform_Stmt_Inject{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Expr_Transform_Stmt_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Expr_Transform_Stmt)
	// stmt
	switch x := m.Stmt.(type) {
	case *Expr_Transform_Stmt_Assign_:
		s := proto.Size(x.Assign)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Transform_Stmt_Let:
		s := proto.Size(x.Let)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Expr_Transform_Stmt_Inject:
		s := proto.Size(x.Inject)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Expr_Transform_Stmt_Assign struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Expr                 *Expr    `protobuf:"bytes,2,opt,name=expr" json:"expr,omitempty"`
	Table                bool     `protobuf:"varint,3,opt,name=table" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Transform_Stmt_Assign) Reset()         { *m = Expr_Transform_Stmt_Assign{} }
func (m *Expr_Transform_Stmt_Assign) String() string { return proto.CompactTextString(m) }
func (*Expr_Transform_Stmt_Assign) ProtoMessage()    {}
func (*Expr_Transform_Stmt_Assign) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 3, 0, 0}
}
func (m *Expr_Transform_Stmt_Assign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Transform_Stmt_Assign.Unmarshal(m, b)
}
func (m *Expr_Transform_Stmt_Assign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Transform_Stmt_Assign.Marshal(b, m, deterministic)
}
func (dst *Expr_Transform_Stmt_Assign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Transform_Stmt_Assign.Merge(dst, src)
}
func (m *Expr_Transform_Stmt_Assign) XXX_Size() int {
	return xxx_messageInfo_Expr_Transform_Stmt_Assign.Size(m)
}
func (m *Expr_Transform_Stmt_Assign) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Transform_Stmt_Assign.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Transform_Stmt_Assign proto.InternalMessageInfo

func (m *Expr_Transform_Stmt_Assign) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Expr_Transform_Stmt_Assign) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *Expr_Transform_Stmt_Assign) GetTable() bool {
	if m != nil {
		return m.Table
	}
	return false
}

type Expr_IfElse struct {
	Cond                 *Expr    `protobuf:"bytes,1,opt,name=cond" json:"cond,omitempty"`
	IfTrue               *Expr    `protobuf:"bytes,2,opt,name=if_true,json=ifTrue" json:"if_true,omitempty"`
	IfFalse              *Expr    `protobuf:"bytes,3,opt,name=if_false,json=ifFalse" json:"if_false,omitempty"`
	Nullsafe             bool     `protobuf:"varint,4,opt,name=nullsafe" json:"nullsafe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_IfElse) Reset()         { *m = Expr_IfElse{} }
func (m *Expr_IfElse) String() string { return proto.CompactTextString(m) }
func (*Expr_IfElse) ProtoMessage()    {}
func (*Expr_IfElse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 4}
}
func (m *Expr_IfElse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_IfElse.Unmarshal(m, b)
}
func (m *Expr_IfElse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_IfElse.Marshal(b, m, deterministic)
}
func (dst *Expr_IfElse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_IfElse.Merge(dst, src)
}
func (m *Expr_IfElse) XXX_Size() int {
	return xxx_messageInfo_Expr_IfElse.Size(m)
}
func (m *Expr_IfElse) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_IfElse.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_IfElse proto.InternalMessageInfo

func (m *Expr_IfElse) GetCond() *Expr {
	if m != nil {
		return m.Cond
	}
	return nil
}

func (m *Expr_IfElse) GetIfTrue() *Expr {
	if m != nil {
		return m.IfTrue
	}
	return nil
}

func (m *Expr_IfElse) GetIfFalse() *Expr {
	if m != nil {
		return m.IfFalse
	}
	return nil
}

func (m *Expr_IfElse) GetNullsafe() bool {
	if m != nil {
		return m.Nullsafe
	}
	return false
}

type Expr_Call struct {
	Func                 string   `protobuf:"bytes,1,opt,name=func" json:"func,omitempty"`
	Arg                  []*Expr  `protobuf:"bytes,2,rep,name=arg" json:"arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Call) Reset()         { *m = Expr_Call{} }
func (m *Expr_Call) String() string { return proto.CompactTextString(m) }
func (*Expr_Call) ProtoMessage()    {}
func (*Expr_Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 5}
}
func (m *Expr_Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Call.Unmarshal(m, b)
}
func (m *Expr_Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Call.Marshal(b, m, deterministic)
}
func (dst *Expr_Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Call.Merge(dst, src)
}
func (m *Expr_Call) XXX_Size() int {
	return xxx_messageInfo_Expr_Call.Size(m)
}
func (m *Expr_Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Call proto.InternalMessageInfo

func (m *Expr_Call) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *Expr_Call) GetArg() []*Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

type Expr_UnExpr struct {
	Op                   Expr_UnExpr_Op `protobuf:"varint,1,opt,name=op,enum=sysl.Expr_UnExpr_Op" json:"op,omitempty"`
	Arg                  *Expr          `protobuf:"bytes,2,opt,name=arg" json:"arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Expr_UnExpr) Reset()         { *m = Expr_UnExpr{} }
func (m *Expr_UnExpr) String() string { return proto.CompactTextString(m) }
func (*Expr_UnExpr) ProtoMessage()    {}
func (*Expr_UnExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 6}
}
func (m *Expr_UnExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_UnExpr.Unmarshal(m, b)
}
func (m *Expr_UnExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_UnExpr.Marshal(b, m, deterministic)
}
func (dst *Expr_UnExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_UnExpr.Merge(dst, src)
}
func (m *Expr_UnExpr) XXX_Size() int {
	return xxx_messageInfo_Expr_UnExpr.Size(m)
}
func (m *Expr_UnExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_UnExpr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_UnExpr proto.InternalMessageInfo

func (m *Expr_UnExpr) GetOp() Expr_UnExpr_Op {
	if m != nil {
		return m.Op
	}
	return Expr_UnExpr_NO_Op
}

func (m *Expr_UnExpr) GetArg() *Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

type Expr_BinExpr struct {
	Op                   Expr_BinExpr_Op `protobuf:"varint,1,opt,name=op,enum=sysl.Expr_BinExpr_Op" json:"op,omitempty"`
	Lhs                  *Expr           `protobuf:"bytes,2,opt,name=lhs" json:"lhs,omitempty"`
	Rhs                  *Expr           `protobuf:"bytes,3,opt,name=rhs" json:"rhs,omitempty"`
	Scopevar             string          `protobuf:"bytes,4,opt,name=scopevar" json:"scopevar,omitempty"`
	AttrName             []string        `protobuf:"bytes,5,rep,name=attr_name,json=attrName" json:"attr_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Expr_BinExpr) Reset()         { *m = Expr_BinExpr{} }
func (m *Expr_BinExpr) String() string { return proto.CompactTextString(m) }
func (*Expr_BinExpr) ProtoMessage()    {}
func (*Expr_BinExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 7}
}
func (m *Expr_BinExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_BinExpr.Unmarshal(m, b)
}
func (m *Expr_BinExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_BinExpr.Marshal(b, m, deterministic)
}
func (dst *Expr_BinExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_BinExpr.Merge(dst, src)
}
func (m *Expr_BinExpr) XXX_Size() int {
	return xxx_messageInfo_Expr_BinExpr.Size(m)
}
func (m *Expr_BinExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_BinExpr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_BinExpr proto.InternalMessageInfo

func (m *Expr_BinExpr) GetOp() Expr_BinExpr_Op {
	if m != nil {
		return m.Op
	}
	return Expr_BinExpr_NO_Op
}

func (m *Expr_BinExpr) GetLhs() *Expr {
	if m != nil {
		return m.Lhs
	}
	return nil
}

func (m *Expr_BinExpr) GetRhs() *Expr {
	if m != nil {
		return m.Rhs
	}
	return nil
}

func (m *Expr_BinExpr) GetScopevar() string {
	if m != nil {
		return m.Scopevar
	}
	return ""
}

func (m *Expr_BinExpr) GetAttrName() []string {
	if m != nil {
		return m.AttrName
	}
	return nil
}

// TODO: Migrate BinExpr ops to RelExpr as appropriate.
type Expr_RelExpr struct {
	Op                   Expr_RelExpr_Op `protobuf:"varint,1,opt,name=op,enum=sysl.Expr_RelExpr_Op" json:"op,omitempty"`
	Target               *Expr           `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Arg                  []*Expr         `protobuf:"bytes,3,rep,name=arg" json:"arg,omitempty"`
	Scopevar             string          `protobuf:"bytes,4,opt,name=scopevar" json:"scopevar,omitempty"`
	Descending           []bool          `protobuf:"varint,5,rep,packed,name=descending" json:"descending,omitempty"`
	AttrName             []string        `protobuf:"bytes,6,rep,name=attr_name,json=attrName" json:"attr_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Expr_RelExpr) Reset()         { *m = Expr_RelExpr{} }
func (m *Expr_RelExpr) String() string { return proto.CompactTextString(m) }
func (*Expr_RelExpr) ProtoMessage()    {}
func (*Expr_RelExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 8}
}
func (m *Expr_RelExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_RelExpr.Unmarshal(m, b)
}
func (m *Expr_RelExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_RelExpr.Marshal(b, m, deterministic)
}
func (dst *Expr_RelExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_RelExpr.Merge(dst, src)
}
func (m *Expr_RelExpr) XXX_Size() int {
	return xxx_messageInfo_Expr_RelExpr.Size(m)
}
func (m *Expr_RelExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_RelExpr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_RelExpr proto.InternalMessageInfo

func (m *Expr_RelExpr) GetOp() Expr_RelExpr_Op {
	if m != nil {
		return m.Op
	}
	return Expr_RelExpr_NO_Op
}

func (m *Expr_RelExpr) GetTarget() *Expr {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Expr_RelExpr) GetArg() []*Expr {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *Expr_RelExpr) GetScopevar() string {
	if m != nil {
		return m.Scopevar
	}
	return ""
}

func (m *Expr_RelExpr) GetDescending() []bool {
	if m != nil {
		return m.Descending
	}
	return nil
}

func (m *Expr_RelExpr) GetAttrName() []string {
	if m != nil {
		return m.AttrName
	}
	return nil
}

type Expr_Tuple struct {
	Attrs                map[string]*Expr `protobuf:"bytes,1,rep,name=attrs" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Expr_Tuple) Reset()         { *m = Expr_Tuple{} }
func (m *Expr_Tuple) String() string { return proto.CompactTextString(m) }
func (*Expr_Tuple) ProtoMessage()    {}
func (*Expr_Tuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{19, 9}
}
func (m *Expr_Tuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Tuple.Unmarshal(m, b)
}
func (m *Expr_Tuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Tuple.Marshal(b, m, deterministic)
}
func (dst *Expr_Tuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Tuple.Merge(dst, src)
}
func (m *Expr_Tuple) XXX_Size() int {
	return xxx_messageInfo_Expr_Tuple.Size(m)
}
func (m *Expr_Tuple) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Tuple.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Tuple proto.InternalMessageInfo

func (m *Expr_Tuple) GetAttrs() map[string]*Expr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_B
	//	*Value_I
	//	*Value_D
	//	*Value_S
	//	*Value_Decimal
	//	*Value_Data
	//	*Value_Enum
	//	*Value_List_
	//	*Value_Map_
	//	*Value_Set
	//	*Value_Null_
	//	*Value_Uuid
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{20}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_B struct {
	B bool `protobuf:"varint,1,opt,name=b,oneof"`
}
type Value_I struct {
	I int64 `protobuf:"varint,2,opt,name=i,oneof"`
}
type Value_D struct {
	D float64 `protobuf:"fixed64,3,opt,name=d,oneof"`
}
type Value_S struct {
	S string `protobuf:"bytes,4,opt,name=s,oneof"`
}
type Value_Decimal struct {
	Decimal string `protobuf:"bytes,11,opt,name=decimal,oneof"`
}
type Value_Data struct {
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3,oneof"`
}
type Value_Enum struct {
	Enum int64 `protobuf:"varint,6,opt,name=enum,oneof"`
}
type Value_List_ struct {
	List *Value_List `protobuf:"bytes,7,opt,name=list,oneof"`
}
type Value_Map_ struct {
	Map *Value_Map `protobuf:"bytes,8,opt,name=map,oneof"`
}
type Value_Set struct {
	Set *Value_List `protobuf:"bytes,9,opt,name=set,oneof"`
}
type Value_Null_ struct {
	Null *Value_Null `protobuf:"bytes,10,opt,name=null,oneof"`
}
type Value_Uuid struct {
	Uuid []byte `protobuf:"bytes,12,opt,name=uuid,proto3,oneof"`
}

func (*Value_B) isValue_Value()       {}
func (*Value_I) isValue_Value()       {}
func (*Value_D) isValue_Value()       {}
func (*Value_S) isValue_Value()       {}
func (*Value_Decimal) isValue_Value() {}
func (*Value_Data) isValue_Value()    {}
func (*Value_Enum) isValue_Value()    {}
func (*Value_List_) isValue_Value()   {}
func (*Value_Map_) isValue_Value()    {}
func (*Value_Set) isValue_Value()     {}
func (*Value_Null_) isValue_Value()   {}
func (*Value_Uuid) isValue_Value()    {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetB() bool {
	if x, ok := m.GetValue().(*Value_B); ok {
		return x.B
	}
	return false
}

func (m *Value) GetI() int64 {
	if x, ok := m.GetValue().(*Value_I); ok {
		return x.I
	}
	return 0
}

func (m *Value) GetD() float64 {
	if x, ok := m.GetValue().(*Value_D); ok {
		return x.D
	}
	return 0
}

func (m *Value) GetS() string {
	if x, ok := m.GetValue().(*Value_S); ok {
		return x.S
	}
	return ""
}

func (m *Value) GetDecimal() string {
	if x, ok := m.GetValue().(*Value_Decimal); ok {
		return x.Decimal
	}
	return ""
}

func (m *Value) GetData() []byte {
	if x, ok := m.GetValue().(*Value_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Value) GetEnum() int64 {
	if x, ok := m.GetValue().(*Value_Enum); ok {
		return x.Enum
	}
	return 0
}

func (m *Value) GetList() *Value_List {
	if x, ok := m.GetValue().(*Value_List_); ok {
		return x.List
	}
	return nil
}

func (m *Value) GetMap() *Value_Map {
	if x, ok := m.GetValue().(*Value_Map_); ok {
		return x.Map
	}
	return nil
}

func (m *Value) GetSet() *Value_List {
	if x, ok := m.GetValue().(*Value_Set); ok {
		return x.Set
	}
	return nil
}

func (m *Value) GetNull() *Value_Null {
	if x, ok := m.GetValue().(*Value_Null_); ok {
		return x.Null
	}
	return nil
}

func (m *Value) GetUuid() []byte {
	if x, ok := m.GetValue().(*Value_Uuid); ok {
		return x.Uuid
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_B)(nil),
		(*Value_I)(nil),
		(*Value_D)(nil),
		(*Value_S)(nil),
		(*Value_Decimal)(nil),
		(*Value_Data)(nil),
		(*Value_Enum)(nil),
		(*Value_List_)(nil),
		(*Value_Map_)(nil),
		(*Value_Set)(nil),
		(*Value_Null_)(nil),
		(*Value_Uuid)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_B:
		t := uint64(0)
		if x.B {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_I:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.I))
	case *Value_D:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.D))
	case *Value_S:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.S)
	case *Value_Decimal:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Decimal)
	case *Value_Data:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Data)
	case *Value_Enum:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Enum))
	case *Value_List_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *Value_Map_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Map); err != nil {
			return err
		}
	case *Value_Set:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Set); err != nil {
			return err
		}
	case *Value_Null_:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Null); err != nil {
			return err
		}
	case *Value_Uuid:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Uuid)
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // value.b
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_B{x != 0}
		return true, err
	case 2: // value.i
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_I{int64(x)}
		return true, err
	case 3: // value.d
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Value_D{math.Float64frombits(x)}
		return true, err
	case 4: // value.s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_S{x}
		return true, err
	case 11: // value.decimal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_Decimal{x}
		return true, err
	case 5: // value.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Value_Data{x}
		return true, err
	case 6: // value.enum
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_Enum{int64(x)}
		return true, err
	case 7: // value.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value_List)
		err := b.DecodeMessage(msg)
		m.Value = &Value_List_{msg}
		return true, err
	case 8: // value.map
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value_Map)
		err := b.DecodeMessage(msg)
		m.Value = &Value_Map_{msg}
		return true, err
	case 9: // value.set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value_List)
		err := b.DecodeMessage(msg)
		m.Value = &Value_Set{msg}
		return true, err
	case 10: // value.null
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value_Null)
		err := b.DecodeMessage(msg)
		m.Value = &Value_Null_{msg}
		return true, err
	case 12: // value.uuid
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Value_Uuid{x}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_B:
		n += 1 // tag and wire
		n += 1
	case *Value_I:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.I))
	case *Value_D:
		n += 1 // tag and wire
		n += 8
	case *Value_S:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.S)))
		n += len(x.S)
	case *Value_Decimal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Decimal)))
		n += len(x.Decimal)
	case *Value_Data:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *Value_Enum:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Enum))
	case *Value_List_:
		s := proto.Size(x.List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_Map_:
		s := proto.Size(x.Map)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_Set:
		s := proto.Size(x.Set)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_Null_:
		s := proto.Size(x.Null)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_Uuid:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Uuid)))
		n += len(x.Uuid)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value_List struct {
	Value                []*Value `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value_List) Reset()         { *m = Value_List{} }
func (m *Value_List) String() string { return proto.CompactTextString(m) }
func (*Value_List) ProtoMessage()    {}
func (*Value_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{20, 0}
}
func (m *Value_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value_List.Unmarshal(m, b)
}
func (m *Value_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value_List.Marshal(b, m, deterministic)
}
func (dst *Value_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value_List.Merge(dst, src)
}
func (m *Value_List) XXX_Size() int {
	return xxx_messageInfo_Value_List.Size(m)
}
func (m *Value_List) XXX_DiscardUnknown() {
	xxx_messageInfo_Value_List.DiscardUnknown(m)
}

var xxx_messageInfo_Value_List proto.InternalMessageInfo

func (m *Value_List) GetValue() []*Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Value_Map struct {
	Items                map[string]*Value `protobuf:"bytes,1,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Value_Map) Reset()         { *m = Value_Map{} }
func (m *Value_Map) String() string { return proto.CompactTextString(m) }
func (*Value_Map) ProtoMessage()    {}
func (*Value_Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{20, 1}
}
func (m *Value_Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value_Map.Unmarshal(m, b)
}
func (m *Value_Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value_Map.Marshal(b, m, deterministic)
}
func (dst *Value_Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value_Map.Merge(dst, src)
}
func (m *Value_Map) XXX_Size() int {
	return xxx_messageInfo_Value_Map.Size(m)
}
func (m *Value_Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Value_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Value_Map proto.InternalMessageInfo

func (m *Value_Map) GetItems() map[string]*Value {
	if m != nil {
		return m.Items
	}
	return nil
}

type Value_Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value_Null) Reset()         { *m = Value_Null{} }
func (m *Value_Null) String() string { return proto.CompactTextString(m) }
func (*Value_Null) ProtoMessage()    {}
func (*Value_Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{20, 2}
}
func (m *Value_Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value_Null.Unmarshal(m, b)
}
func (m *Value_Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value_Null.Marshal(b, m, deterministic)
}
func (dst *Value_Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value_Null.Merge(dst, src)
}
func (m *Value_Null) XXX_Size() int {
	return xxx_messageInfo_Value_Null.Size(m)
}
func (m *Value_Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Value_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Value_Null proto.InternalMessageInfo

type ScopedRef struct {
	Context              *Scope   `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Ref                  *Scope   `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScopedRef) Reset()         { *m = ScopedRef{} }
func (m *ScopedRef) String() string { return proto.CompactTextString(m) }
func (*ScopedRef) ProtoMessage()    {}
func (*ScopedRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{21}
}
func (m *ScopedRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScopedRef.Unmarshal(m, b)
}
func (m *ScopedRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScopedRef.Marshal(b, m, deterministic)
}
func (dst *ScopedRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedRef.Merge(dst, src)
}
func (m *ScopedRef) XXX_Size() int {
	return xxx_messageInfo_ScopedRef.Size(m)
}
func (m *ScopedRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedRef.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedRef proto.InternalMessageInfo

func (m *ScopedRef) GetContext() *Scope {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ScopedRef) GetRef() *Scope {
	if m != nil {
		return m.Ref
	}
	return nil
}

type Scope struct {
	Appname              *AppName `protobuf:"bytes,1,opt,name=appname" json:"appname,omitempty"`
	Path                 []string `protobuf:"bytes,2,rep,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scope) Reset()         { *m = Scope{} }
func (m *Scope) String() string { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()    {}
func (*Scope) Descriptor() ([]byte, []int) {
	return fileDescriptor_sysl_5137add53988ebde, []int{22}
}
func (m *Scope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scope.Unmarshal(m, b)
}
func (m *Scope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scope.Marshal(b, m, deterministic)
}
func (dst *Scope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scope.Merge(dst, src)
}
func (m *Scope) XXX_Size() int {
	return xxx_messageInfo_Scope.Size(m)
}
func (m *Scope) XXX_DiscardUnknown() {
	xxx_messageInfo_Scope.DiscardUnknown(m)
}

var xxx_messageInfo_Scope proto.InternalMessageInfo

func (m *Scope) GetAppname() *AppName {
	if m != nil {
		return m.Appname
	}
	return nil
}

func (m *Scope) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*SourceContext)(nil), "sysl.SourceContext")
	proto.RegisterType((*SourceContext_Location)(nil), "sysl.SourceContext.Location")
	proto.RegisterType((*Module)(nil), "sysl.Module")
	proto.RegisterMapType((map[string]*Application)(nil), "sysl.Module.AppsEntry")
	proto.RegisterMapType((map[string]*Type)(nil), "sysl.Module.TypesEntry")
	proto.RegisterType((*Attribute)(nil), "sysl.Attribute")
	proto.RegisterType((*Attribute_Array)(nil), "sysl.Attribute.Array")
	proto.RegisterType((*AppName)(nil), "sysl.AppName")
	proto.RegisterType((*Application)(nil), "sysl.Application")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.Application.AttrsEntry")
	proto.RegisterMapType((map[string]*Endpoint)(nil), "sysl.Application.EndpointsEntry")
	proto.RegisterMapType((map[string]*Type)(nil), "sysl.Application.TypesEntry")
	proto.RegisterMapType((map[string]*View)(nil), "sysl.Application.ViewsEntry")
	proto.RegisterType((*Endpoint)(nil), "sysl.Endpoint")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.Endpoint.AttrsEntry")
	proto.RegisterType((*Endpoint_RestParams)(nil), "sysl.Endpoint.RestParams")
	proto.RegisterType((*Endpoint_RestParams_QueryParam)(nil), "sysl.Endpoint.RestParams.QueryParam")
	proto.RegisterType((*Param)(nil), "sysl.Param")
	proto.RegisterType((*Statement)(nil), "sysl.Statement")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.Statement.AttrsEntry")
	proto.RegisterType((*Action)(nil), "sysl.Action")
	proto.RegisterType((*Call)(nil), "sysl.Call")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.Call.DONOTUSEAttrsEntry")
	proto.RegisterType((*Call_Arg)(nil), "sysl.Call.Arg")
	proto.RegisterType((*Cond)(nil), "sysl.Cond")
	proto.RegisterType((*Loop)(nil), "sysl.Loop")
	proto.RegisterType((*LoopN)(nil), "sysl.LoopN")
	proto.RegisterType((*Foreach)(nil), "sysl.Foreach")
	proto.RegisterType((*Alt)(nil), "sysl.Alt")
	proto.RegisterType((*Alt_Choice)(nil), "sysl.Alt.Choice")
	proto.RegisterType((*Group)(nil), "sysl.Group")
	proto.RegisterType((*Return)(nil), "sysl.Return")
	proto.RegisterType((*Type)(nil), "sysl.Type")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.Type.AttrsEntry")
	proto.RegisterType((*Type_Enum)(nil), "sysl.Type.Enum")
	proto.RegisterMapType((map[string]int64)(nil), "sysl.Type.Enum.ItemsEntry")
	proto.RegisterType((*Type_Tuple)(nil), "sysl.Type.Tuple")
	proto.RegisterMapType((map[string]*Type)(nil), "sysl.Type.Tuple.AttrDefsEntry")
	proto.RegisterMapType((map[string]*Type_Tuple_Field)(nil), "sysl.Type.Tuple.FUTUREFieldsEntry")
	proto.RegisterType((*Type_Tuple_Field)(nil), "sysl.Type.Tuple.Field")
	proto.RegisterType((*Type_List)(nil), "sysl.Type.List")
	proto.RegisterType((*Type_Map)(nil), "sysl.Type.Map")
	proto.RegisterType((*Type_OneOf)(nil), "sysl.Type.OneOf")
	proto.RegisterType((*Type_Relation)(nil), "sysl.Type.Relation")
	proto.RegisterMapType((map[string]*Type)(nil), "sysl.Type.Relation.AttrDefsEntry")
	proto.RegisterType((*Type_Relation_Key)(nil), "sysl.Type.Relation.Key")
	proto.RegisterType((*Type_Foreign)(nil), "sysl.Type.Foreign")
	proto.RegisterType((*Type_Constraint)(nil), "sysl.Type.Constraint")
	proto.RegisterType((*Type_Constraint_Range)(nil), "sysl.Type.Constraint.Range")
	proto.RegisterType((*Type_Constraint_Length)(nil), "sysl.Type.Constraint.Length")
	proto.RegisterType((*Type_Constraint_Resolution)(nil), "sysl.Type.Constraint.Resolution")
	proto.RegisterType((*Type_NoType)(nil), "sysl.Type.NoType")
	proto.RegisterType((*View)(nil), "sysl.View")
	proto.RegisterMapType((map[string]*Attribute)(nil), "sysl.View.AttrsEntry")
	proto.RegisterMapType((map[string]*View)(nil), "sysl.View.ViewsEntry")
	proto.RegisterType((*Expr)(nil), "sysl.Expr")
	proto.RegisterType((*Expr_GetAttr)(nil), "sysl.Expr.GetAttr")
	proto.RegisterType((*Expr_Navigate)(nil), "sysl.Expr.Navigate")
	proto.RegisterType((*Expr_List)(nil), "sysl.Expr.List")
	proto.RegisterType((*Expr_Transform)(nil), "sysl.Expr.Transform")
	proto.RegisterType((*Expr_Transform_Stmt)(nil), "sysl.Expr.Transform.Stmt")
	proto.RegisterType((*Expr_Transform_Stmt_Assign)(nil), "sysl.Expr.Transform.Stmt.Assign")
	proto.RegisterType((*Expr_IfElse)(nil), "sysl.Expr.IfElse")
	proto.RegisterType((*Expr_Call)(nil), "sysl.Expr.Call")
	proto.RegisterType((*Expr_UnExpr)(nil), "sysl.Expr.UnExpr")
	proto.RegisterType((*Expr_BinExpr)(nil), "sysl.Expr.BinExpr")
	proto.RegisterType((*Expr_RelExpr)(nil), "sysl.Expr.RelExpr")
	proto.RegisterType((*Expr_Tuple)(nil), "sysl.Expr.Tuple")
	proto.RegisterMapType((map[string]*Expr)(nil), "sysl.Expr.Tuple.AttrsEntry")
	proto.RegisterType((*Value)(nil), "sysl.Value")
	proto.RegisterType((*Value_List)(nil), "sysl.Value.List")
	proto.RegisterType((*Value_Map)(nil), "sysl.Value.Map")
	proto.RegisterMapType((map[string]*Value)(nil), "sysl.Value.Map.ItemsEntry")
	proto.RegisterType((*Value_Null)(nil), "sysl.Value.Null")
	proto.RegisterType((*ScopedRef)(nil), "sysl.ScopedRef")
	proto.RegisterType((*Scope)(nil), "sysl.Scope")
	proto.RegisterEnum("sysl.Delta", Delta_name, Delta_value)
	proto.RegisterEnum("sysl.Endpoint_RestParams_Method", Endpoint_RestParams_Method_name, Endpoint_RestParams_Method_value)
	proto.RegisterEnum("sysl.Endpoint_RestParams_QueryParam_Loc", Endpoint_RestParams_QueryParam_Loc_name, Endpoint_RestParams_QueryParam_Loc_value)
	proto.RegisterEnum("sysl.Loop_Mode", Loop_Mode_name, Loop_Mode_value)
	proto.RegisterEnum("sysl.Type_Primitive", Type_Primitive_name, Type_Primitive_value)
	proto.RegisterEnum("sysl.Expr_UnExpr_Op", Expr_UnExpr_Op_name, Expr_UnExpr_Op_value)
	proto.RegisterEnum("sysl.Expr_BinExpr_Op", Expr_BinExpr_Op_name, Expr_BinExpr_Op_value)
	proto.RegisterEnum("sysl.Expr_RelExpr_Op", Expr_RelExpr_Op_name, Expr_RelExpr_Op_value)
}

func init() { proto.RegisterFile("pb/sysl.proto", fileDescriptor_sysl_5137add53988ebde) }

var fileDescriptor_sysl_5137add53988ebde = []byte{
	// 3660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x5a, 0x4f, 0x93, 0xdb, 0xc6,
	0x72, 0x5f, 0x12, 0x7f, 0x48, 0x36, 0x77, 0x57, 0xd0, 0x68, 0x2d, 0xd3, 0x90, 0xad, 0xb7, 0xa2,
	0x25, 0x5b, 0xb2, 0xfd, 0xa8, 0x67, 0xc5, 0x49, 0xa9, 0xfc, 0xaa, 0x5e, 0xc2, 0x5d, 0x42, 0x4b,
	0x3e, 0xf3, 0xcf, 0x0a, 0xc4, 0xca, 0x56, 0x2e, 0x2c, 0x2c, 0x39, 0xa4, 0x90, 0x80, 0x00, 0x1e,
	0x00, 0xca, 0xbb, 0x97, 0x54, 0xaa, 0xf2, 0x05, 0x92, 0x5b, 0x92, 0x5b, 0xae, 0xa9, 0xdc, 0xf2,
	0xef, 0x03, 0xe4, 0x9c, 0x43, 0x4e, 0x39, 0x24, 0xf9, 0x0a, 0xa9, 0x54, 0xe5, 0x96, 0xca, 0x21,
	0xd5, 0x33, 0x03, 0x10, 0x20, 0x41, 0x5b, 0xf5, 0xa4, 0x4a, 0xbd, 0x13, 0x67, 0xba, 0x7f, 0xd3,
	0xd3, 0x33, 0xd3, 0xdd, 0xd3, 0xd3, 0x20, 0x1c, 0x04, 0x97, 0x8f, 0xa3, 0xeb, 0xc8, 0x6d, 0x05,
	0xa1, 0x1f, 0xfb, 0x44, 0xc6, 0x76, 0xf3, 0xdf, 0x4b, 0x70, 0x30, 0xf6, 0x57, 0xe1, 0x94, 0x9e,
	0xfa, 0x5e, 0x4c, 0xaf, 0x62, 0x42, 0x40, 0x9e, 0x3b, 0x2e, 0x6d, 0x94, 0x8e, 0x4b, 0x0f, 0x6b,
	0x26, 0x6b, 0x93, 0x27, 0xa0, 0x44, 0xb1, 0x1d, 0xc6, 0x8d, 0xf2, 0x71, 0xe9, 0x61, 0xfd, 0xc9,
	0x87, 0x2d, 0x26, 0x27, 0x37, 0xae, 0xd5, 0xf7, 0xa7, 0x76, 0xec, 0xf8, 0x9e, 0xc9, 0xa1, 0xa4,
	0x05, 0x12, 0xf5, 0x66, 0x0d, 0xe9, 0x0d, 0x46, 0x20, 0x90, 0xdc, 0x03, 0x65, 0x46, 0xdd, 0xd8,
	0x6e, 0xa8, 0xc7, 0xa5, 0x87, 0x87, 0x4f, 0xea, 0x7c, 0x44, 0x07, 0x49, 0x26, 0xe7, 0xe8, 0x3f,
	0x83, 0x6a, 0x32, 0x06, 0xd5, 0x74, 0x1d, 0x8f, 0xab, 0xa9, 0x98, 0xac, 0x4d, 0x34, 0x90, 0xa6,
	0xbe, 0xcb, 0x94, 0x54, 0x4c, 0x6c, 0x36, 0xff, 0xb1, 0x0c, 0xea, 0xc0, 0x9f, 0xad, 0x5c, 0x4a,
	0x3e, 0x03, 0xd9, 0x0e, 0x82, 0xa8, 0x51, 0x3e, 0x96, 0x1e, 0xd6, 0x9f, 0xdc, 0xe6, 0xe2, 0x39,
	0xaf, 0xd5, 0x0e, 0x82, 0xc8, 0xf0, 0xe2, 0xf0, 0xda, 0x64, 0x18, 0xf2, 0x53, 0x50, 0xe2, 0xeb,
	0x80, 0x46, 0x0d, 0x89, 0x81, 0xdf, 0xcf, 0x81, 0x2d, 0xe4, 0x70, 0x34, 0x47, 0x91, 0xaf, 0xe1,
	0x30, 0x62, 0x2b, 0x9b, 0x4c, 0xf9, 0xd2, 0x1a, 0x53, 0xb6, 0xea, 0x5b, 0x05, 0xab, 0x36, 0x0f,
	0xa2, 0x6c, 0x57, 0xff, 0x25, 0xd4, 0xd2, 0xd9, 0x71, 0x01, 0x7f, 0x48, 0xaf, 0xc5, 0xd6, 0x63,
	0x93, 0x7c, 0x0a, 0xca, 0x6b, 0xdb, 0x5d, 0x51, 0xb1, 0xf3, 0x37, 0xb9, 0xc4, 0x76, 0x10, 0xb8,
	0x4e, 0xb2, 0xdd, 0x8c, 0xff, 0x75, 0xf9, 0x69, 0x49, 0xef, 0x00, 0xac, 0x95, 0x2b, 0x10, 0x76,
	0x9c, 0x17, 0x06, 0x5c, 0x18, 0x0e, 0xc9, 0x48, 0xf9, 0xa5, 0x5c, 0x2d, 0x69, 0xe5, 0xe6, 0xbf,
	0x96, 0xa0, 0xd6, 0x8e, 0xe3, 0xd0, 0xb9, 0x5c, 0xc5, 0x94, 0x1c, 0x42, 0x29, 0x6a, 0xc8, 0x28,
	0xa9, 0xbb, 0x67, 0x96, 0x22, 0xec, 0x3b, 0x0d, 0xe5, 0xb8, 0xf4, 0x50, 0xc2, 0xbe, 0x83, 0x7d,
	0x8f, 0x1d, 0x5c, 0x09, 0xfb, 0x1e, 0x79, 0x00, 0x25, 0xbb, 0x51, 0x61, 0xb3, 0xbc, 0x27, 0x54,
	0x4e, 0x64, 0xb5, 0xda, 0x61, 0x68, 0x5f, 0x23, 0xcc, 0x7e, 0xab, 0x8d, 0xfb, 0x0c, 0x14, 0x26,
	0x89, 0xdc, 0x03, 0x89, 0xba, 0x71, 0xa3, 0xc4, 0x8e, 0xea, 0xc6, 0xc6, 0x6c, 0x26, 0xf2, 0x4e,
	0xea, 0x50, 0xb3, 0x13, 0x4a, 0xf3, 0x23, 0xa8, 0xb4, 0x83, 0x60, 0x68, 0x2f, 0x29, 0x1a, 0x51,
	0x80, 0x66, 0x8d, 0x63, 0x6b, 0x26, 0x6b, 0x37, 0xff, 0x41, 0x85, 0x7a, 0x66, 0x7f, 0xc9, 0x3d,
	0x90, 0x3d, 0x7b, 0xc9, 0x0d, 0xad, 0xfe, 0xe4, 0x20, 0x3d, 0x00, 0x14, 0x60, 0x32, 0x16, 0xb9,
	0x03, 0x35, 0xd7, 0xf7, 0x16, 0x13, 0x86, 0x2b, 0xb3, 0xfd, 0xae, 0x22, 0x81, 0xcd, 0xf1, 0x21,
	0xd4, 0x66, 0xfe, 0x34, 0x8a, 0x43, 0xc7, 0x5b, 0x30, 0x6f, 0xa8, 0x99, 0x6b, 0x02, 0x7a, 0x16,
	0x6a, 0x86, 0x9b, 0x2b, 0xad, 0xfd, 0x24, 0x33, 0x3f, 0x5b, 0x4a, 0x62, 0x6e, 0x0c, 0x4a, 0x7e,
	0x01, 0x35, 0xea, 0xcd, 0x02, 0xdf, 0xf1, 0xe2, 0xa8, 0xa1, 0xb0, 0x71, 0xc7, 0xdb, 0xe3, 0x8c,
	0x04, 0xc2, 0xc7, 0xae, 0x87, 0xe0, 0x9c, 0xdc, 0xba, 0xd5, 0x5d, 0x73, 0x6e, 0x9b, 0xf8, 0x13,
	0x50, 0x5e, 0x3b, 0xf4, 0xfb, 0xa8, 0x01, 0xbb, 0xc6, 0xbc, 0x40, 0xb6, 0x18, 0xc3, 0xa0, 0xe4,
	0x11, 0xa8, 0x4b, 0xe7, 0xca, 0xf1, 0x9e, 0x34, 0xaa, 0x6c, 0x50, 0x81, 0xf1, 0x0a, 0x00, 0xf9,
	0x1c, 0x2a, 0xdf, 0x87, 0x76, 0x10, 0xd0, 0x59, 0xa3, 0xb6, 0xcb, 0xd0, 0x13, 0xc4, 0xdb, 0x58,
	0x0d, 0xf9, 0x0a, 0x0e, 0x3b, 0xa3, 0xe1, 0xc8, 0xba, 0x18, 0x1b, 0x13, 0x36, 0x77, 0xa3, 0xc2,
	0x74, 0xdb, 0x38, 0xd7, 0x83, 0x04, 0x34, 0x40, 0x8c, 0xde, 0x03, 0x58, 0x1f, 0x43, 0x81, 0x63,
	0x3d, 0xc8, 0x3b, 0xd6, 0x96, 0x11, 0x66, 0x7c, 0xb4, 0x0f, 0x87, 0xf9, 0x93, 0x29, 0x10, 0x77,
	0x3f, 0x2f, 0xee, 0x90, 0x8b, 0x4b, 0x86, 0xbd, 0x73, 0x8f, 0x47, 0x29, 0xeb, 0xd3, 0x7b, 0x63,
	0x29, 0x38, 0x24, 0x23, 0xa5, 0xf9, 0xb7, 0x15, 0xa8, 0x26, 0x3a, 0xa2, 0x67, 0xa5, 0x5e, 0x53,
	0x7b, 0x7b, 0x37, 0x79, 0x9c, 0x77, 0x93, 0x0f, 0xf2, 0x3b, 0x52, 0xe0, 0x23, 0x78, 0x8b, 0xb9,
	0xf6, 0x82, 0x99, 0x2b, 0xde, 0x62, 0xae, 0xbd, 0x20, 0x0f, 0x40, 0xe5, 0xc6, 0xc0, 0x22, 0xd7,
	0xd6, 0x99, 0x0b, 0x26, 0xaa, 0xe9, 0x44, 0x93, 0x60, 0x75, 0x19, 0xad, 0x2e, 0x59, 0x4c, 0xab,
	0x9a, 0x55, 0x27, 0x3a, 0x67, 0x7d, 0xbc, 0xa5, 0x02, 0x3b, 0xb4, 0x97, 0x8d, 0x1a, 0x53, 0x44,
	0xdc, 0x52, 0xe7, 0x48, 0x32, 0x39, 0x87, 0x7c, 0x0c, 0x72, 0x14, 0x2f, 0x63, 0x61, 0x58, 0xc2,
	0x16, 0xc6, 0xb1, 0x1d, 0xd3, 0x25, 0xf5, 0x62, 0x93, 0x31, 0xc9, 0xd7, 0x50, 0x0f, 0x69, 0x14,
	0x4f, 0xd8, 0x90, 0xa8, 0x51, 0x65, 0x0a, 0x6d, 0x2e, 0xcb, 0xa4, 0x51, 0xcc, 0x44, 0x47, 0x26,
	0x84, 0x69, 0xfb, 0xad, 0xa2, 0xe6, 0x3b, 0xb4, 0xe4, 0x7f, 0x92, 0x00, 0xd6, 0x1a, 0x92, 0xa7,
	0xa0, 0x2e, 0x69, 0xfc, 0xca, 0x9f, 0x31, 0x71, 0x87, 0x49, 0x48, 0x2a, 0x58, 0x4c, 0x6b, 0xc0,
	0x70, 0xa6, 0xc0, 0xf3, 0x28, 0x1c, 0xbf, 0x12, 0x26, 0xc1, 0xda, 0xc4, 0x80, 0xfa, 0xaf, 0x56,
	0x34, 0xbc, 0xe6, 0x1b, 0x24, 0xee, 0xe1, 0xfb, 0xbb, 0x45, 0x3e, 0x47, 0x30, 0x3f, 0x06, 0xf8,
	0x55, 0xda, 0xd6, 0xff, 0xba, 0x04, 0xb0, 0x66, 0x15, 0x5a, 0xe5, 0x5d, 0x90, 0x31, 0xc4, 0x15,
	0x78, 0x08, 0xa3, 0xe3, 0x1e, 0xb9, 0xfe, 0x94, 0x5d, 0x7e, 0x55, 0x13, 0x9b, 0xe4, 0x41, 0x26,
	0x86, 0x24, 0xea, 0xa1, 0xbc, 0x34, 0x68, 0xb0, 0xc9, 0x9a, 0x8f, 0x41, 0xea, 0xfb, 0x53, 0x52,
	0x03, 0x65, 0x38, 0x9a, 0xf4, 0x3c, 0x6d, 0x0f, 0x9b, 0xcf, 0x2f, 0x0c, 0xf3, 0xa5, 0x56, 0x22,
	0x55, 0x90, 0xcf, 0xdb, 0x56, 0x57, 0x2b, 0x63, 0xeb, 0x64, 0xd4, 0x79, 0xa9, 0x49, 0xcd, 0x08,
	0x54, 0xbe, 0x33, 0xe4, 0x00, 0x6a, 0xc3, 0xd1, 0x84, 0x77, 0xb4, 0x3d, 0x52, 0x01, 0xe9, 0xcc,
	0xb0, 0xb4, 0x12, 0x36, 0xce, 0x2f, 0x2c, 0x4d, 0x62, 0xc3, 0x47, 0x63, 0x4b, 0x93, 0x09, 0x80,
	0xda, 0x31, 0xfa, 0x86, 0x65, 0x68, 0x0a, 0xca, 0x3f, 0x6f, 0x5b, 0xa7, 0x5d, 0x4d, 0x25, 0x47,
	0xa0, 0xa5, 0x3a, 0x8e, 0xce, 0xad, 0xde, 0x68, 0x38, 0xd6, 0x2a, 0xe4, 0x26, 0xa4, 0x3a, 0x4e,
	0xba, 0x46, 0xbb, 0xa3, 0x95, 0x9b, 0x3f, 0x07, 0xe5, 0xd7, 0xde, 0x9b, 0xe6, 0x9f, 0xc9, 0x50,
	0x4b, 0x2d, 0x9b, 0x7c, 0x02, 0xaa, 0x3d, 0xc5, 0x50, 0x2d, 0xee, 0xca, 0x7d, 0x61, 0x3c, 0x8c,
	0xd6, 0xdd, 0x33, 0x05, 0x97, 0x1c, 0x83, 0x3c, 0xb5, 0x5d, 0x37, 0x2f, 0xf5, 0xd4, 0x76, 0xdd,
	0xee, 0x9e, 0xc9, 0x38, 0x0c, 0xe1, 0xa7, 0xc9, 0x63, 0x82, 0xf0, 0xbd, 0x19, 0x43, 0xf8, 0xde,
	0x0c, 0x11, 0xae, 0xef, 0x07, 0xec, 0x58, 0x52, 0x44, 0xdf, 0xf7, 0x03, 0x44, 0x20, 0x87, 0xdc,
	0x07, 0x15, 0x7f, 0x27, 0x9e, 0xf0, 0xf6, 0xfa, 0x1a, 0x33, 0xec, 0xee, 0x99, 0x0a, 0x32, 0x87,
	0xe4, 0x11, 0x54, 0xe6, 0x7e, 0x48, 0xed, 0xe9, 0x2b, 0x71, 0xf1, 0x88, 0xa0, 0xf0, 0x8c, 0x13,
	0xbb, 0x7b, 0x66, 0xc2, 0x27, 0x1f, 0x81, 0x64, 0xbb, 0x31, 0x8b, 0x08, 0xf5, 0x27, 0x35, 0xb1,
	0x36, 0x37, 0xee, 0xee, 0x99, 0x48, 0x27, 0x1f, 0x83, 0xb2, 0x08, 0xfd, 0x55, 0x20, 0xd2, 0x1e,
	0x31, 0xdd, 0x19, 0x92, 0x70, 0x3a, 0xc6, 0x23, 0xc7, 0x20, 0x85, 0x34, 0x16, 0xee, 0x2e, 0xf6,
	0xc7, 0xa4, 0xf1, 0x2a, 0xc4, 0xfd, 0x41, 0x16, 0xf9, 0x59, 0x12, 0xe9, 0xf8, 0x45, 0xab, 0x6f,
	0x84, 0x8f, 0x82, 0x50, 0xf7, 0x9b, 0x11, 0x0e, 0x4e, 0x54, 0x1e, 0xf6, 0x9a, 0xc7, 0xa0, 0xf2,
	0x13, 0x27, 0xb7, 0x53, 0x7b, 0xe0, 0x9e, 0x2d, 0x7a, 0xcd, 0xff, 0x2e, 0x83, 0x8c, 0xc7, 0x8d,
	0x01, 0x39, 0xb6, 0xc3, 0x05, 0x8d, 0x8b, 0x93, 0x2b, 0xc1, 0x24, 0x3a, 0x54, 0x93, 0xe4, 0x25,
	0xb9, 0x36, 0x92, 0x3e, 0x6e, 0xa8, 0x1d, 0x2e, 0xc4, 0xb5, 0x70, 0xb8, 0x36, 0xa5, 0x56, 0x3b,
	0x5c, 0x98, 0xc8, 0x22, 0x9d, 0x8c, 0xb7, 0xf2, 0x9d, 0xe5, 0xc1, 0xe4, 0xa3, 0x0c, 0x38, 0x01,
	0x64, 0x36, 0xf7, 0x20, 0x47, 0xd3, 0x97, 0x20, 0xb5, 0xc3, 0x05, 0x86, 0x7f, 0xbe, 0x1f, 0x52,
	0xf6, 0x90, 0x5f, 0x20, 0x49, 0xec, 0x45, 0xa1, 0x1f, 0x3d, 0xce, 0xf8, 0xdd, 0x0e, 0x87, 0xda,
	0x4f, 0x00, 0xd8, 0xd3, 0x9f, 0x03, 0xd9, 0xd6, 0xe9, 0xad, 0xce, 0xa7, 0xf9, 0xbb, 0x20, 0xa3,
	0x07, 0xa1, 0x7e, 0x31, 0x8d, 0xe2, 0x44, 0x3f, 0x6c, 0xa7, 0x57, 0x56, 0xf9, 0x07, 0xae, 0xac,
	0xe6, 0x5f, 0x94, 0x40, 0x46, 0xef, 0x41, 0xf4, 0xd2, 0x9f, 0x51, 0x11, 0xe7, 0x6f, 0xac, 0xfd,
	0x0a, 0x5f, 0x48, 0xd4, 0x64, 0x4c, 0xbc, 0xcf, 0xa7, 0xa1, 0x13, 0xd3, 0x70, 0x7d, 0xfe, 0x6b,
	0x42, 0x3a, 0xa1, 0xf4, 0x43, 0x13, 0x3e, 0x02, 0x19, 0x05, 0x92, 0x3a, 0x54, 0x30, 0x1a, 0xfa,
	0x33, 0xca, 0x63, 0xe8, 0xb7, 0xdd, 0x5e, 0xdf, 0xd0, 0x4a, 0xd8, 0xbc, 0x18, 0x5a, 0xbd, 0xbe,
	0x56, 0x6e, 0x9e, 0x80, 0xc2, 0x1c, 0x9b, 0x1c, 0x81, 0x32, 0xf5, 0x57, 0x5e, 0x2c, 0xde, 0x85,
	0xbc, 0xf3, 0x66, 0xd3, 0x0d, 0xa1, 0x22, 0xbc, 0x9e, 0xdc, 0x05, 0x98, 0xfa, 0xae, 0x4b, 0xd7,
	0xd1, 0xac, 0x66, 0x66, 0x28, 0x6f, 0x26, 0x2f, 0x04, 0xa9, 0xed, 0xc6, 0xe4, 0x21, 0xa8, 0xd3,
	0x57, 0xbe, 0x33, 0xa5, 0xe2, 0x85, 0xa2, 0xa5, 0x91, 0xa3, 0x75, 0xca, 0xe8, 0xa6, 0xe0, 0xeb,
	0x6d, 0x50, 0x39, 0x05, 0xcf, 0x88, 0xc5, 0x3f, 0x71, 0x46, 0x2c, 0xe2, 0xbd, 0xd1, 0x19, 0x9d,
	0x80, 0xc2, 0x22, 0x0e, 0xee, 0x43, 0xec, 0xc4, 0xe9, 0x33, 0x9e, 0x77, 0xde, 0x4c, 0x46, 0x13,
	0x54, 0x1e, 0x92, 0x48, 0x03, 0x2a, 0x81, 0x7d, 0xed, 0xfa, 0x76, 0xa2, 0x49, 0xd2, 0x6d, 0xfe,
	0x27, 0x01, 0x19, 0x0d, 0x95, 0x7c, 0x05, 0xb5, 0x20, 0x74, 0x96, 0x4e, 0xec, 0xbc, 0x4e, 0x0c,
	0xe2, 0x68, 0x6d, 0xd5, 0xad, 0xf3, 0x84, 0xd7, 0xdd, 0x33, 0xd7, 0x40, 0xf2, 0x00, 0x64, 0xea,
	0xad, 0x96, 0x79, 0xab, 0x65, 0x03, 0x0c, 0x6f, 0xb5, 0xc4, 0x10, 0x8e, 0x6c, 0xf2, 0x10, 0x94,
	0x78, 0x15, 0xb8, 0x89, 0xb7, 0x69, 0x19, 0x9c, 0x85, 0x74, 0x8c, 0xab, 0x0c, 0x80, 0x02, 0x5d,
	0x27, 0x8a, 0xc5, 0x75, 0x90, 0x15, 0xd8, 0x77, 0xa2, 0x98, 0xdd, 0x09, 0x4e, 0x14, 0x93, 0x26,
	0x48, 0x4b, 0x3b, 0x10, 0x17, 0xc2, 0x61, 0x06, 0x35, 0xb0, 0x31, 0x48, 0x23, 0x13, 0x5f, 0x2d,
	0xbe, 0x47, 0x27, 0xfe, 0x5c, 0x44, 0xfa, 0xec, 0xac, 0x23, 0x8f, 0x8e, 0xe6, 0x38, 0xab, 0x8f,
	0x0d, 0xf2, 0x25, 0x54, 0x43, 0xea, 0xb2, 0xd7, 0x89, 0x88, 0xfa, 0xb7, 0x32, 0x60, 0x53, 0xb0,
	0xba, 0x7b, 0x66, 0x0a, 0x23, 0x5f, 0x40, 0x15, 0x03, 0xc0, 0x24, 0xa4, 0x73, 0x71, 0xe1, 0x24,
	0xa7, 0x30, 0xf5, 0x03, 0x3a, 0x33, 0x29, 0x8a, 0xaf, 0x20, 0xc4, 0xa4, 0x73, 0x72, 0x17, 0xa4,
	0x88, 0xc6, 0x8d, 0x83, 0xcd, 0x68, 0x81, 0xba, 0x46, 0x34, 0x26, 0x5f, 0x40, 0xc5, 0xf3, 0x79,
	0x44, 0x39, 0xcc, 0x3e, 0x9b, 0xd8, 0xfc, 0x43, 0x5f, 0x40, 0x55, 0x8f, 0xb5, 0xc8, 0xe7, 0xc9,
	0xd5, 0xc2, 0x9f, 0x63, 0xef, 0x65, 0xb0, 0xdb, 0xb7, 0xca, 0x6f, 0xa3, 0x0b, 0x78, 0x51, 0x1c,
	0xda, 0x18, 0x76, 0x61, 0x6b, 0xc4, 0x69, 0xca, 0x34, 0x33, 0xc0, 0x7c, 0x1a, 0x5f, 0xdf, 0x4c,
	0xe3, 0x35, 0x90, 0xfc, 0x20, 0x6e, 0xec, 0xf3, 0x5c, 0xca, 0x0f, 0xe2, 0xdf, 0x94, 0x5c, 0x36,
	0x04, 0x19, 0x2d, 0x0f, 0x6f, 0x5f, 0x27, 0xa6, 0xcb, 0x48, 0xf8, 0xaa, 0xbe, 0x61, 0x99, 0xad,
	0x1e, 0x32, 0xc5, 0x3e, 0x31, 0xa0, 0xfe, 0x14, 0x60, 0x4d, 0x2c, 0x50, 0xe2, 0x28, 0xab, 0x84,
	0x94, 0x9d, 0xf3, 0x2f, 0x25, 0x50, 0x98, 0x19, 0x93, 0x9f, 0xf3, 0xf2, 0xc4, 0x64, 0x46, 0xe7,
	0xc9, 0xcc, 0x77, 0x37, 0x6d, 0x9d, 0xe9, 0xdd, 0xa1, 0x73, 0x31, 0x7b, 0xd5, 0x16, 0x5d, 0x72,
	0x06, 0x07, 0xcf, 0x2e, 0xac, 0x0b, 0xd3, 0x98, 0xcc, 0x1d, 0xea, 0xce, 0x92, 0x02, 0x57, 0x73,
	0x4b, 0x00, 0x47, 0x3d, 0x63, 0x20, 0x2e, 0x64, 0x3f, 0x4b, 0xd2, 0xcf, 0xe0, 0x20, 0x37, 0xc7,
	0xaf, 0xfd, 0x9c, 0xfc, 0x16, 0x6e, 0x6e, 0xcd, 0x55, 0x20, 0xec, 0x8b, 0xbc, 0xb0, 0xdb, 0xdb,
	0x0a, 0xe3, 0xf0, 0xac, 0x60, 0x07, 0x14, 0x46, 0x4b, 0xf3, 0xd2, 0xd2, 0x8e, 0x9c, 0xfd, 0x27,
	0x50, 0x5f, 0x3a, 0xde, 0x24, 0xa4, 0x01, 0xb5, 0xe3, 0x48, 0x6c, 0x3d, 0x2c, 0x1d, 0xcf, 0xe4,
	0x14, 0x06, 0xb0, 0xaf, 0x52, 0x80, 0x24, 0x00, 0xf6, 0x95, 0x00, 0xe8, 0x9f, 0x80, 0x8c, 0x91,
	0xe3, 0xc7, 0x66, 0xd2, 0x0d, 0x90, 0x06, 0x76, 0x40, 0x3e, 0x5c, 0xaf, 0x2e, 0x8f, 0x7a, 0xb3,
	0x6d, 0xd3, 0x3f, 0x05, 0x85, 0xc5, 0x96, 0xcc, 0x7c, 0x52, 0xe1, 0x7c, 0x7f, 0x57, 0x86, 0x6a,
	0x12, 0x58, 0xc8, 0x2f, 0xb6, 0xed, 0xe6, 0x5e, 0x41, 0x00, 0xda, 0x69, 0x3a, 0x4f, 0xa1, 0x8e,
	0x31, 0xd9, 0x0e, 0xaf, 0x27, 0xa8, 0x3d, 0xd7, 0xee, 0xfd, 0x22, 0x09, 0xdf, 0xd0, 0x6b, 0x13,
	0x04, 0xf6, 0x1b, 0x7a, 0x4d, 0x1e, 0xf1, 0xf5, 0xe6, 0xca, 0xa3, 0xdb, 0x23, 0xd8, 0xe2, 0x6f,
	0x83, 0xea, 0x78, 0x7f, 0x40, 0xa7, 0x31, 0x4b, 0xd2, 0x6a, 0xa6, 0xe8, 0xbd, 0x3b, 0x73, 0x6b,
	0x82, 0x84, 0x2a, 0xdd, 0x11, 0x9b, 0x21, 0x92, 0x2f, 0x9c, 0x8a, 0xad, 0x14, 0x53, 0x49, 0x7d,
	0xca, 0xef, 0x76, 0x67, 0xe1, 0x91, 0x9f, 0x80, 0x64, 0x07, 0x41, 0x71, 0xc6, 0x89, 0x1c, 0x4c,
	0x37, 0xd3, 0xa8, 0x2e, 0xd2, 0xcd, 0x34, 0x7c, 0xe7, 0x26, 0xe1, 0xaf, 0xbe, 0xf5, 0x24, 0x7f,
	0x2e, 0x01, 0xac, 0xc3, 0x22, 0xf9, 0x12, 0x94, 0xd0, 0xf6, 0x16, 0x89, 0xed, 0xdc, 0x29, 0x0c,
	0x9e, 0x2d, 0x13, 0x21, 0x26, 0x47, 0x92, 0xaf, 0x40, 0x75, 0xa9, 0xb7, 0x10, 0x6f, 0xe1, 0xb4,
	0xcc, 0xb6, 0x39, 0xa6, 0xcf, 0x30, 0xa6, 0xc0, 0x92, 0xdf, 0x03, 0x08, 0x69, 0xe4, 0xbb, 0x2b,
	0xa6, 0x32, 0xbf, 0x2b, 0x8f, 0x77, 0xcc, 0x96, 0xe2, 0xcc, 0xcc, 0x18, 0x8c, 0xda, 0x41, 0x48,
	0xa7, 0x4e, 0x84, 0x02, 0x64, 0x96, 0x39, 0xad, 0x09, 0x18, 0xc2, 0xa2, 0xa9, 0xed, 0xf2, 0xb2,
	0x89, 0x62, 0xf2, 0x8e, 0x6e, 0x80, 0xc2, 0x74, 0xc7, 0x77, 0xd1, 0xd2, 0x49, 0xde, 0x7c, 0xb9,
	0x8c, 0x18, 0xe9, 0x8c, 0x6d, 0x5f, 0x89, 0x05, 0x6d, 0xb0, 0xed, 0x2b, 0xfd, 0x0b, 0x50, 0xf9,
	0x72, 0xf0, 0xfc, 0x13, 0x39, 0x12, 0x1f, 0xaa, 0xad, 0x87, 0x4a, 0x1c, 0xfd, 0x3b, 0xac, 0xe4,
	0x90, 0xa8, 0x4d, 0x40, 0xbe, 0xb4, 0xa3, 0xf4, 0x1b, 0x00, 0xb6, 0x51, 0x59, 0xc7, 0x9b, 0xd1,
	0x2b, 0xf1, 0x15, 0x80, 0x77, 0xf4, 0x2a, 0xa8, 0xfc, 0x3a, 0x6c, 0xfe, 0x4d, 0x09, 0x6a, 0x69,
	0x56, 0x42, 0x34, 0xd8, 0x1f, 0x8e, 0x26, 0x69, 0x9f, 0xe7, 0x97, 0xc6, 0xe0, 0xdc, 0x7a, 0xc9,
	0x5f, 0xdb, 0xed, 0xe1, 0xcb, 0xe4, 0x89, 0x3e, 0xea, 0x6b, 0x12, 0x92, 0x7a, 0x43, 0x7c, 0x76,
	0xd7, 0x40, 0x79, 0xd6, 0x1f, 0xb5, 0x2d, 0x4d, 0xc1, 0xf4, 0xb4, 0x63, 0x9c, 0xf6, 0x06, 0xed,
	0xbe, 0xb6, 0x8f, 0xcf, 0xf1, 0xb1, 0x65, 0xf6, 0x86, 0x67, 0x9a, 0x8a, 0x98, 0x93, 0x97, 0x96,
	0x81, 0x0f, 0xef, 0x7d, 0xa8, 0x72, 0xf2, 0xe4, 0xa9, 0x56, 0x45, 0x79, 0x9d, 0xb6, 0x65, 0x68,
	0x35, 0xa4, 0x63, 0xcb, 0xea, 0x0d, 0x0c, 0x0d, 0x50, 0xfa, 0x77, 0x83, 0xbe, 0x56, 0x47, 0xc0,
	0xc5, 0x45, 0xaf, 0xa3, 0x1d, 0xe0, 0xab, 0x8a, 0xbd, 0xb4, 0xff, 0xad, 0x0c, 0xf2, 0x0b, 0x87,
	0x7e, 0xbf, 0x2e, 0x40, 0x95, 0x76, 0x16, 0xa0, 0x1e, 0xa0, 0x01, 0xc7, 0xbb, 0x1e, 0x1a, 0x95,
	0x90, 0xc6, 0x2c, 0x1d, 0xb8, 0x0b, 0x32, 0xbd, 0x0a, 0xc2, 0xfc, 0x23, 0xdb, 0xb8, 0x0a, 0x42,
	0x93, 0xd1, 0x31, 0x5d, 0xe0, 0x25, 0x5f, 0x39, 0x7b, 0xf9, 0xa3, 0x12, 0x05, 0xb5, 0xde, 0xe2,
	0xdc, 0x82, 0x81, 0xb7, 0x72, 0x8b, 0x77, 0x53, 0x6f, 0x7c, 0x87, 0xd7, 0x7f, 0xf3, 0xbf, 0x8e,
	0x40, 0xc6, 0x95, 0x93, 0xa3, 0xec, 0xe3, 0x0d, 0xd3, 0x46, 0xf6, 0x7c, 0xfb, 0x14, 0x2a, 0x2e,
	0xbe, 0x5c, 0x6c, 0xb7, 0xc0, 0x8c, 0x31, 0x5f, 0x13, 0x5c, 0xf2, 0x18, 0xaa, 0x0b, 0x1a, 0xb3,
	0x67, 0xa6, 0xd8, 0x56, 0xb2, 0xde, 0xd6, 0xd6, 0x19, 0x8d, 0x71, 0x76, 0x1c, 0xb0, 0xe0, 0x4d,
	0x4c, 0x9f, 0xe3, 0xd0, 0xf6, 0xa2, 0xb9, 0x1f, 0x2e, 0x45, 0xf2, 0x7a, 0x94, 0x19, 0x61, 0x25,
	0x3c, 0x4c, 0x9f, 0x53, 0x20, 0xf9, 0x1c, 0x54, 0x67, 0x4e, 0xdd, 0x28, 0x29, 0x64, 0xde, 0xcc,
	0x0c, 0xe9, 0xcd, 0x0d, 0x37, 0x62, 0x59, 0x1f, 0x87, 0x60, 0x6a, 0xcc, 0xaa, 0x2d, 0x6a, 0x76,
	0x17, 0x18, 0x34, 0x57, 0x72, 0xf9, 0x1c, 0xd4, 0x95, 0xc7, 0xec, 0xa1, 0xb2, 0x25, 0xf3, 0xc2,
	0xc3, 0x1f, 0x94, 0xc9, 0x21, 0xa4, 0x05, 0x95, 0x4b, 0x87, 0xa3, 0xab, 0x5b, 0xcb, 0x3c, 0x71,
	0x12, 0x78, 0x02, 0x42, 0x7c, 0x48, 0x5d, 0x86, 0xdf, 0xdf, 0xc2, 0x9b, 0xd4, 0x4d, 0xf0, 0x02,
	0x84, 0x89, 0xb5, 0x67, 0xbf, 0x76, 0x16, 0x76, 0x4c, 0x45, 0x96, 0x7c, 0x2b, 0x33, 0x60, 0x28,
	0x58, 0x98, 0x58, 0x27, 0xb0, 0xf4, 0x05, 0x00, 0x5b, 0xcb, 0xcc, 0xbd, 0x00, 0x3e, 0xe6, 0x19,
	0x75, 0x7d, 0x17, 0x8a, 0xa5, 0xd5, 0xe9, 0xbb, 0xe3, 0x30, 0xfb, 0x02, 0xe0, 0x27, 0x92, 0x7f,
	0x77, 0x24, 0xd7, 0xf5, 0xc1, 0x8e, 0xf4, 0x60, 0x09, 0x15, 0x71, 0xea, 0x98, 0x22, 0xd8, 0xe1,
	0x22, 0x9f, 0x22, 0x30, 0x6f, 0x63, 0x55, 0x0a, 0x02, 0x32, 0xb3, 0x1a, 0x51, 0x03, 0xc5, 0x36,
	0x5e, 0x44, 0xde, 0xca, 0x75, 0x23, 0x7b, 0xce, 0xef, 0x9a, 0xaa, 0x99, 0xf6, 0x59, 0x4c, 0xa6,
	0xb1, 0x3f, 0x17, 0x75, 0x49, 0xde, 0xd1, 0xff, 0xb8, 0x04, 0xd5, 0x64, 0x77, 0xfe, 0x3f, 0x26,
	0x44, 0xaf, 0x7b, 0xed, 0xd8, 0xcc, 0x0c, 0x6b, 0x26, 0x36, 0xb3, 0x89, 0x13, 0x3b, 0xef, 0x5c,
	0x22, 0xb3, 0x8e, 0x2e, 0xfa, 0x3f, 0x4b, 0x50, 0x4b, 0xcd, 0xfb, 0x47, 0x74, 0xd5, 0xa1, 0x1a,
	0xe1, 0xf3, 0xe8, 0xb5, 0x9d, 0xe8, 0x9b, 0xf6, 0xc9, 0x4f, 0x73, 0x4f, 0xf1, 0x0f, 0x8a, 0x9c,
	0xa7, 0x35, 0x8e, 0x97, 0x49, 0xdd, 0x1d, 0x2f, 0x70, 0xd7, 0x9d, 0x24, 0x1f, 0x13, 0xd8, 0x1a,
	0x6d, 0xd7, 0x65, 0x81, 0x84, 0xdc, 0x83, 0x7d, 0x7a, 0x35, 0xa5, 0x41, 0x2c, 0xf8, 0x0a, 0xcb,
	0x22, 0xea, 0x9c, 0xc6, 0x21, 0xd9, 0x2d, 0x52, 0xf3, 0x5b, 0xa4, 0xff, 0x4f, 0x09, 0xe4, 0x31,
	0x2f, 0xee, 0xab, 0x76, 0x14, 0x39, 0x8b, 0xe4, 0x52, 0x3c, 0xde, 0xa9, 0x55, 0xab, 0xcd, 0x70,
	0xac, 0x38, 0xca, 0x5a, 0xe4, 0x2b, 0x90, 0x5c, 0x9a, 0x7c, 0x68, 0x7f, 0x93, 0x81, 0x08, 0x27,
	0xf7, 0x33, 0x49, 0xd6, 0xc6, 0x16, 0xb2, 0x50, 0xc0, 0x53, 0x2e, 0x13, 0x54, 0x3e, 0x6c, 0x57,
	0xb1, 0x97, 0x9d, 0x58, 0x79, 0xc7, 0x7d, 0x70, 0x04, 0x4a, 0x6c, 0x5f, 0xba, 0x89, 0x69, 0xf0,
	0x4e, 0x52, 0xf6, 0xd3, 0xff, 0xb4, 0x04, 0x2a, 0x8f, 0x3d, 0x28, 0x28, 0xad, 0x5e, 0x6c, 0x08,
	0x12, 0x95, 0x8c, 0x8a, 0x33, 0x9f, 0xc4, 0xe1, 0x66, 0x80, 0x67, 0x10, 0xd5, 0x99, 0x5b, 0xe1,
	0x0a, 0xfd, 0xb9, 0xea, 0xcc, 0x27, 0x73, 0x1b, 0xa3, 0xdc, 0xf6, 0x0d, 0x55, 0x71, 0xe6, 0xcf,
	0x90, 0x95, 0x3b, 0x0f, 0x79, 0xe3, 0x3c, 0x9e, 0x8a, 0x32, 0x23, 0x01, 0x79, 0xbe, 0xf2, 0xa6,
	0xe9, 0x3f, 0x1a, 0x56, 0xde, 0x34, 0x31, 0xb8, 0xf2, 0x96, 0x75, 0x22, 0x59, 0xff, 0xab, 0x12,
	0xa8, 0x3c, 0xe8, 0x91, 0xfb, 0x50, 0xf6, 0x83, 0x7c, 0x65, 0x23, 0x13, 0x13, 0x5b, 0xa3, 0xc0,
	0x2c, 0xfb, 0xc1, 0x5a, 0x5c, 0x91, 0xfd, 0x36, 0xc7, 0x50, 0x1e, 0x05, 0xe2, 0x43, 0xc0, 0x28,
	0xe0, 0x05, 0xfd, 0xa1, 0x71, 0x26, 0x0a, 0xfa, 0xa3, 0xb1, 0x56, 0x66, 0x94, 0x91, 0x95, 0x64,
	0x18, 0x2f, 0x78, 0x61, 0x7f, 0xdc, 0x1b, 0x9e, 0xf5, 0x0d, 0x4d, 0x21, 0x04, 0x0e, 0x79, 0x7b,
	0x32, 0x32, 0x27, 0xc3, 0x8b, 0x7e, 0x5f, 0x53, 0xf5, 0xff, 0x90, 0xa0, 0x22, 0x42, 0x2d, 0x79,
	0x90, 0x51, 0xf2, 0xbd, 0xed, 0x50, 0x9c, 0xd1, 0xd2, 0x7d, 0x15, 0x15, 0x69, 0xe9, 0xbe, 0x8a,
	0x90, 0x1b, 0xbe, 0x8a, 0x0a, 0x36, 0x1b, 0xc9, 0x39, 0x1f, 0x94, 0x37, 0x7c, 0x30, 0x97, 0x15,
	0x2b, 0xf9, 0xd4, 0xbb, 0xf9, 0xf7, 0xe5, 0xcd, 0xd5, 0xab, 0x50, 0x36, 0x9e, 0x6b, 0x25, 0xfc,
	0x1d, 0x1a, 0x5a, 0x19, 0x7f, 0xfb, 0xb8, 0x74, 0xfc, 0x35, 0x34, 0x19, 0x7f, 0xcf, 0x30, 0xb1,
	0xc2, 0x5f, 0x43, 0x53, 0xf1, 0xb7, 0x37, 0xd4, 0x1a, 0x98, 0x2c, 0x9d, 0x8e, 0x86, 0x56, 0xbb,
	0x37, 0x1c, 0x6b, 0x1f, 0xe0, 0xfe, 0x0c, 0x47, 0xd6, 0xa4, 0x37, 0xd4, 0x74, 0x9e, 0xc6, 0x59,
	0x93, 0x94, 0x7b, 0x87, 0xe5, 0x6e, 0x9d, 0x8e, 0x56, 0xc1, 0xc6, 0xf8, 0xe2, 0x44, 0xab, 0x62,
	0x63, 0x70, 0xd1, 0xd7, 0x6a, 0xd8, 0xe8, 0xf4, 0x5e, 0xf0, 0x74, 0x6b, 0x30, 0xea, 0x68, 0x75,
	0x7e, 0x0a, 0xdf, 0x6a, 0xfb, 0x3c, 0xe3, 0xeb, 0x68, 0x07, 0x38, 0xe5, 0xc8, 0xd4, 0x0e, 0x71,
	0x92, 0x93, 0x0b, 0x0b, 0x4f, 0xe6, 0x3d, 0xd6, 0xee, 0x59, 0xc8, 0xbf, 0xc1, 0x52, 0xbb, 0x9e,
	0x35, 0x32, 0x35, 0x4d, 0x90, 0xbf, 0x1b, 0x99, 0xda, 0x4d, 0xae, 0x61, 0xbb, 0x6f, 0x8c, 0x4f,
	0x0d, 0x8d, 0xf0, 0x52, 0xa5, 0x61, 0x1a, 0xda, 0x2d, 0x72, 0x03, 0xea, 0xd6, 0x68, 0x32, 0x68,
	0x5b, 0xa7, 0x5d, 0xcc, 0x0d, 0x8f, 0xc8, 0x2d, 0xb8, 0x61, 0x8d, 0x26, 0xa8, 0x74, 0x4a, 0x7c,
	0x1f, 0x33, 0xc9, 0x67, 0xfd, 0xb6, 0x65, 0x19, 0x43, 0xed, 0xb6, 0xfe, 0x2f, 0x65, 0xa8, 0x88,
	0xab, 0x71, 0xe7, 0xf9, 0x0a, 0x7e, 0x72, 0xbe, 0xcd, 0xb4, 0x9e, 0x5e, 0xe0, 0x57, 0xa2, 0x98,
	0x2e, 0x2c, 0x55, 0x2a, 0x34, 0xfc, 0x1f, 0x3c, 0xe5, 0xbb, 0x00, 0x33, 0x1a, 0x4d, 0xa9, 0x37,
	0x73, 0xbc, 0x05, 0x3b, 0xe6, 0xaa, 0x99, 0xa1, 0xe4, 0xad, 0x40, 0xdd, 0xb0, 0x82, 0x3f, 0x29,
	0x15, 0xf8, 0xc0, 0xa0, 0x37, 0xe4, 0x3e, 0x30, 0x68, 0x7f, 0xc7, 0x7d, 0x60, 0x7c, 0x31, 0xd0,
	0x24, 0xdc, 0x87, 0xf6, 0x0b, 0xc3, 0x6c, 0x9f, 0xa1, 0x35, 0x68, 0x20, 0xca, 0x14, 0x13, 0xbe,
	0x99, 0xcc, 0x1b, 0x04, 0x25, 0xd9, 0x2d, 0x15, 0x33, 0x66, 0xb3, 0x3d, 0xfc, 0x46, 0xa4, 0xda,
	0xc3, 0xf6, 0xf9, 0xb8, 0x3b, 0xb2, 0xb4, 0x2a, 0xf6, 0x9e, 0xf5, 0xcc, 0xb1, 0x35, 0x39, 0x79,
	0xa9, 0xd5, 0xf0, 0x7e, 0x14, 0x25, 0x97, 0x2f, 0x93, 0x7c, 0x95, 0xdf, 0x4f, 0x77, 0x36, 0xaf,
	0xf8, 0xe2, 0xac, 0xf5, 0x07, 0xf3, 0xcd, 0xe2, 0xac, 0x95, 0xed, 0x6d, 0xfe, 0x33, 0x09, 0x46,
	0xd3, 0xe6, 0xff, 0x4a, 0xa0, 0xb0, 0xfc, 0x91, 0x1c, 0x42, 0xe9, 0x92, 0xc9, 0xa9, 0x76, 0xf7,
	0xcc, 0xd2, 0x25, 0xff, 0x6f, 0x4d, 0x39, 0xf7, 0xdf, 0x1a, 0xfe, 0x25, 0x8c, 0xfd, 0xb7, 0x66,
	0xb6, 0xf5, 0x5f, 0x1c, 0x1d, 0x2a, 0x33, 0x3a, 0x75, 0x96, 0xb6, 0xcb, 0x0b, 0x6e, 0x98, 0x48,
	0x09, 0x02, 0xe6, 0xb3, 0x33, 0x3b, 0xe6, 0x17, 0xf4, 0x3e, 0x26, 0x41, 0xd8, 0x43, 0x2a, 0x2b,
	0xbf, 0xaa, 0x62, 0x12, 0x5e, 0x6d, 0xfd, 0x44, 0x64, 0x50, 0x95, 0x6c, 0xd2, 0xc3, 0x54, 0xdc,
	0x4a, 0xa1, 0x96, 0x76, 0x20, 0x12, 0xbf, 0x1b, 0x59, 0x58, 0xa6, 0x8a, 0x7a, 0x9f, 0xe7, 0x59,
	0xb5, 0x9d, 0xb2, 0x58, 0xa2, 0xf5, 0x09, 0xc8, 0x18, 0xad, 0x45, 0xd2, 0x96, 0x83, 0x0d, 0x57,
	0x3c, 0x39, 0x45, 0x3e, 0x2a, 0xbc, 0x5a, 0x39, 0x33, 0x96, 0x3c, 0xb2, 0x65, 0x60, 0x4f, 0x7f,
	0x24, 0x52, 0x8d, 0xf4, 0xa3, 0x4c, 0xee, 0x49, 0x94, 0xfd, 0x28, 0xa3, 0xff, 0x11, 0x2f, 0xd3,
	0x14, 0x97, 0xf7, 0x52, 0xe5, 0x0b, 0xca, 0x7b, 0xc6, 0x8f, 0x94, 0xf7, 0xee, 0xe5, 0x0f, 0xbd,
	0x60, 0x6e, 0xf6, 0x56, 0x51, 0x41, 0xc6, 0x05, 0x9d, 0x54, 0x04, 0xbc, 0xf9, 0x1c, 0x6a, 0x69,
	0xc5, 0x97, 0x3c, 0x80, 0x4a, 0x52, 0xfd, 0xcc, 0xbd, 0xa2, 0x19, 0xc2, 0x4c, 0x78, 0xf8, 0x92,
	0x0e, 0xe9, 0x3c, 0x3f, 0x13, 0x87, 0x20, 0xbd, 0xd9, 0x01, 0x85, 0xf5, 0xf0, 0xb9, 0x62, 0x07,
	0xc1, 0xee, 0x3f, 0x2d, 0x25, 0xdc, 0xcc, 0x87, 0x77, 0x29, 0xf9, 0xf0, 0xfe, 0xd9, 0x77, 0xa0,
	0xb0, 0xff, 0xdc, 0xa1, 0xe7, 0x0c, 0x47, 0x13, 0xd6, 0xd6, 0xf6, 0xc8, 0x21, 0x40, 0xc7, 0xe8,
	0x5b, 0xed, 0xc9, 0xb8, 0x3d, 0x30, 0xb4, 0x12, 0x7a, 0x25, 0xef, 0x9f, 0x76, 0xdb, 0xc3, 0x33,
	0x8c, 0xe2, 0x07, 0x50, 0xe3, 0x14, 0x8c, 0xbb, 0xd2, 0x1a, 0x60, 0x1a, 0x83, 0xd1, 0x0b, 0x43,
	0x93, 0x4f, 0xe4, 0xdf, 0x2f, 0x07, 0x97, 0x97, 0x2a, 0xfb, 0xf7, 0xe1, 0x6f, 0xfd, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0x1e, 0x7e, 0x44, 0x8e, 0x28, 0x00, 0x00,
}
