package storemanager

import (
	// "context"
	// "encoding/json"
	// "errors"
	// "fmt"
	"log"

	// pb "github.com/joseluis8906/yummies/go-code/protobuf/delivery/storemanagerpb"
	// logl "github.com/joseluis8906/yummies/go-code/internal/app/log"
	// "github.com/joseluis8906/yummies/go-code/internal/product"
	"github.com/joseluis8906/yummies/go-code/internal/store"
	// "github.com/joseluis8906/yummies/go-code/src/pkg/grpc"
	// "github.com/joseluis8906/yummies/go-code/src/pkg/repository"
	"github.com/nats-io/nats.go"
	// "go.opentelemetry.io/otel"
	"go.uber.org/fx"
	// epb "google.golang.org/protobuf/types/known/emptypb"
)

const authEmail string = "x-auth-email"

type (
	// StoreManager represent a store manager.
	// It is responsible for creating and updating stores and products.
	// It is also responsible for reporting that an order has been taken and
	// also if the order is ready to be delivered.
	StoreManager struct {
		// pb.UnimplementedStoreManagerServiceServer
		Nats   *nats.Conn
		Log    *log.Logger
		Stores *store.Repository
	}

	Deps struct {
		fx.In
		Nats   *nats.Conn
		Log    *log.Logger
		Stores *store.Repository
	}
)

// NewDeliveryService returns a new instance of DeliveryService.
func New(deps Deps) *StoreManager {
	s := &StoreManager{
		Nats:   deps.Nats,
		Log:    deps.Log,
		Stores: deps.Stores,
	}

	// s.Nats.Subscribe("storemanager.add_store", s.OnAddStore)
	return s
}

// AddStore creates a new store with the provided data.
// If the store already exists, it will return an error.
// The store's products are discarted.
// func (s *StoreManager) AddStore(ctx context.Context, req *pb.AddStoreRequest) (*epb.Empty, error) {
// 	ctx, span := otel.Tracer("").Start(ctx, "storemanager.StoreManager.AddStore")
// 	defer span.End()

// 	_, err := grpc.Header(ctx, authEmail).ExpectOne()
// 	if err != nil {
// 		s.Log.Printf("getting x-auth-email: %q", err)
// 		return nil, err
// 	}

// 	newStore, err := store.FromPB(req.GetStore())
// 	if err != nil {
// 		return nil, fmt.Errorf("creating new store: %w", err)
// 	}

// 	var page int64 = 1
// 	query := repository.Query("name = ?", fmt.Sprintf("%s", newStore.Name))
// 	stores, err := s.Stores.Get(ctx, query, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("verifing exisiting store: %w", err)
// 	}

// 	if len(stores) > 0 {
// 		return nil, errors.New("store already exists")
// 	}

// 	err = s.Stores.Add(ctx, newStore)
// 	if err != nil {
// 		return nil, fmt.Errorf("persisting new store: %w", err)
// 	}

// 	data, _ := json.Marshal(newStore)
// 	s.Nats.Publish("storemanager.add_store", data)

// 	return &epb.Empty{}, nil
// }

// AddProduct adds a new product to a store.
// If the store does not exist, it will return an error.
// If products already exist, it will update the existing product with new data.
// func (s *StoreManager) AddProduct(ctx context.Context, form *pb.AddProductRequest) (*epb.Empty, error) {
// 	ctx, span := otel.Tracer("").Start(ctx, "storemanager.AddProduct")
// 	defer span.End()

// 	_, err := grpc.Header(ctx, authEmail).ExpectOne()
// 	if err != nil {
// 		s.Log.Printf("getting x-auth-email: %q", err)
// 		return nil, err
// 	}

// 	name, err := store.NewName(form.GetStore().GetName().GetValue())
// 	if err != nil {
// 		return nil, fmt.Errorf("casting store name: %w", err)
// 	}

// 	var page int64 = 1
// 	query := repository.Query("name = ?", fmt.Sprintf("%s", name))
// 	stores, err := s.Stores.Get(ctx, query, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting store: %w", err)
// 	}

// 	if len(stores) != 1 {
// 		return nil, errors.New("store does not exist or there are more than one")
// 	}

// 	store := stores[0]
// 	catalog := make(map[product.Ref]product.Product, len(store.Products))
// 	for _, p := range store.Products {
// 		catalog[p.Ref] = p
// 	}

// 	newPrd, err := product.FromPB(form.GetProduct())
// 	if err != nil {
// 		return nil, fmt.Errorf("creating new product: %w", err)
// 	}

// 	catalog[newPrd.Ref] = newPrd
// 	store.Products = make([]product.Product, 0, len(catalog))
// 	for _, p := range catalog {
// 		store.Products = append(store.Products, p)
// 	}

// 	err = s.Stores.Add(ctx, store)
// 	if err != nil {
// 		return nil, fmt.Errorf("persisting store: %w", err)
// 	}

// 	return &epb.Empty{}, nil
// }

// func (s *StoreManager) OnAddStore(msg *nats.Msg) {
// 	s.Log.Printf(logl.Info("event happend, store added: %s"), msg.Data)
// }
