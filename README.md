# golang-assignment

A simple app written in Go with a GET endpoint to get students records with provided query filter.

***

## Prerequisites

- Docker and docker-compose already installed
- port 8080 and 5434 is free (no app using it)

***

## Setup

1) Clone the repository
2) Open the directory
3) Run `docker-compose up` to build and run the contaianers with showing logs or `docker-compose up -d` without showing logs and `docker-compose logs` to see the logs
4) Run `docker exec -i postgres psql -U postgres -d mezink < dummy.sql` to import the dummy database

***

## Endpoint

```
GET http://localhost:8080/records?startDate=<startDate>&endDate=<endDate>&minCount=<minCount>&maxCount=<minCount>
```
- `<startDate>` and `<endDate>` should be filled with a date in the format yyyy-mm-dd (e.g., 2024-04-01)
- `<minCount>` and `<maxCount>` should be filled with a number (e.g., 300)

### Sample

```
http://localhost:8080/records?2024-04-01startDate=&endDate=2024-04-30&minCount=100&maxCount=500
```

## Test and coverage
Run the following command to execute tests and generate coverage report: 
```
go test ./... -coverprofile=cov.out
```
This will run tests recursively in all directories (./...) and produce a coverage report (cov.out).