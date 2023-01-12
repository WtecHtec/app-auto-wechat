package uitls

import (
	"fmt"
	"serverwechat/logger"

	"github.com/robfig/cron/v3"
)

// 任务调度
func InitTimeTask() {
	c := cron.New()
	logger.Logger.Info("任务调度初始化")
	// 添加一个任务，每 天凌晨0:0:0（0 0 * * *） 执行一次
	// */1 * * * * ? 每1分钟
	// @every 2s 每2秒
	// @daily 每天凌晨0点
	_, ok := c.AddFunc("@daily", func() {
		logger.Logger.Info("任务调度启动")

	})
	if ok != nil {
		logger.Logger.Error(fmt.Sprintf("任务调度失败 %v", ok))
		return
	}
	// 开始执行（每个任务会在自己的 goroutine 中执行）
	c.Start()
}
