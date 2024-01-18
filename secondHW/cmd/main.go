package main

import (
	"log"
	"net/http"
	"secondHW/pkg/api"
)

func main() {
	myApi := api.New("0.0.0.0:8080", http.NewServeMux())
	myApi.FillEndpoints()
	log.Fatal(myApi.ListenAndServe())
}
