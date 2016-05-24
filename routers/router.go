package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaunthium/go-chat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/create", &controllers.MainController{}, "get,post:Create")
	beego.Router("/join", &controllers.MainController{}, "get,post:Join")
	beego.Router("/room/:id", &controllers.MainController{}, "get,post:Room")
}
