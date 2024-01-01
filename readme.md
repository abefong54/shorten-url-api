# Project Title

URL Shortner

## Description

A simple URL shortning API with functionality similar to TinyURL.
We use Redis as a local db for v1.

## Getting Started

### Dependencies

- GoLang - go version go1.19.3
- Docker
- github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
- github.com/go-redis/redis/v8 v8.11.5
- github.com/gofiber/fiber/v2 v2.51.0

### Installing

- Clone repo to your local machine.
- CD into `./shorten-url` directory and run the following command to install dependencies.

```
go mod tidy
```

- Any modifications needed to be made to files/folders

### Executing program

- from the `./shorten-url` directory level run docker containers using the following command

```
docker-compose up -d
```

### Using the API

- After the container is running with the db and the api images,
  use a service like Postman to make a POST call to `http://127.0.0.1:8080/api/v1`

- Provide a JSON body like the one below. custom_short and expiry are optional.

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
