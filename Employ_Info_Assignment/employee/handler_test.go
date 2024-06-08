package employee

import (
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	emp := CreateEmployee("John Doe", "Engineer", 50000)
	if emp.ID != 1 {
		t.Errorf("expected ID 1, got %d", emp.ID)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	CreateEmployee("John Doe", "Engineer", 50000)
	emp, err := GetEmployeeByID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if emp.Name != "John Doe" {
		t.Errorf("expected Name 'John Doe', got %s", emp.Name)
	}
}

func TestUpdateEmployee(t *testing.T) {
	CreateEmployee("John Doe", "Engineer", 50000)
	emp, err := UpdateEmployee(1, "Jane Doe", "Manager", 60000)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if emp.Name != "Jane Doe" {
		t.Errorf("expected Name 'Jane Doe', got %s", emp.Name)
	}
}

func TestDeleteEmployee(t *testing.T) {
	CreateEmployee("John Doe", "Engineer", 50000)
	err := DeleteEmployee(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = GetEmployeeByID(1)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
