package main

import (
	"flag"
	"log"

	"./migration"
)

func main()  {
	var migrate string

	//para recibir los parametros al ejecutar el archivo ejecutable
	flag.StringVar(&migrate,"migrate","no","Genera la migracion a la base de datos") //se ejecutara de la siguiente forma ./goproject --migrate yes
	flag.Parse() //parsea los parametros

	if migrate == "yes"{
		log.Println("Iniciando la migracion de base de datos")
		migration.MigrateDb()
		log.Println("Termino la migracion de base de datos")
	}else {
		log.Println("Iniciando sin migracion de base de datos")
	}

}
