package reflectMap_test

import (
	"reflect"
	"testing"

	"github.com/Drelf2018/reflectMap"
)

type M struct{}

func (*M) Test1(int)   {}
func (M) Test2(string) {}
func (*M) Test3(bool)  {}
func (M) Test4(a, b int) float64 {
	return float64(a + b)
}

func TestMethod(t *testing.T) {
	s := reflect.ValueOf(&M{})
	for k, v := range reflectMap.MethodsOf(s) {
		t.Logf("%v: %v, embedded: %v", k.Name, v.Type(), reflectMap.IsEmbeddedMethod(k))
	}
}

func TestFunctions(t *testing.T) {
	s := reflect.ValueOf(&M{})
	for name, fn := range reflectMap.FunctionsOf(s) {
		switch fn := fn.(type) {
		case func(int, int) float64:
			t.Logf("%s: %v\n", name, fn(114, 514))
		}
	}
}
