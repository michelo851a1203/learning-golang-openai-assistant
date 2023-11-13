package openAiAssistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testf/openAiType"
)

type AssistantFileImpl struct {
	ApiKey string
}

func (assistantFileImpl *AssistantFileImpl) CreateAssistantFile(
	assistantID string,
	createRequest *openAiType.CreateFileAssistantRequest,
) (
	*openAiType.AssistantFileObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileClient := &http.Client{}

	response, err := assistantFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.AssistantFileObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantFileImpl *AssistantFileImpl) GetAssistantFile(
	assistantID string,
	fileID string,
) (
	*openAiType.AssistantFileObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/assistants/%s/files/%s",
			assistantID,
			fileID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileClient := &http.Client{}

	response, err := assistantFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.AssistantFileObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (assistantFileImpl *AssistantFileImpl) DeleteAssistantFile(
	assistantID string,
	fileID string,
) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf(
			"https://api.openai.com/v1/assistants/%s/files/%s",
			assistantID,
			fileID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileDeleteClient := &http.Client{}

	response, err := assistantFileDeleteClient.Do(request)
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

func (assistantFileImpl *AssistantFileImpl) GetAssistantFileList(
	assistantID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.AssistantFileObject],
	error,
) {
	queryString := ""
	if listRequest != nil {
		reflectValue := reflect.ValueOf(listRequest)
		reflectType := reflectValue.Type()
		queryStringValues := url.Values{}

		if reflectValue.Kind() != reflect.Struct {
			return nil, fmt.Errorf("listRequest is not struct")
		}

		for i := 0; i < reflectValue.NumField(); i++ {
			fieldKey := reflectType.Field(i).Name
			fieldValue := reflectValue.Field(i).Interface()
			queryStringValues.Add(fieldKey, fmt.Sprintf("%v", fieldValue))
		}
		queryString = fmt.Sprintf("?%s", queryStringValues.Encode())
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/assistants/%s/files%s",
			assistantID,
			queryString,
		),
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

	result := &openAiType.ListResponse[openAiType.AssistantFileObject]{}
	json.Unmarshal(body, result)

	return result, nil
}
