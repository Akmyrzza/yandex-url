package shortener

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/akmyrzza/yandex-url/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewShortener(strg storage.Storage) *Service {
	return &Service{
		storage: strg,
	}
}

func (s *Service) MakeShortPath(baseURL string, body []byte) string {
	originalURL := strings.TrimSpace(string(body))
	shortURL := getRandomString()
	s.storage.SetValue(shortURL, originalURL)
	url := fmt.Sprintf("%s/%s", baseURL, shortURL)

	return url
}

func (s *Service) GetOriginalURL(ShortURL string) (string, bool) {
	OriginalURL, ok := s.storage.GetValue(ShortURL)
	return OriginalURL, ok
}

func getRandomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
