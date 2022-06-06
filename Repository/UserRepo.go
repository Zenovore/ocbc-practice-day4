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

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbname   = "ocbc-training"
// )

// var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
// 	"password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)
// var db, err = sql.Open("postgres", psqlInfo)

// func init() {

// if err != nil {
// 	return nil, err
// }
// fmt.Println("db Connected")
// }

// defer db.Close()

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

// func GetEmployeeById(id int) model.Employee {
// 	var employee model.Employee
// 	row := db.QueryRow("SELECT * FROM employee WHERE id = $1", id)
// 	row.Scan(&employee.Id, &employee.Name, &employee.Dob, &employee.Address, &employee.Jobs, &employee.JoinDate)
// 	return employee
// }

// func AddEmployee(newEmployee model.Employee) error {
// 	// var newEmployee model.Employee
// 	// if err := c.BindJSON(&newEmployee); err != nil {
// 	// 	return
// 	// }
// 	sqlStatement := `
// 	INSERT INTO employee (name,dob,address,jobs,joinDate)
// 	VALUES ($1,$2,$3,$4,$5)`

// 	row, err := db.Query(sqlStatement, newEmployee.Name, newEmployee.Dob, newEmployee.Address, newEmployee.Jobs, newEmployee.JoinDate)
// 	fmt.Println(row)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func deleteEmployee(id int) error {
// 	sqlStatement := `DELETE FROM employees WHERE id = $1`
// 	_, err := db.Exec(sqlStatement, id)
// 	return err
// }

// func updateEmployee(id int, employee model.Employee) error {
// 	sqlStatement := `UPDATE employees SET name = $2, dob = $3, address = $4, jobs = $5, join_date = $6 WHERE id = $1`
// 	_, err := db.Exec(sqlStatement, id, employee.Name, employee.Dob, employee.Address, employee.Jobs, employee.JoinDate)
// 	return err
// }
