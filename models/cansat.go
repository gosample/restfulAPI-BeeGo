package models

import (
	"errors"
//	"strconv"
//	"time"
//  "reflect"
	"fmt"
)

type Cansat struct{
  Id_cansat   string `json:"id_cansat"`
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

var listCansats []Cansat

func init() {
  listCansats = cansatRequest("")
}

func GetAllCansats() []Cansat {
  listCansats = cansatRequest("")
	return listCansats
}

func GetCansat(Id_cansat string) (cansat []Cansat, err error) {
  listCansats = cansatRequest(Id_cansat)
  if listCansats == nil {
    fmt.Printf("\nNo se encontro Cansat !\n\n")
    return nil, errors.New("Cansat no existe")
  } else {
    return listCansats, nil
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
