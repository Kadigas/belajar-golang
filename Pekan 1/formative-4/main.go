package main

import "fmt"

func main() {
	//soal 1
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d - ", i)
		if i%2 != 0 {
			if i%3 == 0 {
				fmt.Println("I Love Coding")
			} else {
				fmt.Println("Santai")
			}
		} else {
			fmt.Println("Berkualitas")
		}
	}

	//soal 2
	i := 1
	for i <= 7 {
		j := 1
		for j <= i {
			fmt.Printf("#")
			j++
		}
		fmt.Printf("\n")
		i++
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newKalimat = kalimat[2:7]
	fmt.Println(newKalimat)

	//soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	for idx, sayuran := range sayuran {
		fmt.Printf("%d. %s\n", idx+1, sayuran)
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var vBalok = 1
	for k, v := range satuan {
		vBalok *= v
		fmt.Println(k, "=", v)
	}
	fmt.Printf("Volume Balok = %d\n", vBalok)
}
