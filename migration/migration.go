package migration

import (
	"../configuration"
	"../model"
)

//clase para migrar la base de datos de codigo a base de datos
func MigrateDb() {
	conDb := configuration.GetConnectionDb()
	defer conDb.Close()

	//creando los modelos
	conDb.CreateTable(&model.User{})
	conDb.CreateTable(&model.Comment{})
	//agregamos las columnas ids y sus nombres
	conDb.CreateTable(&model.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id","user_id")

}
