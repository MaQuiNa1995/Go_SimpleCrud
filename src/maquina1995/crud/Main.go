package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Hechizo struct {
	id     int64
	nombre string
	mana   int
}

func main() {
	db := connect()
	defer close(db)

	hechizo := Hechizo{
		0,
		"Carámbano Punzante",
		5,
	}

	// Create
	id, _ := create(db, &hechizo)

	// Update
	hechizo.nombre = "Hechizo updateado"
	hechizo.mana = 100
	hechizo.id = id
	update(db, &hechizo)

	hechizos, _ := findAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}

	// Read
	findById(db, hechizo.id)

	// Delete
	delete(db, hechizo.id)

	hechizos, _ = findAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
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

func delete(db *sql.DB, id int64) error {

	fmt.Println()
	log.Println("---- Delete ----")

	query, error := db.Query("DELETE FROM HECHIZO WHERE ID=?", id)
	if error != nil {
		log.Println(error)
		return error
	}
	defer query.Close()
	log.Println("Se ha Eliminado de la BD el hechizo con id: ", id)
	return nil
}

func update(db *sql.DB, hechizoUpdate *Hechizo) error {

	fmt.Println()
	log.Println("---- Consulta Update ----")

	query, error := db.Exec("UPDATE HECHIZO SET NOMBRE=?, MANA=? WHERE ID=?", hechizoUpdate.id, hechizoUpdate.nombre, hechizoUpdate.mana)
	if error != nil {
		log.Println(error)
		return error
	}

	rows, _ := query.RowsAffected()
	log.Printf("Se ha actualizado el item de base de datos %v, registros afectados: %v", hechizoUpdate, rows)

	return nil

}

func findById(db *sql.DB, idHechizo int64) error {

	fmt.Println()
	log.Println("---- Consulta FindById ----")

	query, error := db.Query("SELECT * FROM  HECHIZO WHERE ID=?", idHechizo)
	if error != nil {
		return error
	}
	defer query.Close()

	var id int64
	var mana int
	var nombre string

	query.Next()
	error = query.Scan(&id, &nombre, &mana)
	if error != nil {
		log.Println(error)
		return error
	}

	hechizo := Hechizo{id, nombre, mana}
	log.Println("Se ha obtenido de la BD: ", hechizo)

	return nil
}

func findAll(db *sql.DB) ([]Hechizo, error) {

	fmt.Println()
	log.Println("---- Consulta FindAll ----")

	var id int64
	var mana, resultCounter int
	var nombre string

	countQuery, error := db.Query("SELECT COUNT(*) AS counter FROM HECHIZO")
	if error != nil {
		return make([]Hechizo, 0), error
	}
	defer countQuery.Close()

	countQuery.Scan(&resultCounter)

	query, error := db.Query("SELECT * FROM HECHIZO")
	if error != nil {
		return make([]Hechizo, 0), error
	}

	hechizos := make([]Hechizo, resultCounter)

	for query.Next() {
		errorSelect := query.Scan(&id, &nombre, &mana)
		if errorSelect != nil {
			log.Println(errorSelect)
			continue
		}

		hechizosDb := Hechizo{id, nombre, mana}
		hechizos = append(hechizos, hechizosDb)
	}

	return hechizos, nil
}

func create(db *sql.DB, hechizo *Hechizo) (int64, error) {

	fmt.Println()
	log.Println("---- Consulta Create ----")

	stmt, error := db.Prepare("INSERT INTO HECHIZO (NOMBRE, MANA) VALUES (?, ?)")
	if error != nil {
		log.Println(error)
		return 0, error
	}
	defer stmt.Close()

	query, error := stmt.Exec(hechizo.nombre, hechizo.mana)
	if error != nil {
		log.Println(error)
		return 0, error
	}

	id, _ := query.LastInsertId()
	log.Printf("Se ha insertado una columna en la tabla Hechizo con id:%v !!!", id)
	return id, nil
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
