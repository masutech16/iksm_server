package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type settings struct {
	Port string `json:"port"`
	Iksm string `json:"iksm"`
}

// Data has config data
var Data = new(settings)

// Init should be called before using this package
func Init() {
	bytes, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, Data); err != nil {
		log.Fatal(err)
	}
}
