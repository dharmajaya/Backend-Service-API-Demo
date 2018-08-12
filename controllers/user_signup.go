package controllers

import (
  "src/github.com/astaxie/beego"
  "src/github.com/astaxie/beego/orm"
  "src/github.com/astaxie/beego/validation"
  _ "src/github.com/astaxie/beego/cache/redis"
  //"src/github.com/uuid"
  "github.com/twinj/uuid"
  "../models"
)

 
type SignupController struct {
  BaseController
}


func (this *SignupController) Get() {

  beego.Debug("In SignupController:Get - Start")
  this.LayoutSections["Footer"] = "signupcontroller/footer.tpl"
}


func (this *SignupController) Post() {

  beego.Debug("In SignupController:Post - Start")

  this.LayoutSections["Footer"] = "signupcontroller/footer.tpl"

  flash := beego.NewFlash()
  signupform := models.FormUserSignUp{}
  errormap := make(map[string]string)

  captchaVerification := cpt.VerifyReq(this.Ctx.Request)

  if err := this.ParseForm(&signupform); err != nil {
    beego.Error("SignupController:Post - Got err parsing form", err)
    flash.Error("  Couldn't process your request.")
    flash.Store(&this.Controller)
    return
  }

  this.Data["Form"] = signupform

  valid := validation.Validation{}

  if b, _ := valid.Valid(&signupform); !b {
    beego.Debug("SignupController:Post - Form Validation Errors: ", valid.ErrorsMap)
    this.Data["Errors"] = valid.ErrorsMap
    return
  }

  if !captchaVerification {
    beego.Debug("SignupController:Post - Got wrong captcha")
    errormap["Captcha"] = "Sorry but the characters you endered didn't match. Please try again"
    this.Data["Errors"] = errormap
    return
  }

  password := []byte(signupform.Password)

  bcryptPasswordHash, err := CryptPassword(password)

  if err != nil {
    beego.Error("SignupController:Post - Got err hashing password: ", err)
    flash.Error("We're very sorry but we couldn't process your request.")
    flash.Store(&this.Controller)
    return
  }

  o := orm.NewOrm()
  o.Using("write")

  user := models.User{}
 
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

  beego.Debug("SignupController:Post - After sendEmail finished successfully")
  flash.Notice("Your account has been created.")
  flash.Store(&this.Controller)
  this.Redirect("/user/signin", 303)

}
