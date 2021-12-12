CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

CREATE DATABASE GoRide_Passengers;
USE GoRide_Passengers;

CREATE TABLE Passenger 
(
ID int NOT NULL auto_increment PRIMARY KEY,
FirstName varchar(30),
LastName varchar(30),
ContactNumber varchar(8),
EmailAddress varchar(50),
AccountStatus varchar(20),
AccountUpdated datetime
);

CREATE DATABASE GoRide_Drivers;
USE GoRide_Drivers;

CREATE TABLE Driver 
(
ID int NOT NULL auto_increment PRIMARY KEY,
FirstName varchar(30),
LastName varchar(30),
ContactNumber varchar(8),
EmailAddress varchar(50),
DriverIdentification varchar(10),
LicenseNumber varchar(20),
AccountStatus varchar(20),
AccountUpdated datetime
);

CREATE database GoRide_Trips;
USE GoRide_Trips;

CREATE TABLE Trip
(
ID int NOT NULL auto_increment PRIMARY KEY,
PassengerID varchar(5),
DriverID varchar(5),
PickUp varchar(6),
DropOff varchar(6),
TripStatus varchar(20),
TripStart datetime null,
TripEnd datetime null
);

CREATE DATABASE GoRide_Logins;
USE GoRide_Logins;

CREATE TABLE LoginInformation 
(
ID int NOT NULL auto_increment PRIMARY KEY,
Username varchar(30),
Password varchar(30),
AccountType varchar(9),
AccountStatus varchar(20),
AccountUpdated datetime
);

-- Trip Insertion
INSERT INTO GoRide_Trips.Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus, TripStart, TripEnd) VALUES(1, 1, "211009", "654321", "Ended", '2021-01-23 12:45:56', '2021-01-23 13:15:22');
INSERT INTO GoRide_Trips.Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus, TripStart, TripEnd) VALUES(1, 1, "654321", "211009", "Ended", '2021-01-23 17:30:41', '2021-01-23 16:02:32');
INSERT INTO GoRide_Trips.Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus, TripStart, TripEnd) VALUES(1, 1, "584796", "123456", "Ended", '2021-04-15 23:30:53', '2021-04-16 01:01:24');
INSERT INTO GoRide_Trips.Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus, TripStart, TripEnd) VALUES(1, 2, "739019", "100500", "Ended", '2021-07-18 14:30:23', '2021-07-18 14:55:00');
INSERT INTO GoRide_Trips.Trip(PassengerID, DriverID, PickUp, DropOff, TripStatus, TripStart, TripEnd) VALUES(1, 2, "100500", "739019", "Ended", '2021-07-18 19:07:07', '2021-07-18 19:33:04');

-- Login Insertion
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Pritheev Roshan", "12345", "Passenger", "Active", CURDATE());
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Danny Chan", "12345", "Driver", "Active", CURDATE());
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Caleb Goh", "12345", "Passenger", "Active", CURDATE());
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Dong Kiat", "12345", "Passenger", "Active", CURDATE());
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Kah Ho", "12345", "Driver", "Active", CURDATE());
INSERT INTO GoRide_Logins.LoginInformation(Username, Password, AccountType, AccountStatus, AccountUpdated) VALUES("Kenenth Teo", "12345", "Driver", "Active", CURDATE());

-- Driver Insertion
INSERT INTO GoRide_Drivers.Driver(FirstName, LastName, Password, ContactNumber, EmailAddress, DriverIdentification, LicenseNumber, AccountStatus, AccountUpdated) VALUES("Kenenth", "Teo", "12345678", "Kenneth@gmail.com", "D9400X", "SCL6193H", "Active", CURDATE());
INSERT INTO GoRide_Drivers.Driver(FirstName, LastName, Password, ContactNumber, EmailAddress, DriverIdentification, LicenseNumber, AccountStatus, AccountUpdated) VALUES("Danny", "Chan", "12345678", "Danny@gmail.com", "D1234X", "SGX7562F", "Active", CURDATE());
INSERT INTO GoRide_Drivers.Driver(FirstName, LastName, Password, ContactNumber, EmailAddress, DriverIdentification, LicenseNumber, AccountStatus, AccountUpdated) VALUES("Kah", "Ho", "12345678", "KahHo@gmail.com", "D1867X", "SUJ2396J", "Active", CURDATE());

-- Passenger Insertion
INSERT INTO GoRide_Passengers.Passenger(FirstName, LastName, Password, ContactNumber, EmailAddress,AccountStatus, AccountUpdated) VALUES("Pritheev", "Roshan", "12345678", "Pritheev@gmail.com", "Active", CURDATE());
INSERT INTO GoRide_Passengers.Passenger(FirstName, LastName, Password, ContactNumber, EmailAddress,AccountStatus, AccountUpdated) VALUES("Caleb", "Goh", "12345678", "Caleb@gmail.com", "Active", CURDATE());
INSERT INTO GoRide_Passengers.Passenger(FirstName, LastName, Password, ContactNumber, EmailAddress,AccountStatus, AccountUpdated) VALUES("Dong", "Kiat", "12345678", "DongKiat@gmail.com", "Active", CURDATE());

