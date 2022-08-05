package complex

import (
	"fmt"
	"github.com/zjiajun/go-core/best_practice/error_handing/normal"
)

type errType string

const (
	errNotFound     errType = "item not found"
	errMissingParam errType = "missing param"
	errUnKnown      errType = "unKnownErr"
)

type BusinessError struct {
	errType errType
	msg     string
}

func NewBusinessError(errType errType, msg string) *BusinessError {
	return &BusinessError{errType, msg}
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("%v_%v", e.errType, e.msg)
}

func getItem(key string) (any, error) {
	cn := normal.NewMap()
	val, ok := cn[key]
	if !ok {
		return nil, NewBusinessError(errNotFound, key)
	}
	return val, nil
}

func HandlingErrorWitchCustomError(key string) {
	val, err := getItem(key)
	if err != nil {
		switchErrByTypeChecking(err)
		return
	}
	fmt.Println("未发生异常", val)
	return
}

func switchErrByTypeChecking(err error) {
	if err != nil {
		switch err.(type) {
		case *BusinessError:
			fmt.Println(err.Error())
		default:
			fmt.Println(errUnKnown)
		}
	}
}
