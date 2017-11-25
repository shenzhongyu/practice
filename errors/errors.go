package main

import (
	"strconv"
	"errors"
	"fmt"
)

type argError struct {
	arg int
	err error
}

func (e *argError) Error () string {
	return "[argument:" + strconv.Itoa(e.arg) + "] --- " + e.err.Error()
}

func f1 (arg int) (int, error) {
	if arg == 0 {
		return arg, errors.New("can 't work with it.")
	}
	return arg * 2, nil
}

func f2(arg int) (int, error) {
	if arg == 0 {
		return arg, &argError{arg, errors.New("can 't with it.")}
	}
	return arg * 3, nil
}

func main() {
	for _, i := range []int {3, 4, 0} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, j :=  range []int {2, 4, 0} {
		if r, e :=  f2(j); e != nil {
			fmt.Println("f2: failed:", e)
		} else {
			fmt.Println("f2: worked:", r)
		}
	}
}
