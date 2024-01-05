package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/behatris-fiorentini/curso_alura_go_api_rest/database"
	"github.com/behatris-fiorentini/curso_alura_go_api_rest/models"
	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var p models.Personality
	database.DB.First(&p, id)
	if p.ID == 0 {
		json.NewEncoder(w).Encode("personalidade não encontrada")
	}

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var p models.Personality
	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var personality models.Personality
	database.DB.First(&personality, id)

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	database.DB.Save(personality)

	json.NewEncoder(w).Encode(personality)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}
