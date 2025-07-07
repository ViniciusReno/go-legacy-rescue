package tests

import (
	"context"
	"errors"
	"testing"

	"go-legacy-rescue/v2_refactored/api"

	"github.com/stretchr/testify/assert"
)

type FailingProcessor struct{}

func (p *FailingProcessor) Process(ctx context.Context, userID string) (string, error) {
	return "", errors.New("processing failed")
}

func TestSimpleUserProcessor_Process_Success(t *testing.T) {
	processor := api.NewSimpleUserProcessor()
	ctx := context.Background()
	userID := "12345"

	result, err := processor.Process(ctx, userID)

	assert.NoError(t, err)
	assert.Contains(t, result, userID)
}

func TestSimpleUserProcessor_Process_ContextCancelled(t *testing.T) {
	processor := api.NewSimpleUserProcessor()
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // simulate immediate cancellation

	result, err := processor.Process(ctx, "12345")

	// This test assumes no context handling in processor, so result is still returned
	assert.NoError(t, err)
	assert.Contains(t, result, "12345")
}

func TestFailingProcessor_Process_Error(t *testing.T) {
	processor := &FailingProcessor{}
	ctx := context.Background()

	result, err := processor.Process(ctx, "user-x")

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.EqualError(t, err, "processing failed")
}
