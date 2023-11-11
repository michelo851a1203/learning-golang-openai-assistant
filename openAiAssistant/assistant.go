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
	ApiKey        string
	CreateRequest openAiType.CreateAssistantRequest
	UpdateRequest openAiType.UpdateAssistantRequest
	AssistantId   string
}

func (assistantImpl *AssistantImpl) CreateAssistant() (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(assistantImpl.CreateRequest)
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

func (assistantImpl *AssistantImpl) ModifyAssistant() (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(assistantImpl.UpdateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantImpl.AssistantId),
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

func (assistantImpl *AssistantImpl) DeleteAssistant() (
	*openAiType.DeleteAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantImpl.AssistantId),
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

	result := &openAiType.DeleteAssistantResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) GetAssistantList() (
	*openAiType.OpenAiListAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		"https://api.openai.com/v1/assistants",
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

	result := &openAiType.OpenAiListAssistantResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) GetAssistant() (
	*openAiType.AssistantObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantImpl.AssistantId),
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
