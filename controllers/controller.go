package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, você está na HomePage!")
}
