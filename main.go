package main

import (
	//"io"
	"os"
	"realdummy/controller"
	"realdummy/logger"
	"realdummy/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var dd service.RequestService = service.New()
var reCon controller.Controller = controller.New(dd)

var PortNo string = ":8080"

func main() {

	//clear logs
	os.Truncate("logfile.log", 0)
	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {
		logMessage := "GET" + " |\t" + c.ClientIP() + " |\t" + c.FullPath()
		logger.Logger(logMessage)
		reCon.Respond(c)

	})

	//start server
	logger.Logger("Server start")
	server.Run(PortNo)

}
