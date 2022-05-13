package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandler)
	fmt.Println("Server rodando em :8080...")
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body = []byte("")
	}

	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Echo    string `json:"echo"`
	}{
		Success: true,
		Message: "Testando 1234 !!!",
		Echo:    string(body),
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(encoded))
}
