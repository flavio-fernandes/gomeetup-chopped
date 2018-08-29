package main

import (
	//	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
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
	command := r.Header.Get("Command")
	log.Printf("Key is %s\n", body)
	response := []byte(`You said to "`)
	response = append(response, []byte(command)...)
	response = append(response, '"', '\n')
	w.Write(response)
}
