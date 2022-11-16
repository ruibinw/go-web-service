package main

import (
	"git.epam.com/ryan_wang/go-web-service/cmd/server/di"
)

// @title 			Go Web Service
// @description 	A simple REST Web service written in Go that supports CRUD operations.
// @version 		1.0
// @contact.name 	Ryan_Wang
// @contact.email 	Ryan_Wang@epam.com
// @BasePath        /
func main() {
	server := di.InitServer()
	recordCtrl := di.InitRecordController()
	server.InitRouter()
	server.SetupSwagger("/swagger/*")
	server.RegisterRecordController(recordCtrl)
	server.Start()
}
