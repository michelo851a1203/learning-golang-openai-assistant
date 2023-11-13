package openAiType

import (
	"errors"
	"testf/openAiType/listOrder"
	"testf/openAiType/openAiListObject"
	"testf/openAiType/openAiTool"
)

type FunctionCallingObject struct {
}

type OpenAiTool struct {
	Type     openAiTool.Tool        `json:"type"`
	Function *FunctionCallingObject `json:"function,omitempty"`
}

func (OpenAiTool *OpenAiTool) ValidateFunctionCalling() error {
	if OpenAiTool.Type != openAiTool.FunctionCalling && OpenAiTool == nil {
		return errors.New("OpenAiTool.Type is not function calling")
	}
	return nil
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
	AssistantObject | AssistantFileObject | OpenAiMessagesObject
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
