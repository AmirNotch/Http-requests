package main

import (
	"fmt"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request){
//	w.Header().Set("RequestID", "213124123dadw")
//
//	fmt.Fprintln(w, "You browser is", r.UserAgent())
//	fmt.Fprintln(w, "You accept ", r.Header.Get("Accept"))
//}
func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at: 8080")
	http.ListenAndServe(":8080",nil)
}