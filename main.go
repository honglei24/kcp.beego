package main

import (
	"github.com/astaxie/beego"
	_ "kcp/routers"
	_ "kcp/sysinit"
)

func main() {
	beego.Run()
}
