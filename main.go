package main

import (
	"Backend/Core/App"
	"Backend/Core/App/Configuration"
	"Backend/Core/Globals"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Routes/V1"
)

func main() {
	app := App.App{
		Environment:   &Globals.EnvValues,
		Configuration: Configuration.DebugConfiguration{},
	}
	app.AddDatabases(Mongo.MongoDatabase{})
	app.AddRouters(V1.V1Routes{})
	app.StartApp()
}
