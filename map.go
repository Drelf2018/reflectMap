package reflectMap

import (
	"fmt"
	"reflect"
)

type Alias func(reflect.Type) uintptr

type ElemParser[V any] func(*Map[V], reflect.Type) V

type FieldParser[V any] func(m *Map[V], field reflect.StructField, elem reflect.Type, data *V)

type Map[V any] struct {
	types map[uintptr]V

	ElemParser  ElemParser[V]
	FieldParser FieldParser[V]

	TypeConverter func(reflect.Type) reflect.Type
	Aliases       []Alias
}

func (m *Map[V]) SetConverter(f func(reflect.Type) reflect.Type) *Map[V] {
	m.TypeConverter = f
	return m
}

func (m *Map[V]) Clear() {
	for k := range m.types {
		delete(m.types, k)
	}
}

func (m *Map[V]) GetType(elem reflect.Type) (v V, ok bool) {
	// type check
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	if m.TypeConverter != nil {
		elem = m.TypeConverter(elem)
	}
	if elem.Kind() != reflect.Struct {
		return
	}
	// return an existing value
	elemPtr := Type(elem)
	v, ok = m.types[elemPtr]
	if ok {
		return
	}
	// initial zero value
	m.types[elemPtr] = v
	// parse actual value
	if m.ElemParser != nil {
		v = m.ElemParser(m, elem)
	} else if m.FieldParser != nil {
		for _, field := range FieldsOf(elem) {
			m.FieldParser(m, field, elem, &v)
		}
	}
	// set alias
	m.types[elemPtr] = v
	for _, alias := range m.Aliases {
		m.types[alias(elem)] = v
	}
	return v, true
}

func (m *Map[V]) MustGetType(elem reflect.Type) (v V) {
	if elem != nil {
		v, _ = m.GetType(elem)
	}
	return
}

func (m *Map[V]) Ptr(in uintptr) (v V, ok bool) {
	v, ok = m.types[in]
	return
}

func (m *Map[V]) Get(in any) V {
	if v, ok := m.types[Ptr(in)]; ok {
		return v
	}
	if v, ok := m.GetType(reflect.TypeOf(in)); ok {
		return v
	}
	panic(fmt.Errorf("reflectMap: unsupported type: %#v , maybe you forgot to use an alias", in))
}

func (m *Map[V]) Gets(in ...any) []V {
	v := make([]V, 0, len(in))
	for _, i := range in {
		v = append(v, m.Get(i))
	}
	return v
}

func (m *Map[V]) Init(in ...any) *Map[V] {
	m.Gets(in...)
	return m
}

func New[V any](p ElemParser[V], aliases ...Alias) *Map[V] {
	return &Map[V]{
		types:      make(map[uintptr]V),
		ElemParser: p,
		Aliases:    aliases,
	}
}

func NewFieldParser[V any](p FieldParser[V], aliases ...Alias) *Map[V] {
	return &Map[V]{
		types:       make(map[uintptr]V),
		FieldParser: p,
		Aliases:     aliases,
	}
}

func NewDataParser[V any](p DataParser[V], aliases ...Alias) *Map[[]Data[V]] {
	return New(p.Do, aliases...)
}

func NewTagParser[V any](tag string, f func(string) V, aliases ...Alias) *Map[[]Data[V]] {
	return NewDataParser(func(field reflect.StructField, elem reflect.Type) (v V, ok bool) {
		var val string
		val, ok = field.Tag.Lookup(tag)
		if ok {
			v = f(val)
		}
		return
	}, aliases...)
}

func NewTagOKParser[V any](tag string, f func(string) (v V, ok bool), aliases ...Alias) *Map[[]Data[V]] {
	return NewDataParser(func(field reflect.StructField, elem reflect.Type) (v V, ok bool) {
		var val string
		val, ok = field.Tag.Lookup(tag)
		if ok {
			v, ok = f(val)
		}
		return
	}, aliases...)
}

func NewTagErrorParser[V any](tag string, f func(string) (v V, err error), aliases ...Alias) *Map[[]Data[V]] {
	return NewDataParser(func(field reflect.StructField, elem reflect.Type) (v V, ok bool) {
		var val string
		val, ok = field.Tag.Lookup(tag)
		if ok {
			var err error
			v, err = f(val)
			ok = err == nil
		}
		return
	}, aliases...)
}
