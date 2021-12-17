package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var newCache *cache.Cache

func init() {
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	newCache.Set("Simy", "Zoovie", cache.DefaultExpiration)
}
func getFromCache(w http.ResponseWriter, r *http.Request) {
	Simy, found := newCache.Get("Simy")
	if found {
		log.Print("Key found in cache with value as :: ", Simy.(string))
		fmt.Fprint(w, "Hello "+Simy.(string))
	} else {
		log.Print("Key not found in cache :: ", "Simy")
		fmt.Fprint(w, "key not found in cache.")
	}
}

func main() {
	http.HandleFunc("/", getFromCache)
	log.Println("Starting server at Port 8080...")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
		return
	}
}
