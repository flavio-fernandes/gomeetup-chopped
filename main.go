package main

import (
	"io/ioutil"
	"log"
	"net/http"
	//	"crypto/rsa"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %s", err)
	}
	response := []byte(`You said to "`)
	response = append(response, body...)
	response = append(response, '"', '\n')
	w.Write(response)
}
