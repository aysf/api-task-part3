package main

import (
	"aysf/day6r1/config"
	"aysf/day6r1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
