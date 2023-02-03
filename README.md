# Synapsis Test BE

This repository is a place to ...

## Requirements

to run this application locally

```
- golang v1.19 or higher
- mariadb or mysql database
```

to run with docker

```
- docker server
- docker compose
```

## Setup

Please run the command bellow

- Clone the repository
- Replace config `.env` with yours or if you using unix, you can setup environment with copy this command on your terminal

```
# EXPORT DB_USER=your_db
# EXPORT DB_PASS=your_db_password
# EXPORT DB_NET=tcp
# EXPORT DB_HOST=your_db_host
# EXPORT DB_PORT=your_db_port
# EXPORT DB_NAME=your_db_name
# EXPORT JWT_KEY=generate_your_secret_key
```

## How to run 

If you want to run this application locally

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

## Api Specs Documentation

All this API documentation endpoints is in the following link

- https://swagger.io

## Folder structure