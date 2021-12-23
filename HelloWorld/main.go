package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)


func main (){

	fmt.Println("hello golphers");

	u := models.User {
		ID: 90,
		FirstName: "Marcia",
		LastName: "Peterson",
	}

	fmt.Println(u)

}

