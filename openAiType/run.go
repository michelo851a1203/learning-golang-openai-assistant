package openAiType

import (
	"testf/openAiType/openAiLastError"
	"testf/openAiType/openAiModel"
	"testf/openAiType/openAiRequiredAction"
	"testf/openAiType/openAiRunStatus"
	"testf/openAiType/openAiStep"
	"testf/openAiType/openAiTool"
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

type RequiredActionObject struct {
	Type              openAiRequiredAction.ActionType `json:"type"`
	SubmitToolOutputs *SubmitToolOutputsObject        `json:"submit_tool_outputs,omitempty"`
}

type SubmitToolOutputsObject struct {
	ToolCalls []*ToolCallsObject `json:"tool_calls"`
}

type ToolCallsObject struct {
	ID       string          `json:"id"`
	Type     openAiTool.Tool `json:"type"`
	Function *FunctionObject `json:"function,omitempty"`
}

type FunctionObject struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type OpenAiRunObject struct {
	ID             string                    `json:"id"`
	Object         string                    `json:"object"`
	CreatedAt      int64                     `json:"created_at"`
	AssistantID    string                    `json:"assistant_id"`
	ThreadID       string                    `json:"thread_id"`
	Status         openAiRunStatus.RunStatus `json:"status"`
	StartedAt      int64                     `json:"started_at"`
	ExpiresAt      *int64                    `json:"expires_at,omitempty"`
	CancelledAt    *int64                    `json:"cancelled_at,omitempty"`
	FailedAt       *int64                    `json:"failed_at,omitempty"`
	CompletedAt    *int64                    `json:"completed_at,omitempty"`
	RequiredAction *RequiredActionObject     `json:"required_action,omitempty"`
	LastError      *LastErrorObject          `json:"last_error"`
	Model          openAiModel.OpenAiModel   `json:"model"`
	Instructions   string                    `json:"instructions"`
	Tools          []*OpenAiTool             `json:"tools"`
	FileIDs        []string                  `json:"file_ids"`
	Metadata       *OpenAiMetaData           `json:"metadata"`
}

type OpenAiRunStepObject struct {
	ID          string                    `json:"id"`
	Object      string                    `json:"object"`
	CreatedAt   int64                     `json:"created_at"`
	RunID       string                    `json:"run_id"`
	AssistantID string                    `json:"assistant_id"`
	ThreadID    string                    `json:"thread_id"`
	Type        openAiStep.StepTypeStatus `json:"type"`
	Status      openAiRunStatus.RunStatus `json:"status"`
	CancelledAt *int64                    `json:"cancelled_at,omitempty"`
	CompletedAt *int64                    `json:"completed_at,omitempty"`
	ExpiredAt   *int64                    `json:"expired_at,omitempty"`
	FailedAt    *int64                    `json:"failed_at,omitempty"`
	LastError   *LastErrorObject          `json:"last_error"`
	StepDetails *StepDetails              `json:"step_details"`
	Metadata    *OpenAiMetaData           `json:"metadata,omitempty"`
}

type StepDetails struct {
	Type            openAiStep.StepTypeStatus `json:"type"`
	MessageCreation *MessageCreation          `json:"message_creation,omitempty"`
	ToolCalls       *[]ToolCalls              `json:"tool_calls,omitempty"`
}

func (detail *StepDetails) CheckValid() bool {
	switch detail.Type {
	case openAiStep.MessageCreation:
		return detail.MessageCreation != nil && detail.ToolCalls == nil
	case openAiStep.ToolCalls:
		return detail.MessageCreation == nil && detail.ToolCalls != nil
	}
	return false
}

type MessageCreation struct {
	MessageID string `json:"message_id"`
}

type ToolCalls struct {
	ID              string                     `json:"id"`
	Type            openAiTool.Tool            `json:"type"`
	CodeInterpreter *CodeInterpreterObject     `json:"code_interpreter,omitempty"`
	Retrieval       *map[string]interface{}    `json:"retrieval,omitempty"`
	Function        *FunctionCallingStepObject `json:"function,omitempty"`
}

func (toolCalls *ToolCalls) CheckValid() bool {
	switch toolCalls.Type {
	case openAiTool.CodeInterpreter:
		return toolCalls.CodeInterpreter != nil && toolCalls.Retrieval == nil && toolCalls.Function == nil
	case openAiTool.Retrieval:
		return toolCalls.CodeInterpreter == nil && toolCalls.Retrieval != nil && toolCalls.Function == nil
	case openAiTool.FunctionCalling:
		return toolCalls.CodeInterpreter == nil && toolCalls.Retrieval == nil && toolCalls.Function != nil
	}
	return false
}

type CodeInterpreterObject struct {
	Input   string              `json:"input"`
	Outputs []*ToolOutputObject `json:"outputs"`
}

type FunctionCallingStepObject struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
	Output    *string                `json:"output"`
}
