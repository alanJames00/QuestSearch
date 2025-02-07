// repository layer for fetching questions from mongodb
package repositories

import (
	"context"
	"fmt"
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

func (r *MongoQuestionRepository) GetAllQuestionsPaginated(ctx context.Context, qType string, page int, limit int) ([]models.BaseQuestion, *PageMetaData, error) {
	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	// build a query with filter using qType
	query := bson.M{}

	// get the qType and use it as filter
	if qType != "" && qType != "ALL" {
		query = bson.M{
			"type": qType,
		}
	}

	var questions []models.BaseQuestion

	cursor, err := r.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, &PageMetaData{}, err
	}
	defer cursor.Close(ctx)

	totalResults, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, &PageMetaData{}, err
	}

	// decode the cursor to questions
	err = cursor.All(ctx, &questions)
	if err != nil {
		return nil, &PageMetaData{}, err
	}

	totalPages := (int(totalResults) + limit - 1) / limit // ceil division

	return questions, &PageMetaData{
		TotalResults: int(totalResults),
		TotalPages:   totalPages,
	}, nil
}

func (r *MongoQuestionRepository) SearchQuestionsTitleDyRegex(ctx context.Context, search_term string, qType string, page int, limit int) ([]models.BaseQuestion, *PageMetaData, error) {

	searchRegex := buildTypoTolerantRegex(search_term)
	fmt.Println("generated regex for search_term", searchRegex)

	skip := (page - 1) * limit
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	filter := bson.M{
		"title": bson.M{
			"$regex":   searchRegex,
			"$options": "i", // for case insensitive
		},
	}

	if qType != "" && qType != "ALL" {
		filter = bson.M{
			"type": qType,
			"title": bson.M{
				"$regex":   searchRegex,
				"$options": "i",
			},
		}
	}

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, &PageMetaData{}, err
	}
	defer cursor.Close(ctx)

	totalResults, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, &PageMetaData{}, err
	}

	var questions []models.BaseQuestion

	// decode the cursor to questions
	if err = cursor.All(ctx, &questions); err != nil {
		return nil, &PageMetaData{}, err
	}

	totalPages := (int(totalResults) + limit - 1) / limit // ceil division

	return questions, &PageMetaData{
		TotalResults: int(totalResults),
		TotalPages:   totalPages,
	}, nil
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

func ConvertToContentOnly(baseQuestion models.BaseQuestion) models.ContentOnly {
	return models.ContentOnly{
		ID:        baseQuestion.ID,
		Type:      baseQuestion.Type,
		Title:     baseQuestion.Title,
		SiblingID: baseQuestion.SiblingID,
	}
}

func ConvertToReadAlong(baseQuestion models.BaseQuestion) models.ReadAlong {
	return models.ReadAlong{
		ID:        baseQuestion.ID,
		Type:      baseQuestion.Type,
		Title:     baseQuestion.Title,
		SiblingID: baseQuestion.SiblingID,
	}
}

func ConvertToConversation(baseQuestion models.BaseQuestion) models.Conversation {
	return models.Conversation{
		ID:        baseQuestion.ID,
		Type:      baseQuestion.Type,
		Title:     baseQuestion.Title,
		SiblingID: baseQuestion.SiblingID,
	}
}

func buildTypoTolerantRegex(searchTerm string) string {
	var regex string

	// allow each char of searchTerm or its absence
	for _, char := range searchTerm {
		if char == ' ' {
			// Allow the absence of space or any sequence of characters
			regex += ".*"
		} else {
			// Preserve the character as-is
			regex += string(char)
		}
	}

	// Wrap the entire regex to enforce case-insensitivity and match the full term
	return regex
}
