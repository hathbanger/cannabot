package routers

import (
	"cannabot/cannabot_gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

