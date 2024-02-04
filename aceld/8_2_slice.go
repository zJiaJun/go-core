package main

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func printSlice(myArray []int) {
	// _ 表示匿名变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func sliceCreate() {
	//声明slice1是一个切片, 默认值是1,2,3, 长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//声明slice2是一个切片,但是并没有给slice分配空间
	var slice2 []int
	if slice2 == nil {
		fmt.Println("slice2 是一个空切片")
	}
	slice2 = make([]int, 3) //开辟3个空间
	slice2[0] = 1
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//声明slice3是一个切片
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	//声明slice4是一个切片, := 推导slice4是切片
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)
}

func sliceCap() {
	//长度3, 容量5
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//向numbers切片追加一个元素1, len = 4 [0,0,0,1] cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//向numbers切片追加一个元素2, len = 5 [0,0,0,1,2] cap = 5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//容量已满, 添加元素, 扩展的容量, len = 6, new_cap = cap * 2 = 10
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
}
func sliceCut() {
	numbers := []int{1, 2, 3}

	//[0,2) 取下标0,1元素,2不包括
	n1 := numbers[0:2]
	fmt.Println(n1)

	//[0,1)
	n2 := numbers[:1]
	fmt.Println(n2)
	//[1, len(s))
	n3 := numbers[1:]
	fmt.Println(n3)

	//numbers[0] 也会等于100
	n1[0] = 100
	fmt.Println(n1)
	fmt.Println(numbers)

	//copy 可以将底层数组的slice一起进行拷贝
	np := make([]int, 3)
	//将numbers的拷贝到np中
	copy(np, numbers)
	fmt.Println(np)

}

func SliceExample() {
	defer fmt.Println(common.Separator)
	//动态数组,切片 slice
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray types = %T\n", myArray)
	printSlice(myArray)

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	sliceCreate()
	sliceCap()
	sliceCut()
}
