package main

import (
	"github.com/zjiajun/go-core/best_practice/factory"
)

func main() {
	complexVar := complex(3, -1.18)
	println(complexVar)
	a := imag(complexVar)
	b := real(complexVar)
	println(a, b)
	factory.Run()
}
