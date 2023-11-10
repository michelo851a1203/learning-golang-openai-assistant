package openAiType

import (
	"errors"
	"testf/openAiType/openAiToolType"
)

type FunctionCallingObject struct {
}

type OpenAiTool struct {
	Type     openAiToolType.OpenAiToolType `json:"type"`
	Function *FunctionCallingObject        `json:"function,omitempty"`
}

func (OpenAiTool *OpenAiTool) ValidateFunctionCalling() error {
	if OpenAiTool.Type != openAiToolType.FunctionCalling && OpenAiTool == nil {
		return errors.New("OpenAiTool.Type is not function calling")
	}
	return nil
}

type OpenAiAttachFile struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type OpenAiMetaData map[string]string
