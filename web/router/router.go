package router

import (
	"github.com/capn-o-source/yecaptain/internal/loader"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	return router
}

func AddCoreRoutes(router *gin.Engine, loader loader.Loader) *gin.Engine {
	router = loader.GetRoutes(router)

	return router
}
