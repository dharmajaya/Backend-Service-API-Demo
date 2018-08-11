package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  "../models"
)


type SecurityController struct {
  UserController
}


func (this *SecurityController) Get() {
  beego.Debug("In SecurityController:Get - Start")
  this.Data["SecurityActive"] = true
  this.LayoutSections["Footer"] = "securitycontroller/footer.tpl"
}


func (this *SecurityController) Post() {

  beego.Debug("In SecurityController:Post - Start")

  this.LayoutSections["Footer"] = "securitycontroller/footer.tpl"

  o := orm.NewOrm()
  o.Using("write")

  user := models.User{Uid: this.session["uid"].(string)}

  beego.Debug("In SecurityController:Post - Reading user from the database")

  err := o.Read(&user)

  if err != nil {
    flash := beego.NewFlash()
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
  }

  this.Data["User"] = user
  this.Data["SecurityActive"] = true

  flash := beego.NewFlash()

  passwordUpdateForm := models.FormPasswordUpdate{}

  if err := this.ParseForm(&passwordUpdateForm); err != nil {
    beego.Debug("In SecurityController:Post - Got err parsing the form", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }


  valid := validation.Validation{}

  if v, _ := valid.Valid(&passwordUpdateForm); !v {
    beego.Debug("In SecurityController:Post - Got form validation err")
    this.Data["Errors"] = valid.ErrorsMap
    return
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordUpdateForm.CurrentPassword))

  if err != nil {
    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return
  }

  bcryptPasswordHash, err := CryptPassword([]byte(passwordUpdateForm.NewPassword))

  if err != nil {
    beego.Debug("In SecurityController:Post - Got err hashing the password: ", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  user.Password = string(bcryptPasswordHash)

  _, err = o.Update(&user)

  if err != nil {

    beego.Error("In SecurityController:Post - Gor err updating user in database", err)
    flash.Error("Internal server error - Please try later or let us know that something whent wrong.")
    flash.Store(&this.Controller)
    return
  }

  flash.Notice("Password updated cheers!")
  flash.Store(&this.Controller)

  this.Redirect("/user/security", 303)

}
