package main

import "fmt"

func Pertemuan3Act1() {
	var i, j int
	for j = 0; j < 10; j++ {
		fmt.Print("Input Nilai : ")
		fmt.Scan(&i)
		if i > 10 {
			fmt.Println("Pertanyaan Selesai, karena Angka I yang diinput sudah melebihi ketentuan")
		} else if i%2 == 1 {
			fmt.Println("Ini adalah Bilangan ganjil")
		} else {
			fmt.Println("Ini adalah Bilangan genap")
		}
	}

}
