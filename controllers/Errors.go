package controllers

import (
	"github.com/astaxie/beego"
)

type ErrrosController struct {
	beego.Controller
}

func (c *ErrrosController) Error403()  {
	c.TplName = "errors/403.html"
}

func (c *ErrrosController) Error404()  {
	c.TplName = "errors/404.html"
}
