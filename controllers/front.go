package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
RegisterController function will register all http handels.
*/
func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
