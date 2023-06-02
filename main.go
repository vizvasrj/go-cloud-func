package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// var clientOnce sync.Once
// var client *storage.Client

// func init() {
// 	functions.HTTP("LazyGlobal", LazyGlobal)
// }

// func LazyGlobal(w http.ResponseWriter, r *http.Request) {
// 	clientOnce.Do(func() {
// 		var err error
// 		client, err = storage.NewClient(context.Background())
// 		if err != nil {
// 			http.Error(w, "internal error", http.StatusInternalServerError)
// 			log.Printf("storage.NewClient: %v", err)
// 			return
// 		}
// 	})
// }

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pong)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listning on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	curl, err := http.Get("https://ipinfo.io")
	if err != nil {
		fmt.Fprintf(w, "error?? %s", err.Error())
		return
	}
	defer curl.Body.Close()
	body, err := ioutil.ReadAll(curl.Body)
	if err != nil {
		fmt.Fprintf(w, "error?? %s", err.Error())
		return
	}
	fmt.Fprintf(w, "hello ji bonjour %s", string(body))
}

func pong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(15 * time.Second)
	fmt.Fprintf(w, "pong")
}
