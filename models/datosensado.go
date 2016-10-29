package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])

type datoSensado struct{
  Id_cansat   string  `json:"id_cansat"`
  Id_sensor   string  `json:"id_sensor"`
  Value       float32 `json:"value"`
  Tipo_sensor string  `json:"tipo_sensor"`
  Fecha       string  `json:"fecha"`
  Hora        string  `json:"hora"`
}

var listDato []datoSensado

func init(){
  listDato = datoSensadoRequest("","")
}

func GetAllDatoSensado() []datoSensado{
  listDato = datoSensadoRequest("","")
  return listDato
}

func GetDatoSensado(Tipo_sensor string,Id_usr string) (listUsrxSensor [] datoSensado, err error) {
  listDato = datoSensadoRequest(Tipo_sensor,Id_usr)
  if listDato ==nil {
    fmt.Printf("\nNo se encontro Sensores !\n\n")
    return nil, errors.New("Sensor no existe")
  } else {
    return listDato, nil
  }
/*  var sens []datoSensado
  for i := 0; i < len(listDato); i++ {
    if listDato[i].Tipo_sensor==Tipo_sensor && listDato[i].Id_usr==Id_usr && Id_usr != ""{
      sens = append(sens, listDato[i])
    }
	}
	if sens==nil{
    fmt.Printf("\nNo se encontraron IoTDevices con el tipo Sensor !\n\n")
	  return nil, errors.New("Sensor no existe")
	}
	return sens, nil*/
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
