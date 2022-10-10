package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Vikraam27/go-open-music-api/routes"
)

func main() {
	r := routes.Routes()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting server on the port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
