package controller

import "net/http"

func new() {
	http.HandleFunc("/products")
}
