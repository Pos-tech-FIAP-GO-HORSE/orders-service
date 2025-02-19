package mongodb_repository

import (
	"context"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository/mongodb_repository/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) repository.IOrderRepository {
	return &OrderRepository{collection}
}

func (ref *OrderRepository) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
	record := models.OrderFromDomain(order)
	record.PublicID = uuid.NewString()

	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	result, err := ref.collection.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return ref.FindByID(ctx, id)
}

func (ref *OrderRepository) Find(ctx context.Context) ([]*entity.Order, error) {
	cursor, err := ref.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	orders := make([]*entity.Order, 0)

	for cursor.Next(ctx) {
		var record models.Order
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}

		order := record.ToDomain()
		orders = append(orders, &order)
	}

	return orders, nil
}

func (ref *OrderRepository) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := ref.collection.FindOne(ctx, bson.M{"_id": objectID})

	var record models.Order
	if err = result.Decode(&record); err != nil {
		return nil, err
	}

	order := record.ToDomain()

	return &order, nil
}

func (ref *OrderRepository) FindByPublicID(ctx context.Context, publicID string) (*entity.Order, error) {
	result := ref.collection.FindOne(ctx, bson.M{"publicId": publicID})

	var record models.Order
	if err := result.Decode(&record); err != nil {
		return nil, err
	}

	order := record.ToDomain()

	return &order, nil
}

func (ref *OrderRepository) UpdateByID(ctx context.Context, id string, order entity.Order) (*entity.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	record := models.OrderFromDomain(order)
	record.UpdatedAt = time.Now()

	_, err = ref.collection.UpdateByID(ctx, objectID, record)
	if err != nil {
		return nil, err
	}

	return ref.FindByID(ctx, id)
}

func (ref *OrderRepository) UpdateStatusByID(ctx context.Context, id string, status string) (*entity.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"status":    status,
			"updatedAt": time.Now(),
		},
	}

	_, err = ref.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return nil, err
	}

	return ref.FindByID(ctx, id)
}
