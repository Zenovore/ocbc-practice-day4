package repository

import (
	"database/sql"
	"fmt"

	model "github.com/Zenovore/ocbc-practice-day4/Model"
	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db}
}

func (r *Repository) GetAllEmployees() ([]model.Employee, error) {

	var employees []model.Employee
	rows, err := r.DB.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Dob, &employee.Address, &employee.Jobs, &employee.JoinDate); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
		fmt.Println(employee)
	}
	return employees, nil
}

func (r *Repository) GetEmployeeById(id int) model.Employee {
	var employee model.Employee
	row := r.DB.QueryRow("SELECT * FROM employee WHERE id = $1", id)
	// fmt.Println(row)
	row.Scan(&employee.Id, &employee.Name, &employee.Dob, &employee.Address, &employee.Jobs, &employee.JoinDate)
	// fmt.Println(employee)
	return employee
}

func (r *Repository) AddEmployee(newEmployee model.Employee) error {
	sqlStatement := `
	INSERT INTO employee (name,dob,address,jobs,joinDate)
	VALUES ($1,$2,$3,$4,$5)`

	row, err := r.DB.Query(sqlStatement, newEmployee.Name, newEmployee.Dob, newEmployee.Address, newEmployee.Jobs, newEmployee.JoinDate)
	fmt.Println(row)
	if err != nil {
		return nil
	}
	return err
}

func (r *Repository) DeleteEmployee(id int) error {
	sqlStatement := `DELETE FROM employee WHERE id = $1`
	_, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

// func updateEmployee(id int, employee model.Employee) error {
// 	sqlStatement := `UPDATE employees SET name = $2, dob = $3, address = $4, jobs = $5, join_date = $6 WHERE id = $1`
// 	_, err := db.Exec(sqlStatement, id, employee.Name, employee.Dob, employee.Address, employee.Jobs, employee.JoinDate)
// 	return err
// }
