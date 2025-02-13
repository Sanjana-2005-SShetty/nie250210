package main

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

 type Flight struct {
	Id string
	Num string
	Airline string
	Source string
	Destination string
	Capacity int
	Price float32
 }
 func readallflights(c *gin.Context){
	flights:=[]Flight{{Id: "201",Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Destination",Capacity: 180,Price: 15000.0},
	{Id: "201",Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Destination",Capacity: 180,Price: 15000.0},
}

	c.JSON(http.StatusOK,flights)
 }
 func readflightbyid(c *gin.Context){
	id:=c.Param("id")
	flight:=Flight{Id: id,Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Destination",Capacity: 180,Price: 15000.0}

	c.JSON(http.StatusOK,flight)
 }
 func createflight(c *gin.Context){
	var jbodyflight Flight
	err:=c.BindJSON(&jbodyflight)
	if err!=nil{
	c.JSON(http.StatusInternalServerError,
	gin.H{"error":"Server Error."+err.Error()})
	return
	}
	createflight:=Flight{Id:"1001",Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Abu Dhabi",Capacity: 180,Price:15000.0}
	c.JSON(http.StatusCreated,
	gin.H{"messege":"Flight Created Sucessfully.","flight":createflight})
 }
 func updatedflight(c *gin.Context){
	id := c.Param("id")
	var jbodyflight Flight
	err:=c.BindJSON(&jbodyflight)
	if err!=nil{
	c.JSON(http.StatusInternalServerError,
	gin.H{"error":"Server Error."+err.Error()})
	return
	}
	updatedflight:=Flight{Id:id,Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Abu Dhabi",Capacity: 180,Price:15000.0}
	c.JSON(http.StatusCreated,
	gin.H{"messege":"Flight updated Sucessfully.","flight":updatedflight})
 }
 func deleteflight(c *gin.Context){
	//id := c.Param("id")
	c.JSON(http.StatusOK,
		gin.H{"messege":"Flight delete Sucessfully."})

	}

 func main(){
	// flight1:=Flight{Id: "201",Num: "AI 845",Airline: "Air India",Source: "Mumbai",Destination: "Destination",Capacity: 180,Price: 15000.0}
	// fmt.Println(flight1)
	r:=gin.Default()
	r.GET("/flights",readallflights)
	r.GET("/flights/:id",readflightbyid)
	r.POST("/flights",createflight)
	r.PUT("/flights/:id",updatedflight)
	r.DELETE("/flights/:id",deleteflight)

	
	r.Run(":8080")
 }