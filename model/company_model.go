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
	rows, err := companyModel.DB.Query("select id, name, active, last_payment from company where id = ?", companyID)
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
