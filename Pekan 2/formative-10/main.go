package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

func greetings(str string, year int) {
	fmt.Println(str, year)
}

func kelilingSegitigaSamaSisi(num int, b bool) (string, error) {
	if num != 0 {
		if b == true {
			return "keliling segitiga sama sisinya dengan sisi 2 cm adalah 6 cm", errors.New("")
		} else {
			return fmt.Sprint(num), errors.New("")
		}
	} else {
		if b == true {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			runPanic()
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
}

func runPanic() {
	defer recoverPanic()
	panic(" PANIC ERROR")
}

func recoverPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func tambahAngka(num int, total *int) {
	*total += num
}

func cetakAngka(total *int) {
	fmt.Println(*total)
	fmt.Println()
}

func tambahData(str string, phone *[]string) {
	*phone = append(*phone, str)
}

func display(phone []string) {
	for idx, ph := range phone {
		time.Sleep(time.Second)
		fmt.Printf("%d. %s\n", idx+1, ph)
	}
	fmt.Println()
}

func displayPhone(phone []string, wg *sync.WaitGroup) {
	for idx, ph := range phone {
		time.Sleep(time.Second)
		fmt.Printf("%d. %s\n", idx+1, ph)
	}
	fmt.Println()
	wg.Done()
}

func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}

func main() {
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	//soal 1
	defer greetings("Golang Backend Development", 2021)

	//soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	fmt.Println()

	//soal 3
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//soal 4
	var phones = []string{}
	tambahData("Xiaomi", &phones)
	tambahData("Asus", &phones)
	tambahData("IPhone", &phones)
	tambahData("Samsung", &phones)
	tambahData("Oppo", &phones)
	tambahData("Realme", &phones)
	tambahData("Vivo", &phones)
	sort.Strings(phones)
	display(phones)

	//soal 5
	phones = []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup

	sort.Strings(phones)

	wg.Add(1)
	go displayPhone(phones, &wg)
	wg.Wait()

	//soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
	fmt.Println()
}
