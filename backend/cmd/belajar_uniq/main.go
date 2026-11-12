package main

import "fmt"

type User struct {
	Name     string
	Username string
}

func main() {

	users := []User{
		{Name: "anita", Username: "sdasd"},
		{Name: "dita", Username: "asd"},
		{Name: "anita", Username: "3"},
		{Name: "anita", Username: "1"},
		{Name: "dita", Username: "4"},
		{Name: "j", Username: "fgh"},
	}

	var usersUniq map[string]bool = map[string]bool{}

Parent:
	for _, v := range users {
		if usersUniq[v.Name] == true {
			continue Parent
		}
		usersUniq[v.Name] = true
		fmt.Println("| Nama:", v.Name, "| Username:", v.Username, "|")
	}

	// users := []string{
	// 	"anita",
	// 	"fitri",
	// 	"asdasd",
	// 	"none",
	// 	"fitri",
	// 	"asdasd",
	// 	"james",
	// 	"fitri",
	// 	"asdasd",
	// 	"james",
	// 	"anita",
	// 	"fitri",
	// 	"asdasd",
	// 	"none",
	// 	"fitri",
	// 	"asdasd",
	// 	"james",
	// 	"fitri",
	// 	"asdasd",
	// 	"james",
	// 	"fitri",
	// 	"asdasd",
	// 	"james",
	// 	"anita",
	// 	"fitri",
	// 	"asdasd",
	// 	"none",
	// 	"fitri",
	// 	"asdasd",
	// }

	// var usersUniq []string

	// for _, v := range users {
	// 	var duplikat bool = false

	// 	for _, itm := range usersUniq {
	// 		if itm == v {
	// 			duplikat = true
	// 		}
	// 	}

	// 	if duplikat {
	// 		continue
	// 	} else {
	// 		usersUniq = append(usersUniq, v)
	// 	}
	// }

	// fmt.Println(usersUniq)
}
