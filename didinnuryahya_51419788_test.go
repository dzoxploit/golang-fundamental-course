package main

import "testing"

var (
	segitiga           Segitiga = Segitiga{5, 5}
	luasSeharusnya     float64  = 12.5
	kelilingSeharusnya float64  = 15
)

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", segitiga.Luas())
	if segitiga.Luas() != luasSeharusnya {
		t.Errorf("Salah Jawabanya!")
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", segitiga.Keliling())
	if segitiga.Keliling() != kelilingSeharusnya {
		t.Error("Salah Jawabanya")
	}
}
