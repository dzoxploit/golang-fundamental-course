package main

import "fmt"

func main() {
	var data = map[string]mahasiswa{
		"10115966": {
			"Ariel Noah",
			"4IA20",
		},
		"10115961": {
			"Andhika Kangen Band",
			"4KA01",
		},
	}
	var search string
	fmt.Print("Masukan Npm anda ? ")
	fmt.Scanf("%s", &search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \nKelas anda %s", value.nama, value.kelas)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}

type mahasiswa struct {
	nama  string
	kelas string
}
