package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// CompanyModel Database model for Company Table
type CompanyModel struct {
	DB *sql.DB
}

// GetCompanyByID Get Company Entity by ID
func (companyModel *CompanyModel) GetCompanyByID(companyID *int64) (*entity.Company, error) {
	rows, err := companyModel.DB.Query("select id, name, manager_id, active, last_payment from company where id = ?", companyID)
	if err != nil {
		return nil, err
	}
	var company entity.Company

	var userModel UserModel
	userModel.DB = companyModel.DB

	for rows.Next() {
		var managerID int64
		err := rows.Scan(&company.CompanyID, &company.CompanyName, &managerID, &company.CompanyIsActive, &company.LastPayment)
		if err != nil {
			return nil, err
		}
		company.CompanyManager, err = userModel.GetCompanyManager(&managerID, &company)
		if err != nil {
			return nil, err
		}
	}
	return &company, nil
}
