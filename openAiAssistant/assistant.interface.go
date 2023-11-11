package openAiAssistant

import "testf/openAiType"

type Assistant interface {
	CreateAssistant(createRequest *openAiType.CreateAssistantRequest) (*openAiType.AssistantObject, error)
	ModifyAssistant(assistantID string, updateRequest *openAiType.UpdateAssistantRequest) (*openAiType.AssistantObject, error)
	DeleteAssistant(assistantID string) (*openAiType.DeleteAssistantResponse, error)
	GetAssistantList() (*openAiType.OpenAiListAssistantResponse, error)
	GetAssistant(assistantID string) (*openAiType.AssistantObject, error)
}
