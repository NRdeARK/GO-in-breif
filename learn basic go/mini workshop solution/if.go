package main

import (
	"fmt"
)

func main() {
	// Workshop: if
	// กำหนด: ให้แก้โจทย์นี้ด้วยการใช้ if เท่านั้น

	// Output:
	// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า "Disappointed 😞"
	// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า "Normal 😐"
	// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า "Good 🥰"
	// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า "🤷🤷🤷🤷"

	// TODO: write code below.
	ratings := 13
	if ratings < 5.0 {
		fmt.Println("Disappointed 😞")
	} else if ratings >= 5.0 && ratings < 7.0{
		fmt.Println("Normal 😐")
	} else if ratings >= 7.0 && ratings < 10.0{
		fmt.Println("Good 🥰")
	} else {
		fmt.Println("🤷🤷🤷🤷")
	}

}