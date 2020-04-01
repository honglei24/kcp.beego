package routers

import (
	"kcp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/?:id", &controllers.UserController{})
}
