package openaiThreads

import "testf/openAiType"

type Threads interface {
	CreateThread(createRequest *openAiType.ThreadCreateRequest) (*openAiType.OpenAiThreadObject, error)
	ModifyThread(threadsID string, updateRequest *openAiType.UpdateThreadsRequest) (*openAiType.OpenAiThreadObject, error)
	DeleteThread(threadID string) (*openAiType.DeleteResponse, error)
	GetThread(threadID string) (*openAiType.OpenAiThreadObject, error)
}
