package controllers

import (
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  "../models"
  "github.com/twinj/uuid"
  "github.com/utils"
)


type ResetPasswordController struct {
  BaseController
}

func (this *ResetPasswordController) Get() {
  beego.Debug("In ResetPasswordController:Get - Start")
}


func (this *ResetPasswordController) Post() {

  beego.Debug("In ResetPasswordController:Post - Start")

  flash := beego.NewFlash()

  submitted_reset_uid := this.Ctx.Input.Param(":uuid")

  o := orm.NewOrm()
  o.Using("write")

  user := models.User{Password_reset_uid: submitted_reset_uid}

  user = user.CopySignUpForm(&signupform)

  user.Uid = uuid.NewV5(app_name_space, uuid.Name(user.Email)).String()

  user.Password = string(bcryptPasswordHash)

  // Add user to database with new uuid and send verification email
  registration_uid := uuid.NewV4().String()

  user.Registration_uid = registration_uid

  _, err = o.Insert(&user)
  if err != nil {
    beego.Error("SignupController:Post - Got err inserting user to the database: ", err)
    flash.Error(signupform.Email + " already registered")
    flash.Store(&this.Controller)
    return
  }

  // send activation email
  link := "http://" + base_url + "/accounts/verify/" + registration_uid

  mail := utils.NewEMail(email_config)
  mail.To = []string{signupform.Email}
  mail.From = gmail_account
  mail.Subject = "Beego-Ureg - Account Activation"
  mail.HTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+
          "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"

  err = mail.Send()

  if err != nil {
    beego.Error("SignupController:Post - Unable to send verification email")
    flash.Error("Unable to send verification email")
    flash.Store(&this.Controller)
    return
  }

  beego.Debug("SignupController:Post - After sendEmail finished successfully")
  flash.Notice("Your account has been created. Please check your email to verify and activate the account before login.")
  flash.Store(&this.Controller)
  this.Redirect("/user/signin", 303)

  err := o.Read(&user, "Password_reset_uid")

  if err != nil {
    flash.Notice("Invalid key.")
    flash.Store(&this.Controller)
    this.Redirect("/user/signin", 303)
  }

  new_password := this.GetString("password")
  new_password_confirmation := this.GetString("password2")

  valid := validation.Validation{}
  valid.MinSize(new_password, 10, "new_password")
  valid.Required(new_password_confirmation, "new_password_confirmation")

  if valid.HasErrors() {
    errormap := make(map[string]string)
    for _, err := range valid.Errors {
      errormap[err.Key] = err.Message
    }
    this.Data["Errors"] = errormap
    return
  }

  if new_password != new_password_confirmation {
    flash.Error("Passwords don't match")
    flash.Store(&this.Controller)
    return
  }

  // Hashing the password with the default cost of 10
  bcryptPasswordHash, err := CryptPassword([]byte(new_password))

  if err != nil {
    beego.Error("In ResetPasswordController:Post - err hashing new_password: ", err)
    flash.Error("We're very sorry but we couldn't process your request.")
    flash.Store(&this.Controller)
    return
  }

  user.Password = string(bcryptPasswordHash)

  user.Password_reset_uid = ""

  if _, err := o.Update(&user); err != nil {
    flash.Error("Internal server error")
    flash.Store(&this.Controller)
    return
  }

  flash.Notice("Password updated.")
  flash.Store(&this.Controller)
  this.Redirect("/user/signin", 303)
}
