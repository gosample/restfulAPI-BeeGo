package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

type Sensor struct {
  Id_sensor   string `json:"id_sensor"`
  Id_cansat   string `json:"id_cansat"`
  Tipo_sensor string `json:"tipo_sensor"`
  Unidad      string `json:"unidad"`
  Modelo      string `json:"modelo"`
  F_install   string `json:"f_install"`
  H_install   string `json:"h_install"`
  Prefijo     string `json:"prefijo"`
}

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

var listSensores []Sensor

func init() {
  listSensores = sensoresRequest()
}

func GetAllSensores() []Sensor {
  listSensores = sensoresRequest()
	return listSensores
}

func GetSensorId(Tipo_sensor string) (sensor []Sensor, err error) {
  listSensores = sensoresRequest()
  var sens []Sensor
  for i := 0; i < len(listSensores); i++ {
    if listSensores[i].Tipo_sensor==Tipo_sensor {
      sens = append(sens, listSensores[i])
    }
	}
	if sens==nil{
    fmt.Printf("\nNo se encontraron sensores de ese tipo !\n\n")
	  return nil, errors.New("Sensor no existe")
	}
	return sens, nil
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
