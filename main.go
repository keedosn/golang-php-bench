package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	data, _ := ioutil.ReadFile("php-app/data.txt")
	ioutil.WriteFile("php-app/data-go.txt", data, 0666)

	duration := time.Since(start)

	log.Printf("getting %s took %s ms", r.URL.EscapedPath(), duration)
	w.Write([]byte("getting " + r.URL.EscapedPath() + " took " + string(duration) + "ms"))

}
