package main

import (
	"fmt"
	"github.com/pluralsight/webservice"
)

func main (){
	type user struct {
		ID int 
		FirstName string
		LastName string
}
	var u user
	fmt.Println(u)

	u.ID = 90
	u.FirstName = "Alex"
	u.LastName = "Molina"
	fmt.Println(u)
	fmt.Println(u.FirstName, u.LastName)

	u2 := user {ID: 1, 
		  FirstName: "Margaret", 
		  LastName: "Tudor",
	}
	fmt.Println(u2)
}