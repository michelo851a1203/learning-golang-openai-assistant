package openAiMessages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiError"
	"testf/openAiError/openAiErrorCode"
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
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CreateMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createMessagesClient := &http.Client{}

	response, err := createMessagesClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CreateMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CreateMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	result := &openAiType.OpenAiMessagesObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
				OpenStatusCode: openAiErrorCode.MessageCreateErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "CreateMessages",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageCreateOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "CreateMessages",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesError{},
		}
	}

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
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageModifyRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "ModifyMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
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
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageModifySendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "ModifyMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageModifyReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "ModifyMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	result := &openAiType.OpenAiMessagesObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageModifyResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "ModifyMessages",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
				OpenStatusCode: openAiErrorCode.MessageModifyErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "ModifyMessages",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageModifyOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "ModifyMessages",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesError{},
		}
	}

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
		queryString = listRequest.ToQueryString()
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
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetMessagesList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	MessagesClient := &http.Client{}

	response, err := MessagesClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessagesList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessagesList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.OpenAiMessagesObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessagesList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	if result.FirstID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
				OpenStatusCode: openAiErrorCode.MessageGetListErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetMessagesList",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetListOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetMessagesList",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesError{},
		}
	}

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
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetMessage",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessagesImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailMessagesClient := &http.Client{}

	response, err := detailMessagesClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessage",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}

	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessage",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	result := &openAiType.OpenAiMessagesObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessage",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
				OpenStatusCode: openAiErrorCode.MessageGetErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetMessage",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesError]{
			OpenStatusCode: openAiErrorCode.MessageGetOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetMessage",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesError{},
		}
	}

	return result, nil
}
