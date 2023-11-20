package openAiMessages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiError"
	"testf/openAiError/openAiErrorCode"
	"testf/openAiType"
)

type MessageFileImpl struct {
	ApiKey string
}

func (MessageFileImpl *MessageFileImpl) GetMessageFileList(
	threadID string,
	messageID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiMessagesFileObject],
	error,
) {
	queryString := ""
	if listRequest != nil {
		queryString = listRequest.ToQueryString()
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/messages/%s/files%s",
			threadID,
			messageID,
			queryString,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessageFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	MessageFileClient := &http.Client{}

	response, err := MessageFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.OpenAiMessagesFileObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}

	}

	if result.FirstID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
				OpenStatusCode: openAiErrorCode.GetMessageFileListGetErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetMessageFileList",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListGetOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetMessageFileList",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	return result, nil
}

func (MessageFileImpl *MessageFileImpl) GetMessageFile(
	threadID string,
	messageID string,
	fileID string,
) (
	*openAiType.OpenAiMessagesFileObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/messages/%s/files/%s",
			threadID,
			messageID,
			fileID,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessageFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailMessageFileClient := &http.Client{}

	response, err := detailMessageFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	result := &openAiType.OpenAiMessagesFileObject{}
	err = json.Unmarshal(body, result)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
				OpenStatusCode: openAiErrorCode.GetMessageFileGetErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetMessageFile",
				RawError:       err.Error(),
				Details:        &openAiError.MessagesFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.MessagesFileError]{
			OpenStatusCode: openAiErrorCode.GetMessageFileGetOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetMessageFile",
			RawError:       errorResult.String(),
			Details:        &openAiError.MessagesFileError{},
		}
	}

	return result, nil
}
