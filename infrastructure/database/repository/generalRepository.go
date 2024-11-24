package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Name      string  `json:"name" bson:"name"`
	Price     float64 `json:"price" bson:"price"`
	DtCreated string  `json:"dt_created" bson:"dt_created"`
	DtUpdated string  `json:"dt_updated" bson:"dt_updated"`
}

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client, dbName, collectionName string) *ProductRepository {
	return &ProductRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

// Listar Produtos
func (r *ProductRepository) GetProducts(ctx context.Context, filter bson.M) ([]Product, error) {
	var products []Product

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("erro ao decodificar produto: %w", err)
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Criar Produto
func (r *ProductRepository) CreateProduct(ctx context.Context, product Product) (string, error) {
	product.DtCreated = time.Now().Format("2006-01-02 15:04:05")
	product.DtUpdated = product.DtCreated

	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(string)
	if !ok {
		return "", fmt.Errorf("falha ao converter o ID gerado")
	}

	return id, nil
}
