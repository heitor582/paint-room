# Description

This project is a test for be BackEnd developer at Digital Republic

### Stack

- GoLang 1.17

# Clone

```bash

# Clone github project

$ git clone https://gitlab.com/heitor582/digitalrepublictest.git



# Enter the folder

$ cd ./digitalRepublicTest/back-paint-room

```

# Instalation

## Running the app with Dockerfile

```bash

# Build the image

$ docker build -t back-room-paint .



# Run dockerfile

$ docker run -d -p 8080:8080 back-room-paint

```

## Running the app without Dockerfile

### Start the app

```bash

# Build the project

$ go build -o /back-paint-room



# Run app

$ back-paint-room

```

# Unit Test

```bash

# unit tests

$ go test ./...

```

# Import archive of Postman

### First what is it Postman?

Postman is a program to make, organize and view the result of api requests.

### Running the app

At the root of the project is a json that contains data that the postman program processes and transforms into pre-made requests.

In the postman click import -> Upload Files

