package dao

import (
	"fmt"
	"serverwechat/config"
	"serverwechat/datasource"
	"serverwechat/logger"
	"serverwechat/model"
)

//  注册新用户
func CreateAuto(id string) bool {
	user := &model.Auto{
		Id:       id,
		UserId:   id,
		AutoBot:  "nobot",
		AutoDesc: "正在忙",
	}
	// has, e := datasource.Engine.Where("Id = ?", id).Get(&model.Auto{})
	// if e != nil {
	// 	logger.Logger.Error(fmt.Sprintf("查询用户失败%v, error: %v", id, e.Error()))
	// 	return false
	// }
	// if has == true {
	// 	logger.Logger.Info("配置已存在")
	// 	return true
	// }
	_, err := datasource.Engine.Insert(*user)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("创建配置失败 %v", err))
		return false
	}
	logger.Logger.Info("创建配置成功")
	return true
}

func GetAutoInfoById(Id string) (model.Auto, int) {
	userInfo := make([]model.Auto, 0)
	sf := "id"
	si := Id
	err := datasource.Engine.Cols("auto_reply", "auto_reply_group", "auto_bot", "auto_desc", "tuling_api_key", "enable").Where(sf+" = ?", si).Find(&userInfo)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("获取用户信息失败, Openid: %v", Id))
		return model.Auto{}, -1
	}
	if len(userInfo) == 0 {
		logger.Logger.Error(fmt.Sprintf("获取用户信息成功，数据为空, Openid: %v", Id))
		return model.Auto{}, 0
	}
	logger.Logger.Info(fmt.Sprintf("获取用户信息成功, Openid: %v; Info: %v", Id, userInfo[0]))
	return userInfo[0], 1
}

func UpdateAuto(id string, info *model.Auto) (bool, int) {
	_, err := datasource.Engine.Cols("auto_reply", "auto_reply_group", "auto_bot", "auto_desc", "tuling_api_key").Where("id = ?", id).Update(info)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("更新配置信息失败, Openid: %v, %v", id, err))
		return false, config.STATUS_ERROR
	}
	logger.Logger.Error("更新配置信息成功")
	return true, config.STATUS_SUE
}

func UpdateAutoEnable(id string, enable bool) (bool, int) {
	_, err := datasource.Engine.Cols("enable").Where("id = ?", id).Update(&model.Auto{Enable: enable})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("更新配置状态信息失败, Openid: %v, %v", id, err))
		return false, config.STATUS_ERROR
	}
	logger.Logger.Error("更新配置状态信息成功")
	return true, config.STATUS_SUE
}
