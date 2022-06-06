package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Services.insertEmployee()
	// router.GET("/employees", Services. getAllEmployee)
	// router.GET("/employees/:id", getEmployeeById)
	// router.DELETE("/employees/:id", deleteEmployeeById)
	// router.PUT("/employees/:id", updateEmployeeById) // PUT FOR CREATE
	// router.POST("/add", addEmployee)

	router.Run("localhost:8080")
}
