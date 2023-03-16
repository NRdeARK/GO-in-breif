package main

import "fmt"

func main() {
	var r rune = 'A' // rune can use as pointer to unicode
	var f float64 = 54
	fmt.Println(r)
	fmt.Printf("r: %c\n", r) // use printf to fix format of msg
	fmt.Printf("r: %d\n", r) // use printf to fix format of msg
	fmt.Printf("f: %.2f\n", f)

	fmt.Printf("r: %#v\n", r)//  %#v => can used to all type
	fmt.Printf("r: %T\n", r)// %T => give you type of that var

}
