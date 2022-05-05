package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

//类名首字母大写,表示其他包能够访问
type Hero struct {
	Name  string
	level int //写小,无法访问,私有的
}

func (this *Hero) Show() {
	fmt.Println(this)
}
func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(newName string) {
	//*Hero, 如不加*, 就是副本,不是指针
	this.Name = newName
}

func Oop102Example() {
	defer fmt.Println(common.Separator)
	hero := Hero{Name: "qwe", level: 10}
	hero.Show()
	hero.SetName("newName")
	hero.Show()
}
