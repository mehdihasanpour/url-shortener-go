package main

import (
	"math/rand"
	"net/http"
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
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("only get request supported"))
			return
		}
		url := r.URL.Query().Get("url")
		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("need a url to short"))
			return
		}
		key := shorten(url)
		toClickURL := "localhost:8080/long?key=" + key
		w.Write([]byte(toClickURL))
	})

	http.HandleFunc("/long", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("need a url to short"))
			return
		}
		url := getOriginalUrl(key)
		w.Write([]byte(url))
	})

	http.ListenAndServe(":8080", nil)
}
