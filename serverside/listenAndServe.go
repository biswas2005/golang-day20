package serverside

import (
	"fmt"
	"net/http"
)

func ListenServe() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", routing)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func routing(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Get response"))
	}
	if r.Method == http.MethodPost {
		w.Write([]byte("Post response"))
	}

}
