package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/time", GetTime)
	http.ListenAndServe(":8000", r)
}

//GetTime function
func GetTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time is %s", time.Now())
}
