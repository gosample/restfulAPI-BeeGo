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
	restfulRouter := beego.NewNamespace("/smartcity",
		beego.NSNamespace("/cansats",
			beego.NSInclude(
				&controllers.UsrController{},
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
		beego.NSNamespace("/datossensados",
			beego.NSInclude(
				&controllers.DatoSensadoController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(restfulRouter)
}
