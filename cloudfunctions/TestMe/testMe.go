package p

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	name := v.Get("name")

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello %s, my cool friend.", name)))
}
