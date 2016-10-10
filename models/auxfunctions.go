package models

import(
	"fmt"
  "gopkg.in/mgo.v2"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
)

func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}
    // Create a map of all unique elements.
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
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("smartcity").C("cansats")
  err = c.Find(bson.M{ }).All(&listCansats)
  js, __ := json.Marshal(listCansats)
  //fmt.Printf("%s\n",listCansats[0].Id_sensores[1])
  //fmt.Println(reflect.TypeOf(listCansats))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de cansats:\n%s\n\n",js)
  return listCansats
}

func sensoresRequest() []Sensor{
  var listSensores []Sensor
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
  return listSensores
}

func sensorTipoRequest() []sensorTipo{
  var listTipo []sensorTipo
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("smartcity").C("sensorestipo")
  err = c.Find(bson.M{ }).All(&listTipo)
  js, __ := json.Marshal(listTipo)
  //fmt.Printf("%s\n",listCansats[0].Modelo)
  //fmt.Println(reflect.TypeOf(listCansats))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista delos tipos de sensores:\n%s\n\n",js)
  return listTipo
}

func sensorDatoRequest() []sensorDato{
  var listDato []sensorDato
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c := session.DB("smartcity").C("temperatura")
  err = c.Find(bson.M{ }).All(&listDato)
  js, __ := json.Marshal(listDato)
  //fmt.Printf("%s\n",listCansats[0].Modelo)
  //fmt.Println(reflect.TypeOf(listCansats))
  if __ != nil {
    panic(__)
  }
  fmt.Printf("\nlista de los datos de sensores:\n%s\n\n",js)
  return listDato
}
