package database

import (
	"main/database/database_models"
	"main/handlers/handler_models"

	"gorm.io/gorm"
)

func GetCompanies(conn gorm.DB) ([]database_models.Companies, error) {
	var rowData []database_models.Companies

	err := conn.Find(&rowData)
	if err != nil {
		return rowData, err.Error
	}

	return rowData, nil
}

func SearchCompanies(conn gorm.DB, params handler_models.CompanyParams) ([]database_models.Companies, error) {
	var rowData []database_models.Companies

	err := conn.Where(&params).Find(&rowData)
	if err != nil {
		return rowData, err.Error
	}
	return rowData, nil
}

func GetDiscounts(conn gorm.DB) ([]database_models.Discounts, error) {
	var rowData []database_models.Discounts

	err := conn.Find(&rowData)
	if err != nil {
		return rowData, err.Error
	}

	return rowData, nil
}

func AddUser(conn gorm.DB, login handler_models.Authentication) (bool, error) {
	var rowData = database_models.Auth{
		Username: *login.Username,
		Password: *login.Password,
	}

	result := conn.Create(&rowData)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func GetUser(conn gorm.DB, login handler_models.Authentication) (database_models.Auth, error) {
	var rowData database_models.Auth

	err := conn.Where("? = username", login.Username).Find(&rowData)
	if err != nil {
		return rowData, err.Error
	}

	return rowData, nil
}
