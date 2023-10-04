package main

import (
	"covid-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func updateData() {
	utils.DownloadFile("WHO-COVID-19-global-data.csv", "https://covid19.who.int/WHO-COVID-19-global-data.csv")
}

func main() {
	var countriesReports map[string]utils.Country

	// Update the data every 24 hours
	go func() {
		for {
			updateData()
			countriesReports, _ = utils.ParseCsvFile("WHO-COVID-19-global-data.csv")
			time.Sleep(24 * time.Hour)
		}
	}()

	// Try to read the data every 5 seconds
	func() {
		for {
			if _countriesReports, err := utils.ParseCsvFile("WHO-COVID-19-global-data.csv"); err == nil {
				countriesReports = _countriesReports
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()

	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/countries", func(c echo.Context) error {
		countriesNames := map[string]string{}
		for countryCode, countryData := range countriesReports {
			if countryCode == " " {
				continue
			}
			countriesNames[countryCode] = countryData.CountryName
		}
		return c.JSON(200, countriesNames)
	})

	e.GET("/countries/:countryCode", func(c echo.Context) error {
		countryCode := c.Param("countryCode")
		countryReports := countriesReports[countryCode]
		return c.JSON(200, countryReports)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
