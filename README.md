# Tech stack:
- Go
- Gorilla Mux
```bash
$ go get -u github.com/gorilla/mux
```
- GORM:
```bash
$ go get -u gorm.io/gorm
$ go get gorm.io/driver/postgres # postgres driver
```
- Postgres
- PGAdmin
- docker
## Go mod init:
- go mod init github.com/luciormoraes/go-rest-api
- go mod init go-rest-api

## Docker: Postgres
- to start Postgres:
```bash
$ docker-compose up
```
- to access docker:
```bash
$ docker exec -it pg_goapi bash
$ hostname -i
```
or
```bash
$ docker inspect id_container | grep IPAddres
```

# Day 17/100: Developing a Rest API
- Json, Request, Response and Go. Get returns Json

# Day 18/100: Developing a Rest API
- adding gorilla mux
- creating routes for our new Endpoints: GET Methods
- change celebrities to Personalities
- running postgres in docker

# Day 20/100: using GORM
- GORM installagion
- connection to DB / change controllers to show from DB
- new endpoints to Create, Delete, or Edit a personality
