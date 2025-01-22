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

func (s *QuestionServiceImpl) GetAllQuestionsPaginated(ctx context.Context, page int, limit int) ([]interface{}, error) {
	baseQuestions, err := s.repo.GetAllQuestionsPaginated(ctx, page, limit)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	anagramDetails, err := s.repo.GetAnagramDetails(ctx, anagramIDs)
	if err != nil {
		return nil, err
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
	return result, nil
}

func (s *QuestionServiceImpl) SearchQuestionsTitleDyRegex(ctx context.Context, search_term string, page int, limit int) ([]interface{}, error) {
	baseQuestions, err := s.repo.SearchQuestionsTitleDyRegex(ctx, search_term, page, limit)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	anagramDetails, err := s.repo.GetAnagramDetails(ctx, anagramIDs)
	if err != nil {
		return nil, err
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
	return result, nil

}
