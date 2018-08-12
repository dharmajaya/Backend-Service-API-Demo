package controllers

import (
	"src/github.com/astaxie/beego"
	"src/github.com/astaxie/beego/orm"
	_ "src/github.com/astaxie/beego/cache/redis"
	"sgithub.com/twinj/uuid"
	"../models"
)


//News information controller

type NewsController struct {
	BaseController
}

func (this * NewsController) Get(){
	beego.Debug("News Information ")
	this.Data["ProfileActive"] = true
	this.LayoutSections["Footer"] = "newscontroller/footer.tpl"
}

func (this * NewsController ) Post(){
	beego.Debug("News Controller : Post - start")
	flash := beego.NewFlash()
	News := this.Ctx.Input.Param(":news")
	o := orm.NewOrm()
	o.Using("write")
        market := models.Market{News:News}
	market = market.PromoteNews
	market = market.BoardNews
	market = market.NewPrice
	market = market.AssetName
	market = models.saveNewsForm(&FormNews)

	market.Uid = uuid.NewV5(app_name_space, uuid.Name(market.AssetName)).String()
	_, err = o.Insert(&user)
	if err != nil {
		beego.Error("news controller Post - Got err inserting user to the database: ", err)
		flash.Error(FormNews.news + " already registered")
		flash.Store(&this.Controller)
		return
	}
	flash.Store(&this.Controller)
	this.Redirect("/user/news", 303)
}