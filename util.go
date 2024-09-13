package main

import (
	"crypto/sha1"
	"encoding/base32"
    "math/rand"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortUrl() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 6)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func HashUrl(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	return base32.StdEncoding.EncodeToString(hash.Sum(nil))[:6]
}