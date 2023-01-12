package middleware

import (
	"fmt"
	"time"

	"serverwechat/logger"
	"serverwechat/model"
	"serverwechat/uitls"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var AuthMiddleware *jwt.GinJWTMiddleware

func InitJWT() *jwt.GinJWTMiddleware {
	logger.Logger.Info(fmt.Sprintf("JWT start"))
	// 定义一个Gin的中间件
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "wtechtec zone",       //标识
		SigningAlgorithm: "HS256",               //加密算法
		Key:              []byte("qweronexkfc"), //密钥
		Timeout:          7 * 24 * time.Hour,    // 过期时间
		MaxRefresh:       7 * 24 * time.Hour,    //刷新最大延长时间
		IdentityKey:      uitls.Identity_Key,    //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					uitls.Identity_Key: v.Id,
				}
			}
			logger.Logger.Error(fmt.Sprintf("JWT PayloadFunc"))
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, err := claims[uitls.Identity_Key].(string)
			if !err {
				logger.Logger.Error(fmt.Sprintf("JWT IdentityHandler %v", claims))
			}
			return &model.User{
				Id: id,
			}
		},
		Authenticator: Authenticator, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			if _, ok := data.(*model.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) { //错误时响应
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: handleLoginResponse,
	})

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("JWT Error %v", err.Error()))
	}
	logger.Logger.Info(fmt.Sprintf("JWT succeed"))
	AuthMiddleware = authMiddleware
	return authMiddleware
}
