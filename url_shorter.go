package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type URL struct {
	originalURL  string
	shortenedURL string
}

type URLStore map[string]string

func InitilizeURLShortner() *URLStore {
	return &URLStore{}
}

func (s *URLStore) ShortenUrl(url string) string {
	randId := rand.Intn(900000) + 100000
	shortenedURL := fmt.Sprintf("https://example.com/%d", randId)
	(*s)[shortenedURL] = url
	return shortenedURL
}

func (s *URLStore) GetOriginalURL(shortenedUrl string) (string, error) {
	url, ok := (*s)[shortenedUrl]

	if !ok {
		return "", errors.New("url not found")
	}
	return url, nil
}
