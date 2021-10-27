package main

type Segitiga struct {
	alas, t float64
}

func (k Segitiga) Luas() float64 {
	return k.alas * k.t / 2
}

func (k Segitiga) Keliling() float64 {
	return k.alas * 3
}
