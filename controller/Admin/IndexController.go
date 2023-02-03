package Admin

import (
	"encoding/json"
	"fmt"
	"github.com/deatil/doak-cron/pkg/db"
	"github.com/kataras/iris/v12"
)

type IndexController struct {}

func (index *IndexController) Index(ctx iris.Context)  {
	check(ctx)
	ctx.ViewLayout("admin/layouts/layout.html")
	ctx.View("admin/index/index.html")
}
//验证登录状态
func check(ctx iris.Context)  {
	session := sess.Start(ctx)
	user_json := session.Get("userInfo")
	if user_json != nil{
		var user db.AdminModel
		fmt.Printf("%#v\n", user_json)
		err := json.Unmarshal([]byte(user_json.(string)), &user)
		if err == nil{
			ctx.ViewData("userInfo", user)
		}
	}else{
		ctx.Redirect("/admin/login")
		return
	}
}