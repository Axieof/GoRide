package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type LoginInformation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DriverName struct {
	Username string `json:username`
}

type AccountDetails struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	AccountType    string    `json:"accounttype"`
	AccountStatus  string    `json:"accountstatus"`
	AccountUpdated time.Time `json:"accountupdated"`
}

type PassengerTrip struct {
	PickupLocation  string `json: pickuplocation`
	DropoffLocation string `json: dropofflocation`
	Name            string `json: passengerName`
}

type Passenger struct {
	Firstname     string `json: firstname`
	Lastname      string `json: lastname`
	Password      string `json: password`
	ContactNumber string `json: contactnumber`
	EmailAddress  string `json: emailaddress`
}

type Driver struct {
	Firstname            string `json: firstname`
	Lastname             string `json: lastname`
	Password             string `json: password`
	ContactNumber        string `json: contactnumber`
	EmailAddress         string `json: emailaddress`
	DriverIdentification string `json: driveridentification`
	LicenseNumber        string `json: licensenumebr`
}

type Trip struct {
	ID          string `json: id`
	PassengerID string `json: passengerid`
	DriverID    string `json: driverid`
	PickUp      string `json: pickup`
	DropOff     string `json: dropoff`
	TripStatus  string `json: tripstatus`
	TripStart   string `json: tripstart`
	TripEnd     string `json: tripend`
}

type Account struct {
	Username    string `json: username`
	AccountType string `json: accounttype`
}

func CreateAccount(username string, password string, accounttype string, accountstatus string) {

	log.Printf("Creating new account")

	LoginsDB := OpenLoginsDB()

	Query := "INSERT INTO LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES (?, ?, ?, ?, ?)"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	stmt, err := LoginsDB.PrepareContext(ctx, Query)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
	}
	defer stmt.Close()

	var datetime = time.Now()
	res, err := stmt.ExecContext(ctx, username, password, accounttype, accountstatus, datetime)
	if err != nil {
		log.Printf("Error %s when inserting row into passenger table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
	}
	log.Printf("%d Login Created", rows)

}

func InsertPassenger(c echo.Context) error {
	PassengerDetails := Passenger{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&PassengerDetails)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		//Insert Passenger Record
		PassengerDB := OpenPassengersDB()

		Query := "INSERT INTO Passenger(FirstName, LastName, ContactNumber, EmailAddress, AccountStatus, AccountUpdated) VALUES (?, ?, ?, ?, ?, ?)"

		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		stmt, err := PassengerDB.PrepareContext(ctx, Query)

		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()

		var datetime = time.Now()
		res, err := stmt.ExecContext(ctx, PassengerDetails.Firstname, PassengerDetails.Lastname, PassengerDetails.ContactNumber, PassengerDetails.EmailAddress, "Active", datetime)
		if err != nil {
			log.Printf("Error %s when isnerting row into passenger table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d Passenger Created", rows)

		AccountUsername := PassengerDetails.Firstname + " " + PassengerDetails.Lastname

		log.Printf(AccountUsername)

		CreateAccount(AccountUsername, PassengerDetails.Password, "Passenger", "active")

		return c.String(http.StatusAccepted, "Passenger Created!")
	}

}

