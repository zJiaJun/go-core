package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"io"
	"os"
)

func pair1Example() {
	// a变量中会存在一个pair<static type:string, value:"abc">
	var a string
	a = "abc"

	// pair<type:string, value:"abc">
	var allType interface{}
	//a赋值给allType, allType的pair还是a的pair, 保持不变
	allType = a
	var val, _ = allType.(string)
	fmt.Println(val)
}

func pair2Example() {
	//tty: pair<type:*os.File, value:"dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	//r: pair<type: , value: > type和value都是空
	var r io.Reader
	//r: pair<type:*os.File, value:"dev/tty"文件描述符>
	r = tty

	//w: <type:*os.File, value:"dev/tty"文件描述符>
	var w io.Writer
	w = r.(io.Writer)
	c, err := w.Write([]byte("hello\n"))
	fmt.Println(c)
}

type reader interface {
	readBook()
}
type writer interface {
	writeBook()
}
type bookObj struct {
}

func (this *bookObj) readBook() {
	fmt.Println("read book")
}
func (this *bookObj) writeBook() {
	fmt.Println("write book")
}

func pair3Example() {
	//b: pair<type:*bookObj, value:bookObj指针地址>
	b := &bookObj{}

	//r: pair<type,value>
	var r reader
	//r: pair<type:*bookObj, value:bookObj指针地址>
	r = b
	r.readBook()

	var w writer
	//w: pair<type:*bookObj, value:bookObj指针地址>
	w = r.(writer) //断言成功, w 和 r 具体的type是一致的
	w.writeBook()
}

func Reflect111Example() {
	defer fmt.Println(common.Separator)
	pair1Example()
	pair2Example()
	pair3Example()
}
