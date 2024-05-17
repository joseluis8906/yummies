package home

import (
	"context"
	"log"

	"github.com/joseluis8906/yummies/go-code/src/delivery/pkg/pb"
	"github.com/joseluis8906/yummies/go-code/src/pkg/grpc"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

const authEmail string = "x-auth-email"

type (
	Service struct {
		pb.UnimplementedHomeServiceServer
		Nats *nats.Conn
		Log  *log.Logger
	}

	Deps struct {
		fx.In
		Nats *nats.Conn
		Log  *log.Logger
	}
)

func New(deps Deps) *Service {
	s := &Service{
		Nats: deps.Nats,
		Log:  deps.Log,
		// Stores: deps.Stores,
	}

	return s
}

func (s *Service) Home(ctx context.Context, req *pb.HomeRequest) (*pb.HomeResponse, error) {
	ctx, span := otel.Tracer("").Start(ctx, "HomeService.Home")
	defer span.End()

	_, err := grpc.Header(ctx, authEmail).ExpectOne()
	if err != nil {
		s.Log.Printf("getting x-auth-email: %q", err)
		return nil, err
	}

	res := &pb.HomeResponse{
		Categories: []string{"Meals", "Slides", "Snacks", "Drinks", "Healthy", "Fast Food"},
		TodaysSpecialOffer: &pb.HomeTodaysSpecialOffer{
			Name: "Yummies Special Burger",
			Price: &pb.HomeMoney{
				Amount:   275,
				Currency: "USD",
				Decimals: 2,
			},
			Discount: 10.0,
			PriceDiscounted: &pb.HomeMoney{
				Amount:   250,
				Currency: "USD",
				Decimals: 2,
			},
		},
		PopularNow: []*pb.HomePopularNow{
			{
				Name: "Beef Salad",
				Price: &pb.HomeMoney{
					Amount:   425,
					Decimals: 2,
					Currency: "USD",
				},
				IsFavorite: true,
			},
			{
				Name: "Spicy Noodless",
				Price: &pb.HomeMoney{
					Amount:   570,
					Decimals: 2,
					Currency: "USD",
				},
			},
			{
				Name: "Lasagna",
				Price: &pb.HomeMoney{
					Amount:   630,
					Decimals: 2,
					Currency: "USD",
				},
			},
		},
	}

	return res, nil
}