func InsertDriver(c echo.Context) error {
	DriverDetails := Driver{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&DriverDetails)
	log.Printf(DriverDetails.Firstname)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		//Insert Driver Record
		DriverDB := OpenDriversDB()

		Query := "INSERT INTO Driver(FirstName, LastName, ContactNumber, EmailAddress, DriverIdentification, LicenseNumber, AccountStatus, AccountUpdated) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		stmt, err := DriverDB.PrepareContext(ctx, Query)

		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()

		var datetime = time.Now()
		res, err := stmt.ExecContext(ctx, DriverDetails.Firstname, DriverDetails.Lastname, DriverDetails.ContactNumber, DriverDetails.EmailAddress, DriverDetails.DriverIdentification, DriverDetails.LicenseNumber, "Active", datetime)
		if err != nil {
			log.Printf("Error %s when isnerting row into passenger table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d Driver Created", rows)

		AccountUsername := DriverDetails.Firstname + " " + DriverDetails.Lastname

		log.Printf(AccountUsername)

		CreateAccount(AccountUsername, DriverDetails.Password, "Driver", "active")

		return c.String(http.StatusAccepted, "Driver Created!")
	}

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
	LoginDetails := LoginInformation{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	LoginDB := OpenLoginsDB()

	accountExists := CheckAccount(&LoginDB, LoginDetails.Username, LoginDetails.Password)
	log.Println(accountExists)

	username := LoginDetails.Username

	firstname := strings.Fields(username)
	log.Printf(username)

	if accountExists[0] == "true" {
		log.Printf(accountExists[0])

		http.Redirect(c.Response(), c.Request(), "http://localhost:9000/homepage/"+firstname[0]+"/"+accountExists[1], http.StatusSeeOther)
		return c.String(http.StatusOK, accountExists[1])
	} else {
		http.Redirect(c.Response(), c.Request(), "http://localhost:9000/login", http.StatusSeeOther)
		return c.String(http.StatusNotAcceptable, accountExists[0])
	}

}

func getPassengerName(PassengerID string) string {
	PassengersDB := OpenPassengersDB()
	var FirstName string
	var LastName string

	Query := "SELECT FirstName, LastName FROM GoRide_Passengers.Passenger WHERE ID = '" + PassengerID + "'"
	fmt.Println(Query)
	results := PassengersDB.QueryRow(Query)

	switch err := results.Scan(&FirstName, &LastName); err {
	case sql.ErrNoRows:
		PassengersDB.Close()
		return "Empty"
	case nil:
		PassengersDB.Close()
		log.Printf("Row Get!")
		return FirstName + " " + LastName
	default:
		panic(err)
	}
}

func checkrequests(DriverName string) string {
	TripsDB := OpenTripsDB()
	PassengerRequest := PassengerTrip{}

	driverID := getDriverID(DriverName)

	Query := "SELECT PassengerID, PickUp, DropOff FROM GoRide_Trips.trip WHERE DriverID = '" + driverID + "'  AND TripStatus = 'Pending'"
	fmt.Println(Query)
	results := TripsDB.QueryRow(Query)

	switch err := results.Scan(&PassengerRequest.Name, &PassengerRequest.PickupLocation, &PassengerRequest.DropoffLocation); err {

	case sql.ErrNoRows:
		TripsDB.Close()
		return "Empty"
	case nil:
		TripsDB.Close()
		log.Printf("Row Get!")
		passengerName := getPassengerName(PassengerRequest.Name)

		log.Printf("PassengerName is %s", passengerName)
		fullstring := passengerName + "," + PassengerRequest.PickupLocation + "," + PassengerRequest.DropoffLocation

		return fullstring
	default:
		panic(err)
	}
}

func checktriprequests(c echo.Context) error {
	log.Printf("Checkk Trip Requests Accessed")
	Drivername := DriverName{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&Drivername)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		//If error is nil, check if driver has any trip requests
		reply := checkrequests(Drivername.Username)
		log.Printf("Check Trip Requests reply is: %s", reply)
		return c.String(http.StatusOK, reply)
	}
}

func assignDriverID() string {
	DriversDB := OpenDriversDB()
	var ID string
	Driver := Driver{}

	Query := "SELECT ID, FirstName, LastName, ContactNumber, EmailAddress, DriverIdentification, LicenseNumber FROM GoRide_Drivers.Driver d WHERE d.ID NOT IN (SELECT DriverID FROM GoRide_Trips.Trip) LIMIT 1"
	fmt.Println(Query)
	results := DriversDB.QueryRow(Query)

	switch err := results.Scan(&ID, &Driver.Firstname, &Driver.Lastname, &Driver.ContactNumber, &Driver.EmailAddress, &Driver.DriverIdentification, &Driver.LicenseNumber); err {
	case sql.ErrNoRows:
		DriversDB.Close()
		return "Empty"
	case nil:
		DriversDB.Close()
		log.Printf("Row Get!")
		log.Printf(ID)
		return ID
	default:
		panic(err)
	}
}

func getDriverID(DriverName string) string {
	DriversDb := OpenDriversDB()
	var ID string

	Query := "SELECT ID FROM GoRide_Drivers.Driver WHERE FirstName = '" + DriverName + "'"
	fmt.Println(Query)
	results := DriversDb.QueryRow(Query)

	switch err := results.Scan(&ID); err {
	case sql.ErrNoRows:
		DriversDb.Close()
		return "Empty"
	case nil:
		DriversDb.Close()
		log.Printf("Row Get!")
		log.Printf(ID)
		return ID
	default:
		panic(err)
	}
}

func getPassengerID(PassengerName string) string {
	PassengersDB := OpenPassengersDB()
	var ID string

	Query := "SELECT ID FROM GoRide_Passengers.Passenger WHERE FirstName = '" + PassengerName + "'"
	fmt.Println(Query)
	results := PassengersDB.QueryRow(Query)

	switch err := results.Scan(&ID); err {
	case sql.ErrNoRows:
		PassengersDB.Close()
		return "Empty"
	case nil:
		PassengersDB.Close()
		log.Printf("Row Get!")
		log.Printf(ID)
		return ID
	default:
		panic(err)
	}
}

func createtrip(c echo.Context) error {
	PassengerRequest := PassengerTrip{}
	driverID := assignDriverID()

	name := c.Param("name")
	log.Printf("Name from param is %s", name)

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&PassengerRequest)

	log.Printf("Passenger Start: %s", PassengerRequest.PickupLocation)
	log.Printf("Passenger End: %s", PassengerRequest.DropoffLocation)
	log.Printf("Passenger Name: %s", PassengerRequest.Name)

	passengerID := getPassengerID(name)
	log.Printf("Passenger ID is %s", passengerID)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		//If error is nil, get a driver and create trip

		log.Printf("Driver ID is %s", driverID)
		TripsDB := OpenTripsDB()
		Query := "INSERT INTO Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus) VALUES (?, ?, ?, ?, ?)"

		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		stmt, err := TripsDB.PrepareContext(ctx, Query)

		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx, passengerID, driverID, PassengerRequest.PickupLocation, PassengerRequest.DropoffLocation, "Pending")
		if err != nil {
			log.Printf("Error %s when isnerting row into passenger table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d Trip Created", rows)

		return c.String(http.StatusAccepted, "Trip Created!")

	}

	return c.String(http.StatusOK, driverID)
}

