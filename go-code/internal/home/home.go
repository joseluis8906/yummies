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

	Category struct {
		Name string `bson:"name"`
	}

	TodaysSpecialOffer struct {
		Name     string  `bson:"name"`
		Img      string  `bson:"img"`
		Price    Money   `bson:"price"`
		Discount float32 `bson:"discount"`
	}

	PopularNow struct {
		Name       string `bson:"name"`
		Img        string `bson:"img"`
		Price      Money  `bson:"price"`
		IsFavorite bool   `bson:"is_favorite"`
	}

	Money struct {
		Amount   uint64 `bson:"amount"`
		Decimals uint8  `bson:"decimals"`
		Currency string `bson:"currency"`
	}
)

func (m Money) PB() *pb.HomeMoney {
	return &pb.HomeMoney{
		Amount:   m.Amount,
		Decimals: uint32(m.Decimals),
		Currency: m.Currency,
	}
}

func (tso TodaysSpecialOffer) PB() *pb.HomeTodaysSpecialOffer {
	return &pb.HomeTodaysSpecialOffer{
		Name:     tso.Name,
		Img:      tso.Img,
		Price:    tso.Price.PB(),
		Discount: tso.Discount,
	}
}

func (pn PopularNow) PB() *pb.HomePopularNow {
	return &pb.HomePopularNow{
		Name:       pn.Name,
		Img:        pn.Img,
		Price:      pn.Price.PB(),
		IsFavorite: pn.IsFavorite,
	}
}

func New(deps Deps) *Service {
	s := &Service{
		Nats: deps.Nats,
		Log:  deps.Log,
		DB:   deps.DB.Database("yummies"),
	}

	return s
}

func (s *Service) Home(ctx context.Context, req *pb.HomeRequest) (*pb.HomeResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "HomeService.Home")
	defer span.End()

	_, err := grpc.Header(ctx, authEmail).ExpectOne()
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

	res := &pb.HomeResponse{
		Categories:         categories,
		TodaysSpecialOffer: todaysSpecialOffer,
		PopularNow:         popularNow,
	}

	return res, nil
}
