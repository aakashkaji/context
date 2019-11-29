package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
)

// do some context handling on the server
func main() {


	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log.Printf("handler started")
	defer log.Printf("handler ended")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "hellow aakash")
	case <- ctx.Done():
		log.Print(ctx.Err())


	}
}
