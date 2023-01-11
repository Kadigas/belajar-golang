package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// soal 1
func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NilaiMhs = NilaiMahasiswa{ID: uint(len(nilaiNilaiMahasiswa) + 1)}
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&NilaiMhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			nama := r.PostFormValue("nama")
			matakuliah := r.PostFormValue("mata_kuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			NilaiMhs = NilaiMahasiswa{
				Nama:       nama,
				MataKuliah: matakuliah,
				Nilai:      uint(nilai),
			}
		}

		if NilaiMhs.Nilai > 100 {
			w.Write([]byte("Nilai tidak boleh lebih dari 100!"))
			return
		}

		if NilaiMhs.Nilai >= 80 {
			NilaiMhs.IndeksNilai = "A"
		} else if NilaiMhs.Nilai >= 70 {
			NilaiMhs.IndeksNilai = "B"
		} else if NilaiMhs.Nilai >= 60 {
			NilaiMhs.IndeksNilai = "C"
		} else if NilaiMhs.Nilai >= 50 {
			NilaiMhs.IndeksNilai = "D"
		} else if NilaiMhs.Nilai < 50 {
			NilaiMhs.IndeksNilai = "E"
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NilaiMhs)
		dataNilaiMahasiswa, _ := json.Marshal(NilaiMhs) // to byte
		w.Write(dataNilaiMahasiswa)                     // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username dan Password tidak boleh kosong!"))
			return
		}

		if username == "Admin" && password == "admin123" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau Password salah!"))
		return
	})
}

// soal 2
func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilaiMhs, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(nilaiMhs)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Coba API")
	})

	//soal 1
	http.Handle("/post-nilai", Auth(http.HandlerFunc(PostNilaiMahasiswa)))

	//soal 2
	http.HandleFunc("/nilai", GetNilaiMahasiswa)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
