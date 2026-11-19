package main

import (
	"fmt"
)

// TODO: Definisikan struct Rectangle dengan field width dan height
type Rectangle struct {
	width  int
	height int
}

// TODO: Tambahkan method "Area" untuk menghitung luas persegi ke struct Rectangle
func (r Rectangle) Area() int {
	return r.width * r.height
}

// TODO: Tambahkan method "Perimeter" ke struct Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

// TODO: Tambahkan method "Scale" ke struct Rectangle yang mengalikan dimensi (width dan height) dengan faktor (konstanta)
func (r *Rectangle) Scale(factor int) {
	r.width *= factor
	r.height *= factor
}

//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	r := Rectangle{width: 5, height: 3}
	fmt.Println("Rectangle:", r)
	
	area := r.Area()
	fmt.Println("Area:", area) // Seharusnya mencetak: Area: 15
	
	perimeter := r.Perimeter()
	fmt.Println("Perimeter:", perimeter) // Seharusnya mencetak: Perimeter: 16
	
	r.Scale(2)
	fmt.Println("After scaling:", r) // Seharusnya mencetak: After scaling: {10 6}
	fmt.Println("New area:", r.Area()) // Seharusnya mencetak: New area: 60
}