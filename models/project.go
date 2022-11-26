package models

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type _Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectRequest struct {
	Name        string `json:"project_name"`
	Description string `json:"descript"`
}

// @Title Create Project
// @Description Creates Projects & Returns a Project based on the request
// @Param request body _Project true "Create Project Request"
// @Router /projects/create [post]
func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project Project
	json.NewDecoder(r.Body).Decode(&project)
	DB.Create(&project)
	json.NewEncoder(w).Encode(project)
}

// @Title Get All Project
// @Description Get All Projects based on the request
// @Router /projects [get]
func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var projects []Project
	DB.Find(&projects)
	json.NewEncoder(w).Encode(projects)
}

// @Title Get _Project by ID
// @Description Get Projects by ID based on the request
// @Param id path int true "Get Projects by ID Request"
// @Router /projects/{id} [get]
func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.First(&project, params["id"])
	json.NewEncoder(w).Encode(project)
}

// @Title Update _Project By ID
// @Description Update Projects based on the request
// @Param request body _Project true "Update Project Request"
// @Param id path int true "Update Project Request"
// @Router /projects/{id} [patch]
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.First(&project, params["id"])
	json.NewDecoder(r.Body).Decode(&project)
	DB.Save(&project)
	json.NewEncoder(w).Encode(project)
}

// @Title Delete Project By ID
// @Description Delete Projects based on the request
// @Param id path int true "Delete Project Request"
// @Router /projects/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.Delete(&project, params["id"])
	json.NewEncoder(w).Encode("This Client %v has been successfully deleted")
}
