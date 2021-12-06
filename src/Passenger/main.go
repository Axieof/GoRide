package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ServeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GoRide/1.0")

		return next(c)
	}
}

func GetPassenger() {

}

func CreatePassenger() {

}

func DeletePassenger() {

}

func UpdatePassenger() {

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
	g.GET("/passenger", GetPassenger)
	g.POST("/passenger", CreatePassenger)
	g.DELETE("/passenger", DeletePassenger)
	g.PUT("/passenger", UpdatePassenger)

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
