package util

import "github.com/astaxie/beego"

var (
	MYSQL_URL string
)

func init() {
	runmode := beego.AppConfig.String("runmode")
	config, err := beego.AppConfig.GetSection(runmode)
	if err != nil {
		panic("配置文件初始化失败")
	}
	MYSQL_URL = config["mysql_url"]
}
