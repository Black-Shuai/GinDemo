package Router

import (
	"GinDemo/Controllers"
	"GinDemo/Middlewares"
	"GinDemo/Sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("mysession", Sessions.Store))
	v1 := router.Group("/api/user")
	{
		v1.POST("/findUser", Controllers.UserFind)
	}

	router.Run(":8080")
}
