package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "hello ji %s", name)
}
