package api

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/routers"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()

	routers.NewAuthRouter(r, *controllers.NewAuthController()).Initialize()
	routers.NewMenuRouter(r, *controllers.NewMenuController()).Initialize()
	routers.NewUserRouter(r, *controllers.NewUserController()).Initialize()
	routers.NewMealRouter(r, *controllers.NewMealController()).Initialize()
	routers.NewOrderRouter(r, *controllers.NewOrderController()).Initialize()

	r.Run()

	return nil
}
