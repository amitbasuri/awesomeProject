package main

import (
	"github.com/RichardKnop/uuid"
	"time"

	"awesomeProject/awesome"
	"awesomeProject/awesome/mongo"
	pb "awesomeProject/awesome/proto/item"
	"github.com/golang/protobuf/ptypes"
	"github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	svc mongo.Service
}

func NewItemSvc(svc mongo.Service) pb.ItemResolverServer {
	return &Server{svc}
}

func (svc *Server) Item(ctx context.Context, in *pb.QueryItemReq) (*pb.QueryItemRes, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be blank")
	}
	item, err := svc.svc.GetItemByID(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	t, _ := ptypes.TimestampProto(item.CreatedAt.Time)

	res := &pb.QueryItemRes{
		Id:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		CreatedAt:   t,
	}

	return res, nil
}

func (svc *Server) CreateItem(ctx context.Context, in *pb.CreateItemReq) (*pb.QueryItemRes, error) {
	item := &awesome.Item{
		ID:          uuid.New(),
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		CreatedAt:   graphql.Time{Time: time.Now()},
	}
	err := svc.svc.CreateItem(item)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	t, _ := ptypes.TimestampProto(item.CreatedAt.Time)

	res := &pb.QueryItemRes{
		Id:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		CreatedAt:   t,
	}

	return res, nil
}
