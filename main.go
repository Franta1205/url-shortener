package main

import (
	"flag"
    "fmt"
    "log"
    "os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: url-shortener [shorten|resolve] [options]")
		return
	}

	store, err := NewURLStore()

	if err != nil {
		log.Fatalf("Failed to initialize store: %v", err)
	}

	command := os.Args[1]

	switch command {
	case "shorten":
		shortenCmd := flag.NewFlagSet("shorten", flag.ExitOnError)
		url := shortenCmd.String("url", "", "URL to shorten")
		shortenCmd.Parse(os.Args[2:])

		if *url == "" {
			fmt.Println("URL is required")
            return
		}

		shortUrl := GenerateShortUrl()
		store.SaveUrl(shortUrl, *url)
		fmt.Printf("Shortened URL: http://localhost:8080/%s\n", shortUrl)

	case "resolve":
		resolveCmd := flag.NewFlagSet("resolve", flag.ExitOnError)
		shortUrl := resolveCmd.String("short", "", "Short URL to resolve")
		resolveCmd.Parse(os.Args[2:])

		if *shortUrl == "" {
			fmt.Println("Short URL is required") // Print an error message if no short URL is provided.
            return
		}

		longUrl, exists := store.GetUrl(*shortUrl)

		if !exists {
			fmt.Println("Short URL not found")
            return
		}

		fmt.Printf("Original URL: %s\n", longUrl) 
	}
}