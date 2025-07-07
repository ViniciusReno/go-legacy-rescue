package api

import (
	"context"
	"fmt"
)

type UserProcessor interface {
	Process(ctx context.Context, userID string) (string, error)
}

type SimpleUserProcessor struct{}

func NewSimpleUserProcessor() *SimpleUserProcessor {
	return &SimpleUserProcessor{}
}

func (p *SimpleUserProcessor) Process(ctx context.Context, userID string) (string, error) {
	// Example logic (mocked)
	return fmt.Sprintf("Processed user with ID: %s", userID), nil
}

// This would be the equivalent of a handler calling the processor directly
func HandleUserRequest(ctx context.Context, processor UserProcessor, userID string) error {
	result, err := processor.Process(ctx, userID)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