func tripstatus(c echo.Context) error {
	status := c.Param("status")
	drivername := c.Param("drivername")
	driverID := getDriverID(drivername)
	TripsDB := OpenTripsDB()

	if status == "start" {
		statusToUpdate := "Started"
		log.Printf("Status to update is: %s", statusToUpdate)

		Query := "UPDATE GoRide_Trips.Trip SET TripStatus = 'Started', TripStart = CurDate() WHERE DriverID = '" + driverID + "' AND TripStatus = 'Pending'"
		log.Printf("Query is: %s", Query)
		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		stmt, err := TripsDB.PrepareContext(ctx, Query)

		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx)
		if err != nil {
			log.Printf("Error %s when updating trips table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d Trips Updated", rows)
		return c.String(http.StatusOK, "Status updated")

	} else if status == "end" {
		statusToUpdate := "Ended"
		log.Printf("Status to update is: %s", statusToUpdate)

		Query := "UPDATE GoRide_Trips.Trip SET TripStatus = 'Ended', TripEnd = CurDate() WHERE DriverID = '" + driverID + "' AND TripStatus = 'Started'"
		log.Printf("Query is: %s", Query)
		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		stmt, err := TripsDB.PrepareContext(ctx, Query)

		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx)
		if err != nil {
			log.Printf("Error %s when updating trips table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d Trips Updated", rows)

		return c.String(http.StatusOK, "Status updated")

		url := "http://localhost:9000/homepage/" + drivername + "/Driver"

		c.Redirect(http.StatusFound, url)
	}

	return c.String(http.StatusNotAcceptable, "Status not received")

}

func getTrips(username string, accounttype string) []Trip {
	TripsDb := OpenTripsDB()
	var TripList []Trip

	Query := ""

	if accounttype == "Passenger" {
		ID := getPassengerID(username)
		Query = "SELECT * FROM GoRide_Trips.Trip WHERE PassengerID = '" + ID + "' " + "ORDER BY TripStart DESC"
	} else if accounttype == "Driver" {
		ID := getDriverID(username)
		Query = "SELECT * FROM GoRide_Trips.Trip WHERE DriverID = '" + ID + "' " + "ORDER BY TripStart DESC"
	}

	rows, err := TripsDb.Query(Query)

	if err != nil {
		log.Printf(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var TripTemp Trip
		err := rows.Scan(&TripTemp.ID, &TripTemp.PassengerID, &TripTemp.DriverID, &TripTemp.PickUp, &TripTemp.DropOff, &TripTemp.TripStatus, &TripTemp.TripStart, &TripTemp.TripEnd)

		if err != nil {
			log.Printf(err.Error())
		}
		TripList = append(TripList, TripTemp)
	}
	return TripList
}

func viewtrips(c echo.Context) error {

	log.Printf("View trips accessed")

	Account := Account{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&Account)
	log.Printf("Name retrieved: %s", Account.Username)
	log.Printf("Accoutntype retrieved: %s", Account.AccountType)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		TripRecords := getTrips(Account.Username, Account.AccountType)
		var trips string

		for i := 0; i < len(TripRecords); i++ {
			temp := TripRecords[i].ID + "," + TripRecords[i].PassengerID + "," + TripRecords[i].DriverID + "," + TripRecords[i].PickUp + "," + TripRecords[i].DropOff + "," + TripRecords[i].TripStatus + "," + TripRecords[i].TripStart + "," + TripRecords[i].TripEnd
			trips = trips + "/" + temp
			log.Printf("Trips is: %s", trips)
		}

		log.Printf(trips)

		return c.String(http.StatusOK, trips)

		//jsonresults, _ := json.Marshal(TripRecords)
		//log.Println(string(jsonresults))

		//return c.JSON(http.StatusOK, jsonresults)
	}

}

func ServeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GoRide/1.0")

		return next(c)
	}
}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Use custom server header dispalying applciation version
	e.Use(ServeHeader)

	//Group API version one routes together
	g := e.Group("/api/V1")

	//Routes
	//Send Get Request in form-data with keys 'Username' and 'Password'
	g.POST("/checkuser", Checkuser)
	g.POST("/database/createpassenger", InsertPassenger)
	g.POST("/database/createdriver", InsertDriver)
	g.POST("/database/createtrip/:name", createtrip)
	g.POST("/checktriprequests", checktriprequests)
	g.POST("/database/tripstatus/:status/:drivername", tripstatus)
	g.POST("/database/viewtrips/:name/:accountype", viewtrips)

	go func() {
		if err := e.Start(":8001"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	//Gracefully shutdown the server if an error happens
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
