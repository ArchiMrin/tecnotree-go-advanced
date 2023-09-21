package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleDefault) // this is expecting the url
	fmt.Println("server running on port 4001")
	http.ListenAndServe(":4001", nil)
	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		fmt.Println("Server up & running")
	} else {
		fmt.Println("Server not running")
	}
}
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Welcome to http in GO</h1>")
}
