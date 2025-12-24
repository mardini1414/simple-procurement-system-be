# Simple procurement system be

Simple procurement system be is a web api for Simple procurement system fe

---

## Tech Stack

- Golang
- Fiber
- Postgres

---

## Prerequisites

Make sure the following tools are installed on your machine:

- Golang 1.24
- postgres database

## How to run

- copy file .env.example and rename to .env
- change the contents of the configuration file such as the database and others
- run this command for install dependencies

```bash
go mod tidy
```

- run this command for database migration

```bash
go run ./cmd/migrate/migrate.go
```

- run this command for insert user

```bash
go run ./cmd/seeder/seeder.go
```

- default username is `admin` and password `admin123`
- run

```bash
go run ./cmd/main.go
```

- or if you using air run

```bash
air
```

- import postman collection from /doc folder to test
