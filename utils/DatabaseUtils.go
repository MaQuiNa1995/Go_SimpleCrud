package utils

import (
	"database/sql"
	"log"
)

/*
	Se abre la conexión a la base de datos
*/
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "crud_user:pass@tcp(localhost:3306)/go_application")
	if err != nil {
		log.Panic("Error en la conexión a la base de datos...", err)
	}
	log.Println("Se abre la conexión a la base de datos")
	return db
}

/*
	Se cierra la conexión a la base de datos
*/
func Close(database *sql.DB) error {

	if database == nil {
		return nil
	}

	log.Println("Se cierra la conexión")
	return database.Close()
}
