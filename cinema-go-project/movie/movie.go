package movie

import (
	"fmt"

	"github.com/NRdeARK/cinema-go-project/ticket"//cant import in cycle (cyclic dependency)
)


func Review(name string, rating float64) {// when you want to used in other file(global) func name need to start with a cap letter 
	ticket.BuyTicket(name)
	fmt.Printf("I reviewed %s and it's rating is %.2f\n", name, rating)
}

func init(){
	fmt.Println("init : movie")
}

func FindNameByID(imdbID string) string {

	switch imdbID {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	}

	return "not found."
}