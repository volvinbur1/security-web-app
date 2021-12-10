package main

import (
	"fmt"
	"github.com/volvinbur1/security-web-app/internal/auth"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
)

func main() {
	dbManager := db.New()
	defer dbManager.Disconnect()

	usr := cmn.User{
		Login:    "test1",
		Name:     "test1",
		Surname:  "test1",
		Password: "test123445678",
	}
	err := auth.Register(dbManager, usr)
	if err != nil {
		panic(err)
	}

	users, err := dbManager.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	err = auth.LoginUser(dbManager, usr)
	if err != nil {
		panic(err)
	}

	fmt.Println("user validated")
}
