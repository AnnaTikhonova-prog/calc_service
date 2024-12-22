package application

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(address string) {
	http.HandleFunc("/api/v1/calculate", CalculateHandler)

	if address == "" {
		address = ":8080"
	}

	fmt.Printf("Server started at http://localhost%s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
