package main

import "fmt"

type human struct {
	name  string
	job   string
	money int64
}

func main() {
	// name := "NR"
	// job := "unemployed"
	// money := 9999
	c := human{
		name:  "NR",
		job:   "unemployed",
		money: 9999,
	}
	// c2 := human{"NR","unemployed",9999}\
	// c3 := human{name:"NR"}


	fmt.Println("name:", c.name)
	fmt.Println("job:", c.job)
	fmt.Println("money:", c.money)

	c.name = "test"

	fmt.Println("name:", c.name)
	fmt.Println("job:", c.job)
	fmt.Println("money:", c.money)
}

