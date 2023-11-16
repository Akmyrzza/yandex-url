package storage

type Storage struct {
	StorageURL map[string]string
}

func New() *Storage {
	return &Storage{
		StorageURL: make(map[string]string),
	}
}

func (s *Storage) SetValue(shortURL, OriginalURL string) {
	s.StorageURL[shortURL] = OriginalURL
}

func (s *Storage) GetValue(shortURL string) (string, bool) {
	originalURL, ok := s.StorageURL[shortURL]
	return originalURL, ok
}
