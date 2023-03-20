package main

import "fmt"
import "github.com/NRdeARK/cinema-go-project/movie"
import "github.com/NRdeARK/cinema-go-project/ticket"

//start by "go mod init"
//then "go mod tidy"


func main() {
	movieName := movie.FindNameByID("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName,8.4)
	ticket.BuyTicket(movieName)
}