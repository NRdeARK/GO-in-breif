package main

import (
	"errors"
	"fmt"
)

// type error interface { already in go
// 	Error() string
// }

type ErrorType1 struct {}

func (e ErrorType1) Error() string {
	return "type1"
}

var type2 = errors.New("type2")

func divide(a, b int) (int,error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	s := a/b
	return s, nil
}

func main() {
	r, err := divide(1,1)
	if err != nil {
		fmt.Println("handler error:",err)
		return
	}
	fmt.Println(r,err)
}