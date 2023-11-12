package openaiThreads

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiType"
)

type ThreadsImpl struct {
	ApiKey string
}

func (threadsImpl *ThreadsImpl) CreateThread(
	createRequest *openAiType.ThreadCreateRequest,
) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/threads",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadsClient := &http.Client{}

	response, err := createThreadsClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiThreadObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadsImpl *ThreadsImpl) ModifyThread(
	threadsID string,
	updateRequest *openAiType.UpdateThreadsRequest,
) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadsID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateThreadsClient := &http.Client{}

	response, err := updateThreadsClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiThreadObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadsImpl *ThreadsImpl) DeleteThread(threadID string) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	deleteThreadsClient := &http.Client{}

	response, err := deleteThreadsClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.DeleteResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (threadsImpl *ThreadsImpl) GetThread(threadID string) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailThreadsClient := &http.Client{}

	response, err := detailThreadsClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiThreadObject{}
	json.Unmarshal(body, result)

	return result, nil
}
