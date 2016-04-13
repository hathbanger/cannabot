package routers

import (
	"canna-bot/cannabot_gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
