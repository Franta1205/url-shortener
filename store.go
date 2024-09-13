package main

import (
	"encoding/json"
    "os"
)

const storeFile = "urls.json"

type URLStore struct {
	Urls map[string]string
}


func NewURLStore() (*URLStore, error) {
	store := &URLStore{Urls: make(map[string]string)}

	if err := store.load(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *URLStore) load() error {
	file, err := os.Open(storeFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&s.Urls)
}

func (s *URLStore) save() error {
	file, err := os.Create(storeFile)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(s.Urls)
}

func (s *URLStore) SaveUrl(shortUrl, longUrl string) {
	s.Urls[shortUrl] = longUrl
	s.save()
}

func (s *URLStore) GetUrl(shortUrl string) (string, bool) {
	longUrls, exists := s.Urls[shortUrl]
	return longUrls, exists
}
