package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"kcp/models"
)

type ClusterController struct {
	beego.Controller
}

func (c *ClusterController) Get() {
	id, err := c.GetInt(":id")
	fmt.Println(id)
	if err != nil {
		c.Ctx.WriteString("error")
		c.Abort("400")
	}
	o := orm.NewOrm()
	u := models.User{Id: id}
	err = o.Read(&u)
	c.Data["json"] = &u
	c.ServeJSON()
}