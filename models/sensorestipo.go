package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

type sensorTipo struct{
  Prefijo     string `json:"prefijo"`
  Tipo_sensor string `json:"tipo_sensor"`
}

var listTipo []sensorTipo

func init(){
  listTipo = sensorTipoRequest()
}

func GetAllSensorTipo() []sensorTipo{
  listTipo = sensorTipoRequest()
  return listTipo
}

func GetSensorTipo(Tipo_sensor string) (listCansatxSensor []string, err error) {
  listTipo = sensorTipoRequest()
  var sens []string
  for i := 0; i < len(listSensores); i++ {
    if listSensores[i].Tipo_sensor==Tipo_sensor {
      sens = append(sens, listSensores[i].Id_cansat)
    }
	}
	if sens==nil{
    fmt.Printf("\nNo se encontraron cansats con el tipo Sensor !\n\n")
	  return nil, errors.New("Sensor no existe")
	}
	return sens, nil
}

func GetAllSensorTipo1() []string {
  listTipo = sensorTipoRequest()
  var elementsTipo []string
  var elementsPrefijo []string
  var sens []string
  for i := 0; i < len(listSensores); i++ {
    elementsTipo = append(elementsTipo,listSensores[i].Tipo_sensor)
    elementsPrefijo = append(elementsPrefijo,listSensores[i].Prefijo)
  }
  sensTipo := removeDuplicatesUnordered(elementsTipo)
  sensPrefijo := removeDuplicatesUnordered(elementsPrefijo)

  fmt.Printf("%v\n",sensTipo)
  fmt.Printf("%v\n",sensPrefijo)
  for i := 0; i < len(sensTipo); i++ {
    str:="{"+"tipo_sensor:"+sensTipo[i]+",prefijo:"+sensPrefijo[i]+"}"
    sens = append(sens,str)
  }
  fmt.Printf("%v\n",sens)
  return sens
}

/*
func AddSensor(Sensor Sensor) (SensorId string) {
	Sensor.SensorId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Sensors[Sensor.SensorId] = &Sensor
	return Sensor.SensorId
}

func UpdateSensor(SensorId string, Score string) (err error) {
	if v, ok := Sensors[SensorId]; ok {
		v.Modelo = Score
		return nil
	}
	return errors.New("SensorId Not Exampleist")
}

func DeleteSensor(SensorId string) {
	delete(Sensors, SensorId)
}
*/
