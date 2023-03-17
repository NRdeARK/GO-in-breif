package main

import "fmt"

func compute() (func() float64, func() float64) { //this compute func is called higher order func
	sum := 0.0
	return func() float64 {
			sum = sum + 1
			fmt.Printf("1: %v\n", sum)
			return sum
		}, func() float64 {
			fmt.Printf("2: %v\n", sum)
			return sum
		}
}

func main() {
	inc, curr := compute() // sum is in scope of inc,curr so it will not reset state(sum) until end of this func
	v1 := inc()
	v2 := curr()
	// sum = 1
	fmt.Println(v1, v2)
	v := inc() // sum = 2 so it make this func still can ref to var sum
	fmt.Println(v)
	y := curr()
	fmt.Println(y)

}
