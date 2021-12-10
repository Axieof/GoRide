package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

type Driver struct {
	Password             string
	FirstName            string
	LastName             string
	MobileNumber         string
	EmailAddress         string
	IdentificationNumber string
	CarLicenseNumber     string
}

type Passenger struct {
	Password     string
	FirstName    string
	LastName     string
	MobileNumber string
	EmailAddress string
}

var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("app"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "app")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func login(c echo.Context) error {
	log.Printf("Login accessed")
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func homepage(c echo.Context) error {
	log.Printf("Login accessed")
	return c.Render(http.StatusOK, "homepage.html", map[string]interface{}{})
}

func register(c echo.Context) error {
	account := c.QueryParam("account")

	if account == "passenger" {
		log.Printf("Passenger account registration accessed")

		return c.Render(http.StatusOK, "registerPassenger.html", map[string]interface{}{})

	} else if account == "driver" {
		log.Printf("Driver account registration accessed")

		return c.Render(http.StatusOK, "registerDriver.html", map[string]interface{}{})

	} else {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
	}

}

func getFormValue(c echo.Context) error {
	account := c.QueryParam("account")

	log.Printf("Account details posted")

	if account == "driver" {
		driverdetails := Driver{
			Password:             c.FormValue("password"),
			FirstName:            c.FormValue("firstname"),
			LastName:             c.FormValue("lastname"),
			MobileNumber:         c.FormValue("mobilenumber"),
			EmailAddress:         c.FormValue("emailaddress"),
			IdentificationNumber: c.FormValue("idnumber"),
			CarLicenseNumber:     c.FormValue("carlicensenumber"),
		}

		log.Printf("Details are %s", driverdetails.FirstName)

		postBody, _ := json.Marshal(map[string]string{
			"firstname":            driverdetails.FirstName,
			"lastname":             driverdetails.LastName,
			"password":             driverdetails.Password,
			"contactnumber":        driverdetails.MobileNumber,
			"emailaddress":         driverdetails.EmailAddress,
			"driveridentification": driverdetails.IdentificationNumber,
			"licensenumber":        driverdetails.CarLicenseNumber,
		})

		driver_json, _ := json.Marshal(postBody)

		driver_data := bytes.NewBuffer(driver_json)

		response, err := http.Post("http://localhost:8003/api/V1/driver", "application/json", driver_data)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {

			reply, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			fmt.Printf("Reply is: %s", reply)

			c.Redirect(http.StatusFound, "http://localhost:9000/homepage")
		}

		return c.String(http.StatusOK, "Successful")

	} else if account == "passenger" {
		passengerdetails := Passenger{
			Password:     c.FormValue("password"),
			FirstName:    c.FormValue("firstname"),
			LastName:     c.FormValue("lastname"),
			MobileNumber: c.FormValue("mobilenumber"),
			EmailAddress: c.FormValue("emailaddress"),
		}

		postBody, _ := json.Marshal(map[string]string{
			"firstname":     passengerdetails.FirstName,
			"lastname":      passengerdetails.LastName,
			"password":      passengerdetails.Password,
			"contactnumber": passengerdetails.MobileNumber,
			"emailaddress":  passengerdetails.EmailAddress,
		})

		passenger_json, _ := json.Marshal(postBody)

		passenger_data := bytes.NewBuffer(passenger_json)

		response, err := http.Post("http://localhost:8003/api/V1/passenger", "application/json", passenger_data)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {

			reply, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			fmt.Printf("Reply is: %s", reply)

			c.Redirect(http.StatusFound, "http://localhost:9000/homepage")
		}

		return c.String(http.StatusOK, "Successful")
	} else {

		return c.String(http.StatusUnprocessableEntity, "Not Successful")
	}
}

func main() {

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}

	e.Renderer = renderer

	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/", fs)

	e.GET("/", index)
	e.GET("/login", login)
	e.GET("/register", register)
	e.POST("/register", getFormValue)

	e.GET("/homepage", homepage)

	log.Printf("Frontend Service started")

	e.Logger.Fatal(e.Start(":9000"))

}
