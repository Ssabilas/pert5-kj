package main

import (
	"fmt"
)

// TODO: Implementasikan fungsi "processFile" yang mensimulasikan pembukaan dan penutupan file
// menggunakan defer untuk memastikan file ditutup
func processFile(filename string) {
	fmt.Println("Opening file:", filename)
	
	// Menunda penutupan file sampai fungsi selesai
	defer fmt.Println("Closing file:", filename)
	
	// Memproses file
	fmt.Println("Processing file...")
	
	// Operasi lain akan diletakkan di sini
	// Walaupun ada terjadi error di bawah nanti, fungsi yang di-defer tetap akan dipanggil
}

// TODO: Lengkapi fungsi "executeTransaction" yang mensimulasikan transaksi database dengan keyword defer

func executeTransaction(userId string, success bool) {
	fmt.Println("Starting transaction for user:", userId)
	
	defer fmt.Println("Transaction finished")
	
	// Connect to the database
	fmt.Println("Connected to database")
	defer fmt.Println("Disconnected from database")
	
	// Begin transaction
	fmt.Println("Beginning transaction")
	
	// Register rollback only if needed
	if !success {
		defer fmt.Println("Rolling back transaction")
	}
	
	// Process transaction
	fmt.Println("Processing transaction...")
	
	// Commit only if successful
	if success {
		fmt.Println("Committing transaction")
		defer fmt.Println("Transaction committed")
	} else {
		fmt.Println("Transaction failed")
		return // Walaupun ini kita lakukan early return, seluruh perintah defer akan tetap dijalankan.
	}
}

//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	// Panggil processFile untuk melihat defer beraksi
	processFile("example.txt")
	// Seharusnya mencetak:
	// Opening file: example.txt
	// Processing file...
	// Closing file: example.txt
	
	// Panggil executeTransaction dengan parameter sukses
	fmt.Println("\nSuccessful transaction:")
	executeTransaction("user123", true)
	// Seharusnya mencetak urutan operasi dan cleanup tanpa rollback:
	// 	Successful transaction:
	// Starting transaction for user: user123
	// Connected to database
	// Beginning transaction
	// Processing transaction...
	// Committing transaction
	// Transaction committed
	// Disconnected from database
	// Transaction finished
	
	// Panggil executeTransaction dengan parameter gagal
	fmt.Println("\nFailed transaction:")
	executeTransaction("user456", false)
	// Seharusnya mencetak urutan operasi dan rollback:
	// Failed transaction:
	// Starting transaction for user: user456
	// Connected to database
	// Beginning transaction
	// Processing transaction...
	// Transaction failed
	// Rolling back transaction
	// Disconnected from database
	// Transaction finished
}