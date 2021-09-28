package repository

import (
	"database/sql"
	"fmt"
	"log"
	entity "maquina1995/crud/entities"
)

const (
	tabla = "HECHIZO"
)

func Delete(db *sql.DB, id int64) error {

	fmt.Println()
	log.Println("---- Delete ----")

	query, error := db.Query("DELETE FROM ? WHERE ID=?", tabla, id)
	if error != nil {
		log.Println(error)
		return error
	}
	defer query.Close()
	log.Println("Se ha Eliminado de la BD el hechizo con id: ", id)
	return nil
}

func Update(db *sql.DB, hechizoUpdate *entity.Hechizo) error {

	fmt.Println()
	log.Println("---- Consulta Update ----")

	query, error := db.Exec("UPDATE ? SET NOMBRE=?, MANA=? WHERE ID=?", tabla, hechizoUpdate.Id, hechizoUpdate.Nombre, hechizoUpdate.Mana)
	if error != nil {
		log.Println(error)
		return error
	}

	rows, _ := query.RowsAffected()
	log.Printf("Se ha actualizado el item de base de datos %v, registros afectados: %v", hechizoUpdate, rows)

	return nil

}

func FindById(db *sql.DB, idHechizo int64) error {

	fmt.Println()
	log.Println("---- Consulta FindById ----")

	query, error := db.Query("SELECT * FROM ? WHERE ID=?", tabla, idHechizo)
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

	hechizo := entity.Hechizo{Id: id, Nombre: nombre, Mana: mana}
	log.Println("Se ha obtenido de la BD: ", hechizo)

	return nil
}

func FindAll(db *sql.DB) ([]entity.Hechizo, error) {

	fmt.Println()
	log.Println("---- Consulta FindAll ----")

	var id int64
	var mana, resultCounter int
	var nombre string

	countQuery, error := db.Query("SELECT COUNT(*) AS counter FROM ?", tabla)
	if error != nil {
		return make([]entity.Hechizo, 0), error
	}
	defer countQuery.Close()

	countQuery.Scan(&resultCounter)

	query, error := db.Query("SELECT * FROM ?", tabla)
	if error != nil {
		return make([]entity.Hechizo, 0), error
	}

	hechizos := make([]entity.Hechizo, resultCounter)

	for query.Next() {
		errorSelect := query.Scan(&id, &nombre, &mana)
		if errorSelect != nil {
			log.Println(errorSelect)
			continue
		}

		hechizosDb := entity.Hechizo{Id: id, Nombre: nombre, Mana: mana}
		hechizos = append(hechizos, hechizosDb)
	}

	return hechizos, nil
}

func Create(db *sql.DB, hechizo *entity.Hechizo) (int64, error) {

	fmt.Println()
	log.Println("---- Consulta Create ----")

	stmt, error := db.Prepare("INSERT INTO ? (NOMBRE, MANA) VALUES (?, ?)")
	if error != nil {
		log.Println(error)
		return 0, error
	}
	defer stmt.Close()

	query, error := stmt.Exec(tabla, hechizo.Nombre, hechizo.Mana)
	if error != nil {
		log.Println(error)
		return 0, error
	}

	id, _ := query.LastInsertId()
	log.Printf("Se ha insertado una columna en la tabla %v con id:%v !!!", tabla, id)
	return id, nil
}
