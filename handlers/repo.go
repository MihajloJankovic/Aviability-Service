package handlers

import (
	"context"
	"errors"
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

func (pr *AviabilityRepo) DeleteByUser(xtx context.Context, in *protos.DeleteRequestb) (*protos.Emptyb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	profileCollection := pr.getCollection()

	// Define the filter to find documents where 'created' is greater than the specified date
	filter := bson.M{"uid": in.GetUid()}
	// Perform the find operation
	_, err := profileCollection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, errors.New("Coundnt delete")
	}

	return new(protos.Emptyb), nil
}
func (pr *AviabilityRepo) GetAccommodationCheck(xtx context.Context, in *protos.CheckRequest) (*protos.CheckSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	profileCollection := pr.getCollection()

	// Define the filter to find documents where 'created' is greater than the specified date
	filter := bson.D{
		{"from", bson.D{
			{"$lt", in.GetFrom()},
		}},
		{"to", bson.D{
			{"$gt", in.GetTo()},
		}},
	}
	// Perform the find operation
	accommodationAviabilityCursor, err := profileCollection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	defer accommodationAviabilityCursor.Close(ctx)

	for accommodationAviabilityCursor.Next(ctx) {
		var accommodation protos.CheckSet
		if err := accommodationAviabilityCursor.Decode(&accommodation); err != nil {
			pr.logger.Println(err)
			return nil, err
		}
		return &accommodation, nil
	}
	return nil, errors.New("No result")
}
func (pr *AviabilityRepo) GetAllforAccomendation(xtx context.Context, in *protos.GetAllRequest) ([]*protos.CheckSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationCollection := pr.getCollection()
	var accommodationAviabilitysSlice []*protos.CheckSet

	// Assuming you have a filter based on the email, modify the filter as needed
	filter := bson.M{"uid": in.GetId()}

	accommodationAviabilityCursor, err := accommodationCollection.Find(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	defer func(accommodationCursor *mongo.Cursor, ctx context.Context) {
		err := accommodationCursor.Close(ctx)
		if err != nil {
			pr.logger.Println(err)
		}
	}(accommodationAviabilityCursor, ctx)

	for accommodationAviabilityCursor.Next(ctx) {
		var accommodation protos.CheckSet
		if err := accommodationAviabilityCursor.Decode(&accommodation); err != nil {
			pr.logger.Println(err)
			return nil, err
		}
		accommodationAviabilitysSlice = append(accommodationAviabilitysSlice, &accommodation)
	}

	if err := accommodationAviabilityCursor.Err(); err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	return accommodationAviabilitysSlice, nil
}
func (pr *AviabilityRepo) SetAccommodationAviability(xtx context.Context, in *protos.CheckSet) (*protos.Emptyb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	avaCollection := pr.getCollection()

	result, err := avaCollection.InsertOne(ctx, &in)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return new(protos.Emptyb), nil
}
func (pr *AviabilityRepo) getCollection() *mongo.Collection {
	profileDatabase := pr.cli.Database("mongoAviability")
	profileCollection := profileDatabase.Collection("aviability")
	return profileCollection
}

func (pr *AviabilityRepo) GetallbyIDandPrice(ctx context.Context, in *protos.PriceAndIdRequest) (*protos.DummyLista3, error) {
	avaCollection := pr.getCollection()

	// Prilagodi filter prema tvojim potrebama
	filter := bson.M{
		"uid": in.GetId(),
		"$and": []bson.M{
			{"price_per_person": bson.M{"$gte": in.GetMinPrice()}},
			{"price_per_person": bson.M{"$lte": in.GetMaxPrice()}},
		},
	}

	cursor, err := avaCollection.Find(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var accommodationAvailabilities []*protos.CheckSet

	for cursor.Next(ctx) {
		var accommodation protos.CheckSet
		if err := cursor.Decode(&accommodation); err != nil {
			pr.logger.Println(err)
			return nil, err
		}
		accommodationAvailabilities = append(accommodationAvailabilities, &accommodation)
	}

	if err := cursor.Err(); err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	dummyResponse := &protos.DummyLista3{
		Dummy: accommodationAvailabilities,
	}

	return dummyResponse, nil
}
