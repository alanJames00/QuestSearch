package repositories

import (
	"context"
	"questsearch/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// repository interface definition that should be implemented by mongodb
type QuestionRepository interface {
	// TODO: later for extended question types
	GetAllQuestionsPaginated(ctx context.Context, page int, limit int) ([]models.BaseQuestion, error)
	SearchQuestionsTitleDyRegex(ctx context.Context, search_term string, page int, limit int) ([]models.BaseQuestion, error)

	GetMCQDetails(ctx context.Context, ids []primitive.ObjectID) (map[primitive.ObjectID]models.MCQ, error)
	GetAnagramDetails(ctx context.Context, ids []primitive.ObjectID) (map[primitive.ObjectID]models.Anagram, error)
}
