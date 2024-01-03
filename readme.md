# Project Title

URL Shortner

## Description

A simple URL shortning API with functionality similar to TinyURL.
We use Redis as a lightweight db for v1.

We deploy using (Railway) [https://docs.railway.app/overview/about-railway] for infrastructure automation. Railway is a deployment platform designed to streamline the software development life-cycle, starting with instant deployments and effortless scale.

We use (apianalytics)[https://github.com/tom-draper/api-analytics], A free lightweight API analytics solution, complete with a dashboard, to track requests to the API.

API Analytics Dashboard (here) [https://www.apianalytics.dev/dashboard/ee19cf67f5a44a2c8218b07075eb472f]
Note - you will need the API key associated to our API to access the dashboard!

## Getting Started

### Dependencies

- GoLang - go version go1.19.3
- Docker (for local development/testing)

### Installing

- Clone repo to your local machine.
- This application requires you to have Docker installed locally.
- The Dockerfile included will build the main go build for the application to run and a local Redis container for DB functionality.

### Executing program

- from the top directory level run docker containers using the following command:

```
docker-compose up -d
```

### Using the API

- After the container is running with the db and the api images,
  use a service like Postman to make a POST call to `http://127.0.0.1:8080/api/v1`

- Provide a JSON body like the one below. Note: custom_short and expiry are optional.
  If custom short is not provided, we will

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

- If the `custom_short`` value has already been used before, you will see an error message like the one below:

```
{
    "error": "Custom short URL already in use."
}
```

## Authors

Contributors names and contact info

ex. Abraham Fong (abefong54@gmail.com)

## Version History

- 0.1
  - Initial Release
