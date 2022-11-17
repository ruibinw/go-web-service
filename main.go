package main

// @title           Go Web Service
// @description     A simple REST Web service written in Go that supports CRUD operations.
// @version         1.0
// @contact.name    Ryan_Wang
// @contact.email   Ryan_Wang@epam.com
// @BasePath        /
func main() {
	recordCtrl := InitRecordController()
	server := InitServer()
	server.InitRouter()
	server.SetupSwagger("/swagger/*")
	server.RegisterRecordController(recordCtrl)
	server.Start()
}
