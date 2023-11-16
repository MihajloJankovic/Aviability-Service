package handlers

import (
	"context"
	"fmt"
	protos "github.com/MihajloJankovic/Aviability-Service/protos/main"
	"log"
	"os"
	"time"
	// NoSQL: module containing Mongo api client

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type AviabilityRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*AviabilityRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &AviabilityRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (pr *AviabilityRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *AviabilityRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	// Print available databases
	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}
func GetAccommodationCheck(xtx context.Context, in *protos.CheckRequest) (*protos.Emptyb, error) {
	return nil, nil
}
func GetAllforAccomendation(xtx context.Context, in *protos.Emptyb) (*protos.DummyList, error) {
	return nil, nil
}
func SetAccommodationAviability(xtx context.Context, in *protos.CheckSet) (*protos.Emptyb, error) {
	return nil, nil
}
