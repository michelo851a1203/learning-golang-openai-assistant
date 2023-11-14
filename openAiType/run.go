package openAiType

import (
	"testf/openAiType/openAiLastError"
	"testf/openAiType/openAiModel"
	"testf/openAiType/openAiRunStatus"
)

type ToolOutputObject struct {
	ToolCallID string `json:"tool_call_id,omitempty"`
	OutputName string `json:"output_name,omitempty"`
}

type SubmitOutputsAndRunRequest struct {
	ToolOutputs []*ToolOutputObject `json:"tool_outputs"`
}

type LastErrorObject struct {
	Code    openAiLastError.LastErrorCode `json:"code"`
	Message string                        `json:"message"`
}

type OpenAiRunObject struct {
	ID           string                    `json:"id"`
	Object       string                    `json:"object"`
	CreatedAt    int64                     `json:"created_at"`
	AssistantID  string                    `json:"assistant_id"`
	ThreadID     string                    `json:"thread_id"`
	Status       openAiRunStatus.RunStatus `json:"status"`
	StartedAt    int64                     `json:"started_at"`
	ExpiresAt    *int64                    `json:"expires_at,omitempty"`
	CancelledAt  *int64                    `json:"cancelled_at,omitempty"`
	FailedAt     *int64                    `json:"failed_at,omitempty"`
	CompletedAt  *int64                    `json:"completed_at,omitempty"`
	LastError    *LastErrorObject          `json:"last_error"`
	Model        openAiModel.OpenAiModel   `json:"model"`
	Instructions string                    `json:"instructions"`
	Tools        []*OpenAiTool             `json:"tools"`
	FileIDs      []*OpenAiFileObject       `json:"file_ids"`
	Metadata     *OpenAiMetaData           `json:"metadata"`
}
