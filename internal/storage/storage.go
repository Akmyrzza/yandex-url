package storage

type Storage struct {
	StorageURL map[string]string
}

func New() *Storage {
	return &Storage{
		StorageURL: make(map[string]string),
	}
}
