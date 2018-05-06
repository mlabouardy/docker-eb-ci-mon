package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Avenger struct {
	Character string `json:"character"`
	Name      string `json:"name"`
}

var avengers []Avenger

func init() {
	data, _ := ioutil.ReadFile("avengers.json")
	json.Unmarshal(data, &avengers)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(avengers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
