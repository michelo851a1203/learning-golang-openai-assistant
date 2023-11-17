package openAiThreadRun

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

type ThreadRunImpl struct {
	ApiKey string
}

func (threadRunImpl *ThreadRunImpl) CreateRun(
	threadID string,
	createRequest *openAiType.CreateThreadRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(createRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs",
			threadID,
		),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunCreateNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CreateRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunCreateSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CreateRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunCreateReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CreateRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunCreateResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CreateRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) ModifyRun(
	threadID,
	runID string,
	updateRequest *openAiType.UpdateThreadRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunModifyRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "ModifyRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s",
			threadID,
			runID,
		),

		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunModifyNewRequestError,
			Message:        "NewRequest Error",
			Method:         "ModifyRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	updateThreadRunClient := &http.Client{}

	response, err := updateThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunModifySendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "ModifyRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunModifyReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "ModifyRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunModifyResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "ModifyRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) GetRunList(
	threadID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiRunObject],
	error,
) {
	queryString := ""
	if listRequest != nil {
		queryString = listRequest.ToQueryString()
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs%s",
			threadID,
			queryString,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetRunList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	ThreadRunClient := &http.Client{}

	response, err := ThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetRunList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetRunList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.OpenAiRunObject]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetRunList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) GetRun(threadID, runID string) (
	*openAiType.OpenAiRunObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,

		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s",
			threadID,
			runID,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailThreadRunClient := &http.Client{}

	response, err := detailThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.RunGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) SubmitToolOutputToRun(
	threadID,
	runID string,
	toolOutputRequest *openAiType.SubmitOutputsAndRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(toolOutputRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.SubmitToolOutputToRunCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "SubmitToolOutputToRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/submit_tool_outputs",
			threadID,
			runID,
		),
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.SubmitToolOutputToRunGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "SubmitToolOutputToRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.SubmitToolOutputToRunGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "SubmitToolOutputToRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.SubmitToolOutputToRunGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "SubmitToolOutputToRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.SubmitToolOutputToRunGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "SubmitToolOutputToRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) CancelRun(
	threadID,
	runID string,
) (
	*openAiType.OpenAiRunObject,
	error,
) {

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/cancel",
			threadID,
			runID,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CancelRunNewRequestError,
			Message:        "NewRequest Error",
			Method:         "CancelRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CancelRunSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "CancelRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CancelRunReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "CancelRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CancelRunResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "CancelRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) CreateThreadAndRun(
	threadAndRunRequest *openAiType.ThreadAndRunRequest,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	requestInfo, err := json.Marshal(threadAndRunRequest)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CreateThreadAndRunCreateRequestJSONError,
			Message:        "Request Marshal JSON Error",
			Method:         "CreateThreadAndRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/threads/runs",
		bytes.NewBuffer(requestInfo),
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CreateThreadAndRunGetNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CreateThreadAndRunGetSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CreateThreadAndRunGetReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}

	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.CreateThreadAndRunGetResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetRun",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil

}

func (threadRunImpl *ThreadRunImpl) GetRunStep(
	threadID, runID, stepID string,
) (
	*openAiType.OpenAiRunObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/steps/%s",
			threadID,
			runID,
			stepID,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetRunStep",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetRunStep",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetRunStep",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.OpenAiRunObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetRunStep",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}

func (threadRunImpl *ThreadRunImpl) GetRunStepList(
	threadID,
	runID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiRunStepObject],
	error,
) {
	queryString := ""
	if listRequest != nil {
		queryString = listRequest.ToQueryString()
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/runs/%s/steps%s",
			threadID,
			runID,
			queryString,
		),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetRunStepList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", threadRunImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	createThreadRunClient := &http.Client{}

	response, err := createThreadRunClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetRunStepList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetRunStepList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	result := &openAiType.ListResponse[openAiType.OpenAiRunStepObject]{}
	fmt.Println(string(body))
	err = json.Unmarshal(body, result)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.RunAssistantError]{
			OpenStatusCode: openAiErrorCode.GetRunStepListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetRunStepList",
			RawError:       err.Error(),
			Details:        &openAiError.RunAssistantError{},
		}
	}

	return result, nil
}
