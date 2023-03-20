package main

import "fmt"

func main() {
	defer fmt.Println("a")// defer will work when func end
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i :=0; i< 10 ; i++ {
		defer fmt.Println(i)
		defer func() {
			fmt.Println(i)// you will see 10 because 10 is point to value i when finish loop
		}()
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

	fmt.Println("test 1,2,3")
}