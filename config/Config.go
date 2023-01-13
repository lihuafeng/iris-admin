package Config

import "time"

const (
	SERVER_NAME = "lris-admin"
	SERVER_PORT = 8080

	//验证码长度
	CAPTCHA = 4

	EMAIL_CODE_LENGTH = 4

	SESSION_EXPIRE_TIME = time.Hour

	SERVER_TIME_FORMAT = "2006-01-02 15:04:05"
	SERVER_CHARSET     = "UTF-8"

	LOG_DIR    = "Runtime/Logs"
	VIEW_DIR   = "Public/views"
	VIEW_EXT   = ".html"
	STATIC_DIR = "Public"
	ICON_PATH  = "favicon.ico"

	DB_DRIVE = "mysql"
	DB_DNS   = "root:root@tcp(127.0.0.1:3306)/goadmin?charset=utf8&parseTime=true&loc=Local"

	REIDS_ADDR = "127.0.0.1:6379"
	REDIS_PWD  = ""
	REDIS_DB   = 0
)
