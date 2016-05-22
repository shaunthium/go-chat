package routers

import (
	"github.com/shaunthium/go-chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
