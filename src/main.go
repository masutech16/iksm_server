package main

import (
	"log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* private settings data (ex:port number, iksm_session) */
type Settings struct {
	Port string "json:'port'"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	bytes, err := ioutil.ReadFile("../settings.json")
	if err != nil {
		log.Fatal(err)
	}

	//decode
	var settings Settings
	if err := json.Unmarshal(bytes, &settings); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+settings.Port, nil)
}
