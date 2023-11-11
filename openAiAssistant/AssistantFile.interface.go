package openAiAssistant

import "testf/openAiType"

type AssistantFile interface {
	CreateAssistantFile(assistantID string, createRequest *openAiType.CreateFileAssistantRequest) (*openAiType.OpenAiAssistantFileResponse, error)
	GetAssistantFile(assistantID string, fileID string) (*openAiType.OpenAiAssistantFileResponse, error)
	DeleteAssistantFile(assistantID string, fileID string) (*openAiType.OpenAiAssistantFileDeleteResponse, error)
	GetAssistantFileList(assistantID string) (*openAiType.OpenAiAssistantFileListResponse, error)
}
