package openaiThreads

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

type ThreadsImpl struct {
	ApiKey string
}

func (threadsImpl *ThreadsImpl) CreateThread(
	createRequest *openAiType.ThreadCreateRequest,
) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}

	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/threads",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadCreateNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CreateThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadsClient := &http.Client{}

	response, err := createThreadsClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadCreateSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CreateThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadCreateReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CreateThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	result := &openAiType.OpenAiThreadObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	return result, nil
}

func (threadsImpl *ThreadsImpl) ModifyThread(
	threadsID string,
	updateRequest *openAiType.UpdateThreadsRequest,
) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadModifyRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "ModifyThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadsID),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadModifyNewRequestError,
			Message:        "NewRequest Error",
			Method:         "ModifyThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateThreadsClient := &http.Client{}

	response, err := updateThreadsClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadModifySendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "ModifyThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadModifyReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "ModifyThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	result := &openAiType.OpenAiThreadObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadModifyResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "ModifyThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	return result, nil
}

func (threadsImpl *ThreadsImpl) DeleteThread(threadID string) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadDeleteNewRequestError,
			Message:        "NewRequest Error",
			Method:         "DeleteThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	deleteThreadsClient := &http.Client{}

	response, err := deleteThreadsClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadDeleteSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "DeleteThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadDeleteReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "DeleteThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	result := &openAiType.DeleteResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadDeleteResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "DeleteThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	return result, nil
}

func (threadsImpl *ThreadsImpl) GetThread(threadID string) (
	*openAiType.OpenAiThreadObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadsImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailThreadsClient := &http.Client{}

	response, err := detailThreadsClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	result := &openAiType.OpenAiThreadObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.ThreadsError]{
			OpenStatusCode: openAiErrorCode.ThreadGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetThread",
			RawError:       err.Error(),
			Details:        &openAiError.ThreadsError{},
		}
	}

	return result, nil
}
