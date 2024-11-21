package home

import (
	"context"
	"fmt"
	"log"

	"github.com/joseluis8906/yummies/go-code/pkg/pb"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

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
	wg := new(errgroup.Group)

	var categories []string
	wg.Go(func() error {
		cur, err := s.DB.Collection("home.categories").Find(ctx, bson.D{})
		if err != nil {
			s.Log.Printf("getting categories: %v", err)
			return fmt.Errorf("getting categories: %v", err)
		}

		var dbCategories []Category
		err = cur.All(ctx, &dbCategories)
		if err != nil {
			s.Log.Printf("parsing categories: %v", err)
			return fmt.Errorf("parsing categories: %v", err)
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
		err := s.DB.Collection("home.todaysspecialoffer").FindOne(ctx, bson.D{}).Decode(&dbTso)
		if err != nil {
			return fmt.Errorf("getting todaysspecialoffer: %v", err)
		}

		todaysSpecialOffer = dbTso.PB()
		return nil
	})

	var popularNow []*pb.HomePopularNow
	wg.Go(func() error {
		cur, err := s.DB.Collection("home.popularnow").Find(ctx, bson.D{})
		if err != nil {
			return fmt.Errorf("getting popularnow: %v", err)
		}

		var dbPn []PopularNow
		err = cur.All(ctx, &dbPn)
		if err != nil {
			return fmt.Errorf("parsing popularnow: %v", err)
		}

		popularNow = make([]*pb.HomePopularNow, len(dbPn))
		for i, v := range dbPn {
			popularNow[i] = v.PB()
		}

		return nil
	})

	if err := wg.Wait(); err != nil {
		s.Log.Println(err)
		return nil, err
	}

	res := &pb.HomeIndexResponse{
		Categories:         categories,
		TodaysSpecialOffer: todaysSpecialOffer,
		PopularNow:         popularNow,
	}

	return res, nil
}
