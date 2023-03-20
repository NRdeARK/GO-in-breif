package main

import "fmt"

func show (val int){
	fmt.Printf("val value:%#v type: %T\n",val,val)
}

func show2 (val any){
	val2 := val.(int)// if you dont change type you cant do operation same as value  of var
	val2 += 1
	fmt.Printf("val value:%#v type: %T\n",val2,val2)
}

func show3 (val any){
	switch v:= val.(type){
	case string:
		i := v + "test"
		fmt.Println("string",i)
	case int:
		i := v + 2
		fmt.Println("int :",i)
	
	default:
		fmt.Println("type not support")
	}
}

func sale(val interface{}){
	fmt.Println("sale:", val)
}

type promotion interface{// only require you can also put other interface in this
	discount() int
}

func sale2(val promotion){// any struc that satisfied interface req 
	fmt.Println("sale:", val ,"discount:",val.discount())// you can use the function that have in promotion
}

type course struct{
	price int
}

func (c course) discount() int {
	c.price -= 1000
	return c.price
}

func main() {
	var v interface{}// ist the same as any type in other languague
	v = 26
	fmt.Printf("v value:%#v type: %T\n",v,v)
	v = "test"
	fmt.Printf("v value:%#v type: %T\n",v,v)

	//go >1.8 change interface => any

	var f any
	f = 26
	fmt.Printf("f value:%#v type: %T\n",f,f)
	f = "test"
	fmt.Printf("f value:%#v type: %T\n",f,f)
	f= 34

	show(f.(int))//type assertion

	show2(f)

	show3(v)

	var gc = course{1999}
	sale2(gc)
}