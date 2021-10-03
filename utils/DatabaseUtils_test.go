package utils

import (
	"log"
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
	err := Close(db)
	if err != nil {
		log.Println(err)
		t.Error("la desconexión a la base de datos ha fallado", err)
	}
}
