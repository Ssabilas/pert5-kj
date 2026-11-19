package main

import (
	"fmt"
)

// TODO: Implementasikan fungsi "operate" yang menerima dua bilangan bulat dan
// parameter fungsi yang mendefinisikan operasi yang akan dilakukan pada keduanya
func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}


//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	// Gunakan operate dengan penjumlahan
	result := operate(5, 3, func(a, b int) int {
		return a + b
	})
	fmt.Println("5 + 3 = ", result) // Seharusnya mencetak: 5 + 3 = 8
	
	// Gunakan operate dengan perkalian
	result = operate(5, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println("5 * 3 = ", result) // Seharusnya mencetak: 5 * 3 = 15
	
	//! Kode di atas tidak boleh dimodifikasi.

	// TODO: Buat fungsi anonim baru untuk operasi pengurangan dan panggil dengan operate
	result = operate(5, 3, func(x, y int) int {
		return x - y    // modifikasi return value sehingga menjadi operasi pengurangan
	})
	fmt.Println("5 - 3 = ", result)
}