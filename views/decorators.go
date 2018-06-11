package views

import (
	"net/http"
	"log"
	"runtime/debug"
)

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// capture panic error
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("panic in %v: %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()

		// call real handler
		fn(w, r)
	}

}
