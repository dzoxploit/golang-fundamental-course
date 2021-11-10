package gajisederhana

import "fmt"

func naikanGaji(gajiAwal int, gajiAkhir int) int {
	gajiAwal = gajiAkhir
}
func main() {
	var gajiSekarang, ekspektasiGaji int
	fmt.Print("Masukan gaji anda : ")
	fmt.Scan(&gajiSekarang)
	fmt.Print("Masukan gaji yang anda inginkan : ")
	fmt.Scan(&ekspektasiGaji)

	naikanGaji(gajiSekarang, ekspektasiGaji)

	fmt.Printf("\nGaji anda sekarang %d\n", ekspektasiGaji)

}
