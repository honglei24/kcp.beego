package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
