package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6}
	odd := filter(slice, isOdd)
	fmt.Println(odd)
	even := filter(slice, isEven)
	fmt.Println(even)

	var P person
	P.name = "jiajun"
	P.age = 20
	Q := person{human{120, 123},
		"name", 10, 456}
	fmt.Println(Q.phone)
	fmt.Println(Q.human.phone)
	Q.name, Q.age = "leon", 20
	W := person{age: 20, name: "300"}
	E := new(person)
	E.human.weight = 1
	fmt.Println(Q)
	fmt.Println(W)
	fmt.Println(E)
}

func filter(slice []int, f cfunc) []int {
	var result []int
	for val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

type cfunc func(int) bool

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

type person struct {
	human
	name  string
	age   int
	phone int
}

type human struct {
	weight int
	phone  int
}
