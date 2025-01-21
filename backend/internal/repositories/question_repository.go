// repository layer for fetching questions from mongodb
package repositories

import (
	"context"
	"questsearch/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoQuestionRepository struct {
	collection *mongo.Collection
}

// instantiate a new instance of MongoQuestionRepository
func NewMongoQuestionRepository(db *mongo.Database, collectionName string) *MongoQuestionRepository {
	return &MongoQuestionRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoQuestionRepository) GetAllQuestionsPaginated(ctx context.Context, page int, limit int) ([]models.BaseQuestion, error) {
	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	var questions []models.BaseQuestion

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// decode the cursor to questions
	err = cursor.All(ctx, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// get mcqType typed questions from ad id slice of them
func (r *MongoQuestionRepository) GetMCQDetails(ctx context.Context, ids []primitive.ObjectID) (map[primitive.ObjectID]models.MCQ, error) {

	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	mcqMap := make(map[primitive.ObjectID]models.MCQ)
	var mcqs []models.MCQ

	if err := cursor.All(ctx, &mcqs); err != nil {
		return nil, err
	}

	for _, mcq := range mcqs {
		mcqMap[mcq.ID] = mcq
	}

	return mcqMap, nil

}

func (r *MongoQuestionRepository) GetAnagramDetails(ctx context.Context, ids []primitive.ObjectID) (map[primitive.ObjectID]models.Anagram, error) {

	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	anagramMap := make(map[primitive.ObjectID]models.Anagram)
	var anagrams []models.Anagram

	if err = cursor.All(ctx, &anagrams); err != nil {
		return nil, err
	}

	for _, anagram := range anagrams {
		anagramMap[anagram.ID] = anagram
	}

	return anagramMap, nil

}
