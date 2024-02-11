package storage

type URLShortener struct {
	URLs map[string]string
}

func New() *URLShortener {
	return &URLShortener{
		URLs: make(map[string]string),
	}
}
