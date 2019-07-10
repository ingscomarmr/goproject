package main

import (
	"fmt"
	"log"

	"./configuration"
)

func main() {
	log.Println("Iniciar conexion de base datos")
	conDb := configuration.GetConnectionDb()
	defer conDb.Close() //con el defer cerramos la conexion
	fmt.Println("SE CONECTO A LA BASE DE DATOS")

}
