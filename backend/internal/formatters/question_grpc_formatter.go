package formatters

import (
	"context"
	"fmt"
	"questsearch/internal/models"
	"questsearch/internal/services"
	"questsearch/proto/pb/question"
)

type GrpcFormatterService struct {
	service services.QuestionService
}

func NewGrpcFormatterService(service services.QuestionService) *GrpcFormatterService {
	return &GrpcFormatterService{
		service: service,
	}
}

// formats generic []interface{} result from service layer for grpc request response
func (s *GrpcFormatterService) GetAllQuestionsPaginated(ctx context.Context, req *question.GetQuestionsRequest) (*question.GetQuestionsResponse, error) {

	questionSlice, err := s.service.GetAllQuestionsPaginated(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, err
	}

	return formatQuestionsToGrpc(questionSlice)
}

func (s *GrpcFormatterService) SearchQuestionsTitleDyRegex(ctx context.Context, req *question.SearchQuestionRequest) (*question.GetQuestionsResponse, error) {

	questionSlice, err := s.service.SearchQuestionsTitleDyRegex(ctx, req.SearchTerm, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, err
	}

	return formatQuestionsToGrpc(questionSlice)
}

// generic private function to convert service results to grpc responses
func formatQuestionsToGrpc(questionSlice []interface{}) (*question.GetQuestionsResponse, error) {

	// prepare the grpc response slice
	var grpcQuestions []*question.Question

	for _, item := range questionSlice {
		switch v := item.(type) {
		case models.Anagram:

			// nil check on SiblingId
			SafesiblingId := ""
			if v.SiblingID != nil {
				SafesiblingId = v.SiblingID.String()
			}

			grpcQuestions = append(grpcQuestions, &question.Question{
				Id:        v.ID.Hex(),
				Type:      v.Type,
				SiblingId: SafesiblingId,
				Title:     v.Title,
				QuestionDetails: &question.Question_Anagram{
					Anagram: &question.AnagramQuestion{
						Blocks:   mapBlocksToGrpc(v.Blocks),
						Solution: v.Solution,
					},
				},
			})

		case models.ContentOnly:
			SafesiblingId := ""
			if v.SiblingID != nil {
				SafesiblingId = v.SiblingID.String()
			}

			grpcQuestions = append(grpcQuestions, &question.Question{
				Id:        v.ID.Hex(),
				Type:      v.Type,
				SiblingId: SafesiblingId,
				Title:     v.Title,
			})

		case models.Conversation:
			SafesiblingId := ""
			if v.SiblingID != nil {
				SafesiblingId = v.SiblingID.String()
			}

			grpcQuestions = append(grpcQuestions, &question.Question{
				Id:        v.ID.Hex(),
				Type:      v.Type,
				SiblingId: SafesiblingId,
				Title:     v.Title,
			})

		case models.MCQ:
			SafesiblingId := ""
			if v.SiblingID != nil {
				SafesiblingId = v.SiblingID.String()
			}

			grpcQuestions = append(grpcQuestions, &question.Question{
				Id:        v.ID.Hex(),
				Type:      v.Type,
				SiblingId: SafesiblingId,
				Title:     v.Title,
				QuestionDetails: &question.Question_Mcq{
					Mcq: &question.MCQQuestion{
						Options: mapOptionsToGrpc(v.Options),
					},
				},
			})

		case models.ReadAlong:
			SafesiblingId := ""
			if v.SiblingID != nil {
				SafesiblingId = v.SiblingID.String()
			}

			grpcQuestions = append(grpcQuestions, &question.Question{
				Id:        v.ID.Hex(),
				Type:      v.Type,
				SiblingId: SafesiblingId,
				Title:     v.Title,
			})

		default:
			fmt.Printf("Unknown type: %T\n", v)
		}

	}

	return &question.GetQuestionsResponse{Questions: grpcQuestions}, nil
}
