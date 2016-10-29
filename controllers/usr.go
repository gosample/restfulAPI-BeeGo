package controllers

import (
	"restfulAPI-BeeGo/models"
//	"encoding/json"
	"github.com/astaxie/beego"
)

type UsrController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all usrs
// @Success 200 {usr} models.usr
// @Failure 403 :Id_usr is empty
// @router / [get]
func (o *UsrController) GetAll() {
  o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
  o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	obs := models.GetAllUsrs()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Get
// @Description find usr by id_usr
// @Param	Id_usr		path 	string	true		"the id_usr you want to get"
// @Success 200 {usr} models.usr
// @Failure 403 :Id_usr is empty
// @router /:Id_usr [get]
func (o *UsrController) Get() {
  o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
  o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	Id_usr := o.Ctx.Input.Param(":Id_usr")
	if Id_usr != "" {
		ob, err := models.GetUsr(Id_usr)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}
/*
// @Title Create
// @Description create sensor
// @Param	body		body 	models.Sensor	true		"The sensor content"
// @Success 200 {string} models.Sensor.Id
// @Failure 403 body is empty
// @router / [post]
func (o *SensorController) Post() {
	var ob models.Sensor
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	sensorid := models.AddSensor(ob)
	o.Data["json"] = map[string]string{"SensorId": sensorid}
	o.ServeJSON()
}

// @Title Update
// @Description update the sensor
// @Param	sensorId		path 	string	true		"The sensorid you want to update"
// @Param	body		body 	models.sensor	true		"The body"
// @Success 200 {sensor} models.sensor
// @Failure 403 :sensorId is empty
// @router /:sensorId [put]
func (o *SensorController) Put() {
	sensorId := o.Ctx.Input.Param(":sensorId")
	var ob models.Sensor
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.UpdateSensor(sensorId, ob.Modelo)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the sensor
// @Param	sensorId		path 	string	true		"The sensorId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 sensorId is empty
// @router /:sensorId [delete]
func (o *SensorController) Delete() {
	sensorId := o.Ctx.Input.Param(":sensorId")
	models.DeleteSensor(sensorId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
*/
