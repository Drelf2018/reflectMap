package reflectMap

import (
	"reflect"
	"unsafe"
)

// 反射加速
//
// 参考: https://www.cnblogs.com/cheyunhua/p/16642488.html
type eface struct {
	Type  unsafe.Pointer
	Value unsafe.Pointer
}

func (e *eface) Ptr() uintptr {
	return uintptr(e.Type)
}

func (e *eface) Word() uintptr {
	return uintptr(e.Value)
}

func EFace(in any) *eface {
	return (*eface)(unsafe.Pointer(&in))
}

func Ptr(in any) uintptr {
	return EFace(in).Ptr()
}

func Word(in any) uintptr {
	return EFace(in).Word()
}

func Type(typ reflect.Type) uintptr {
	return Word(typ)
}

func Addr(typ reflect.Type) uintptr {
	return Type(reflect.PointerTo(typ))
}

func Slice(typ reflect.Type) uintptr {
	return Type(reflect.SliceOf(typ))
}

func SlicePtr(typ reflect.Type) uintptr {
	return Type(reflect.SliceOf(reflect.PointerTo(typ)))
}

func AddrSlice(typ reflect.Type) uintptr {
	return Addr(reflect.SliceOf(typ))
}

func AddrSlicePtr(typ reflect.Type) uintptr {
	return Addr(reflect.SliceOf(reflect.PointerTo(typ)))
}
