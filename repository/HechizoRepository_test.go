package repository

import (
	"fmt"
	entity "maquina1995/crud/entities"
	utils "maquina1995/crud/utils"
	"testing"
)

func TestCreate(t *testing.T) {

	// Given
	hechizo := createHechizo()

	// When
	id, error := Create(&hechizo)

	// Then
	utils.AssertNotEquals(t, id, 0)
	utils.AssertNotError(t, error)

}

func TestFindById(t *testing.T) {

	// Given
	hechizo := createHechizo()
	id := executeTestInsert(&hechizo)

	// When
	record, error := FindById(id)

	// Then
	utils.AssertNotError(t, error)
	utils.AssertEquals(t, id, record.Id)
	utils.AssertEquals(t, hechizo.Mana, record.Mana)
	utils.AssertEquals(t, hechizo.Nombre, record.Nombre)
}

func TestFindAll(t *testing.T) {

	// Given
	hechizos, _ := FindAll()
	numeroRegistros := len(hechizos)

	hechizo := createHechizo()
	executeTestInsert(&hechizo)
	hechizo.Nombre = utils.GenerarCadena(15)
	executeTestInsert(&hechizo)
	hechizo.Nombre = utils.GenerarCadena(15)
	executeTestInsert(&hechizo)

	// When
	records, error := FindAll()

	// Then
	utils.AssertNotError(t, error)
	utils.AssertEquals(t, len(records), numeroRegistros+3)
}

func TestUpdate(t *testing.T) {

	// Given
	hechizo := createHechizo()
	id := executeTestInsert(&hechizo)

	hechizo.Id = id
	hechizo.Mana = utils.GenerarNumero()
	hechizo.Nombre = utils.GenerarCadena(15)

	// When
	error := Update(&hechizo)

	// Then
	hechizoComparar, _ := FindById(id)

	utils.AssertNotError(t, error)
	utils.AssertEquals(t, hechizoComparar, hechizo)
}

func TestDelete(t *testing.T) {

	// Given
	hechizo := createHechizo()
	id := executeTestInsert(&hechizo)

	// When
	error := Delete(id)

	// Then
	utils.AssertNotError(t, error)
	entidad, _ := FindById(id)
	utils.AssertEquals(t, entity.Hechizo{}, entidad)
}

func createHechizo() entity.Hechizo {
	return entity.Hechizo{
		Nombre: utils.GenerarCadena(15),
		Mana:   utils.GenerarNumero(),
	}
}

func executeTestInsert(hechizo *entity.Hechizo) int64 {
	db := utils.Connect()
	defer utils.Close(db)

	query := fmt.Sprintf("INSERT INTO %v (NOMBRE, MANA) VALUES ('%v', %v)", tabla, hechizo.Nombre, hechizo.Mana)
	stmt, _ := db.Exec(query)
	id, _ := stmt.LastInsertId()

	return id
}
