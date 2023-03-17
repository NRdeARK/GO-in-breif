package main

import "fmt"

func add(x float64, y float64) float64 { // to rertune val nedd to assign type
	v := x+y  
	return v // can't directly return x+y you need to return v
}

func add2(x float64, y float64) (float64,float64,float64){ // to return multiple var at once
	v := x+y  
	return v,x,y
}

func split(sum float64,)(x float64, y float64){
	x = sum /4
	y = 3*sum/4
	return// this called naaked return using this by assign the var that will return then you just use return to return those val CAUTION : this shouldn't used with long func
}

var add3 = func(x float64, y float64) (float64,float64,float64){ // assingn by using
	v := x+y  
	return v,x,y
}

func compute(fn func(float64,float64)float64) (float64){// you can used function as param to used in fuction
	v := fn(34,45)
	return v
}

func main() {
	fmt.Println(add(42, 34))
	fmt.Println(add2(42, 34))
	fmt.Println(split(16))
	fmt.Println(add3(44,45))
	fmt.Printf("%T\n",add3)//you will see the type of this func as func(float64, float64) (float64, float64, float64)
	fmt.Println(compute(add))
	//so as you can see this func is just some var or val in go
}