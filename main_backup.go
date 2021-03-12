package main

import (
	"fmt"
	"time"

	"github.com/tjgurwara99/golang_database_utility/config"
	"github.com/tjgurwara99/golang_database_utility/entity"
	"github.com/tjgurwara99/golang_database_utility/model"
	"github.com/tjgurwara99/golang_database_utility/service"
)

func main() {
	db, err := service.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	companyModel := model.CompanyModel{DB: db}

	company := entity.Company{
		CompanyName:     "Crowntech",
		CompanyIsActive: true,
		LastPayment:     time.Now(),
	}

	_, err = companyModel.CreateCompany(&company)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Company ID: %d, Company Name: %s, Company Is Active: %v, Last Payment: %v", company.CompanyID, company.CompanyName, company.CompanyIsActive, company.LastPayment)

}
