// Package repository
package repository

import (
	"context"
	"time"

	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/domain"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookRepository interface {
	Create(ctx context.Context, b *domain.Book) (*domain.Book, error)
	GetByID(ctx context.Context, id string) (*domain.Book, error)
	List(ctx context.Context, limit int64) ([]*domain.Book, error)
	Delete(ctx context.Context, id string) error
}

type mongoRepo struct {
	col *mongo.Collection
	log logger.Logger
}

func NewMongoRepo(db *mongo.Database, logger logger.Logger) BookRepository {
	col := db.Collection("books")
	// ensure index on ownerId if needed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "ownerId", Value: 1}},
	})
	if err != nil {
		logger.Error("failed to create index on ownerId", err)
	}

	return &mongoRepo{col: col, log: logger}
}

func (m *mongoRepo) Create(ctx context.Context, b *domain.Book) (*domain.Book, error) {
	if b.ID == "" {
		b.ID = primitive.NewObjectID().Hex()
	}

	bsonDoc, err := toBSON(b)
	if err != nil {
		return nil, err
	}
	res, err := m.col.InsertOne(ctx, bsonDoc)
	if err != nil {
		return nil, err
	}
	b.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return b, nil
}

func (m *mongoRepo) GetByID(ctx context.Context, id string) (*domain.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var out domain.Book
	err = m.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&out)
	if err != nil {
		return nil, err
	}
	out.ID = objID.Hex() // keep string ID
	return &out, nil
}

func (m *mongoRepo) List(ctx context.Context, limit int64) ([]*domain.Book, error) {
	opts := options.Find().SetLimit(limit).SetSort(bson.D{{Key: "_id", Value: -1}})
	cur, err := m.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*domain.Book
	for cur.Next(ctx) {
		var b domain.Book
		if err := cur.Decode(&b); err != nil {
			return nil, err
		}
		results = append(results, &b)
	}
	return results, nil
}

func (m *mongoRepo) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.col.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func toBSON(b *domain.Book) (bson.M, error) {
	// convert ID if present
	var id interface{}
	if b.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(b.ID); err == nil {
			id = oid
		}
	}
	return bson.M{
		"_id":      id,
		"ownerId":  b.OwnerID,
		"title":    b.Title,
		"author":   b.Author,
		"photos":   b.Photos,
		"price":    b.Price,
		"lend":     b.Lend,
		"location": bson.M{"lat": b.Location.Lat, "lng": b.Location.Lng},
	}, nil
}
