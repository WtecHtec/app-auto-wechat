package model

import (
	"fmt"
	"serverwechat/datasource"
	"serverwechat/logger"
)

func InitModel() {
	logger.Logger.Info("DataTable init start")
	err := datasource.Engine.Sync2(new(User), new(Auto))
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("DataTable error %v", err))
		return
	}
	logger.Logger.Info("DataTable init success")
}
