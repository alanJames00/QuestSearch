package services

import (
	"context"
	"questsearch/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionServiceImpl struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(repo repositories.QuestionRepository) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		repo: repo,
	}
}

func (s *QuestionServiceImpl) GetAllQuestionsPaginated(ctx context.Context, qType string, page int, limit int) ([]interface{}, *repositories.PageMetaData, error) {
	baseQuestions, pageMetaData, err := s.repo.GetAllQuestionsPaginated(ctx, qType, page, limit)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	// Group questions by type
	mcqIDs := []primitive.ObjectID{}
	anagramIDs := []primitive.ObjectID{}

	for _, question := range baseQuestions {
		switch question.Type {
		case "MCQ":
			mcqIDs = append(mcqIDs, question.ID)
		case "ANAGRAM":
			anagramIDs = append(anagramIDs, question.ID)
		}
	}

	// Fetch details for each type
	mcqDetails, err := s.repo.GetMCQDetails(ctx, mcqIDs)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	anagramDetails, err := s.repo.GetAnagramDetails(ctx, anagramIDs)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	// Combine base questions with details
	var result []interface{}
	for _, question := range baseQuestions {
		switch question.Type {
		case "MCQ":
			if mcq, exists := mcqDetails[question.ID]; exists {
				result = append(result, mcq)
			}
		case "ANAGRAM":
			if anagram, exists := anagramDetails[question.ID]; exists {
				result = append(result, anagram)
			}
		case "CONVERSATION":
			conversation := repositories.ConvertToConversation(question)
			result = append(result, conversation)
		case "READ_ALONG":
			readalong := repositories.ConvertToReadAlong(question)
			result = append(result, readalong)
		case "CONTENT_ONLY":
			contentonly := repositories.ConvertToContentOnly(question)
			result = append(result, contentonly)

		default:
			result = append(result, question)
		}
	}
	return result, pageMetaData, nil
}

func (s *QuestionServiceImpl) SearchQuestionsTitleDyRegex(ctx context.Context, search_term string, qType string, page int, limit int) ([]interface{}, *repositories.PageMetaData, error) {
	baseQuestions, pageMetaData, err := s.repo.SearchQuestionsTitleDyRegex(ctx, search_term, qType, page, limit)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	// Group questions by type
	mcqIDs := []primitive.ObjectID{}
	anagramIDs := []primitive.ObjectID{}

	for _, question := range baseQuestions {
		switch question.Type {
		case "MCQ":
			mcqIDs = append(mcqIDs, question.ID)
		case "ANAGRAM":
			anagramIDs = append(anagramIDs, question.ID)
		}
	}

	// Fetch details for each type
	mcqDetails, err := s.repo.GetMCQDetails(ctx, mcqIDs)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	anagramDetails, err := s.repo.GetAnagramDetails(ctx, anagramIDs)
	if err != nil {
		return nil, &repositories.PageMetaData{}, err
	}

	// Combine base questions with details
	var result []interface{}
	for _, question := range baseQuestions {
		switch question.Type {
		case "MCQ":
			if mcq, exists := mcqDetails[question.ID]; exists {
				result = append(result, mcq)
			}
		case "ANAGRAM":
			if anagram, exists := anagramDetails[question.ID]; exists {
				result = append(result, anagram)
			}
		case "CONVERSATION":
			conversation := repositories.ConvertToConversation(question)
			result = append(result, conversation)
		case "READ_ALONG":
			readalong := repositories.ConvertToReadAlong(question)
			result = append(result, readalong)
		case "CONTENT_ONLY":
			contentonly := repositories.ConvertToContentOnly(question)
			result = append(result, contentonly)

		default:
			result = append(result, question)
		}
	}
	return result, pageMetaData, nil

}
