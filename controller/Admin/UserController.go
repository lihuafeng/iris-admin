package Admin

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/deatil/doak-cron/pkg/db"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
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
		res := db.Db.First(&user, "name=? and password=?", username, fmt.Sprintf("%x",md5.Sum(md5_password)))
		if res.Error != nil{
			ctx.Redirect("/admin/login")
			return
		}
		//更新登录时间
		row := db.Db.Exec("update admin set last_login_time=? where id=?", time.Now().Unix(), user.Id)
		if row.Error != nil{
			ctx.Redirect("/admin/login")
			return
		}
		//写入session
		session := sess.Start(ctx)
		user_json,_ := json.Marshal(user)
		session.Set("userInfo", user_json)

		ctx.Redirect("/admin")
		return
	}else{
		ctx.Redirect("/admin/login")
		return
	}
}
//个人信息页面
func (user *UserController) Profile(ctx iris.Context)  {
	ctx.View("admin/user/profile.html")
}
