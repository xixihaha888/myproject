package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/register", &controllers.UserController{},"get:ShowRegister")
	beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")


}
