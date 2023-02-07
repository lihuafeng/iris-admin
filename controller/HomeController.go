package Controller

import (
	"github.com/kataras/iris/v12"
)

type HomeController struct {}

func (h *HomeController) Home(ctx iris.Context) {

	ctx.View("m/home/index.html")
}