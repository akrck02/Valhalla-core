package main

import (
	"runtime"
	"strings"

	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/services"
)

func main() {

	var _, current_execution_dir, _, _ = runtime.Caller(0)
	var BASE_PATH = current_execution_dir
	var _ = configuration.SetBasePath(BASE_PATH)

	// substract the last 1 directories
	BASE_PATH = BASE_PATH[:strings.LastIndex(BASE_PATH, "/")] + "/"

	configuration.SetBasePath(BASE_PATH)
	configuration.LoadConfiguration()
	services.Start()
}
