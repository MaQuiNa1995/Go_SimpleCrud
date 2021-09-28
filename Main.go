package main

import (
	"log"
	entity "maquina1995/crud/entities"
	repo "maquina1995/crud/repository"
	dbUtils "maquina1995/crud/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbUtils.Connect()
	defer dbUtils.Close(db)

	hechizo := entity.Hechizo{
		Id:     0,
		Nombre: "Tsunami",
		Mana:   200,
	}

	// Create
	id, _ := repo.Create(db, &hechizo)

	// Read
	repo.FindById(db, hechizo.Id)

	// Update
	hechizo.Nombre = "Hechizo updateado"
	hechizo.Mana = 100
	hechizo.Id = id
	repo.Update(db, &hechizo)

	hechizos, _ := repo.FindAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}

	// Delete
	repo.Delete(db, hechizo.Id)

	hechizos, _ = repo.FindAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}
}
