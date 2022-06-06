package services

import (
	// "fmt"
	// "strconv"

	// model "github.com/Zenovore/ocbc-practice-day4/Model"
	"fmt"

	config "github.com/Zenovore/ocbc-practice-day4/Config"
	model "github.com/Zenovore/ocbc-practice-day4/Model"
	repository "github.com/Zenovore/ocbc-practice-day4/Repository"
	"github.com/gin-gonic/gin"
)

// type Users interface{
// GetAllEmployees() ([]Employee, error) LIST FUNCTION YG DIBUTUHKAN
// }

type UserService struct {
	repository *repository.Repository
}

type UserServiceInterface interface {
	GetAllEmployee(c *gin.Context) []model.Employee
}

func NewService() UserServiceInterface {
	db := config.GetDB()
	repo := repository.NewRepository(db)
	fmt.Println("db Connected")
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) GetAllEmployee(c *gin.Context) []model.Employee {
	// func (s *UserService) GetAllEmployee() {
	item, _ := s.repository.GetAllEmployees()
	if item != nil {
		return item
	}
	return nil
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
