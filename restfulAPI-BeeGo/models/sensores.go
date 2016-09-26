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

type Sensor struct {
    Id_sensor   string `json:"id_sensor"`
    Id_cansat   string `json:"id_cansat"`
    Tipo_sensor string `json:"tipo_sensor"`
    Unidad      string `json:"unidad"`
    Tipo_valor  string `json:"tipo_valor"`
    Modelo      string `json:"modelo"`
    F_install   string `json:"f_install"`
    H_install   string `json:"h_install"`
    Prefijo     string `json:"prefijo"`
}

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

var listSensores []Sensor

func init() {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("smartcity").C("sensores")
  err = c.Find(bson.M{ }).All(&listSensores)
  js, __ := json.Marshal(listSensores)
  //fmt.Printf("%s\n",listCansats[0].Modelo)
  //fmt.Println(reflect.TypeOf(listCansats))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de sensores:\n%s\n\n",js)
}

func GetAllSensores() []Sensor {
	return listSensores
}

func GetSensorId(Id_sensor string) (sensor Sensor, err error) {
  sensor = listSensores[0]
  for i := 0; i < len(listSensores); i++ {
    if listSensores[i].Id_sensor==Id_sensor {
      return listSensores[i], nil
    }
	}
  fmt.Printf("\nNo se encontro Sensor !\n\n")
	return sensor, errors.New("Sensor no existe")
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