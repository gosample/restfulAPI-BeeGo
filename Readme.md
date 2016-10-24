# RESTful Web Services using Bee-API with Go language

Plataforma de desarrollo usando el lenguaje de programación Go, el driver para MongoDB de Go y el framework BeeGo.

Esta plataforma es el servicio Restful (Get, Post, Delete, Update) para el servicio frontend: https://github.com/juanmorenomotta/angular-beego

# Requisitos

## Golang:

Para mayor información consultar: https://golang.org/doc/install

## BeeGo:

WebFramework para el lenguaje Go, para mayor información consultar: http://beego.me/quickstart

Instalar:

    go get github.com/astaxie/beego
    go get github.com/beego/bee

## Mango (mgo):

Driver de MongoDB para el lenguaje Go, para mayor información consultar: https://labix.org/mgo

Instalar:

    go get gopkg.in/mgo.v2
    go get gopkg.in/check.v1

### Otra referencia (personal)

Revisar: https://github.com/Jenazad/developConfig/ en la carpeta Golang.

**Levantar el los servicios RESTful web**

    bee run

**Ingresamos a la aplicación**

    http://localhost:6568


