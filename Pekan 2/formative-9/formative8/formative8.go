package formative8

import "fmt"

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return t.JariJari * t.JariJari * t.Tinggi * 22 / 7
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * t.JariJari * (t.JariJari + t.Tinggi) * 22 / 7
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * ((float64(b.Panjang) * float64(b.Lebar)) + (float64(b.Panjang) * float64(b.Tinggi)) + (float64(b.Lebar) * float64(b.Tinggi)))
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneData interface {
	Display() string
}

func (p Phone) Display() string {
	str := "name : " + p.Name + "\nbrand : " + p.Brand + "\nyear : " + fmt.Sprint(p.Year) + "\ncolors : "
	for idx, color := range p.Colors {
		if idx == 0 {
			str += color
		} else {
			str += ", " + color
		}
	}
	return str
}

func LuasPersegi(num int, b bool) interface{} {
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
