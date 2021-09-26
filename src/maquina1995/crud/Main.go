package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Hechizo struct {
	id     int
	nombre string
	mana   int
}

func main() {
	db := connect()
	defer close(db)

	hechizo := Hechizo{
		0,
		"Rayo Paralizante",
		13,
	}

	error := create(db, &hechizo)
	if error != nil {
		log.Println(error)
	}

	hechizos, error := findAll(db)
	if error == nil {
		for _, hechizo := range hechizos {
			log.Println(hechizo)
		}
	}

	// http.Handl eFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	// http.ListenAndServe(":8080", nil)
}

func findAll(db *sql.DB) ([]Hechizo, error) {

	countQuery, error := db.Query("SELECT COUNT(*) AS counter FROM HECHIZO")
	if error != nil {
		return nil, error
	}

	rows, error := db.Query("SELECT * FROM HECHIZO")
	if error != nil {
		return nil, error
	}

	var id, mana, counter int
	var nombre string
	hechizo := Hechizo{}

	countQuery.Scan(&counter)
	fmt.Print(counter)

	respuesta := make([]Hechizo, counter)

	for rows.Next() {
		errorSelect := rows.Scan(&id, &nombre, &mana)
		if errorSelect != nil {
			log.Println(errorSelect)
			continue
		}

		hechizo.id = id
		hechizo.nombre = nombre
		hechizo.mana = mana
		respuesta = append(respuesta, hechizo)
	}

	return respuesta, nil
}

func create(db *sql.DB, hechizo *Hechizo) error {
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

	id, _ := query.LastInsertId()
	log.Printf("Se ha insertado una columna en la tabla Hechizo con id:%v !!!", id)
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
