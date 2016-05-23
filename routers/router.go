package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaunthium/go-chat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/create", &controllers.MainController{}, "get,post:Create")
	beego.Router("/join", &controllers.MainController{}, "get:Join")
	beego.Router("/room/:id([0-9]+)", &controllers.MainController{}, "get:Room")
}
