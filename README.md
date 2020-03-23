# awesomeProject

[![Go Report Card](https://goreportcard.com/badge/github.com/amitbasuri/awesomeProject)](https://goreportcard.com/report/github.com/amitbasuri/awesomeProject)

A simple app of gRPC + GraphQL + Go + MongoDB.

## microservices:

1. First microservice accepts GraphQL requests from Front-End (Library: 
github.com/graph-gophers/graphql-go) and send to the second microservice via gRPC.

2. Second microservice accepts gRPC requests from the first one and gets/saves 
it from Mongodb.

### Run:
```bash
MONGO_CONNECT_STRING='mongodb://localhost/awesomedb' docker-compose up --build
```

### Create an item
```bash
curl -i -H 'Content-Type: application/json'  -X POST -d '{"query": "mutation {CreateItem(name:\"Iphone 11\",description:\"Iphone 11 desc\",price:855.69) {id,name,description,price,created_at}}"}' localhost:8080/item 
```
### get an item by id
```bash
curl -i -H 'Content-Type: application/json'  -X POST -d '{"query": "query {Item(Id: \"2ba56bfe-5505-416c-a5ba-baaacaffbd88\") {id,name,description,price,created_at}}"}' localhost:8080/item 
```
