package main

import "fmt"

func test(ar [3]string) {// array param that is used need to have the same size
 	fmt.Println("Hello")
}

func main() {
	var skills = [3]string{"js", "ate", ""}// arrray in go is staic type in other word its cant increase or decrease size 
	s1 := skills[1]
	fmt.Println(s1)
	fmt.Println(len(skills))
	skills[1] = "serf"
	fmt.Println(skills[1])
	test(skills)
}
