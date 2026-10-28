package core

import (
	"context"
	"encoding/json"
)

type Provider interface {
	Complete(ctx context.Context, req CompletionRequest) (CompletionResponse, error)
}

type ToolExecutor interface {
	Definition() []ToolDefinition
	Execute(ctx context.Context, input json.RawMessage) (content string, isError bool)
}
