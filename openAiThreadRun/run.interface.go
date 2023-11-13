package openAiThreadRun

import "testf/openAiType"

type ThreadRun interface {
	CreateRun(threadID string, createRequest *openAiType.CreateAssistantRequest) (*openAiType.OpenAiRunObject, error)
	ModifyRun(threadID, runID string, updateRequest *openAiType.UpdateAssistantRequest) (*openAiType.OpenAiRunObject, error)
	GetRunList(threadID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.OpenAiRunObject], error)
	GetRun(threadID, runID string) (*openAiType.OpenAiRunObject, error)
}
