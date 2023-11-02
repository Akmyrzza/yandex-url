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

	longUrl := string(body)
	shortUrl := GetRandomURL()
	urlStorage[shortUrl] = longUrl

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + r.Host + "/" + shortUrl))
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
	shortUrl := r.URL.Path[1:]
	longUrl, ok := urlStorage[shortUrl]
	if ok {
		w.Header().Set("Location", longUrl)
		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}
}
