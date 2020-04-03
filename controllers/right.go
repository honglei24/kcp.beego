package controllers

import (
	"github.com/astaxie/beego"
	"kcp/utils"
)

type RightController struct {
	beego.Controller
}

// @Title get right
// @Description get right info by id
// @Param   right_id     path    string  true        "right id"
// @Success 200 {object} models.Right
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [get]
func (c *RightController) Get() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title create right
// @Description get right info by id
// @Param   ClusterId     body   int false       "right id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "right name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.Right
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *RightController) Post() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title delete right
// @Description get right info by id
// @Param   id     path    string  true        "right id"
// @Success 200 {object} models.Right
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [delete]
func (c *RightController) Delete() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title update right
// @Description get right info by id
// @Param   Id     body   int false       "right id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "right name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.Right
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [post]
func (c *RightController) Update() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title list right
// @Description get right list
// @Success 200 {object} models.Right
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router / [get]
func (c *RightController) List() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}