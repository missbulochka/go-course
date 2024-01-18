package api

import "net/http"

func multiplication(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
