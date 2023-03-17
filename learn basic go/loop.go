package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {//in go there is other loop than for
		fmt.Println("i:",i)
	}

	sum := 0
	for sum < 10 {
		sum += 1
		fmt.Println("sum:",sum)
	}

	skills := [2]string{"tes1","tes2"}

	for i:= range skills {// i will be index of skill
		fmt.Println("i:",i,skills[i])
	}

	for i,val := range skills {// assign i as index and val as value of array
		fmt.Println("i:",i,val)
	}

	for _,val := range skills {// if you dont want to used index you can do like this by replace i with _ because go will force you to used every var
		fmt.Println("i:",val)
	}
}