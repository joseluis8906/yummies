package home

import "github.com/joseluis8906/yummies/go-code/pkg/pb"

type (
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
