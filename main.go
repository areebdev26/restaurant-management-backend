package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := http.NewServeMux()

	http.ListenAndServe(port, router)

}
