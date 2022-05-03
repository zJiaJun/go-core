package goweb

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func init() {
	fmt.Println("function init running...")
}

func filter(slice []int, f numFunction) []int {
	var result []int
	for val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

type numFunction func(int) bool

func isOdd(val int) bool {
	if val%2 == 0 {
		return false
	}
	return true
}
func isEven(val int) bool {
	if val%2 == 0 {
		return true
	}
	return false
}

func functionAsTypeExample() {
	defer fmt.Println(common.Separator)
	var slice = []int{1, 2, 3, 4, 5, 6}
	odd := filter(slice, isOdd)
	fmt.Println("FunctionExample", odd)
	even := filter(slice, isEven)
	fmt.Println("FunctionExample", even)
}
