package main

import (
	"awesomeProject/flourish/mongo"
	"awesomeProject/flourish/proto/item"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// Environment Variables
const (
	MongoConnStr = "MONGO_CONNECT_STRING"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	mongoConnStr := MustGetEnv(MongoConnStr)
	session := mongo.NewMongoConnection(mongoConnStr)
	defer session.Close()
	mongoDB := session.DB("")
	mongoSvc := mongo.Service{DB: mongoDB}

	srv := grpc.NewServer()

	itemResolverSvc := NewItemSvc(mongoSvc)
	item.RegisterItemResolverServer(srv, itemResolverSvc)

	log.Print("Starting up the userService")
	log.Printf("starting up Server: %v", srv.Serve(lis))
}

// MustGetEnv gets an environment variable or panics.
func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("%s missing", key))
	}
	return v
}
