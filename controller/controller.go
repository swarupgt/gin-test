package controller

import (
	"encoding/json"
	"realdummy/logger"
	"realdummy/requestdata"
	"realdummy/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controller interface {
	Respond(ctx *gin.Context)
}

type controller struct {
	service service.RequestService
}

//constructor
func New(s service.RequestService) Controller {
	return &controller{
		service: s,
	}
}

func (c *controller) Respond(ctx *gin.Context) {
	var newReq requestdata.RequestData
	var newLoc requestdata.LocationData

	err1 := ctx.ShouldBindBodyWith(&newLoc, binding.JSON)
	if err1 != nil {
		err2 := ctx.ShouldBindBodyWith(&newReq, binding.JSON)
		if err2 != nil {
			ctx.JSON(400, gin.H{"message": "Bad request, invalid API request body"})
			logger.Logger("400 Bad request, invalid API request body")
			return
		}
		ob1, _ := json.Marshal(newReq)
		logger.Logger("Recieved from GET: " + string(ob1))
		ctx.JSON(200, c.service.RespondText(newReq))
		logger.Logger("200 Text Response")
		return
	}
	ob2, _ := json.Marshal(newLoc)
	logger.Logger("Recieved from GET: " + string(ob2))
	ctx.JSON(200, c.service.RespondLoc(newLoc))
	logger.Logger("200 Location Response")

}
