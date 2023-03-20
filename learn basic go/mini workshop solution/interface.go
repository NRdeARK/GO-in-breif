package main

import "fmt"

// TODO: write code below.


func vote(v voter, rating float64) {
	v.addVote(rating)
}

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

type voter interface {
	addVote(float64)
}

func main() {
	// Workshop: pointer of struct
	// กำหนด: 1. ให้สร้าง interface voter เพื่อให้สามารถนำ movie มาใช้ในการโหวตได้
	//
	// Output:
	// votes: [7 8 9 10 6 9 9 10 8]

	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	vote(eg, 8)
	fmt.Println("votes:", eg.votes)
}