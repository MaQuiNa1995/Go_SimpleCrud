package main

import (
	"log"
	entity "maquina1995/crud/entities"
	repo "maquina1995/crud/repository"
	utils "maquina1995/crud/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := utils.Connect()
	defer utils.Close(db)

	hechizo := entity.Hechizo{
		Id:     0,
		Nombre: "Tsunami",
		Mana:   200,
	}

	// Create
	id, _ := repo.Create(db, &hechizo)

	// Update
	hechizo.Nombre = "Hechizo updateado"
	hechizo.Mana = 100
	hechizo.Id = id
	repo.Update(db, &hechizo)

	hechizos, _ := repo.FindAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}

	// Read
	repo.FindById(db, hechizo.Id)

	// Delete
	repo.Delete(db, hechizo.Id)

	hechizos, _ = repo.FindAll(db)
	for _, hechizo := range hechizos {
		log.Println(hechizo)
	}

	// http.Handl eFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	// http.ListenAndServe(":8080", nil)
}
