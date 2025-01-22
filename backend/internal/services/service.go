// service layer for questions handling
package services

import (
	"context"
	"questsearch/internal/repositories"
)

// interface of methods on QuestionService
type QuestionService interface {
	GetAllQuestionsPaginated(ctx context.Context, page int, limit int) ([]interface{}, *repositories.PageMetaData, error)
	SearchQuestionsTitleDyRegex(ctx context.Context, search_term string, page int, limit int) ([]interface{}, *repositories.PageMetaData, error)
}
