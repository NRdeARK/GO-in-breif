package main

import "fmt"

func main() {
	//var ss []string = []string{}
	var m map[string]int = map[string]int{}

	fmt.Println("the value", m)
	m["answer"] = 42
	fmt.Printf("The value %#v\n",m)

	v:= m["answer"]
	fmt.Printf("The value %#v\n",v)

	m["nong"] = 15
	m["nong"] = 20

	fmt.Printf("The value %#v\n",m["nong"])

	delete(m, "answer")
	fmt.Printf("The value %#v\n",m)

	n,ok := m["nong"]

	fmt.Printf("The value %#v %#v\n",n,ok)

	n,ok = m["test"]

	fmt.Printf("The value %#v %#v\n",n,ok)



	


}