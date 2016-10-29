package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

type Usr struct{
  Id_usr   string `json:"id_cansat"`
  Modelo      string `json:"modelo"`
  F_install   string `json:"f_install"`
  H_install   string `json:"h_install"`
  Longuitud   float64 `json:"longuitud"`
  Latitud     float64`json:"latitud"`
  Lugar       string `json:"lugar"`
  Sensores  []SensCan `json:"sensores"`
}

type SensCan struct{
  Tipo_Sensor  string `json:"tipo_sensor"`
  Id_Sensor    string `json:"id_sensor"`
}
// db.cansats.aggregate([{$unwind:"$id_sensores"},{$project:{_id:0,id_sensores:1}}])
// db.cansats.aggregate([{$unwind:"$id_sensores"},{$match:{"id_cansat":"C1001"}},{$project:{_id:0,id_sensores:1}}])

var listUsrs []Usr

func init() {
  listUsrs = usrRequest("")
}

func GetAllUsrs() []Usr {
  listUsrs = usrRequest("")
	return listUsrs
}

func GetUsr(Id_usr string) (usr []Usr, err error) {
  listUsrs = usrRequest(Id_usr)
  if listUsrs == nil {
    fmt.Printf("\nNo se encontro IoTDevice !\n\n")
    return nil, errors.New("IotDevice no existe")
  } else {
    return listUsrs, nil
  }
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
