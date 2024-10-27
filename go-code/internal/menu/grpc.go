package menu

import (
	"context"
	"fmt"
	"log"

	"github.com/joseluis8906/yummies/go-code/pkg/grpc"
	"github.com/joseluis8906/yummies/go-code/pkg/pb"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

const authEmail string = "x-auth-email"

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
	ctx, span := otel.Tracer("").Start(ctx, "yummies.MenuService.Index")
	defer span.End()

	_, err := grpc.Header(ctx, authEmail)
	if err != nil {
		s.Log.Printf("getting x-auth-email: %q", err)
		return nil, fmt.Errorf("getting x-auth-email header: %w", err)
	}

	cur, err := s.DB.Collection("menu").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var dbProducts []Product
	err = cur.All(ctx, &dbProducts)
	if err != nil {
		return nil, err
	}

	products := make([]*pb.MenuProduct, len(dbProducts))
	for i, v := range dbProducts {
		products[i] = v.PB()
	}

	return &pb.MenuIndexResponse{Products: products}, nil
}
