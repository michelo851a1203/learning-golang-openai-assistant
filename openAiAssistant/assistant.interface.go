package openAiAssistant

import "testf/openAiType"

type Assistant interface {
	CreateAssistant(createRequest *openAiType.CreateAssistantRequest) (*openAiType.AssistantObject, error)
	ModifyAssistant(assistantID string, updateRequest *openAiType.UpdateAssistantRequest) (*openAiType.AssistantObject, error)
	DeleteAssistant(assistantID string) (*openAiType.DeleteResponse, error)
	GetAssistantList() (*openAiType.ListResponse[openAiType.AssistantObject], error)
	GetAssistant(assistantID string) (*openAiType.AssistantObject, error)
}
