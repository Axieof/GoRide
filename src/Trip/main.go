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

func ServeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GoRide/1.0")

		return next(c)
	}
}

func createtrip(c echo.Context) error {
	name := c.Param("name")

	Start := c.FormValue("pickuplocation")
	End := c.FormValue("dropofflocation")

	log.Printf("Name is %s", name)

	log.Printf("Start is %s", Start)
	log.Printf("End is %s", End)

	postBody, _ := json.Marshal(map[string]string{
		"pickuplocation":  Start,
		"dropofflocation": End,
		"passengerName":   name,
	})

	responsebody := bytes.NewBuffer(postBody)

	url := "http://localhost:8001/api/V1/database/createtrip/" + name

	resp, err := http.Post(url, "application/json", responsebody)

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

	//Get free driver
	//Create trip
	//

	return nil

}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Use custom server header dispalying applciation version
	e.Use(ServeHeader)

	//Group API version one routes together
	g := e.Group("/api/V1")

	//Routes
	g.POST("/createtrip/:name", createtrip)

	go func() {
		if err := e.Start(":8004"); err != nil && err != http.ErrServerClosed {
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
