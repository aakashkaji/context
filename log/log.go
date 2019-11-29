package log

import (
	"context"
	"log"
	"net/http"
	"math/rand"
)

type key int64
const  requestIdKey  = key(42)

func Println(ctx context.Context, msg string) {

	id, ok := ctx.Value(requestIdKey).(int64)

	if !ok {
		log.Println("could not find find request id context")
		return
	}
	log.Printf("[%d], %s", id, msg)
}


func Decorate(f http.HandlerFunc) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		// it get context from request
		// receiving a context and sending back these context with value
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIdKey, id)
		f(w, r.WithContext(ctx))
	}
}
