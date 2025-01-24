// grpc handlers for questions
package handlers

import (
	"context"
	"log"
	"questsearch/internal/formatters"
	"questsearch/proto/pb/question"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Printf("Received GetQuestions request: qType =%s Page=%d, Limit=%d\n", req.QType, req.Page, req.Limit)

	// validate inputs
	if (req.Page <= 0) || (req.Limit <= 0) {
		return nil, status.Errorf(codes.InvalidArgument, "page and limit must be greater than 0")
	}

	// validate qType input
	if req.QType != "" {
		valid := false

		validTypes := []string{"ALL", "ANAGRAM", "MCQ", "CONVERSATION", "CONTENT_ONLY", "READ_ALONG"}

		for _, v := range validTypes {
			if v == req.QType {
				valid = true
				break
			}
		}

		if !valid {
			return nil, status.Errorf(codes.InvalidArgument, "invalid question type: %v", req.QType)
		}
	}

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
	log.Printf("Received SearchQuestion request: search_term = %s qType = %s page = %d limit = %d\n", req.SearchTerm, req.QType, req.Page, req.Limit)

	// validate inputs
	if (req.Page <= 0) || (req.Limit <= 0) {
		return nil, status.Errorf(codes.InvalidArgument, "page and limit must be greater than 0")
	}

	// validate qType input
	if req.QType != "" {
		valid := false

		validTypes := []string{"ALL", "ANAGRAM", "MCQ", "CONVERSATION", "CONTENT_ONLY", "READ_ALONG"}

		for _, v := range validTypes {
			if v == req.QType {
				valid = true
				break
			}
		}

		if !valid {
			return nil, status.Errorf(codes.InvalidArgument, "invalid question type: %v", req.QType)
		}
	}

	resp, err := h.formatterService.SearchQuestionsTitleDyRegex(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
