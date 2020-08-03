package middleware

import (
	"authorizater/models" // models package where User schema is defined
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable
	"strconv"  // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// AccountPost create a user in the postgres db
func AccountPost(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// create an empty user of type models.User
	var user models.AccountIdDetail

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println("Unable to decode the request body.  %v", err)
		failRes := response{
			Message: "Unable to decode the request body.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// call insert user function and pass the user
	insertID := insertUser(user)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// AccountIDGet will return a single user by its id
func AccountIDGet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Unable to convert the string into int.  %v", err)
		failRes := response{
			Message: "Unable to convert the string into int.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// call the getUser function with user id to retrieve a single user
	user, err := getUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(user)
}

// AccountsGet will return all the users
func AccountsGet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// get all the users in the db
	users, err := getAllUsers()

	if err != nil {
		log.Println("Unable to get all user. %v", err)
		failRes := response{
			Message: "Unable to get all user.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

// AccountIDPut update user's detail in the postgres db
func AccountIDPut(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Unable to convert the string into int.  %v", err)
		failRes := response{
			Message: "Unable to convert the string into int.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// create an empty user of type models.User
	var user models.AccountIdDetail

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println("Unable to decode the request body. %v", err)
		failRes := response{
			Message: "Unable to decode the request body.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// call update user to update the user
	updatedRows := updateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// AccountIDDelete delete user's detail in the postgres db
func AccountIDDelete(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Unable to convert the string into int.  %v", err)
		failRes := response{
			Message: "Unable to convert the string into int.",
		}
		json.NewEncoder(w).Encode(failRes)
	}

	// call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// CustomerAccountGet create user by email
func CustomerAccountGet(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := r.URL.Query()

	email, ok := params["username"]
	if ok {
		user, err := getUserByEmail(email[0])

		if err != nil {
			log.Fatalf("Unable to get user. %v", err)
		}

		// send the response
		json.NewEncoder(w).Encode(user)
	} else {
		log.Println("Param 'username' is missing")
		failRes := response{
			Message: "Param 'username' is missing",
		}
		json.NewEncoder(w).Encode(failRes)
	}

}

//------------------------- handler functions ----------------
// insert one user in the DB
func insertUser(user models.AccountIdDetail) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning id will return the id of the inserted user
	sqlStatement := `INSERT INTO users (email, company, dashboard, spc, adjustment_guiding, parameter_failure_pediction, process_status_analysis) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Email, user.Company, user.Dashboard, user.Spc, user.AdjustmentGuiding, user.ParameterFailurePediction, user.ProcessStatusAnalysis).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query.  %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one user from the DB by its id
func getUser(id int64) (models.AccountIdDetail, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var user models.AccountIdDetail

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Email, &user.Dashboard, &user.Spc, &user.AdjustmentGuiding, &user.ParameterFailurePediction, &user.ProcessStatusAnalysis, &user.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get one user from the DB by its email
func getUserByEmail(email string) (models.AccountIdDetail, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var user models.AccountIdDetail

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE email=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, email)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Email, &user.Dashboard, &user.Spc, &user.AdjustmentGuiding, &user.ParameterFailurePediction, &user.ProcessStatusAnalysis, &user.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get all users
func getAllUsers() ([]models.AccountIdDetail, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var users []models.AccountIdDetail

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.AccountIdDetail

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Email, &user.Dashboard, &user.Spc, &user.AdjustmentGuiding, &user.ParameterFailurePediction, &user.ProcessStatusAnalysis, &user.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}

// update user in the DB
func updateUser(id int64, user models.AccountIdDetail) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET email=$2, company=$3, dashboard=$4, spc=$5, adjustment_guiding=$6, parameter_failure_pediction=$7, process_status_analysis=$8 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Email, user.Company, user.Dashboard, user.Spc, user.AdjustmentGuiding, user.ParameterFailurePediction, user.ProcessStatusAnalysis)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
