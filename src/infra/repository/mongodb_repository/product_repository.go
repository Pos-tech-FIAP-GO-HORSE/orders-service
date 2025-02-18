package mongodb_repository

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository/mongodb_repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) repository.IProductRepository {
	return &ProductRepository{collection}
}

func (ref *ProductRepository) Create(ctx context.Context, product entity.Product) (*entity.Product, error) {
	record := models.ProductFromDomain(product)

	result, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return ref.FindByID(ctx, id)
}

func (ref *ProductRepository) Find(ctx context.Context) ([]*entity.Product, error) {
	cursor, err := ref.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	products := make([]*entity.Product, 0)

	for cursor.Next(ctx) {
		var record models.Product
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}

		product := record.ToDomain()
		products = append(products, &product)
	}

	return products, nil
}

func (ref *ProductRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})

	var record models.Product
	if err = result.Decode(&record); err != nil {
		return nil, err
	}

	product := record.ToDomain()

	return &product, nil
}

func (ref *ProductRepository) FindByPublicID(ctx context.Context, publicID string) (*entity.Product, error) {
	result := ref.collection.FindOne(ctx, bson.M{"public_id": publicID})

	var record models.Product
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	product := record.ToDomain()

	return &product, nil
}

func (ref *ProductRepository) UpdateByID(ctx context.Context, id string, product entity.Product) (*entity.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	record := models.ProductFromDomain(product)

	_, err = ref.collection.UpdateByID(ctx, objectID, record)
	if err != nil {
		return nil, err
	}

	return ref.FindByID(ctx, id)
}

func (ref *ProductRepository) DeleteByID(ctx context.Context, id string) (*entity.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result := ref.collection.FindOne(ctx, filter)

	var record models.Product
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	product := record.ToDomain()

	_, err = ref.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
