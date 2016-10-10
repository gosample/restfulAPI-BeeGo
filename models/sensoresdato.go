package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

type sensorDato struct{
  Id_cansat  string `json:"id_cansat"`
  Id_sensor  string `json:"id_sensor"`
  Value      float32 `json:"value"`
  Fecha      string `json:"fecha"`
  Hora       string `json:"hora"`
}

var listDato []sensorDato

func init(){
  listDato = sensorDatoRequest()
}

func GetAllSensorDato() []sensorDato{
  listDato = sensorDatoRequest()
  return listDato
}

func GetSensorDato(Id_sensor string) (listCansatxSensor [] float32, err error) {
  listDato = sensorDatoRequest()
  var sens []float32
  for i := 0; i < len(listSensores); i++ {
    if listDato[i].Id_sensor==Id_sensor {
      sens = append(sens, listDato[i].Value)
    }
	}
	if sens==nil{
    fmt.Printf("\nNo se encontraron cansats con el tipo Sensor !\n\n")
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
