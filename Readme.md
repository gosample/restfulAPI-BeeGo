# RestulAPI-BeeGo

Plataforma de desarrollo usando el lenguaje de programación Go, el driver para MongoDB de Go y el framework BeeGo.

Esta plataforma es el servicio Restful (Get, Post, Delete, Update) para el servicio frontend: https://github.com/juanmorenomotta/angular-beego

# Requisitos

## Golang:
Escribir en consola:

wget https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz

Descomprimir en /usr/local
tar -C /usr/local -xzf go1.7.linux-amd64.tar.gz

escribir el en fichero /etc/profile escribir:

  export PATH=$PATH:/usr/local/go/bin
  
y añadir en el fichero .profile:

  export GOPATH=$HOME/golang

  export GOROOT=$HOME/go

  export PATH=$PATH:$GOROOT/bin

## BeeGo:
Escribir en consola

go get github.com/astaxie/beego

go get github.com/beego/bee

mover el ejecutable bee al directorio go/bin

## Mango (mgo):
Escribir en consola:

go get gopkg.in/mgo.v2

go get gopkg.in/check.v1
