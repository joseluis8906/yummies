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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
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

	wg := new(errgroup.Group)
	var categories []string
	wg.Go(func() error {
		ctx, span := otel.Tracer("").Start(ctx, "mongodb.Find")
		defer span.End()
		span.AddEvent("get categories", trace.WithAttributes(attribute.String("query", "db.home.categories.find({})")))
		cur, err := s.DB.Collection("home.categories").Find(ctx, bson.D{})
		if err != nil {
			return err
		}

		var dbCategories []Category
		err = cur.All(ctx, &dbCategories)
		if err != nil {
			return err
		}

		categories := make([]string, len(dbCategories))
		for i, v := range dbCategories {
			categories[i] = v.Name
		}

		return nil
	})

	var todaysSpecialOffer *pb.HomeTodaysSpecialOffer
	wg.Go(func() error {
		var dbTso TodaysSpecialOffer
		ctx, span := otel.Tracer("").Start(ctx, "mongodb.FindOne")
		defer span.End()
		span.AddEvent("get todaysspecialoffer", trace.WithAttributes(attribute.String("query", "db.home.todaysspecialoffer.findOne({})")))
		err = s.DB.Collection("home.todaysspecialoffer").FindOne(ctx, bson.D{}).Decode(&dbTso)
		if err != nil {
			return err
		}

		todaysSpecialOffer = dbTso.PB()
		return nil
	})

	var popularNow []*pb.HomePopularNow
	wg.Go(func() error {
		ctx, span := otel.Tracer("").Start(ctx, "mongodb.Find")
		defer span.End()
		span.AddEvent("get popularnow", trace.WithAttributes(attribute.String("query", "db.home.popularnow.find({})")))
		cur, err := s.DB.Collection("home.popularnow").Find(ctx, bson.D{})
		if err != nil {
			return err
		}

		var dbPn []PopularNow
		err = cur.All(ctx, &dbPn)
		if err != nil {
			return err
		}

		popularNow = make([]*pb.HomePopularNow, len(dbPn))
		for i, v := range dbPn {
			popularNow[i] = v.PB()
		}

		return nil
	})

	wg.Wait()
	res := &pb.HomeIndexResponse{
		Categories:         categories,
		TodaysSpecialOffer: todaysSpecialOffer,
		PopularNow:         popularNow,
	}

	return res, nil
}
