package main

import (
	"fmt"
	"log"
	"net/http"
)

type News struct {
	Title   string `bson:"Title"`
	Content string `bson:"Content"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	var str = []byte(`[
                       {"Title": "News Service Complete", "Content": "Congratulations:Your News Service Complete"},
                       {"Title": "Total Ticket System Complete", "Content": "Just a total test"}
                    ]`)

	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	fmt.Fprintf(w, string(str))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":12862", nil))
}
