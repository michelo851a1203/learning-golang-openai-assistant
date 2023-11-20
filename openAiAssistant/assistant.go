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

type AssistantImpl struct {
	ApiKey string
}

func (assistantImpl *AssistantImpl) CreateAssistant(
	createRequest *openAiType.CreateAssistantRequest,
) (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/assistants",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CreateAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createAssistantClient := &http.Client{}

	response, err := createAssistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CreateAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CreateAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	result := &openAiType.AssistantObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
				OpenStatusCode: openAiErrorCode.AssistantCreateErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "CreateAssistant",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantCreateOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "CreateAssistant",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantError{},
		}
	}

	return result, nil
}

func (assistantImpl *AssistantImpl) ModifyAssistant(
	assistantID string,
	updateRequest *openAiType.UpdateAssistantRequest,
) (
	*openAiType.AssistantObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifyRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "ModifyAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifyNewRequestError,
			Message:        "NewRequest Error",
			Method:         "ModifyAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateAssistantClient := &http.Client{}

	response, err := updateAssistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifySendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "ModifyAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifyReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "ModifyAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	result := &openAiType.AssistantObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifyResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "ModifyAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
				OpenStatusCode: openAiErrorCode.AssistantModifyErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "ModifyAssistant",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantModifyOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "ModifyAssistant",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantError{},
		}
	}

	return result, nil
}

func (assistantImpl *AssistantImpl) DeleteAssistant(assistantID string) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteNewRequestError,
			Message:        "NewRequest Error",
			Method:         "DeleteAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	deleteAssistantClient := &http.Client{}

	response, err := deleteAssistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "DeleteAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "DeleteAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	result := &openAiType.DeleteResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "DeleteAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
				OpenStatusCode: openAiErrorCode.AssistantDeleteErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "DeleteAssistant",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantDeleteOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "DeleteAssistant",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantError{},
		}
	}

	return result, nil
}

func (assistantImpl *AssistantImpl) GetAssistantList(
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.AssistantObject],
	error,
) {
	queryString := ""
	if listRequest != nil {
		queryString = listRequest.ToQueryString()
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants%s", queryString),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetAssistantList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	assistantClient := &http.Client{}

	response, err := assistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetAssistantList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetAssistantList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.AssistantObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistantList",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	if result.FirstID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
				OpenStatusCode: openAiErrorCode.AssistantGetListErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetAssistantList",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetListOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetAssistantList",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantError{},
		}
	}

	return result, nil
}

func (assistantImpl *AssistantImpl) GetAssistant(assistantID string) (
	*openAiType.AssistantObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/assistants/%s", assistantID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", assistantImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailAssistantClient := &http.Client{}

	response, err := detailAssistantClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	result := &openAiType.AssistantObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetAssistant",
			RawError:       err.Error(),
			Details:        &openAiError.AssistantError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
				OpenStatusCode: openAiErrorCode.AssistantGetErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetAssistant",
				RawError:       err.Error(),
				Details:        &openAiError.AssistantError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.AssistantError]{
			OpenStatusCode: openAiErrorCode.AssistantGetOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetAssistant",
			RawError:       errorResult.String(),
			Details:        &openAiError.AssistantError{},
		}
	}

	return result, nil
}
