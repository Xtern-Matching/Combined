package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
)

func GetOrganizations(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	organizations, keys, err := services.GetOrganizations(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	type Response struct {
		Keys []*datastore.Key			`json:"keys"`
		Organizations []models.Organization	`json:"organizations"`
	}
	response := Response{Keys: keys, Organizations: organizations}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	name := dat["name"].(string)

	_, err := services.NewOrganization(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetOrganizationStudents(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	org, err := services.GetOrganization(ctx,orgKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	students := make([]models.Student,0)
	keys :=make([]*datastore.Key,0)
	for _, studentRank := range org.StudentRanks {
		student, err := services.GetStudent(ctx, studentRank.Student)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		students = append(students, student)
		keys  = append(keys, studentRank.Student)
	}
	

	type Response struct {
		Keys []*datastore.Key		`json:"keys"`
		Students []models.Student	`json:"students"`
	}
	response := Response{Keys: keys, Students: students}


	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddStudentToOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = services.AddStudentToOrganization(ctx, orgKey, studentKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveStudentFromOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = services.RemoveStudentFromOrganization(ctx, orgKey, studentKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SwitchStudentsInOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey1, err := datastore.DecodeKey(dat["studentKey1"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey2, err := datastore.DecodeKey(dat["studentKey2"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = services.SwitchStudentsInOrganization(ctx, orgKey, studentKey1, studentKey2)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetCurrentOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	company, err := services.GetOrganization(ctx, orgKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}