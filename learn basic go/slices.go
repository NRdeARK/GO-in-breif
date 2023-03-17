package main

import "fmt"

func test(slices []string) {// array param that is used need to have the same size
	fmt.Println("Hello")
}

func main() {
	skills := []string{"str", "int", "float"}// the main ciiference between slices and array is array was static and slices is dynamics type in other word is slices can increase their own size
	fmt.Println(skills)
	skills2 := append(skills, "PuNefuckup", "WTF", "are")
	fmt.Println(skills,skills2)

	s1 := skills[0:2]// the reason we called this slices because wes accually can slices these in to peices

	fmt.Println(s1)
	fmt.Println(skills[:2])
	fmt.Println(skills[1:])
	fmt.Println(skills[1:])
	fmt.Println(skills[:len(skills)-1])
	var ss []int
	fmt.Printf("%#v\n",ss)// zero values of slices is nil == null value
	// if you tried to access [0] or [>n] you will get error
	ss = append(ss,44)
	fmt.Printf("%#v\n",ss)// append 44 to ss nil will gone automatically

	sss := make([]int,3)// make three slot in slices
	fmt.Printf("%#v\n",sss)
}
