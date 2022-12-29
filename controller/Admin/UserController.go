package Admin

import (
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

type UserController struct {}
//登录页面
func (user *UserController) Login(ctx iris.Context)  {
	ctx.View("admin/user/login.html")
}
//登录操作
func (user *UserController) DoLogin(ctx iris.Context)  {
	//username := ctx.FormValue("username")
	//password := ctx.FormValue("password")
	captchaCode := ctx.FormValue("captcha")
	if captcha.VerifyString(ctx.GetCookie("captchaId"), captchaCode){
		ctx.RemoveCookie("captchaId")
		ctx.Redirect("/admin")
	}else{
		ctx.Redirect("/admin/login")
	}
}
//个人信息页面
func (user *UserController) Profile(ctx iris.Context)  {
	ctx.View("admin/user/profile.html")
}

