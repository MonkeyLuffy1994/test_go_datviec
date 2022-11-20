package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test_go_datviec/pkg/config"
	"test_go_datviec/pkg/routes"
	"test_go_datviec/pkg/service"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	sv := service.Server{Cof: c}

	r := gin.Default()

	routes.RouteUpload(r, &sv)

	r.Run(c.Port)
}
