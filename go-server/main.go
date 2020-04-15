package main

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// func main() {
// Using Gin
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "I am using Gin !!",
// 		})
// 	})
// 	r.Run()
// }

// // Using Gorilla-mux
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/time", GetTime)
	http.ListenAndServe(":8000", r)
}

//GetTime function
func GetTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	// fmt.Println(r.Header)
	// fmt.Println(r.Proto)
	fmt.Fprintf(w, "Current time is %s", time.Now())
}
