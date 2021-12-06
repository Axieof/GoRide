package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginInformation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func requestCheck(username string, password string) string {
	postBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	responsebody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8001/api/checkuser", "application/json", responsebody)

	if err != nil {
		log.Fatalf("An error occured %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}

//e.GET("/login/api", GetLogin)
func GetLogin(c echo.Context) error {
	var reply string

	LoginDetails := LoginInformation{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&LoginDetails)

	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else {
		//If error is nil, check if username and  password exists in database.
		reply = requestCheck(LoginDetails.Username, LoginDetails.Password)
		log.Printf(reply)
	}

	return c.String(http.StatusOK, reply)
}

func main() {
	//Create Echo HTTP Server
	e := echo.New()

	//Routes
	//Listen to POST Request with keys 'Username' and 'Password'
	e.POST("/api/login", GetLogin)

	e.Logger.Fatal(e.Start(":8000"))
}
