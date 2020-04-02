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
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
				//beego.NSGet("/:id", func(ctx *context.Context) {
				//	ctx.Output.Body([]byte("shopinfo"))
				//}),
			),
			beego.NSNamespace("/cluster",
				beego.NSInclude(
					&controllers.ClusterController{},
				),
			),
		)
	beego.AddNamespace(ns)

    beego.Router("/", &controllers.MainController{})
	//beego.Router("/user/?:id", &controllers.UserController{})
}
