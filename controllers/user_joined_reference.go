package controllers

import (
	"src/github.com/astaxie/beego"
	"src/github.com/astaxie/beego/orm"
	"src/github.com/astaxie/beego/validation"
	_ "src/github.com/astaxie/beego/cache/redis"
	"../models"
)

type JoinedReferenceLinkController struct {
	BaseController
}

func (this * JoinedReferenceLinkController) Get(){
	beego.Debug("Joined with Reference ")
	this.Data["ProfileActive"] = true

}
func (this * JoinedReferenceLinkController ) Post(){

	beego.Debug("In Joined Reference Controller:Post - Start")
	flash := beego.NewFlash()
	Username := this.GetString("Username")
	Reference_id := this.GetString("Reference_id")

	valid := validation.Validation{}
	valid.Required(Reference_id, "Reference_id")

	this.Data["Username"] = Username
        this.Data["Reference_id"] = Reference_id
	if valid.HasErrors() {
		errormap := make(map[string]string)
		for _, err := range valid.Errors {
			errormap[err.Key] = err.Message
		}
		this.Data["Errors"] = errormap
		return
	}

	//******** Read user account from the database
	o := orm.NewOrm()
	o.Using("default")

	user := models.Profile{Username: Username, Reference_id:Reference_id}

	err := o.Read(&user, "Reference_id")

	if err == orm.ErrNoRows || err == orm.ErrMissPK {

		flash.Error("Doesnt Exists Reference ID")
		flash.Store(&this.Controller)
		return

	} else if err != nil {

		flash.Error("Internal server error -Please try again.")
		flash.Store(&this.Controller)
		return
	}

	//******** Create session and go back to previous page
	m := make(map[string]interface{})

	m["uid"] = user.Uid
	m["UserName"] = user.Username
	m["Reference_id"] = user.Reference_id
	m["Registers_count"] = user.Registers_count

	beego.Debug("In Joined Referened Controller:Post - Creating new session")
	this.SetSession("session", m)

	if Reference_id != "" {
		this.Redirect("/user/signin", 303)
	}

	this.Redirect("/user/joinedWithRefer", 303)

}
