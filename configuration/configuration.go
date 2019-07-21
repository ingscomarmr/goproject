package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //solo se ejecuta la funcion init que es la requerida para gorm
	"github.com/jinzhu/gorm"
)

//guarda y transporta la configuracion
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	DataBase string
	Dialect  string
	PathPublicKeyRsa string
	PathPrivateKeyRsa string
}

//Obtiene la configuracion del archivo config.json
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json") //abrimos el archivo
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //cerramos el archivo antes de finalizar

	//llenar la struct de configurarion
	err = json.NewDecoder(file).Decode(&c)
	log.Println(c)
	return c
}

//regresa una conexion de base de datos
func GetConnectionDb() *gorm.DB {
	config := GetConfiguration() //obtener la configuracion
	//obtener la cadena de conexion
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Server,
		config.Port,
		config.DataBase)

	//regresar la conexion
	db, err := gorm.Open(config.Dialect, dns)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
