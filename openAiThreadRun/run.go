package openAiThreadRun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type ThreadRunImpl struct {
	ApiKey string
}

func (ThreadRunImpl *ThreadRunImpl) CreateRun(
	createRequest *openAiType.CreateThreadRunRequest,
) (
	*openAiType.ThreadRunObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/ThreadRuns",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ThreadRunImpl.ApiKey))
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

	result := &openAiType.ThreadRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (ThreadRunImpl *ThreadRunImpl) ModifyRun(
	ThreadRunID string,
	updateRequest *openAiType.UpdateThreadRunRequest,
) (
	*openAiType.ThreadRunObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/ThreadRuns/%s", ThreadRunID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ThreadRunImpl.ApiKey))
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

	result := &openAiType.ThreadRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (ThreadRunImpl *ThreadRunImpl) GetRunList(
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.ThreadRunObject],
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
		fmt.Sprintf("https://api.openai.com/v1/ThreadRuns%s", queryString),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ThreadRunImpl.ApiKey))
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

	result := &openAiType.ListResponse[openAiType.ThreadRunObject]{}
	json.Unmarshal(body, result)

	return result, nil
}

func (ThreadRunImpl *ThreadRunImpl) GetRun(ThreadRunID string) (
	*openAiType.ThreadRunObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/ThreadRuns/%s", ThreadRunID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ThreadRunImpl.ApiKey))
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

	result := &openAiType.ThreadRunObject{}
	json.Unmarshal(body, result)

	return result, nil
}
