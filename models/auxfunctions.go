package models

import(
	"fmt"
  "gopkg.in/mgo.v2"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
)

var ip = "127.0.0.1"
var portdb = "27017"
var db = "smartcity"
var usr_collection = "cansats"
var sensor_collection = "sensores"
var sensado_collection = "datosensado"

var url = ip+":"+portdb

func removeDuplicatesUnordered(elements []string) []string {
  encountered := map[string]bool{}
  for v:= range elements {
  	encountered[elements[v]] = true
  }
  result := []string{}
  for key, _ := range encountered {
  	result = append(result, key)
  }
  return result
}

func usrRequest(Id string) []Usr{
  var listUsrs []Usr
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(usr_collection)
  if Id == ""{
    err = c.Find(bson.M{ }).Sort("id_cansat").All(&listUsrs)
  } else {
    err = c.Find(bson.M{"id_cansat": Id}).All(&listUsrs)
  }
  js, __ := json.Marshal(listUsrs)
  //fmt.Printf("%s\n",listUsrs[0].Modelo)
  //fmt.Println(reflect.TypeOf(listUsrs))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de IoTDevices:\n%s\n\n",js)
  return listUsrs
}

func sensoresRequest(Tipo string) []Sensor{
  var listSensores []Sensor
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(sensor_collection)
  if Tipo == ""{
    err = c.Find(bson.M{ }).Sort("id_sensor").All(&listSensores)
  } else {
    err = c.Find(bson.M{"tipo_sensor": Tipo}).Sort("id_sensor").All(&listSensores)
  }
  js, __ := json.Marshal(listSensores)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de sensores:\n%s\n\n",js)
  return listSensores
}

func sensorTipoRequest(Id_user string) sensorTipo{
  var listTipo1 []sensorTipo1
  var listTipo2 []string
  var listTipo sensorTipo
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)

  c := session.DB(db).C("sensores")
  c1 := session.DB(db).C("sensorestipo")
  c2 := session.DB(db).C("sensores")

  err = c.Find(bson.M{ }).Distinct("tipo_sensor",&listTipo.Tipo_sensor)
  err = c1.Find(bson.M{ }).All(&listTipo1)
  err = c2.Find(bson.M{}).Distinct("tipo_sensor",&listTipo2)

  js , __ := json.Marshal(listTipo)
  js1, __ := json.Marshal(listTipo1)
  js2 , __ := json.Marshal(listTipo2)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js)
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js1)
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js2)
  return listTipo
}

func datoSensadoRequest(Tipo_sensor string, Id_usr string) []datoSensado{
  var listDato []datoSensado
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(sensado_collection)
  if Tipo_sensor == "" && Id_usr ==""{
    err = c.Find(bson.M{ }).All(&listDato)
  } else if Tipo_sensor !=""{
    if Id_usr == ""{
      err = c.Find(bson.M{"tipo_sensor": Tipo_sensor}).All(&listDato)
    } else if Id_usr != ""{
      err = c.Find(bson.M{"tipo_sensor": Tipo_sensor,"id_cansat": Id_usr}).All(&listDato)
    }
  } else if Tipo_sensor == "" && Id_usr != "" {
    err = c.Find(bson.M{"id_cansat": Id_usr}).All(&listDato)
  }
  js, __ := json.Marshal(listDato)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de los datos de sensores:\n%s\n\n",js)
  return listDato
}
