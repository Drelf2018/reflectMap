package reflectMap_test

import (
	"reflect"
	"testing"

	"github.com/Drelf2018/reflectMap"
)

type A struct {
	Name string
	Age  int64
}

var u uintptr
var a = A{
	Name: "12138",
	Age:  17,
}
var typeOfA = reflect.TypeOf(a)

func OldAddr(typ reflect.Type) uintptr {
	return reflectMap.Ptr(reflect.New(typ).Interface())
}

func OldType(typ reflect.Type) uintptr {
	return reflectMap.Ptr(reflect.Zero(typ).Interface())
}

func BenchmarkPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = reflectMap.Ptr(a)
	}
}

func BenchmarkType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = reflectMap.Type(typeOfA)
	}
}

func BenchmarkOldType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = OldType(typeOfA)
	}
}

func BenchmarkAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = reflectMap.Addr(typeOfA)
	}
}

func BenchmarkOldAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u = OldAddr(typeOfA)
	}
}

// BenchmarkPtr-8       	1000000000	         1.174 ns/op	       0 B/op	       0 allocs/op
// BenchmarkType-8      	1000000000	         0.7794 ns/op	       0 B/op	       0 allocs/op
// BenchmarkOldType-8   	179699146	         6.905 ns/op	       0 B/op	       0 allocs/op
// BenchmarkAddr-8      	153470544	         7.938 ns/op	       0 B/op	       0 allocs/op
// BenchmarkOldAddr-8   	29307964	         40.51 ns/op	      24 B/op	       1 allocs/op
