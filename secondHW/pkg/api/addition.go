package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Terms struct {
	Term1 int `json:"term1"`
	Term2 int `json:"term2"`
}

var operationsHistory = make(map[int]int)

func addition(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var err error
		if len(operationsHistory) == 0 {
			_, err = w.Write([]byte("Your map is empty\n"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		for _, value := range operationsHistory {
			_, err = w.Write([]byte(strconv.Itoa(value) + " "))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	case http.MethodPost:
		pair := &Terms{Term1: 1, Term2: 2}
		err := json.NewDecoder(r.Body).Decode(pair)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = w.Write([]byte(strconv.Itoa(pair.Term1 + pair.Term2)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		operationsHistory[len(operationsHistory)+1] = pair.Term1 + pair.Term2
	}
}
