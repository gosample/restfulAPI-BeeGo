package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:CansatController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:CansatController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:CansatController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:CansatController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Id_cansat`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Tipo_sensor/:Id_cansat`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:DatoSensadoController"],
		beego.ControllerComments{
			Method: "GetTypes",
			Router: `/:Tipo_sensor`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Tipo_sensor`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorTipoController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorTipoController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:SensorTipoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Tipo_sensor`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["restfulAPI-BeeGo/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
