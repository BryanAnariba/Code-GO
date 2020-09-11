package main
import (
	"net/http"
)

func main () {
	//println("Hola mundo")

	// Routes
	http.HandleFunc("/" , homeRoute)
	http.HandleFunc("/contact" , contactRoute)

	// Starting server
	http.ListenAndServe(":3500" , nil)
}

func homeRoute(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("Hello World"))
}

func contactRoute(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("Contact Works"))
}
