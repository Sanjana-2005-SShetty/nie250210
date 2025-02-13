package main

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

 type Employee struct {
	Id string
	Name string
	/*Airline string
	Source string
	Destination string
	Capacity int
	Price float32*/
 }
 func readallemployee(c *gin.Context){
	employee:= Employee{Id: "201","Name":"ram"}
	
}

	c.JSON(http.StatusOK,employee)
 }
 func main() {
    r := gin.Default()

    // Define your routes
    r.GET("/Employee", handleEmployee)

    r.Run(":8080")
}

func handleEmployee(c *gin.Context) {
    // Your handler logic
    c.JSON(http.StatusOK, gin.H{"message": "Employee route"})
}
