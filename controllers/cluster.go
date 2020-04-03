package controllers

import (
	"github.com/astaxie/beego"
	"kcp/utils"
)

type ClusterController struct {
	beego.Controller
}

// @Title get cluster
// @Description get cluster info by id
// @Param   cluster_id     path    string  true        "cluster id"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [get]
func (c *ClusterController) Get() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title create cluster
// @Description get cluster info by id
// @Param   ClusterId     body   int false       "cluster id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "cluster name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @router / [post]
func (c *ClusterController) Post() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title delete cluster
// @Description get cluster info by id
// @Param   id     path    string  true        "cluster id"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [delete]
func (c *ClusterController) Delete() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title update cluster
// @Description get cluster info by id
// @Param   Id     body   int false       "cluster id"
// @Param   RoleId    body   int true       "role id"
// @Param   Username   body   string  true       "cluster name"
// @Param   Password body   string  true       "password"
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router /:id [post]
func (c *ClusterController) Update() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}

// @Title list cluster
// @Description get cluster list
// @Success 200 {object} models.User
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
// @router / [get]
func (c *ClusterController) List() {
	c.Data["json"] = utils.RET_MSG_OK
	c.ServeJSON()
}