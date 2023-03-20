package main

import (
	"encoding/json"
	"fmt"
)

type User struct { // Name need to be capital
	Name       string   `json:"name"`
	Surname    string   `json:"surename"`
	Level      string   `json:"level"`
	Money      float64  `json:"money"` // fix name when convert
	Languagues []string `json:"languagues"`
}

func main() { //convert struct to json

	c := User{
		Name:       "nataphan",
		Surname:    "rakvidhyasatra",
		Level:      "newbie",
		Money:      0,
		Languagues: []string{"c", "python", "javascript", "java"},
	}

	b, err := json.Marshal(c)
	fmt.Printf("type: %T\n", b)
	fmt.Printf("byte: %v\n", b)
	fmt.Printf("string: %s\n", b)
	fmt.Printf("err: %#v\n", err)
	
	jsonData := []byte(`{"Name":"nataphan","Surname":"rakvidhyasatra","Level":"newbie","Money":0,"Languagues":["c","python","javascript","java"]}`)
	
	var u User;
	
	err = json.Unmarshal(jsonData,&u)
	fmt.Printf("%#v\n",u)
	fmt.Printf("err: %#v\n",err)
}
