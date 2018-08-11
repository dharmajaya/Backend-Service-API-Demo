package controllers

import (
	"src/github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error401() {
    c.Data["Content"] = "Unauthorized Token"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error403() {
    c.Data["Content"] = "Forbidden page"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error404() {
    c.Data["Content"] = "Erro 404 page"
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error500() {
    c.Data["Content"] = "Oops ! "
    c.TplName = "error_page.tpl"
}


func (c *ErrorController) Error503() {
    c.Data["Content"] = "Service Unavailable, Please Try again"
    c.TplName = "error_page.tpl"
}
