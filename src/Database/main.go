package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type LoginInformation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountDetails struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	AccountType    string    `json:"accounttype"`
	AccountStatus  string    `json:"accountstatus"`
	AccountUpdated time.Time `json:"accountupdated"`
}

func OpenPassengersDB() sql.DB {
	//Open Passengers Database
	Passengerdb, Passengerdberr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/GoRide_Passengers")

	if Passengerdberr != nil {
		panic(Passengerdberr.Error())
	} else {
		fmt.Println("GoRide_Passengers Database Opened!")
	}

	return *Passengerdb
}

func OpenLoginsDB() sql.DB {
	//Open Passengers Database
	Loginsdb, Logindberr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/GoRide_Logins?parseTime=true")

	if Logindberr != nil {
		panic(Logindberr.Error())
	} else {
		fmt.Println("GoRide_Logins Database Opened!")
	}

	return *Loginsdb
}

func OpenDriversDB() sql.DB {
	//Open Passengers Database
	Driversdb, Driverdberr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/GoRide_Drivers")

	if Driverdberr != nil {
		panic(Driverdberr.Error())
	} else {
		fmt.Println("GoRide_Drivers Database Opened!")
	}

	return *Driversdb
}

func OpenTripsDB() sql.DB {
	//Open Passengers Database
	Tripsdb, Tripdberr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/GoRide_Trips")

	if Tripdberr != nil {
		panic(Tripdberr.Error())
	} else {
		fmt.Println("GoRide_Trips Database Opened!")
	}

	return *Tripsdb
}

func CheckAccount(db *sql.DB, username string, password string) [2]string {
	var LoginInfo AccountDetails
	var returnResult [2]string
	Query := "SELECT * FROM GoRide_Logins.LoginInformation WHERE Username = '" + username + "' AND Password = '" + password + "'"
	fmt.Println(Query)
	results := db.QueryRow(Query)

	switch err := results.Scan(&LoginInfo.ID, &LoginInfo.Username, &LoginInfo.Password, &LoginInfo.AccountType, &LoginInfo.AccountStatus, &LoginInfo.AccountUpdated); err {
	case sql.ErrNoRows:
		db.Close()
		returnResult[0] = "false"
		returnResult[1] = "NULL"
		return returnResult
	case nil:
		db.Close()
		log.Printf("Row Get!")
		returnResult[0] = "true"
		returnResult[1] = LoginInfo.AccountType
		return returnResult
	default:
		panic(err)
	}
}

//e.Post("/login/api", Checkuser)
func Checkuser(c echo.Context) error {
	// Get name and password
	LoginDetails := LoginInformation{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&LoginDetails)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		LoginDB := OpenLoginsDB()

		passengerExists := CheckAccount(&LoginDB, LoginDetails.Username, LoginDetails.Password)
		log.Println(passengerExists)

		if passengerExists[0] == "true" {
			log.Printf(passengerExists[0])
			return c.String(http.StatusOK, passengerExists[1])
		} else {
			return c.String(http.StatusNotAcceptable, passengerExists[0])
		}
	}

}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Routes
	//Send Get Request in form-data with keys 'Username' and 'Password'
	e.POST("/api/checkuser", Checkuser)

	e.Logger.Fatal(e.Start(":8001"))
}
