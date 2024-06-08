// employee/employee.go
package employee

import (
	"errors"
	"sync"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

var (
	mu        sync.Mutex
	employees = make(map[int]Employee)
	nextID    = 1
)

func CreateEmployee(name, position string, salary float64) Employee {
	mu.Lock()
	defer mu.Unlock()

	employee := Employee{
		ID:       nextID,
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	employees[nextID] = employee
	nextID++
	return employee
}

func GetEmployeeByID(id int) (Employee, error) {
	mu.Lock()
	defer mu.Unlock()

	employee, exists := employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return employee, nil
}

func UpdateEmployee(id int, name, position string, salary float64) (Employee, error) {
	mu.Lock()
	defer mu.Unlock()

	employee, exists := employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}

	employee.Name = name
	employee.Position = position
	employee.Salary = salary
	employees[id] = employee
	return employee, nil
}

func DeleteEmployee(id int) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := employees[id]; !exists {
		return errors.New("employee not found")
	}
	delete(employees, id)
	return nil
}
