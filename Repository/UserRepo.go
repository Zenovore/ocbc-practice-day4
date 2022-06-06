package Repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hydralisk123"
	dbname   = "ocbc-training"
)

// func dbInit() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
// }

// defer db.Close()

func getAllEmployees() ([]Employee, error) {
	var employees []Employee
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var employee Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func getEmployeeById(id int) (Employee, error) {
	var employee Employee
	err := db.QueryRow("SELECT * FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone)
	return employee, err
}

func addEmployee(newEmployee Employee) (Employee, error) {
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}
	sqlStatement := `
	INSERT INTO employee (name,dob,address,jobs,joinDate)
	VALUES ($1,$2,$3,$4,$5)`

	row, err := db.Query(sqlStatement, newEmployee.Name, newEmployee.Dob, newEmployee.Address, newEmployee.Jobs, newEmployee.JoinDate)
	fmt.Println(row)
	if err != nil {
		panic(err)
	}
}

func deleteEmployee(id int) error {
	sqlStatement := `DELETE FROM employees WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func updateEmployee(id int, employee Employee) error {
	sqlStatement := `UPDATE employees SET name = $2, dob = $3, address = $4, jobs = $5, join_date = $6 WHERE id = $1`
	_, err := db.Exec(sqlStatement, id, employee.Name, employee.Dob, employee.Address, employee.Jobs, employee.JoinDate)
	return err
}
