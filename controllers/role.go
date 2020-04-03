package controllers

import (
	"github.com/astaxie/beego"
	"kcp/utils"
)

type RoleController struct {
	beego.Controller
}

// @Title get role
// @Description get role info by id
// @Param   role_id     path    string  true        "role id"
// @Success 200 {object} models.Role
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [get]
func (c *RoleController) Get() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title create role
// @Description get role info by id
// @Param   ClusterId     body   int false       "role id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "role name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.Role
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *RoleController) Post() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title delete role
// @Description get role info by id
// @Param   id     path    string  true        "role id"
// @Success 200 {object} models.Role
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [delete]
func (c *RoleController) Delete() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title update role
// @Description get role info by id
// @Param   Id     body   int false       "role id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "role name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.Role
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [post]
func (c *RoleController) Update() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title list role
// @Description get role list
// @Success 200 {object} models.Role
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router / [get]
func (c *RoleController) List() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}