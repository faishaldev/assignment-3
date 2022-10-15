package main

import (
	"assignment-3/models"
	"assignment-3/routers"
)

func main() {
	models.Init()
	routers.StartServer().Run(":8080")
}
