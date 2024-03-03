package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"urlshortener/internal/app/config"
)

type URLShortener struct {
	URLs map[string]string
}

type Event struct {
	UUID        string `json:"uuid"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}

func New(cfg *config.Config) *URLShortener {
	Urls := make(map[string]string)

	// restore urls from file
	FileUrls, err := RestoreFromFile(cfg.FileStoragePath)
	if err == nil && FileUrls != nil {
		Urls = FileUrls
	}

	return &URLShortener{
		URLs: Urls,
	}
}

func RestoreFromFile(filename string) (map[string]string, error) {
	URLs := make(map[string]string)

	if filename == "" {
		return nil, errors.New("data-file not found")
	}

	// create file in not exist
	_, err := os.Stat(filename)
	if err != nil {
		f, _ := os.Create(filename)
		f.Close()
		return nil, err
	}

	// read file if exist
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		// parse line and update map
		var event Event
		err := json.Unmarshal(line, &event)
		if err != nil {
			continue
		}

		URLs[event.ShortUrl] = event.OriginalUrl

	}

	if err = scanner.Err(); err != nil {
		return URLs, err
	}

	return URLs, nil
}

func (u URLShortener) UpdateFile(filename string, key string) error {
	if v, ok := u.URLs[key]; ok {
		e := Event{
			UUID:        fmt.Sprintf("%v", len(u.URLs)),
			ShortUrl:    key,
			OriginalUrl: v,
		}

		r, err := json.Marshal(e)
		if err != nil {
			return err
		}

		r = append(r, '\n')

		file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0o644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(r)
		return err
	} else {
		return errors.New("key not found")
	}
}
