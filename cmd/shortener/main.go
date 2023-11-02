package main

import (
	"io"
	"math/rand"
	"net/http"
)

var urlStorage = make(map[string]string)

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
		SaveURL(w, r)
	case "GET":
		ReturnURL(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SaveURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	longURL := string(body)
	shortURL := GetRandomURL()
	urlStorage[shortURL] = longURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + r.Host + "/" + shortURL))
}

func GetRandomURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, rand.Intn(8)+1)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ReturnURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	longURL, ok := urlStorage[shortURL]
	if ok {
		w.Header().Set("Location", longURL)
		http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}
}
