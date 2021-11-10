package main

import (
	"didinnuryahya_51419788_pert4/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handler.API)

	log.Println("localhost : 8088")
	http.ListenAndServe(":8088", nil)
}
