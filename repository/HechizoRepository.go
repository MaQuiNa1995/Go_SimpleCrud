package repository

import (
	"fmt"
	"log"
	entity "maquina1995/crud/entities"
	dbUtils "maquina1995/crud/utils"
)

const (
	tabla = "HECHIZO"
)

/*
	Crea un Hechizo en base de datos
*/
func Create(hechizo *entity.Hechizo) (int64, error) {

	log.Println("---- Consulta Create ----")

	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	query := fmt.Sprintf("INSERT INTO %v (NOMBRE, MANA) VALUES (?, ?)", tabla)
	stmt, error := db.Prepare(query)
	if error != nil {
		log.Println(error)
		return 0, error
	}
	defer stmt.Close()

	insertQuery, error := stmt.Exec(hechizo.Nombre, hechizo.Mana)
	if error != nil {
		log.Println(error)
		return 0, error
	}

	id, _ := insertQuery.LastInsertId()
	log.Printf("Se ha insertado una columna en la tabla %v con id:%v !!!", tabla, id)
	return id, nil
}

/*
	Cuenta los registros en la tabla de Hechizo
*/
func Count() int {

	fmt.Println()
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	var resultCounter int

	query := fmt.Sprintf("SELECT COUNT(*) AS counter FROM %v", tabla)

	countQuery, error := db.Query(query)
	if error != nil {
		return 0
	}
	defer countQuery.Close()

	countQuery.Scan(&resultCounter)

	return resultCounter
}

/*
	Busca por id en tabla de Hechizo
*/
func FindById(idHechizo int64) (entity.Hechizo, error) {

	log.Println("---- Consulta FindById ----")
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	query := fmt.Sprintf("SELECT * FROM %v WHERE ID=?", tabla)
	selectQuery, error := db.Query(query, idHechizo)
	if error != nil {
		log.Println(error)
		return entity.Hechizo{}, error
	}
	defer selectQuery.Close()

	var id int64
	var mana int
	var nombre string

	selectQuery.Next()
	error = selectQuery.Scan(&id, &nombre, &mana)
	if error != nil {
		log.Println(error)
		return entity.Hechizo{}, error
	}

	hechizo := entity.Hechizo{Id: id, Nombre: nombre, Mana: mana}
	log.Println("Se ha obtenido de la BD: ", hechizo)

	return hechizo, nil
}

/*
	Recupera todos los registros de la tabla de Hechizo
*/
func FindAll() ([]entity.Hechizo, error) {

	log.Println("---- Consulta FindAll ----")
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	var id int64
	var mana, counter int
	var nombre string

	query := fmt.Sprintf("SELECT * FROM  %v", tabla)
	resultQuery, error := db.Query(query)

	hechizos := make([]entity.Hechizo, 0)
	if error != nil {
		return hechizos, error
	}

	for resultQuery.Next() {

		errorSelect := resultQuery.Scan(&id, &nombre, &mana)
		if errorSelect != nil {
			log.Println(errorSelect)
			continue
		}
		counter++
		hechizosDb := entity.Hechizo{Id: id, Nombre: nombre, Mana: mana}
		hechizos = append(hechizos, hechizosDb)
	}
	return hechizos, nil
}

/*
	Actualiza un Hechizo por id
*/
func Update(hechizoUpdate *entity.Hechizo) error {

	log.Println("---- Consulta Update ----")
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	query := fmt.Sprintf("UPDATE %v SET NOMBRE=?, MANA=? WHERE ID=?", tabla)

	updateQuery, error := db.Exec(query, hechizoUpdate.Nombre, hechizoUpdate.Mana, hechizoUpdate.Id)
	if error != nil {
		log.Println(error)
		return error
	}

	rows, _ := updateQuery.RowsAffected()
	log.Printf("Se ha actualizado el item de base de datos %v, registros afectados: %v", hechizoUpdate, rows)

	return nil
}

/*
	Elimina hechizo por id
*/
func Delete(id int64) error {

	fmt.Println()
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	log.Println("---- Delete ----")

	query := fmt.Sprintf("DELETE FROM %v WHERE ID=?", tabla)
	deleteQuery, error := db.Query(query, id)
	if error != nil {
		log.Println(error)
		return error
	}
	defer deleteQuery.Close()
	log.Println("Se ha Eliminado de la BD el hechizo con id: ", id)
	return nil
}
