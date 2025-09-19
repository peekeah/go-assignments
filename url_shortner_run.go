package main

import "fmt"

func RunURLShortner() {
	// Initilize
	urlShortner := InitilizeURLShortner()

	// 1. Shorten URL
	testUri := "https://google.com"
	shortenUri := urlShortner.ShortenUrl(testUri)
	fmt.Printf("shorten URL for %s is %s\n", testUri, shortenUri)

	// 2. Get shorten url
	originalUrl, err := urlShortner.GetOriginalURL(shortenUri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("orignial URL:", originalUrl)
}
