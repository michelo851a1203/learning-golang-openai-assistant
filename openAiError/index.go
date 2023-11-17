package openAiError

import (
	"fmt"
	"testf/openAiError/openAiErrorCode"
)

type OpenAiErrorGenericType interface {
	ThreadsError | AssistantError | AssistantFileError | MessagesError | RunAssistantError | FileError | MessagesFileError
}

type OpenAiError[T OpenAiErrorGenericType] struct {
	OpenStatusCode openAiErrorCode.OpenAiErrorCode
	Message        string
	Method         string
	RawError       string
	Details        *T
	NativeApiError *OpenAiNativeApiError
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

type OpenAiNativeApiError struct {
	Message string  `json:"message"`
	Type    string  `json:"type"`
	Param   *string `json:"param"`
	Code    string  `json:"code"`
}

func (openAiNativeApiError *OpenAiNativeApiError) String() string {
	return fmt.Sprintf(`
{
  "error": {
    "message": %s,
    "type": %s,
    "param": %s,
    "code":%s 
  }
}
	`,
		openAiNativeApiError.Message,
		openAiNativeApiError.Type,
		*openAiNativeApiError.Param,
		openAiNativeApiError.Code,
	)
}
