package main

import (
	"fmt"
)

func main() {
	// var nama string

	// fmt.Print("Masukkan nama anda : ")
	// _, err := fmt.Scan(&nama)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Halo", nama)

	// 	jawaban := "Y"

	// Parent:
	// 	for jawaban == "Y" || jawaban == "y" {

	// 		var user, err = NewUser()
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		// Filter Data
	// 		b, err := FilterUser(user)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if b == false {
	// 			fmt.Println("Kamu tidak boleh daftar")
	// 			continue Parent
	// 		}

	// 		// Simpan Data
	// 		check, err := AddData(&user)
	// 		if err != nil {
	// 			fmt.Println("Data gagal disimpan")
	// 		}

	// 		if !check {
	// 			fmt.Println("Config yaml salah")
	// 		} else {
	// 			fmt.Println("Data berhasil disimpan")
	// 		}

	//		fmt.Println("Lanjut (Y/n) : ")
	//		_, err = fmt.Scan(&jawaban)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	var jawaban string

Parent:
	for {
		fmt.Println("Pilihan :")
		fmt.Print("1. Daftar\n2. Login\n3. GetProfile\n4. Quit\nJawaban : ")
		fmt.Scan(&jawaban)

		switch jawaban {
		case "1":
			var user, err = NewUser()
			if err != nil {
				panic(err)
			}

			// Filter Data
			b, err := FilterUser(user)
			if err != nil {
				panic(err)
			}
			if b == false {
				fmt.Println("Kamu tidak boleh daftar")
				continue Parent
			}

			// Simpan Data
			check, err := AddData(&user)
			if err != nil {
				fmt.Println("Data gagal disimpan")
				continue Parent
			}

			if !check {
				fmt.Println("Config yaml salah")
				continue Parent
			}
			fmt.Println("Data berhasil disimpan")
		case "2":
			err := Login()
			if !err {
				fmt.Println("gagal login")
				continue Parent
			}
			fmt.Println("berhasil login")
		case "3":
			var nama string
			fmt.Print("Masukkan nama : ")
			fmt.Scan(&nama)
			user, err := Profile(nama)
			if err != nil {
				panic(err)
			}
			fmt.Println("Nama:", user.Name)
			fmt.Println("Username:", user.Username)
			fmt.Println("Email:", user.Email)
		case "4":
			fmt.Println("Keluar")
			break Parent
		default:
			fmt.Println("Input salah")
			break Parent
		}
	}
}
