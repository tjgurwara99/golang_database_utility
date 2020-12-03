package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// UserModel for database connection to user table
type UserModel struct {
	DB *sql.DB
}

// CreateUser Create a new user record in the database
func (userModel *UserModel) CreateUser(newUser *entity.User) (*entity.User, error) {
	_, err := userModel.DB.Exec(`
	insert into user (first_name, last_name, last_login, is_superuser, username, password,
	email, is_staff, is_active, date_joined, birth_date, is_manager, is_owner, company_id )
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newUser.FirstName, newUser.LastName, newUser.LastLogin, newUser.IsSuperuser,
		newUser.Username, newUser.Password, newUser.Email, newUser.IsStaff, newUser.IsActive,
		newUser.DateJoined, newUser.BirthDate, newUser.IsManager, newUser.IsOwner, newUser.CompanyID)
	if err != nil {
		return nil, err
	}
	return newUser, nil

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
			&user.LastLogin, &user.IsSuperuser,
			&user.Username, &user.Email,
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
			&user.LastLogin, &user.IsSuperuser,
			&user.Username, &user.Email,
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

	err = user.CheckPassword(password)

	if err != nil {
		return nil, err
	}

	return user, nil

}
