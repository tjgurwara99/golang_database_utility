package test

import (
	"testing"
	"time"

	"github.com/tjgurwara99/golang_database_utility/config"
	"github.com/tjgurwara99/golang_database_utility/entity"
	"github.com/tjgurwara99/golang_database_utility/model"
	"github.com/tjgurwara99/golang_database_utility/service"
)

// TestNewCompany Tests that the NewCompany Constructor works as expected
func TestNewCompany(t *testing.T) {
	companyName := "Apple"
	companyIsActive := true
	lastPayment := time.Now()
	company, _ := entity.NewCompany(&companyName, &companyIsActive, &lastPayment)
	testCompany := entity.Company{
		CompanyName:     companyName,
		CompanyIsActive: companyIsActive,
		LastPayment:     lastPayment,
	}
	if testCompany.CompanyName != company.CompanyName {
		t.Errorf(`company and newCompany's CompanyNames do not match`)
	}
	if testCompany.CompanyIsActive != company.CompanyIsActive {
		t.Errorf("company and newCompany's CompanyIsActive field do not mathc")
	}
	if testCompany.LastPayment != company.LastPayment {
		t.Errorf("company and testCompany's LastPayments do not match")
	}
}

// TestCreateCompany Tests whether the CompanyModel's CreatCompany Function works
func TestCreateCompany(t *testing.T) {
	db, err := service.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	companyModel := model.CompanyModel{DB: db}
	// Company Entity Creation
	companyName := "Apple"
	companyIsActive := true
	lastPayment := time.Now()
	company, _ := entity.NewCompany(&companyName, &companyIsActive, &lastPayment)

	// Test Only When Database Doesn't contain the record
	// Beats the purpose of having tests. Can be solved
	// if we make a delete record function
	company, err = companyModel.CreateCompany(company)

	if err != nil {
		t.Errorf(err.Error())
	}

	if company.CompanyID == 0 {
		t.Errorf("The company ID was not set")
	}
}

func TestGetCompanyByID(t *testing.T) {
	db, err := service.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	companyModel := model.CompanyModel{DB: db}
	var companyID int64 = 1
	_, err = companyModel.GetCompanyByID(&companyID)
	if err != nil {
		t.Errorf(err.Error())
	}

}
