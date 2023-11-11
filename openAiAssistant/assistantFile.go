package openAiAssistant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiType"
)

type AssistantFileImpl struct {
	ApiKey      string
	AssistantID string
}

func (assistantFileImpl *AssistantFileImpl) CreateAssistantFile() (
	*openAiType.OpenAiListAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantFileImpl.AssistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
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

func (assistantFileImpl *AssistantFileImpl) GetAssistantFile() (
	*openAiType.OpenAiListAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantFileImpl.AssistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
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

func (assistantFileImpl *AssistantFileImpl) DeleteAssistantFile() (
	*openAiType.OpenAiListAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantFileImpl.AssistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
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

func (assistantFileImpl *AssistantFileImpl) GetAssistantFileList() (
	*openAiType.OpenAiListAssistantResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantFileImpl.AssistantID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
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
