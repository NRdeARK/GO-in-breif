package main

import (
	"fmt"
)

type course struct {
	name, instructor string
	price int
}

func (c course) discount(d int) int {// way to create method c(this in many lang) -> reciver 
 	p := c.price - d
	fmt.Println("discount:",p)
	return p
}

func main() {
	c := course{"Basic Go","anu",999}
	fmt.Printf("%#v\n", c)

	d := c.discount(199)
	fmt.Println("discounted price:",d)
}