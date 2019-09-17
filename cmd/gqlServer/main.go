package main

import (
	"awesomeProject/flourish/gql"
	pb "awesomeProject/flourish/proto/item"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func main() {
	// connect to gRPC server of userService
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing userService grpc server: %v", err)
	}

	defer conn.Close()

	itemSvcGRpcClient := pb.NewItemResolverClient(conn)

	resolver := gql.NewResolver(itemSvcGRpcClient)

	schema := graphql.MustParseSchema(gql.Schema, resolver)

	mux := NewRouter(schema)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Starting up the gqlServer")

	log.Fatal(srv.ListenAndServe())
}

func NewRouter(schema *graphql.Schema) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/item", &relay.Handler{Schema: schema})

	return mux
}
