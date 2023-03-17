package main

import "fmt"

func inc(a float64) float64 {
	return a + 1
}

func curr(a float64) float64 {
	return a
}

func compute() (func(float64) float64, func(float64) float64) {//this compute func is called higher order func
	return inc, curr
}

func main() {
	i, c := compute()
	fmt.Println(i)//i is func by calling like that you will get address
	fmt.Println(i(1))
	fmt.Println(c)
	fmt.Println(c(1))

}