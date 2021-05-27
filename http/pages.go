package main

import (
	"fmt"
	"net/http"
)

//func handler( w http.ResponseWriter, r *http.Request){
//	fmt.Fprintln(w,"Main Page")
//}

func main() {
	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			   fmt.Fprintln(w, "Single Page:", r.URL.String())
		})

	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			   fmt.Fprintln(w, "Multiple pages:", r.URL.String())
		})
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080",nil)
}
