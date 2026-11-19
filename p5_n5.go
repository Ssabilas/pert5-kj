package main

import (
	"fmt"
)

// TODO: Definisikan struct Counter dengan field count
type Counter struct {
	count int
}

// TODO: Tambahkan method "Increment" dengan value receiver yang menambah counter tetapi tidak memodifikasi original
func (c Counter) Increment() {
	c.count++
	// Modifikasi ini hanya mempengaruhi salinan, bukan yang asli
}

// TODO: Tambahkan method "IncrementByPointer" dengan pointer receiver yang menambah counter dan memodifikasi original
func (c *Counter) IncrementByPointer() {
	c.count++
	// Modifikasi ini mempengaruhi yang asli
}

// TODO: Tambahkan method "Value" yang mengembalikan nilai count saat ini
func (c Counter) Value() int {
	return c.count
}

//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	c := Counter{count: 0}
	fmt.Println("Initial counter:", c.Value()) // Seharusnya mencetak: Initial counter: 0
	
	// Panggil Increment dengan penerima nilai
	c.Increment()
	fmt.Println("After Increment:", c.Value()) // Seharusnya mencetak: After Increment: 0
	
	// Panggil IncrementByPointer dengan penerima pointer
	c.IncrementByPointer()
	fmt.Println("After IncrementByPointer:", c.Value()) // Seharusnya mencetak: After IncrementByPointer: 1
	
	// Tes dengan pointer eksplisit
	counter := &Counter{count: 10}
	counter.Increment()
	fmt.Println("Pointer counter after Increment:", counter.Value()) // Seharusnya tetap mencetak: Pointer counter after Increment: 10
	
	counter.IncrementByPointer()
	fmt.Println("Pointer counter after IncrementByPointer:", counter.Value()) // Seharusnya mencetak: Pointer counter after IncrementByPointer: 11
}