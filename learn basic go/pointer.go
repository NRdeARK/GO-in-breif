package main

import "fmt"

func changePriceByValue(price int){//this call value type reciever
	price = price -499
	//fmt.Printf("%#v %#v\n",price, &price)
}

func changePriceByAddress(price *int){
	*price = *price -499
	//fmt.Printf("%#v %#v\n",*price, price)
}

type course struct {
	name, instructor string
	price int
}

func discount(c *course) int {//this call pointer type
 	c.price = c.price - 1000//method will automatically 
	fmt.Println("discount:",c.price)
	return c.price
}

func main() {
	var price int = 99999
	var addr *int = &price// using  & to get pointer from var

	fmt.Println(price, addr)
	fmt.Println("old", *addr, price)

	*addr = 92123 // deref from pointer to var 

	fmt.Println("new", *addr, price)// 

	v := *addr

	changePriceByValue(price)// change price by passing value
	fmt.Printf("change from %#v %#v to %#v %#v\n", price, addr, v, &v)// var val isn't change
	fmt.Printf("change from %#v %#v to", price, addr)
	changePriceByAddress(addr)// change price using address
	fmt.Printf(" %#v %#v\n", price, addr)

	c := &course{"test","tester",8878}// store pointer of struct
	fmt.Println("C price:")
	discount(c)

	d := new(course)// another way this can we used when we using new create new struct it will pass pointer
	fmt.Println("D price:")
	discount(d)

	//zero pointer of pointer is nil
}

	
