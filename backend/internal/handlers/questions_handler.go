// grpc handlers for questions
package handlers

import (
	"fmt"
	"context"
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
	fmt.Printf("Received GetQuestions request: Page=%d, Limit=%d\n", req.Page, req.Limit)

	resp, err := h.formatterService.FormatQuestionsToGrpc(ctx, req)
	if err != nil {
		fmt.Printf("Error formatting questions: %v\n", err)
		return nil, err
	}

	return resp, nil
}

// TODO: 
func (h *GrpcQuestionHandler) GetQuestionById(ctx context.Context, req *question.GetQuestionByIdRequest) (*question.GetQuestionByIdResponse, error) {
	fmt.Printf("Received GetQuestionById request: Id = %s\n", string(req.Id))

	resp := &question.GetQuestionByIdResponse{
		Question: &question.Question{
			Id: "1234",
			Type: "TODO",
			Title: "todo title",
			SiblingId: "todo SiblingId",
		},	
	}

	return resp, nil
} 
