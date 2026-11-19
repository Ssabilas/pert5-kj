package main

import (
	"errors"
	"fmt"
)

// TODO: Implementasikan fungsi "calculateStats" yang menerima array bilangan bulat dan
// mengembalikan nilai minimum, maksimum, dan rata-rata sebagai named return value
func calculateStats(numbers []int) (min, max int, avg float64) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	sum := 0
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}
	
	avg = float64(sum) / float64(len(numbers))
	return // Menggunakan return kosong (naked return) karena kita sudah mendeklarasikan named return value pada saat mendeklarasikan function calculateStats di line 9
}


// TODO: Implementasikan fungsi "createUser" yang menerima firstName, lastName (keduanya bertipe string) dan age (int)
// dan mengembalikan formatted string (perhatikan kode di bawah menggunakan fmt.Printf)
func createUser(firstName, lastName string, age int) string {
	return fmt.Sprintf("%s %s (%d)", firstName, lastName, age)
}

// TODO: Implementasikan fungsi "divide" yang menerima dua bilangan bulat (dividend/pembilang dan divisor/penyebut)
// dan mengembalikan hasil dan error jika penyebut adalah nol
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}

//! Kode di bawah tidak boleh dimodifikasi.
func main() {
	nums := []int{5, 2, 9, 1, 7}
	min, max, avg := calculateStats(nums)
	fmt.Printf("Stats - Min: %d, Max: %d, Average: %.2f\n", min, max, avg)
	// Seharusnya mencetak semacam: Stats - Min: 1, Max: 9, Average: 4.80
	
	user := createUser("John", "Doe", 30)
	fmt.Println(user) // Seharusnya mencetak: John Doe (30)
	
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 ÷ 2 =", result) // Seharusnya mencetak: 10 ÷ 2 = 5
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Seharusnya mencetak pesan error
	} else {
		fmt.Println("10 ÷ 0 =", result)
	}
}