package routers

import (
	"gowebsocket/controllers/home"
	"gowebsocket/controllers/systems"
	"gowebsocket/controllers/user"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	// 用户组
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.GET("/online", user.Online)
		userRouter.POST("/sendMessage", user.SendMessage)
		userRouter.POST("/sendMessageAll", user.SendMessageAll)
	}

	// 系统
	systemRouter := router.Group("/system")
	{
		systemRouter.GET("/state", systems.Status)
	}

	// home
	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/index", home.Index)
	}

	// router.POST("/user/online", user.Online)
}
