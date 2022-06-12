package services

import (
	// "fmt"
	// "strconv"

	// model "github.com/Zenovore/ocbc-practice-day4/Model"
	"fmt"
	"strconv"

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
	GetEmployeeById(c *gin.Context) model.Employee
	AddEmployee(c *gin.Context) error
	DeleteEmployeeById(c *gin.Context) error
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

func (s *UserService) GetEmployeeById(c *gin.Context) model.Employee {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("get employee by id")
	item := s.repository.GetEmployeeById(id)
	return item
}

func (s *UserService) AddEmployee(c *gin.Context) error {
	var newEmployee model.Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		return nil
	}
	fmt.Println("insert employee")
	item := s.repository.AddEmployee(newEmployee)
	return item
}

func (s *UserService) DeleteEmployeeById(c *gin.Context) error {
	fmt.Println("delete employee id")
	id, _ := strconv.Atoi(c.Param("id"))
	err := s.repository.DeleteEmployee(id)
	return err
}

// func UpdateEmployeeById(c *gin.Context) {
// 	fmt.Println("update employee id")
// }
