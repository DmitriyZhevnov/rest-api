package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewUserMongo(database *mongo.Database, collection string, logger *logging.Logger) *userMongo {
	return &userMongo{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *userMongo) Create(ctx context.Context, user model.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", apperror.NewInternalServerError(fmt.Sprintf("failed to create user due to error: %v", err), "45645234")
	}

	d.logger.Debug("conver insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", apperror.NewInternalServerError(fmt.Sprintf("failed to convert objectID to hex. probably oid: %s", oid), "3456543456")
}

func (d *userMongo) FindAll(ctx context.Context) (u []model.User, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, apperror.NewInternalServerError(fmt.Sprintf("failed to find all users due to error: %v", err), "2234234")
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, apperror.NewInternalServerError(fmt.Sprintf("failed to read all documents from cursor. error: %v", err), "245646")
	}

	return u, nil
}

func (d *userMongo) FindOne(ctx context.Context, id string) (u model.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, apperror.NewErrNotFound(fmt.Sprintf("failed to convert hex to ObjectID. hex: %s", id), "456462")
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperror.NewErrNotFound("user not exists", "2546461")
		}

		return u, apperror.NewInternalServerError(fmt.Sprintf("failed to find one user by id: %s due to error: %v", id, err), "234234247")
	}

	if err = result.Decode(&u); err != nil {
		return u, apperror.NewInternalServerError(fmt.Sprintf("failed to decode user(id: %s) from DB due to error: %v", id, err), "2345676543")
	}

	return u, nil
}

func (d *userMongo) Update(ctx context.Context, user model.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return apperror.NewErrNotFound(fmt.Sprintf("failed to convert userID to ObjectID. ID=%s", user.ID), "5654234234")
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to marshal user. Error: %v", err), "345423765")
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to unmarshal user bytes. error: %v", err), "234457603")
	}

	delete(updateUserObj, "_id")

	if user.Email == "" {
		delete(updateUserObj, "email")
	}
	if user.PasswordHash == "" {
		delete(updateUserObj, "password")
	}
	if user.Username == "" {
		delete(updateUserObj, "username")
	}

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to execute update user query. error: %v", err), "94530904858")
	}
	if result.MatchedCount == 0 {
		return apperror.NewErrNotFound("user not found", "34765342")
	}

	d.logger.Trace("Matched %d documents and modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *userMongo) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return apperror.NewErrNotFound(fmt.Sprintf("failed to convert userID to ObjectID. ID=%s", id), "2094852")
	}
	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to execute query. error: %v", err), "9495854932")
	}
	if result.DeletedCount == 0 {
		return apperror.NewErrNotFound("user not found", "2346234")
	}

	d.logger.Trace("Deleted %d documents", result.DeletedCount)

	return nil
}
