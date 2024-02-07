package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Terms struct {
	Term1 int `json:"term1"`
	Term2 int `json:"term2"`
}

var operations = make([]string, 0)
var mu = sync.Mutex{}

func addition(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	switch r.Method {
	case http.MethodGet:
		var err error
		if len(operations) == 0 {
			_, err = w.Write([]byte("Your map is empty\n"))
		} else {
			res := strings.Join(operations, " ")
			_, err = w.Write([]byte(res))
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodPost:
		pair := &Terms{Term1: 1, Term2: 2}
		err := json.NewDecoder(r.Body).Decode(pair)
		if err != nil {
			fmt.Println("here")
			log.Fatal(err)
			return
		}

		res := pair.Term1 + pair.Term2
		operations = append(operations, strconv.Itoa(res))
		_, err = w.Write([]byte(operations[len(operations)-1]))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
