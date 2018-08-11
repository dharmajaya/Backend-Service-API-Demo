package routers

import (
	"github.com/astaxie/beego"
	"../controllers"
)


func init() {

        beego.Router("/", &controllers.IndexController{})

	beego.Router("/user/signup", &controllers.SignupController{})

	beego.Router("/user/signin", &controllers.SigninController{})

	beego.Router("/user/signout", &controllers.SignoutController{})

	beego.Router("/user/verify/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.AccountVerifyController{})

        beego.Router("/user/profile", &controllers.ProfileController{})

	beego.Router("/user/phoneverify", &controllers.PhoneVerifyController{})

	beego.Router("/user/delete", &controllers.DeleteController{})

        beego.Router("/user/security", &controllers.SecurityController{})

	beego.Router("user/generateRefer", &controllers.SendReferenceLinkController{})

	beego.Router("user/joinedWithRefer", &controllers.JoinedReferenceLinkController{})

	beego.Router("user/news", &controllers.NewsController{})

	beego.Router("/user/resetpassword/:uuid([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})", &controllers.ResetPasswordController{})
}
