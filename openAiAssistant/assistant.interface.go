package openAiAssistant

import "testf/openAiType"

type Assistant interface {
	CreateAssistant() (*openAiType.AssistantObject, error)
	ModifyAssistant() (*openAiType.AssistantObject, error)
	DeleteAssistant() (*openAiType.DeleteAssistantResponse, error)
	GetAssistantList() (*openAiType.OpenAiListAssistantResponse, error)
	GetAssistant() (*openAiType.AssistantObject, error)
}
