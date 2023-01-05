package main

import "fmt"

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return t.jariJari * t.jariJari * t.tinggi * 22 / 7
}

func (t tabung) luasPermukaan() float64 {
	return 2 * t.jariJari * (t.jariJari + t.tinggi) * 22 / 7
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * ((float64(b.panjang) * float64(b.lebar)) + (float64(b.panjang) * float64(b.tinggi)) + (float64(b.lebar) * float64(b.tinggi)))
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneData interface {
	display() string
}

func (p phone) display() string {
	str := "name : " + p.name + "\nbrand : " + p.brand + "\nyear : " + fmt.Sprint(p.year) + "\ncolors : "
	for idx, color := range p.colors {
		if idx == 0 {
			str += color
		} else {
			str += ", " + color
		}
	}
	return str
}

func luasPersegi(num int, b bool) interface{} {
	if num != 0 {
		if b == true {
			return "luas persegi dengan sisi 2 cm adalah " + fmt.Sprint(num)
		} else {
			return num
		}
	} else {
		if b == true {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
}

func main() {

	//soal 1
	var bangunDatar hitungBangunDatar
	bangunDatar = segitigaSamaSisi{8, 4}
	fmt.Printf("Segitiga Sama Sisi\n=> Luas\t\t: %d\n=> Keliling\t: %d\n", bangunDatar.luas(), bangunDatar.keliling())
	bangunDatar = persegiPanjang{8, 4}
	fmt.Printf("Persegi Panjang\n=> Luas\t\t: %d\n=> Keliling\t: %d\n", bangunDatar.luas(), bangunDatar.keliling())

	var bangunRuang hitungBangunRuang
	bangunRuang = tabung{7, 10}
	fmt.Printf("Tabung\n=> Volume\t\t: %.2f\n=> Luas Permukaan\t: %.2f\n", bangunRuang.volume(), bangunRuang.luasPermukaan())
	bangunRuang = balok{15, 6, 8}
	fmt.Printf("Balok\n=> Volume\t\t: %.2f\n=> Luas Permukaan\t: %.2f\n", bangunRuang.volume(), bangunRuang.luasPermukaan())

	//soal 2
	var hp phoneData
	hp = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(hp.display())

	//soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

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
