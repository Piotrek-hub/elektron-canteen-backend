package api

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() error {
	log.Println(os.Getenv("JWT_SECRET_KEY"))
	r := gin.Default()

	authRouter := routers.NewAuthRouter(r, *controllers.NewAuthController())
	authRouter.Initialize()

	menuRouter := routers.NewMenuRouter(r, *controllers.NewMenuController())
	menuRouter.Initialize()

	r.Run()

	return nil
}
