package product

import (
	"fmt"
	"regexp"

	"github.com/joseluis8906/yummies/go-code/src/pkg/financial"
	"github.com/joseluis8906/yummies/go-code/src/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	Ref         types.StringValue
	Name        types.StringValue
	Price       struct{ financial.Money }
	Description types.StringValue
	Raiting     types.FloatValue
)

func NewRef(value string) (Ref, error) {
	re := regexp.MustCompile(`^[A-Z]{3}\-[0-9]{3}$`)
	if !re.MatchString(value) {
		return Ref{}, fmt.Errorf("invalid product reference")
	}

	return Ref{value, true}, nil
}

func NewName(value string) (Name, error) {
	if len(value) < 3 {
		return Name{}, fmt.Errorf("invalid product name")
	}

	return Name{value, true}, nil
}

func NewPrice(amount int64, currency string) (Price, error) {
	v, err := financial.NewMoney(amount, financial.Currency(currency))
	if err != nil {
		return Price{}, fmt.Errorf("invalid price: %w", err)
	}

	return Price{v}, nil
}

func (n Ref) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Value)
}

func (n *Ref) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &n.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	n.Valid = true

	return nil
}

func (n Name) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(n.Value)
}

func (n *Name) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	if t != bson.TypeString {
		return fmt.Errorf("invalid bson type '%s'", t.String())
	}

	if err := bson.UnmarshalValue(t, val, &n.Value); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	n.Valid = true

	return nil
}

func (p Price) MarshalBSON() ([]byte, error) {
	price := struct {
		Amount   int64
		Currency string
	}{
		Amount:   p.Amount,
		Currency: string(p.Currency),
	}

	return bson.Marshal(price)
}

func (n *Price) UnmarshalBSON(val []byte) error {
	var price struct {
		Amount   int64
		Currency string
	}

	if err := bson.Unmarshal(val, &price); err != nil {
		return fmt.Errorf("unmarshaling value: %w", err)
	}

	v, err := financial.NewMoney(price.Amount, financial.Currency(price.Currency))
	if err != nil {
		return fmt.Errorf("invalid price: %w", err)
	}

	n.Money = v

	return nil
}

func NewDesctiption(val string) (Description, error) {
	return Description{val, true}, nil
}

func NewRaiting(val float32) (Raiting, error) {
	if val < 0 || val > 5 {
		return Raiting{}, fmt.Errorf("invalid raiting value")
	}
	return Raiting{val, true}, nil
}
