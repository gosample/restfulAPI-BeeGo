package models

type sensorTipo struct{
  Tipo_sensor []string `json:"tipo_sensor"`
}

type sensorTipo1 struct{
  Prefijo     string `json:"prefijo"`
  Tipo_sensor string `json:"tipo_sensor"`
}

var listTipo sensorTipo

func init(){
  listTipo = sensorTipoRequest()
}

func GetAllSensorTipo() sensorTipo{
  listTipo = sensorTipoRequest()
  return listTipo
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
