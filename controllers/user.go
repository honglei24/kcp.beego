package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"kcp/models"
	"kcp/utils"
	"log"
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
	u := models.User{UserId: id}
	err = o.Read(&u)
	if err != nil {
		c.Ctx.WriteString("User not found, id:" + string(id))
		c.Abort("404")
	}
	c.Data["json"] = &u
	c.ServeJSON()
}

// @Title create user
// @Description create user
// @Param   RoleId    body   int false       "role id"
// @Param   username   body   string  true       "user name"
// @Param   password body   string  true       "password"
// @Param   mobile body   string  false       "mobile"
// @Param   email body   string  false       "email"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *UserController) Post() {
	var u models.User
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &u); err == nil {
		logs.Debug(&u)

		valid := validation.Validation{}
		valid.Required(u.Username, "username")
		valid.Required(u.Password, "password")
		valid.MaxSize(u.Username, 50, "username")
		valid.MaxSize(u.Password, 50, "password")

		if valid.HasErrors() {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
			c.Abort("400")
		}
		_, err = models.UserAdd(&u)
		if err == nil {
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]string{"msg": err.Error()}
			//c.ServeJSON()
			c.Abort("409")
		}
	} else {
		c.Data["json"] = map[string]string{"msg": err.Error()}
		c.ServeJSON()
		c.Abort("401")
	}
}

// @Title delete user
// @Description get user info by id
// @Param   id     path    string  true        "user id"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [delete]
func (c *UserController) Delete() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title update user
// @Description get user info by id
// @Param   Id     body   int false       "user id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "user name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [post]
func (c *UserController) Update() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title list user
// @Description get user list
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router / [get]
func (c *UserController) List() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}