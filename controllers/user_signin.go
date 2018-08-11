package controllers

import (
  "golang.org/x/crypto/bcrypt"
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  "../models"
)


type SigninController struct {
  BaseController
}


func (this *SigninController) Get() {
  beego.Debug("In SigninController:Get - Start")
}


func (this *SigninController) Post() {

  beego.Debug("In SigninController:Post - Start")

  flash := beego.NewFlash()
  email := this.GetString("email")
  submittedPassword := this.GetString("password")
  redirect_after_login := this.GetString("redirect_after_login")

  valid := validation.Validation{}

  valid.Email(email, "email")
  valid.Required(submittedPassword, "password")

  this.Data["Email"] = email

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  o := orm.NewOrm()
  o.Using("default")

  user := models.User{Email: email}

  err := o.Read(&user, "Email")
  
  if err == orm.ErrNoRows || err == orm.ErrMissPK {

    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return

  } else if err != nil {

    flash.Error("Internal server error -Please try again.")
    flash.Store(&this.Controller)
    return
  }
  // Check the phone is verified or not
  if user.PhoneVerify != "" {

    flash.Error("Please check your email to activate your mobile verficcation account before login.")
    flash.Store(&this.Controller)
    return
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(submittedPassword))

  if err != nil {
    flash.Error("The email or password you entered is incorrect.")
    flash.Store(&this.Controller)
    return
  }

  m := make(map[string]interface{})

  m["uid"] = user.Uid
  m["UserName"] = user.Username
  m["NickName"] = user.Nickname
  m["Email"] = email

  beego.Debug("In SigninController:Post - Creating new session")
  this.SetSession("session", m)

  if redirect_after_login != "" {
    this.Redirect(redirect_after_login, 303)
  }

  this.Redirect("/user/profile", 303)

}
