package main

import (
	"errors"
	"fmt"
)

type error1 struct {
	msg string
}

func newError1(msg string) error {
	return error1{msg}
}

func (e error1) Error() string {
	return e.msg
}

func main() {
	err := newError1("error1")
	err = fmt.Errorf("error2: %w", err)

	if _, ok := err.(error1); ok {
		fmt.Println("err is error1")
	} else {
		fmt.Println("err1 is not error1")
	}

	var err1 *error1
	errors.As(err, &err1)

	if err1 == nil {
		fmt.Println("err is error1")
	} else {
		fmt.Println("err1 is not error1")
	}
}
