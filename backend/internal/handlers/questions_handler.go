// grpc handlers for questions
package handlers

import (
	"context"
	"log"
	"questsearch/internal/formatters"
	"questsearch/proto/pb/question"
)

// GrpcQuestionHandler handles gRPC requests for questions
type GrpcQuestionHandler struct {
	formatterService *formatters.GrpcFormatterService
	question.UnimplementedQuestionServiceServer
}

// func to create instance of GrpcQuestionHandler
func NewGrpcHandler(formatter *formatters.GrpcFormatterService) *GrpcQuestionHandler {
	return &GrpcQuestionHandler{
		formatterService: formatter,
	}
}

// GetQuestions handles the gRPC request to retrieve questions

func (h *GrpcQuestionHandler) GetQuestions(ctx context.Context, req *question.GetQuestionsRequest) (*question.GetQuestionsResponse, error) {
	// Use the formatter service to handle the request and format the response
	log.Printf("Received GetQuestions request: Page=%d, Limit=%d\n", req.Page, req.Limit)

	resp, err := h.formatterService.GetAllQuestionsPaginated(ctx, req)
	if err != nil {
		log.Printf("Error formatting questions: %v\n", err)
		return nil, err
	}

	return resp, nil
}

// TODO:
func (h *GrpcQuestionHandler) GetQuestionById(ctx context.Context, req *question.GetQuestionByIdRequest) (*question.GetQuestionByIdResponse, error) {
	log.Printf("Received GetQuestionById request: Id = %s\n", string(req.Id))

	resp := &question.GetQuestionByIdResponse{
		Question: &question.Question{
			Id:        "1234",
			Type:      "TODO",
			Title:     "todo title",
			SiblingId: "todo SiblingId",
		},
	}

	return resp, nil
}

func (h *GrpcQuestionHandler) SearchQuestion(ctx context.Context, req *question.SearchQuestionRequest) (*question.GetQuestionsResponse, error) {
	log.Printf("Received SearchQuestion request: search_term = %s page = %d limit = %d\n", req.SearchTerm, req.Page, req.Limit)

	resp, err := h.formatterService.SearchQuestionsTitleDyRegex(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
