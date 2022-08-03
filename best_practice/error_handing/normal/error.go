package normal

import (
	"errors"
	"fmt"
)

// 简单错误 - 原生errors package + 直接比较

var (
	errNotFound     = errors.New("item not found")
	errMissingParam = errors.New("missing param")
	errUnKnown      = "unKnownErr"
)

func NewMap() map[string]string {
	resMap := make(map[string]string)
	resMap["name"] = "Leon"
	resMap["age"] = "30"
	return resMap
}

func HandingErrorInSimpleCase(key string) {
	val, err := getItem(key)
	if err != nil {
		switchErrByComparison(err)
		return
	}
	fmt.Println("未发生异常", val)
	return
}

func getItem(key string) (any, error) {
	cn := NewMap()
	val, ok := cn[key]
	if !ok {
		return nil, errNotFound
	}
	return val, nil
}

func switchErrByComparison(err error) {
	if err == nil {
		return
	}
	switch err {
	case errNotFound:
		//记录、上报错误日志
		fmt.Println(errNotFound.Error())
	case errMissingParam:
		fmt.Println(errMissingParam.Error())
	default:
		fmt.Println(errUnKnown)
	}
}
