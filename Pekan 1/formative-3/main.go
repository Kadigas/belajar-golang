package main

import (
	"fmt"
	"strconv"
)

func main() {

	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPP, err1 := strconv.Atoi(panjangPersegiPanjang)
	lebarPP, err2 := strconv.Atoi(lebarPersegiPanjang)
	alasS, err3 := strconv.Atoi(alasSegitiga)
	tinggiS, err4 := strconv.Atoi(tinggiSegitiga)

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		var luasPersegiPanjang int = panjangPP * lebarPP
		var kelilingPersegiPanjang int = 2 * (panjangPP + lebarPP)
		var luasSegitiga int = alasS * tinggiS / 2

		fmt.Printf("%d %d %d\n", luasPersegiPanjang, kelilingPersegiPanjang, luasSegitiga)
	}

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var gradeJohn, gradeDoe string

	if nilaiJohn >= 80 {
		gradeJohn = "A"
	} else if nilaiJohn >= 70 {
		gradeJohn = "B"
	} else if nilaiJohn >= 60 {
		gradeJohn = "C"
	} else if nilaiJohn >= 50 {
		gradeJohn = "D"
	} else {
		gradeJohn = "E"
	}

	if nilaiDoe > 80 {
		gradeDoe = "A"
	} else if nilaiDoe >= 70 {
		gradeDoe = "B"
	} else if nilaiDoe >= 60 {
		gradeDoe = "C"
	} else if nilaiDoe >= 50 {
		gradeDoe = "D"
	} else {
		gradeDoe = "E"
	}

	fmt.Println(gradeJohn, gradeDoe)

	//soal 3
	var tanggal = 27
	var bulan = 3
	var tahun = 2002

	var bln string

	switch bulan {
	case 1:
		bln = "Januari"
	case 2:
		bln = "Februari"
	case 3:
		bln = "Maret"
	case 4:
		bln = "April"
	case 5:
		bln = "Mei"
	case 6:
		bln = "Juni"
	case 7:
		bln = "Juli"
	case 8:
		bln = "Agustus"
	case 9:
		bln = "September"
	case 10:
		bln = "Oktober"
	case 11:
		bln = "November"
	case 12:
		bln = "Desember"
	}

	fmt.Println(tanggal, bln, tahun)

	//soal 4
	if tahunKelahiran := 2002; tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahunKelahiran <= 1979 {
		fmt.Println("Generasi X")
	} else if tahunKelahiran <= 1994 {
		fmt.Println("Generasi Y (Milenials)")
	} else if tahunKelahiran <= 2015 {
		fmt.Println("Generasi Z")
	}
}
