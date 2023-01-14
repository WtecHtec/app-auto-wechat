package dao

import (
	"fmt"
	"serverwechat/config"
	"serverwechat/datasource"
	"serverwechat/logger"
	"serverwechat/model"
)

//  注册新用户
func CreateUser(id string, pbOpenId string) bool {
	user := &model.User{
		Id:       id,
		Enable:   true,
		PbOpenId: pbOpenId,
	}
	has, e := datasource.Engine.Where("pb_openid = ?", pbOpenId).Get(&model.User{})
	if e != nil {
		logger.Logger.Error(fmt.Sprintf("查询用户失败%v, error: %v", pbOpenId, e.Error()))
		return false
	}
	if has == true {
		logger.Logger.Info("用户已存在")
		return true
	}
	_, err := datasource.Engine.Insert(*user)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("创建新用户失败 %v", err))
		return false
	}
	logger.Logger.Info("创建新用户成功")
	return true
}

// 根据 公众号openid 查询用户信息
func GetUserInfoByPbOpenId(pbopenId string, apopendId string, Id string, wxUnId string) (model.User, int) {
	userInfo := make([]model.User, 0)
	sf := "pb_openid"
	si := pbopenId
	if apopendId != "" {
		sf = "ap_openid"
		si = apopendId
	} else if Id != "" {
		sf = "id"
		si = Id
	} else if wxUnId != "" {
		sf = "wx_unid"
		si = wxUnId
	}
	err := datasource.Engine.Cols("id", "wx_unid", "ap_openid", "enable", "create_time").Where(sf+" = ?", si).Find(&userInfo)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("获取用户信息失败, Openid: %v", pbopenId))
		return model.User{}, -1
	}
	if len(userInfo) == 0 {
		logger.Logger.Error(fmt.Sprintf("获取用户信息成功，数据为空, Openid: %v", pbopenId))
		return model.User{}, 0
	}
	logger.Logger.Info(fmt.Sprintf("获取用户信息成功, Openid: %v; Info: %v", pbopenId, userInfo[0]))
	return userInfo[0], 1
}

// 修改数据
func UpdateUser(id string, info *model.User) (bool, int) {
	_, err := datasource.Engine.Cols("ap_openid", "nick_name").Where("id = ?", id).Update(&model.User{
		ApOpenId: info.ApOpenId,
		NickName: info.NickName,
	})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("更新用户信息失败, Openid: %v, %v", id, err))
		return false, config.STATUS_ERROR
	}
	logger.Logger.Error("更新用户信息成功")
	return true, config.STATUS_SUE
}
