# GORM (Object Oriented Mapping)

GORM is a library that makes it easy for us to connect to the database so we don't need to use native queries where we initialize the query first, but by using this we can simplify operations and communication into the database. Some of them are needed. For more details, check [here](https://gorm.io/).

## Requirement

1. You have PostgreSQL in your computer or you can use the cloud instead.
2. For this project, we will need some packages. The dependencies that you need:

```go
  go get gorm.io/gorm
  go get gorm.io/driver/postgres
  go get github.com/joho/godotenv
```

- GORM is package ORM in Golang
- Postgres is package that required to run GORM with Postgres
- Godotenv is package to read the `.env` file

## How to use

1. You can clone this project.
2. You can run:
   ```go
    go mod tidy
   ```
   to install the dependencies.
3. And finally, you can run with
   ```go
    go run main.go
   ```
4. You can just run the handler that do you want in `main.go`.
