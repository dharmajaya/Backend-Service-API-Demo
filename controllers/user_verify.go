package controllers

import (
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  _ "src/github.com/astaxie/beego/cache/redis"
  "../models"
)


type AccountVerifyController struct {
  BaseController
}


func (this *AccountVerifyController) Get() {

  beego.Debug("In AccountVerifyController:Get - Start")

  flash := beego.NewFlash()

	uuid := this.Ctx.Input.Param(":uuid")

	o := orm.NewOrm()
	o.Using("write")

	user := models.User{Registration_uid: uuid}

	err := o.Read(&user, "Registration_uid")

  if err != nil {
    flash.Error("Unable to verify your account.Please try again.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
  }

  user.Registration_uid = ""

  if _, err := o.Update(&user); err != nil {

    beego.Error("In AccountVerifyController:Get - Got err updating user", err)
    flash.Error("Unable to verify your account.Please try again.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)

  }

  flash.Notice("Your account successfully activated, feel free to sign in.")
  flash.Store(&this.Controller)
  this.Redirect("/user/signin", 303)
}
