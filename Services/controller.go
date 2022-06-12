package services

import (
	"fmt"

	model "github.com/Zenovore/ocbc-practice-day4/Model"
	"github.com/gin-gonic/gin"
)

func GetAllEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		// func (s *UserService) GetAllEmployee() {
		// item := UserServiceInterface.GetAllEmployee()
		item := userInterface.GetAllEmployee(c)
		fmt.Println(item)
		if item != nil {
			c.JSON(200, item)
		} else {
			c.JSON(404, gin.H{"message": "not found"})
		}
	}
}

func GetEmployeeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		item := userInterface.GetEmployeeById(c)
		if (item == model.Employee{}) {
			c.JSON(404, gin.H{"message": "id not found"})
		} else {
			c.JSON(200, item)
		}
	}
}

func AddEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := userInterface.AddEmployee(c)
		if err != nil {
			c.JSON(400, gin.H{"message": "error"})
		} else {
			c.JSON(200, gin.H{"message": "success"})
		}
	}
}

func DeleteEmployeeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := userInterface.DeleteEmployeeById(c)
		if err != nil {
			c.JSON(400, gin.H{"message": "error"})
		} else {
			c.JSON(200, gin.H{"message": "success"})
		}
	}
}
