# Covid-api
## Run using docker
### Compose
- It's the docker-compose.yml file up there. If you don't want to build the image, remove `build` part from the file and replace `image: covid-api` with `image: ghcr.io/Liminova/covid-api/covid-api:latest`

### Run
- `docker run --name covid-api -p 8080:8080 --restart unless-stopped covid-api`

## Run using docker-compose, build from source
- Clone the repo
- `docker-compose up --build`

## API
- `/countries` - Get data for all countries
```jsonc
{
    "VN": "Vietnam",
    "US": "United States",
    "WW": "Worldwide"
    // ...
}
```

- `/countries/{countryCode}` - Get data for a specific country, country code is the key from `/countries`
```jsonc
{
    "CountryName": "Viet Nam",
    "WHORegion": "WPRO",
    "Reports": [
        {
            "Date": "2020-01-03",
            "NewCases": 0,
            "NewDeaths": 0,
            "CumulativeCases": 0,
            "CumulativeDeaths": 0
        },
        // ...
    ]
}
```