package main

import "fmt"

func main() {
	var msg string// variable without init is called Zero value so when variable is created in go it already has value
	var msgInt int
	fmt.Println(msg)
	fmt.Printf("%q\n",msg)
	fmt.Println(msgInt)
}