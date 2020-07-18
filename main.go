package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Start Server(8888)...")
	http.ListenAndServe(":8888", nil)
}
