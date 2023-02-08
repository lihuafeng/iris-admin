package Controller

import (
	"github.com/kataras/iris/v12"
)

type HomeController struct {}

/**
首页
 */
func (h *HomeController) Home(ctx iris.Context) {
	ctx.View("m/home/index.html")
}
/**
租房信息页
 */
func (h *HomeController) HouseInfo(ctx iris.Context)  {
	ctx.View("m/house/house_info.html")
}

/**
搜索列表页
 */
func (h *HomeController) Search(ctx iris.Context)  {
	ctx.View("m/house/house_search_list.html")
}

/**
小区房源页
 */
func (h *HomeController) Building(ctx iris.Context)  {
	ctx.View("m/house/building.html")
}