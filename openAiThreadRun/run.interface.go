package openAiThreadRun

import "testf/openAiType"

type ThreadRun interface {
	CreateRun(threadID string, createRequest *openAiType.CreateThreadRunRequest) (*openAiType.OpenAiRunObject, error)
	ModifyRun(threadID, runID string, updateRequest *openAiType.UpdateThreadRunRequest) (*openAiType.OpenAiRunObject, error)
	GetRunList(threadID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.OpenAiRunObject], error)
	GetRun(threadID, runID string) (*openAiType.OpenAiRunObject, error)
	SubmitToolOutputToRun(threadID, runID string, toolOutputRequest *openAiType.SubmitOutputsAndRunRequest) (*openAiType.OpenAiRunObject, error)
	CancelRun(threadID, runID string) (*openAiType.OpenAiRunObject, error)
	CreateThreadAndRun(threadAndRunRequest *openAiType.ThreadAndRunRequest) (*openAiType.OpenAiRunObject, error)
	GetRunStep(threadID, runID, stepID string) (*openAiType.OpenAiRunStepObject, error)
	GetRunStepList(threadID, runID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.OpenAiRunStepObject], error)
}
