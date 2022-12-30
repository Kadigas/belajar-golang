package main

import "fmt"

func luasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) (volume int) {
	volume = panjang * lebar * tinggi
	return
}

func introduce(name string, gender string, job string, age string) (message string) {
	if gender == "laki-laki" {
		message = "Pak "
	} else {
		message = "Bu "
	}
	message += name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	return
}

func buahFavorit(name string, buah ...string) string {
	message := "halo nama saya " + name + " dan buah favorit saya adalah"
	for idx, value := range buah {
		if idx == 0 {
			message += ` "`
		} else {
			message += `, "`
		}
		message += value + `"`
	}
	return message
}

func main() {

	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(title string, duration string, genre string, year string) {
		dataFilm = append(dataFilm, map[string]string{"genre": genre, "jam": duration, "tahun": year, "title": title})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
