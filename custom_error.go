package main

import (
	"errors"
	"fmt"
)

type argErr struct {
	arg     int
	message string
}

func (e *argErr) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func egenerator(arg int) (int, error) {
	if arg == 42 {
		return -1, &argErr{arg, "can't work with that"}
	}
	return arg + 3, nil
}

func main() {
	_, err := egenerator(42)
	var ae *argErr
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
