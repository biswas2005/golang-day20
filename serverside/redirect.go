package serverside

import "net/http"

func Redirect() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/redirect", redirect)
	http.ListenAndServe(":8080", nil)

}

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello welcome"))
}
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}
