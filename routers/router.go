package routers

import (
	"restfulAPI-BeeGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	sensCan := beego.NewNamespace("/smartcity",
		beego.NSNamespace("/cansats",
			beego.NSInclude(
				&controllers.CansatController{},
			),
		),
		beego.NSNamespace("/sensores",
			beego.NSInclude(
				&controllers.SensorController{},
			),
		),
		beego.NSNamespace("/sensorestipo",
			beego.NSInclude(
				&controllers.SensorTipoController{},
			),
		),
		beego.NSNamespace("/sensoresdato",
			beego.NSInclude(
				&controllers.SensorDatoController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(sensCan)
}
