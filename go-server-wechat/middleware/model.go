package middleware

//用于登录code
type login struct {
	Code      string `form:"code" json:"code" binding:"required"`
	NickName  string `form:"nickName" json:"nickName"`
	LoginType string `form:"loginType" json:"login_type"`
	Id        string `form:"id" json:"id"`
}

// type User struct {
// 	OpenId      string `json:"openid"`
// 	UserName    string
// 	PhoneNumber string
// 	HeadeImg    string
// }
