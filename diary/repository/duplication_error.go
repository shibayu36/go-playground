package repository

import "github.com/go-sql-driver/mysql"

type DuplicationError struct {
	message string
}

func NewDuplicationError(message string) *DuplicationError {
	return &DuplicationError{message: message}
}

func IsDuplicationError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
		return true
	} else {
		return false
	}
}

func (e *DuplicationError) Error() string {
	return e.message
}
