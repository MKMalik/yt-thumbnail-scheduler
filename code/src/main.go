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

// https://excalidraw.com/#room=67ce4b2c35022cf242fa,YfqrTC_EZTsNXaALnYH6Aw
