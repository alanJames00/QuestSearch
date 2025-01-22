// contains helper functions for converting Go structs for questions to grpc messagep
package formatters

import (
	"questsearch/internal/models"
	"questsearch/proto/pb/question"
)

// takes a blocks slice of Go struct and convert to grpc question block repeated slice
func mapBlocksToGrpc(blocks []models.Block) []*question.Block {
	grpcBlocks := make([]*question.Block, len(blocks))

	for i, block := range blocks {
		grpcBlocks[i] = &question.Block{
			Text:         block.Text,
			ShowInOption: block.ShowInOption,
			IsAnswer:     block.IsAnswer,
		}
	}

	return grpcBlocks
}

func mapOptionsToGrpc(options []models.Option) []*question.Option {
	grpcOptions := make([]*question.Option, len(options))

	for i, option := range options {
		grpcOptions[i] = &question.Option{
			Text:            option.Text,
			IsCorrectAnswer: option.IsCorrectAnswer,
		}
	}

	return grpcOptions
}
