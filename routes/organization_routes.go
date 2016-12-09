package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetOrganizationRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/organization", handlers.GetOrganizations).Methods("GET").Name("GetOrganization")
	router.HandleFunc("/organization", handlers.AddOrganization).Methods("POST").Name("AddOrganization")
	router.HandleFunc("/organization/addStudent", handlers.AddStudentToOrganization).Methods("POST").Name("AddStudentToOrganization")
	router.HandleFunc("/organization/removeStudent", handlers.RemoveStudentFromOrganization).Methods("DELETE").Name("RemoveStudentFromOrganization")
	router.HandleFunc("/organization/moveStudent", handlers.MoveStudentInOrganization).Methods("PUT").Name("MoveStudentInOrganization")
	return router
}