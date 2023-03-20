package main

import "fmt"

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

// TODO: write code below.
func (eg *movie) addVote(r float64){
	eg.votes = append(eg.votes, r)
}


func main() {
	// Workshop: pointer of struct
	// กำหนด: 1. ให้สร้าง method addVote รับพารามิเตอร์ rating เป็น float64
	// 	  2. ให้ method addVote เพิ่มค่า rating เข้าไปใน slice votes ของ struct movie
	//
	// Output:
	// votes: [7 8 9 10 6 9 9 10 8 8]

	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	eg.addVote(8)
	fmt.Println("votes:", eg.votes)
}