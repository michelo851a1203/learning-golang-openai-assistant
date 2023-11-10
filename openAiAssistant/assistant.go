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
	AssistantId   string
}

func (assistantImpl *AssistantImpl) Create() (
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

func (assistantImpl *AssistantImpl) Update() (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(assistantImpl.CreateRequest)
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

func (assistantImpl *AssistantImpl) Delete() (
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

	result := &openAiType.DeleteAssistantResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) ListAll() (
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

	result := &openAiType.OpenAiListAssistantResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantImpl *AssistantImpl) Detail() (
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
