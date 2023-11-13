package openAiType

import "testf/openAiType/openAiRole"

type CreateMessagesRequest struct {
	Role    openAiRole.Role `json:"role"`
	Content string          `json:"content"`
}
