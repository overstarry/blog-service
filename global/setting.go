package global

import (
	"github.com/overstarry/blog-service/pkg/logger"
	"github.com/overstarry/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
