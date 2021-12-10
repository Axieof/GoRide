package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Passenger struct {
	Firstname     string `json: firstname`
	Lastname      string `json: lastname`
	Password      string `json: password`
	ContactNumber string `json: contactnumber`
	EmailAddress  string `json: emailaddress`
}

func CreatePassenger(c echo.Context) error {

	log.Printf("Details posted to passenger")

	PassengerDetails := Passenger{
		Password:      c.FormValue("password"),
		Firstname:     c.FormValue("firstname"),
		Lastname:      c.FormValue("lastname"),
		ContactNumber: c.FormValue("mobilenumber"),
		EmailAddress:  c.FormValue("emailaddress"),
	}

	postBody, _ := json.Marshal(map[string]string{
		"firstname":     PassengerDetails.Firstname,
		"lastname":      PassengerDetails.Lastname,
		"password":      PassengerDetails.Password,
		"contactnumber": PassengerDetails.ContactNumber,
		"emailaddress":  PassengerDetails.EmailAddress,
	})

	responsebody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8001/api/V1/database/createpassenger", "application/json", responsebody)

	if err != nil {
		log.Fatalf("An error occured %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return c.String(http.StatusOK, sb)
}

func ServeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GoRide/1.0")

		return next(c)
	}
}

func GetPassenger() string {
	return "Lol"
}

func DeletePassenger() string {
	return "Lol"
}

func UpdatePassenger() string {
	return "Lol"
}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Use custom server header dispalying applciation version
	e.Use(ServeHeader)

	//Group API version one routes together
	g := e.Group("/api/V1")

	//Routes
	//Listen to POST Request with keys 'Username' and 'Password'
	//g.GET("/passenger", GetPassenger)
	g.POST("/passenger", CreatePassenger)
	//g.DELETE("/passenger", DeletePassenger)
	//g.PUT("/passenger", UpdatePassenger)

	go func() {
		if err := e.Start(":8002"); err != nil && err != http.ErrServerClosed {
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
