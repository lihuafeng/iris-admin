package Admin

import "github.com/kataras/iris/v12"

type IndexController struct {}

func (index *IndexController) Index(ctx iris.Context)  {
	ctx.View("admin/index/index.html")
}