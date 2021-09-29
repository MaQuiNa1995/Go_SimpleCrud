package repository

import (
	entity "maquina1995/crud/entities"
	"testing"
)

func TestCreate(t *testing.T) {

	// Given
	var hechizo = entity.Hechizo{
		Id:     0,
		Nombre: "nombreTest",
		Mana:   50,
	}

	// When
	id, error := Create(&hechizo)

	// Then
	if id == 0 {
		t.Error("Hubo un error al intentar insertar el Hechizo en la base de datos: ", error)
	}

}
