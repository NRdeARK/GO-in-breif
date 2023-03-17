package main

import "fmt"

func test(slices []string) {// array param that is used need to have the same size
	fmt.Println("Hello")
}

func show(tag string, sk[]string) {
	// slices have 2 number len(lenght) and cap(capacity)
	// len >> how many things in here are
	// cap >> how many seat is left that have create
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s: len: %d cap: %d -- show: %#v\n", tag, l, c, sk)
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
	// always aware that slices is pass by ref that ref from array
	// so if you have slices "A" that link from another slices "B"
	// modifing "B" will affect "A"
	
	show("s1",s1)//this happen because it will get cap from original slices

	s2 := skills[:]
	show("s2",s2)//this happen beacause where it start is index 1 it will remove index 0 so cap will see as 2 

	s2 = append(s2, "append")// from this when append is called it will go reserve new arr themn replace s2 with with new ref from new arr

	fmt.Println()
	show("skills",skills)
	show("s2",s2)

	s2[1] = ""
	show("skills",skills)// your will see that nothing is affected when s2 is been modified beacuse it is difference ref
	show("s2",s2)

}
