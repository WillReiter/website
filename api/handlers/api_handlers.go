package apihandlers

import (
	"errors"
	"log"
	"main/database"
	"main/database/database_models"
	"main/handlers/handler_models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCompanies(dbConnection gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var companies []database_models.Companies

		companies, err := database.GetCompanies(dbConnection)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				c.JSON(200, gin.H{"message": "No matching companies found!"})
				return
			}
			c.JSON(500, gin.H{"message": "Error occured fetching data!"})
			return
		}

		c.JSON(200, companies)

	}
}

func SearchCompanies(dbConnection gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var companies []database_models.Companies
		var companyParams handler_models.CompanyParams
		c.BindQuery(companyParams)

		log.Printf("params: %+v", companyParams)

		companies, err := database.SearchCompanies(dbConnection, companyParams)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				c.JSON(200, gin.H{"message": "No matching companies found!"})
				return
			}
			c.JSON(500, gin.H{"message": "Error occured fetching data!"})
			return
		}
		c.JSON(200, companies)
	}
}

func GetDiscounts(dbConnection gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var discounts []database_models.Discounts

		discounts, err := database.GetDiscounts(dbConnection)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				c.JSON(200, gin.H{"message": "No matching deals found!"})
				return
			}
			c.JSON(500, gin.H{"message": "Error occured fetching data!"})
			return
		}

		c.JSON(200, discounts)

	}
}

func AuthenticateUser(dbConnection gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var auth database_models.Auth
		var login handler_models.Authentication

		c.BindJSON(&login)

		auth, err := database.GetUser(dbConnection, login)
		if err != nil {
			if errors.As(err, &gorm.ErrRecordNotFound) {
				c.JSON(200, gin.H{"message": "Invalid username!"})
				return
			}
			c.JSON(500, gin.H{"message": "Error occured authenticating login data!"})
			return
		}
		if !strings.EqualFold(auth.Password, *login.Password) {
			c.JSON(500, gin.H{"message": "Invalid password!"})
			return
		}
		c.JSON(200, gin.H{"message": "Successfully validated login!"})
	}
}

func CreateUser(dbConnection gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var login handler_models.Authentication

		if err := c.BindJSON(&login); err != nil {
			return
		}

		isUserCreated, err := database.AddUser(dbConnection, login)
		if err != nil {
			c.JSON(500, gin.H{"message": err})
			return
		}

		if !isUserCreated {
			c.JSON(500, gin.H{"message": "Unable to create user!"})
			return
		}

		c.JSON(200, gin.H{"message": "User successfully created!"})
	}
}
