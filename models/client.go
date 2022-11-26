package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// var validate = validator.New()

type Client struct {
	gorm.Model
	Name        string `json:"client_name"`
	Description string `json:"description"`
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var client Client
	json.NewDecoder(r.Body).Decode(&client)
	var check_client Client
	DB.Table("clients").Where("name = ?", client.Name).Scan(&check_client)
	//DB.Table("users").Where("username=?, password=?",user.username, user.password).
	fmt.Printf("%#v\n", check_client)

	if check_client.Name == "" {
		DB.Create(&client)
		fmt.Fprintln(w, "new user created")
		json.NewEncoder(w).Encode(client)
	} else {
		fmt.Fprintln(w, "user already exist")
	}

}

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var clients []Client
	DB.Find(&clients)
	json.NewEncoder(w).Encode(clients)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var client Client
	DB.First(&client, params["id"])
	json.NewEncoder(w).Encode(client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var client Client
	params := mux.Vars(r)
	DB.First(&client, params["id"])
	json.NewDecoder(r.Body).Decode(&client)
	DB.Save(&client)
	json.NewEncoder(w).Encode(client)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var client Client
	DB.Delete(&client, params["id"])
	json.NewEncoder(w).Encode("The Client Has been successfully deleted")
}
