package storage

type URLShortener struct {
	Urls map[string]string
}

func New() *URLShortener {
	return &URLShortener{
		Urls: make(map[string]string),
	}
}
