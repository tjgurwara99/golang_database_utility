package test

import (
	"testing"
	"time"

	"github.com/tjgurwara99/golang_database_utility/config"
	"github.com/tjgurwara99/golang_database_utility/entity"
	"github.com/tjgurwara99/golang_database_utility/model"
	"github.com/tjgurwara99/golang_database_utility/service"
)

func TestNewUser(t *testing.T) {
	var username, password, firstName, lastName, email string
	var isSuperuser, isStaff, isActive, isManager, isOwner bool
	birthDate := time.Now()

	companyName := "Apple"
	companyIsActive := true
	lastPayment := time.Now()
	company, err := entity.NewCompany(&companyName, &companyIsActive, &lastPayment)
	if err != nil {
		t.Errorf(err.Error())
	}

	username = "test"
	password = "blahblah123"
	firstName = "test"
	lastName = "test"
	email = "test@test.test"
	isSuperuser = false
	isStaff = false
	isActive = true
	isManager = false
	isOwner = false
	birthDate = time.Now()
	user, err := entity.NewUser(username, password, firstName,
		lastName, email, isSuperuser, isStaff, isActive, isManager,
		isOwner, company, birthDate)
	if err != nil {
		t.Errorf(err.Error())
	}

	if user.Password == password {
		t.Errorf("Something is wrong with password generate function in User entity")
	}
}

func TestCreateUser(t *testing.T) {

	db, err := service.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userModel := model.UserModel{DB: db}

	// User fields
	var username, password, firstName, lastName, email string
	var isSuperuser, isStaff, isActive, isManager, isOwner bool
	birthDate := time.Now()

	username = "test"
	password = "blahblah123"
	firstName = "test"
	lastName = "test"
	email = "test@test.test"
	isSuperuser = false
	isStaff = false
	isActive = true
	isManager = false
	isOwner = false
	birthDate = time.Now()

	companyModel := model.CompanyModel{DB: db}

	company, err := companyModel.GetCompanyByName("Apple")
	if err != nil {
		t.Errorf(err.Error())
	}
	user, err := entity.NewUser(username, password, firstName,
		lastName, email, isSuperuser, isStaff, isActive, isManager,
		isOwner, company, birthDate)

	// Test Only When Database doesn't contain the record.
	// Beats the purpose of having tests - can be solved
	// if we make a delete record function
	user, err = userModel.CreateUser(user)

	if err != nil {
		t.Errorf(err.Error())
	}

}
