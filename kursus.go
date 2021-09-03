package main

import "fmt"

func main() {
	var kursus []string

	kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
	kursus = append(kursus[0:4], "Managemen Informatika")

	fmt.Println("Isi dari Kursus adalah : ", kursus)
	fmt.Printf("Panjang kursus %d dan kapasitas %d\n", cap(kursus), cap(kursus))

	kursus_saya := kursus[1:4]
	kursus_saya = append(kursus_saya, "Managemen Informatika")

	fmt.Println("Isi dari Kursus saya adalah : ", kursus_saya)
	fmt.Printf("Panjang kursus %d dan kapasitas %d\n", cap(kursus_saya), cap(kursus_saya))

}
