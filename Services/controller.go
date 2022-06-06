package services

import (
	"fmt"

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
