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
			OpenStatusCode: openAiErrorCode.AssistantFileCreateRequestJSONError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileCreateNewRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileCreateSendHTTPRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileCreateReadResponseBodyError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
				OpenStatusCode: openAiErrorCode.AssistantFileCreateErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "CreateAssistantFile",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantFileCreateOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "CreateAssistantFile",
			RawError:       errorResult.String(),
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetNewRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetSendHTTPRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetReadResponseBodyError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
				OpenStatusCode: openAiErrorCode.AssistantFileGetErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetAssistantFile",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantFileGetOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetAssistantFile",
			RawError:       errorResult.String(),
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
			OpenStatusCode: openAiErrorCode.AssistantFileDeleteNewRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileDeleteSendHTTPRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileDeleteReadResponseBodyError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileDeleteResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "DeleteAssistantFile",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
				OpenStatusCode: openAiErrorCode.AssistantFileDeleteErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "DeleteAssistantFile",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantFileDeleteOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "DeleteAssistantFile",
			RawError:       errorResult.String(),
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetListNewRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetListSendHTTPRequestError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetListReadResponseBodyError,
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
			OpenStatusCode: openAiErrorCode.AssistantFileGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistantFileList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	if result.FirstID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
				OpenStatusCode: openAiErrorCode.AssistantFileGetListErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetAssistantFileList",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantFileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantFileError]{
			OpenStatusCode: openAiErrorCode.AssistantFileGetListOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetAssistantFileList",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantFileError{},
		}
	}

	return result, nil
}
