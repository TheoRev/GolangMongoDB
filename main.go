package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/TheoRev/GolangMongoDB/data"
	"github.com/TheoRev/GolangMongoDB/models/user"
	"github.com/kr/pretty"
)

func main() {
	data.InitData()
	user.SetStorage(data.UserStorage{})
	// u := &user.User{
	// 	UserName: "theoRev",
	// 	Email:    "hrevillafernandez79@gmail.com",
	// 	Password: "1234",
	// }

	nombres := []string{
		"Theo",
		"Almendra",
		"Jose",
		"Christian",
		"Marco",
	}

	for _, nombre := range nombres {
		u := &user.User{
			UserName: nombre,
			Email:    strings.ToLower(nombre) + "@gmail.com",
			Password: "123456",
		}

		// fmt.Println("CREATE")
		err := u.Create()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Usuario %s creado exitosamente\n", u.UserName)
		// pretty.Printf("%# v", u)
	}

	// FIND_ALL
	fmt.Println("FIND_BY ID")
	u := user.New()
	users, err := u.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	pretty.Printf("%# v", users)

	// // CREATE
	// fmt.Println("CREATE")
	// err := u.Create()
	// if err != nil {
	// 	log.Println(err)
	// }
	// pretty.Printf("%# v", u)

	// // UPDATE
	// fmt.Println("UPDATE")
	// u.UserName = "almendra"
	// u.Email = "alm@gmail.com"
	// err = u.Update()
	// if err != nil {
	// 	log.Println(err)
	// }
	// pretty.Printf("%# v", u)

	// // FIND_BY ID
	// fmt.Println("FIND_BY ID")
	// u2 := user.New()
	// // u2.ID = u.ID
	// u2.ID = bson.ObjectIdHex("5a737f28cae7c6349c0f793a")
	// err := u2.GetByID()
	// if err != nil {
	// 	fmt.Println("ERROR: ", err)
	// 	fmt.Println(err)
	// }
	// pretty.Printf("%# v", u2)

	// // DELETE
	// fmt.Println("DELETE")
	// u2 = user.New()
	// u2.ID = bson.ObjectIdHex("5a737f28cae7c6349c0f793a")
	// err = u2.Delete()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println("Eliminado correctamente")
	// }
	// pretty.Printf("%# v", u2)
}
