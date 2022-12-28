package Controller

import (
	"github.com/deatil/doak-cron/controller/Admin"
	"github.com/deatil/doak-cron/controller/Api"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/versioning"
)

func RouterHandler(app *iris.Application)  {
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	//路由
	app.Get("/", new(IndexController).Index)
	app.Get("/add", new(IndexController).Add)
	app.Post("/save", new(IndexController).Save)
	app.Post("/modify", new(IndexController).Modify)
	//admin
	app.PartyFunc("/admin", func(admin iris.Party) {
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
