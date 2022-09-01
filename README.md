# Introducción a Go

Este repositorio contiene una base para comenzar a realizar servicios REST con Go y puede ser utilizado como curso introductorio o como una plantilla para inicializar otros proyectos.

## Inicialización del proyecto (módulo)

Lo primero que se debe hacer es inicializar el proyecto donde vamos a trabajar, para esto creamos la carpeta del proyecto y dentro de la carpeta se ejecuta el siguiente comando:

`go mod init github.com/nombre-del-proyecto`

Este comando creará el archivo go.mod el cual contiene todos los datos del proyecto junto con todas las dependencias que vayamos agregando.

## Creación de un endpoint REST

Para exponer un endpoint rest se pueden utilizar las libreras que trae incluidas Go, sin embargo por facilidad y mayores prestaciones en los servicios se recomienda utilizar una libre externa adicional, entre las que existen se encuentra:

- [gin-gonic](https://github.com/gin-gonic/gin)

Para este proyecto de ejemplo vamos a utilizar gin-gonic, para esto primero agregamos la dependencia ejecutando el siguiente comando:

`go get -u github.com/gin-gonic/gin`

En Go las aplicaciones se ejecuta a partir de la función main que se debe encontrar en el paquete main, para esto, creamos en la raíz del proyecto el archivo main.go con el siguiente contenido:

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main() {
   r := gin.Default()
   r.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "message": "pong",
      })
   })
   r.Run() // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}
```

## Ejecución del proyecto

Para compilar y ejecutar el proyecto ejecutamos el siguiente comando:

`go run main.go`

## Ejecución de pruebas

Para ejecutar todas las pruebas del correcto se debe ejecutar el siguiente comando:

`go test ./…`

## Imagen de docker

Para crear la imagen de docker se ejecuta el siguiente comando:

`docker build -t go-rest-api:latest .`

Luego de haber creado la imagen se puede ejecutar el contenedor de la siguiente forma:

`docker run -d -p 8080:8080 go-rest-api`