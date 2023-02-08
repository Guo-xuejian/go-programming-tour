package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS   // 服务端设置
	AppSetting      *setting.AppSettingS      // 应用设置
	DatabaseSetting *setting.DatabaseSettingS // 数据库设置
	Logger          *logger.Logger            // 日志
)
