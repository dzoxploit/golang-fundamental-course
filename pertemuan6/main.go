package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/matakuliah", user)
	fmt.Println("Start server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

type matakuliah struct {
	Nama  string `json:"nama_mata_kuliah"`
	Nilai int    `json:"nilai_mata_kuliah"`
}

var data_mata_kuliah = []matakuliah{
	{
		Nama:  "Golang",
		Nilai: 100,
	},
	{
		Nama:  "Java",
		Nilai: 100,
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mata_kuliah)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}
