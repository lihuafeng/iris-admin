package core

import (
	Config "github.com/deatil/doak-cron/config"
	Controller "github.com/deatil/doak-cron/controller"
	"github.com/kataras/iris/v12"
	"strconv"
)

func Run()  {
	//Lris
	httpapp := iris.Default()
	httpapp.Logger().SetLevel("error")
	httpapp.Use(myMiddleware)
	// 加载视图模板地址
	httpapp.RegisterView(iris.HTML("./views", ".html"))
	//提供静态文件服务
	httpapp.HandleDir("/static", "./static")
	httpapp.HandleDir("/uploads", "./uploads")

	//加载路由
	Controller.RouterHandler(httpapp)

	//前置操作
	httpapp.Use(before)
	//后置操作
	//httpapp.Use(after)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	httpapp.Run(iris.Addr(":"+strconv.Itoa(Config.SERVER_PORT)))
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

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
