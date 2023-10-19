package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/gin-gonic/gin"
)

var dirPath string = "./cores/"

type Loader interface {
	GetRoutes(router *gin.Engine) *gin.Engine
}

func SearchForCores() ([]string, error) {
	var cores []string

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(dirPath+file.Name()) != ".so" {
			continue
		}

		cores = append(cores, dirPath+file.Name())
	}

	return cores, nil
}

func LoadCores(core string) (Loader, error) {
	plugin, err := plugin.Open(core)
	if err != nil {
		return nil, fmt.Errorf("error loading plugin %s - %v", core, err)
	}

	symPlugin, err := plugin.Lookup("Loader")
	if err != nil {
		return nil, fmt.Errorf("error loading plugin %s - %v", core, err)
	}

	var loader Loader
	loader, ok := symPlugin.(Loader)
	if !ok {
		return nil, fmt.Errorf("error loading plugin %s - %v", core, err)
	}

	return loader, nil
}
