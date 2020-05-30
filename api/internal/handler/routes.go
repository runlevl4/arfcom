package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// GitCommit stores the hash from HEAD.
var GitCommit string
const content = "Content-Type"
const contentJSON = "application/json"

//Health responds if the service is ready to take traffic.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "HEALTHY"}`)	
}

//Info responds if the service is up.
func Info(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "UP"}`)
}

// Chili tries to answer the age-old question on Arfcom...beans or no beans.
func Chili(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	var s string

	w.Header().Set(content, contentJSON)
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

// Fu speaks to our infamous alum.
func Fu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	t := `{"message": "FU AROCK!"}`
	fmt.Fprintf(w, t)
	fmt.Fprintf(w, "\n")
}

// Caliber tries to steer the caller to the right pistol choice.
func Caliber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusInternalServerError)
	t := `{"error": "divide by zero"}`
	fmt.Fprintf(w, t)
	fmt.Fprintf(w, "\n")
}
