package openAiAssistant

import "testf/openAiType"

type Assistant interface {
	Create() (*openAiType.AssistantObject, error)
	Update() (*openAiType.AssistantObject, error)
	Delete() (*openAiType.DeleteAssistantResponse, error)
	ListAll() (*openAiType.OpenAiListAssistantResponse, error)
	Detail() (*openAiType.AssistantObject, error)
}
