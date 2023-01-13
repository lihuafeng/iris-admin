package Controller

import (
	"github.com/dchest/captcha"
	Config "github.com/deatil/doak-cron/config"
	"github.com/deatil/doak-cron/controller/Admin"
	"github.com/deatil/doak-cron/controller/Api"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/versioning"
	"time"
)

func RouterHandler(app *iris.Application)  {
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	//路由
	app.Get("/", new(IndexController).Index)
	app.Get("/add", new(IndexController).Add)
	app.Post("/save", new(IndexController).Save)
	app.Post("/modify", new(IndexController).Modify)
	//验证码
	app.Get("/captcha", func(ctx iris.Context) {
		captchaId := captcha.NewLen(Config.CAPTCHA)
		ctx.SetCookieKV("captchaId", captchaId, iris.CookieExpires(time.Duration(60)*time.Second))
		_ = captcha.WriteImage(ctx, captchaId, 130, 40)
	})
	//admin
	app.PartyFunc("/admin", func(admin iris.Party) {
		admin.Get("/login", new(Admin.UserController).Login).Name = "admin.login"
		admin.Post("/login", new(Admin.UserController).DoLogin).Name = "admin.doLogin"
		admin.Get("/loginout", new(Admin.UserController).LoginOut).Name = "admin.loginOut"

		admin.Get("/profile", new(Admin.UserController).Profile) // 个人信息
		admin.Post("/send_email_code", new(Admin.UserController).SendEmailCode) // 发送邮箱验证码
		//后台首页
		admin.Get("/", new(Admin.IndexController).Index).Name = "admin"
	})
	//api
	app.PartyFunc("api", func(api router.Party) {
		api.Get("/", versioning.NewMatcher(versioning.Map{
			"1.0":              new(Api.IndexController).Test1,
			">= 2, < 3":         new(Api.IndexController).Test2,
		}))
	})
}

//404
func notFound(ctx iris.Context)  {
	ctx.View("errors/404.html")
}
//500
func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}
