package menu

import (
	"context"
	"log"

	"github.com/joseluis8906/yummies/go-code/pkg/pb"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In
		Nats *nats.Conn
		Log  *log.Logger
		DB   *mongo.Client
	}

	Service struct {
		pb.UnimplementedMenuServiceServer
		Nats *nats.Conn
		Log  *log.Logger
		DB   *mongo.Database
	}
)

func New(deps Deps) *Service {
	return &Service{
		Nats: deps.Nats,
		Log:  deps.Log,
		DB:   deps.DB.Database("yummies"),
	}
}

func (s *Service) Index(ctx context.Context, req *pb.MenuIndexRequest) (*pb.MenuIndexResponse, error) {
	cur, err := s.DB.Collection("menu").Find(ctx, bson.D{})
	if err != nil {
		s.Log.Printf("getting menu from db: %v", err)
		return nil, err
	}

	var dbProducts []Product
	if err := cur.All(ctx, &dbProducts); err != nil {
		s.Log.Printf("marshaling products: %v", err)
		return nil, err
	}

	products := make([]*pb.MenuProduct, len(dbProducts))
	for i, v := range dbProducts {
		products[i] = v.PB()
	}

	return &pb.MenuIndexResponse{Products: products}, nil
}
