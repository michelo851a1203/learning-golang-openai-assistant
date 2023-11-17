package openAiAssistant

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
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s/files", assistantID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileClient := &http.Client{}

	response, err := assistantFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}

	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	result := &openAiType.AssistantFileObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

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
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileClient := &http.Client{}

	response, err := assistantFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	result := &openAiType.AssistantFileObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

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
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteNewRequestError,
			Message:        "NewRequest Error",
			Method:         "DeleteAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantFileDeleteClient := &http.Client{}

	response, err := assistantFileDeleteClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "DeleteAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "DeleteAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	result := &openAiType.DeleteResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "DeleteAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

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
		queryString = listRequest.ToQueryString()
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
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetAssistantFileList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantClient := &http.Client{}

	response, err := assistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetAssistantFileList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetAssistantFileList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.AssistantFileObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistantFileList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	return result, nil
}
