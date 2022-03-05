package dbrepo


import (
	"root/gleam/golang/model"
	"fmt"
	"github.com/jackc/pgconn"
	"errors"
	"os"
	"time"

)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}


func (m *postgresDBRepo) SaveNewUser(u model.User, hashedPwd string) (int64, error) {

	var userID int64
	t := time.Now()
	sqlStatement := `INSERT INTO users(firstName, lastName, email, password, createdAt, updatedAt) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := m.DB.QueryRow(sqlStatement, u.FirstName, u.LastName, u.Email, hashedPwd, t, t).Scan(&userID)
		
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "unique_email" {
				fmt.Fprintf(os.Stderr, "Unable to create user email, because email already taken: %v\n", pgErr)
				
			} else {
				fmt.Fprintf(os.Stderr, "Unexpected postgres error trying to create user: %v\n", pgErr)
				
			}

		//   fmt.Println(pgErr.Message) // => syntax error at end of input
		//   fmt.Println(pgErr.Code) // => 42601
		}else {
			fmt.Fprintf(os.Stderr, "Unexpected error trying to create user mwood: %v\n", err)
			
		}
	}


	// if err != nil {
	// 	var pgerr *pgconn.PgError
	// 	if pgerr, ok := errors.As(err, &pgerr); ok {
	// 		if pgerr.ConstraintName == "unique_email" {
	// 			fmt.Fprintf(os.Stderr, "Unable to create user email, because email already taken: %v\n", pgerr)
	// 		} else {
	// 			fmt.Fprintf(os.Stderr, "Unexpected postgres error trying to create user: %v\n", pgerr)
	// 		}
	// 	} else {
	// 		fmt.Fprintf(os.Stderr, "Unexpected error trying to create user mwood: %v\n", err)
	// 	}
	// 	os.Exit(1)
	// }
	// fmt.Printf("Successfully created user with id %v\n", userID)


	return userID, err
}
