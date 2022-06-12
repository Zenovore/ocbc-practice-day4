package main

import (
	services "github.com/Zenovore/ocbc-practice-day4/Services"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/employees", services.GetAllEmployee())
	// Oper Context, langsung pake ke user service aja

	router.GET("/employees/:id", services.GetEmployeeById())
	router.POST("/add", services.AddEmployee())
	router.DELETE("/employees/:id", services.DeleteEmployeeById())
	// router.PUT("/employees/:id", services.UpdateEmployeeById) // PUT FOR CREATE

	router.Run("localhost:8082")
	// NEVER SET LOCALHOST, USE ENV
}
