package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["kcp/controllers:ClusterController"] = append(beego.GlobalControllerRouter["kcp/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:ClusterController"] = append(beego.GlobalControllerRouter["kcp/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:ClusterController"] = append(beego.GlobalControllerRouter["kcp/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:ClusterController"] = append(beego.GlobalControllerRouter["kcp/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:ClusterController"] = append(beego.GlobalControllerRouter["kcp/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RightController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RightController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RightController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RightController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RightController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleRightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleRightController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleRightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleRightController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleRightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleRightController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleRightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleRightController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:RoleRightController"] = append(beego.GlobalControllerRouter["kcp/controllers:RoleRightController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kcp/controllers:UserController"] = append(beego.GlobalControllerRouter["kcp/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
