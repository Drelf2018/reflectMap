package reflectMap_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Drelf2018/reflectMap"
)

type Tag struct {
	Tag    string
	Fields []Tag
}

func (t Tag) String() string {
	b := strings.Builder{}
	b.WriteString("Tag(")
	b.WriteString(t.Tag)

	for _, f := range t.Fields {
		b.WriteString(", ")
		b.WriteString(f.String())
	}

	b.WriteByte(')')
	return b.String()
}

type TagParser string

func (t TagParser) Simple(m *reflectMap.Map[[]string], typ reflect.Type) (s []string) {
	for _, field := range reflectMap.FieldsOf(typ) {
		s = append(s, field.Tag.Get(string(t)))
	}
	return
}

func (t TagParser) Complex(m *reflectMap.Map[[]Tag], typ reflect.Type) (tags []Tag) {
	for _, field := range reflectMap.FieldsOf(typ) {
		tags = append(tags, Tag{
			Tag:    field.Tag.Get(string(t)),
			Fields: m.MustGetType(field.Type),
		})
	}
	return
}

func (t TagParser) FieldParser(m *reflectMap.Map[[]Tag], field reflect.StructField, elem reflect.Type, data *[]Tag) {
	val, ok := field.Tag.Lookup(string(t))
	if !ok {
		return
	}
	*data = append(*data, Tag{
		Tag:    val,
		Fields: m.MustGetType(field.Type),
	})
}

func NewTagString(tag string, aliases ...reflectMap.Alias) *reflectMap.Map[[]string] {
	return reflectMap.New(TagParser(tag).Simple, aliases...)
}

func NewTagStruct(tag string, aliases ...reflectMap.Alias) *reflectMap.Map[[]Tag] {
	return reflectMap.New(TagParser(tag).Complex, aliases...)
}

func NewTagStructFieldParser(tag string, aliases ...reflectMap.Alias) *reflectMap.Map[[]Tag] {
	return reflectMap.NewFieldParser(TagParser(tag).FieldParser, aliases...)
}

type Struct1 struct {
	Struct2 struct {
		d1 string `ref:"1"`
		d2 int64  `ref:"14"`
	} `ref:"514"`

	Struct3 struct {
		d3 bool    `ref:"19"`
		d4 float64 `ref:"19"`
	}

	D5 *Struct1 `ref:"Struct1"`
	D6 string

	Struct4 struct {
		d7 *Struct1 `ref:"Struct7"`
	} `ref:"Struct4"`
}

var m = NewTagStruct("ref", reflectMap.AddrSlicePtr).Init(Struct1{})
var get = m.Get
var temp []Tag

func BenchmarkMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		temp = m.Get(&[]*Struct1{})
	}
}

func BenchmarkVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		temp = get(&[]*Struct1{})
	}
}

func TestFieldParser(t *testing.T) {
	tags := NewTagStructFieldParser("ref").Get(Struct1{})
	for idx, val := range tags {
		t.Logf("#%d: %v\n", idx, val)
	}
}
