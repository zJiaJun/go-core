package main

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func mapCreate() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是空map", myMap1)
	}
	//在使用map前, 需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "golang"
	fmt.Println(myMap1)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "golang"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "golang",
	}
	fmt.Println(myMap3)
}

func printMap(cityMap map[string]string) {
	//cityMap 引用传递
	for key, value := range cityMap {
		fmt.Println("key =", key, ",value =", value)
	}
}

func MapExample() {
	defer fmt.Println(common.Separator)
	mapCreate()

	cityMap := make(map[string]string)
	cityMap["China"] = "Shanghai"
	cityMap["Japan"] = "Toyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)
	//删除
	delete(cityMap, "USA")
	printMap(cityMap)

}
