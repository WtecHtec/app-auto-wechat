package router

import (
	"net/http"
	"serverwechat/dao"
	"serverwechat/datasource"
	"serverwechat/middleware"
	"serverwechat/model"
	"serverwechat/store"
	"serverwechat/uitls"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitRouter(r *gin.Engine) {
	// 跨域中间件
	r.Use(middleware.Cors())
	authMiddleware := middleware.AuthMiddleware
	//登录接口
	r.POST("/login", authMiddleware.LoginHandler)

	// 0 : 初始化； 1 : 扫码； [unid]:确认登录
	r.GET("/qr", func(ctx *gin.Context) {
		singcode := ctx.Query("singcode")
		if singcode != "" {
			b, vsingcode := datasource.GetRedisByString(singcode)
			if b == false {
				ctx.JSON(http.StatusOK, gin.H{"code": 404, "singcode": ""})
			} else {
				if vsingcode == "1" {
					ctx.JSON(http.StatusOK, gin.H{"code": 201, "singcode": ""})
				} else if vsingcode != "0" {
					// 确定登录 立即删除
					datasource.DelRedisByString(singcode)
					ctx.JSON(http.StatusOK, gin.H{"code": 202, "singcode": vsingcode})
				}
			}
		} else {
			uuid := uuid.New()
			key := uuid.String()
			datasource.SetRedisByString(key, "0", 5*time.Minute)
			ctx.JSON(http.StatusOK, gin.H{"code": 200, "singcode": key})
		}
	})
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
		auth.POST("/check", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 200})
		})

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
			_, stid := dao.GetUserInfoByPbOpenId("", "", id, "")
			if stid == -1 || stid == 0 {
				ctx.JSON(http.StatusUnauthorized, "-1")
				return
			}
			info, st := dao.GetAutoInfoById(id)
			if st == -1 {
				ctx.JSON(http.StatusBadRequest, "-1")
				return
			}
			// fmt.Println(id, st)
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
			_, st := dao.GetUserInfoByPbOpenId("", "", id, "")
			if st == -1 || st == 0 {
				ctx.JSON(http.StatusUnauthorized, "-1")
				return
			}
			dao.UpdateAuto(id, &autoInfo)
			if store.BOTS[id] != nil {
				store.BOTS[id].Auto = &autoInfo
			}
			ctx.JSON(http.StatusOK, "")
		})

		auth.GET("/checkscancode", func(ctx *gin.Context) {
			singcode := ctx.Query("singcode")
			b, lcode := datasource.GetRedisByString(singcode)
			if singcode == "" || b == false || (b == true && lcode != "0") {
				ctx.JSON(http.StatusBadRequest, "singcode"+lcode)
				return
			}
			datasource.SetRedisByString(singcode, "1", 1*time.Minute)
			ctx.JSON(http.StatusOK, gin.H{"code": 200})
		})
		auth.GET("/scancode", func(ctx *gin.Context) {
			singcode := ctx.Query("singcode")
			b, lcode := datasource.GetRedisByString(singcode)
			if singcode == "" || b == false || (b == true && lcode != "1") {
				ctx.JSON(http.StatusBadRequest, "singcode"+lcode)
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
			info, st := dao.GetUserInfoByPbOpenId("", "", id, "")
			if st == -1 || st == 0 {
				ctx.JSON(http.StatusBadRequest, "-1")
				return
			}
			uuid := uuid.New()
			key := uuid.String()
			datasource.SetRedisByString(key, info.Id, 5*time.Minute)
			datasource.SetRedisByString(singcode, key, 5*time.Minute)
			ctx.JSON(http.StatusOK, key)
		})
	}

}
