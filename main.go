package main

import (
	"fmt"
	"math/rand"
	"time"
)

var db = make(map[string]string)

func getOriginalUrl(key string) string {
	return db[key]
}

func randString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func shorten(url string) string {
	key := randString(5)
	db[key] = url
	return key
}

func main() {
	shortenKey := shorten("http://google.com")
	fmt.Println(shortenKey)
	fmt.Println(getOriginalUrl(shortenKey))
}
