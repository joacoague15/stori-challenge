# StoriApp Challenge

StoriApp is a Golang application to process transaction data from a CSV file and send information via email. This project includes Docker easy building and running.

## Prerequisites

- [Go](https://golang.org/doc/install) (if you plan to run or build locally without Docker)
- [Docker](https://docs.docker.com/get-docker/) (for building and running the application within a Docker container)

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development, testing purposes, or production environment.

### Local Setup

1. Clone the repository:

```sh
git clone https://yourrepositoryurl.com/StoriApp.git
cd StoriApp
```

2. Build the application:

```sh
go build -o storiapp 
```

3. Run the application:
```sh
./storiapp
```

### Running with Docker

1. Build the Docker image:

```sh
docker build -t storiapp .
```

2. Run the Docker container:

```sh
docker run storiapp
```

### What things can be added in the future?
#### 1. Front-end for selecting the email to send the data and choose different csv files to summarize
#### 2. Some kind of statistics to analyze the behavior between debit and credit transactions or check how many of them are invalid
