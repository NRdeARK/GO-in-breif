package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	var f float64 = float64(i)
	fmt.Println(f)
	var u uint8 = uint8(f)
	fmt.Println(u) // when covert from bigger type in to smaller one value = 0
	v := "2222g"
	s, err := strconv.Atoi(v)// string to int
	fmt.Println(s, err)

	v = strconv.Itoa(i)// int to string

	fmt.Println(v)
}
