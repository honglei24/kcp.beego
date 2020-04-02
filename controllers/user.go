package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"kcp/models"
)

type UserController struct {
	beego.Controller
}

//func (c *UserController) URLMapping() {
//	c.Mapping("Get", c.Get)
//}

// @Title get user
// @Description get user info by id
// @Param   id     path    string  true        "user id"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [get]
func (c *UserController) Get() {
	id, err := c.GetInt(":id")
	fmt.Println(id)
	if err != nil {
		c.Ctx.WriteString("Invalid id supplied")
		c.Abort("400")
	}
	o := orm.NewOrm()
	u := models.User{Id: id}
	err = o.Read(&u)
	if err != nil {
		c.Ctx.WriteString("User not found, id:" + string(id))
		c.Abort("404")
	}
	c.Data["json"] = &u
	c.ServeJSON()
}

// @Title create user
// @Description get user info by id
// @Param   Id     body   int false       "user id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "user name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *UserController) Create() {

	var u models.User
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &u); err == nil {

		o := orm.NewOrm()
		_, err := o.Insert(&u)
		if err != nil {

			c.Data["json"] = "{{\"Add\":\" successed\"}}"
		} else {

		}
	} else {
		c.Data["json"] = err.Error()
	}
	//c.ServeJSON()
	c.Ctx.WriteString("OK")
}
