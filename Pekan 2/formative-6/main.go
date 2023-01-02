package main

import "fmt"

func luasKelilingLingkaran(luas *float64, keliling *float64, r float64) {
	*luas = r * r * 22 / 7
	*keliling = 2 * r * 22 / 7
}

func introduce(str *string, name string, gender string, job string, age string) {
	if gender == "laki-laki" {
		*str = "Pak "
	} else {
		*str = "Bu "
	}
	*str += name + " adalah seorang " + job + " yang berusia " + age + " tahun"
}

func tambahBuah(str *[]string) {
	*str = []string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}
}

func tambahDataFilm(title string, duration string, genre string, year string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{"genre": genre, "duration": duration, "year": year, "title": title})
}

func main() {

	//soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	var r float64 = 7.0
	luasKelilingLingkaran(&luasLingkaran, &kelilingLingkaran, r)
	fmt.Println("Luas lingkaran:", luasLingkaran)
	fmt.Println("Keliling lingkaran:", kelilingLingkaran)

	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{}
	tambahBuah(&buah)
	for idx, fruit := range buah {
		fmt.Printf("%d. %s\n", idx+1, fruit)
	}

	//soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for idx, film := range dataFilm {
		fmt.Printf("%d. ", idx+1)
		for key, value := range film {
			fmt.Println(key, " : ", value)
		}
		fmt.Println("")
	}
}
