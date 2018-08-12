package controllers

import (
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  "golang.org/x/crypto/bcrypt"
  "../models"
)


type DeleteController struct {
  UserController
}

func (this *DeleteController) Get() {
  beego.Debug("In DeleteController:Get - Start")

  this.Data["ProfileActive"] = true
}


func (this *DeleteController) Post() {
  beego.Debug("In DeleteController:Post - Start")

  this.Data["ProfileActive"] = true

  currentpassword := this.GetString("currentpassword")
  valid := validation.Validation{}
  valid.Required(currentpassword, "currentpassword")

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  flash := beego.NewFlash()

  //******** Read account from the database
  o := orm.NewOrm()
  o.Using("write")
  user := models.User{Uid: this.session["uid"].(string)}

  err := o.Read(&user)

  if err != nil {
    beego.Error("In DeleteController:Post - Error reading user from DB: ", err)
    flash.Error("Internal server error, Please try again later")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
    return
  }

  //******** Compare submitted password with the saved hash
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentpassword))

  if err != nil {
    flash.Error("The current password seems to be incorrect, please try again.")
    flash.Store(&this.Controller)
    return
  }


 //******** Delete user record from the database
  _, err = o.Delete(&user)
  if err != nil {
    flash.Error("Internal server error")
    flash.Store(&this.Controller)
    this.DelSession("session")
    this.Redirect("/user/signin", 303)
    return
  }

  flash.Notice("Your account is deleted.")
  flash.Store(&this.Controller)
  this.DelSession("session")
  this.Redirect("/user/signin", 303)

}//end DeleteController:Post() func
