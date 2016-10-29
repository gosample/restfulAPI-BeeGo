# RESTful Web Services using Bee-Go Framework API with Go language

Implementación de un API de desarrollo usando el lenguaje de programación Go, el driver para MongoDB de Go y el framework BeeGo.

Esta plataforma es el servicio Restful (Get, Post, Delete, Update) para el servicio web [webAppServer-Go](https://github.com/juanmorenomotta/angular-beego)

# Requisitos

## Go language

Revisar [Golang](https://golang.org/dl/) para instalar Go language en tu maquina.

o tambien revisar el repositorio [developConfig/Golang](https://github.com/Jenazad/developConfig/tree/master/golang) para ver como instalar go language.

## BeeGo:

WebFramework para el lenguaje Go, para mayor información consultar la web [Beego](http://beego.me/quickstart)

Instalar:

    go get github.com/astaxie/beego
    go get github.com/beego/bee

## Mango (mgo):

Driver de MongoDB para el lenguaje Go, para mayor información consultar la web de [mgo](https://labix.org/mgo)

Instalar:

    go get gopkg.in/mgo.v2
    go get gopkg.in/check.v1

### Otra referencia (personal)

Revisar el repositorio [developConfig/Golang](https://github.com/Jenazad/developConfig/tree/master/golang) para ver como instalar go language y demas drivers.

###Levantar el los servicios RESTful web

    bee run

###Ingresamos a la aplicación

    http://localhost:6568


