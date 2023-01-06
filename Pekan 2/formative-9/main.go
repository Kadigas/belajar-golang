package main

import (
	"fmt"
	. "formative-9/formative8"
)

func main() {

	//soal 1
	var bangunDatar HitungBangunDatar
	bangunDatar = SegitigaSamaSisi{Alas: 8, Tinggi: 4}
	fmt.Printf("Segitiga Sama Sisi\n=> Luas\t\t: %d\n=> Keliling\t: %d\n", bangunDatar.Luas(), bangunDatar.Keliling())
	bangunDatar = PersegiPanjang{Panjang: 8, Lebar: 4}
	fmt.Printf("Persegi Panjang\n=> Luas\t\t: %d\n=> Keliling\t: %d\n", bangunDatar.Luas(), bangunDatar.Keliling())

	var bangunRuang HitungBangunRuang
	bangunRuang = Tabung{JariJari: 7, Tinggi: 10}
	fmt.Printf("Tabung\n=> Volume\t\t: %.2f\n=> Luas Permukaan\t: %.2f\n", bangunRuang.Volume(), bangunRuang.LuasPermukaan())
	bangunRuang = Balok{Panjang: 15, Lebar: 6, Tinggi: 8}
	fmt.Printf("Balok\n=> Volume\t\t: %.2f\n=> Luas Permukaan\t: %.2f\n", bangunRuang.Volume(), bangunRuang.LuasPermukaan())

	//soal 2
	var hp PhoneData = Phone{Name: "Samsung Galaxy Note 20", Brand: "Samsung Galaxy Note 20", Year: 2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(hp.Display())

	//soal 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

	//soal 4

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// tulis jawaban anda disini
	var str = prefix.(string)
	var total = 0
	for idx, num := range kumpulanAngkaPertama.([]int) {
		if idx == 0 {
			str += fmt.Sprint(num)
		} else {
			str += " + " + fmt.Sprint(num)
		}
		total += num
	}
	for _, num := range kumpulanAngkaKedua.([]int) {
		str += " + " + fmt.Sprint(num)
		total += num
	}
	fmt.Println(str, "=", total)
}
