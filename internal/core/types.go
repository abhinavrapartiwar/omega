package core

import "encoding/json"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type StopReason string

const (
	StopReasonEndTurn   StopReason = "end_turn"
	StopReasonToolUse   StopReason = "tool_use"
	StopReasonMaxTokens StopReason = "max_tokens"
)

type BlockType string

const (
	BlockTypeToolUse BlockType = "tool_use"
	BlockTypeText    BlockType = "text"
	BlockTypeResult  BlockType = "tool_result"
)

type ContentBlock struct {
	Type BlockType `json:"type"`

	// text
	Text string `json:"text,omitempty"`

	// tool_use
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Input json.RawMessage `json:"input,omitempty"`

	// tool_result
	ToolUseID string `json:"tool_use_id,omitempty"`
	Content   string `json:"content,omitempty"`
	IsError   bool   `json:"is_error,omitempty"`
}

func ToolUseBlock(id, name string, input json.RawMessage) ContentBlock {
	return ContentBlock{
		Type:  BlockTypeToolUse,
		ID:    id,
		Name:  name,
		Input: input,
	}
}

func TextBlock(text string) ContentBlock {
	return ContentBlock{
		Text: text,
		Type: BlockTypeText,
	}
}

func ToolResult(toolUseId string, content string, isError bool) ContentBlock {
	return ContentBlock{
		Type:      BlockTypeResult,
		ToolUseID: toolUseId,
		Content:   content,
		IsError:   isError,
	}
}

type Message struct {
	Role    Role
	Content []ContentBlock
}

type ToolDefinition struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"input_schema"`
}

type Usage struct {
	InputTokens     int
	OutputTokens    int
	CacheReadTokens int
}

type CompletionRequest struct {
	System    string
	Messages  []Message
	tools     []ToolDefinition
	maxTokens int
}

type CompletionResponse struct {
	Content    ContentBlock
	StopReason StopReason
	Usage      Usage
}
