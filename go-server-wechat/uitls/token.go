package uitls

import (
	"serverwechat/logger"

	jwt "github.com/appleboy/gin-jwt/v2"
)

var Identity_Key = "info"

// 解析token
func PaserToken(token string, authMiddleware *jwt.GinJWTMiddleware) string {
	tk, ok := authMiddleware.ParseTokenString(token)
	if ok != nil {
		logger.Logger.Error("paserToken error")
		return ""
	}
	claims := jwt.ExtractClaimsFromToken(tk)
	return claims[Identity_Key].(string)
}
