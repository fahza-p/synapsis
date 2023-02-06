# Synapsis Test BE

This repository for doing test from PT. Synapsis Sinergi Digital

## Requirements

to run this application locally

```
- golang v1.19 or higher
- mariadb or mysql database
```

to run with docker

```
- docker
- docker compose
```

## Setup

Please run the command bellow

- Clone the repository
- Replace config `.env` with yours or if you using unix, you can setup environment with copy this command on your terminal

```
# EXPORT DB_USER=your_db
# EXPORT DB_PASS=your_db_password
# EXPORT DB_ADDR=your_host:your_port
# EXPORT DB_NAME=your_db_name
# EXPORT JWT_KEY=generate_your_secret_key
# EXPORT PORT=3000
```

- Create your database and copy value from `script/migration/migrate.sql`. (this step can be skipped when you using docker)

## How to run 

#### If you want to run this application locally

- Load all golang module

```
go mod tidy
```

- Run application

```
go run main.go
```

- Run application with hot reload using nodemon (optional)

```
nodemon --exec go run main.go --signal SIGTERM
```

#### If you using docker

```
docker-compose up --build
```

## Api Specs Documentation

All this API documentation endpoints is in the following link

- https://swagger.io

## Folder structure