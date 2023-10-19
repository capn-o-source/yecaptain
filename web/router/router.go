package router

import (
	"github.com/capn-o-source/cronfusion/internal/loader"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	return router
}

func AddPluginRoutes(router *gin.Engine, loader loader.Loader) *gin.Engine {
	router = loader.GetRoutes(router)

	return router
}
