package openAiError

import (
	"fmt"
	"testf/openAiError/openAiErrorCode"
)

type OpenAiErrorGenericType interface {
	ThreadsError | AssistantError | MessagesError | RunAssistantError | FileError | MessagesFile
}

type OpenAiError[T OpenAiErrorGenericType] struct {
	OpenStatusCode openAiErrorCode.OpenAiErrorCode
	Message        string
	Method         string
	RawError       string
	Details        *T
}

func (e *OpenAiError[T]) Error() string {
	return fmt.Sprintf(
		"\n  ========== \n  [openAiError]:\n  code : %v\n  message : %s\n  method : %s\n RawError : %s\n  Detail : %v\n  ==========",
		e.OpenStatusCode,
		e.Message,
		e.Method,
		e.RawError,
		e.Details,
	)
}
