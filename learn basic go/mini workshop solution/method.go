package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

// TODO: write code below.
func (m movie) info() {
	fmt.Printf("%s - %.2f\n",m.title,m.rating)
	fmt.Printf("Genres:\n")
	for _,v := range m.genres{
		fmt.Println("			",v)
	}
	return
}


func main() {
	// Workshop: method
	// กำหนด: 1. ให้สร้างmethod info สำหรับ movie เพื่อเก็บแสดงผลรายละเอียด โดยประกอบด้วย ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool).
	//
	// Output:
	// Avengers: Endgame (2019) - 8.40
	// Genres:
	// 				Action
	// 				Drama

	ae := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}


	ae.info()
}
