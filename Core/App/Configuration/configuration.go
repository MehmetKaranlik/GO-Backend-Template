package Configuration

import (
	"log"
	"os"
	"path/filepath"
)

const (
	DevEnv      = "dev"
	ProdEnv     = "prod"
	devEnvPath  = "Env/dev.env"
	prodEnvPath = "Env/prod.env"
	prodPort    = ":8080"
	devPort     = "localhost:8080"
)

type IAppConfiguration interface {
	GetAppPort() string
	GetAppEnv() string
	GetEnvDir() string
}

type DebugConfiguration struct {
}

func (d DebugConfiguration) GetEnvDir() string {
	return devEnvPath
}

func (d DebugConfiguration) GetAppPort() string {
	return devPort
}

func (d DebugConfiguration) GetAppEnv() string {
	return DevEnv
}

type ReleaseConfiguration struct {
}

func (r ReleaseConfiguration) GetEnvDir() string {
	dir, dirErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	return dir + prodEnvPath
}

func (r ReleaseConfiguration) GetAppPort() string {
	return prodPort
}

func (r ReleaseConfiguration) GetAppEnv() string {
	return ProdEnv
}
