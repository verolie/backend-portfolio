package server

import (
	"project/portofolio/handler"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	e := gin.Default()
	
	regisServer(e)

	e.Run()
}


func regisServer(e *gin.Engine){
	e.GET("job/list", getJobList)
	e.POST("job/list/info", postJobListInfo)
}

func getJobList(c *gin.Context) {
	handler.ProcessJobList(c)
}

func postJobListInfo(c *gin.Context) {
	handler.ProcessJobListInfo(c)
}