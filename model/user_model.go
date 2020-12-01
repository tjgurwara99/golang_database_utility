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

// GetCompanyManagers Get company managers returns a pointer to a slice of entity.User
func (userModel *UserModel) GetCompanyManagers(company *entity.Company) (*[]entity.User, error) {
	rows, err := userModel.DB.Query(`
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
func (userModel *UserModel) GetCompanyOwners(company *entity.Company) (*[]entity.User, error) {
	rows, err := userModel.DB.Query(`
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
