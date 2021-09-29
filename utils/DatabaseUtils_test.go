package utils

import "testing"

func TestConnect(t *testing.T) {
	db := Connect()
	defer db.Close()

	if db.Ping() != nil {
		t.Errorf("La conexión a la base de datos ha fallado")
	}
}

func TestClose(t *testing.T) {
	db := Connect()

	if Close(db) != nil {
		t.Errorf("la desconexión a la base de datos ha fallado")
	}
}
