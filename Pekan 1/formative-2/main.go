package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1

	var text1 = "Bootcamp"
	var text2 = "Digital"
	var text3 = "Skill"
	var text4 = "Sanbercode"
	var text5 = "Golang"
	text := text1 + " " + text2 + " " + text3 + " " + text4 + " " + text5
	fmt.Println(text)

	// soal 2
	halo := "Halo Dunia"
	newHalo := strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(newHalo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	kataKedua = strings.Replace(kataKedua, "s", strings.ToUpper("s"), 1)
	kataKetiga = strings.Replace(kataKetiga, "r", strings.ToUpper("r"), 1)
	kataKeempat = strings.ToUpper(kataKeempat)
	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	num1, err1 := strconv.Atoi(angkaPertama)
	num2, err2 := strconv.Atoi(angkaKedua)
	num3, err3 := strconv.Atoi(angkaKetiga)
	num4, err4 := strconv.Atoi(angkaKeempat)
	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		fmt.Printf("%d\n", num1+num2+num3+num4) // 124
	}

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	num := strconv.Itoa(angka)
	kalimat = `"` + kalimat + `"` + " - " + num
	fmt.Println(kalimat)

}
