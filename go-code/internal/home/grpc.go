package home

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
		pb.UnimplementedHomeServiceServer
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

func (s *Service) Index(ctx context.Context, req *pb.HomeIndexRequest) (*pb.HomeIndexResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "yummies.HomeService.Index")
	defer span.End()

	_, err := grpc.Header(ctx, authEmail)
	if err != nil {
		s.Log.Printf("getting x-auth-email: %q", err)
		return nil, fmt.Errorf("getting x-auth-email header: %w", err)
	}

	cur, err := s.DB.Collection("home.categories").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var dbCategories []Category
	err = cur.All(ctx, &dbCategories)
	if err != nil {
		return nil, err
	}

	var dbTso TodaysSpecialOffer
	err = s.DB.Collection("home.todaysspecialoffer").FindOne(ctx, bson.D{}).Decode(&dbTso)
	if err != nil {
		return nil, err
	}

	cur, err = s.DB.Collection("home.popularnow").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var dbPn []PopularNow
	err = cur.All(ctx, &dbPn)
	if err != nil {
		return nil, err
	}

	categories := make([]string, len(dbCategories))
	for i, v := range dbCategories {
		categories[i] = v.Name
	}

	todaysSpecialOffer := dbTso.PB()
	popularNow := make([]*pb.HomePopularNow, len(dbPn))
	for i, v := range dbPn {
		popularNow[i] = v.PB()
	}

	res := &pb.HomeIndexResponse{
		Categories:         categories,
		TodaysSpecialOffer: todaysSpecialOffer,
		PopularNow:         popularNow,
	}

	return res, nil
}
