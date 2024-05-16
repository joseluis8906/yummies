package store

import (
	"context"
	"fmt"
	"log"

	"github.com/joseluis8906/go-code/src/pkg/repository"
	"go.opentelemetry.io/otel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/fx"
)

// CatalogRepository represents the repository for the catalog.
type (
	Deps struct {
		fx.In

		Conn *mongo.Client
		Logs *log.Logger
	}

	Repository struct {
		logs *log.Logger

		client     *mongo.Client
		db         string
		collection string
	}
)

// NewCatalogRepository creates a new assistant repository instance.
func NewRepository(deps Deps) (*Repository, error) {
	db := "delivery"
	collection := "stores"
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "name",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := deps.Conn.Database(db).
		Collection(collection).
		Indexes().
		CreateMany(context.TODO(), indexes)
	if err != nil {
		log.Printf("error! : %v", err)
		return nil, err
	}

	repo := &Repository{
		logs:       deps.Logs,
		client:     deps.Conn,
		db:         db,
		collection: collection,
	}

	return repo, nil
}

// Matchig returns the assistant for the given system.
func (r *Repository) Get(ctx context.Context, criteria repository.Criteria, page int64) ([]Store, error) {
	ctx, span := otel.Tracer("").Start(ctx, "store.Repository.Get")
	defer span.End()

	var result []Store
	opts := options.Find().SetLimit(repository.Limit).SetSkip(repository.Page(page))
	query := bson.D{
		{
			Key: criteria.Field,
			Value: bson.D{
				{
					Key:   criteria.Operator,
					Value: criteria.Value,
				},
			},
		},
	}

	ev := fmt.Sprintf("db.stores.find({\"%s\": {\"%s\": \"%s\"}})", criteria.Field, criteria.Operator, criteria.Value)
	span.AddEvent(ev)
	cursor, err := r.client.Database(r.db).Collection(r.collection).Find(ctx, query, opts)
	if err != nil {
		return nil, (fmt.Errorf("searching in mongo: %w", err))
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, (fmt.Errorf("decoding from mongo cursor: %w", err))
	}

	return result, nil
}

func (r *Repository) Add(ctx context.Context, aStore Store) error {
	filter := bson.D{
		{
			Key:   Fields().Name,
			Value: aStore.Name.Value,
		},
	}

	_, err := r.client.Database(r.db).Collection(r.collection).
		ReplaceOne(ctx, filter, aStore, options.Replace().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("creating or replacing in mongo: %w", err)
	}

	return nil
}
