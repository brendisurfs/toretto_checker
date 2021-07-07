package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HomePage func
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}


// Password type
type Password struct {
	password string `json:password`
}

// TestPassword func
func TestPassword(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var password Password
	json.Unmarshal(reqBody, &password)
}