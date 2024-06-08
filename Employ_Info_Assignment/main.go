// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"Employ_Info_Assignment/employee" // Import the employee package

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//restapi to adds a new employee
	router.HandleFunc("/employees", employee.CreateEmployeeHandler).Methods("POST")
	//restapi to  Retrieves an employee by ID.
	router.HandleFunc("/employees/{id}", employee.GetEmployeeByIDHandler).Methods("GET")
	//restapi to Update the details of an existing employee on bases of id
	router.HandleFunc("/employees/{id}", employee.UpdateEmployeeHandler).Methods("PUT")
	//restapi to Deletes an employee from the database
	router.HandleFunc("/employees/{id}", employee.DeleteEmployeeHandler).Methods("DELETE")
	//rest api for listing employee records with pagination
	//url : http://localhost:8080/employees?page=1&pageSize=1
	router.HandleFunc("/employees", employee.ListEmployeesHandler).Methods("GET")
	fmt.Println("8080 server started ................")

	log.Fatal(http.ListenAndServe(":8080", router))
}
