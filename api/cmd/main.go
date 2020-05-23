package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error|shutting down|%s", err)
		os.Exit(1)
	}
}

func run() error {
	var log = log.New(os.Stdout, "arfcom|", log.LstdFlags|log.Lmicroseconds)

	r := mux.NewRouter()
	r.HandleFunc("/chili", loggingMiddleware(Chili)).Methods("GET")
	r.HandleFunc("/fu", loggingMiddleware(Fu)).Methods("GET")
	r.HandleFunc("/caliber/9mmOr45", loggingMiddleware(Caliber)).Methods("GET")

	log.Printf("%s : Starting server on port 8000", time.Now().Format(time.RFC3339))
	log.Fatal(http.ListenAndServe(":8000", r))

	return nil
}

// HandleFunc provides our logging middleware functionality.
type HandleFunc func(w http.ResponseWriter, r *http.Request)

func loggingMiddleware(handler HandleFunc) HandleFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s : %s : %s", time.Now().Format(time.RFC3339), r.Method, r.URL)
		handler(w, r)
	}
	return fn
}

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

// Fu speaks to our infamous alum.
func Fu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t := `{"message": "FU AROCK!"}`
	fmt.Fprintf(w, t)
	fmt.Fprintf(w, "\n")
}

// Caliber tries to steer the caller to the right pistol choice.
func Caliber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	t := `{"error": "divide by zero"}`
	fmt.Fprintf(w, t)	
	fmt.Fprintf(w, "\n")
}
