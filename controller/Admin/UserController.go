package Admin

import (
	"crypto/md5"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/deatil/doak-cron/pkg/db"
	"github.com/kataras/iris/v12"
)

type UserController struct {}
//登录页面
func (user *UserController) Login(ctx iris.Context)  {
	ctx.View("admin/user/login.html")
}
//登录操作
func (user *UserController) DoLogin(ctx iris.Context)  {
	captchaCode := ctx.FormValue("captcha")
	if captcha.VerifyString(ctx.GetCookie("captchaId"), captchaCode){
		ctx.RemoveCookie("captchaId")
		username := ctx.FormValue("username")
		password := ctx.FormValue("password")
		var user db.AdminModel
		md5_password :=[]byte(password)
		db.Db.First(&user, "name=? and password=?", username, fmt.Sprintf("%x",md5.Sum(md5_password)))
		fmt.Print(user)
		if user != (db.AdminModel{}){
			ctx.Redirect("/admin/login")
		}
		ctx.Redirect("/admin")
	}else{
		ctx.Redirect("/admin/login")
	}
}
//个人信息页面
func (user *UserController) Profile(ctx iris.Context)  {
	ctx.View("admin/user/profile.html")
}

