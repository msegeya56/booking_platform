package main

import (
	"fmt"
	"net/http"

	"github.com/wander4747/booking/pkg/handlers"
)

const portNumber = ":9998"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
