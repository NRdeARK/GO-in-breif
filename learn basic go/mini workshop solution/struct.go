package main

import "fmt"

func main() {
	// Workshop: struct
	// กำหนด: 1. ให้นิยามโครงสร้างข้อมูล movie เพื่อเก็บ ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool)
	// 	  2. ให้ประกาศตัวแปรเพื่อเก็บหนัง Avengers: Endgame, ปี:2019, เรตติ้ง:8.4, ประเภท:["Action", "Drama"] และ isSuperHero:true
	//	  3. ให้ประกาศตัวแปรเพื่อเก็บหนัง Infinity War, ปี:2018, เรตติ้ง:8.4, ประเภท:["Action", "Sci-Fi"] และ isSuperHero:true
	//	  4. ให้เก็บหนังทั้งสองเรื่องไว้ในตัวแปร mvs
	// 	  5. ทำการวนลูปเพื่อแสดงผลหนังทีละเรื่อง
	//
	// Output:
	// main.movie{title:"Avengers: Endgame", year:2019, rating:8.4, genres:[]string{"Action", "Drama"}, isSuperHero:true}
	// main.movie{title:"Avengers: Infinity War", year:2018, rating:8.4, genres:[]string{"Action", "Sci-Fi"}, isSuperHero:true}

	// TODO: write code below.
	type movie struct {
		titles string
		year int64
		rating float64
		category []string
		isSuperHero bool
	}
	a := movie{
		titles : "Avengers: Endgame",
		year : 2019,
		rating : 8.4 ,
		category : []string{"Action", "Drama"},
		isSuperHero : true,
	}

	b := movie{
		titles : "Avengers: Infinity War",
		year : 2018,
		rating : 8.4 ,
		category : []string{"Action", "Sci-Fi"},
		isSuperHero : true,
	}

	var mvs []movie

	mvs = append(mvs,a,b)

	for i := 0; i < len(mvs); i++ {
		fmt.Printf("%#v",mvs[i])
	}

}
