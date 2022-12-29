package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

//定时任务表结构
type CronModel struct {
	Id int `gorm:"not null" json:"id"`
	UniueCode string  `gorm:"not null" json:"uniue_code"`
	CronType int `gorm:"not null" json:"cron_type"`
	CronTime string `gorm:"not null" json:"cron_time"`
	Command string `gorm:"not null" json:"command"`
	RunStatus int `gorm:"not null" json:"run_status"`
	Status int `gorm:"not null" json:"status"`
	CreatedAt string `gorm:"null" json:"created_at"`
	LastRuntime int `gorm:"not null" json:"last_runtime"`
	NextRuntime int `gorm:"not null" json:"next_runtime"`
}

//自定义表名
func (cm *CronModel) TableName() string {
	return "cron"
}

//用户表结构
type AdminModel struct {

}

func init()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/cron?charset=utf8mb4&parseTime=True&loc=Local"
	datetimePrecision := 2

	Db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize: 256, // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision: true, // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision: &datetimePrecision, // default datetime precision
		DontSupportRenameIndex: true, // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	/**
	db.DB() 是获得db连接对象
	SetMaxIdleConns 是设置空闲时的最大连接数
	SetMaxOpenConns 设置与数据库的最大打开连接数
	SetConnMaxLifetime 每一个连接的生命周期等信息
	 */
	//con,err := Db.DB()
	//if err == nil {
	//	con.SetMaxIdleConns(1000)
	//	con.SetMaxOpenConns(100000)
	//	con.SetConnMaxLifetime(-1)
	//}

}

func New() *gorm.DB {
	return Db
}

