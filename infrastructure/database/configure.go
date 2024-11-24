package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func HandleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func Connector() *mongo.Client {
	err := godotenv.Load()
	HandleError("Erro ao carregar o arquivo .env", err)

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	HandleError("Erro ao conectar ao MongoDB", err)

	err = client.Ping(context.Background(), nil)
	HandleError("Erro ao verificar a conexão com o MongoDB", err)

	log.Println("Conexão com MongoDB estabelecida com sucesso")
	return client
}
