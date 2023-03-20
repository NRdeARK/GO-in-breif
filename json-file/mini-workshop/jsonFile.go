package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title       string
	Year        int
	Rating      float32
	Genres      []string
	IsSuperHero bool
}

func main() {
	// Workshop: json to struct
	// กำหนด: 1. ให้ทำการนำค่า json body ไป decode เป็น struct แล้วเก็บไว้ในตัวแปร movies
	//
	// Output:
	// []main.movie{main.movie{Title:"Avengers: Endgame", Year:2019, Rating:8.4, Genres:[]string{"Action", "Drama"}, IsSuperHero:true}, main.movie{Title:"Avengers: Infinity War", Year:2018, Rating:8.4, Genres:[]string{"Action", "Sci-Fi"}, IsSuperHero:true}}

	body := `[
  {
    "imdbID": "tt4154796",
    "title": "Avengers: Endgame",
    "year": 2019,
    "rating": 8.4,
    "genres": ["Action", "Drama"],
    "isSuperHero": true
  },
  {
    "imdbID": "tt4154756",
    "title": "Avengers: Infinity War",
    "year": 2018,
    "rating": 8.4,
    "genres": ["Action", "Sci-Fi"],
    "isSuperHero": true
  }
]`

  jsonBody := []byte(body)
  ms := []movie{}
  json.Unmarshal(jsonBody,&ms)


	// TODO: write code below.

	fmt.Printf("%#v\n", ms)
}