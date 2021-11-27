package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type LoginInformation struct {
	ID             int
	Username       string
	Password       string
	AccountType    string
	AccountStatus  string
	AccountUpdated time.Time
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

func CheckPassenger(db *sql.DB, username string, password string) bool {
	var LoginInfo LoginInformation
	Query := "SELECT * FROM GoRide_Logins.LoginInformation WHERE Username = '" + username + "' AND Password = '" + password + "'"
	fmt.Println(Query)
	results := db.QueryRow(Query)

	switch err := results.Scan(&LoginInfo.ID, &LoginInfo.Username, &LoginInfo.Password, &LoginInfo.AccountType, &LoginInfo.AccountStatus, &LoginInfo.AccountUpdated); err {
	case sql.ErrNoRows:
		db.Close()
		return false
	case nil:
		db.Close()
		return true
	default:
		panic(err)
	}
}

//e.GET("/login/api/passenger", GetPassengerLogin)
func GetPassengerLogin(c echo.Context) error {
	// Get name and password
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(username, password)

	LoginDB := OpenLoginsDB()

	passengerExists := CheckPassenger(&LoginDB, username, password)

	if passengerExists {
		return c.String(http.StatusOK, "true")
	} else {
		return c.String(http.StatusNotAcceptable, "false")
	}

}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Routes
	//Send GetRequest in form-data with keys 'Username' and 'Password'
	e.GET("/login/api/passenger", GetPassengerLogin)

	e.Logger.Fatal(e.Start(":8000"))
}
