package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "kcp/routers"
	_ "kcp/sysinit"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
