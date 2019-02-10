package routers

import (
	"ormTest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/addUser", &controllers.MainController{}, "post:AddUser")
	beego.Router("/updateUser", &controllers.MainController{}, "post:UpdateUser")
	beego.Router("/addProduct", &controllers.MainController{}, "post:AddProduct")
}
