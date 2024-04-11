package routes

import (
	"encoding/json"
	"net/http"

	"github.com/pciruelos/go-gorm-restapi/db"
	"github.com/pciruelos/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
