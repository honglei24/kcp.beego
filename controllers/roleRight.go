package controllers

import (
	"github.com/astaxie/beego"
	"kcp/utils"
)

type RoleRightController struct {
	beego.Controller
}

// @Title get roleRight
// @Description get roleRight info by id
// @Param   roleRight_id     path    string  true        "roleRight id"
// @Success 200 {object} models.RoleRight
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [get]
func (c *RoleRightController) Get() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title create roleRight
// @Description get roleRight info by id
// @Param   ClusterId     body   int false       "roleRight id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "roleRight name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.RoleRight
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *RoleRightController) Post() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title delete roleRight
// @Description get roleRight info by id
// @Param   id     path    string  true        "roleRight id"
// @Success 200 {object} models.RoleRight
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [delete]
func (c *RoleRightController) Delete() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title update roleRight
// @Description get roleRight info by id
// @Param   Id     body   int false       "roleRight id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "roleRight name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.RoleRight
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [post]
func (c *RoleRightController) Update() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title list roleRight
// @Description get roleRight list
// @Success 200 {object} models.RoleRight
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router / [get]
func (c *RoleRightController) List() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}