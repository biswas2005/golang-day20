package serverside

import (
	"fmt"
	"net/http"
	"strings"
)

func ManualPath() {
	http.HandleFunc("/user/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) >= 4 {
		id := parts[2]
		name := parts[3]

		response := fmt.Sprintf("User ID:%s Name:%s", id, name)
		w.Write([]byte(response))
	} else {
		w.Write([]byte("Invalid path"))
	}
}
