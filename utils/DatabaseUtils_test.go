package utils

import (
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	db := Connect()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		log.Println(err)
		t.Error("La conexión a la base de datos ha fallado", err)
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
