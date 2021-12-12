# GoRide Ride Sharing Platform

## Project Description

This project is an assignment of my Emergin Trends in IT(ETI) Module. The assignment was to design a Ride-Sharing Platform with the use of Microservice Architecture, written in Go with the use of MYSQL as a database and HTML/CSS for the FrontEnd

## Setup Instructions

Each Service requires the Go [Echo Framework](https://echo.labstack.com/guide/) package, with the database service needing the mysql sql driver package

FrontEnd Service Setup Commands

```
cd src/Frontend
go get "github.com/labstack/echo/v4"
go run main.go
```

Authentication Service Setup Commands

```
cd src/Authentication
go get "github.com/labstack/echo/v4"
go run main.go
```

Passenger Service Setup Commands

```
cd src/Passenger
go get "github.com/labstack/echo/v4"
go run main.go
```

Driver Service Setup Commands

```
cd src/Driver
go get "github.com/labstack/echo/v4"
go run main.go
```

Trip Service Setup Commands

```
cd src/Trip
go get "github.com/labstack/echo/v4"
go run main.go
```

Database Service Setup Commands

```
cd src/Database
go get "github.com/labstack/echo/v4"
go get "github.com/go-sql-driver/mysql"
go run main.go
```
Database setup scripts can be found in setup/scripts/GoRide_DB_Initialization.sql
Open the file in your MySQL Workbench and run it

## Design & Architecture
