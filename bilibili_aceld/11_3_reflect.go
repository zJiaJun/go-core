package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"reflect"
)

type resume struct {
	//字段绑定两个标签, key = info, value = name
	name string `info:"name" doc:"名字"`
	sex  string `info:"sex"`
}

func findTag(r interface{}) {
	//得到当前结构体的元素
	t := reflect.TypeOf(r).Elem()
	for i := 0; i < t.NumField(); i++ {
		t := t.Field(i).Tag
		tInfo := t.Get("info")
		tDoc := t.Get("doc")
		fmt.Println("info:", tInfo, "doc:", tDoc)
		value, ok := t.Lookup("info")
		fmt.Println("value:", value, "ok:", ok)
	}
}

func Reflect113Example() {
	defer fmt.Println(common.Separator)
	r := resume{"name", "sex"}
	findTag(r)
}
