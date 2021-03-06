# GoRide Ride Sharing Platform

## Project Description

This project is an assignment of my Emerging Trends in IT(ETI) Module. The assignment was to design a Ride-Sharing Platform with the use of Microservice Architecture, written in Go with the use of MYSQL as a database and HTML/CSS for the FrontEnd

---
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

---
## Design & Architecture

![GoRide Architecture](https://github.com/Axieof/GoRide/blob/master/setup/img/GoRide_Architecture_Diagram.PNG)

GoRide consists of 6 services, that all communciate with each other, or through each other, to perform tasks. The main service is the frontend service, which renders html for the user as an itnerface to perform tasks. The frontend then calls the corresponding services based on the tasks requested, such as Authentication service for logging in, Trip service for creating trips and viewing trips. The frontend also communicates with the Passenger and Driver service which communciates with the database service for certain tasks. GoRide was designed with each service built to perform specifics tasks related to a certain category, such as the category they are named after. 

### Authentication Service
- POST (http://localhost:8000/api/V1/login)
```
The Authentication Service listens to any posts request to the login route, whenevr a user tries to login. 
The Frontend service then posts the details the user has posted in the login form to the route, 
which then takes in teh values and checks again the database if the login exists. More was planend 
for the Authentication service, such as JWT authentication, but was not able to be implemented by the deadline.
```

---
### Frontend Service
- GET (http://localhost:9000/)
```
This route is when the user attempts the access the GoRide service, by reaching the default landing page. 
The Frontend service then renders a html offering options to login, register as a passenger or driver.
```
- GET (http://localhost:9000/login)
```
This route is when the user selects the login option, to which they are presented with a login screen
where they can enter their login details in a form. The details are then sent to the authentication 
service where it is checked against the database through the database service to ensure teh login details are correct.
```
- GET (http://localhost:9000/register)
```
This route is for when the user attempts to register an account. Depending on the option they selected, 
Register as a passenger or Register as a Driver, query parameters in the form of ?account=accounttype 
are sent to the route, to render corresponding registration page as both passenger and driver 
have different registration details that they need to fill in.
```
- POST (http://localhost:9000/register)
```
This route is for when the user clicks the register button, which posts the data to the same 
route to be processed and sent to the corresponding Passenger,Driver service to create an 
account, which then communicates with the database service to create the account as well 
as passenger/driver record into the database.
```
- GET (http://localhost:9000/homepage/:name/:accounttype)
```
When a user logs in into a valid account, they are directed to this route, where a html page 
is rendered with their name, as well as options catering to their account type, such as Check Trips 
for driver, Book Trip for passenger, and View Trips for both
```
- GET (http://localhost:9000/booktrip/:name)
```
This route is for when a passenger attempts to book a trip, which then renders a html form for 
them to input their pick up and drop off postal codes.
```
- GET (http://localhost:9000/viewtrips/:name/:accounttype)
```
This route is for when the user, regardless of passenger or driver, wants to view their trip history
in reverse chronological order. The Frontend service then communicates with the trip service which 
then communicates with the database service to retrieve the records and render them onto the screen.
```
- GET (http://localhost:9000/checktrips/:username)
```
This route is for when the driver wants to check if any trip requests have been assigned to them,
where they can then start the trip if they have picked up the passenger, and end trip when they 
have dropped off the passenger.
```

---
### Database Service
- POST (http://localhost:8001/api/V1/checkuser)
```
This route is for when the authentication service posts the suer login information, where it is then
sent to this rotue to be checked against the database.
```
- POST (http://localhost:8001/api/V1/database/createpassenger)
```
This route is for when the user creates a new passenger account, where the Frontend service then 
posts their details to the passenger service which then posts it to this rotue to create a pasenger
```
- POST (http://localhost:8001/api/V1/databse/createdriver)
```
This route is for when the user creates a new driver account, where the Frontend service then 
posts their details to the driver service which then posts it to this rotue to create a driver
```
- POST (http://localhost:8001/api/V1/createtrip/:name)
```
This route is for when the passenger user books a new trip, which sends their inforamtion to 
the trip service, which then sends it to this route to create the trip record. A free driver 
is assigned to the trip at this point and sent back to the driver service so that the driver 
can choose to start the trip or end it.
```
- POST (http://localhost:8001/api/V1/checktriprequests)
```
This route is for when the driver user wants to check if they have been assigned into any trips, 
in which the database service checks against the database and sends the relevant information 
to the Frontend service for the user to view it.
```
- POST (http://localhost:8001/api/V1/database/tripstatus/:status/:drivername)
```
This route is for when the driver user starts or ends the trip, to which the information is 
sent to this route from the Frontend to the Trip and then udpates the database depending on 
the case, such as changing status of trip to "Started" if the driver starts it or changing 
status of trip to "Ended" if driver ends it.
```
- POST (http://localhost:8001/api/V1/database/viewtrips/:name/:accounttype)
```
This route is for when the user wants to view their trip, where their information is sent 
from the Frontend to this route, which then returns the trips of the user for it to be rendered 
into a html page.
```

---
### Trip Service
- POST (http://localhost:8004/api/V1/createtrip/:name)
```
This route is for when the user wants to create a trip, where information is passed from the
Frontend service to here, where the information is processed then sent to the database service 
for a record to be created
```

---
### Driver Service
- POST (http://localhost:8003/api/V1/driver)
```
This route is for when the user creates a driver account, where the the details are sent here 
to be processed and sent to the database sercice for a record to be created
```

---
### Passenger Service
- POST (http://localhost:8002/api/V1/passenger)
```
This route is for when the user creates a passenger account, where the the details are sent here 
to be processed and sent to the database sercice for a record to be created
```

---
