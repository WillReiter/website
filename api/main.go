package main

import (
	apihandlers "main/handlers"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConfig, err := utils.ParseDBConfig("config/dbconfig.yaml")
	if err != nil {
		panic(err)
	}

	dbConnection, err := utils.GetDatabaseConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/companies", apihandlers.GetCompanies(dbConnection))
	r.GET("/search/companies", apihandlers.SearchCompanies(dbConnection))
	r.GET("/discounts", apihandlers.GetDiscounts(dbConnection))

	r.POST("/login", apihandlers.AuthenticateUser(dbConnection))
	r.POST("/login/create", apihandlers.CreateUser(dbConnection))

	r.Run() // listen and serve on 0.0.0.0:8080
}
