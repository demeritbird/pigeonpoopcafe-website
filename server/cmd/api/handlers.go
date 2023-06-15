package main

import (
	"fmt"
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
