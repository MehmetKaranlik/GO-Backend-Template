package App

import (
	"Backend/Core/App/Configuration"
	"Backend/Core/App/Database"
	"Backend/Core/App/Environment"
	"Backend/Core/App/Router"
	"Backend/Core/Middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type App struct {
	Environment   *Environment.Env
	databases     []Database.IDatabase
	routers       []Router.IRouteContainer
	Configuration Configuration.IAppConfiguration
}

func (self *App) init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func (self *App) StartApp() {
	self.init()
	self.Environment.LoadEnvironment(self.Configuration)
	self.Environment.MapEnvironmentValues()

	for _, database := range self.databases {
		err := database.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Initialize Routes
	for _, versionRouter := range self.routers {
		Router.InitializeRoutes(versionRouter.MakeRoutes())
	}
	self.listen(self.Configuration.GetAppPort())

}

func (self *App) listen(port string) {
	_ = http.ListenAndServe(port, new(Middleware.LoggerMiddleWare))
}

func (self *App) AddDatabases(databases ...Database.IDatabase) {
	self.databases = databases
}

func (self *App) AddRouters(routers ...Router.IRouteContainer) {
	self.routers = routers
}
