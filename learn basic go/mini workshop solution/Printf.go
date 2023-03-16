package main

import "fmt"

func main() {
	// Workshop: Printf

	// Output:
	// เรื่อง: Avengers: Endgame
	// ปี: 2019
	// เรตติ้ง: 8.4
	// ประเภท: Sci-Fi
	// ซุปเปอร์ฮีโร่: true
	// เรื่องโปรด: ⭐

	// TODO: write code below.
	key,value := "เรื่อง:", " Avengers: Endgame"
	fmt.Printf("%s%s\n",key,value)

	key,valueD := "ปี:", 2019
	fmt.Printf("%s%d\n",key,valueD)
	
	key,valueF := "เรตติ้ง:", 8.4
	fmt.Printf("%s%.1f\n",key,valueF)
	
	key,value = "ประเภท:", " Sci-Fi"
	fmt.Printf("%s%s\n",key,value)
	
	key,valueB := "ซุปเปอร์ฮีโร่:", true
	fmt.Printf("%s%t\n",key,valueB)
	
	key,valueC := "เรื่องโปรด:", " ⭐"
	fmt.Printf("%s%s\n",key,valueC)

}