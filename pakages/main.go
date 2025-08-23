package main

import (
	"fmt"

	"github.com/aayushrajputz/ecommerce/auth"
	"github.com/aayushrajputz/ecommerce/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("aayushrajput", "password12")

	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "aayush rajput",
	}

	fmt.Println(user.Email, user.Name)
	color.HiGreen("fuck you")

}
