package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Chili tries to answer the age-old question on Arfcom...beans or no beans.
func Chili(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	var s string

	w.WriteHeader(http.StatusOK)

	if n%2 == 0 {
		s = "beans"
	} else {
		s = "no beans"
	}
	t := `{"answer": "%s"}`

	fmt.Fprintf(w, t, s)
	fmt.Fprintf(w, "\n")
}
