package main

import (
	"fmt"
	"net/http"
)

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return t.jariJari * t.jariJari * t.tinggi * 22 / 7
}

func (t tabung) luasAlas() float64 {
	return t.jariJari * t.jariJari * 22 / 7
}

func (t tabung) kelilingAlas() float64 {
	return 2 * t.jariJari * 22 / 7
}

type hitungBangunRuang interface {
	volume() float64
	luasAlas() float64
	kelilingAlas() float64
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini adalah web-server pertama Andhika Ditya Bagaskara D menggunakan Golang.")
}

func main() {
	//soal 4
	var jariJari float64 = 7
	var tinggi float64 = 10

	var bangunRuang hitungBangunRuang = tabung{jariJari, tinggi}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "jari-jari : %.0f, tinggi : %.0f, volume : %.0f, luas alas : %.0f, keliling alas : %.0f\n", jariJari, tinggi, bangunRuang.volume(),
			bangunRuang.luasAlas(), bangunRuang.kelilingAlas())
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)

}
