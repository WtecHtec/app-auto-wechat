package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"serverwechat/config"
	"serverwechat/dao"
	"serverwechat/datasource"
	"serverwechat/logger"
	"serverwechat/middleware"
	"serverwechat/model"
	"serverwechat/router"
	"serverwechat/uitls"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TOKEN = "wtechtec"

// 定义接收数据的结构体
type WxPublic struct {
	Signature string `form:"signature" json:"signature" uri:"signature" xml:"signature"`
	Timestamp string `form:"timestamp" json:"timestamp" uri:"timestamp" xml:"timestamp"`
	Nonce     string `form:"nonce" json:"nonce" uri:"nonce" xml:"nonce"`
	Echostr   string `form:"echostr" json:"echostr" uri:"echostr" xml:"echostr"`
}

type PublicMsg struct {
	MsgType      string `form:"MsgType" json:"MsgType" uri:"MsgType" xml:"MsgType"`
	Content      string `form:"Content" json:"Content" uri:"Content" xml:"Content"`
	ToUserName   string `form:"ToUserName" json:"ToUserName" uri:"ToUserName" xml:"ToUserName"`
	FromUserName string `form:"FromUserName" json:"FromUserName" uri:"FromUserName" xml:"FromUserName"`
}

func main() {
	// 初始化日志
	logger.InitLogger()
	// 初始化配置
	config.InitConfig()
	// 初始化MySQL
	datasource.InitMysqlXORM()
	// 初始化Redis
	datasource.InitRedis()
	// 同步数据结构
	model.InitModel()
	// 初始化JWT
	middleware.InitJWT()

	// 任务调度初始化
	uitls.InitTimeTask()
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {

		// wx := wxbot.NewWxBot()
		// go func() {
		// 	wx.Login()
		// }()
		// wx.Login()
		// c.Request.Body

		// 声明接收的变量
		wxPublic := &WxPublic{
			Signature: c.Query("signature"),
			Timestamp: c.Query("timestamp"),
			Nonce:     c.Query("nonce"),
			Echostr:   c.Query("echostr"),
		}
		if wxPublic.Signature == "" {
			c.String(http.StatusOK, "hello, NO Signature")
			return
		}
		c.String(http.StatusOK, wxPublic.Echostr)
	})

	r.POST("/", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		var publicMsg PublicMsg
		if err := xml.Unmarshal(data, &publicMsg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Println("post publicMsg==", publicMsg)
		// 限制一个用户 5 分钟内只能申请一个
		has, txt := datasource.GetRedisByString(publicMsg.FromUserName)
		if publicMsg.MsgType == "text" && publicMsg.Content == "体验码" && has == false && txt == "empty" {
			datasource.SetRedisByString(publicMsg.FromUserName, 1, 5*time.Minute)
			_, status := dao.GetUserInfoByPbOpenId(publicMsg.FromUserName, "", "")
			uuid := uuid.New()
			key := uuid.String()
			desc := "体验码"
			if status == -1 {
				c.String(http.StatusBadRequest, "")
				return
			} else if status == 1 {
				desc = "登录指令"
			} else {
				// 创建一个用户
				dao.CreateUser(key, publicMsg.FromUserName)
			}
			datasource.SetRedisByString(key, 1, 5*time.Minute)
			replyText := formatPbTxtMsg(&publicMsg, fmt.Sprintf("您好,%v: %v ,时效5分钟.感谢使用.", desc, key))
			c.String(http.StatusOK, replyText)
		}
	})

	router.InitRouter(r)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":" + config.BASE_CONFIG.Port)
}

func formatPbTxtMsg(publicMsg *PublicMsg, txt string) string {
	return fmt.Sprintf(`<xml>
				<ToUserName><![CDATA[%v]]></ToUserName>
				<FromUserName><![CDATA[%v]]></FromUserName>
				<CreateTime>%v</CreateTime>
				<MsgType><![CDATA[text]]></MsgType>
				<Content><![CDATA[%v]]></Content>
			</xml>`, publicMsg.FromUserName, publicMsg.ToUserName, int(time.Now().Unix()), txt)
}
