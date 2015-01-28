package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	file, err := ioutil.ReadFile("./responses/template.json")
	checkErr(err, "File read error")
	w.Write(file)
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
