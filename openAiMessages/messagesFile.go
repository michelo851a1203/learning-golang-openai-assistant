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
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}

	result := &openAiType.ListResponse[openAiType.OpenAiMessagesFileObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessageFileList",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
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
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessageFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailMessageFileClient := &http.Client{}

	response, err := detailMessageFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}

	result := &openAiType.OpenAiMessagesFileObject{}
	err = json.Unmarshal(body, result)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.MessagesFile]{
			OpenStatusCode: openAiErrorCode.GetMessageFileResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetMessageFile",
			RawError:       err.Error(),
			Details:        &openAiError.MessagesFile{},
		}
	}

	return result, nil
}
