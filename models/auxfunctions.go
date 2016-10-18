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

func cansatRequest() []Cansat{
  var listCansats []Cansat
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(usr_collection)
  err = c.Find(bson.M{ }).All(&listCansats)
  js, __ := json.Marshal(listCansats)
  //fmt.Printf("%s\n",listCansats[0].Modelo)
  //fmt.Println(reflect.TypeOf(listCansats))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de cansats:\n%s\n\n",js)
  return listCansats
}

func sensoresRequest() []Sensor{
  var listSensores []Sensor
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(sensor_collection)
  err = c.Find(bson.M{ }).All(&listSensores)
  js, __ := json.Marshal(listSensores)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de sensores:\n%s\n\n",js)
  return listSensores
}

func sensorTipoRequest() []sensorTipo{
  var listTipo []sensorTipo
  var listTipo1 []string
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C("sensorestipo")
  c1 := session.DB(db).C("sensores")
  err = c1.Find(bson.M{ }).Distinct("tipo_sensor",&listTipo1)
  err = c.Find(bson.M{ }).All(&listTipo)
  js1 , __ := json.Marshal(listTipo1)
  js, __ := json.Marshal(listTipo)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js)
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js1)
  return listTipo
}

func datoSensadoRequest() []datoSensado{
  var listDato []datoSensado
  session, err := mgo.Dial(url)
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(db).C(sensado_collection)
  err = c.Find(bson.M{ }).All(&listDato)
  js, __ := json.Marshal(listDato)
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de los datos de sensores:\n%s\n\n",js)
  return listDato
}
