package employeecontroller

import (

	"net/http"
	"virtualmachine/models"
	"virtualmachine/helper"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"strconv"
	"encoding/json"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError


func Index(w http.ResponseWriter, r *http.Request){
	var employees[]models.Employee

	if err:= models.DB.Find(&employees).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	ResponseJson(w, http.StatusOK, employees)


}
func Show(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var employee models.Employee
	if err := models.DB.First(&employee, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Employee no found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, employee)
}
func Create(w http.ResponseWriter, r *http.Request){
	var employee models.Employee

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&employee).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, employee)
	
}
func Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var employee models.Employee

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if  models.DB.Where("id = ?", id).Updates(&employee).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, " failed to updating employee")
		return
	}

	employee.Id = id

	ResponseJson(w, http.StatusOK, employee)

}
func Delete(w http.ResponseWriter, r *http.Request){
	input := map[string]string{"id":""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var employee models.Employee
	if models.DB.Delete(&employee, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, " employee can't deleted" )
		return
	}

	response := map[string]string{"message":"Succes to Deleted Employee"}
	ResponseJson(w, http.StatusOK, response)
	
}

