package openAiAssistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiType"
)

type AssistantImpl struct {
	ApiKey string
}

func (assistantImpl *AssistantImpl) CreateAssistant(
	createRequest *openAiType.CreateAssistantRequest,
) (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/assistants",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createAssistantClient := &http.Client{}

	response, err := createAssistantClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.AssistantObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) ModifyAssistant(
	assistantID string,
	updateRequest *openAiType.UpdateAssistantRequest,
) (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateAssistantClient := &http.Client{}

	response, err := updateAssistantClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.AssistantObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) DeleteAssistant(assistantID string) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	deleteAssistantClient := &http.Client{}

	response, err := deleteAssistantClient.Do(request)
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

func (assistantImpl *AssistantImpl) GetAssistantList(
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.AssistantObject],
	error,
) {

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants%s", listRequest.ToQueryString()),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantClient := &http.Client{}

	response, err := assistantClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListResponse[openAiType.AssistantObject]{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) GetAssistant(assistantID string) (
	*openAiType.AssistantObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailAssistantClient := &http.Client{}

	response, err := detailAssistantClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.AssistantObject{}
	json.Unmarshal(body, result)

	return result, nil
}
