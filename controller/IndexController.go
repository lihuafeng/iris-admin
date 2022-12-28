package controller

import (
	"fmt"
	"github.com/deatil/doak-cron/pkg/db"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"time"
)

type IndexController struct {}

func (index *IndexController) Index(ctx iris.Context)  {
	var crons []db.CronModel
	//db.Db.Where("status=1").Find(&crons)
	db.Db.Find(&crons)
	ctx.ViewData("crons", crons)
	// Render template file: ./views/index.html
	ctx.View("index.html")
}

func (index *IndexController) Add(ctx iris.Context)  {
	ctx.View("add.html")
}

func (index *IndexController) Save(ctx iris.Context)  {
	CronTime := ctx.FormValue("CronTime")
	command := ctx.FormValue("command")
	cron_type, _ := ctx.PostValueInt("type")
	fmt.Printf("CronTime:%s, cron_type:%d, command:%s", CronTime, cron_type, command)
	db.Db.Create(&db.CronModel{
		UniueCode:uuid.NewString(),
		CronType:cron_type,
		CronTime:CronTime,
		Command:command,
		RunStatus:0,
		Status:1,
		CreatedAt:time.Now().Format("2006-01-02 15:04:05"),
	})
	ctx.Redirect("/")
}

func (index *IndexController) Modify(ctx iris.Context)  {
	uniue_code := ctx.PostValue("uniue_code")
	status,err := ctx.PostValueInt("status")
	if err!=nil{
		ctx.JSON(iris.Map{"code":1,"msg":"缺少参数"})
	}
	var cron db.CronModel
	db.Db.First(&cron, "uniue_code=?",uniue_code)
	db.Db.Model(&cron).Update("status", status)
	ctx.JSON(iris.Map{"code":0,"msg":"修改成功"})
}