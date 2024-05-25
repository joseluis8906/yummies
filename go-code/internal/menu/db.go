package menu

import "github.com/joseluis8906/yummies/go-code/pkg/pb"

type (
	Product struct {
		Name       string `bson:"name"`
		Img        string `bson:"img"`
		Price      Price  `bson:"price"`
		IsFavorite bool   `bson:"is_favorite"`
	}

	Price struct {
		Amount   uint64 `bson:"amount"`
		Decimals uint8  `bson:"decimals"`
		Currency string `bson:"currency"`
	}
)

func (p Product) PB() *pb.MenuProduct {
	return &pb.MenuProduct{
		Name:       p.Name,
		Img:        p.Img,
		Price:      p.Price.PB(),
		IsFavorite: p.IsFavorite,
	}
}

func (p Price) PB() *pb.MenuMoney {
	return &pb.MenuMoney{
		Amount:   p.Amount,
		Decimals: uint32(p.Decimals),
		Currency: p.Currency,
	}
}
