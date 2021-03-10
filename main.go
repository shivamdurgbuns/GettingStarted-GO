package main

import (
	"fmt"

	"github.com/shivamdurgbuns/webservices/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Shivam",
		LastName:  "Durgbuns",
	}
	fmt.Println(u)
}
