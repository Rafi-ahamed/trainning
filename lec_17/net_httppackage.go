package main

import (
	"fmt"
	"net/http"
)

//type Handler interface{
//ServeHTTP(ResponseWriter,*Request)
//}
func main() {
	//var name datatype
	// var x string
	// var handler func(ResponseWriter,*Request)
	http.Handle("/", http.FileServer(http.Dir("./img")))
	//http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `this is my first golang webpage`)
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `welcome to my about page`)
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "this is my first contact page<img src=\"capture.jpg\">")
}
