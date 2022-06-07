package main

import (
	"fmt"

	"main.go/database"
	"main.go/routers"
)

func main() {
	if err := database.InitDB(); err != nil {
		fmt.Println("Database error on init!")
		fmt.Println(err.Error())
		return
	}
	app := routers.InitGin()
	app.Run()
}
