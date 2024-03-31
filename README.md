# reflectMap
 
反射字典，以及更好的反射。

原项目 [Reflect](https://github.com/Drelf2018/TypeGo/tree/main/Reflect)

### 使用

```go
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
	for _, data := range m.Get(&College{}) {
		t.Log(data)
	}
}
```

#### 控制台

```
parser_test.go:31: Data#0(./td[1]//span/text())
parser_test.go:31: Data#1(./td[2]//a/@href)
parser_test.go:31: Data#3(//td//span[a] | //form//li[a] | //ul[@style='list-style-type:none;']//li[a] | //ul[@class='attach']//li[a])
```

### 写在最后

写给自己用的，如果真的真的有人用可以直接在 issues 问我咋用🥺🥺