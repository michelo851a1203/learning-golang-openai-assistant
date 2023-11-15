package openAiType

import (
	"fmt"
	"net/url"
	"testf/openAiType/listOrder"
	"testf/openAiType/openAiListObject"
	"testf/openAiType/openAiTool"
)

type FunctionCallingParameters struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required,omitempty"`
}

type FunctionCallingObject struct {
	Description string                       `json:"description"`
	Name        string                       `json:"name"`
	Parameters  []*FunctionCallingParameters `json:"parameters"`
}

type OpenAiTool struct {
	Type     openAiTool.Tool        `json:"type"`
	Function *FunctionCallingObject `json:"function,omitempty"`
}

func (tool *OpenAiTool) checkValid() bool {
	switch tool.Type {
	case openAiTool.CodeInterpreter:
		return tool.Function == nil
	case openAiTool.Retrieval:
		return tool.Function == nil
	case openAiTool.FunctionCalling:
		return tool.Function != nil
	}
	return false
}

type OpenAiFileObject struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type OpenAiMetaData map[string]string

type DeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type ListResult interface {
	AssistantObject | AssistantFileObject | OpenAiMessagesObject | OpenAiMessagesFileObject | OpenAiRunObject | OpenAiRunStepObject
}

type ListResponse[T ListResult] struct {
	Object  openAiListObject.ObjectStatus `json:"object"`
	Data    []*T                          `json:"data"`
	FirstID string                        `json:"first_id"`
	LastID  string                        `json:"last_id"`
	HasMore bool                          `json:"has_more"`
}

type QueryListRequest struct {
	Limit  int                 `json:"limit,omitempty"`
	Order  listOrder.ListOrder `json:"order,omitempty"`
	After  string              `json:"after,omitempty"`
	Before string              `json:"before,omitempty"`
}

func (queryListRequest *QueryListRequest) ToQueryString() string {
	queryValue := url.Values{}

	if queryListRequest.Limit > 0 {
		queryValue.Set("limit", fmt.Sprintf("%d", queryListRequest.Limit))
	}

	if queryListRequest.Order != "" {
		queryValue.Set("order", string(queryListRequest.Order))
	}

	if queryListRequest.After != "" {
		queryValue.Set("after", queryListRequest.After)
	}

	if queryListRequest.Before != "" {
		queryValue.Set("before", queryListRequest.Before)
	}
	queryRawString := queryValue.Encode()
	if queryRawString == "" {
		return ""
	}
	return fmt.Sprintf("?%s", queryRawString)
}
