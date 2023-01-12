package router

import (
	"net/http"
	"serverwechat/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	authMiddleware := middleware.AuthMiddleware
	//登录接口
	r.POST("/login", authMiddleware.LoginHandler)
	// 跨域中间件
	r.Use(middleware.Cors())
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	r.StaticFS("/upload", http.Dir("upload"))
	r.StaticFS("/qrcodes", http.Dir("qrcodes"))
	auth := r.Group("/auth")
	//退出登录
	auth.POST("/logout", authMiddleware.LogoutHandler)
	// 刷新token，延长token的有效期
	auth.POST("/refresh_token", authMiddleware.RefreshHandler)
	// JWT中间件
	auth.Use(authMiddleware.MiddlewareFunc())
	{

	}

}
