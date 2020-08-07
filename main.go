package main

import (
	"f1api/routes"
	"fmt"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 5000...")
	http.ListenAndServe(":5000", r)
}
