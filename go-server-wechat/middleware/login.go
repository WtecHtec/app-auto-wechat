package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"serverwechat/config"
	"serverwechat/dao"
	"serverwechat/datasource"
	"serverwechat/logger"
	"serverwechat/model"
	"serverwechat/uitls"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var URL = "https://api.weixin.qq.com/sns/jscode2session?grant_type=authorization_code"

// 登录逻辑
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	loginId := loginVals.Code
	loginStatus := false
	fmt.Println("loginVals==", loginVals)
	// 微信小程序登录 方式
	if loginVals.LoginType == "wx_app" && loginVals.Code != "" {
		openid := getWXOpen(loginVals.Code, loginVals.NickName)
		fmt.Println("loginVals== openid", openid)
		if loginVals.Id == "" {
			userInfo, st := dao.GetUserInfoByPbOpenId("", openid, "", "")
			if st == 1 {
				loginId = userInfo.Id
				loginStatus = true
			}
		} else {
			// 绑定
			userInfo, st := dao.GetUserInfoByPbOpenId("", "", loginVals.Id, "")
			if st == 1 && userInfo.ApOpenId == openid {
				loginId = userInfo.Id
				loginStatus = true
			} else if st == 1 && userInfo.ApOpenId == "" {
				b, _ := dao.UpdateUser(loginVals.Id, &model.User{ApOpenId: openid, NickName: loginVals.NickName})
				if b == true {
					loginId = userInfo.Id
					loginStatus = true
				}
			}
		}
	} else {
		// 体验码 、登录指令 登录方式
		t, id := datasource.GetRedisByString(loginId)
		fmt.Println("登录id===", id)
		if t == true {
			userInfo, st := dao.GetUserInfoByPbOpenId("", "", id, "")
			// 登录完 移除redis
			datasource.DelRedisByString(loginId)
			if st != -1 {
				loginStatus = true
				if st == 1 {
					loginId = userInfo.Id
				}
			}
		}
	}

	if loginId != "" {
		if loginStatus {
			return &model.User{
				Id: loginId,
			}, nil
		} else {
			return nil, jwt.ErrFailedAuthentication
		}
	}
	logger.Logger.Error(fmt.Sprintf("登陆失败, OpenId %v", loginVals))
	return nil, jwt.ErrFailedAuthentication
}

func getWXOpen(code string, nickName string) string {
	wxConfig := config.BASE_CONFIG.WxConfig
	fmt.Println("wxConfig==", wxConfig)
	url := fmt.Sprintf("%v&appid=%v&secret=%v&js_code=%v", URL, wxConfig.Appid, wxConfig.AppSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("获取openId 失败：%v", nickName))
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	user := &model.User{}
	jer := json.Unmarshal(body, user)
	if jer != nil {
		logger.Logger.Error(fmt.Sprintf("获取openId 失败：%v", jer))
		return ""
	}
	if resp.StatusCode == 200 {
		logger.Logger.Info(fmt.Sprintf("获取openId 成功：%v", user))
		return user.OpenId
	}
	logger.Logger.Error(fmt.Sprintf("获取openId 失败：%v", nickName))
	return ""
}

func handleLoginResponse(c *gin.Context, code int, message string, time time.Time) {
	id := uitls.PaserToken(message, AuthMiddleware)
	info, ok := dao.GetUserInfoByPbOpenId("", "", id, "")
	if ok != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"minikey": message,
		"code":    http.StatusOK,
		"info":    &model.User{Id: info.Id},
	})
}
