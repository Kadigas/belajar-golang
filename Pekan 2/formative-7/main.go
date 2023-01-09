package main

import (
	"fmt"
	"reflect"
)

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func printBuah(fruit buah) {
	if fruit.adaBijinya == true {
		fmt.Printf("Nama: %s, Warna: %s, Ada bijinya: Ada, Harga: %d\n", fruit.nama, fruit.warna, fruit.harga)
	} else {
		fmt.Printf("Nama: %s, Warna: %s, Ada bijinya: Tidak, Harga: %d\n", fruit.nama, fruit.warna, fruit.harga)
	}
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (obj segitiga) luas() int {
	return obj.alas * obj.tinggi / 2
}

func (obj persegi) luas() int {
	return obj.sisi * obj.sisi
}

func (obj persegiPanjang) luas() int {
	return obj.panjang * obj.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (hp *phone) addColors(str string) {
	hp.colors = append(hp.colors, str)
}

type movie struct {
	title, genre   string
	year, duration int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title: title, duration: duration, genre: genre, year: year})
}

func main() {

	//soal 1
	var nanas = buah{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000}
	var jeruk = buah{"Jeruk", "Oranye", true, 8000}
	var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = buah{"Pisang", "Kuning", false, 5000}
	printBuah(nanas)
	printBuah(jeruk)
	printBuah(semangka)
	printBuah(pisang)

	//soal 2
	var triangle = segitiga{3, 4}
	var square = persegi{4}
	var rectangle = persegiPanjang{4, 8}
	fmt.Println("Luas Segitiga dengan alas", triangle.alas, "dan tinggi", triangle.tinggi, "=", triangle.luas())
	fmt.Println("Luas Persegi dengan sisi", square.sisi, "=", square.luas())
	fmt.Println("Luas Persegi Panjang dengan panjang", rectangle.panjang, "dan lebar", rectangle.lebar, "=", rectangle.luas())

	//soal 3
	var pocoF4 = phone{name: "POCO F4", brand: "POCO", year: 2022}
	pocoF4.addColors("Moonlight Silver")
	pocoF4.addColors("Night Black")
	pocoF4.addColors("Nebula Green")
	fmt.Printf("Nama: %s; Brand: %s; Tahun: %d; Colors: ", pocoF4.name, pocoF4.brand, pocoF4.year)
	for idx, color := range pocoF4.colors {
		if idx == 0 {
			fmt.Printf("%s", color)
		} else {
			fmt.Printf(", %s", color)
		}
	}
	fmt.Println("")

	//soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for idx, film := range dataFilm {
		fmt.Printf("%d. ", idx+1)
		values := reflect.ValueOf(film)
		types := values.Type()
		for i := 0; i < values.NumField(); i++ {
			if types.Field(i).Name == "duration" {
				fmt.Println(types.Field(i).Name, ":", film.duration/60, "jam")
			} else {
				fmt.Println(types.Field(i).Name, ":", values.Field(i))
			}
		}
	}
}
