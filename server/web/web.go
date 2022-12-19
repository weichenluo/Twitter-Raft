package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "localhost:" + os.Args[1]

	fmt.Printf("Starting server at port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
