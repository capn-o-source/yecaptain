package main

import (
	"github.com/capn-o-source/yecaptain/internal/loader"
	"github.com/capn-o-source/yecaptain/web/router"
)

func main() {
	cores, err := loader.SearchForCores()
	if err != nil {
		panic(err)
	}

	apiServer := router.InitRouter()

	for _, core := range cores {
		coreLoader, err := loader.LoadCores(core)
		if err != nil {
			panic(err)
		}

		router.AddCoreRoutes(apiServer, coreLoader)
	}

	apiServer.Run(":8080")
}
