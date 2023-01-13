package router

import (
	"fmt"
	"net/http"
	"serverwechat/dao"
	"serverwechat/middleware"
	"serverwechat/model"
	"serverwechat/store"
	"serverwechat/uitls"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 跨域中间件
	r.Use(middleware.Cors())
	authMiddleware := middleware.AuthMiddleware
	//登录接口
	r.POST("/login", authMiddleware.LoginHandler)
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	r.StaticFS("/public", http.Dir("public"))
	r.StaticFS("/qrcodes", http.Dir("qrcodes"))
	auth := r.Group("/auth")
	//退出登录
	auth.POST("/logout", authMiddleware.LogoutHandler)
	// 刷新token，延长token的有效期
	auth.POST("/refresh_token", authMiddleware.RefreshHandler)
	// JWT中间件
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/check", func(c *gin.Context) {})

		auth.GET("/getautoconfig", func(ctx *gin.Context) {
			tokenString := ctx.GetHeader("authorization")
			if tokenString == "" {
				ctx.JSON(http.StatusBadRequest, "TOKEN")
				return
			}

			id := uitls.PaserToken(strings.Split(tokenString, "Bearer ")[1], authMiddleware)
			if id == "" {
				ctx.JSON(http.StatusBadRequest, tokenString+"==ID=="+strings.Split(tokenString, "Bearer ")[1])
				return
			}
			info, st := dao.GetAutoInfoById(id)
			if st == -1 {
				ctx.JSON(http.StatusBadRequest, "-1")
				return
			}
			fmt.Println(id, st)
			if st == 0 {
				dao.CreateAuto(id)
				info, _ = dao.GetAutoInfoById(id)
			}
			if store.BOTS[id] != nil {
				store.BOTS[id].Auto = &info
			}
			ctx.JSON(http.StatusOK, info)
		})

		auth.POST("/updateautoconfig", func(ctx *gin.Context) {
			var autoInfo model.Auto
			if err := ctx.ShouldBindJSON(&autoInfo); err != nil {
				// 返回错误信息
				// gin.H封装了生成json数据的工具
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			tokenString := ctx.GetHeader("authorization")
			if tokenString == "" {
				ctx.JSON(http.StatusBadRequest, "TOKEN")
				return
			}
			id := uitls.PaserToken(strings.Split(tokenString, "Bearer ")[1], authMiddleware)
			if id == "" {
				ctx.JSON(http.StatusBadRequest, tokenString+"==ID=="+strings.Split(tokenString, "Bearer ")[1])
				return
			}
			dao.UpdateAuto(id, &autoInfo)
			if store.BOTS[id] != nil {
				store.BOTS[id].Auto = &autoInfo
			}
			ctx.JSON(http.StatusOK, "")
		})
	}

}
