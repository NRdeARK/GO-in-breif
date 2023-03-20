package main

import "fmt"

const PI = 3.14
func skills(xs ...string) {
	for _,v := range xs {
		fmt.Println(v)
	}
}
func main() {

	const hello = "hello world"
	fmt.Printf("hello: %s\n", hello)
	fmt.Printf("hello: %T\n", hello)

	// const (
	// 	sunday    = 1
	// 	monday    = 2
	// 	tuesday   = 3
	// 	wednesday = 4
	// 	thursday  = 5
	// 	friday    = 6
	// 	saturday  = 7
	// )

	type day int 

	const (
		_ = iota
		sunday day = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Printf("hello: %#v\n", saturday	)
	fmt.Printf("hello: %T\n", saturday)

	//valadict function
	

	skills("1","2","3","4","5","6","7","8","9")

}
