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

type Driver struct {
	Firstname            string `json: firstname`
	Lastname             string `json: lastname`
	ContactNumber        string `json: contactnumber`
	EmailAddress         string `json: emailaddress`
	DriverIdentification string `json: driveridentification`
	LicenseNumber        string `json: licensenumber`
}

func ServeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GoRide/1.0")

		return next(c)
	}
}

func CreateDriver(c echo.Context) error {
	DriverDetails := Driver{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&DriverDetails)

	log.Printf(DriverDetails.DriverIdentification)
	log.Printf(DriverDetails.LicenseNumber)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		postBody, _ := json.Marshal(map[string]string{
			"firstname":            DriverDetails.Firstname,
			"lastname":             DriverDetails.Lastname,
			"contactnumber":        DriverDetails.ContactNumber,
			"emailaddress":         DriverDetails.EmailAddress,
			"driveridentification": DriverDetails.DriverIdentification,
			"licensenumber":        DriverDetails.LicenseNumber,
		})

		responsebody := bytes.NewBuffer(postBody)

		resp, err := http.Post("http://localhost:8001/api/V1/database/createdriver", "application/json", responsebody)

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
	//g.GET("/driver", Getdriver)
	g.POST("/driver", CreateDriver)
	//g.DELETE("/driver", Deletedriver)
	//g.PUT("/driver", Updatedriver)

	go func() {
		if err := e.Start(":8003"); err != nil && err != http.ErrServerClosed {
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
