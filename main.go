package main

import (
	"github.com/amatsuzero/ginchat/router"
	"github.com/amatsuzero/ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081")
}
