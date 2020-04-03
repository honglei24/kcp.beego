// @APIVersion 1.0.0
// @Title kcp API
// @Description kubernetes control platform APIs.
// @Contact honglei
package routers

import (
	"kcp/controllers"
	"github.com/astaxie/beego"
//"github.com/astaxie/beego/context"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			//beego.NSBefore(auth),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/cluster",
				beego.NSInclude(
					&controllers.ClusterController{},
				),
			),
			beego.NSNamespace("/role",
				beego.NSInclude(
					&controllers.RoleController{},
				),
			),
			beego.NSNamespace("/right",
				beego.NSInclude(
					&controllers.RightController{},
				),
			),
			beego.NSNamespace("/roleRight",
				beego.NSInclude(
					&controllers.RoleRightController{},
				),
			),
		)
	beego.AddNamespace(ns)

    beego.Router("/", &controllers.MainController{})
	//beego.Router("/user/?:id", &controllers.UserController{})
}
