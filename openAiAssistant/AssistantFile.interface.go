package openAiAssistant

import "testf/openAiType"

type AssistantFile interface {
	CreateAssistantFile(assistantID string, createRequest *openAiType.CreateFileAssistantRequest) (*openAiType.AssistantFileObject, error)
	GetAssistantFile(assistantID string, fileID string) (*openAiType.AssistantFileObject, error)
	DeleteAssistantFile(assistantID string, fileID string) (*openAiType.DeleteResponse, error)
	GetAssistantFileList(assistantID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.AssistantFileObject], error)
}
