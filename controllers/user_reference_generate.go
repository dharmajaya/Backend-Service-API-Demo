package controllers

import (
"src/github.com/astaxie/beego"
"src/github.com/astaxie/beego/orm"
_ "src/github.com/astaxie/beego/cache/redis"
"../models"
"github.com/twinj/uuid"
"github.com/utils"
)


type SendReferenceLinkController struct {
	BaseController
}

func (this * SendReferenceLinkController ) Get() {
	beego.Debug("In ResetPasswordController:Get - Start")
	this.Data["ProfileActive"] = true
}

func (this * SendReferenceLinkController ) Post() {

	beego.Debug("In  SendReferenceLinkController :Post - Start")

	flash := beego.NewFlash()
	submitted_Uid := this.Ctx.Input.Param(":uuid")

	o := orm.NewOrm()
	o.Using("write")

	user := models.User{Registration_uid: submitted_Uid}

	user = user.CopySignUpForm(&signupform)

	user.Uid = uuid.NewV5(app_name_space, uuid.Name(user.Email)).String()


	// Add user to database with new uuid and send verification email
	registration_uid := uuid.NewV4().String()

	user.Reference_Link = registration_uid

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
	mail.Subject = " Account Activation"
	mail.HTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+
		"\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"

	err = mail.Send()

	if err != nil {
		beego.Error("Unable to send verification email")
		flash.Error("Unable to send verification email")
		flash.Store(&this.Controller)
		return
	}

	beego.Debug(" After sendEmail finished successfully")
	flash.Notice("Your account has been created. Please check your email to verify and activate the account before login.")
	flash.Store(&this.Controller)
	this.Redirect("/user/generateRefer", 303)

	err := o.Read(&user, "Password_reset_uid")

	if err != nil {
		flash.Notice("Invalid key.")
		flash.Store(&this.Controller)
		this.Redirect("/user/generateRefer", 303)
	}
	flash.Store(&this.Controller)
	this.Redirect("/user/generateRefer", 303)
}

