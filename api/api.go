package api

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/routers"

	"github.com/gin-gonic/gin"
)

func Start() error {
  r := gin.Default()

  authRouter := routers.NewAuthRouter(r, *controllers.NewAuthController())
  authRouter.Initialize()

  r.Run()
  return nil
}	
