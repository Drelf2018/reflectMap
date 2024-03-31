package reflectMap

import "reflect"

type Value struct {
	reflect.Value
}

func (v Value) Call(in ...reflect.Value) []reflect.Value {
	return v.Value.Call(in)
}

func (v Value) CallSlice(in ...reflect.Value) []reflect.Value {
	return v.Value.CallSlice(in)
}

func (v Value) CallAny(in ...any) []reflect.Value {
	return v.Value.Call(ValuesOf(in...))
}

func (v Value) CallAnySlice(in ...any) []reflect.Value {
	return v.Value.CallSlice(ValuesOf(in...))
}

func ValueOf(i any) Value {
	return Value{reflect.ValueOf(i)}
}

func ValuesOf(in ...any) []reflect.Value {
	val := make([]reflect.Value, 0, len(in))
	for _, i := range in {
		val = append(val, reflect.ValueOf(i))
	}
	return val
}
