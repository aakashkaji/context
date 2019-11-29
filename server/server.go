package main

import (
	"net/http"
	"time"
	"fmt"
	"sms/log"
	"flag"
)

// do some context handling on the server
func main() {

	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	//ctx = context.WithValue(ctx, int(42), int64(100))

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "hellow aakash")
	case <- ctx.Done():
		log.Println(ctx,  ctx.Err().Error())


	}
}
