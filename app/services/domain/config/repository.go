package config

import (
	"context"

	"github.com/saifwork/portfolio-service.git/app/services/domain/config/dtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PortfolioRepository handles MongoDB operations for Config
type PortfolioRepository struct {
	collection *mongo.Collection
}

// NewPortfolioRepository initializes the repository
func NewPortfolioRepository(client *mongo.Client, dbName string) *PortfolioRepository {
	return &PortfolioRepository{
		collection: client.Database(dbName).Collection("contact"),
	}
}

// Insert adds a new config document
func (r *PortfolioRepository) Insert(ctx context.Context, contact *dtos.ContactReqDto) (*mongo.InsertOneResult, error) {
	contact.ID = primitive.NewObjectID()
	return r.collection.InsertOne(ctx, contact)
}
