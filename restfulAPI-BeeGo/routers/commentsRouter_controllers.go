package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["smartCity/controllers:CansatController"] = append(beego.GlobalControllerRouter["smartCity/controllers:CansatController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:CansatController"] = append(beego.GlobalControllerRouter["smartCity/controllers:CansatController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Id_cansat`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smartCity/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smartCity/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smartCity/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smartCity/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smartCity/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:SensorController"] = append(beego.GlobalControllerRouter["smartCity/controllers:SensorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:SensorController"] = append(beego.GlobalControllerRouter["smartCity/controllers:SensorController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Id_sensor`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:SensorTipoController"] = append(beego.GlobalControllerRouter["smartCity/controllers:SensorTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:SensorTipoController"] = append(beego.GlobalControllerRouter["smartCity/controllers:SensorTipoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Tipo_sensor`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["smartCity/controllers:UserController"] = append(beego.GlobalControllerRouter["smartCity/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
