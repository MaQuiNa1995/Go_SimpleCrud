package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Se abre la conexión a la base de datos
*/
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/go_application")
	if err != nil {
		log.Panic(err)
		panic("Error en la conexión a la base de datos...")
	}
	log.Println("Se abre la conexión a la base de datos")
	return db
}

/*
	Se cierra la conexión a la base de datos
*/
func Close(database *sql.DB) error {

	if database == nil {
		log.Println("Se ha intentado cerrar la conexión cuando no se ha abierto")
		return nil
	}

	log.Println("Se cierra la conexión")
	fmt.Println("")
	return database.Close()
}
