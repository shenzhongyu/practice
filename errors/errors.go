package main

import (
	"fmt"
	"strconv"
	"errors"
)

var fp = fmt.Println

type argError struct {
	arg int
	err error
}

// go报错自定义， 自由定制错误格式
func (e *argError) Error() string {
	return "[argument:" + strconv.Itoa(e.arg) + "] --- " + e.err.Error()
}

func f1(arg int) (int, error) {
	if arg == 0 {
		return arg, errors.New("can't work with it.")
	}
	return arg * 3, nil
}

func f2(arg int) (int, error) {
	if arg == 0 {
		return arg, &argError{arg, errors.New("can't work with it.")}
	}
	return arg * 4, nil
}

func main () {
	for _, i := range []int{1, 3, 0} {
		if r, e := f1(i); e != nil {
			fp("f1 failed:", e)
		} else {
			fp("f1 worked:", r)
		}
	}

	for _, j := range []int{3, 8, 0} {
		if r, e :=  f2(j); e != nil {
			fp("f2 failed:", e)
		} else {
			fp("f2 worked:", r)
		}
	}
}