package utils

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := Connect()
	defer db.Close()

	if db.Ping() != nil {
		t.Error("La conexión a la base de datos ha fallado")
	}
}

func TestClose(t *testing.T) {
	db := Connect()

	if Close(db) != nil {
		t.Error("la desconexión a la base de datos ha fallado")
	}
}
