package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var response []Avenger

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func init() {
	data, _ := ioutil.ReadFile("avengers.json")
	json.Unmarshal(data, &response)
}

func TestInit(t *testing.T) {
	assertEqual(t, len(response), 15)
}
