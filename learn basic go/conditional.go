package main

import "fmt"

func main() {
	var num int = -65
	if num == 14 {
		fmt.Println(num , "is 14")
		
	} else if(num>14) || (num<14 && num>=-1) {// &&, || is used as and, or
		fmt.Println(num, "is pos int")
	} else {// "{" is still fixed placed
		fmt.Println(num, "is not pos int")
	}
	//go also have sw case

	switch key := "tac"; key {// the difference from normal switch case is the sw case in go doesn't need to break it will only go to that cond and end we can also use regular ex in case "case c<b"
	case "tic":
		fmt.Println("this is tic")
	case "tac","toe":// can put two conf in one case
		fmt.Println("this is tac or toe")
	default:// like else in if
		fmt.Println("WTF")
	}

}