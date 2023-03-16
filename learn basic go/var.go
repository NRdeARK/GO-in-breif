package main

import "fmt"

var msg4 string = "test4" // declaire global var
func main() {
	var msg1 string= "test1" // normal declair/init
	var msg2 = "test2" // short go will select type by value on right side
	msg3 := "test3" // more short but can't declair global like this
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println(msg4)
	//msg1 = 55 *** // can't change type of var because go is  statically typed
	msg1 = "test5"// can change value by the way ~
	fmt.Println(msg1)
	msg2 = msg1 + msg2// can concat string
	fmt.Println(msg2)
	msg5 := 55// can assign int
	fmt.Println("msg5 is contains :",msg5)// multiple print
	msg6,msg7 := 22 , 44//multiple declairation
	fmt.Println(msg6 + msg7)
}
