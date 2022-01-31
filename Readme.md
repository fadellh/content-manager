# Content Manager

## Run the test

```sh
go test -v -cover ./...
```

## Run Project with Docker

```sh
docker-compose up --build
```

## Run App Manually

1. Install gcc for sqlite database based on environment

Read more on this link https://github.com/mattn/go-sqlite3#google-cloud-platform

Example for Ubuntu

```sh
sudo apt-get install build-essential
```

2. Run app

```sh
go run main.go
```
