package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  "../models"
)


type ProfileController struct {
  UserController
}


func (this *ProfileController) Get() {
  beego.Debug("In ProfileController:Get - Start")

  o := orm.NewOrm()
  o.Using("default")

  user := models.User{Uid: this.session["uid"].(string)}

  beego.Debug("In ProfileController:Get - Reading user from the database")

  err := o.Read(&user)

  if err != nil {
    flash := beego.NewFlash()
    flash.Error(" Server Error - Try Again.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
  }

  this.Data["User"] = user
  this.Data["ProfileActive"] = true
  this.LayoutSections["Footer"] = "profilecontroller/footer.tpl"

}//end ProfileController:Get() func


func (this *ProfileController) Post() {
  beego.Debug("In ProfileController:Post - Start")

  this.Data["ProfileActive"] = true
  this.LayoutSections["Footer"] = "profilecontroller/footer.tpl"

  flash := beego.NewFlash()

  o := orm.NewOrm()
  o.Using("write")

  user := models.User{Uid: this.session["uid"].(string)}

  beego.Debug("In ProfileController:Post - Read ")

  err := o.Read(&user)

  if err != nil {
    flash.Error("Server Error  - Try again.")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
  }

  this.Data["User"] = user

  accountUpdateForm := models.FormUserUpdate{}

  if err := this.ParseForm(&accountUpdateForm); err != nil {
    beego.Debug("In ProfileController:Post - Got err parsing the form", err)
    flash.Error("Server Error - Try Again.")
    flash.Store(&this.Controller)
    return
  }

  user = user.CopyUpdateForm(&accountUpdateForm)
  this.Data["User"] = user

  valid := validation.Validation{}

  if v, _ := valid.Valid(&accountUpdateForm); !v {
    beego.Debug("In ProfileController:Post - Got form validation err")
    this.Data["Errors"] = valid.ErrorsMap
    return
  }

  // Compare the password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(accountUpdateForm.CurrentPassword))

  if err != nil {
    errormap := make(map[string]string)
    errormap["Current"] = "The current password seems to be incorrect, please try again."
    this.Data["Errors"] = errormap
    flash.Error("The current password seems to be incorrect, please try again.")
    flash.Store(&this.Controller)
    return
  }

  // Save user info deatails to database
  user = user.CopyUpdateForm(&accountUpdateForm)

  _, err = o.Update(&user)

  if err != nil {

    beego.Error("In ProfileController:Post - Gor err updating user in database", err)
    flash.Error("Server Error - Try Again..")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
  }

  flash.Notice("Profile updated")
  flash.Store(&this.Controller)

  //******** update session details
  m := make(map[string]interface{})

  m["uid"] = user.Uid
  m["UserName"] = accountUpdateForm.Username
  m["NickName"] = accountUpdateForm.Nickname
  m["Email"] = accountUpdateForm.Email

  beego.Debug("In SigninController:Post - Creating new session")
  this.SetSession("session", m)

  this.session["Email"] = accountUpdateForm.Email
  this.session["UserName"] = accountUpdateForm.Username

  this.Data["User"] = user
  this.Redirect("/user/profile", 303)

}
