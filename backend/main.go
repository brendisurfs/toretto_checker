package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/antzucaro/matchr"

	"github.com/rs/cors"
)

// TODO:
//  - Import match.go module into the main fn so we can watch strings.
//  - handle user input from the frontend and match. DONE
//  - FIXME:return the result of the match back to the frontend.

// this package is Family inspired


func main() {
	router()
}

// http router
func router() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/check", TestPassword)

	handler := cors.Default().Handler(mux)

http.ListenAndServe(":4000", handler)
}

// HomePage func
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}


// TestPassword func
func TestPassword(w http.ResponseWriter, r *http.Request) {
	// parses raw string data.
	r.ParseForm()
	log.Println(r.Form)

	for key := range r.Form {
		returnMatch := match(key)
		fmt.Fprintf(w, "%v", returnMatch)
	}
}

// match func: match the http POST request password to Family.
func match(s1 string) float64 {
	s2 := "Family"
	r := matchr.Jaro(s1, s2)
	roundMatch := math.Round(r * 100)/100
	return roundMatch
}