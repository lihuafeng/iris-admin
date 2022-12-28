package Controller

import (
	"github.com/deatil/doak-cron/controller/Admin"
	"github.com/kataras/iris/v12"
)

func RouterHandler(app *iris.Application)  {
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	//路由
	app.Get("/", new(IndexController).Index)
	app.Get("/add", new(IndexController).Add)
	app.Post("/save", new(IndexController).Save)
	app.Post("/modify", new(IndexController).Modify)

	app.PartyFunc("/admin", func(admin iris.Party) {
		admin.Get("/", new(Admin.IndexController).Index).Name = "admin"
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
