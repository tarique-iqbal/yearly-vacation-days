# Yearly vacation days calculation
Determine the number of vacation days for each of its employees for a given year. The script should take the year of interest as its only input argument and output the name of each employee along with the respective number of vacation days for the given year.

### Requirements
- Each employee has a minimum of 26 vacation days
- A special contract can overwrite the minimum amount of vacation days
- Employees with an age >= 30 years get one additional vacation day every 5 years of
employment
- Contracts may start on the 1st or 15th of the month

## Prerequisites
```
Go 1.24.1
Docker (Optional for Containerized Development)
```

## Environment Setup
- Clone the repository:
```shell
$ git clone https://github.com/tarique-iqbal/yearly-vacation-days.git
$ cd /path/to/project/directory
```
- Ensure Go modules are initialized:
```shell
$ go mod tidy
```

## Installation & Running the Code
### Run the Code Locally (Without Docker)
- Build the application
```shell
$ go build -o yearly-vacation-days cmd/main.go
```
- Run the application for a specific year (e.g., 2025)
```shell
$ ./yearly-vacation-days 2025
Vacation Days Report:
Hans Müller: 26
Angelika Fringe: 26
Peter Klever: 28
Marina Helter: 26
Sepp Meier: 2.16
```
- Using `go run` (Without Building)
```shell
$ go run cmd/main.go 2024
Vacation Days Report:
Hans Müller: 26
Angelika Fringe: 26
Peter Klever: 27
Marina Helter: 24.91
Sepp Meier: Not applicable
```

### Run the Code in Docker
- Build the Docker Image
```shell
$ docker build -t yearly_vacation_days .
```
- Run the Container and Code
```shell
$ docker run --rm yearly_vacation_days 2024
```
- For Local Development (Mounting Code as a Volume)
```shell
$ docker run -it --rm -v $(pwd):/home/app -w /home/app --entrypoint bash yearly_vacation_days
```

### Running the Tests
- Run All Tests
```shell
$ go test ./tests/...
```
- Run Tests in Docker
```shell
$ docker run -it --rm -v $(pwd):/home/app -w /home/app --entrypoint bash yearly_vacation_days
$ go test ./tests/...
```
