package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO:
//  - Import match.go module into the main fn so we can watch strings.
//  - handle user input from the frontend and match.
//  - return the result of the match back to the frontend.

// this package is Family inspired

func main() {

	// fmt.Println(match.match())
	
}


// retrieve the user form data from the frontend >
// match with Family.



// Password type
type Password struct {
	password string `json:password`
}

func testPassword(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var password Password
	json.Unmarshal(reqBody, &password)
}

func router() {
	s := mux.NewRouter()
	s.HandleFunc("/check", testPassword).Method("POST")
}


