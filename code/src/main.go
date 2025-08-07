package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Go application!");
  });
    
  log.Fatal(http.ListenAndServe(":8080", nil));
}
