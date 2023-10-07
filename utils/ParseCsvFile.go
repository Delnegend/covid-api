package utils

import (
	"encoding/csv"
	"os"
)

type DailyReport struct {
	Date             string
	NewCases         int
	NewDeaths        int
	CumulativeCases  int
	CumulativeDeaths int
}

type Country struct {
	CountryName string
	WHORegion   string
	Reports     []DailyReport
}

func ParseCsvFile(filePath string) (map[string]Country, error) {
	// Open the CSV file
	file, err := os.Open("WHO-COVID-19-global-data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader reading from the opened file
	reader := csv.NewReader(file)

	// Read in all of the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create an array of countries
	reports := map[string]Country{}

	worldwide := Country{
		CountryName: "Worldwide",
		WHORegion:   "WW",
		Reports:     []DailyReport{},
	}

	// Iterate through the records, starting at row 2
	for i, row := range records {
		if i == 0 {
			continue
		}

		// Parse the data
		date := row[0]
		newCases := parseInt(row[4])
		newDeaths := parseInt(row[6])
		cumulativeCases := parseInt(row[5])
		cumulativeDeaths := parseInt(row[7])
		countryCode := row[1]
		countryName := row[2]
		whoRegion := row[3]
		report := DailyReport{
			Date:             date,
			NewCases:         newCases,
			NewDeaths:        newDeaths,
			CumulativeCases:  cumulativeCases,
			CumulativeDeaths: cumulativeDeaths,
		}

		// Check if we already have a country for this country code
		country, ok := reports[countryCode]
		if !ok {
			country = Country{
				CountryName: countryName,
				WHORegion:   whoRegion,
			}
		}

		// Check if we already have a daily report for this date for worldwide
		dateExists := false
		for i, wwReport := range worldwide.Reports {
			if wwReport.Date == date {
				worldwide.Reports[i].NewCases += newCases
				worldwide.Reports[i].NewDeaths += newDeaths
				worldwide.Reports[i].CumulativeCases += cumulativeCases
				worldwide.Reports[i].CumulativeDeaths += cumulativeDeaths
				dateExists = true
				break
			}
		}
		if !dateExists {
			worldwide.Reports = append(worldwide.Reports, report)
		}

		// Append the daily report to the country's report list
		country.Reports = append(country.Reports, report)

		// Store the country in the countries map
		reports[countryCode] = country
	}

	reports["WW"] = worldwide

	return reports, nil
}
