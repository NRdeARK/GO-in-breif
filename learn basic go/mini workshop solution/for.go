package main

import "fmt"

func main() {
	// Workshop: for, for-range
	// กำหนด: 1. ให้ใช้ for loop ทำการเปลี่ยนค่าในอาร์เรย์ genres โดยเพิ่มคำนำหน้า "Movie: " ทุกค่า ดังนี้ "Movie: Action", "Movie: Adventure", "Movie: Fantasy"
	// 	  2. ให้แสดงผลค่า genres ทีละค่า โดยใช้ for-range

	// Output:
	// before for loop: [3]string{"Action", "Adventure", "Fantasy"}
	// after  for loop: [3]string{"Movie: Action", "Movie: Adventure", "Movie: Fantasy"}
	// genres[0]: Movie: Action
	// genres[1]: Movie: Adventure
	// genres[2]: Movie: Fantasy

	// TODO: write code below.
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genres)
	// for loop here.
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie: " + genres[i]
	}
	fmt.Printf("after  for loop: %#v\n", genres)

	// for-rage here.
	for _,val := range genres {
		fmt.Println(val)
	}
}