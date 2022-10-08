package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vikraam27/go-open-music-api/routes"
)

func main() {
	r := routes.Routes()

	fmt.Println("Starting server on the port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
