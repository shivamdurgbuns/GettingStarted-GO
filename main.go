package main

import (
	"net/http"

	"github.com/shivamdurgbuns/webservices/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
