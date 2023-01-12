package middleware

//用于登录code
type login struct {
	Code      string `form:"code" json:"code" binding:"required"`
	NickName  string `form:"nickName" json:"nickName" binding:"required"`
	LoginType string `form:"loginType" json:"login_type" binding:"required"`
}

// type User struct {
// 	OpenId      string `json:"openid"`
// 	UserName    string
// 	PhoneNumber string
// 	HeadeImg    string
// }
