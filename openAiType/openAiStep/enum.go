package openAiStep

type StepTypeStatus string

const (
	MessageCreation StepTypeStatus = "message_creation"
	ToolCalls       StepTypeStatus = "tool_calls"
)
