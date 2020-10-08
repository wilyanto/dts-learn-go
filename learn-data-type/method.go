package main

import "fmt"

type Sisi struct {
	Panjang int
	Lebar int
}

func main() {
	s := Sisi{3,4}
	s.ScaleUp(2)
	fmt.Println(s.Luas())
}

// Cara 1 u/ menggunakan method
// func Luas(s Sisi) int {
// 	return s.Panjang * s.Lebar
// }

// Cara 2 u/ menggunakan method
func (s Sisi) Luas() int {
	return s.Panjang * s.Lebar
}

// Merubah nilai LANGSUNG ke alamat variabel menggunakan pointer
func (s *Sisi) ScaleUp(i int) {
	s.Lebar = s.Lebar*i
	s.Panjang = s.Panjang*i
}