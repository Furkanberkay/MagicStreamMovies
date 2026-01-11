package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Connect() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning :unable to find .env file")
	}

	mongoDb := os.Getenv("MONGODB_URI")
}
