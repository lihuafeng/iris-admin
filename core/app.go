package core

import (
	"github.com/kataras/iris/v12"
	"github.com/deatil/doak-cron/controller"
)

func Run()  {
	//Lris
	httpapp := iris.Default()
	httpapp.Logger().SetLevel("error")
	httpapp.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	httpapp.Use(myMiddleware)
	// 加载视图模板地址
	httpapp.RegisterView(iris.HTML("./views", ".html"))
	//提供静态文件服务
	httpapp.HandleDir("/static", "./static")

	//前置操作
	httpapp.Use(before)
	//后置操作
	//httpapp.Use(after)

	httpapp.Get("/", controller.Index)
	httpapp.Get("/add", controller.Add)
	httpapp.Post("/save", controller.Save)
	httpapp.Post("/modify", controller.Modify)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	httpapp.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context)  {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context)  {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
