package main

import (
	"fmt"
	"groupie-tracker/groupie"
	"net/http"
)

func main() {
	groupie.StartTheSite()
	http.HandleFunc("/", groupie.Handler)
	fmt.Print("Listen And Serve Port 8080")
	http.ListenAndServe(":8080", nil)
}
