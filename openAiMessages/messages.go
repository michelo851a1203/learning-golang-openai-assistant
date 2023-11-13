package openAiMessages

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

type MessagesImpl struct {
	ApiKey string
}

func (MessagesImpl *MessagesImpl) CreateMessages(
	threadID string,
	createRequest *openAiType.CreateMessagesRequest,
) (
	*openAiType.OpenAiMessagesObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createMessagesClient := &http.Client{}

	response, err := createMessagesClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiMessagesObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (MessagesImpl *MessagesImpl) ModifyMessages(
	threadID string,
	messagesID string,
	updateRequest *openAiType.UpdateMessagesRequest,
) (
	*openAiType.OpenAiMessagesObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/messages/%s",
			threadID,
			messagesID,
		),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateMessagesClient := &http.Client{}

	response, err := updateMessagesClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiMessagesObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (MessagesImpl *MessagesImpl) GetMessagesList(
	threadID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiMessagesObject],
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
			"https://api.openai.com/v1/threads/%s/messages%s",
			threadID,
			queryString,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	MessagesClient := &http.Client{}

	response, err := MessagesClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListResponse[openAiType.OpenAiMessagesObject]{}
	json.Unmarshal(body, result)

	return result, nil
}

func (MessagesImpl *MessagesImpl) GetMessages(
	threadID string,
	MessagesID string,
) (
	*openAiType.OpenAiMessagesObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages%s", threadID, MessagesID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailMessagesClient := &http.Client{}

	response, err := detailMessagesClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiMessagesObject{}
	json.Unmarshal(body, result)

	return result, nil
}
