package main

import (
	"fmt"
	"strings"
)

// TODO: write code below.
func WordCount(s string) map[string]int {
	r := map[string]int{}
	a := strings.Fields(s)
	for _, v := range a{
		r[v] = r[v] +1
	}
	return r
}

func main() {
	// Workshop: maps
	// กำหนด: 1. ให้สร้างเขียนฟังก์ชัน WordCount เพื่อนำคำซ้ำในประโยค
	//           strings.Fields น่าจะเป็นตัวช่วยได้ https://pkg.go.dev/strings#Fields
	//
	// Output:
	// map[string]int{"If":1, "a":4, "and":1, "duck":4, "is":1, "it":2, "like":3, "looks":1, "probably":1, "quacks":1, "swims":1, "then":1}

	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	w := WordCount(s)

	fmt.Printf("%#v\n", w)
}
