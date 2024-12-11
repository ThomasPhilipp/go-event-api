
# go-rest-api

## Endpoints

### Get a list of events
GET /events

### Get an event
GET /events/<id>

### Create an event
POST /events (auth required)

### Update an event
PUT /events/<id> (auth required, creator scoped by JWT)

### Delete an event
DELETE /events/<id> (auth required, creator scoped by JWT)

### Create new user
POST /signup

### Authenticate user
POST /login (JWT token)

### Register user for event
POST /events/<id>/register (auth required)

### Cancel registration
DELETE /events/<id>/register (auth required)

## Libraries
* Gin: a HTTP web framework

## Tools
* REST Client for Viusal Studio Code(https://marketplace.visualstudio.com/items?itemName=humao.rest-client)


## Project Setup

* Create folder: $ mkdir go-event-api
* Switch to folder: $ cd ./go-event-api
* Create project: $ go mod init zwickit.com/go-event-api
* Install packages:
 * $ go get -u github.com/gin-gonic/gin

