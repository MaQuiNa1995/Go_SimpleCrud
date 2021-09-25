package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Hechizo struct {
	id     int
	nombre string
	mana   int
}

func main() {
	// log.Println("Server started on: http://localhost:8080")
	db := connect()
	defer close(db)

	hechizo := Hechizo{
		0,
		"Torrente Ígneo",
		30,
	}
	insert(db, hechizo)

	// http.HandleFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	// http.ListenAndServe(":8080", nil)
}

func insert(db *sql.DB, hechizo Hechizo) error {
	stmt, error := db.Prepare("INSERT INTO HECHIZO (NOMBRE, MANA) VALUES (?, ?)")
	if error != nil {
		log.Println(error)
		return error
	}
	defer stmt.Close()

	query, err := stmt.Exec(hechizo.nombre, hechizo.mana)
	if err != nil {
		log.Println(err)
		return err
	}

	_, id := query.RowsAffected()
	log.Printf("Se han insertado %v en la tabla Hechizo !!!", id)
	return nil
}

/*
	Se abre la conexión a la base de datos
*/
func connect() *sql.DB {
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
func close(database *sql.DB) error {

	if database == nil {
		return nil
	}

	log.Println("Se cierra la conexión")
	return database.Close()
}
