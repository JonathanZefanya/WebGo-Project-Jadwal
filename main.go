package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/LutfiEkaprima/Goproject/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Login)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/profile", authcontroller.Profile)
	http.HandleFunc("/uts", authcontroller.UTSpage)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)

}