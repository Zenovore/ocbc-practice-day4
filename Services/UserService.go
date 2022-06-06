package services

import (
	// "fmt"
	// "strconv"

	// model "github.com/Zenovore/ocbc-practice-day4/Model"
	repository "github.com/Zenovore/ocbc-practice-day4/Repository"
	"github.com/gin-gonic/gin"
)

// type Users interface{
// GetAllEmployees() ([]Employee, error) LIST FUNCTION YG DIBUTUHKAN
// }

// type

func GetAllEmployee(c *gin.Context) {
	item, _ := repository.GetAllEmployees()
	if item != nil {
		c.JSON(200, item)
	} else {
		c.JSON(404, gin.H{"message": "no data"})
	}
}

// func GetEmployeeById(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	fmt.Println("get employee by id")
// 	item := repository.GetEmployeeById(id)
// 	c.JSON(200, item)
// }

// func AddEmployee(c *gin.Context) {
// 	var newEmployee model.Employee
// 	if err := c.BindJSON(&newEmployee); err != nil {
// 		return
// 	}
// 	fmt.Println("insert employee")
// 	item := repository.AddEmployee(newEmployee)
// }

// func DeleteEmployeeById(c *gin.Context) {
// 	fmt.Println("delete employee id")
// }
// func UpdateEmployeeById(c *gin.Context) {
// 	fmt.Println("update employee id")
// }
