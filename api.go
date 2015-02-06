package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mojotech/feedbag/feedbag/template"
)

var (
	port = flag.String("port", "3000", "Port to run mock server on")
)

func main() {
	flag.Parse()
	http.HandleFunc("/api/activity", activityHandler)
	http.HandleFunc("/templates", templateHandler)
	panic(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./responses/activity.json")
	checkErr(err, "File read error")
	w.Write(file)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseDir("./templates")
	checkErr(err, "Failed to read templates")

	res, err := json.Marshal(templates)
	checkErr(err, "Failed to marshal templates")

	w.Write(res)
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
