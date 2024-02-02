#!/bin/bash

# CatFact API Server

## Introduction

CatFact API Server is a simple Go application that provides cat facts through a RESTful API.

## Features

- Fetches cat facts from an external API
- Logs request duration and errors using Logrus
- Exposes a simple RESTful API for retrieving cat facts

## Installation

To install the CatFact API Server, ensure you have Go installed on your machine:

``bash
go get -u github.com/karalakrepp/MicroServicesLogic





## Usage
Run the CatFact API Server:

``bash
    cd $GOPATH/src/github.com/karalakrepp/MicroServicesLogic
    go run main.go



### API Endpoints
. Get Cat Fact 
URL: /catfact
Method: GET
Response:
Status Code: 200 OK on success
Status Code: 500 Internal Server Error on failure
Body: JSON containing a cat fact or an error message
Example Request: