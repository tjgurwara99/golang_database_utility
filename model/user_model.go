package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// UserModel for database connection to user table
type UserModel struct {
	DB *sql.DB
}

// GetCompanyUsers active users of a company for UserModel - restrict it to per company
func (userModel *UserModel) GetCompanyUsers(userCompany entity.Company) (*[]entity.User, error) {
	rows, err := userModel.DB.Query("select first_name, last_name, email from user where (company_id = ?) and (is_active=true) ", userCompany.CompanyID)
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

// GetUserByID Get User by UserID
func (userModel *UserModel) GetUserByID(userID int64) (*entity.User, error) {
	rows, err := userModel.DB.Query(`
	select id, password, first_name, last_name, last_login, is_superuser, username,
	email, is_staff, is_active, date_joined, company_id, birth_date
	from user where id = ?`, userID)

	if err != nil {
		return nil, err
	}
	var user entity.User

	var companyModel CompanyModel
	companyModel.DB = userModel.DB

	var companyID int64

	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Password,
			&user.FirstName, &user.LastName,
			&user.LastLogin, &user.IsSuperUser,
			&user.UserName, &user.Email,
			&user.IsStaff, &user.IsActive,
			&user.DateJoined, &companyID,
			&user.BirthDate)

		if err != nil {
			return nil, err
		}

		company, err2 := companyModel.GetCompanyByID(&companyID)
		if err != nil {
			return nil, err2
		}
		user.Company = company
	}
	return &user, nil

}

// GetCompanyManager Get company manager
func (userModel *UserModel) GetCompanyManager(managerID *int64, company *entity.Company) (*entity.User, error) {
	rows, err := userModel.DB.Query(`
	select id, password, first_name, last_name, last_login, is_superuser, username,
	email, is_staff, is_active, date_joined, birth_date
	from user where id = ?`, managerID)

	if err != nil {
		return nil, err
	}

	var manager entity.User

	for rows.Next() {
		err := rows.Scan(&manager.UserID, &manager.Password,
			&manager.FirstName, &manager.LastName,
			&manager.LastLogin, &manager.IsSuperUser,
			&manager.UserName, &manager.Email,
			&manager.IsStaff, &manager.IsActive,
			&manager.DateJoined, &manager.BirthDate)
		if err != nil {
			return nil, err
		}
		manager.Company = company
	}

	return &manager, nil

}

//GetUserByUsername Get User by Username
func (userModel *UserModel) GetUserByUsername(username string) (*entity.User, error) {
	rows, err := userModel.DB.Query(
		`select id, password, first_name, last_name, last_login, is_superuser, username,
		email, is_staff, is_active, date_joined, company_id, birth_date
		from user where username = ?`, username)
	if err != nil {
		return nil, err
	}

	var user entity.User

	var companyModel CompanyModel

	companyModel.DB = userModel.DB

	var companyID int64

	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Password,
			&user.FirstName, &user.LastName,
			&user.LastLogin, &user.IsSuperUser,
			&user.UserName, &user.Email,
			&user.IsStaff, &user.IsActive,
			&user.DateJoined, &companyID,
			&user.BirthDate)

		if err != nil {
			return nil, err
		}

		company, err2 := companyModel.GetCompanyByID(&companyID)
		if err != nil {
			return nil, err2
		}
		user.Company = company
	}
	return &user, nil
}

// Authenticate authentication
func (userModel UserModel) Authenticate(username, password string) (*entity.User, error) {
	if username == "" {
		return nil, entity.ErrUsernamePassword
	}
	user, err := userModel.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = user.ValidatePassword(password)

	if err != nil {
		return nil, err
	}

	return user, nil

}
