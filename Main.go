package main

import (
	"log"
	entity "maquina1995/crud/entities"
	repo "maquina1995/crud/repository"
)

func main() {

	hechizo := entity.Hechizo{
		Id:     0,
		Nombre: "Tsunami",
		Mana:   200,
	}

	// Create
	id, _ := repo.Create(&hechizo)

	// Read
	repo.FindById(id)

	// Update
	hechizo.Nombre = "Hechizo updateado"
	hechizo.Mana = 100
	hechizo.Id = id
	repo.Update(&hechizo)

	hechizos, _ := repo.FindAll()
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}

	// Delete
	repo.Delete(hechizo.Id)

	hechizos, _ = repo.FindAll()
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}
}
