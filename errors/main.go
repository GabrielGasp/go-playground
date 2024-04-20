package main

import (
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() pkgErrors.StackTrace
}

func main() {
	err1 := errors.New("error 1")
	err2 := pkgErrors.New("error 2")

	fmt.Println(err1)
	fmt.Println(err2)

	if tracer, ok := err2.(stackTracer); ok {
		fmt.Printf("%s: %+v\n", err2, tracer.StackTrace())
	}
}
