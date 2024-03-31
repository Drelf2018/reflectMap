package reflectMap_test

import (
	"testing"

	"github.com/Drelf2018/reflectMap"
)

type Attachment struct {
	CollegeName string
	Name        string `xpath:"./a//text()"`
	Url         string `xpath:"./a/@href"`
}

type College struct {
	Name        string `xpath:"./td[1]//span/text()" gorm:"primaryKey"`
	Url         string `xpath:"./td[2]//a/@href"`
	Temp        string
	Attachments []Attachment `xpath:"//td//span[a] | //form//li[a] | //ul[@style='list-style-type:none;']//li[a] | //ul[@class='attach']//li[a]"`
}

func TestTagParser(t *testing.T) {
	m := reflectMap.NewTagParser("xpath", func(s string) string { return s })
	// .SetConverter(func(fieldType reflect.Type) reflect.Type {
	// 	if fieldType.Kind() == reflect.Slice {
	// 		return fieldType.Elem()
	// 	}
	// 	return fieldType
	// })
	for _, data := range m.Get(&College{}) {
		t.Log(data)
	}
}
