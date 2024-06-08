// employee/handlers.go
package employee

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employee := CreateEmployee(emp.Name, emp.Position, emp.Salary)
	json.NewEncoder(w).Encode(employee)
}

func GetEmployeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	employee, err := GetEmployeeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee, err := UpdateEmployee(id, emp.Name, emp.Position, emp.Salary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := DeleteEmployee(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ListEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		http.Error(w, "Invalid page size", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	var employeeList []Employee
	for _, emp := range employees {
		employeeList = append(employeeList, emp)
	}

	// Sort the employee list by ID
	sort.Slice(employeeList, func(i, j int) bool {
		return employeeList[i].ID < employeeList[j].ID
	})

	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(employeeList) {
		start = len(employeeList)
	}
	if end > len(employeeList) {
		end = len(employeeList)
	}

	paginatedEmployees := employeeList[start:end]
	if len(paginatedEmployees) == 0 {
		message := "No employee data available"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": message})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedEmployees)
}
