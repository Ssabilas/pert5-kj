package main

import (
	"fmt"
	"strings"
)

// TODO: Implementasikan fungsi "add" yang menerima dua bilangan bulat dan mengembalikan jumlahnya
func add(a, b int) int {
	return a + b
}

// TODO: Implementasikan fungsi "repeat" yang menerima string dan bilangan bulat
// dan mengembalikan string yang diulang sebanyak bilangan tersebut
func repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// TODO: Implementasikan fungsi "getDimensions" yang mengembalikan width dan height (keduanya int)
func getDimensions() (int, int) {
	return 100, 200
}

//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	// Tes fungsi add
	sum := add(5, 3)
	fmt.Println("Sum:", sum) // Seharusnya mencetak: Sum: 8
	
	// Tes fungsi repeat
	repeated := repeat("Go", 3)
	fmt.Println("Repeated:", repeated) // Seharusnya mencetak: Repeated: GoGoGo
	
	// Gunakan identifier kosong untuk mengabaikan height
	width, _ := getDimensions()
	fmt.Println("Width only:", width) // Seharusnya mencetak: Width only: 100
}