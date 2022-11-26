package main

import (
	"log"
	"net/http"
	"restapi/docs"
	"restapi/models"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// @title Projects API
// @version 1.0
// @description This is a Service for Managing Projects
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html\
// @host      localhost:8080
// @BasePath /
func initializeRouter() {
	r := mux.NewRouter()
	//for JWT
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	//	r.PathPrefix("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:1323/swagger/doc.json").Methods("GET")

	r.HandleFunc("/login", models.Login).Methods("POST")
	r.HandleFunc("/home", models.Home).Methods("GET")
	r.HandleFunc("/refresh", models.Refresh).Methods("POST")

	//For USERS
	r.HandleFunc("/users", models.GetUser).Methods("GET")
	//r.HandleFunc("/users/{id}", models.GetUser).Methods("GET")
	// r.HandleFunc("/users/{id}", models.DeleteUser).Methods("DELETE")
	// r.HandleFunc("/users/{id}", models.UpdateUser).Methods("PATCH")
	r.HandleFunc("/register", models.CreateUser).Methods("POST")
	// 	//For Authors
	r.HandleFunc("/clients", models.GetClients).Methods("GET")
	r.HandleFunc("/clients/{id}", models.GetClient).Methods("GET")
	r.HandleFunc("/clients/{id}", models.DeleteClient).Methods("DELETE")
	r.HandleFunc("/clients/{id}/edit", models.UpdateClient).Methods("PATCH")
	r.HandleFunc("/clients/create", models.CreateClient).Methods("POST")
	// 	//For Books
	r.HandleFunc("/projects/create", models.CreateProject).Methods("POST")
	r.HandleFunc("/projects", models.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", models.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}/edit", models.UpdateProject).Methods("PATCH")
	r.HandleFunc("/projects/{id}", models.DeleteProject).Methods("DELETE")
	docs.SwaggerInfo.Host = "localhost:8080"
	log.Fatal(http.ListenAndServe(":8080", r))
	// }
}
func main() {
	models.InitialMigration()
	initializeRouter()
}
