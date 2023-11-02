package main

import (
	"net/http"

	"github.com/akmyrzza/yandex-url/internal/handler"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Handle)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic(err)
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handler.SaveURL(w, r)
	case "GET":
		handler.ReturnURL(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
