package openAiThreadRun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiType"
)

type ThreadRunImpl struct {
	ApiKey string
}

func (threadRunImpl *ThreadRunImpl) CreateRun(
	threadID string,
	createRequest *openAiType.CreateThreadRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs",
			threadID,
		),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) ModifyRun(
	threadID,
	runID string,
	updateRequest *openAiType.UpdateThreadRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s",
			threadID,
			runID,
		),

		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateThreadRunClient := &http.Client{}

	response, err := updateThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) GetRunList(
	threadID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiRunObject],
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs%s",
			threadID,
			listRequest.ToQueryString(),
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	ThreadRunClient := &http.Client{}

	response, err := ThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListResponse[openAiType.OpenAiRunObject]{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) GetRun(threadID, runID string) (
	*openAiType.OpenAiRunObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,

		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s",
			threadID,
			runID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailThreadRunClient := &http.Client{}

	response, err := detailThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) SubmitToolOutputToRun(
	threadID,
	runID string,
	toolOutputRequest *openAiType.SubmitOutputsAndRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {

	requestInfo, err := json.Marshal(toolOutputRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/submit_tool_outputs",
			threadID,
			runID,
		),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) CancelRun(
	threadID,
	runID string,
) (
	*openAiType.OpenAiRunObject,
	error,
) {

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/cancel",
			threadID,
			runID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) CreateThreadAndRun(
	threadAndRunRequest *openAiType.ThreadAndRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(threadAndRunRequest)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/threads/runs",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil

}

func (threadRunImpl *ThreadRunImpl) GetRunStep(
	threadID, runID, stepID string,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/steps/%s",
			threadID,
			runID,
			stepID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiRunObject{}
	json.Unmarshal(body, result)

	return result, nil

}

func (threadRunImpl *ThreadRunImpl) GetRunStepList(
	threadID,
	runID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiRunStepObject],
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/steps%s",
			threadID,
			runID,
			listRequest.ToQueryString(),
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListResponse[openAiType.OpenAiRunStepObject]{}
	json.Unmarshal(body, result)

	return result, nil
}
