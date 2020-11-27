package services

import (
	"database/sql"
	// "log"
	//"errors"
	_ "github.com/go-sql-driver/mysql"
	//"reflect"
	//"strings"
)

// OpenDatabase returns database object if successful
func OpenDatabase(databaseProgram string, username string, password string, hostname string, databaseName string) (*sql.DB, error) {
	db, err := sql.Open(databaseProgram, username+":"+password+"@tcp("+hostname+")/"+databaseName)
	return db, err
}

// Exec executes a query
func Exec(db *sql.DB, query interface{ InsertQuery() string }) {
	db.Exec(query.InsertQuery())
}

//type queryError struct {
//	s string
//}

//fun (e *queryError) Error() string {
//	return e.s
//}

// RaiseQueryError raises an error in text format
//func RaiseQueryError(text string) error {
//	return &queryError{text}
//}

//func CreateQuery(object interface{}) Error {
//	if reflect.ValueOf(object).Kind() == reflect.Struct {
//		objectType := strings.ToLower(reflect.TypeOf(object).Name())
//		query := fmt.Sprintf("insert into %s values(", objectType)
//		value := reflect.ValueOf(object)
//		for i := 0; i < value.NumField(); i++ {
//			switch value.Field(i).Kind() {
//			case reflect.Int:
//				if i == 0 {
//					query = fmt.Sprintf("%s%d", query, value.Field(i).Int())
//				} else {
//					query = fmt.Sprintf("%s, %d", query, value.Field(i).Int())
//				}
//			case reflect.String:
//				if i == 0 {
//					query = fmt.Sprintf("%s\"%s\"", query, value.Field(i).String())
//				} else {
//					query = fmt.Sprintf("%s, \"%s\"", query, value.Field(i).String())
//				}
//			default:
//				return RaiseQueryError("Unsupported Type")
//			}
//		}
//		query = fmt.Sprintf("%s)", query)
//
//
// It seems that reflect is not a good way to go about thing when we think about maintainability
//
//
