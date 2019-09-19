package gql

import (
	"context"
	"github.com/amitbasuri/awesomeProject/awesome"
	pb "github.com/amitbasuri/awesomeProject/awesome/proto/item"
	"github.com/golang/protobuf/ptypes"
	"github.com/graph-gophers/graphql-go"
)

var Schema = `
   schema {
       query: Query
       mutation: Mutation
   }
   # The Query type represents all of the entry points.
   type Query {
       Item(Id: String!): Item
   }

   type Mutation {
		CreateItem(name: String!, description: String!, price: Float!): Item
	}
   scalar Time
   type Item {
       id: String!
       name: String!
       description: String!
       price: Float!
       created_at: Time!
   }
   `

type Resolver struct {
	ItemSvc pb.ItemResolverClient
}

func NewResolver(itemClient pb.ItemResolverClient) *Resolver {
	return &Resolver{
		ItemSvc: itemClient,
	}
}

type ItemResolver struct {
	I *awesome.Item
}

// CreateItem Resolves the CreateItem mutation
func (r *Resolver) CreateItem(ctx context.Context, args struct {
	Name, Description string
	Price             float64
}) (*ItemResolver, error) {
	request := &pb.CreateItemReq{
		Name:        args.Name,
		Description: args.Description,
		Price:       args.Price,
	}

	res, err := r.ItemSvc.CreateItem(ctx, request)
	if err != nil {
		return nil, err
	}
	t, _ := ptypes.Timestamp(res.CreatedAt)
	item := &awesome.Item{
		ID:          res.Id,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		CreatedAt:   graphql.Time{Time: t},
	}
	return &ItemResolver{I: item}, err
}

func (r *Resolver) Item(ctx context.Context, args awesome.Item) (*ItemResolver, error) {
	request := &pb.QueryItemReq{
		Id: args.ID,
	}

	res, err := r.ItemSvc.Item(ctx, request) //(ctx, request)
	if err != nil {
		return nil, err
	}
	t, _ := ptypes.Timestamp(res.CreatedAt)
	item := &awesome.Item{
		ID:          res.Id,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		CreatedAt:   graphql.Time{Time: t},
	}
	return &ItemResolver{I: item}, err
}

// Resolve each field to respond to queries.
func (r *ItemResolver) ID() string {
	return r.I.ID
}

func (r *ItemResolver) Name() string {
	return r.I.Name
}

func (r *ItemResolver) Description() string {
	return r.I.Description
}

func (r *ItemResolver) Price() float64 {
	return r.I.Price
}

func (r *ItemResolver) CreatedAt() graphql.Time {
	return r.I.CreatedAt
}
