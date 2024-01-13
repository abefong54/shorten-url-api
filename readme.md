# URL SHORTENER SERVICE

## Description

A simple URL shortning API with functionality similar to TinyURL.
We use Redis as a lightweight db for v1.

We deploy using Railway, [https://docs.railway.app/overview/about-railway] for infrastructure automation. Railway is a deployment platform designed to streamline the software development life-cycle, starting with instant deployments and effortless scale.

We use[https://github.com/tom-draper/api-analytics], A free lightweight API analytics solution, complete with a dashboard, to track requests to the API.

API Analytics Dashboard: [https://www.apianalytics.dev/dashboard/efaefbb3cf2643159d5b8e63798cc4ce]
Note - you will need the API key associated to our API to access the dashboard!

## Getting Started

### Dependencies

- GoLang - go version go1.19.3
- Docker (for local development/testing)

### Installing

- Clone repo to your local machine.
- This application requires you to have Docker installed locally.
- The Dockerfile included will build the main go build for the application to run and a local Redis container for DB functionality.

### Executing program locally

- from the top directory level run docker containers using the following command:

```
docker-compose up -d
```

## Using the API

### PUBLIC API DETAILS

The service is available at the following endpoint.

- PUBLIC_URL: `https://shorten-url-api-production.up.railway.app`

### AVAILABLE ROUTES

#### 1. POST request to create a new short url.

- ENDPOINT: PUBLIC_URL + `/api/v1`

- Example Payload:

```
{
    "url":"https://medium.com/rupesh-tiwari/publishing-merged-code-coverage-report-of-nx-workspace-in-azure-ci-pipeline-70b44dbff1d9",
    "custom_short" : "my-custom-short-id",
    "expiry" : 1,
}
```

NOTES:

- expiry time value is saved in hours.
- custom_short and expiry time are optional.
- if custom short is not given, a random id will be generated.
- if expiry time is not given, a default time of 24 hours will be saved.

#### 2. GET request to resolve an existing, unexpired short url.

- ENDPOINT: PUBLIC_URL + `/my-custom-short-id`

- Example Payload: Notice that the payload should be provided in the URL (not via json) and should be a valid ID of a valid short url.

## LOCAL/DEVELOPMENT API

- After the container is running with the db and the api images,
  use a service like Postman to make a POST call to `http://127.0.0.1:8080/api/v1`

- Provide a JSON body like the one below. Note: custom_short and expiry are optional.
  If custom short is not provided, we will generate one with a random 6 digit id.

```
{
    "url":"https://medium.com/rupesh-tiwari/publishing-merged-code-coverage-report-of-nx-workspace-in-azure-ci-pipeline-70b44dbff1d9",
    "custom_short" : "mycoolerdomain",
    "expiry": 100
}
```

- A successfull request will result in something like this

```
{
    "URL": "https://medium.com/rupesh-tiwari/publishing-merged-code-coverage-report-of-nx-workspace-in-azure-ci-pipeline-70b44dbff1d9",
    "CustomShort": "localhost:8080/mycoolerdomain",
    "Expiry": 100,
    "XRateRemaining": 93,
    "XRateLimitReset": 25
}

```

- If the `custom_short` value has already been used before, you will see an error message like the one below:

```
{
    "error": "Custom short URL already in use."
}
```

- If the `short_url` value has expired, it will be removed from the database and you will receive the following message:

```
{
    "error": "ShortURL not found."
}
```

## Authors

Contributors names and contact info

ex. Abraham Fong (abefong54@gmail.com)

## Version History

- 0.1
  - Initial Release
