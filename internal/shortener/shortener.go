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

func New(strg storage.Storage) *Service {
	return &Service{
		storage: strg,
	}
}

func (s *Service) MakeShortPath(BaseURL string, body []byte) string {
	OriginalURL := strings.TrimSpace(string(body))
	ShortURL := GetRandomURL()
	s.storage.SetValue(ShortURL, OriginalURL)
	url := fmt.Sprintf("%s/%s", BaseURL, ShortURL)

	return url
}

func (s *Service) GetOriginalURL(ShortURL string) (string, bool) {
	OriginalURL, ok := s.storage.GetValue(ShortURL)
	return OriginalURL, ok
}

func GetRandomURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
