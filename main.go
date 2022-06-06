package main

import (
	services "github.com/Zenovore/ocbc-practice-day4/Services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/employees", services.GetAllEmployee)
	// router.GET("/employees/:id", services.GetEmployeeById)
	// router.DELETE("/employees/:id", services.DeleteEmployeeById)
	// router.PUT("/employees/:id", services.UpdateEmployeeById) // PUT FOR CREATE
	// router.POST("/add", services.AddEmployee)

	router.Run("localhost:8082")
}
