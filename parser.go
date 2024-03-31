package reflectMap

import (
	"fmt"
	"reflect"
	"strings"
)

type Data[V any] struct {
	Index  int
	V      V
	Fields []Data[V]
}

func (d Data[V]) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("Data#%d(%v", d.Index, d.V))

	for _, f := range d.Fields {
		b.WriteString(", ")
		b.WriteString(f.String())
	}

	b.WriteByte(')')
	return b.String()
}

type DataParser[V any] func(field reflect.StructField, elem reflect.Type) (v V, ok bool)

func (d DataParser[V]) Do(m *Map[[]Data[V]], elem reflect.Type) (data []Data[V]) {
	for idx, field := range FieldsOf(elem) {
		v, ok := d(field, elem)
		if ok {
			data = append(data, Data[V]{
				V:      v,
				Index:  idx,
				Fields: m.MustGetType(field.Type),
			})
		}
	}
	return
}
