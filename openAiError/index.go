package openAiError

import "testf/openAiError/openAiErrorCode"

type OpenAiErrorGenericType interface {
	ThreadsError | AssistantError | MessagesError | RunAssistantError
}

type OpenAiError[T OpenAiErrorGenericType] struct {
	OpenStatusCode openAiErrorCode.OpenAiErrorCode
	Message        string
	Details        T
}

func (openAiError *OpenAiError[T OpenAiErrorGenericType]) Error()
