package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// CompanyModel Database model for Company Table
type CompanyModel struct {
	DB *sql.DB
}

// CreateCompany Create a new company record in the database
func (companyModel *CompanyModel) CreateCompany(newCompany *entity.Company) (*entity.Company, error) {
	_, err := companyModel.DB.Exec(`
	insert into company (name, is_active, last_payment)
	values (?, ?, ?)`, newCompany.CompanyName, newCompany.CompanyIsActive, newCompany.LastPayment)
	if err != nil {
		return nil, err
	}

	row := companyModel.DB.QueryRow(`
	select id from company where name = ?`, newCompany.CompanyName)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&newCompany.CompanyID)

	if err != nil {
		return nil, err
	}

	return newCompany, nil

}

// GetCompanyByID Get Company Entity by ID
func (companyModel *CompanyModel) GetCompanyByID(companyID *int64) (*entity.Company, error) {
	rows, err := companyModel.DB.Query("select id, name, is_active, last_payment from company where id = ?", companyID)
	if err != nil {
		return nil, err
	}
	var company entity.Company

	for rows.Next() {
		err := rows.Scan(&company.CompanyID, &company.CompanyName, &company.CompanyIsActive, &company.LastPayment)
		if err != nil {
			return nil, err
		}
	}
	return &company, nil
}

// GetCompanyByName Get Company entity by Name
func (companyModel *CompanyModel) GetCompanyByName(companyName string) (*entity.Company, error) {
	rows, err := companyModel.DB.Query("select id, name, is_active, last_payment from company where name = ?", companyName)
	if err != nil {
		return nil, err
	}
	var company entity.Company

	for rows.Next() {
		err := rows.Scan(&company.CompanyID, &company.CompanyName, &company.CompanyIsActive, &company.LastPayment)
		if err != nil {
			return nil, err
		}
	}
	return &company, nil
}

// GetEmployees active users of a company for CompanyModel - restrict it to per company
func (companyModel *CompanyModel) GetEmployees(userCompany *entity.Company) (*[]entity.User, error) {

	rows, err := companyModel.DB.Query(
		`select id, password, first_name, last_name, last_login, is_superuser, username,
		email, is_staff, is_active, date_joined, company_id, birth_date 
		from user where (company_id = ?) and (is_active=true)`, userCompany.CompanyID)
	if err != nil {
		return nil, err
	}

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return &users, nil
}

// GetCompanyManagers Get company managers returns a pointer to a slice of entity.User
func (companyModel *CompanyModel) GetCompanyManagers(company *entity.Company) (*[]entity.User, error) {
	rows, err := companyModel.DB.Query(`
	select id, password, first_name, last_name, last_login, is_superuser, username,
	email, is_staff, is_active, date_joined, birth_date, is_owner
	from user where (company_id = ?) and (is_manager = true)`, company.CompanyID)

	if err != nil {
		return nil, err
	}

	var managers []entity.User

	for rows.Next() {
		var manager entity.User
		manager.IsManager = true
		err := rows.Scan(&manager.UserID, &manager.Password,
			&manager.FirstName, &manager.LastName,
			&manager.LastLogin, &manager.IsSuperuser,
			&manager.Username, &manager.Email,
			&manager.IsStaff, &manager.IsActive,
			&manager.DateJoined, &manager.BirthDate, &manager.IsOwner)
		if err != nil {
			return nil, err
		}
		manager.Company = company
		managers = append(managers, manager)
	}

	return &managers, nil

}

// GetCompanyOwners Get Company Owners from the database
func (companyModel *CompanyModel) GetCompanyOwners(company *entity.Company) (*[]entity.User, error) {
	rows, err := companyModel.DB.Query(`
	select id, password, first_name, last_name, last_login, is_superuser, username,
	email, is_staff, is_active, date_joined, birth_date, is_manager
	from user where (company_id = ?) and (is_owner = true)`, company.CompanyID)

	if err != nil {
		return nil, err
	}

	var owners []entity.User

	for rows.Next() {
		var owner entity.User
		owner.IsOwner = true
		err := rows.Scan(&owner.UserID, &owner.Password,
			&owner.FirstName, &owner.LastName,
			&owner.LastLogin, &owner.IsSuperuser,
			&owner.Username, &owner.Email,
			&owner.IsStaff, &owner.IsActive,
			&owner.DateJoined, &owner.BirthDate,
			&owner.IsManager)
		if err != nil {
			return nil, err
		}
		owner.Company = company
		owners = append(owners, owner)
	}
	return &owners, nil
}
