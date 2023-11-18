package storage

type Storage struct {
	StorageURL map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		StorageURL: make(map[string]string),
	}
}

func (s *Storage) SetValue(shortURL, originalURL string) {
	s.StorageURL[shortURL] = originalURL
}

func (s *Storage) GetValue(shortURL string) (string, bool) {
	originalURL, ok := s.StorageURL[shortURL]
	return originalURL, ok
}
