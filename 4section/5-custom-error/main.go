package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrorDivisionByZero = errors.New("divide by zero")
var ErrorNumTooLarge = errors.New("number is too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func newOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func DoSomething() error {
	return newOpError("doSomething", 100, "doSomething failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrorDivisionByZero
	}

	if a > 1000 {
		return 0, ErrorNumTooLarge
	}

	return a / b, nil
}

func main() {
	value, err := divide(1001, 0)
	if err != nil {
		if errors.Is(err, ErrorDivisionByZero) {
			fmt.Println("divide by zero")
		} else if errors.Is(err, ErrorNumTooLarge) {
			fmt.Println("number is too large")
		}
		return
	}

	fmt.Println(value)
}
