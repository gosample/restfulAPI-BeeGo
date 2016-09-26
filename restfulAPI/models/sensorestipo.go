package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
  "gopkg.in/mgo.v2"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
)

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

type optionSensor struct{
  Prefijo     string `json:"prefijo"`
  Tipo_sensor string `json:"tipo_sensor"`
}

var listOption []optionSensor

func GetAllSensorTipo() []optionSensor{
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("smartcity").C("tiposensores")
  err = c.Find(bson.M{}).All(&listOption)
  js, __ := json.Marshal(listOption)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de sensores:\n%s\n\n",js)
  return listOption
}

func GetSensorTipo(Tipo_sensor string) (listCansatxSensor []string, err error) {
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
