package handler

import "math/rand"

type Handler struct {
	StorageURL map[string]string
}

func New() *Handler {
	return &Handler{
		StorageURL: make(map[string]string),
	}
}

func GetRandomURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
